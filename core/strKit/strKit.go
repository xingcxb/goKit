package strKit

import (
	"fmt"
	"github.com/xingcxb/goKit/core/arrayKit"
	"math"
	"strconv"
	"strings"
	"unicode/utf8"
)

const (
	EMPTY = ""  // 字符串常量：空字符串 ""
	SPACE = " " // 字符串常量：空格符 " "
)

// Length 获取字符串长度（使用len方法获取字符串的字节数）
// @param 字符串
// @return 返回字符串长度
func Length(str string) int {
	return utf8.RuneCountInString(str)
}

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

// StrParamsToMapParams URL字符串参数转map
// @param str URL字符串参数
// @return 返回map类型的参数
func StrParamsToMapParams(str string) map[string]string {
	paramsMap := make(map[string]string)
	if str == "" {
		return paramsMap
	}
	params := strings.Split(str, "&")
	for _, param := range params {
		paramArr := strings.Split(param, "=")
		paramsMap[paramArr[0]] = paramArr[1]
	}
	return paramsMap
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
	length := Length(changeStr)
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

// SliceToStr 切片转字符串，用逗号分隔
// @param strs 切片
// @return 字符串
func SliceToStr(strs []string) string {
	newStr := ""
	for _, str := range strs {
		newStr = Splicing(newStr, str, ",")
	}
	return newStr[:len(newStr)-1]
}

// CleanStrSymbol 清除字符串中的html标签
func CleanStrSymbol(str string) string {
	return strings.ReplaceAll(str, "</?[^>]+>", "")
}

// StrToAscii 字符串转ASCII码
// @param str 字符串
// @param separator 分隔符，如果为空则默认为逗号
func StrToAscii(str, separator string) string {
	StrAscii := ""
	if separator == "" {
		separator = ","
	}
	strChar := []byte(str)
	for i := 0; i < len(strChar); i++ {
		StrAscii = Splicing(StrAscii, strconv.Itoa(int(strChar[i])), separator)
	}
	return StrAscii[:len(StrAscii)-1]
}

// AsciiToStr ASCII码转字符串
// @param str ASCII码
// @param separator 分隔符，如果为空则默认为逗号
func AsciiToStr(str, separator string) string {
	if separator == "" {
		// 如果没有传递分隔符，默认采用逗号
		separator = ","
	}
	// 将字符串转换为字符串数组
	strArr := strings.Split(str, separator)
	newStr := ""
	for _, s := range strArr {
		// 将字符串转换为int类型
		i, _ := strconv.Atoi(s)
		// 将int类型的ascii转换为字符类型
		newStr = Splicing(newStr, string(rune(i)), "")
	}
	return newStr
}
