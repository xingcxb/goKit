package runTimeKit

import (
	"fmt"
	"testing"
)

func TestShell(t *testing.T) {
	fmt.Println(ExecuteCmd("curl", "cip.cc/172.42.102.198"))
}
