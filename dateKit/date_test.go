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
	fmt.Println(ToStr(time.Now(), DefaultLayout))
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
