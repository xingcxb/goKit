package core

import (
	"fmt"
	"github.com/xingcxb/goKit/core/regKit"
	"testing"
)

func TestContains(t *testing.T) {
	fmt.Println(regKit.Contains(regKit.FilePattern, "www.baidu.com/abc.doc"))
}

func TestFindAll(t *testing.T) {
	fmt.Println(regKit.FindAll(regKit.WordPattern, "www.baidu.com/abc.doc"))
}

func TestCount(t *testing.T) {
	fmt.Println(regKit.Count(regKit.WordPattern, "www.baidu.com/abc.doc"))
}

func TestFindAllMatchStr(t *testing.T) {
	fmt.Println(regKit.FindAllMatchStr(regKit.WordPattern, "www.baidu.com/abc.doc"))
}

func TestIsMatch(t *testing.T) {
	fmt.Println(regKit.IsMatch(regKit.NumbersPattern, "1www.baidu.com/abc.doc"))
}

func TestIndex(t *testing.T) {
	fmt.Println(regKit.Index("baidu", "www.baidu.com/abc.baidu.doc"))
}

func TestGet(t *testing.T) {
	fmt.Println(regKit.Get(regKit.WordPattern, "www.baidu.com/abc.doc", 6))
}
