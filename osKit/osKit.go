package osKit

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/host"
	"goKit/dateKit"
	"time"
)

type BootTimeInfo struct {
	BootTime string `json:"bootTime"` //开机时间
	RunTime  string `json:"runTime"`  // 运行时间
}

// BootTime 开机时间
func BootTime() string {
	timestamp, _ := host.BootTime()
	t := time.Unix(int64(timestamp), 0)
	return dateKit.ToDateTimeStr(t)
}

func OsInfo() {
	version, _ := host.KernelVersion()
	fmt.Println(version)

	platform, family, version, _ := host.PlatformInformation()
	if platform == "darwin" {
		platform = "macOS"
	}
	fmt.Println("platform:", platform)
	fmt.Println("family:", family)
	fmt.Println("version:", version)
}
