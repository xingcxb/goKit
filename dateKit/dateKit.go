package dateKit

import (
	"errors"
	"fmt"
	"time"
)

var (
	// DefaultLayout 格式化时间的默认模板
	DefaultLayout = "2006-01-02 15:04:05"
	// DateLayout 格式化时间的日期模版
	DateLayout = "2006-01-02"
)

// Today 获取当前时间字符串
func Today() string {
	return time.Now().Format(DefaultLayout)
}

// GetYear 获取当前时间的年
func GetYear() string {
	return fmt.Sprintf("%v", time.Now().Year())
}

// GetMonth 获取当前时间的月份
func GetMonth() string {
	return fmt.Sprintf("%v", int(time.Now().Month()))
}

// GetDay 获取当前时间的号数
func GetDay() string {
	return fmt.Sprintf("%v", time.Now().Day())
}

// MillisecondOfToStr 毫秒时间戳转字符串
// @param timeMillis 时间戳，毫秒数
func MillisecondOfToStr(timeMillis int64) string {
	return time.UnixMilli(timeMillis).Format(DefaultLayout)
}

// SecondOfToStr 秒·时间戳转换为字符串
// @param timeSecond 时间戳，秒数
func SecondOfToStr(timeSecond int64) string {
	return time.Unix(timeSecond, 0).Format(DefaultLayout)
}

// OffsetMillisecond 偏移毫秒数
// @param date – 日期
// @param offset – 偏移毫秒数，正数向未来偏移，负数向历史偏移
// @return 偏移后的日期
func OffsetMillisecond(dateTime time.Time, offset int) (time.Time, error) {
	return OffSet(dateTime, TimeMillisecond, offset)
}

// OffsetSecond 偏移秒数
// @param date – 日期
// @param offset – 偏移秒数，正数向未来偏移，负数向历史偏移
// @return 偏移后的日期
func OffsetSecond(dateTime time.Time, offset int) (time.Time, error) {
	return OffSet(dateTime, TimeSeconds, offset)
}

// OffsetMinute 偏移分钟
// @param date – 日期
// @param offset – 偏移分钟数，正数向未来偏移，负数向历史偏移
// @return 偏移后的日期
func OffsetMinute(dateTime time.Time, offset int) (time.Time, error) {
	return OffSet(dateTime, TimeMinute, offset)
}

// OffsetHour 偏移小时
// @param date – 日期
// @param offset – 偏移小时数，正数向未来偏移，负数向历史偏移
// @return 偏移后的日期
func OffsetHour(dateTime time.Time, offset int) (time.Time, error) {
	return OffSet(dateTime, TimeHour, offset)
}

// OffsetDay 偏移天
// @param date – 日期
// @param offset – 偏移天数，正数向未来偏移，负数向历史偏移
// @return 偏移后的日期
func OffsetDay(dateTime time.Time, offset int) (time.Time, error) {
	return OffSet(dateTime, TimeDay, offset)
}

// OffsetWeek 偏移周
// @param date – 日期
// @param offset – 偏移周数，正数向未来偏移，负数向历史偏移
// @return 偏移后的日期
func OffsetWeek(dateTime time.Time, offset int) (time.Time, error) {
	return OffSet(dateTime, TimeWeek, offset)
}

// OffSet 获取指定日期偏移指定时间后的时间，生成的偏移日期不影响原日期
//
// @param dateTime 基准日期<br/>
// @param timeUnit 偏移的粒度大小（秒、分、小时、天、周） dateKit下的TimeSeconds、TimeMinute、TimeHour、TimeDay、TimeWeek<br/>
// @param offset 偏移量，正数为向后偏移，负数为向前偏移<br/>
// @return 偏移后的日期，错误
func OffSet(dateTime time.Time, timeUnit TimeUnit, offset int) (time.Time, error) {
	if timeUnit == "" {
		return time.Time{}, errors.New("timeUnit参数丢失")
	}
	// 将时间转换为时间戳(毫秒)
	timeStamp := dateTime.UnixMilli()
	var addMillisecond int64 = 0
	switch timeUnit {
	case TimeMillisecond:
		// 毫秒
		addMillisecond = int64(offset)
	case TimeSeconds:
		// 秒数
		addMillisecond = int64(offset) * 1e3
	case TimeMinute:
		// 分钟
		addMillisecond = int64(offset) * 60 * 1e3
	case TimeHour:
		//小时
		addMillisecond = int64(offset) * 60 * 60 * 1e3
	case TimeDay:
		// 天
		addMillisecond = int64(offset) * 24 * 60 * 60 * 1e3
	case TimeWeek:
		// 周
		addMillisecond = int64(offset) * 7 * 24 * 60 * 60 * 1e3
	default:
		return time.Time{}, errors.New("timeUnit暂不支持")
	}
	newTimeStamp := timeStamp + addMillisecond
	// 将新的时间戳转换为时间
	return time.UnixMilli(newTimeStamp), nil
}
