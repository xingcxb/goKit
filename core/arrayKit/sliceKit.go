package arrayKit

// Compare 比较两个字符串数组是否相等
/**
 * @param a 字符串数组a
 * @param b 字符串数组b
 * @return true:相等 false:不相等
 */
func Compare(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// SliceDiff 切片差集
/*
 * @param originalSlice 待比较的切片
 * @param compareSlice 比较切片
 * @return added, deleted 新增的数据[基于待比较切片而言], 删除的数据[基于待比较切片而言]
 */
func SliceDiff(originalSlice, compareSlice []string) (added, deleted []string) {
	coreMap := make(map[string]bool)
	compareMap := make(map[string]bool)
	for _, v := range originalSlice {
		coreMap[v] = true
	}
	for _, v := range compareSlice {
		compareMap[v] = true
		if !coreMap[v] {
			added = append(added, v)
		}
	}
	for k := range coreMap {
		if !compareMap[k] {
			deleted = append(deleted, k)
		}
	}
	return added, deleted
}
