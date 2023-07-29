package core

import (
	"fmt"
	"github.com/xingcxb/goKit/core/arrayKit"
	"testing"
)

func TestIndexOf(t *testing.T) {
	fmt.Println(arrayKit.IndexOf([]string{"1", "2", "3"}, "2"))
}

func TestArrayContains(t *testing.T) {
	fmt.Println(arrayKit.Contains([]string{"1", "2", "3"}, "2"))
}

func TestBubbleDescSort(t *testing.T) {
	fmt.Println(arrayKit.BubbleDescSort([]string{"3", "1", "2"}))
}

func TestBubbleAscSort(t *testing.T) {
	fmt.Println(arrayKit.BubbleAscSort([]string{"3", "1", "2"}))
}

func TestJoinStringsInASCII(t *testing.T) {
	params := make(map[string]string)
	params["trade_no"] = "1111"
	params["new_ip"] = "171.42.100.153"
	params["reset"] = "1"
	fmt.Println(arrayKit.JoinStringsInASCII(params, "&", false, false))
}

func TestCompare(t *testing.T) {
	fmt.Println(arrayKit.Compare([]string{"2", "1", "3"}, []string{"1", "2", "3"}))
	fmt.Println(arrayKit.Compare([]string{"1", "2", "3"}, []string{"1", "2", "3"}))
}
