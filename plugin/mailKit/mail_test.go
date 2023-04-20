package mailKit

import (
	"fmt"
	"testing"
)

func TestSend(t *testing.T) {
	account := &MailAccount{
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
