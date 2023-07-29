package regKit

import (
	"regexp"
)

// Get 获取匹配的字符串
/**
 * @params regex 正则表达式
 * @params content 待匹配的字符串
 * @params index 匹配到内容的字符串 下标从0开始
 * @return 匹配后得到的字符串，未匹配返回空字符串
 */
func Get(regex, content string, index int) string {
	if index-1 < 0 {
		index = 0
	}
	strs := FindAll(regex, content)
	if len(strs) < index {
		return ""
	}
	return strs[index-1]
}

// FindAll 取得内容中匹配的所有结果
/**
 * @params regex 正则表达式
 * @params content 待匹配的字符串
 * @return 匹配后的字符串切片，未匹配返回空切片
 */
func FindAll(regex, content string) []string {
	if regex == "" || content == "" {
		// 如果正则表达式或内容都为空，则返回空数组
		return []string{}
	}
	compile, err := regexp.Compile(regex)
	if err != nil {
		return []string{}
	}
	return compile.FindAllString(content, -1)
}

// Index 返回第一个匹配的字符串的起始位置，如果没有匹配则返回-1(下标从0开始)
/**
 * @params regex 正则表达式
 * @params content 待匹配的字符串
 * @return 位置
 */
func Index(regex, content string) int {
	if regex == "" || content == "" {
		// 如果正则表达式或内容都为空，则返回-1
		return -1
	}
	compile, err := regexp.Compile(regex)
	if err != nil {
		return -1
	}
	findStringIndexs := compile.FindStringIndex(content)
	if len(findStringIndexs) > 0 {
		return findStringIndexs[0]
	} else {
		return -1
	}
}

// IsMatch 指定内容中是否有表达式匹配的内容
/**
 * @params regex 正则表达式
 * @params content 待匹配的字符串
 * @return 匹配返回true;不匹配返回false
 */
func IsMatch(regex, content string) bool {
	if regex == "" || content == "" {
		// 如果正则表达式或内容都为空，则返回false
		return false
	}
	compile, err := regexp.Compile(regex)
	if err != nil {
		return false
	}
	return compile.MatchString(content)
}

// Deprecated: Contains函数废弃，请使用IsMatch
/**
 * @params regex 正则表达式
 * @params content 待匹配的字符串
 * @return 匹配返回true;不匹配返回false
 */
func Contains(regex, content string) bool {
	if regex == "" || content == "" {
		// 如果正则表达式或内容都为空，则返回false
		return false
	}
	compile, err := regexp.Compile(regex)
	if err != nil {
		return false
	}
	return compile.MatchString(content)
}

// Count 计算指定字符串中，匹配pattern的个数
/**
 * @params regex 正则表达式
 * @params content 待匹配的字符串
 * @return 匹配个数
 */
func Count(regex, content string) int {
	if regex == "" || content == "" {
		// 如果正则表达式或内容都为空，则返回0
		return 0
	}
	compile, err := regexp.Compile(regex)
	if err != nil {
		return 0
	}
	return len(compile.FindAllString(content, -1))
}

// ReplaceAll 替换所有匹配的字符串
/**
 * @params regex 正则表达式
 * @params replaceStr 替换的字符串
 * @params content 待匹配的字符串
 * @return 替换后的字符串
 */
func ReplaceAll(regex, replaceStr, content string) string {
	if regex == "" || content == "" {
		// 如果正则表达式或内容都为空，则返回0
		return ""
	}
	compile, err := regexp.Compile(regex)
	if err != nil {
		return ""
	}
	return compile.ReplaceAllString(content, replaceStr)
}
