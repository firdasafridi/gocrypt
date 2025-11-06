package gocrypt

import (
	"crypto/rc4"
	"encoding/hex"

	"github.com/pkg/errors"
)

// RC4Opt is structure of RC4 option
type RC4Opt struct {
	secret []byte
}

// NewRC4Opt is function to create new configuration of RC4 algorithm option
// the secret is used directly as bytes (not hex-encoded)
func NewRC4Opt(secret string) (*RC4Opt, error) {
	return &RC4Opt{
		secret: []byte(secret),
	}, nil
}

// Encrypt encrypts the first block in src into dst.
// Dst and src may point at the same memory.
func (rc4Opt *RC4Opt) Encrypt(src []byte) (string, error) {
	if rc4Opt == nil || rc4Opt.secret == nil {
		return "", errors.New("RC4Opt is not properly initialized")
	}
	/* #nosec */
	cipher, err := rc4.NewCipher(rc4Opt.secret)
	if err != nil {
		return "", err
	}
	dst := make([]byte, len(src))
	cipher.XORKeyStream(dst, src)
	return hex.EncodeToString(dst), nil
}

// Decrypt decrypts the first block in src into dst.
// Dst and src may point at the same memory.
func (rc4Opt *RC4Opt) Decrypt(disini []byte) (string, error) {
	if rc4Opt == nil || rc4Opt.secret == nil {
		return "", errors.New("RC4Opt is not properly initialized")
	}
	src, err := hex.DecodeString(string(disini))
	if err != nil {
		return "", err
	}

	/* #nosec */
	cipher, err := rc4.NewCipher(rc4Opt.secret)
	if err != nil {
		return "", err
	}
	dst := make([]byte, len(src))
	cipher.XORKeyStream(dst, src)
	return string(dst), nil
}
