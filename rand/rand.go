package rand

import (
	"math/rand"
	"time"
)

type Rand struct {
	*rand.Rand
}

func NewRand() *Rand {
	source := rand.NewSource(time.Now().UnixNano())
	return &Rand{rand.New(source)}
}

func (r *Rand) StringFromEnglishAndNumbers(n int) string {
	return r.RandStringFromAlphabet(n, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
}

func (r *Rand) StringFromNumbers(n int) string {
	return r.RandStringFromAlphabet(n, "1234567890")
}

func (r *Rand) RandStringFromAlphabet(n int, letters string) string {
	buf := make([]rune, n)
	runes := []rune(letters)

	for i := range buf {
		buf[i] = runes[r.Intn(len(runes))]
	}

	return string(buf)
}
