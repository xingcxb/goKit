package fileKit

import (
	"fmt"
	"testing"
)

func TestFileDirSize(t *testing.T) {
	fmt.Println(FileDirSize("/Users/symbol/Downloads/Gyroflow-mac-universal.dmg"))
	fmt.Println(GetFileTotalLines("/Users/symbol/Downloads/巨量.json"))
}
