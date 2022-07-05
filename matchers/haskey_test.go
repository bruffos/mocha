package matchers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasKey(t *testing.T) {
	m := map[string]any{
		"name": "someone",
		"age":  34,
		"address": map[string]any{
			"street": "very nice place",
		},
		"options": []string{},
		"active":  false,
		"none":    nil,
		"zero":    0,
	}

	result, err := HasKey[any]("name").Matches(m, emptyArgs())
	assert.True(t, result)
	assert.Nil(t, err)

	result, err = HasKey[any]("age").Matches(m, emptyArgs())
	assert.True(t, result)
	assert.Nil(t, err)

	result, err = HasKey[any]("address").Matches(m, emptyArgs())
	assert.True(t, result)
	assert.Nil(t, err)

	result, err = HasKey[any]("active").Matches(m, emptyArgs())
	assert.True(t, result)
	assert.Nil(t, err)

	result, err = HasKey[any]("zero").Matches(m, emptyArgs())
	assert.True(t, result)
	assert.Nil(t, err)

	result, err = HasKey[any]("options").Matches(m, emptyArgs())
	assert.True(t, result)
	assert.Nil(t, err)

	result, err = HasKey[any]("none").Matches(m, emptyArgs())
	assert.False(t, result)
	assert.Nil(t, err)
}