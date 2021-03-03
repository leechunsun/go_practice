package coreStruct

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

func PKCS7Padding(text []byte, blockSize int) []byte {
	padding := blockSize - len(text) % blockSize
	paddingByte := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(text, paddingByte...)
}

func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	padding := int(origData[length-1])
	return origData[:(length - padding)]
}

func AESEncrypt(data, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil{
		return nil, err
	}
	bs := block.BlockSize()
	pdData := PKCS7Padding(data, bs)
	ciph := cipher.NewCBCEncrypter(block, key[:bs])
	crypted := make([]byte, len(pdData))
	ciph.CryptBlocks(crypted, pdData)
	return crypted, nil
}

func AESDecrypt(enData, key []byte) ([]byte, error){
	block, err := aes.NewCipher(key)
	if err != nil{
		return nil, err
	}
	bs := block.BlockSize()
	ciph := cipher.NewCBCDecrypter(block, key[:bs])
	decrypt := make([]byte, len(enData))
	ciph.CryptBlocks(decrypt, enData)
	reDecrypt := PKCS7UnPadding(decrypt)
	return reDecrypt, nil
}



