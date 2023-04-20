package numKit

import "fmt"

const (
	UNIT_STRING_WAN = "万"
	UNIT_STRING_YI  = "亿"
)

// FormatNumUnit 将数字格式化携带单位
// @param num 数字
func FormatNumUnit(num int) string {
	if num < 10000 {
		//如果小于1万
		return fmt.Sprintf("%v", num)
	}
	if num < 100000000 {
		//如果大于1万小于1亿
		return fmt.Sprintf("%v%v", num/10000, UNIT_STRING_WAN)
	}
	return fmt.Sprintf("%v%v", num/100000000, UNIT_STRING_YI)
}
