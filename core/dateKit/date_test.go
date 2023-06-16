package dateKit

import (
	"fmt"
	"testing"
	"time"
)

func TestDate(t *testing.T) {
	fmt.Println(Today())
	fmt.Println(SecondOfToStr(1676257173))
	fmt.Println(MillisecondOfToStr(1676257190509))
	fmt.Println(OffSet(time.Now(), TimeWeek, -1))
	fmt.Println(Quarter(time.Now()))
	fmt.Println(Constellation(time.Now()))
	fmt.Println(OffsetMinute(time.Now(), -5))
	fmt.Println(LeapYear(2020))
	fmt.Println(Constellation(time.Now()))
	fmt.Println(Now())
	fmt.Println(Today())
	fmt.Println(GetYear(time.Now()))
	fmt.Println(GetMonth(time.Now()))
	fmt.Println(GetDay(time.Now()))
	fmt.Println(Quarter(time.Now()))
	fmt.Println(MillisecondOfToStr(time.Now().UnixMilli()))
	fmt.Println(SecondOfToStr(time.Now().Unix()))
	fmt.Println(Yesterday())
	fmt.Println(Tomorrow())
	fmt.Println(LastWeek())
	fmt.Println(NextWeek())
	fmt.Println(OffsetMillisecond(time.Now(), 1))
	fmt.Println(OffsetSecond(time.Now(), 1))
	fmt.Println(OffsetHour(time.Now(), 1))
	fmt.Println(OffsetMonth(time.Now(), 1))
	fmt.Println(OffsetYear(time.Now(), 1))
}

func TestOffset(t *testing.T) {
	fmt.Println(OffSet(ParseDate("2023-05-31"), TimeMonth, 1))
}

func TestCD(t *testing.T) {
	y, m, d := GregorianToLunar(2023, 2, 27)
	fmt.Println(y, m, d)
}

func TestLengthOfMonth(t *testing.T) {
	fmt.Println(LengthOfMonth(ParseDate("2020-02-05")))
}

// 创建公历转换到农历的映射
var lunarMap = [...]int{
	21, 33, 45, 57, 68, 79, 90, 101, 113, 124,
	135, 146, 158, 169, 181, 192, 203, 214, 226, 237,
	248, 259, 270, 282, 293, 304, 315, 326, 338, 349,
	360, 371, 383, 394, 405, 416, 428, 439, 450, 461,
	472, 484, 495, 506, 517, 528, 540, 551, 562, 573,
	585, 596, 607, 618, 629, 641, 652, 663, 674, 686,
	697, 708, 719, 731, 742, 753, 764, 775, 787, 798,
}

// 声明天数常量
const (
	mDay = 30  // 一个月有30天
	yDay = 365 // 一年有365天
)

// 计算公历日期对应的农历日期
func GregorianToLunar(year int, month int, day int) (lyear int, lmonth int, lday int) {
	// 计算该天是今年的第几天
	days := day + (mDay * (month - 1))
	// 如果是闰年，且大于 2 月
	if LeapYear(year) && month > 2 {
		days++
	}
	// 获取农历年
	lyear = year
	for ; days > yDay; lyear++ {
		if LeapYear(lyear) {
			days -= yDay + 1
		} else {
			days -= yDay
		}
	}
	// 计算农历月
	for i := 0; days > lunarMap[i]; i++ {
		lmonth = i + 1
	}
	// 计算农历日
	lday = days - lunarMap[lmonth-1] + 1
	return
}
