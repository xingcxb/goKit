package osKit

import (
	"fmt"
	"testing"
)

func TestCpu(t *testing.T) {
	fmt.Println(BootTime())
	fmt.Println(OsInfo())
	fmt.Println(GetCpuInfo())
	fmt.Println(GetCPUPercent())
	fmt.Println(GetSysLoad())
	fmt.Println(GetMemoryInfo())
	fmt.Println(GetDiskInfo())
	fmt.Println(GetMacOSFanSpeed())
	fmt.Println(GetPower())
	fmt.Println(GetVoltage())
	fmt.Println(GetTemperature())
}
