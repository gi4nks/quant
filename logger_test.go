package quant

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestIfLoggerIsInstatiated(t *testing.T) {
	var l = NewLogger()
	assert.True(t, l != nil, "Object has not been instantiated")
}

func TestIfLoggerIsNotInitialized(t *testing.T) {
	var l = NewLogger()
	l.Print("test-1", "test-2", "test-3")
	assert.True(t, l != nil, "Object has not been instantiated")
}