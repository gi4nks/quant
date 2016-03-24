package quant

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParrot(t *testing.T) {
	var p = NewVerboseParrot("parrot_test")
	p.Debug("ciao")
	p.Debug("come", "stai")

	assert.True(t, p != nil, "Object has not been instantiated")
}
