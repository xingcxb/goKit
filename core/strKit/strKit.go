package strKit

import (
	"fmt"
	"github.com/xingcxb/goKit/core/arrayKit"
	"math"
	"strings"
	"unicode/utf8"
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
// @param str 原始字符串
// @param chars 要剔除的字符列表
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

// ReplaceIndex 通过下标替换值
// @param start 开始下标
// @param end 结束下标
// @param str 要替换字符串
// @param replaceStr 替换的字符串
// @return
func ReplaceIndex(start, end int, str, replaceStr string) string {
	return strings.Replace(str, string([]rune(str)[start:end]), replaceStr, 1)
}

// AutoReplaceMiddle 替换中间字符为*
//
//	如果为单字符的则无法隐藏eg:a  ,  a@qq.com
//
// @param str 待替换的字符
// @return 返回替换后的字符串
func AutoReplaceMiddle(str string) string {
	if len(str) < 2 {
		return str
	}
	// 要改变的字符串
	changeStr := ""
	// 邮箱尾缀
	mailSuffix := ""
	if strings.Contains(str, "@") {
		// 邮箱部分处理
		mailAddLength := strings.Index(str, "@")
		// 邮箱后缀
		mailSuffix = str[mailAddLength:]
		// 邮箱地址
		changeStr = str[:mailAddLength]
	} else {
		changeStr = str
	}
	length := utf8.RuneCountInString(changeStr)
	num := length/2 - 1
	length = length - num
	beginIndex := int(math.Ceil(float64(length / 2)))
	oldReStr := ""
	for i := 0; i < num+1; i++ {
		oldReStr = Splicing(oldReStr, "*")
	}
	return fmt.Sprintf("%v%v", ReplaceIndex(beginIndex, beginIndex+num+1, changeStr, oldReStr), mailSuffix)
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

// FirstUpper 首字母大写
// @param 要处理的字符
// @return 返回首字母大写的字符串
func FirstUpper(str string) string {
	return fmt.Sprintf("%v%v", strings.ToUpper(str[:1]), str[1:])
}
