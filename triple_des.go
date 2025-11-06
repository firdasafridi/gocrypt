package gocrypt

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"crypto/rand"
	"encoding/base64"
	"io"

	"github.com/pkg/errors"
)

// DESOpt contains all DES session option
type DESOpt struct {
	block     cipher.Block
	blockSize int
}

// NewDESOpt is function to create new configuration of des algorithm option
// the secret must be 24 character
func NewDESOpt(secret string) (*DESOpt, error) {
	if len(secret) != 24 {
		return nil, errors.New("secret must be 24 char")
	}

	/* #nosec */
	block, err := des.NewTripleDESCipher([]byte(secret))
	if err != nil {
		return nil, errors.Wrap(err, "NewDESOpt.des.NewTripleDESCipher")
	}

	return &DESOpt{
		block:     block,
		blockSize: des.BlockSize,
	}, nil
}

// Encrypt is function to encrypt data using DES algorithm
func (desOpt *DESOpt) Encrypt(plainText []byte) (string, error) {
	if desOpt == nil || desOpt.block == nil {
		return "", errors.New("DESOpt is not properly initialized")
	}
	block := desOpt.block
	blockSize := desOpt.blockSize

	// Generate a random IV for each encryption
	iv := make([]byte, blockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", errors.Wrap(err, "Encrypt.io.ReadFull")
	}

	origData := pkcs5Padding(plainText, blockSize)
	mode := cipher.NewCBCEncrypter(block, iv)
	encrypted := make([]byte, len(origData))
	mode.CryptBlocks(encrypted, origData)

	// Prepend IV to ciphertext (like AES does with nonce)
	ciphertext := append(iv, encrypted...)
	return base64.URLEncoding.EncodeToString(ciphertext), nil
}

// Decrypt is function to decypt data using DES algorithm
func (desOpt *DESOpt) Decrypt(cipherText []byte) (string, error) {
	if desOpt == nil || desOpt.block == nil {
		return "", errors.New("DESOpt is not properly initialized")
	}
	block := desOpt.block
	blockSize := desOpt.blockSize

	rbyte, err := base64.URLEncoding.DecodeString(string(cipherText))
	if err != nil {
		return "", errors.Wrap(err, "Decrypt.base64.URLEncoding.DecodeString")
	}

	// Extract IV from the beginning of the ciphertext
	if len(rbyte) < blockSize {
		return "", errors.New("ciphertext too short to contain IV")
	}
	iv := rbyte[:blockSize]
	ciphertext := rbyte[blockSize:]

	decrypter := cipher.NewCBCDecrypter(block, iv)
	decrypted := make([]byte, len(ciphertext))
	decrypter.CryptBlocks(decrypted, ciphertext)
	decrypted, err = pkcs5Unpadding(decrypted)
	if err != nil {
		return "", errors.Wrap(err, "Decrypt.pkcs5Unpadding")
	}
	return string(decrypted), nil
}

func pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func pkcs5Unpadding(origData []byte) ([]byte, error) {
	length := len(origData)
	if length == 0 {
		return nil, errors.New("pkcs5Unpadding: data is empty")
	}

	unpadding := int(origData[length-1])
	if unpadding <= 0 || unpadding > length {
		return nil, errors.New("pkcs5Unpadding: invalid padding value")
	}

	// Validate that all padding bytes are the same
	for i := length - unpadding; i < length; i++ {
		if origData[i] != byte(unpadding) {
			return nil, errors.New("pkcs5Unpadding: invalid padding")
		}
	}

	return origData[:(length - unpadding)], nil
}
