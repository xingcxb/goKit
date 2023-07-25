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
