package core

import (
	"fmt"
	"github.com/xingcxb/goKit/core/strKit"
	"testing"
)

func TestLength(t *testing.T) {
	fmt.Println(strKit.Length("𠮷"))
	fmt.Println(strKit.Length("👩‍👩‍👦‍👦"))
}

func TestIsNum(t *testing.T) {
	fmt.Println(strKit.IsNum("123"))
	fmt.Println(strKit.IsNum("123.123"))
	fmt.Println(strKit.IsNum("123.123.123"))
}

func TestSplicing(t *testing.T) {
	fmt.Println(strKit.Splicing("aa", "bb", "cc"))
}

func TestSubString(t *testing.T) {
	fmt.Println(strKit.SubString("abcdefg", 1, 3))
}

func TestRemoveAll(t *testing.T) {
	fmt.Println(strKit.RemoveAll("abcdefghijklmn", "n", "m"))
}

func TestMapParamsToUrlParams(t *testing.T) {
	fmt.Println(strKit.MapParamsToUrlParams(map[string]string{"a": "1", "b": "2"}))
}

func TestStrParamsToMapParams(t *testing.T) {
	fmt.Println(strKit.StrParamsToMapParams("a=1&b=2"))
}

func TestReplaceIndex(t *testing.T) {
	fmt.Println(strKit.ReplaceIndex(0, 2, "ababab", "12"))
}

func TestAutoReplaceMiddle(t *testing.T) {
	fmt.Println(strKit.AutoReplaceMiddle("abc"))
}

func TestIsEnLetter(t *testing.T) {
	fmt.Println(strKit.IsEnLetter("你"))
	fmt.Println(strKit.IsEnLetter("a"))
}

func TestReverse(t *testing.T) {
	fmt.Println(strKit.Reverse("1234567890qwertyuiopasdfghjkzxcvbnm,"))
}

func TestFirstUpper(t *testing.T) {
	fmt.Println(strKit.FirstUpper("abv"))
}

func TestFirstLower(t *testing.T) {
	fmt.Println(strKit.FirstLower("ABV"))
}

func TestSplitterToHump(t *testing.T) {
	fmt.Println(strKit.SplitterToHump("abc_hello", "_"))
	fmt.Println(strKit.SplitterToHump("你_好_nice", ""))
}

func TestSliceToStr(t *testing.T) {
	fmt.Println(strKit.SliceToStr([]string{"a", "b", "c"}))
}

func TestCleanStrSymbol(t *testing.T) {
	fmt.Println(strKit.CleanStrSymbol("a<div>b</div>c"))
}

func TestStrToAscii(t *testing.T) {
	fmt.Println(strKit.StrToAscii("a", ""))
}

func TestAsciiToStr(t *testing.T) {
	fmt.Println(strKit.AsciiToStr("97|98", "|"))
}

func TestCheckIdCardValid(t *testing.T) {
	fmt.Println(strKit.CheckIdCardValid("500154199301135886", true))
}
