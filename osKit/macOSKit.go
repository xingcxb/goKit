// go:build darwin
package osKit

import (
	"encoding/json"
	"fmt"
	"github.com/shirou/gopsutil/v3/cpu"
)

func Cpu() {
	infos, _ := cpu.Info()
	for _, info := range infos {
		data, _ := json.MarshalIndent(info, "", " ")
		fmt.Print(string(data))
	}
}
