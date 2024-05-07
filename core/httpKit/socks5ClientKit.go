package httpKit

import (
	"errors"
	"fmt"
	"golang.org/x/net/proxy"
	"io"
	"net/http"
)

const (
	socks5Version  = 0x05 // socks5版本号
	cmdConnect     = 0x01 // 请求类型
	addrTypeIPV4   = 0x01 // 地址类型，ipv4
	addrTypeIPV6   = 0x04 // 地址类型，ipv6
	addrTypeDomain = 0x03 // 地址类型，域名
)

// Socks5ProxyClient socks5代理请求
/*
 *@params urlStr: 请求地址
 *@params proxyIpPort: 代理ip:端口
 *@params username: 用户名
 *@params password: 密码
 *@return
 */
func Socks5ProxyClient(urlStr, proxyIpPort, username, password string) (string, error) {
	if proxyIpPort == "" {
		return "", errors.New("proxyIpPort Parameters must be passed ")
	}
	var auth *proxy.Auth
	if username != "" && password != "" {
		auth = &proxy.Auth{
			User:     username,
			Password: password,
		}
	} else {
		auth = nil
	}
	// 创建socks5代理Dialer
	dialer, err := proxy.SOCKS5("tcp", proxyIpPort, auth, proxy.Direct)
	if err != nil {
		fmt.Println("Failed to create socks5 dialer:", err)
		return "", err
	}
	// 创建HTTP客户端并设置代理
	httpClient := &http.Client{
		Transport: &http.Transport{
			Dial: dialer.Dial,
		},
	}
	// 发送HTTP GET请求
	resp, err := httpClient.Get(urlStr)
	if err != nil {
		fmt.Println("Failed to send HTTP request:", err)
		return "", err
	}
	defer resp.Body.Close()
	// 读取响应内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to read response body:", err)
		return "", err
	}
	return string(body), nil
}
