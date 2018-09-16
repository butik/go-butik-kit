package stringutil

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
)

func TestPaddingLeft(t *testing.T) {
	assert.Equal(t, "1", PaddingLeft("1", "0", 0))
	assert.Equal(t, "1", PaddingLeft("1", "0", 1))
	assert.Equal(t, "01", PaddingLeft("1", "0", 2))
	assert.Equal(t, "**1", PaddingLeft("1", "**", 3))
	assert.Equal(t, "*1", PaddingLeft("1", "**", 2))
}

func ExamplePaddingLeft() {
	fmt.Println(PaddingLeft("1", "0", 0))
	fmt.Println(PaddingLeft("1", "0", 1))
	fmt.Println(PaddingLeft("1", "0", 2))
	fmt.Println(PaddingLeft("1", "**", 3))
	fmt.Println(PaddingLeft("1", "**", 2))
	// Output: 1
	// 1
	// 01
	// **1
	// *1
}

func TestPaddingRight(t *testing.T) {
	assert.Equal(t, "1", PaddingRight("1", "0", 0))
	assert.Equal(t, "1", PaddingRight("1", "0", 1))
	assert.Equal(t, "10", PaddingRight("1", "0", 2))
	assert.Equal(t, "1**", PaddingRight("1", "**", 3))
	assert.Equal(t, "1*", PaddingRight("1", "**", 2))
}

func ExamplePaddingRight() {
	fmt.Println(PaddingRight("1", "0", 0))
	fmt.Println(PaddingRight("1", "0", 1))
	fmt.Println(PaddingRight("1", "0", 2))
	fmt.Println(PaddingRight("1", "**", 3))
	fmt.Println(PaddingRight("1", "**", 2))
	// Output: 1
	// 1
	// 10
	// 1**
	// 1*
}
