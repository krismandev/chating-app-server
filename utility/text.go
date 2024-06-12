package utility

import (
	"math/rand"
	"strings"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyz0123456789"

func RandString(n int) string {
	sb := strings.Builder{}
	sb.Grow(n)
	for i := 0; i < n; i++ {
		sb.WriteByte(letterBytes[rand.Intn(len(letterBytes))])
	}
	return sb.String()
}
