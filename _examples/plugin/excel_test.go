package plugin

import (
	"fmt"
	"github.com/xingcxb/goKit/plugin/office/excelKit"
	"testing"
)

func TestCreateExcel(t *testing.T) {
	var excelHeaders []string
	excelHeaders = append(excelHeaders, "姓名")
	excelHeaders = append(excelHeaders, "年龄")
	excelHeaders = append(excelHeaders, "性别")

	var excelData [][]interface{}
	for i := 0; i < 10; i++ {
		var data []interface{}
		data = append(data, "张三")
		data = append(data, 18+i)
		data = append(data, "男")
		excelData = append(excelData, data)
	}
	filePath, err := excelKit.CreateExcel("test.xlsx", "测试", excelHeaders, excelData)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(filePath)
}

func TestExcel2Vcf(t *testing.T) {
	fmt.Println(excelKit.Excel2Cvf("/Users/symbol/Downloads/无标题.xlsx", "sheet1", "/Users/symbol/Downloads/通讯录3.vcf"))
}
