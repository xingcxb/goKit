package core

import (
	"fmt"
	"github.com/xingcxb/goKit/core/randomKit"
	"testing"
)

func TestRandomLong(t *testing.T) {
	fmt.Println(randomKit.RandomLong(10))
}

func TestRandomInt(t *testing.T) {
	fmt.Println(randomKit.RandomInt(0, 20))
}

func TestRandomBool(t *testing.T) {
	fmt.Println(randomKit.RandomBool())
}

func TestRandomStringWithoutStr(t *testing.T) {
	fmt.Println(randomKit.RandomStringWithoutStr(10, ""))
}

func TestRandomString(t *testing.T) {
	fmt.Println(randomKit.RandomTradeNo(10, true))
}

func TestRandomNumbers(t *testing.T) {
	fmt.Println(randomKit.RandomNumbers(10))
}

func TestRandomStr(t *testing.T) {
	fmt.Println(randomKit.RandomStr(10))
}

func TestRandomStrBasic(t *testing.T) {
	fmt.Println(randomKit.RandomStrBasic(randomKit.BaseCharNumber, 1))
}
