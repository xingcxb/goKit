// Package dateKit 处理中国农历
package dateKit

// LeapYear 判断是否为闰年
/**
 * @param year 年份
 * @return 闰年返回true 否则返回false
 */
func LeapYear(year int) bool {
	if year%400 == 0 {
		return true
	} else if year%100 == 0 {
		return false
	} else if year%4 == 0 {
		return true
	}
	return false
}
