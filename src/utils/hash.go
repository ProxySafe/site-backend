package utils

import (
	"crypto/sha1"
	"encoding/base64"
)

const (
	salt = ")skjf$djkvn^#@dkjvndf%&878kj"
)

func GetPasswordHash(password string) string {
	hasher := sha1.New()
	hasher.Write([]byte(password))
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	return sha + salt
}
