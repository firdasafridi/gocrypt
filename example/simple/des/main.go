package main

import (
	"fmt"
	"log"

	"github.com/firdasafridi/gocrypt"
)

const (
	// it's character 24 bit
	key = "12345678" + "12345678" + "12345678"
)

func main() {
	// sample plain text
	sampleText := "Halo this is encrypted text!!!"

	// define DES option
	desOpt, err := gocrypt.NewDESOpt(key)
	if err != nil {
		log.Println("ERR", err)
		return
	}

	// Encrypt text using DES algorithm
	cipherText, err := desOpt.Encrypt([]byte(sampleText))
	if err != nil {
		log.Println("ERR", err)
		return
	}
	fmt.Println("Encrypted data", string(cipherText))

	// Decrypt text using DES algorithm
	plainText, err := desOpt.Decrypt([]byte(cipherText))
	if err != nil {
		log.Println("ERR", err)
		return
	}
	fmt.Println("Decrypted data", string(plainText))

}
