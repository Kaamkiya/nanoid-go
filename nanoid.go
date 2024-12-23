package nanoid

import (
	"math/big"
	"crypto/rand"
)

const DefaultLength = 22
const DefaultCharset = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz123456789-_"

func Nanoid(length int, charset string) string {
	id := ""
	charsetLen := big.NewInt(int64(len(charset)))
	for i := 0; i < length; i++ {
		n, _ := rand.Int(rand.Reader, charsetLen)
		id += string(charset[n.Int64()])
	}
	return id
}
