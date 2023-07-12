package dateKit

import (
	"errors"
	"fmt"
	"github.com/xingcxb/goKit/core/strKit"
	"strconv"
	"time"
)

var (
	// DefaultLayout 格式化时间的默认模板
	// DefaultLayout = "2006-01-02 15:04:05"
	// DateLayout 格式化时间的日期模版
	// DateLayout = "2006-01-02"

	// DateLayoutYM 格式化时间的年月模版
	DateLayoutYM = "2006-01"
	// DateLayoutYMD 格式化时间的年月日模版
	DateLayoutYMD = "2006-01-02"
	// DateLayoutHMS 格式化时间的时分秒模版
	DateLayoutHMS = "15:04:05"
	// DateLayoutYMDHMS 格式化时间的年月日时分秒模版
	DateLayoutYMDHMS = "2006-01-02 15:04:05"
	// DateLayoutYMDHMSS 格式化时间的年月日时分秒毫秒模版
	DateLayoutYMDHMSS = "2006-01-02 15:04:05.000"
	// DateLayoutPureYMD 格式化时间的纯年月日模版
	DateLayoutPureYMD = "20060102"
	// DateLayoutPureYMDHMS 格式化时间的纯年月日时分秒模版
	DateLayoutPureYMDHMS = "20060102150405"
	// DateLayoutPureYMDHMSS 格式化时间的纯时分秒毫秒模版
	DateLayoutPureYMDHMSS = "20060102150405000"
)

// ============获取时间================

// Now 当前时间，格式 yyyy-MM-dd HH:mm:ss
// @return 当前时间的标准形式字符串
func Now() string {
	return time.Now().Format(time.DateTime)
}

// Today 获取当前日期，格式 yyyy-MM-dd
// @return 当前时间的标准形式字符串
func Today() string {
	return time.Now().Format(time.DateOnly)
}

// BeginOfDay 获取某天的开始时间
// @param 日期
// @return 返回输入日期的开始时间
func BeginOfDay(date time.Time) string {
	beginDate := Format(date, time.DateOnly)
	return strKit.Splicing(beginDate, " 00:00:00")
}

// EndOfDay 获取某天的结束时间
// @param 日期
// @return 返回输入日期的结束时间
func EndOfDay(date time.Time) string {
	endDate := Format(date, time.DateOnly)
	return strKit.Splicing(endDate, " 23:59:59")
}

// GetYear 获得年的部分
// @param dateTime 日期
// @return 返回年份
func GetYear(dateTime time.Time) string {
	return fmt.Sprintf("%v", dateTime.Year())
}

// GetMonth 获取指定时间的月份
// @param dateTime 日期
// @return 返回月份
func GetMonth(dateTime time.Time) string {
	return fmt.Sprintf("%v", int(dateTime.Month()))
}

// GetDay 获取当前时间的号数
// @param dateTime 日期
// @return 返回号数
func GetDay(dateTime time.Time) string {
	return fmt.Sprintf("%v", dateTime.Day())
}

// Quarter 获得指定日期所属季度，从1开始计数 公历形式的划分3-5为春
// @param dateTime 日期
// @return 返回第几个季度 1：春、2：夏、3：秋、4：冬
func Quarter(dateTime time.Time) string {
	month := GetMonth(dateTime.Local())
	switch month {
	case "3", "4", "5":
		// 春
		return "1"
	case "6", "7", "8":
		// 夏
		return "2"
	case "9", "10", "11":
		// 秋
		return "3"
	case "12", "1", "2":
		// 冬
		return "4"
	default:
		// 不属于任何季度，直接返回-1
		return "-1"
	}
}

// Yesterday 昨天
// @return 昨天
func Yesterday() (time.Time, error) {
	return OffsetDay(time.Now(), -1)
}

// Tomorrow 明天
// @return 明天
func Tomorrow() (time.Time, error) {
	return OffsetDay(time.Now(), 1)
}

// LastWeek 上周
// @return 上周
func LastWeek() (time.Time, error) {
	return OffsetWeek(time.Now(), -1)
}

// NextWeek 下周
// @return 下周
func NextWeek() (time.Time, error) {
	return OffsetWeek(time.Now(), 1)
}

// LengthOfMonth 获取指定日期的月份的天数
// @param dateTime 日期
func LengthOfMonth(dateTime time.Time) int {
	return time.Date(dateTime.Year(), dateTime.Month()+1, 0, 0, 0, 0, 0, dateTime.Location()).Day()
}

// LengthOfYear 获取指定日期的年份的天数
// @param dateTime 日期
func LengthOfYear(dateTime time.Time) int {
	if IsLeapYear(dateTime) {
		return 366
	}
	return 365
}

// IsLeapYear 判断是否是闰年
// @param dateTime 日期
// @return 是否是闰年
func IsLeapYear(dateTime time.Time) bool {
	year := dateTime.Year()
	return (year%4 == 0 && year%100 != 0) || year%400 == 0
}

// ======================== 格式化日期 ========================

// FormatDateTime 时间转换为默认格式 yyyy-MM-dd HH:mm:ss
// @param dateTime 日期
// @return "yyyy-MM-dd HH:mm:ss" 格式字符串
func FormatDateTime(dateTime time.Time) string {
	return Format(dateTime, time.DateTime)
}

// FormatDate 时间转换为默认格式 yyyy-MM-dd
// @param dateTime 日期
// @return "yyyy-MM-dd" 格式字符串
func FormatDate(dateTime time.Time) string {
	return Format(dateTime, time.DateOnly)
}

// Format 时间转换为字符串
// @param dateTime 日期
// @return 格式化后字符串
func Format(dateTime time.Time, format string) string {
	return dateTime.Local().Format(format)
}

// ParseDate 日期时间字符串转time类型
// @param str 日期时间字符串 yyyy-MM-dd
// @return 返回时间类型数据
func ParseDate(str string) time.Time {
	t, err := time.ParseInLocation(time.DateOnly, str, time.Local)
	if nil == err && !t.IsZero() {
		return t
	}
	return time.Time{}
}

// ParseDateTime 日期时间字符串转time类型
// @param str 日期时间字符串 yyyy-MM-dd HH:mm:ss
// @return 返回时间类型数据
func ParseDateTime(str string) time.Time {
	t, err := time.ParseInLocation(time.DateTime, str, time.Local)
	if nil == err && !t.IsZero() {
		return t
	}
	return time.Time{}
}

// DateTimeStrToMillStr 日期时间字符串转毫秒字符串
// @param str 日期时间字符串 yyyy-MM-dd HH:mm:ss
func DateTimeStrToMillStr(str string) string {
	t := ParseDateTime(str)
	// 获取时间对象对应的 Unix 时间戳，并将其转换为字符串形式
	return strconv.FormatInt(t.UnixMilli(), 10)
}

// DateTimeStrToSecondStr 日期时间字符串转秒字符串
// @param str 日期时间字符串 yyyy-MM-dd HH:mm:ss
func DateTimeStrToSecondStr(str string) string {
	t := ParseDateTime(str)
	// 获取时间对象对应的 Unix 时间戳，并将其转换为字符串形式
	return strconv.FormatInt(t.Unix(), 10)
}

// MillisecondOfToStr 毫秒时间戳转字符串
// @param timeMillis 时间戳，毫秒数
// @return "yyyy-MM-dd HH:mm:ss" 格式字符串
func MillisecondOfToStr(timeMillis int64) string {
	return time.UnixMilli(timeMillis).Format(time.DateTime)
}

// SecondOfToStr 秒·时间戳转换为字符串
// @param timeSecond 时间戳，秒数
// @return "yyyy-MM-dd HH:mm:ss" 格式字符串
func SecondOfToStr(timeSecond int64) string {
	return time.Unix(timeSecond, 0).Format(time.DateTime)
}

// ======================== 日期比较 ========================

// IsSameDay 判断两个日期是否为同一天
// @param date1 日期1
// @param date2 日期2
// @return true：是同一天，false：不是同一天
func IsSameDay(date1, date2 time.Time) bool {
	date1Str := FormatDate(date1)
	date2Str := FormatDate(date2)
	if date1Str == date2Str {
		return true
	}
	return false
}

// IsSameMonth 判断两个日期是否为同一月
// @param date1 日期1
// @param date2 日期2
// @return true：是同一月，false：不是同一月
func IsSameMonth(date1, date2 time.Time) bool {
	date1Str := Format(date1, DateLayoutYM)
	date2Str := Format(date2, DateLayoutYM)
	if date1Str == date2Str {
		return true
	}
	return false
}

// ================= 日期偏移 =================

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

// OffsetMonth 偏移月
// @param date – 日期
// @param offset – 偏移月数，正数向未来偏移，负数向历史偏移
// @return 偏移后的日期
func OffsetMonth(dateTime time.Time, offset int) (time.Time, error) {
	return OffSet(dateTime, TimeMonth, offset)
}

// OffsetYear 偏移年
// @param date – 日期
// @param offset – 偏移年数，正数向未来偏移，负数向历史偏移
// @return 偏移后的日期
func OffsetYear(dateTime time.Time, offset int) (time.Time, error) {
	return OffSet(dateTime, TimeYear, offset)
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
	switch timeUnit {
	case TimeMillisecond:
		// 毫秒
		return dateTime.Add(time.Duration(offset) * time.Millisecond), nil
	case TimeSeconds:
		// 秒数
		return dateTime.Add(time.Duration(offset) * time.Second), nil
	case TimeMinute:
		// 分钟
		return dateTime.Add(time.Duration(offset) * time.Minute), nil
	case TimeHour:
		//小时
		return dateTime.Add(time.Duration(offset) * time.Hour), nil
	case TimeDay:
		// 天
		return dateTime.AddDate(0, 0, offset), nil
	case TimeWeek:
		// 周
		return dateTime.AddDate(0, 0, offset*7), nil
	case TimeMonth:
		// 月
		return dateTime.AddDate(0, offset, 0), nil
	case TimeYear:
		// 年
		return dateTime.AddDate(offset, 0, 0), nil
	default:
		return time.Time{}, errors.New("timeUnit暂不支持")
	}
}

// =============其它================

// SpendNt 计时，常用于记录某段代码的执行时间，单位：纳秒
// @param preTime 之前记录的时间
// @return 两次记录时间的差值，单位：纳秒
func SpendNt(preTime int64) int64 {
	return time.Now().UnixNano() - preTime
}

// SpendMs 计时，常用于记录某段代码的执行时间，单位：毫秒
// @param preTime 之前记录的时间
// @return 两次记录时间的差值，单位：毫秒
func SpendMs(preTime int64) int64 {
	return time.Now().UnixNano()/1e6 - preTime
}
