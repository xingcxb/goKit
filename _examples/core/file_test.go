package core

import (
	"context"
	"fmt"
	"github.com/xingcxb/goKit/core/fileKit"
	"testing"
)

func TestHomeDir(t *testing.T) {
	fmt.Println(fileKit.HomeDir(context.Background()))
}

func TestExists(t *testing.T) {
	fmt.Println(fileKit.Exists("/Users/symbol/Desktop/1.txt"))
}

func TestCreateFile(t *testing.T) {
	fmt.Println(fileKit.CreateFile("/Users/symbol/Desktop/1.txt"))
}

func TestCreateLazyFile(t *testing.T) {
	fmt.Println(fileKit.CreateLazyFile("/Users/symbol/Downloads/abc/1.txt"))
}

func TestSaveFile(t *testing.T) {
	fmt.Println(fileKit.SaveFile("/Users/symbol/Desktop", "1.txt", "hello"))
}

func TestGetFileTotalLines(t *testing.T) {
	fmt.Println(fileKit.GetFileTotalLines("/Users/symbol/Desktop/1.txt"))
}

func TestFileDirSize(t *testing.T) {
	fmt.Println(fileKit.FileDirSize("/Users/symbol/Desktop/test-1.0.jar"))
	fmt.Println(fileKit.FileDirSize("/Users/symbol/Desktop/"))
}

func TestGetCurrentAbPath(t *testing.T) {
	fmt.Println(fileKit.GetCurrentAbPath())
}

func TestUnzip(t *testing.T) {
	// 如果要解压到当前目录下可以使用 ./
	fmt.Println(fileKit.Unzip("/Users/symbol/Downloads/qts_linux_amd64.zip", "/Users/symbol/Downloads"))
}
