package stringutil

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseStrings(t *testing.T) {
	strings := []string{"foo", "bar", "bar2"}
	assert.Equal(t, ReverseStrings(strings), []string{"bar2", "bar", "foo"})
	assert.Equal(t, strings, []string{"foo", "bar", "bar2"})
}
