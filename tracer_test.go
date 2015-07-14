package quant

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestIfTracerIsInstatiated(t *testing.T) {
	var tracer = NewTrace("test_1")
	assert.True(t, tracer != nil, "Object has not been instantiated")
}

func TestIfTracerIsNotInitialized(t *testing.T) {
	var tracer = NewTrace("test_1")
	err := tracer.News("news message to be written here: " + "working with message concatenation also")
	assert.True(t, err != nil, "Tracer has not been initialized")
}

func TestIfTracerIsInitialized(t *testing.T) {
	var tracer = NewTrace("test_1")
	tracer.Full()
	err := tracer.News("news message to be written here: " + "working with message concatenation also")
	assert.True(t, err == nil, "Tracer has not been initialized")
}