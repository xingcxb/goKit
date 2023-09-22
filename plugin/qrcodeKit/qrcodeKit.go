// Package qrcodeKit 二维码生成器
package qrcodeKit

import (
	"bytes"
	"github.com/xingcxb/goKit/core/cryptoKit"
	"github.com/yeqown/go-qrcode/v2"
	"github.com/yeqown/go-qrcode/writer/standard"
	"io"
)

type nopCloser struct {
	io.Writer
}

func (nopCloser) Close() error { return nil }

type QrCodeParams struct {
	Content    string // * 内容
	QrColor    string // 二维码颜色,默认黑色
	QrColorBg  string // 二维码背景颜色,默认白色
	QRWidth    uint8  // 指定每个qr块的宽度
	QrFillPath string // 二维码填充的图片
	LogoPath   string // 中间logo的位置 图片大小166*166
	SavePath   string // 保存的位置
}

// CreateQrCodeFile 创建二维码
/*
 * @param content 内容
 * @param qrFillPath 二维码填充的图片
 * @param logoPath 中间logo的位置 图片大小166*166
 * @param savePath 保存的位置
 */
func (q QrCodeParams) CreateQrCodeFile() error {
	// 创建参数
	qrc, err := qrcode.New(q.Content)
	if err != nil {
		return err
	}
	imageOptions := make([]standard.ImageOption, 0)
	if q.QRWidth == 0 {
		q.QRWidth = 30
	}
	imageOptions = append(imageOptions, standard.WithQRWidth(q.QRWidth)) // 指定每个qr块的宽度
	if q.QrColor == "" {
		q.QrColor = "#000000"
	}
	imageOptions = append(imageOptions, standard.WithFgColorRGBHex(q.QrColor)) // 二维码格子颜色
	if q.QrColorBg == "" {
		q.QrColorBg = "#ffffff"
	}
	imageOptions = append(imageOptions, standard.WithBgColorRGBHex(q.QrColorBg)) // 二维码背景颜色

	if q.QrFillPath != "" {
		imageOptions = append(imageOptions, standard.WithHalftone(q.QrFillPath)) // 图片融入到二维码中
	}
	if q.LogoPath != "" {
		imageOptions = append(imageOptions, standard.WithLogoImageFilePNG(q.LogoPath)) // 中间logo
	}
	// 配置qr参数
	w, err := standard.New(
		q.SavePath, // 保存地址
		imageOptions...,
	)
	if err = qrc.Save(w); err != nil {
		return err
	}
	return nil
}

// CreateQrCodeBase64Str 创建二维码返回base64字符串
/*
 * @param content 内容
 * @param qrFillPath 二维码填充的图片
 * @param logoPath 中间logo的位置
 */
func (q QrCodeParams) CreateQrCodeBase64Str() (string, error) {
	// 创建参数
	qrc, err := qrcode.NewWith(q.Content)
	if err != nil {
		return "", err
	}
	imageOptions := make([]standard.ImageOption, 0)
	if q.QRWidth == 0 {
		q.QRWidth = 30
	}
	imageOptions = append(imageOptions, standard.WithQRWidth(q.QRWidth)) // 指定每个qr块的宽度
	if q.QrColor == "" {
		q.QrColor = "#000000"
	}
	imageOptions = append(imageOptions, standard.WithFgColorRGBHex(q.QrColor)) // 二维码格子颜色
	if q.QrColorBg == "" {
		q.QrColorBg = "#ffffff"
	}
	imageOptions = append(imageOptions, standard.WithBgColorRGBHex(q.QrColorBg)) // 二维码背景颜色

	if q.QrFillPath != "" {
		imageOptions = append(imageOptions, standard.WithHalftone(q.QrFillPath)) // 图片融入到二维码中
	}
	if q.LogoPath != "" {
		imageOptions = append(imageOptions, standard.WithLogoImageFilePNG(q.LogoPath)) // 中间logo
	}
	buf := bytes.NewBuffer(nil)
	wr := nopCloser{Writer: buf}
	w2 := standard.NewWithWriter(
		wr,
		imageOptions...,
	)
	if err = qrc.Save(w2); err != nil {
		return "", err
	}
	return cryptoKit.Base64Encode(buf.String()), nil
}
