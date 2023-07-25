package fileKit

import (
	"bufio"
	"errors"
	"github.com/xingcxb/goKit/core/strKit"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

// Exists 文件或文件夹是否存在
/**
 * @param path 文件或文件夹路径
 * @return bool true 存在 false 不存在
 */
func Exists(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

// CreateFile 创建文件
/**
 * @param filePath 文件路径
 * @return error
 */
func CreateFile(filePath string) error {
	// 判断文件是否存在
	if Exists(filePath) {
		return errors.New("文件已存在")
	}
	// 创建文件
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()
	return nil
}

// SaveFile 保存信息到文件
/**
 * @param filePath 文件夹路径
 * @param fileName 文件名
 * @param content 文件内容
 * @return error
 */
func SaveFile(filePath, fileName, content string) error {
	// 文件夹路径
	folderPath := strKit.Splicing(filePath, "/", fileName)
	// 判断文件是否存在
	if !Exists(folderPath) {
		err := CreateFile(folderPath)
		if err != nil {
			return err
		}
	}
	f, err := os.OpenFile(folderPath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	f.WriteString(content)
	defer f.Close()
	return nil
}

// GetFileTotalLines 获取文件总行数
/**
 * 注意：
 * 1. 读取大文件时，会消耗大量内存
 * 2. 读取商业加密文本时行数上面会异常
 * 3. 空白行不会被计算
 * @param filePath 文件路径
 * @return 文件总行数, 错误信息
 */
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
/**
 * @param path 文件/文件夹路径
 * @return 文件大小[byte], 错误信息
 */
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

// GetCurrentAbPath 获取当前执行文件绝对路径
func GetCurrentAbPath() string {
	dir := getCurrentAbPathByExecutable()
	tmpDir, _ := filepath.EvalSymlinks(os.TempDir())
	if strings.Contains(dir, tmpDir) {
		return getCurrentAbPathByCaller()
	}
	return dir
}

// 获取当前执行文件绝对路径
func getCurrentAbPathByExecutable() string {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	res, _ := filepath.EvalSymlinks(filepath.Dir(exePath))
	return res
}

// 获取当前执行文件绝对路径（go run）
func getCurrentAbPathByCaller() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}
