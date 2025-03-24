package nanoid

import (
	"crypto/rand"
	"math/big"
)

const (
	// Sensible, URL safe default values.
	DefaultLength  = 22
	DefaultCharset = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz123456789-_"
)

// Nanoid generates a cryptographically secure user ID.
func Nanoid(length int, charset string) string {
	id := ""
	charsetLen := big.NewInt(int64(len(charset)))
	for range length {
		n, _ := rand.Int(rand.Reader, charsetLen)
		id += string(charset[n.Int64()])
	}
	return id
}

// Default generates a nanoid using the default values.
func Default() string {
	return Nanoid(DefaultLength, DefaultCharset)
}
