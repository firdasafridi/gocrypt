package gocrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"

	"github.com/pkg/errors"
)

// AES256GCMOpt contains all aes-256-gcm session option
// This is a separate implementation to ensure compatibility with other languages
// and to maintain backward compatibility with existing aes tag
type AES256GCMOpt struct {
	aesGCM cipher.AEAD
}

// NewAES256GCMOpt is function to create new configuration of aes-256-gcm algorithm option
// the secret must be hexa a-f & 0-9, 64 characters (32 bytes = 256 bits)
func NewAES256GCMOpt(secret string) (*AES256GCMOpt, error) {
	if len(secret) != 64 {
		return nil, errors.New("Secret must be 64 character (256 bits)")
	}
	key, err := hex.DecodeString(secret)
	if err != nil {
		return nil, errors.Wrap(err, "NewAES256GCMOpt.hex.DecodeString")
	}

	// Create a new Cipher Block from the key
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, errors.Wrap(err, "NewAES256GCMOpt.aes.NewCipher")
	}

	// Create a new GCM - https://en.wikipedia.org/wiki/Galois/Counter_Mode
	// https://golang.org/pkg/crypto/cipher/#NewGCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, errors.Wrap(err, "NewAES256GCMOpt.cipher.NewGCM")
	}

	return &AES256GCMOpt{
		aesGCM: aesGCM,
	}, nil
}

// Encrypt is function to encrypt data using AES-256-GCM algorithm
// Format: nonce (12 bytes) + ciphertext, all hex encoded
// This format is compatible with JavaScript crypto.subtle API
func (aesOpt *AES256GCMOpt) Encrypt(plainText []byte) (string, error) {
	if aesOpt == nil || aesOpt.aesGCM == nil {
		return "", errors.New("AES256GCMOpt is not properly initialized")
	}

	// Create a nonce. Nonce should be from GCM (12 bytes for AES-GCM)
	nonceSize := aesOpt.aesGCM.NonceSize()
	nonce := make([]byte, nonceSize)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", errors.Wrap(err, "Encrypt.io.ReadFull")
	}

	// Encrypt the data using aesGCM.Seal
	// The nonce is prefixed to the encrypted data for compatibility with JavaScript
	ciphertext := aesOpt.aesGCM.Seal(nonce, nonce, plainText, nil)
	return fmt.Sprintf("%x", ciphertext), nil
}

// Decrypt is function to decrypt data using AES-256-GCM algorithm
// Format: nonce (12 bytes) + ciphertext, all hex encoded
// This format is compatible with JavaScript crypto.subtle API
func (aesOpt *AES256GCMOpt) Decrypt(cipherText []byte) (string, error) {
	if aesOpt == nil || aesOpt.aesGCM == nil {
		return "", errors.New("AES256GCMOpt is not properly initialized")
	}

	enc, err := hex.DecodeString(string(cipherText))
	if err != nil {
		return "", errors.Wrap(err, "Decrypt.hex.DecodeString")
	}

	// Get the nonce size (12 bytes for AES-GCM)
	nonceSize := aesOpt.aesGCM.NonceSize()
	if len(enc) < nonceSize {
		return "", errors.New("The data can't be decrypted: ciphertext too short")
	}

	// Extract the nonce from the encrypted data
	nonce, ciphertext := enc[:nonceSize], enc[nonceSize:]

	// Decrypt the data
	plainText, err := aesOpt.aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", errors.Wrap(err, "Decrypt.aesGCM.Open")
	}

	return string(plainText), nil
}
