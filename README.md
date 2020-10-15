# gocrypt
`gocrypt` is encryption/decryption library for golang. 

Package gocrypt provide a library for do encryption in struct with go field.Package gocrypt provide in struct tag encryption or inline encryption and decryption

## Overview
[![Go Doc](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/firdasafridi/gocrypt)
[![Go Report Card](https://goreportcard.com/badge/github.com/firdasafridi/gocrypt)](https://goreportcard.com/report/github.com/firdasafridi/gocrypt)


## The package supported:

### **DES3** — Triple Data Encryption Standard
The AES cipher is the current U.S. government standard for all software, and is recognized worldwide.

### **AES** — Advanced Encryption Standard
The DES ciphers are primarily supported for PBE standard that provides the option of generating an encryption key based on a passphrase.

### **RC4** — stream chipper
The RC4 is supplied for situations that call for fast encryption, but not strong encryption. RC4 is ideal for situations that require a minimum of encryption.

## The Idea
### Sample JSON Payload
Before the encryption:
```
{
    "a": "Sample plain text",
    "b": {
        c: "Sample plain text 2"
    }
}
```
After the encryption:
```
{
    "a": "akldfjiaidjfods==",
    "b": {
        c: "Ijdsifsjiek18239=="
    }
}
```
### Struct Tag Field Implementation
```go
type A struct {
    A string    `json:"a" gocrypt:"aes"`
    B B         `json:"b"`
}

type B struct {
    C string    `json:"c" gocrypt:"aes"`
}
```

## Full Sample
```go
package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/firdasafridi/gocrypt"
)

// ABC is sample structure
type A struct {
	A   string `json:"a" gocrypt:"des"`
	B   *B     `json:"b"`
}

// B is sample structure
type B struct {
	D   string `json:"d" gocrypt:"des"`
}

const (
	// it's character 24 bit
	deskey = "xxxxxxxxxxxxxxxxxxxxxxxx"
)

func main() {
	// define DES option
	desOpt, err := gocrypt.NewDESOpt(deskey)
	if err != nil {
		log.Println("ERR", err)
		return
	}	return
	}

	cryptRunner := gocrypt.New(&gocrypt.Option{
		DESOpt: desOpt,
	})
	a := &ABC{
		A: "Halo this is encrypted RC4!!!",
		B: &B{
			D: "Halo this is encrypted des!!!",
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

```

