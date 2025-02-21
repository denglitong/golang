package testify_demo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAssert(t *testing.T) {
	assert.Equal(t, 123, 123, "should be equal")
	assert.NotEqualf(t, 123, 456, "should be not equal")
	assert.Nil(t, nil, "should be nil")
	assert.NotNil(t, "", "should be not nil")

	obj := struct {
		Value string
	}{
		"value",
	}

	if assert.NotNil(t, obj) {
		assert.NotEmptyf(t, obj.Value, "value should be not empty")
	}
}
