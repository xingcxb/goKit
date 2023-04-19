package cryptoKit

import (
	"fmt"
	"testing"
)

func TestMd5(t *testing.T) {
	fmt.Println(Md5("12345678901234567890"))
	fmt.Println(Md5ToUpper("12345678901234567890"))
	fmt.Println(Base64Encode("12345678901234567890"))
	fmt.Println(Base64Decode("MTIzNDU2Nzg5MDEyMzQ1Njc4OTA="))
	fmt.Println(UrlEncode("http://www.oschina.net/search?scope=bbs&q=C语言"))
	fmt.Println(UrlDecode("http://www.oschina.net/search?scope=bbs&q=C%E8%AF%AD%E8%A8%80"))
	//fmt.Println(AESEncode([]byte("1000000000000000"), []byte("1111")))
	//fmt.Println(AESDecode("GOol+ab8C1YsiT5sOd5raw==", "1000000000000000"))
}
