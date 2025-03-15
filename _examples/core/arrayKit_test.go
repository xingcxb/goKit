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

func TestBinarySearchIndexOf(t *testing.T) {
	fmt.Println(arrayKit.BinarySearchIndexOf([]string{"1", "2", "3"}, "2"))
}

func TestBubbleDescSort(t *testing.T) {
	fmt.Println(arrayKit.BubbleDescSort([]string{"3", "1", "2"}))
}

func TestBubbleAscSort(t *testing.T) {
	fmt.Println(arrayKit.BubbleAscSort([]string{"10", "2", "33", "4", "5"}))
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

func TestSliceDiff(t *testing.T) {
	a := []string{
		"1.1.1.1",
		"2.2.2.2"}
	b := []string{
		"1.1.1.1",
		"2.2.2.2",
		"203.209.242.97"}
	add, del := arrayKit.SliceDiff(a, b)
	fmt.Println(add, "\n", del)
}
