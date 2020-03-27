package stringutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseStrings(t *testing.T) {
	strings := []string{"foo", "bar", "bar2"}
	assert.Equal(t, ReverseStrings(strings), []string{"bar2", "bar", "foo"})
	assert.Equal(t, strings, []string{"foo", "bar", "bar2"})
}
