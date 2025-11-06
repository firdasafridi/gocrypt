package gocrypt

import "github.com/pkg/errors"

// GocryptInterface is facing the format gocrypt option library
type GocryptInterface interface {
	Encrypt(interface{}) error
	Decrypt(interface{}) error
}

// GocryptOption is facing the format encryption and decryption format
type GocryptOption interface {
	Encrypt(plainText []byte) (string, error)
	Decrypt(cipherText []byte) (string, error)
}

// Option contains an option from initial algorithm encryptioin & decryption.
type Option struct {
	AESOpt       GocryptOption
	AES256GCMOpt GocryptOption
	DESOpt       GocryptOption
	RC4Opt       GocryptOption
	Custom       map[string]GocryptOption
	Prefix       string
	Postfix      string
}

// New create and initialize new option for struct field encryption.
//
// It needs option from aes, rc4, or des for initialitaion
func New(opt *Option) *Option {
	return opt
}

// Encrypt is function to set struct field encrypted
func (opt *Option) Encrypt(structVal interface{}) error {
	return read(structVal, opt.encrypt)
}

// Decrypt is function to set struct field decrypted
func (opt *Option) Decrypt(structVal interface{}) error {
	return read(structVal, opt.decrypt)
}

func (opt *Option) encrypt(algo string, plainText string) (string, error) {
	// Convert to []byte only once
	plainByte := []byte(plainText)

	switch algo {
	case "aes":
		if opt.AESOpt == nil {
			return "", errors.New("AESOpt is not initialized")
		}
		return opt.AESOpt.Encrypt(plainByte)
	case "aes256gcm":
		if opt.AES256GCMOpt == nil {
			return "", errors.New("AES256GCMOpt is not initialized")
		}
		return opt.AES256GCMOpt.Encrypt(plainByte)
	case "des":
		if opt.DESOpt == nil {
			return "", errors.New("DESOpt is not initialized")
		}
		return opt.DESOpt.Encrypt(plainByte)
	case "rc4":
		if opt.RC4Opt == nil {
			return "", errors.New("RC4Opt is not initialized")
		}
		return opt.RC4Opt.Encrypt(plainByte)
	default:
		if opt.AESOpt == nil {
			return "", errors.New("AESOpt is not initialized")
		}
		return opt.AESOpt.Encrypt(plainByte)
	}
}

func (opt *Option) decrypt(algo string, cipherText string) (string, error) {
	// Convert to []byte only once
	cipherByte := []byte(cipherText)

	switch algo {
	case "aes":
		if opt.AESOpt == nil {
			return "", errors.New("AESOpt is not initialized")
		}
		return opt.AESOpt.Decrypt(cipherByte)
	case "aes256gcm":
		if opt.AES256GCMOpt == nil {
			return "", errors.New("AES256GCMOpt is not initialized")
		}
		return opt.AES256GCMOpt.Decrypt(cipherByte)
	case "des":
		if opt.DESOpt == nil {
			return "", errors.New("DESOpt is not initialized")
		}
		return opt.DESOpt.Decrypt(cipherByte)
	case "rc4":
		if opt.RC4Opt == nil {
			return "", errors.New("RC4Opt is not initialized")
		}
		return opt.RC4Opt.Decrypt(cipherByte)
	default:
		if opt.AESOpt == nil {
			return "", errors.New("AESOpt is not initialized")
		}
		return opt.AESOpt.Decrypt(cipherByte)
	}
}
