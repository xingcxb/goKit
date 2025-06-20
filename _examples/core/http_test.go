package core

import (
	"fmt"
	"github.com/xingcxb/goKit/core/dateKit"
	"github.com/xingcxb/goKit/core/httpKit"
	"github.com/xingcxb/goKit/core/httpKit/uaKit"
	"github.com/xingcxb/goKit/core/strKit"
	"net/http"
	"testing"
	"time"
)

func TestHttpDownload(t *testing.T) {
	fmt.Println(httpKit.HttpDownload("https://dl.google.com/chrome/mac/universal/stable/GGRO/googlechrome.dmg", "/Users/symbol/Downloads", "", true))
}

func TestHttpGet(t *testing.T) {
	fmt.Println(httpKit.HttpGet("https://www.xingcxb.com"))
}

func TestHttpGetFull(t *testing.T) {
	fmt.Println(httpKit.HttpGetFull("https://www.xingcxb.com", nil, nil, nil, 300))
}

func TestHttpPost(t *testing.T) {
	fmt.Println(httpKit.HttpPost("https://www.xingcxb.com", nil))
}

func TestHttpPostFull(t *testing.T) {
	fmt.Println(httpKit.HttpPostFull("https://www.xingcxb.com", nil, nil, nil, 300))
}

func TestHttpBasic(t *testing.T) {
	fmt.Println(httpKit.HttpBasic("https://www.xingcxb.com", http.MethodGet, nil, nil, nil, 300))
}

func TestHttpProxyGet(t *testing.T) {
	//for i := 0; i < 100; i++ {
	//	fmt.Println("第", i, "次，时间为：", dateKit.Now())
	_, s, err := httpKit.HttpProxyGet("https://xingcxb.com", "117.68.38.158:23816")
	fmt.Println(strKit.SubString(s, 0, 100), err)
	//}
}

func TestHttpProxyGetFull(t *testing.T) {
	fmt.Println(httpKit.HttpProxyGetFull("https://cip.cc", nil, nil,
		nil, 300, "socks5", "u", "p",
		"222.222.222.222:2222"))
}

func TestHttpProxyPost(t *testing.T) {
	fmt.Println(httpKit.HttpProxyPost("https://cip.cc", nil,
		"255.255.255.255:52724"))
}

func TestHttpProxyPostFull(t *testing.T) {
	fmt.Println(httpKit.HttpProxyPostFull("https://cip.cc", nil, nil,
		nil, 300, "http", "u", "p", ""))
}

func TestHttpProxyBasic(t *testing.T) {
	fmt.Println(httpKit.HttpProxyBasic("https://cip.cc", http.MethodGet,
		nil, nil, nil, 300, "http",
		"", "", "255.255.255.255:29093"))
}

func TestUA(t *testing.T) {
	fmt.Println(uaKit.ParseUA("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36"))
}

func TestGetIps(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("第", i, "次，时间为：", dateKit.Now())
			v, err := httpKit.HttpGet("http://1.1.1.1")
			fmt.Println(v, err)
		}()
		go func() {
			fmt.Println("第", i, "次，时间为：", dateKit.Now())
			v, err := httpKit.HttpGet("http://1.1.1.1")
			fmt.Println(v, err)
		}()
	}
	time.Sleep(time.Hour * 10)
}

func TestSocks5Client(t *testing.T) {
	fmt.Println(httpKit.Socks5ProxyClient("https://cip.cc", "1.1.1.1:123456", "123456", "123456"))
}
