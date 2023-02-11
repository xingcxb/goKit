package strKit

import "strings"

// Splicing 字符串拼接
// @param str 待拼凑的字符串
func Splicing(str ...string) string {
	var newStr = strings.Builder{}
	for i := 0; i < len(str); i++ {
		newStr.WriteString(str[i])
	}
	return newStr.String()
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
