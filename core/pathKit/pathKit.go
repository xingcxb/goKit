package pathKit

import (
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

// GetAbsolutePackagePath 获取当前调用函数的绝对路径文件夹
// @return 绝对路径, 错误信息
func GetAbsolutePackagePath() (string, error) {
	dir, err := getCurrentAbPathByExecutable()
	if err != nil {
		return "", err
	}
	if strings.Contains(dir, getTmpDir()) {
		return getCurrentAbPathByCaller()
	}
	return dir, err
}

// 获取系统临时目录，兼容go run
func getTmpDir() string {
	dir := os.Getenv("TEMP")
	if dir == "" {
		dir = os.Getenv("TMP")
	}
	res, _ := filepath.EvalSymlinks(dir)
	return res
}

// 获取当前执行文件绝对路径(go build)
func getCurrentAbPathByExecutable() (string, error) {
	_, err := os.Executable()
	if err != nil {
		return "", err
	}
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return dir, nil
}

// 获取当前执行文件绝对路径（go run）
func getCurrentAbPathByCaller() (string, error) {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath, nil
}
