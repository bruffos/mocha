package mock

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/vitorsalgado/mocha/matchers"
)

func TestDebug(t *testing.T) {
	mk := New()
	mk.Name = "test"
	result, err := Debug("equalTo", *mk, matchers.EqualTo("test")).Matches("test", matchers.Args{})

	assert.Nil(t, err)
	assert.True(t, result)
}

func TestDebugErr(t *testing.T) {
	mk := New()
	mk.Name = "test"
	result, err := Debug("err", *mk, matchers.Fn(
		func(v string, params matchers.Args) (bool, error) {
			return false, fmt.Errorf("failed")
		})).
		Matches("test", matchers.Args{})

	assert.NotNil(t, err)
	assert.False(t, result)
}

func TestDebugNotMatched(t *testing.T) {
	mk := New()
	mk.Name = "test"
	result, err := Debug("equalTo", *mk, matchers.EqualTo("test")).Matches("dev", matchers.Args{})

	assert.Nil(t, err)
	assert.False(t, result)
}