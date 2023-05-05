package fileKit

import (
	"bufio"
	"errors"
	"os"
)

// GetFileTotalLines 获取文件总行数<br/>
// 注意：<br/>
//  1. 读取大文件时，会消耗大量内存
//  2. 读取商业加密文本时行数上面会异常
//  3. 空白行不会被计算
//
// @param filePath 文件路径
func GetFileTotalLines(filePath string) (int, error) {
	if filePath == "" {
		return 0, errors.New("文件路径不能为空")
	}
	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var lines int
	for scanner.Scan() {
		lines++
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	return lines, nil
}

// FileDirSize 获取文件/文件夹下所有文件的大小
// @param path 文件/文件夹路径
// @return 文件大小[byte], 错误信息
func FileDirSize(path string) (int, error) {
	// 打开文件
	f, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	// 获取文件信息
	fi, err := f.Stat()
	if err != nil {
		return 0, err
	}
	// 判断是否为文件夹
	if fi.IsDir() {
		// 获取文件夹下所有文件
		fis, err := f.Readdir(-1)
		if err != nil {
			return 0, err
		}
		// 定义文件大小
		var size int
		// 遍历文件夹下所有文件
		for _, fi := range fis {
			// 判断是否为文件夹
			if fi.IsDir() {
				// 递归调用
				s, err := FileDirSize(path + "/" + fi.Name())
				if err != nil {
					return 0, err
				}
				// 累加文件大小
				size += s
			} else {
				// 累加文件大小
				size += int(fi.Size())
			}
		}
		return size, nil
	} else {
		// 返回文件大小
		return int(fi.Size()), nil
	}
}
