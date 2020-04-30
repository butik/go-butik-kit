package sliceutils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type Example struct {
	Value    string
	Priority int
}

func TestMax(t *testing.T) {
	slice := []Example{
		{
			Value:    "value 1",
			Priority: 1,
		},
		{
			Value:    "value 2",
			Priority: 8,
		},
		{
			Value:    "value 1",
			Priority: 10,
		},
	}
	greater := func(i int, j int) bool {
		return slice[i].Priority < slice[j].Priority
	}
	i := Max(len(slice), greater)
	require.Equal(t, 2, i)

	var emptySlice []Example
	j := Max(len(emptySlice), greater)
	require.GreaterOrEqual(t, len(emptySlice), j)
}

func TestMin(t *testing.T) {
	slice := []Example{
		{
			Value:    "value 2",
			Priority: 8,
		},
		{
			Value:    "value 1",
			Priority: 1,
		},
		{
			Value:    "value 1",
			Priority: 10,
		},
	}
	less := func(i int, j int) bool {
		return slice[i].Priority > slice[j].Priority
	}
	i := Min(len(slice), less)
	require.Equal(t, 1, i)

	var emptySlice []Example
	j := Max(len(emptySlice), less)
	require.GreaterOrEqual(t, len(emptySlice), j)
}
