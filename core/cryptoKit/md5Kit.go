// Package cryptoKit 加密工具包
package cryptoKit

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

// Md5 Md5加密
/**
 * @param str 加密字符串
 * @return 加密结果 32个字符小写
 */
func Md5(str string) string {
	sum := md5.Sum([]byte(str))
	return hex.EncodeToString(sum[:])
}

// Md5ToUpper Md5加密大写
/**
 * @param str 加密字符串
 * @return 加密结果 32个字符大写
 */
func Md5ToUpper(str string) string {
	sumStr := Md5(str)
	return strings.ToUpper(sumStr)
}
