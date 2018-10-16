package hash

import (
	"crypto/md5" // #nosec
	"encoding/hex"
)

// GetMD5Hash return md5 hash of the string
func GetMD5Hash(text string) (string, error) {
	hasher := md5.New() // #nosec
	if _, err := hasher.Write([]byte(text)); err != nil {
		return "", err
	} // #nosec
	return hex.EncodeToString(hasher.Sum(nil)), nil
}
