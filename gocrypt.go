package gocrypt

type GocryptInterface interface {
	Encrypt(interface{}) error
	Decrypt(interface{}) error
	EncryptString(string) (string, error)
	DecryptString(string) (string, error)
}

type Option struct {
	Aes *AesOpt
}

func New(opt *Option) *Option {
	return opt
}

func (opt *Option) Encrypt(structVal interface{}) error {
	return read(structVal, opt.encrypt)
}

func (opt *Option) Decrypt(structVal interface{}) error {
	return read(structVal, opt.decrypt)
}

func (opt *Option) encrypt(algo string, plain string) (string, error) {
	plaintext := []byte(plain)
	switch algo {
	case "aes":
		return opt.Aes.encryptAES(plaintext)
	default:
		return opt.Aes.encryptAES(plaintext)
	}
}

func (opt *Option) decrypt(algo string, plaintext string) (string, error) {
	switch algo {
	case "aes":
		return opt.Aes.decryptAES(plaintext)
	default:
		return opt.Aes.decryptAES(plaintext)
	}
}
