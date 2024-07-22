package plugin

import (
	"fmt"
	"github.com/xingcxb/goKit/plugin/pdfKit"
	"testing"
)

// TestPdfEncrypt pdf 加密
func TestPdfEncrypt(t *testing.T) {
	inFile := "/Users/symbol/Downloads/test.pdf"
	outFile := "/Users/symbol/Downloads/testEn.pdf"
	_, err := pdfKit.PdfEncrypt(inFile, outFile, "123456", "778899")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("加密成功")
}
