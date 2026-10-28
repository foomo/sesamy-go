package mpv2

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

// SHA256Hash is a hex-encoded SHA-256 hash, as Google expects for user-provided
// data such as email addresses and phone numbers.
type SHA256Hash string

// NewSHA256Hash returns the SHA256Hash of s after trimming whitespace and
// converting it to lower case.
func NewSHA256Hash(s string) SHA256Hash {
	h := sha256.New()
	h.Write([]byte(strings.TrimSpace(strings.ToLower(s))))

	return SHA256Hash(hex.EncodeToString(h.Sum(nil)))
}

// String returns the hex-encoded hash.
func (s SHA256Hash) String() string {
	return string(s)
}
