package dateKit

import (
	"fmt"
	"time"
)

// FormatWeChatTimeStr 格式化时间字符串为微信时间
/*
 * @param chatTimeStr 时间字符串
 * @return 仿微信时间
 */
func FormatWeChatTimeStr(chatTimeStr string) string {
	// 将时间字符串转换为时间类型
	chatTime := ParseDateTime(chatTimeStr)
	// 获取当前时间
	now := time.Now()
	// 获取间隔时间
	duration := now.Sub(chatTime)
	// 获取传入时间的年月日
	year, month, day := chatTime.Date()
	// 获取传入时间的时分
	hour, minute, _ := chatTime.Clock()
	// 获取当前时间的年月日
	nowYear, nowMonth, nowDay := now.Date()
	if duration < time.Minute {
		// 如果一分钟以内就是现在
		return "现在"
	}
	if duration < time.Hour {
		// 如果聊天时间在一小时之内，返回"x分前"
		return fmt.Sprintf("%d分钟前", int(duration.Minutes()))
	}
	if year == nowYear && month == nowMonth && day == nowDay {
		// 如果聊天时间在今天之内，则返回"HH:MM"
		return fmt.Sprintf("%02d:%02d", hour, minute)
	}
	if year == nowYear && month == nowMonth && day == nowDay-1 {
		//如果聊天时间在昨天之内，请返回"昨天HH:MM"
		return fmt.Sprintf("昨天 %02d:%02d", hour, minute)
	}
	//如果聊天时间在本周内，请返回"周xHH:MM"
	if year == nowYear && month == nowMonth && day > nowDay-7 {
		weekday := chatTime.Weekday().String()
		switch weekday {
		case "Monday":
			weekday = "一"
		case "Tuesday":
			weekday = "二"
		case "Wednesday":
			weekday = "三"
		case "Thursday":
			weekday = "四"
		case "Friday":
			weekday = "五"
		case "Saturday":
			weekday = "六"
		case "Sunday":
			weekday = "日"
		default:
			return ""
		}
		return fmt.Sprintf("周%s %02d:%02d", weekday, hour, minute)
	}
	// 其它,直接返回 "YYYY/MM/DD HH:MM"
	return fmt.Sprintf("%04d/%02d/%02d %02d:%02d", year, month, day, hour, minute)
}
