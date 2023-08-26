package plugin

import (
	"fmt"
	"github.com/xingcxb/goKit/plugin/translationKit"
	"testing"
)

func TestDeepL(t *testing.T) {
	l := &translationKit.DeepL{}
	fmt.Println(l.Translation("hello goKit", "en", "zh"))
}
