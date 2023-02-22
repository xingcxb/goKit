package strKit

import (
	"goKit/arrayKit"
	"strings"
)

const (
	EMPTY = ""  // 字符串常量：空字符串 ""
	SPACE = " " // 字符串常量：空格符 " "
)

// Splicing 字符串拼接
// @param str 待拼凑的字符串
func Splicing(str ...string) string {
	var newStr = strings.Builder{}
	for i := 0; i < len(str); i++ {
		newStr.WriteString(str[i])
	}
	return newStr.String()
}

// RemoveAll 去除字符串中指定的多个字符，如有多个则全部去除
// @param str 字符串
// @param chars 字符列表
// @return 去除后的字符
func RemoveAll(str string, chars ...string) string {
	sb := strings.Builder{}
	charStr := strings.Split(str, "")
	for _, s := range charStr {
		if false == arrayKit.Contains(chars, s) {
			sb.WriteString(s)
		}
	}
	return sb.String()
}

// MapParamsToUrlParams 请求链接中的params转字符串
// @param paramsMap 请求map类型的请求参数，通常来说请求参数都是字符串
func MapParamsToUrlParams(paramsMap map[string]string) string {
	urlParams := ""
	if paramsMap == nil {
		return ""
	}
	for k, v := range paramsMap {
		urlParams = Splicing(urlParams, k, "=", v, "&")
	}
	urlParams = urlParams[:len(urlParams)-1]
	return urlParams
}

// Reverse 反转字符串 例如：abcd =》dcba
// @param str – 被反转的字符串
// @return 反转后的字符串
func Reverse(str string) string {
	strArray := strings.Split(str, "")
	for i, j := 0, len(strArray)-1; i < j; i, j = i+1, j-1 {
		strArray[i], strArray[j] = strArray[j], strArray[i]
	}
	return strings.Join(strArray, "")
}
