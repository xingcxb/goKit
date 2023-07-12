package core

import (
	"fmt"
	"github.com/xingcxb/goKit/core/dateKit"
	"testing"
	"time"
)

// 获取当前时间
func TestNow(t *testing.T) {
	fmt.Println(dateKit.Now())
}

// 获取当前日期
func TestToday(t *testing.T) {
	fmt.Println(dateKit.Today())
}

// 获取某天的开始时间
func TestBeginOfDay(t *testing.T) {
	fmt.Println(dateKit.BeginOfDay(time.Now()))
}

// 获取某天的结束时间
func TestEndOfDay(t *testing.T) {
	fmt.Println(dateKit.EndOfDay(time.Now()))
}

// 获取指定时间的年份
func TestGetYear(t *testing.T) {
	fmt.Println(dateKit.GetYear(time.Now()))
}

// 获取指定时间的月份
func TestGetMonth(t *testing.T) {
	fmt.Println(dateKit.GetMonth(time.Now()))
}

// 获取指定时间的号数
func TestGetDay(t *testing.T) {
	fmt.Println(dateKit.GetDay(time.Now()))
}

// 获取指定时间的季度
func TestQuarter(t *testing.T) {
	fmt.Println(dateKit.Quarter(time.Now()))
}

// 获取昨天
func TestYesterday(t *testing.T) {
	fmt.Println(dateKit.Yesterday())
}

// 获取明天
func TestTomorrow(t *testing.T) {
	fmt.Println(dateKit.Tomorrow())
}

// 获取上周
func TestLastWeek(t *testing.T) {
	fmt.Println(dateKit.LastWeek())
}

// 获取下周
func TestNextWeek(t *testing.T) {
	fmt.Println(dateKit.NextWeek())
}

// 获取指定日期的月份天数
func TestLengthOfMonth(t *testing.T) {
	fmt.Println(dateKit.LengthOfMonth(dateKit.ParseDate("2020-02-05")))
}

// 获取指定日期的年份天数
func TestLengthOfYear(t *testing.T) {
	fmt.Println(dateKit.LengthOfYear(dateKit.ParseDate("2020-02-05")))
}

// 判断是否是闰年
func TestLeapYear(t *testing.T) {
	fmt.Println(dateKit.LeapYear(2022))
}

// 时间转换为日期和时间
func TestDateTimeToStr(t *testing.T) {
	fmt.Println(dateKit.FormatDateTime(time.Now()))
}

// 时间转换为日期
func TestDateToStr(t *testing.T) {
	fmt.Println(dateKit.FormatDate(time.Now()))
}

// 时间转换为指定格式
func TestFormat(t *testing.T) {
	fmt.Println(dateKit.Format(time.Now(), dateKit.DateLayoutYM))
}

// 字符串时间格式化为标准日期类型
func TestParseDate(t *testing.T) {
	fmt.Println(dateKit.ParseDate("2020-02-05"))
}

// 字符串时间格式化为标准日期时间类型
func TestParseDateTime(t *testing.T) {
	fmt.Println(dateKit.ParseDateTime("2020-02-05 12:12:12"))
}

// 日期时间字符串转为毫秒字符串
func TestDateTimeToMillisecondStr(t *testing.T) {
	fmt.Println(dateKit.DateTimeStrToMillStr("2020-02-05 12:12:12"))
}

// 日期时间字符串转为秒字符串
func TestDateTimeToSecondStr(t *testing.T) {
	fmt.Println(dateKit.DateTimeStrToSecondStr("2020-02-05 12:12:12"))
}

// 毫秒时间戳转换为日期时间
func TestMillisecondToDateTime(t *testing.T) {
	fmt.Println(dateKit.MillisecondOfToStr(1580892732000))
}

// 秒时间戳转换为日期时间
func TestSecondToDateTime(t *testing.T) {
	fmt.Println(dateKit.SecondOfToStr(1580892732))
}

// 判断两个日期是否为同一天
func TestIsSameDay(t *testing.T) {
	fmt.Println(dateKit.IsSameDay(time.Now(), time.Now()))
}

// 判断两个日期是否为同一月
func TestIsSameMonth(t *testing.T) {
	fmt.Println(dateKit.IsSameMonth(time.Now(), time.Now()))
}

// 日期偏移毫秒数
func TestAddMillisecond(t *testing.T) {
	fmt.Println(dateKit.OffsetMillisecond(time.Now(), 1000))
}

// 日期偏移秒数
func TestAddSecond(t *testing.T) {
	fmt.Println(dateKit.OffsetSecond(time.Now(), 1000))
}

// 日期偏移分钟数
func TestAddMinute(t *testing.T) {
	fmt.Println(dateKit.OffsetMinute(time.Now(), 1000))
}

// 日期偏移小时数
func TestAddHour(t *testing.T) {
	fmt.Println(dateKit.OffsetHour(time.Now(), 1000))
}

// 日期偏移天数
func TestAddDay(t *testing.T) {
	fmt.Println(dateKit.OffsetDay(time.Now(), 1000))
}

// 日期偏移周数
func TestAddWeek(t *testing.T) {
	fmt.Println(dateKit.OffsetWeek(time.Now(), 1))
}

// 日期偏移月数
func TestAddMonth(t *testing.T) {
	fmt.Println(dateKit.OffsetMonth(time.Now(), 1))
}

// 日期偏移年数
func TestAddYear(t *testing.T) {
	fmt.Println(dateKit.OffsetYear(time.Now(), 1))
}

// 日期偏移
func TestAdd(t *testing.T) {
	fmt.Println(dateKit.OffSet(time.Now(), dateKit.TimeYear, 1))
}

// 计时纳秒
func TestNt(t *testing.T) {
	fmt.Println(dateKit.SpendNt(1580892732 * 1e6))
}

// 计算毫秒
func TestMs(t *testing.T) {
	fmt.Println(dateKit.SpendMs(1580892732 * 1e3))
}
