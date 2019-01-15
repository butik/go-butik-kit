package intutil

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMin(t *testing.T) {
	minInt := 0
	maxInt := 1
	assert.Equal(t, Min(minInt, maxInt), minInt)
}

func TestMax(t *testing.T) {
	minInt := 0
	maxInt := 1
	assert.Equal(t, Max(minInt, maxInt), maxInt)
}
