package process

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"

	"unsafe"
)

func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

//AES加密
func AesEncrypt(body string) (string, error) {

	origData := []byte(body)
	block, err := aes.NewCipher(AESKEY)
	if err != nil {
		return "", nil
	}
	blockSize := block.BlockSize()
	origData = PKCS7Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, AESKEY[:blockSize])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	result := *(*string)(unsafe.Pointer(&crypted))
	return result, nil
}

//AES解密
func AesDecrypt(body string) ([]byte, error) {
	crypted := []byte(body)
	block, err := aes.NewCipher(AESKEY)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, AESKEY[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS7UnPadding(origData)
	return origData, nil
}
