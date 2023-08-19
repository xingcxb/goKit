package core

import (
	"fmt"
	"github.com/xingcxb/goKit/core/osKit"
	"testing"
)

// ---os---

// 开机时间
func TestBootTime(t *testing.T) {
	fmt.Println(osKit.GetBootTime())
}

// 获取系统信息
func TestOsInfo(t *testing.T) {
	fmt.Println(osKit.GetOsInfo())
}

// 获取本地ip
func TestGetLocalIp(t *testing.T) {
	fmt.Println(osKit.GetLocalIp())
}

// 获取硬盘容量信息
func TestDiskSpaceInfo(t *testing.T) {
	fmt.Println(osKit.GetDiskSpaceInfo())
}

// ---macos---

// 获取cpu信息
func TestGetCpuInfo(t *testing.T) {
	fmt.Println(osKit.GetCpuInfo())
}

// 获取cpu使用率
func TestGetCPUPercent(t *testing.T) {
	fmt.Println(osKit.GetCPUPercent())
}

// 获取系统负载
func TestGetSysLoad(t *testing.T) {
	fmt.Println(osKit.GetSysLoad())
}

// 获取内存信息
func TestGetMemInfo(t *testing.T) {
	fmt.Println(osKit.GetMemoryInfo())
}

// 获取磁盘信息
func TestGetDiskInfo(t *testing.T) {
	fmt.Println(osKit.GetDiskInfo())
}

// 获取风扇信息
func TestGetMacOSFanSpeed(t *testing.T) {
	fmt.Println(osKit.GetFanSpeed())
}

// 获取温度信息
func TestGetTemperature(t *testing.T) {
	fmt.Println(osKit.GetTemperature())
}

// 获取电源信息
func TestPower(t *testing.T) {
	fmt.Println(osKit.GetPower())
}

// 获取电压信息
func TestVoltage(t *testing.T) {
	fmt.Println(osKit.GetVoltage())
}

// 获取电池信息
func TestGetBattery(t *testing.T) {
	fmt.Println(osKit.GetBattery())
}

// 获取电流信息
func TestGetCurrent(t *testing.T) {
	fmt.Println(osKit.GetCurrent())
}
