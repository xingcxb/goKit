package cryptoKit

import (
	"crypto/hmac"
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

// Hmac256 Hmac256加密
/*
 * @param str 加密字符串
 * @param secret 加密密钥
 * @return 加密结果
 */
func Hmac256(str, secret string) string {
	// 创建 HMAC 哈希对象
	hmac256 := hmac.New(sha256.New, []byte(secret))
	// 计算 HMAC-SHA256 哈希值
	hmac256.Write([]byte(str))
	hash := hmac256.Sum(nil)
	return Base64Encode(string(hash))
}
