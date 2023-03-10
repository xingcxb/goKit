package ipKit

import (
	"fmt"
	"github.com/xingcxb/goKit/core/ipKit/ipService"
	"testing"
)

func TestIp(t *testing.T) {
	//fmt.Println(ipService.GetIpInfoCZ("171.42.102.199"))
	fmt.Println(ipService.GetIpInfoCip("178.42.102.199"))
}
