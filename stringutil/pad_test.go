package stringutil

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"testing/quick"
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

func TestGeneratedPaddingLeft(t *testing.T) {
	fn := func(s string, i uint8) bool {
		r := PaddingLeft(s, "*", int(i))
		if len(s) > int(i) {
			return len(r) > int(i)
		}
		return len(r) == int(i)
	}
	if err := quick.Check(fn, nil); err != nil {
		t.Error(err)
	}
}

func TestGeneratedPaddingRight(t *testing.T) {
	fn := func(s string, i uint8) bool {
		r := PaddingRight(s, "*", int(i))
		if len(s) > int(i) {
			return len(r) > int(i)
		}
		return len(r) == int(i)
	}
	if err := quick.Check(fn, nil); err != nil {
		t.Error(err)
	}
}