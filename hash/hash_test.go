package hash

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHash_TestMD4(t *testing.T) {
	str := "some string"
	hash, err := GetMD5Hash(str)
	assert.NoError(t, err)
	assert.Equal(t, "5ac749fbeec93607fc28d666be85e73a", hash)
}
