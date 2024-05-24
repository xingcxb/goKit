package arrayKit

// GetVByIndex 通过下标取map中的值
/*
 * @param m map[any]any 取值的map
 * @param index int 下标
 */
func GetVByIndex(m map[any]any, index int) any {
	_v := make([]any, len(m))
	for k, v := range m {
		_v[k.(int)] = v
	}
	return _v[index]
}
