package osKit

import (
	"encoding/json"
	"fmt"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/xingcxb/goKit/core/dateKit"
	"net"
)

type BootTimeInfo struct {
	BootTime string `json:"bootTime"` //开机时间
	RunTime  string `json:"runTime"`  // 运行时间
}

// GetBootTime 开机时间
func GetBootTime() BootTimeInfo {
	timestamp, _ := host.BootTime()
	dateTimeStr := dateKit.SecondOfToStr(int64(timestamp))
	useTimestampS, _ := host.Uptime()
	return BootTimeInfo{
		BootTime: dateTimeStr,
		RunTime:  fmt.Sprintf("%v", useTimestampS),
	}
}

// GetOsInfo 获取系统信息
func GetOsInfo() (string, error) {
	hostInfo, err := host.Info()
	if err != nil {
		return "", err
	}
	if hostInfo.OS == "darwin" {
		hostInfo.OS = "macOS"
		hostInfo.Platform = "macOS"
	}
	_b, err := json.Marshal(hostInfo)
	if err != nil {
		return "", err
	}
	return string(_b), nil
}

// GetDiskSpaceInfo 获取硬盘容量信息
func GetDiskSpaceInfo() (string, error) {
	diskSpaceInfo, err := disk.IOCounters()
	if err != nil {
		return "", err
	}
	_b, err := json.Marshal(diskSpaceInfo)
	return string(_b), err
}

// GetLocalIp 获取本地ip
func GetLocalIp() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	for _, addr := range addrs {
		ipAddr, ok := addr.(*net.IPNet)
		if !ok {
			continue
		}
		if ipAddr.IP.IsLoopback() {
			continue
		}
		if !ipAddr.IP.IsGlobalUnicast() {
			continue
		}
		return ipAddr.IP.String(), nil
	}
	return "", nil
}
