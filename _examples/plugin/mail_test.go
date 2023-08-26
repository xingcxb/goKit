package plugin

import (
	"fmt"
	"github.com/xingcxb/goKit/plugin/mailKit"
	"testing"
)

func TestSend(t *testing.T) {
	account := &mailKit.MailAccount{
		FromAddress: "123456@juliangip.com",
		Password:    "123456",
		SmtpHost:    "smtpdm.aliyun.com",
		SmtpPort:    "465",
	}
	err := account.SendMail("123456@qq.com", "测试邮件", "测试邮件内容")
	if err != nil {
		fmt.Println(err)
	}
}
