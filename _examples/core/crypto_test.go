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
