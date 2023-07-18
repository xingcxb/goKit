package ipKit

import (
	"fmt"
	"github.com/xingcxb/goKit/core/ipKit/ipService"
	"net"
	"testing"
)

func TestIp(t *testing.T) {
	fmt.Println(ipService.GetIpInfoCZ("171.42.102.199"))
	fmt.Println(ipService.GetIpInfoCip("178.42.102.199"))
}

func TestLIp(t *testing.T) {
	ifaces, err := net.Interfaces()
	if err != nil {
		fmt.Println(err)
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp != 0 && iface.Flags&net.FlagLoopback == 0 {
			addrs, err := iface.Addrs()
			if err != nil {
				fmt.Println(err)
			}
			for _, addr := range addrs {
				ipNet, ok := addr.(*net.IPNet)
				if ok && !ipNet.IP.IsLoopback() {
					if ipNet.IP.To4() != nil {
						fmt.Println(ipNet.IP.String())
					}
				}
			}
		}
	}
	fmt.Println("end")
}

func TestLocalIp(t *testing.T) {
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("获取网络接口失败：", err)
		return
	}

	// 遍历所有网络接口
	for _, iface := range interfaces {
		// 排除回环接口和无效接口
		if iface.Flags&net.FlagLoopback != 0 || iface.Flags&net.FlagUp == 0 {
			continue
		}

		// 获取接口的地址列表
		addrs, err := iface.Addrs()
		if err != nil {
			fmt.Println("获取接口地址失败：", err)
			continue
		}

		// 遍历接口地址
		for _, addr := range addrs {
			// 检查是否为 IP 地址
			ipNet, ok := addr.(*net.IPNet)
			if !ok {
				continue
			}

			// 排除私有 IP 地址和 IPv6 地址
			if ipNet.IP.IsPrivate() || ipNet.IP.To4() == nil {
				continue
			}

			// 打印外网 IP 地址
			fmt.Println("外网 IP:", ipNet.IP.String())
		}
	}
}
