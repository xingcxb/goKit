package cryptoKit

import "encoding/base64"

// Base64Encode Base64加密
// @param str 加密字符串
// @return 加密结果
func Base64Encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

// Base64Decode Base64解密
// @param str 解密字符串
// @return 解密结果
func Base64Decode(str string) (string, error) {
	_b, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return "", err
	}
	return string(_b), nil
}
