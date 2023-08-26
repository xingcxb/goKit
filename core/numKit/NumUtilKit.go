package numKit

import "fmt"

const (
	UnitStringWan = "万"
	UnitStringYi  = "亿"
)

// FormatNumUnit 将数字格式化携带单位
/**
 * @param num 数字
 * @return string
 */
func FormatNumUnit(num int) string {
	if num < 10000 {
		//如果小于1万
		return fmt.Sprintf("%v", num)
	}
	if num < 100000000 {
		//如果大于1万小于1亿
		return fmt.Sprintf("%v%v", num/10000, UnitStringWan)
	}
	return fmt.Sprintf("%v%v", num/100000000, UnitStringYi)
}
