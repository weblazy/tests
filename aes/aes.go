package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

type Aes struct {
	key     []byte
	vector  []byte
	block   cipher.Block
	mode    *CBCMode
	padding *PKCS7Padding
}

func NewAes(key []byte) *Aes {
	length := len(key)
	if length >= 32 {
		key = key[:32]
	} else {
		arr := make([]byte, 32)
		for k1 := range key {
			arr[k1] = key[k1]
		}
		key = arr
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil
	}
	return &Aes{
		key:     key,
		block:   block,
		vector:  key[:block.BlockSize()],
		mode:    &CBCMode{},
		padding: &PKCS7Padding{},
	}
}

// Encrypt aes加密
func (this *Aes) Encrypt(origData string) (string, error) {
	data := []byte(origData)
	// padding
	data = this.padding.Padding(data, this.block.BlockSize())
	// mode
	ciphertext, err := this.mode.Encrypt(this.block, data, this.key, this.vector)
	if err != nil {
		return "", err
	}
	// format
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// Decrypt aes解密
func (this *Aes) Decrypt(data string) (string, error) {
	// format
	ciphertext, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}
	// mode
	origData, err := this.mode.Decrypt(this.block, ciphertext, this.key, this.vector)
	if err != nil {
		return "", err
	}
	// padding
	return string(this.padding.UnPadding(origData)), nil
}

type PKCS7Padding struct {
}

func (this *PKCS7Padding) Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func (this *PKCS7Padding) UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

type CBCMode struct {
}

func (this *CBCMode) Encrypt(block cipher.Block, data, key, iv []byte) ([]byte, error) {
	mode := cipher.NewCBCEncrypter(block, iv)
	ciphertext := make([]byte, len(data))
	mode.CryptBlocks(ciphertext, data)
	return ciphertext, nil
}

func (this *CBCMode) Decrypt(block cipher.Block, data, key, iv []byte) ([]byte, error) {
	mode := cipher.NewCBCDecrypter(block, iv)
	plaintext := make([]byte, len(data))
	mode.CryptBlocks(plaintext, data)
	return plaintext, nil
}
