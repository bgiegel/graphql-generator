package graphql

import (
	"reflect"
	"fmt"
	"strings"
)

type GraphQLType struct {
	Name   string
	Fields []GraphQLField
}

type GraphQLField struct {
	Name     string
	Type     string
	Required bool
}

func GenerateType(s interface{}) string {
	t := reflect.TypeOf(s)

	gqlType := GraphQLType{}
	gqlType.Name = t.Name()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		gqlType.Fields = append(gqlType.Fields, buildField(field))
	}

	var typeStr []string
	typeStr = append(typeStr, fmt.Sprintf("type %s {", gqlType.Name))
	for _, field := range gqlType.Fields {
		typeStr = append(typeStr, fmt.Sprintf("\t%s: %s", field.Name, field.Type))
	}
	typeStr = append(typeStr, "}")

	result := strings.Join(typeStr, "\n")

	return result
}
func buildField(field reflect.StructField) GraphQLField {
	gqlField := GraphQLField{}
	if field.Tag != "" {
		gqlField.Name = field.Tag.Get("graphql")
	} else {
		gqlField.Name = field.Name
	}

	typeStr := field.Type.Name()
	switch typeStr {
	case "string":
		gqlField.Type = "String"
	case "int":
		gqlField.Type = "Int"
	case "bool":
		gqlField.Type = "Boolean"
	default:
		gqlField.Type = "String"
	}
	return gqlField
}
