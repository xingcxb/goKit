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
}
