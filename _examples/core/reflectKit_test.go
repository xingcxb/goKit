package core

import (
	"fmt"
	"github.com/xingcxb/goKit/core/reflectKit"
	"testing"
)

// 读取结构体tag中为default的值
func TestStructDefault(t *testing.T) {
	type User struct {
		Name  string  `default:"xingcxb.com"`
		Power float64 `default:"100.01"`
	}
	var u User

	err := reflectKit.StructDefault(&u)
	if err != nil {
		fmt.Println("Uh oh: ", err)
	}

	fmt.Println(u.Name)  // xingcxb.com
	fmt.Println(u.Power) // 9000.01
}

func TestRet(t *testing.T) {
	ret := &reflectKit.Ret{}
	ret.Ok().Set("code", 200).Set("message", "success")
	fmt.Println(ret.IsOk(), ret.Get("message"))
	ret.Delete("message")
	fmt.Println(ret.ToJSON())
}

func TestStructToMapSS(t *testing.T) {

}
