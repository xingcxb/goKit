package pdfKit

import (
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

// Merge 合并两个pdf文件，file1是文章的开始
/*
 * @param mergeFilePath 要合并pdf文件路径
 * @param dividerPage 分隔页
 * @param outFilePath 输出文件
 */
func Merge(mergeFilePath []string, dividerPage bool, outFilePath string) error {
	err := api.MergeCreateFile(mergeFilePath, outFilePath, dividerPage, nil)
	if err != nil {
		return err
	}
	return nil
}

// Split 拆分pdf，将pdf中的页面拆分成多个pdf
/*
 * @param file pdf文件
 */
func Split(file string) string {
	return ""
}

// Compress 压缩pdf
/*
 * @param file pdf文件
 */
func Compress(file string) string {
	return ""
}

// Pdf2Office 将pdf转换为office
/*
 * @param file pdf文件
 * @param officeType office类型 word, excel, ppt
 * @param path 保存路径
 */
func Pdf2Office(file, officeType, path string) bool {
	return false
}

// Office2Pdf 将office转换为pdf
/*
 * @param file office文件
 * @param path 保存路径
 */
func Office2Pdf(file, path string) bool {
	return false
}

// Jpg2Pdf 将jpg转换为pdf
/*
 * @param file {[]string} jpg文件
 * @param path {string} 保存路径
 */
func Jpg2Pdf(filePath []string, outPath string) error {
	err := api.ImportImagesFile(filePath, outPath, nil, nil)
	if err != nil {
		return err
	}
	return nil
}

// Pdf2Jpg 将pdf转换为jpg
/*
 * @param file pdf文件
 * @param path 保存路径
 */
func Pdf2Jpg(file, path string) bool {
	return false
}

// PdfEncrypt pdf加密
/*
 * @param file pdf文件
 * @param outFile 输出文件
 * @param userPwd 用户密码(仅拥有查看的权限)，若为空，则默认为userPwd
 * @param ownerPwd 所有者密码
 */
func PdfEncrypt(inFile, outFile, userPwd, ownerPwd string) (bool, error) {
	if ownerPwd == "" {
		// 如果所有者密码为空，则默认为用户密码
		ownerPwd = userPwd
	}
	// 设置加密配置
	conf := model.NewAESConfiguration(userPwd, ownerPwd, 256)
	conf.Permissions = model.PermissionsNone

	//加密PDF文件
	err := api.EncryptFile(inFile, outFile, conf)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Watermark pdf添加水印
/*
 * @param file pdf文件
 * @param imgPath 水印
 */
func Watermark(file, imgPath string) bool {
	return false
}
