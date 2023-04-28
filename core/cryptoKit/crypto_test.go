package cryptoKit

import (
	"fmt"
	"testing"
)

func TestMd5(t *testing.T) {
	//fmt.Println(Md5("12345678901234567890"))
	//fmt.Println(Md5ToUpper("12345678901234567890"))
	//fmt.Println(Base64Encode("12345678901234567890"))
	//fmt.Println(Base64Decode("MTIzNDU2Nzg5MDEyMzQ1Njc4OTA="))
	//fmt.Println(UrlEncode("http://www.oschina.net/search?scope=bbs&q=C语言"))
	//fmt.Println(UrlDecode("http://www.oschina.net/search?scope=bbs&q=C%E8%AF%AD%E8%A8%80"))
	//fmt.Println(AESEncode([]byte("1000000000000000"), []byte("1111")))
	//fmt.Println(AESDecode("GOol+ab8C1YsiT5sOd5raw==", "1000000000000000"))
	fmt.Println(UnicodeDecode("\\u5341\\u5927\\u6237\\uffe5\\u0040\\uff01\\u0023\\u0025\\u2026\\u2026\\u0026\\u2026\\u2026\\u002a\\uff08\\uff09\\u2014\\u2014\\u002b\\u300a\\u300b\\u3001\\uff0c\\u3002\\u3001\\uff1b\\u2018\\u3001\\u914d\\u3010\\u3011"))
	fmt.Println(UnicodeEncode("十大户￥@！#%……&……*（）——+《》、，。、；‘、配【】"))
}
