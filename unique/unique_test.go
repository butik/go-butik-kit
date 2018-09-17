package unique

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUniqueStrings(t *testing.T) {
	assert.Equal(t, []string{"a", "b", "c"}, Strings([]string{"a", "a", "b", "c"}))
	assert.Equal(t, []string{"a", "b", "c"}, Strings([]string{"a", "b", "c"}))
	assert.Equal(t, []string{}, Strings([]string{}))
}

func ExampleStrings() {
	fmt.Println(Strings([]string{"a", "a", "b", "c"}))
	// Output: [a b c]
}
