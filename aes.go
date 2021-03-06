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

// AESOpt contains all aes session option
type AESOpt struct {
	aesGCM cipher.AEAD
}

// NewAESOpt is function to create new configuration of aes algorithm option
// the secret must be hexa a-f & 0-9
func NewAESOpt(secret string) (*AESOpt, error) {
	if len(secret) != 64 {
		return nil, errors.New("Secret must be 64 character")
	}
	key, err := hex.DecodeString(secret)
	if err != nil {
		return nil, errors.Wrap(err, "NewAESOpt.hex.DecodeString")
	}

	//Create a new Cipher Block from the key
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, errors.Wrap(err, "NewAESOpt.aes.NewCipher")
	}

	//Create a new GCM - https://en.wikipedia.org/wiki/Galois/Counter_Mode
	//https://golang.org/pkg/crypto/cipher/#NewGCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, errors.Wrap(err, "NewAESOpt.cipher.NewGCM")
	}

	return &AESOpt{
		aesGCM: aesGCM,
	}, nil
}

// Encrypt is function to encrypt data using AES algorithm
func (aesOpt *AESOpt) Encrypt(plainText []byte) (string, error) {

	//Create a nonce. Nonce should be from GCM
	nonce := make([]byte, aesOpt.aesGCM.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", errors.Wrap(err, "encryptAES.io.ReadFull")
	}

	//Encrypt the data using aesGCM.Seal
	//Since we don't want to save the nonce somewhere else in this case, we add it as a prefix to the encrypted data. The first nonce argument in Seal is the prefix.
	ciphertext := aesOpt.aesGCM.Seal(nonce, nonce, plainText, nil)
	return fmt.Sprintf("%x", ciphertext), nil
}

// Decrypt is function to decypt data using AES algorithm
func (aesOpt *AESOpt) Decrypt(chiperText []byte) (string, error) {

	enc, _ := hex.DecodeString(string(chiperText))

	//Get the nonce size
	nonceSize := aesOpt.aesGCM.NonceSize()
	if len(enc) < nonceSize {
		return "", errors.New("The data can't be decrypted")
	}
	//Extract the nonce from the encrypted data
	nonce, ciphertext := enc[:nonceSize], enc[nonceSize:]

	//Decrypt the data
	plainText, err := aesOpt.aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", errors.Wrap(err, "decryptAES.aesGCM.Open")
	}

	return fmt.Sprintf("%s", plainText), nil
}
