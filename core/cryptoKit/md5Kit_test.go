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
}
