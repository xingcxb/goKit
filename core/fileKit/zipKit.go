package fileKit

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

func Unzip(zipFile, destDir string) error {
	// 打开 zip 文件
	r, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}
	defer r.Close()
	// 解压每个文件
	for _, file := range r.File {
		// 获取解压后的文件路径
		fpath := filepath.Join(destDir, file.Name)

		// 如果是目录则创建
		if file.FileInfo().IsDir() {
			err := os.MkdirAll(fpath, os.ModePerm)
			if err != nil {
				return err
			}
			continue
		}

		// 创建包含文件的目录（如果不存在）
		if err := os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return err
		}

		// 打开 zip 文件中的内容
		inFile, err := file.Open()
		if err != nil {
			return err
		}
		defer inFile.Close()

		// 创建要写入的文件
		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}
		defer outFile.Close()

		// 将内容写入目标文件
		_, err = io.Copy(outFile, inFile)
		if err != nil {
			return err
		}
	}
	return nil
}
