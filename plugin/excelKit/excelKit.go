package excelKit

import (
	"github.com/xingcxb/goKit/core/fileKit"
	"github.com/xingcxb/goKit/core/strKit"
	"github.com/xuri/excelize/v2"
)

// CreateExcel 生成excel
/**
 * @param filePath 文件路径
 * @param fileName 文件名
 * @param sheetName sheet名
 * @param headers 表头
 * @param contents 内容 双切片按数据库的行列来理解，第一层切片为数据库单条数据，第二层切片为单条数据中的各个元素；
 * <br/> 例如：[[1,2,3],[4,5,6]] 生成的excel为：
 * <br/> 	A B C
 * <br/> 1	1 2 3
 * <br/> 2	4 5 6
 * @return 返回生成的excel文件路径
 */
func CreateExcel(fileName, sheetName string, headers []string, contents [][]interface{}) (string, error) {
	// 强制文件路径
	filePath := fileKit.GetCurrentAbPath()
	// 创建一个excel文件
	f := excelize.NewFile()
	// 创建一个工作表
	if sheetName == "" {
		// 默认sheet名
		sheetName = "Sheet1"
	}
	index, err := f.NewSheet(sheetName)
	if err != nil {
		return "", err
	}
	if sheetName != "Sheet1" {
		// 删除默认创建的sheet1
		if err = f.DeleteSheet("Sheet1"); err != nil {
			return "", err
		}
	}
	// 设置头部信息
	for i, header := range headers {
		// 将数字转换为可识别的坐标，如：1,1转换为A1
		coordinate, err := excelize.CoordinatesToCellName(i+1, 1)
		if err != nil {
			return "", err
		}
		err = f.SetCellValue(sheetName, coordinate, header)
		if err != nil {
			return "", err
		}
	}
	// 设置内容
	for y, content := range contents {
		// 大循环用于循环每一行的数据
		for x, v := range content {
			if v == nil {
				continue
			}
			// 小循环用于循环每一列的数据
			// 将数字转换为可识别的坐标，如：1,1转换为A1
			coordinate, err := excelize.CoordinatesToCellName(x+1, y+2)
			if err != nil {
				return "", err
			}
			f.SetCellValue(sheetName, coordinate, v)
		}
	}

	// 设置工作簿的默认工作表
	f.SetActiveSheet(index)
	// 保存文件
	if err = f.SaveAs(strKit.Splicing(filePath, "/", fileName)); err != nil {
		return "", err
	}
	return strKit.Splicing(filePath, "/", fileName), nil
}
