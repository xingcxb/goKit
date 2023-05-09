package reflectKit

import (
	"fmt"
	"testing"
)

func TestDefaultValue(t *testing.T) {
	type User struct {
		Name  string  `default:"xingcxb.com"`
		Power float64 `default:"100.01"`
	}
	var u User

	err := StructDefault(&u)
	if err != nil {
		fmt.Println("Uh oh: ", err)
	}

	fmt.Println(u.Name)  // Goku
	fmt.Println(u.Power) // 9000.01
}

func TestRetKit(t *testing.T) {
	ret := &Ret{}
	ret.Ok().Set("code", 200).Set("message", "success")
	fmt.Println(ret.IsOk(), ret.Get("message"))
	fmt.Println(ret.ToJSON())
}
