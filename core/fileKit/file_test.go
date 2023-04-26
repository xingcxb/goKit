package fileKit

import (
	"fmt"
	"testing"
)

func TestFileDirSize(t *testing.T) {
	fmt.Println(FileDirSize("/Users/symbol/Downloads/Gyroflow-mac-universal.dmg"))
}
