package ipKit

import (
	"encoding/json"
	"github.com/xingcxb/goKit/core/ipKit/ipService"
	"net"
)

// GetIpAddressInfo 获取ip物理地址信息
/**
 * @param ip 需要查询的ip地址
 * @return string,error
 */
func GetIpAddressInfo(ip string) (string, error) {
	ipInfoObj, err := ipService.RandomQuery(ip)
	if err != nil {
		return "", err
	}
	ipInfoJson, err := json.Marshal(ipInfoObj)
	if err != nil {
		return "", err
	}
	return string(ipInfoJson), nil
}

// IsPrivateIP 判断是否为局域网 IP
/*
 * @param ip 需要判断的ip地址
 * @return bool true 内网 false 局域网
 */
func IsPrivateIP(ip net.IP) bool {
	privateIPBlocks := []string{
		"10.0.0.0/8",
		"172.16.0.0/12",
		"192.168.0.0/16",
	}

	for _, cidr := range privateIPBlocks {
		_, subnet, _ := net.ParseCIDR(cidr)
		if subnet.Contains(ip) {
			return true
		}
	}
	return false
}
