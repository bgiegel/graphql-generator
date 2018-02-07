package graphql

import "testing"
import "github.com/stretchr/testify/assert"

type Person struct {
	Name   string `graphql:"name"`
	Age    int    `graphql:"age"`
	Female bool   `graphql:"isFemale"`
}

const expectedType =
`type Person {
	name: String
	age: Int
	isFemale: Boolean
}`

func TestGenerateType(t *testing.T) {
	actualType := GenerateType(Person{})
	assert.Equal(t, expectedType, actualType)
}
