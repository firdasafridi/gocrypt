package main

import (
	"fmt"
	"log"

	"github.com/firdasafridi/gocrypt"
)

const (
	// it's random string must be hexa a-f & 0-9, 64 characters (256 bits)
	key = "fa89277fb1e1c344709190deeac4465c2b28396423c8534a90c86322d0ec9dcf"
)

func main() {
	// sample plain text
	sampleText := "Halo this is encrypted text!!!"

	// define AES-256-GCM option
	aesOpt, err := gocrypt.NewAES256GCMOpt(key)
	if err != nil {
		log.Println("ERR", err)
		return
	}

	// Encrypt text using AES-256-GCM algorithm
	cipherText, err := aesOpt.Encrypt([]byte(sampleText))
	if err != nil {
		log.Println("ERR", err)
		return
	}
	fmt.Println("Encrypted data:", cipherText)

	// Decrypt text using AES-256-GCM algorithm
	plainText, err := aesOpt.Decrypt([]byte(cipherText))
	if err != nil {
		log.Println("ERR", err)
		return
	}
	fmt.Println("Decrypted data:", plainText)
}
