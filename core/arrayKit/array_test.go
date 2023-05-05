package arrayKit

import (
	"fmt"
	"testing"
)

func TestSliceCompare(t *testing.T) {
	a := []string{"a", "b", "c"}
	b := []string{"a", "b", "c"}
	fmt.Println(Compare(a, b))
}
