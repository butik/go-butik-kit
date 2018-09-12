package unique

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestUniqueStrings(t *testing.T) {
	assert.Equal(t, []string{"a", "b", "c"}, Strings([]string{"a", "a", "b", "c"}))
	assert.Equal(t, []string{"a", "b", "c"}, Strings([]string{"a", "b", "c"}))
	assert.Equal(t, []string{}, Strings([]string{}))
}
