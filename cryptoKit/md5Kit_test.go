package cryptoKit

import (
	"fmt"
	"testing"
)

func TestMd5(t *testing.T) {
	fmt.Println(Md5("12345678901234567890"))
}
