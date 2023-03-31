package osKit

import (
	"fmt"
	"testing"
)

func TestCpu(t *testing.T) {
	fmt.Println("开机时间：", BootTime())
	osInfo, err := OsInfo()
	fmt.Println("系统信息：", osInfo, err)
	cpuInfo, err := GetCpuInfo()
	fmt.Println("Cpu基础信息：", cpuInfo, err)
	cpuPercent, err := GetCPUPercent()
	fmt.Println("获取CPU的占用率：", cpuPercent, err)
	sysLoad, err := GetSysLoad()
	fmt.Println("系统负载：", sysLoad, err)
	memory, err := GetMemoryInfo()
	fmt.Println("内存信息：", memory, err)
	diskInfo, err := GetDiskInfo()
	fmt.Println("硬盘信息：", diskInfo, err)
	diskSpaceInfo, err := DiskSpaceInfo()
	fmt.Println("硬盘空间信息：", diskSpaceInfo, err)
	fan, err := GetMacOSFanSpeed()
	fmt.Println("风扇转速：", fan, err)
	power, err := GetPower()
	fmt.Println("电源信息：", power, err)
	voltage, err := GetVoltage()
	fmt.Println("电压信息：", voltage, err)
	temperature, err := GetTemperature()
	fmt.Println("温度信息：", temperature, err)
	battery, err := GetBattery()
	fmt.Println("电池信息：", battery, err)
	current, err := GetCurrent()
	fmt.Println("当前信息：", current, err)
}
