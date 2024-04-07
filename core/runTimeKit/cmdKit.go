package runTimeKit

import (
	"os/exec"
)

// ExecuteCmd 执行shell命令
/**
 * 如要执行 `ls -l /home/user` 传值应该是 name=ls，arg="-l","/home/user"
 * @param shell 命令
 * @param parameter 参数
 * @return string,err 终端返回的信息，错误信息
 */
func ExecuteCmd(shell string, parameter ...string) (string, error) {
	cmd := exec.Command(shell, parameter...)
	output, err := cmd.CombinedOutput()
	return string(output), err
}
