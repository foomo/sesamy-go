package mpv2

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

type SHA256Hash string

func (s SHA256Hash) String() string {
	return string(s)
}

func NewSHA256Hash(s string) SHA256Hash {
	h := sha256.New()
	h.Write([]byte(strings.TrimSpace(strings.ToLower(s))))
	return SHA256Hash(hex.EncodeToString(h.Sum(nil)))
}
