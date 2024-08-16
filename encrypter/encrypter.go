package encrypter

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"os"
)

type Encrypter struct {
	Key string
}

func NewEncrypter() *Encrypter {
	key := os.Getenv("KEY")

	if key == "" {
		panic("Не передан параметр KEY в переменные окружения")
	}

	return &Encrypter{
		Key: key,
	}
}

func (enc *Encrypter) Encrypt(plainString []byte) []byte {
	block, err := aes.NewCipher([]byte(enc.Key))
	if err != nil {
		panic(err.Error())
	}

	aesGSN, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	nonce := make([]byte, aesGSN.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		panic(err.Error())
	}

	return aesGSN.Seal(nonce, nonce, plainString, nil)

}

func (enc *Encrypter) Decrypt(encrypString []byte) []byte {
	block, err := aes.NewCipher([]byte(enc.Key))
	if err != nil {
		panic(err.Error())
	}

	aesGSN, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	nonceSize := aesGSN.NonceSize()
	nonce, cipherText := encrypString[:nonceSize], encrypString[nonceSize:]
	plainText, err := aesGSN.Open(nil, nonce, cipherText, nil)
	if err != nil {
		panic(err.Error())
	}
	return plainText
}
