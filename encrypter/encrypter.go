package encrypter

import "os"

type Encrypter struct {
	Key string
}

func NewEncrypter() *Encrypter {
	key := os.Getenv("KEy")

	if key == "" {
		panic("Не передан параметр KEY в переменные окружения")
	}

	return &Encrypter{
		Key: key,
	}
}

func (enc *Encrypter) Encrypt(plainString string) string {
	return ""
}

func (enc *Encrypter) Decrypt(encrypString string) string {
	return ""
}
