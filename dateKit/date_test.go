package dateKit

import (
	"fmt"
	"testing"
	"time"
)

func TestDate(t *testing.T) {
	//fmt.Println(Today())
	//fmt.Println(SecondOfToStr(1676257173))
	//fmt.Println(MillisecondOfToStr(1676257190509))
	fmt.Println(OffSet(time.Now(), TimeWeek, -1))
}
