// Package cryptoKit 本文件来自于gogf/gf的crypto/gaes/gaes.go
package cryptoKit

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"errors"
)

// AESEncryptCBC AES使用CBC模式加密
/*
 * 注意，密钥必须是16/24/32位长度
 * @param plainText 待加密内容
 * @param key 密钥
 * @param iv 向量
 */
func AESEncryptCBC(plainText []byte, key []byte, iv ...[]byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	plainText = PKCS5Padding(plainText, blockSize)
	ivValue := ([]byte)(nil)
	if len(iv) > 0 {
		ivValue = iv[0]
	} else {
		return nil, errors.New("Iv must exist")
	}
	blockMode := cipher.NewCBCEncrypter(block, ivValue)
	cipherText := make([]byte, len(plainText))
	blockMode.CryptBlocks(cipherText, plainText)

	return cipherText, nil
}

// AESDecryptCBC AES使用CBC模式解密
/*
 * 注意，密钥必须是16/24/32位长度
 * @param cipherText 加密字符串
 * @param key 密钥
 * @param iv 向量
 */
func AESDecryptCBC(cipherText []byte, key []byte, iv ...[]byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	if len(cipherText) < blockSize {
		return nil, errors.New("cipherText too short")
	}
	ivValue := ([]byte)(nil)
	if len(iv) > 0 {
		ivValue = iv[0]
	} else {
		return nil, errors.New("Iv must exist")
	}
	if len(cipherText)%blockSize != 0 {
		return nil, errors.New("cipherText is not a multiple of the block size")
	}
	blockModel := cipher.NewCBCDecrypter(block, ivValue)
	plainText := make([]byte, len(cipherText))
	blockModel.CryptBlocks(plainText, cipherText)
	plainText, e := PKCS5UnPadding(plainText, blockSize)
	if e != nil {
		return nil, e
	}
	return plainText, nil
}

// PKCS5Padding 填充
func PKCS5Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

// PKCS5UnPadding 反填充
func PKCS5UnPadding(src []byte, blockSize int) ([]byte, error) {
	length := len(src)
	if blockSize <= 0 {
		return nil, errors.New("invalid blocklen")
	}

	if length%blockSize != 0 || length == 0 {
		return nil, errors.New("invalid data len")
	}

	unpadding := int(src[length-1])
	if unpadding > blockSize || unpadding == 0 {
		return nil, errors.New("invalid padding")
	}

	padding := src[length-unpadding:]
	for i := 0; i < unpadding; i++ {
		if padding[i] != byte(unpadding) {
			return nil, errors.New("invalid padding")
		}
	}
	return src[:(length - unpadding)], nil
}

// AESEncryptCFB AEC CFB模式加密
/*
 * 注意，密钥必须是16/24/32位长度。
 * @param plainText 明文
 * @param key 密钥
 * @param padding 填充
 * @param iv 向量
 */
func AESEncryptCFB(plainText []byte, key []byte, padding *int, iv ...[]byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	plainText, *padding = ZeroPadding(plainText, blockSize)
	ivValue := ([]byte)(nil)
	if len(iv) > 0 {
		ivValue = iv[0]
	} else {
		return nil, errors.New("Iv must exist")
	}
	stream := cipher.NewCFBEncrypter(block, ivValue)
	cipherText := make([]byte, len(plainText))
	stream.XORKeyStream(cipherText, plainText)
	return cipherText, nil
}

// AESDecryptCFB AES CFB解密
/*
 * 注意，密钥必须是16/24/32位长度。
 * @param cipherText 密文
 * @param key 密钥
 * @param unPadding 去填充
 * @param iv 向量
 */
func AESDecryptCFB(cipherText []byte, key []byte, unPadding int, iv ...[]byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(cipherText) < aes.BlockSize {
		return nil, errors.New("cipherText too short")
	}
	ivValue := ([]byte)(nil)
	if len(iv) > 0 {
		ivValue = iv[0]
	} else {
		return nil, errors.New("Iv must exist")
	}
	stream := cipher.NewCFBDecrypter(block, ivValue)
	plainText := make([]byte, len(cipherText))
	stream.XORKeyStream(plainText, cipherText)
	plainText = ZeroUnPadding(plainText, unPadding)
	return plainText, nil
}

// ZeroPadding 零填充
/*
 * @param cipherText 密文
 * @param blockSize 块大小
 */
func ZeroPadding(cipherText []byte, blockSize int) ([]byte, int) {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(0)}, padding)
	return append(cipherText, padText...), padding
}

// ZeroUnPadding 零未填充
/*
 * @param plaintext 明文
 * @param unPadding 去除数据中的填充
 */
func ZeroUnPadding(plaintext []byte, unPadding int) []byte {
	length := len(plaintext)
	return plaintext[:(length - unPadding)]
}
