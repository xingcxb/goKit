package runTimeKit

import (
	"os/exec"
)

// ExecuteCmd 执行shell命令
// @param shell 命令
// @param parameter 参数
func ExecuteCmd(shell string, parameter ...string) (string, error) {
	cmd := exec.Command(shell, parameter...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(output), nil
}
