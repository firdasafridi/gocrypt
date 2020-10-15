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
