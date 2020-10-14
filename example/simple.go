package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/firdasafridi/gocrypt"
)

type ABC struct {
	A   string `json:"a"`
	B   int    `json:"b"`
	C   int64  `json:"c"`
	DEF *DEF   `json:"def"`
}

type DEF struct {
	D   string `json:"d" gocrypt:"aes"`
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
	sample = "fa89277fb1e1c344709190deeac4465c2b28396423c8534a90c86322d0ec9dcf"
)

func main() {
	aesOpt, err := gocrypt.NewAESOpt(sample)
	if err != nil {
		log.Println("ERR", err)
		return
	}
	cryptRunner := gocrypt.New(&gocrypt.Option{
		Aes: aesOpt,
	})
	a := &ABC{
		A: "halo",
		DEF: &DEF{
			GHI: &GHI{
				G: "halo",
			},
			D: "halo",
			F: 1,
		},
	}

	err = cryptRunner.Encrypt(a)
	if err != nil {
		log.Println("ERR", err)
		return
	}
	strEncrypt, _ := json.Marshal(a)
	fmt.Println("Result the struct encrypted:", string(strEncrypt))

	err = cryptRunner.Decrypt(a)
	if err != nil {
		log.Println("ERR", err)
		return
	}
	strDecrypted, _ := json.Marshal(a)
	fmt.Println("Result the struct decrypted:", string(strDecrypted))
}
