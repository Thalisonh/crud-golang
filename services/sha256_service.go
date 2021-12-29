package services

import (
	"crypto/sha256"
	"fmt"
)

func Sha256Encoder(a string) string {
	str := sha256.Sum256([]byte(a))

	return fmt.Sprintf("%x", str)
}
