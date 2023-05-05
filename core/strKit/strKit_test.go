package strKit

import (
	"fmt"
	"testing"
	"time"
)

func TestSplicing(t *testing.T) {
	//fmt.Println(Splicing("aa", "bb", "cc"))
	beginTime := time.Now().UnixMicro()
	fmt.Println(Reverse("1234567890qwertyuiopasdfghjkzxcvbnm,"))
	endTime := time.Now().UnixMicro()
	fmt.Println("消耗时间", endTime-beginTime)
	fmt.Println(FirstUpper("abv"))
	fmt.Println(RemoveAll("abcdefghijklmn", "n", "m"))
	fmt.Println(AutoReplaceMiddle("abc"))
	fmt.Println(ReplaceIndex(1, 2, "xingcxb@hotmail.com", "3"))
	fmt.Println(SliceToStr([]string{"a", "b", "c"}))
	fmt.Println(StrToAscii("ab", "|"))
	fmt.Println(AsciiToStr("97|98", "|"))
}
