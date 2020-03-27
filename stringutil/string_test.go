package stringutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringRef(t *testing.T) {
	textRef := StringRef("some text")
	assert.NotNil(t, textRef)
}
