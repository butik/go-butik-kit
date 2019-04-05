package intutil

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseInts(t *testing.T) {
	ints := []int{5, 4, 2, 4, 5, 1, 3}
	assert.Equal(t, ReverseInts(ints), []int{3, 1, 5, 4, 2, 4, 5})
	assert.Equal(t, ints, []int{5, 4, 2, 4, 5, 1, 3})
}
