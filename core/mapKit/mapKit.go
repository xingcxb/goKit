package mapKit

import "github.com/xingcxb/goKit/core/randomKit"

// GetVByIndex 通过下标取map中的值
/*
 * @param m map[any]any 取值的map
 * @param index int 下标
 */
func GetVByIndex[K comparable, V any](m map[K]V, index int) any {
	keys := make([]any, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	return keys[index]
}

// GetRandomValue 随机获取map中的value
/*
 * @param m map[any]any
 */
func GetRandomValue[K comparable, V any](m map[K]V) V {
	keys := make([]any, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	index := randomKit.RandomLong(len(keys))
	randomKey := keys[index]
	return m[randomKey]
}
