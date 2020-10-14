package gocrypt

import (
	"crypto/rc4"
	"encoding/hex"
)

// RC4Opt is tructure of aes option
type RC4Opt struct {
	cipher *rc4.Cipher
	secret string
}

// NewRC4Opt is function to create new configuration of aes algorithm option
// the secret must be hexa a-f & 0-9
func NewRC4Opt(secret string) (*RC4Opt, error) {
	return &RC4Opt{
		secret: secret,
	}, nil
}

// Encrypt encrypts the first block in src into dst.
// Dst and src may point at the same memory.
func (rc4Opt *RC4Opt) Encrypt(src []byte) (string, error) {

	cipher, err := rc4.NewCipher([]byte(rc4Opt.secret))
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
	src, err := hex.DecodeString(string(disini))
	if err != nil {
		return "", err
	}

	cipher, err := rc4.NewCipher([]byte(rc4Opt.secret))
	if err != nil {
		return "", err
	}
	dst := make([]byte, len(src))
	cipher.XORKeyStream(dst, src)
	return string(dst), nil
}
