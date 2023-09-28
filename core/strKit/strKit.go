package strKit

import (
	"fmt"
	"github.com/xingcxb/goKit/core/arrayKit"
	"github.com/xingcxb/goKit/core/regKit"
	"math"
	"strconv"
	"strings"
	"unicode/utf8"
)

const (
	EMPTY     = ""  // 字符串常量：空字符串 ""
	SPACE     = " " // 字符串常量：空格符 " "
	UNDERLINE = "_" // 字符串常量：下划线 _
)

// Length 获取字符串长度
/**
 * @param str 字符串
 * @return 返回字符串长度
 */
func Length(str string) int {
	return utf8.RuneCountInString(str)
}

// Splicing 字符串拼接
/**
 * 	@param str 待拼凑的字符串
 * 	@return 返回拼凑后的字符串
 */
func Splicing(str ...string) string {
	var newStr = strings.Builder{}
	for i := 0; i < len(str); i++ {
		newStr.WriteString(str[i])
	}
	return newStr.String()
}

// SubString 字符串截断
/**
 * @param s 原始字符串
 * @param start 开始位置 0 <= start < len(s)
 * @param end 结束位置
 * @return 返回截取后的字符串
 */
func SubString(s string, start int, end int) string {
	if start < 0 || start >= len(s) || end < 0 || end > len(s) || start > end {
		return s // 参数无效，返回原字符串
	}
	return s[start:end] // 使用切片操作获取子串
}

// Str2Int 字符串转int
/*
 * @param str 字符串
 */
func Str2Int(str, intType string) interface{} {
	switch intType {
	case "int8":
		i, _ := strconv.ParseInt(str, 10, 8)
		return int8(i)
	case "int16":
		i, _ := strconv.ParseInt(str, 10, 16)
		return int16(i)
	case "int32":
		i, _ := strconv.ParseInt(str, 10, 32)
		return int32(i)
	case "int64":
		i, _ := strconv.ParseInt(str, 10, 64)
		return i
	case "uint8":
		i, _ := strconv.ParseUint(str, 10, 8)
		return uint8(i)
	case "uint16":
		i, _ := strconv.ParseUint(str, 10, 16)
		return uint16(i)
	case "uint32":
		i, _ := strconv.ParseUint(str, 10, 32)
		return uint32(i)
	case "uint64":
		i, _ := strconv.ParseUint(str, 10, 64)
		return i
	default:
		i, _ := strconv.Atoi(str)
		return i
	}
}

// RemoveAll 去除字符串中指定的多个字符，如有多个则全部去除
/**
 * @param str 原始字符串
 * @param chars 要剔除的字符列表
 * @return 去除后的字符
 */
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
/**
 * @param paramsMap 请求map类型的请求参数，通常来说请求参数都是字符串
 * @return 返回拼凑后的字符串
 */
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
/**
 * @param str URL字符串参数
 * @return 返回map类型的参数
 */
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

// ReplaceIndex 通过下标指定区域替换值
/**
 * @param start 开始下标
 * @param end 结束下标
 * @param str 要替换字符串
 * @param replaceStr 替换的字符串
 * @return 返回替换后的字符串
 */
func ReplaceIndex(start, end int, str, replaceStr string) string {
	return strings.Replace(str, string([]rune(str)[start:end]), replaceStr, 1)
}

// AutoReplaceMiddle 替换中间字符为* 如果为单字符的则无法隐藏eg:a  ,  a@qq.com
/**
 * @param str 待替换的字符
 * @return 返回替换后的字符串
 */
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

// IsEnLetter 判断单个字符是否为英文，如果字符长度超出单个字符直接判定为false
/**
 * @param str 待判断字符串
 * @return 返回是否为英文
 */
func IsEnLetter(str string) bool {
	strLen := Length(str)
	if strLen > 2 {
		return false
	}
	sStr := StrToAscii(str, "")
	s, _ := strconv.ParseInt(sStr, 10, 64)
	return (s >= 97 && s <= 122) || (s >= 65 && s <= 90)
}

// Reverse 反转字符串 例如：abcd =》dcba
/**
 * @param str 被反转的字符串
 * @return 反转后的字符串
 */
func Reverse(str string) string {
	strArray := strings.Split(str, "")
	for i, j := 0, len(strArray)-1; i < j; i, j = i+1, j-1 {
		strArray[i], strArray[j] = strArray[j], strArray[i]
	}
	return strings.Join(strArray, "")
}

// FirstUpper 首字母大写
/**
 * @param str 要处理的字符
 * @return 返回首字母大写的字符串
 */
func FirstUpper(str string) string {
	if str == "" {
		return ""
	}
	return fmt.Sprintf("%v%v", strings.ToUpper(str[:1]), str[1:])
}

// FirstLower 首字母小写
/**
 * @param str 要处理的字符
 * @return 返回首字母大写的字符串
 */
func FirstLower(str string) string {
	if str == "" {
		return ""
	}
	return fmt.Sprintf("%v%v", strings.ToLower(str[:1]), str[1:])
}

// SplitterToHump 分割符转换为驼峰
/**
 * @param str 待处理的字符
 * @param splitter 分割符，默认为下划线 _
 * @return 返回驼峰字符串
 */
func SplitterToHump(str, splitter string) string {
	if splitter == "" {
		splitter = UNDERLINE
	}
	strs := strings.Split(str, splitter)
	newView := ""
	for _, s := range strs {
		newView = Splicing(newView, strings.Title(s))
	}
	return newView
}

// SliceToStr 切片转字符串，用逗号分隔
/**
 * 	@param strs 切片
 * 	@return 字符串
 */
func SliceToStr(strs []string) string {
	newStr := ""
	for _, str := range strs {
		newStr = Splicing(newStr, str, ",")
	}
	return newStr[:len(newStr)-1]
}

// CleanStrSymbol 清除字符串中的html标签
/**
 * @param str 字符串
 * @return 返回清除后的字符串
 */
func CleanStrSymbol(str string) string {
	return regKit.ReplaceAll("</?[^>]+>", "", str)
}

// StrToAscii 字符串转ASCII码，并用指定的分隔符分隔
/**
 * @param str 字符串
 * @param separator 返回值分隔符，如果为空则默认为逗号
 * @return 返回ASCII码
 */
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
/**
 * @param str ASCII码
 * @param separator 分隔符，如果为空则默认为逗号
 * @return 返回字符串
 */
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
