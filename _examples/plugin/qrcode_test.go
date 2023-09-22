package plugin

import (
	"fmt"
	"github.com/xingcxb/goKit/plugin/qrcodeKit"
	"testing"
)

func TestCreateQrcode(t *testing.T) {
	qrParmas := qrcodeKit.QrCodeParams{
		Content:   "https://xingcxb.com",
		QrColor:   "#000000",
		QrColorBg: "#ffffff",
		//QRWidth:   uint8(30),
		QrFillPath: "./22.png",
		//LogoPath: "./logo.png",
		SavePath: "./2.png",
	}
	err := qrParmas.CreateQrCodeFile()
	if err != nil {
		fmt.Println(err)
	}
	qrParams2 := qrcodeKit.QrCodeParams{
		Content:   "https://xingcxb.com",
		QrColor:   "#000000",
		QrColorBg: "#ffffff",
		//QRWidth:   uint8(30),
		QrFillPath: "./22.png",
		//LogoPath: "./logo.png",
	}
	fmt.Println(qrParams2.CreateQrCodeBase64Str())
}
