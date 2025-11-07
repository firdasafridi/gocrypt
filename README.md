# gocrypt
`gocrypt` is encryption/decryption library for golang. 

Library gocrypt provides a library to do encryption instruct with go field. The library gocrypt provide instruct tag encryption or inline encryption and decryption

## Overview
[![Go Doc](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/firdasafridi/gocrypt)
[![Go Report Card](https://goreportcard.com/badge/github.com/firdasafridi/gocrypt)](https://goreportcard.com/report/github.com/firdasafridi/gocrypt)


## The package supported:

### **DES3** — Triple Data Encryption Standard
The DES ciphers are primarily supported for PBE standard that provides the option of generating an encryption key based on a passphrase.

### **AES** — Advanced Encryption Standard
The AES cipher is the current U.S. government standard for all software and is recognized worldwide.

### **AES-256-GCM** — Advanced Encryption Standard with 256-bit key and Galois/Counter Mode
The AES-256-GCM implementation (tag: `aes256gcm`) is designed for cross-language compatibility. Data encrypted with this implementation can be decrypted in other languages (e.g., JavaScript) using the same format. See [jscrypt](../jscrypt/) for the JavaScript companion library.

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
    "profile": "akldfjiaidjfods==",
    "id": {
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
	"time"

	"github.com/firdasafridi/gocrypt"
)

// Data contains identity and profile user
type Data struct {
	Profile  *Profile  `json:"profile"`
	Identity *Identity `json:"identity"`
}

// Profile contains name and phone number user
type Profile struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number" gocrypt:"aes"`
}

// Identity contains id, license number, and expired date
type Identity struct {
	ID            string    `json:"id" gocrypt:"aes"`
	LicenseNumber string    `json:"license_number" gocrypt:"aes"`
	ExpiredDate   time.Time `json:"expired_date"`
}

const (
	// it's random string must be hexa  a-f & 0-9
	aeskey = "fa89277fb1e1c344709190deeac4465c2b28396423c8534a90c86322d0ec9dcf"
)

func main() {

	// define AES option
	aesOpt, err := gocrypt.NewAESOpt(aeskey)
	if err != nil {
		log.Println("ERR", err)
		return
	}

	data := &Data{
		Profile: &Profile{
			Name:        "Batman",
			PhoneNumber: "+62123123123",
		},
		Identity: &Identity{
			ID:            "12345678",
			LicenseNumber: "JSKI-123-456",
		},
	}

	cryptRunner := gocrypt.New(&gocrypt.Option{
		AESOpt: aesOpt,
	})

	err = cryptRunner.Encrypt(data)
	if err != nil {
		log.Println("ERR", err)
		return
	}
	strEncrypt, _ := json.Marshal(data)
	fmt.Println("Encrypted:", string(strEncrypt))

	err = cryptRunner.Decrypt(data)
	if err != nil {
		log.Println("ERR", err)
		return
	}
	strDecrypted, _ := json.Marshal(data)
	fmt.Println("Decrypted:", string(strDecrypted))
}

```

### Others
`gocrypt` also supported inline encryption/decryption.

```go
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
        fmt.Println("Encrypted data:", cipherText)

        // Decrypt text using DES algorithm
        plainText, err := desOpt.Decrypt([]byte(cipherText))
        if err != nil {
            log.Println("ERR", err)
            return
        }
        fmt.Println("Decrypted data:", plainText)
    }
```

### AES-256-GCM Example (Cross-Language Compatible)

```go
package main

import (
	"fmt"
	"log"

	"github.com/firdasafridi/gocrypt"
)

const (
	// Secret must be 64 hex characters (256 bits)
	key = "fa89277fb1e1c344709190deeac4465c2b28396423c8534a90c86322d0ec9dcf"
)

func main() {
	// Define AES-256-GCM option
	aesOpt, err := gocrypt.NewAES256GCMOpt(key)
	if err != nil {
		log.Println("ERR", err)
		return
	}

	// Encrypt text using AES-256-GCM algorithm
	plainText := "Hello from Go!"
	cipherText, err := aesOpt.Encrypt([]byte(plainText))
	if err != nil {
		log.Println("ERR", err)
		return
	}
	fmt.Println("Encrypted:", cipherText)

	// This cipherText can be decrypted in JavaScript using jscrypt
	// See: ../jscrypt/ for the JavaScript library
}
```

### Struct Tag with AES-256-GCM

```go
type Profile struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number" gocrypt:"aes256gcm"`
}

func main() {
	aesOpt, _ := gocrypt.NewAES256GCMOpt(key)
	
	cryptRunner := gocrypt.New(&gocrypt.Option{
		AES256GCMOpt: aesOpt,
	})

	data := &Profile{
		Name:        "John Doe",
		PhoneNumber: "+1234567890",
	}

	cryptRunner.Encrypt(data)
	// PhoneNumber is now encrypted and can be decrypted in JavaScript
}
```

## Cross-Language Compatibility

The `aes256gcm` tag provides full cross-language compatibility. Data encrypted in Go can be decrypted in JavaScript (and vice versa) using the same secret key.

### Format
- **Nonce**: 12 bytes (prefixed to ciphertext)
- **Encoding**: Hexadecimal
- **Key**: 64 hex characters (32 bytes = 256 bits)
- **Algorithm**: AES-256-GCM

## Limitation
`gocrypt` only supports the string type. Need more research & development to support the library for more type data.
