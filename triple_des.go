package gocrypt

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"

	"github.com/pkg/errors"
)

// DESOpt containts all DES session option
type DESOpt struct {
	block     cipher.Block
	blockSize []byte
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

	blockSize := secret[:des.BlockSize]

	return &DESOpt{
		block:     block,
		blockSize: []byte(blockSize),
	}, nil
}

// Encrypt is function to encrypt data using DES algorithm
func (desOpt *DESOpt) Encrypt(plainText []byte) (string, error) {
	block := desOpt.block
	iv := desOpt.blockSize
	origData := pkcs5Padding(plainText, block.BlockSize())
	mode := cipher.NewCBCEncrypter(block, iv)
	encrypted := make([]byte, len(origData))
	mode.CryptBlocks(encrypted, origData)
	return base64.URLEncoding.EncodeToString(encrypted), nil
}

// Decrypt is function to decypt data using DES algorithm
func (desOpt *DESOpt) Decrypt(chiperText []byte) (string, error) {
	block := desOpt.block
	rbyte, err := base64.URLEncoding.DecodeString(string(chiperText))
	if err != nil {
		return "", err
	}
	iv := desOpt.blockSize
	decrypter := cipher.NewCBCDecrypter(block, iv)
	decrypted := make([]byte, len(rbyte))
	decrypter.CryptBlocks(decrypted, rbyte)
	decrypted = pkcs5Unpadding(decrypted)
	return string(decrypted), nil
}

func pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func pkcs5Unpadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
