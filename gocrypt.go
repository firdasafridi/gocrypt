package gocrypt

// GocryptInterface is facing the format gocrypt option library
type GocryptInterface interface {
	Encrypt(interface{}) error
	Decrypt(interface{}) error
}

// GocryptOption is facing the format encryption and decryption format
type GocryptOption interface {
	Encrypt(plainText []byte) (string, error)
	Decrypt(chipherText []byte) (string, error)
}

//
// Option contains an option from initial algorithm encryptioin & decryption.
//
type Option struct {
	AESOpt  GocryptOption
	DESOpt  GocryptOption
	RC4Opt  GocryptOption
	Custom  map[string]GocryptOption
	Prefix  string
	Postfix string
}

//
// New create and initialize new option for struct field encryption.
//
// It needs option from aes, rc4, or des for initialitaion
//
func New(opt *Option) *Option {
	return opt
}

//
// Encrypt is function to set struct field encrypted
//
func (opt *Option) Encrypt(structVal interface{}) error {
	return read(structVal, opt.encrypt)
}

//
// Decrypt is function to set struct field decrypted
//
func (opt *Option) Decrypt(structVal interface{}) error {
	return read(structVal, opt.decrypt)
}

func (opt *Option) encrypt(algo string, plainText string) (string, error) {
	plainByte := []byte(plainText)
	switch algo {
	case "aes":
		return opt.AESOpt.Encrypt(plainByte)
	case "des":
		return opt.DESOpt.Encrypt(plainByte)
	case "rc4":
		return opt.RC4Opt.Encrypt(plainByte)
	default:
		return opt.AESOpt.Encrypt(plainByte)
	}
}

func (opt *Option) decrypt(algo string, chipperText string) (string, error) {
	chipperByte := []byte(chipperText)
	switch algo {
	case "aes":
		return opt.AESOpt.Decrypt(chipperByte)
	case "des":
		return opt.DESOpt.Decrypt(chipperByte)
	case "rc4":
		return opt.RC4Opt.Decrypt(chipperByte)
	default:
		return opt.AESOpt.Decrypt(chipperByte)
	}
}
