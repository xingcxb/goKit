package core

import (
	"fmt"
	"github.com/xingcxb/goKit/core/httpKit"
	"net/http"
	"testing"
)

func TestHttpDownload(t *testing.T) {
	fmt.Println(httpKit.HttpDownload("https://dl.google.com/chrome/mac/universal/stable/GGRO/googlechrome.dmg", "/Users/symbol/Downloads", "", false))
}

func TestHttpGet(t *testing.T) {
	fmt.Println(httpKit.HttpGet("https://www.baidu.com"))
}

func TestHttpGetFull(t *testing.T) {
	fmt.Println(httpKit.HttpGetFull("https://www.baidu.com", nil, nil, "", 300))
}

func TestHttpPost(t *testing.T) {
	fmt.Println(httpKit.HttpPost("https://www.baidu.com", nil))
}

func TestHttpPostFull(t *testing.T) {
	fmt.Println(httpKit.HttpPostFull("https://www.baidu.com", nil, nil, "", 300))
}

func TestHttpBasic(t *testing.T) {
	fmt.Println(httpKit.HttpBasic("https://www.baidu.com", http.MethodGet, nil, nil, "", 300))
}

func TestHttpProxyGet(t *testing.T) {
	fmt.Println(httpKit.HttpProxyGet("https://cip.cc", "255.255.255.255:52724"))
}

func TestHttpProxyGetFull(t *testing.T) {
	fmt.Println(httpKit.HttpProxyGetFull("https://cip.cc", nil, nil,
		"", 300, "http", "u", "p",
		"255.255.255.255:52724"))
}

func TestHttpProxyPost(t *testing.T) {
	fmt.Println(httpKit.HttpProxyPost("https://cip.cc", nil,
		"255.255.255.255:52724"))
}

func TestHttpProxyPostFull(t *testing.T) {
	fmt.Println(httpKit.HttpProxyPostFull("https://cip.cc", nil, nil,
		"", 300, "http", "u", "p", ""))
}

func TestHttpProxyBasic(t *testing.T) {
	fmt.Println(httpKit.HttpProxyBasic("https://cip.cc", http.MethodGet,
		nil, nil, "", 300, "http",
		"", "", "255.255.255.255:52724"))
}
