package osKit

import (
	"encoding/json"
	"fmt"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/xingcxb/goKit/core/dateKit"
)

type BootTimeInfo struct {
	BootTime string `json:"bootTime"` //开机时间
	RunTime  string `json:"runTime"`  // 运行时间
}

// BootTime 开机时间
func BootTime() BootTimeInfo {
	timestamp, _ := host.BootTime()
	dateTimeStr := dateKit.SecondOfToStr(int64(timestamp))
	useTimestampS, _ := host.Uptime()
	return BootTimeInfo{
		BootTime: dateTimeStr,
		RunTime:  fmt.Sprintf("%v", useTimestampS),
	}
}

// OsInfo 获取系统信息
func OsInfo() (string, error) {
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

// DiskSpaceInfo 获取硬盘容量信息
func DiskSpaceInfo() (string, error) {
	diskSpaceInfo, err := disk.IOCounters()
	if err != nil {
		return "", err
	}
	_b, err := json.Marshal(diskSpaceInfo)
	return string(_b), err
}
