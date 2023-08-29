package arrayKit

import (
	"fmt"
	"sort"
	"strings"
)

const (
	IndexNotFound = -1 // 数组中元素未找到的下标，值为-1
)

// IndexOf 返回数组中指定元素所在位置，未找到返回IndexNotFound
/**
 * @param strs 字符串数组
 * @param char 被检查的元素
 * @return 数组中指定元素所在位置，未找到返回IndexNotFound
 */
func IndexOf(strs []string, char string) int {
	for i, str := range strs {
		if str == char {
			return i
		}
	}
	return IndexNotFound
}

// Contains 数组中是否包含元素
/**
 * @param strs 数组
 * @param value 被检查的元素
 * @return 包含返回true，否则返回false
 */
func Contains(strs []string, char string) bool {
	return IndexOf(strs, char) > IndexNotFound
}

// BinarySearchIndexOf 二分查找是否存在元素
func BinarySearchIndexOf(strs []string, char string) bool {
	// 首先对切片排序
	sort.Strings(strs)
	// 然后进行二分查找
	low := 0
	high := len(strs) - 1
	for low <= high {
		mid := (low + high) / 2
		if strs[mid] == char {
			return true
		} else if strs[mid] < char {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return false
}

// BubbleDescSort 冒泡排序 倒序
/**
 * @param values 待排序的字符串数组
 * @return 排序后的字符串数组
 */
func BubbleDescSort(strs []string) []string {
	n := len(strs)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if strs[j] < strs[j+1] {
				strs[j], strs[j+1] = strs[j+1], strs[j]
			}
		}
	}
	return strs
}

// BubbleAscSort 冒泡排序 正序
/**
 * @param values 待排序的字符串数组
 * @return 排序后的字符串数组
 */
func BubbleAscSort(strs []string) []string {
	n := len(strs)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if strs[j] > strs[j+1] {
				strs[j], strs[j+1] = strs[j+1], strs[j]
			}
		}
	}
	return strs
}

// JoinStringsInASCII 将map数据key以ASCII码从小到大排序后拼接
/**
 * @param data 待拼接的数据
 * @param sep 连接符
 * @param onlyValues 是否只包含参数值，true则不包含参数名，否则参数名和参数值均有
 * @param includeEmpty 是否包含空值，true则包含空值，否则不包含，注意此参数不影响参数名的存在
 * @param exceptKeys 被排除的参数名，不参与排序及拼接
 * @return 返回URL类型的参数字符串
 */
func JoinStringsInASCII(data map[string]string, sep string, onlyValues, includeEmpty bool, exceptKeys ...string) string {
	var list []string
	var keyList []string
	m := make(map[string]int)
	if len(exceptKeys) > 0 {
		for _, except := range exceptKeys {
			m[except] = 1
		}
	}
	for k := range data {
		if _, ok := m[k]; ok {
			continue
		}
		value := data[k]
		if !includeEmpty && value == "" {
			continue
		}
		if onlyValues {
			keyList = append(keyList, k)
		} else {
			list = append(list, fmt.Sprintf("%s=%s", k, value))
		}
	}
	if onlyValues {
		sort.Strings(keyList)
		for _, v := range keyList {
			list = append(list, data[v])
		}
	} else {
		sort.Strings(list)
	}
	return strings.Join(list, sep)
}
