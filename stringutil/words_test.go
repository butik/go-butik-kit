package stringutil

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWords_UCFirst(t *testing.T) {
	str := "this is some string"
	assert.Equal(t, "This is some string", UCFirst(str))
}
