package dryvers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBrightness_Get(t *testing.T) {
	b := NewBrightness()
	if b.mode == noBacklight {
		t.Log("Brightness does not support current computer")
		return
	}

	value, err := b.Get()
	assert.NoError(t, err)
	assert.Greater(t, value, float64(0)) // the screen should really not be off!
	assert.LessOrEqual(t, value, float64(1))
}
