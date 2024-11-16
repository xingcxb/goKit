//go:build darwin

package osKit

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dkorunic/iSMC/smc"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/load"
	"github.com/shirou/gopsutil/v4/mem"
	"time"
)

// CpuInfos cpu信息集合
type CpuInfos struct {
	Cores     string //核心数
	ModelName string // cpu牌子
}

// GetCpuInfo 获取CPU信息
func GetCpuInfo() (string, error) {
	infos, _ := cpu.Info()
	var cpuInfos []CpuInfos
	// 可能存在多个CPU
	for i := 0; i < len(infos); i++ {
		cpuInfo := CpuInfos{
			Cores:     fmt.Sprintf("%v", infos[i].Cores),
			ModelName: infos[i].ModelName,
		}
		cpuInfos = append(cpuInfos, cpuInfo)
	}
	cpuInfosByte, err := json.Marshal(cpuInfos)
	if err != nil {
		return "", err
	}
	return string(cpuInfosByte), nil
}

// GetCPUPercent 获取CPU的占用率
func GetCPUPercent() (string, error) {
	cpuPercent, err := cpu.Percent(time.Second, false)
	if err != nil {
		return "", err
	}
	if len(cpuPercent) > 0 {
		return fmt.Sprintf("%v", cpuPercent[0]), nil
	}
	return "", errors.New("none cpu")
}

// GetSysLoad 系统负载
/**
 * 系统负荷是指系统中正在使用的CPU数量，或者说正在使用的CPU百分比。
 * 系统负荷为0.0，意味着系统中所有的CPU都处于空闲状态。
 * 系统负荷为1.0，意味着系统中有一个CPU被完全使用，而其它CPU处于空闲状态。
 * 系统负荷为2.0，意味着系统中有2个CPU被完全使用，而其它CPU处于空闲状态。
 * 系统负荷为4.0，意味着系统中有4个CPU被完全使用，而其它CPU处于空闲状态。
 */
func GetSysLoad() (string, error) {
	sysLoad, _ := load.Avg()
	_loadByte, err := json.Marshal(sysLoad)
	if err != nil {
		return "", err
	}
	return string(_loadByte), nil
}

// GetMemoryInfo 获取内存信息
func GetMemoryInfo() (string, error) {
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		return "", err
	}
	memInfoByte, err := json.Marshal(memInfo)
	if err != nil {
		return "", err
	}
	return string(memInfoByte), nil
}

// GetDiskInfo 获取硬盘信息
func GetDiskInfo() (string, error) {
	parts, err := disk.Partitions(true)
	if err != nil {
		return "", err
	}
	partsByte, err := json.Marshal(parts)
	if err != nil {
		return "", err
	}
	return string(partsByte), nil
}

// AllFansInfo 风扇所有信息
type AllFansInfo struct {
	FanCount int       `json:"fanCount"` // 风扇个数
	FanInfos []FanInfo `json:"fanInfos"` // 风扇信息
}

// FanInfo 单个风扇信息
type FanInfo struct {
	CurrentSpeed string `json:"currentSpeed"` // 当前转速
	MinimalSpeed string `json:"minimalSpeed"` // 最小转速
	MaximumSpeed string `json:"maximumSpeed"` // 最大转速
	TargetSpeed  string `json:"targetSpeed"`  // 目标转速
}

const (
	FanCountKey = "Fan Count" // 风扇个数
)

// GetFanSpeed 获取macOS风扇转速
func GetFanSpeed() (string, error) {
	osFans := smc.GetFans()
	allFansInfo := AllFansInfo{}
	// 获取风扇个数
	fansCount := osFans[FanCountKey].(map[string]interface{})
	allFansInfo.FanCount = int(fansCount["value"].(uint32))
	// 设置默认最低为一个风扇
	fansInfo := make([]FanInfo, 0)
	for fc := 1; fc <= allFansInfo.FanCount; fc++ {
		fanInfo := FanInfo{}
		// 当前转速
		cSpeedMap := osFans[fmt.Sprintf("Fan %v Current Speed", fc)].(map[string]interface{})
		fanInfo.CurrentSpeed = cSpeedMap["value"].(string)
		// 最小转速
		minSpeedMap := osFans[fmt.Sprintf("Fan %v Minimal Speed", fc)].(map[string]interface{})
		fanInfo.MinimalSpeed = minSpeedMap["value"].(string)
		// 最大转速
		maxSpeedMap := osFans[fmt.Sprintf("Fan %v Maximum Speed", fc)].(map[string]interface{})
		fanInfo.MaximumSpeed = maxSpeedMap["value"].(string)
		// 目标转速
		tSpeedMap := osFans[fmt.Sprintf("Fan %v Target Speed", fc)].(map[string]interface{})
		fanInfo.TargetSpeed = tSpeedMap["value"].(string)
		fansInfo = append(fansInfo, fanInfo)
	}
	allFansInfo.FanInfos = fansInfo
	allFansInfoByte, err := json.Marshal(allFansInfo)
	if err != nil {
		return "", err
	}
	return string(allFansInfoByte), nil

}

// GetTemperature 获取设备温度
func GetTemperature() (string, error) {
	temperature := smc.GetTemperature()
	temperatureJson, err := json.Marshal(temperature)
	return string(temperatureJson), err
}

// GetPower 获取电源信息
func GetPower() (string, error) {
	power := smc.GetPower()
	powerJson, err := json.Marshal(power)
	return string(powerJson), err
}

// GetVoltage 获取电压信息
func GetVoltage() (string, error) {
	voltage := smc.GetVoltage()
	voltageJson, err := json.Marshal(voltage)
	return string(voltageJson), err
}

// GetBattery 获取电池信息
func GetBattery() (string, error) {
	battery := smc.GetBattery()
	batteryJson, err := json.Marshal(battery)
	return string(batteryJson), err
}

// GetCurrent 获取电流信息
func GetCurrent() (string, error) {
	current := smc.GetCurrent()
	currentJson, err := json.Marshal(current)
	return string(currentJson), err
}
