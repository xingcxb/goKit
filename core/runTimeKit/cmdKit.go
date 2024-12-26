package runTimeKit

import (
	"fmt"
	"github.com/xingcxb/goKit/core/dateKit"
	"os/exec"
	"time"
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

// DelayExecCmd 延迟执行命令
/*
 * @param nextSecond 下次执行秒数
 * @param command 执行的命令
 */
func DelayExecCmd(nextSecond int, command string) (string, error) {
	timeAfter, _ := dateKit.OffsetSecond(time.Now(), nextSecond)
	atCmd := fmt.Sprintf("echo '%s' | at %s", command, timeAfter)
	output, err := ExecuteCmd(atCmd)
	return string(output), err
}
