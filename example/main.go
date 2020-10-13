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
	GHI *GHI   `json:"ghi"`
	D   string `json:"d" gocrypt:"true"`
	E   int    `json:"e"`
	F   int64  `json:"f"`
}

type GHI struct {
	G string `json:"g" gocrypt:"true"`
	H int    `json:"h"`
	I int64  `json:"i"`
}

func main() {
	aesOpt, err := gocrypt.NewAESOpt()
	if err != nil {
		log.Println("ERR", err)
		return
	}
	cryptRunner := gocrypt.New(&gocrypt.Option{
		Aes: aesOpt,
	})
	a := &ABC{
		A: "here A",
		DEF: &DEF{
			GHI: &GHI{
				G: "here",
			},
			D: "a",
			F: 1,
		},
	}
	cryptRunner.Encrypt(a)
	strEncrypt, _ := json.Marshal(a)
	fmt.Println("Result the struct encrypted:", string(strEncrypt))

	cryptRunner.Decrypt(a)
	strDecrypted, _ := json.Marshal(a)
	fmt.Println("Result the struct decrypted:", string(strDecrypted))
}
