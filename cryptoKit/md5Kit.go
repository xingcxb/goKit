// Package cryptoKit 加密工具包
package cryptoKit

import (
	"crypto/md5"
	"encoding/hex"
)

// Md5 加密
// @param 加密字符串
// @return 加密结果 32个字符小写
func Md5(str string) string {
	sum := md5.Sum([]byte(str))
	return hex.EncodeToString(sum[:])
}
