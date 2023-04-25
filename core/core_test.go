package core

import (
	"fmt"
	"github.com/xingcxb/goKit/core/cryptoKit"
	"github.com/xingcxb/goKit/core/httpKit"
	"github.com/xingcxb/goKit/core/otherKit"
	"testing"
)

func TestHttpPost(t *testing.T) {
	params := make(map[string]string)
	params["trade_no"] = "1111"
	params["new_ip"] = "171.42.100.153"
	params["reset"] = "1"
	value := otherKit.JoinStringsInASCII(params, "&", false, false)
	value = value + "&key=123123"
	sign := cryptoKit.Md5(value)
	params["sign"] = sign
	response, err := httpKit.HttpPostFull("http://v2.api.juliangip.com/dynamic/replaceWhiteIp", nil, params, "", -1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response)
}

func TestDownload(t *testing.T) {
	fmt.Println(httpKit.HttpDownload("https://juliangip.com/upload/ipInfo.xlsx", "/Users/symbol/Downloads", "", false))
}
