package ipKit

import (
	"encoding/json"
	"goKit/core/ipKit/ipService"
)

// GetIpAddressInfo 获取ip物理地址信息
// @param ip 需要查询的ip地址
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
