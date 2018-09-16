package stringutil

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestPaddingLeft(t *testing.T) {
	assert.Equal(t, "1", PaddingLeft("1", "0", 0))
	assert.Equal(t, "1", PaddingLeft("1", "0", 1))
	assert.Equal(t, "01", PaddingLeft("1", "0", 2))
	assert.Equal(t, "**1", PaddingLeft("1", "**", 3))
	assert.Equal(t, "*1", PaddingLeft("1", "**", 2))
}

func TestPaddingRight(t *testing.T) {
	assert.Equal(t, "1", PaddingRight("1", "0", 0))
	assert.Equal(t, "1", PaddingRight("1", "0", 1))
	assert.Equal(t, "10", PaddingRight("1", "0", 2))
	assert.Equal(t, "1**", PaddingRight("1", "**", 3))
	assert.Equal(t, "1*", PaddingRight("1", "**", 2))
}
