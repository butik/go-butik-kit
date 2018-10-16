package stringutil

import (
	"math/rand"
	"reflect"
	"testing"
	"testing/quick"
	"unicode"

	"github.com/stretchr/testify/assert"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzабвгдежзийклмнопрстуфхцчшщьэюя")

type Words string

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func (Words) Generate(r *rand.Rand, size int) reflect.Value {
	str := Words(RandStringRunes(size))
	return reflect.ValueOf(str)
}

func TestWords_UCFirst(t *testing.T) {
	str := "this is some string"
	assert.Equal(t, "This is some string", UCFirst(str))
}

func TestGeneratedUCFirst(t *testing.T) {
	fn := func(s Words) bool {
		str := UCFirst(string(s))
		for _, v := range str {
			return unicode.IsUpper(v) && str[len(string(v)):] == string(s)[len(string(v)):]
		}

		return false
	}

	if err := quick.Check(fn, &quick.Config{}); err != nil {
		t.Error(err)
	}
}
