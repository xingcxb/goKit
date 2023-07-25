package cryptoKit

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

// Sha256 Sha256加密
/**
 * @param str 加密字符串
 * @return 加密结果 32个字符小写
 */
func Sha256(str string) string {
	sum := sha256.Sum256([]byte(str))
	return hex.EncodeToString(sum[:])
}

// Sha256ToUpper Sha256加密大写
/**
 * @param str 加密字符串
 * @return 加密结果 32个字符大写
 */
func Sha256ToUpper(str string) string {
	sumStr := Sha256(str)
	return strings.ToUpper(sumStr)
}
