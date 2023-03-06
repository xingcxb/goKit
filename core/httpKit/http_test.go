package httpKit

import (
	"fmt"
	"goKit/core/ipKit"
	"testing"
)

func TestHttpGet(t *testing.T) {
	fmt.Println(HttpGetFull("https://xingcxb.com", nil, nil, "", 2000))
}

func TestAddr(t *testing.T) {
	fmt.Println(ipKit.GetIpAddressInfo("127.0.0.1"))
}
