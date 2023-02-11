package dateKit

import "time"

// Year 获取当前时间的年
func Year() string {
	return string(time.Now().Year())
}
