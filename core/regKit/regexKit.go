package regKit

import (
	"regexp"
)

// core 正则核心方法
//func core(ctx context.Context, regex, content string) bool {
//	compile, err := regexp.Compile(regex)
//	if err != nil {
//		return false
//	}
//}

// FindAll 取得内容中匹配的所有结果
func FindAll(regex, content string) []string {
	if regex == "" || content == "" {
		// 如果正则表达式或内容都为空，则返回空数组
		return []string{}
	}
	compile, err := regexp.Compile(regex)
	if err != nil {
		return nil
	}
	return compile.FindAllString(content, -1)
}

// Index 返回第一个匹配的字符串的起始位置，如果没有匹配则返回-1(下标从0开始)
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

// IsMatch 判断内容是否匹配正则表达式
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

// Contains 判断内容是否包含正则表达式
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
