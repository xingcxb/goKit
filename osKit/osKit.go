package osKit

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/host"
	"time"
)

type BootTimeInfo struct {
	BootTime string `json:"bootTime"` //开机时间
	RunTime  string `json:"runTime"`  // 运行时间
}

// BootTime 开机时间
func BootTime() {
	timestamp, _ := host.BootTime()
	t := time.Unix(int64(timestamp), 0)
	fmt.Println(t.Local().Format("2006-01-02 15:04:05"))
}

func OsInfo() {
	version, _ := host.KernelVersion()
	fmt.Println(version)

	platform, family, version, _ := host.PlatformInformation()
	fmt.Println("platform:", platform)
	fmt.Println("family:", family)
	fmt.Println("version:", version)
}
