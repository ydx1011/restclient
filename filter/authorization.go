package filter

import (
	"crypto/rand"
	"encoding/base64"
)

func RandomId(length int) string {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}
