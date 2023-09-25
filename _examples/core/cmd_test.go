package core

import (
	"fmt"
	"github.com/xingcxb/goKit/core/runTimeKit"
	"testing"
)

func TestShell(t *testing.T) {
	fmt.Println(runTimeKit.ExecuteCmd("curl", "cip.cc/172.42.102.198"))
}
