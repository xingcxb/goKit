package core

import (
	"fmt"
	"github.com/xingcxb/goKit/core/cryptoKit"
	"testing"
)

func TestBase64Encode(t *testing.T) {
	fmt.Println(cryptoKit.Base64Encode("123123"))
}

func TestBase64Decode(t *testing.T) {
	fmt.Println(cryptoKit.Base64Decode("MTIzMTIz"))
}

func TestMd5(t *testing.T) {
	fmt.Println(cryptoKit.Md5("123123"))
}

func TestMd5ToUpper(t *testing.T) {
	fmt.Println(cryptoKit.Md5ToUpper("123123"))
}

func TestSha256(t *testing.T) {
	fmt.Println(cryptoKit.Sha256("123123"))
}

func TestSha256ToUpper(t *testing.T) {
	fmt.Println(cryptoKit.Sha256ToUpper("123123"))
}

func TestUnicodeEncode(t *testing.T) {
	fmt.Println(cryptoKit.UnicodeEncode("123123"))
}

func TestUnicodeDecode(t *testing.T) {
	fmt.Println(cryptoKit.UnicodeDecode("\\u0031\\u0032\\u0033\\u0031\\u0032\\u0033"))
}

func TestUrlEncode(t *testing.T) {
	fmt.Println(cryptoKit.UrlEncode("https://www.baidu.com/s?wd=你好"))
}

func TestUrlDecode(t *testing.T) {
	fmt.Println(cryptoKit.UrlDecode("https%3A%2F%2Fwww.baidu.com%2Fs%3Fwd%3D%E4%BD%A0%E5%A5%BD"))
}

// aes cbc模式加密
func TestAESCBCEncrypt(t *testing.T) {
	v, err := cryptoKit.AESEncryptCBC([]byte("pibigstar"), []byte("1234567891234567"), []byte("1231231231231231"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cryptoKit.Base64Encode(string(v)))
	fmt.Println(string(v))
}

// aes cbc模式解密
func TestAESCBCDecrypt(t *testing.T) {
	value, _ := cryptoKit.Base64Decode("MlW3thrpN09qa6gUzyDI+g==")
	v, err := cryptoKit.AESDecryptCBC([]byte(value), []byte("1234567891234567"), []byte("1231231231231231"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(v)
	fmt.Println(string(v))
}

func TestAESEncryptCFB(t *testing.T) {
	padding := 0
	v, err := cryptoKit.AESEncryptCFB([]byte("pibigstar"), []byte("1234567891234567"), &padding, []byte("1231231231231231"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cryptoKit.Base64Encode(string(v)))
	fmt.Println(string(v))
}

func TestAESDecryptCFB(t *testing.T) {
	value, _ := cryptoKit.Base64Decode("qDJPT/XQZNJo5YRfitRUCg==")
	v, err := cryptoKit.AESDecryptCFB([]byte(value), []byte("1234567891234567"), 16-len(value), []byte("1231231231231231"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(v)
	fmt.Println(string(v))
}
