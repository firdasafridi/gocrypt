package gocrypt

type GocryptInterface interface {
	Encrypt(interface{}) error
	Decrypt(interface{}) error
	EncryptString(string) (string, error)
	DecryptString(string) (string, error)
}

type GocryptOption interface {
	Encrypt(plainText []byte) (string, error)
	Decrypt(chipherText []byte) (string, error)
}

type Option struct {
	AESOpt  GocryptOption
	DESOpt  GocryptOption
	Custom  map[string]GocryptOption
	Prefix  string
	Postfix string
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

func (opt *Option) encrypt(algo string, plainText string) (string, error) {
	plainByte := []byte(plainText)
	switch algo {
	case "aes":
		return opt.AESOpt.Encrypt(plainByte)
	case "des":
		return opt.DESOpt.Encrypt(plainByte)
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
	default:
		return opt.AESOpt.Decrypt(chipperByte)
	}
}
