package pathKit

import "os"

// GetAbsolutePackagePath 获取当前调用函数的绝对路径文件夹
// @return 绝对路径, 错误信息
func GetAbsolutePackagePath() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return dir, nil
}
