package plugin

import (
	"fmt"
	"github.com/xingcxb/goKit/plugin/pdfKit"
	"testing"
)

// TestMerge 合并pdf
func TestMerge(t *testing.T) {
	filePath := []string{"/Users/symbol/Downloads/test.pdf", "/Users/symbol/Downloads/test2.pdf"}
	err := pdfKit.Merge(filePath, false, "/Users/symbol/Downloads/test3.pdf")
	if err != nil {
		fmt.Println(err)
		return
	}
}

// TestJpg2Pdf 图片转pdf
func TestJpg2Pdf(t *testing.T) {
	imagePath := []string{"/Users/symbol/Downloads/Snipaste_2024-12-06_11-10-58.png", "/Users/symbol/Downloads/Snipaste_2024-11-23_14-20-03.png"}
	err := pdfKit.Jpg2Pdf(imagePath, "/Users/symbol/Downloads/test2.pdf")
	if err != nil {
		fmt.Println(err)
		return
	}
}

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
