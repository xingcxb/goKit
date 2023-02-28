package arrayKit

const (
	IndexNotFound = -1 // 数组中元素未找到的下标，值为-1
)

// IndexOf 返回数组中指定元素所在位置，未找到返回IndexNotFound
// @param strs 字符串数组
// @param char 被检查的元素
// @return 数组中指定元素所在位置，未找到返回IndexNotFound
func IndexOf(strs []string, char string) int {
	for i, str := range strs {
		if str == char {
			return i
		}
	}
	return IndexNotFound
}

// Contains 数组中是否包含元素
// @param strs 数组
// @param value 被检查的元素
// @return 是否包含
func Contains(strs []string, char string) bool {
	return IndexOf(strs, char) > IndexNotFound
}
