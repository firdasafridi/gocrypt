package main

import (
	"fmt"
	"log"

	"github.com/firdasafridi/gocrypt"
)

const (
	// it's character 24 bit
	key = "123456781234567812345678"
)

func main() {
	// sample plain text
	sampleText := "Halo this is encrypted text!!!"

	// define RC4 option
	rc4Opt, err := gocrypt.NewRC4Opt(key)
	if err != nil {
		log.Println("ERR", err)
		return
	}

	// Encrypt text using RC4 algorithm
	cipherText, err := rc4Opt.Encrypt([]byte(sampleText))
	if err != nil {
		log.Println("ERR", err)
		return
	}
	fmt.Println("Encrypted data:", cipherText)

	// Decrypt text using RC4 algorithm
	plainText, err := rc4Opt.Decrypt([]byte(cipherText))
	if err != nil {
		log.Println("ERR", err)
		return
	}
	fmt.Println("Decrypted data:", plainText)
}
