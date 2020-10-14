package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/firdasafridi/gocrypt"
)

type ABC struct {
	A   string `json:"a" gocrypt:"rc4"`
	B   int    `json:"b"`
	C   int64  `json:"c"`
	DEF *DEF   `json:"def"`
}

type DEF struct {
	D   string `json:"d" gocrypt:"des"`
	E   int    `json:"e"`
	F   int64  `json:"f"`
	GHI *GHI   `json:"ghi"`
}

type GHI struct {
	G string `json:"g" gocrypt:"aes"`
	H int    `json:"h"`
	I int64  `json:"i"`
}

const (
	// it's random string must be hexa  a-f & 0-9
	aeskey = "fa89277fb1e1c344709190deeac4465c2b28396423c8534a90c86322d0ec9dcf"
	// it's character 24 bit
	deskey = "123456781234567812345678"
	// random string key for rc4
	rc4key = "adfasd123123ksdfsd"
)

func main() {

	// define AES option
	aesOpt, err := gocrypt.NewAESOpt(aeskey)
	if err != nil {
		log.Println("ERR", err)
		return
	}

	// define DES option
	desOpt, err := gocrypt.NewDESOpt(deskey)
	if err != nil {
		log.Println("ERR", err)
		return
	}

	// define DES option
	rc4Opt, err := gocrypt.NewRC4Opt(rc4key)
	if err != nil {
		log.Println("ERR", err)
		return
	}

	cryptRunner := gocrypt.New(&gocrypt.Option{
		AESOpt: aesOpt,
		DESOpt: desOpt,
		RC4Opt: rc4Opt,
	})
	a := &ABC{
		A: "Halo this is encrypted RC4!!!",
		DEF: &DEF{
			GHI: &GHI{
				G: "Halo this is encrypted aes!!!",
			},
			D: "Halo this is encrypted des!!!",
			F: 1,
		},
	}

	err = cryptRunner.Encrypt(a)
	if err != nil {
		log.Println("ERR", err)
		return
	}
	strEncrypt, _ := json.Marshal(a)
	fmt.Println("Encrypted:", string(strEncrypt))

	err = cryptRunner.Decrypt(a)
	if err != nil {
		log.Println("ERR", err)
		return
	}
	strDecrypted, _ := json.Marshal(a)
	fmt.Println("Decrypted:", string(strDecrypted))
}
