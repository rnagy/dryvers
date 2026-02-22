package dryvers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBattery_Get(t *testing.T) {
	b := NewBattery()

	value, err := b.Get()
	if err != nil {
		t.Log("Unable to get battery info", err)
		return
	}

	assert.GreaterOrEqual(t, value, float64(0))
	assert.LessOrEqual(t, value, float64(1))
}

func TestBattery_IsLow(t *testing.T) {
	b := NewBattery()

	value, err := b.Get()
	if err != nil {
		t.Log("Unable to get battery info", err)
		return
	}

	if value > 0.1 {
		low, err := b.IsLow()
		assert.NoError(t, err)
		assert.False(t, low)
	} else {
		low, err := b.IsLow()
		assert.NoError(t, err)
		assert.True(t, low)
	}
}
