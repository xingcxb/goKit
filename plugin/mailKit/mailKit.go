package mailKit

import (
	"crypto/tls"
	"fmt"
	"github.com/xingcxb/goKit/core/strKit"
	"net"
	"net/smtp"
)

type MailAccount struct {
	FromAddress string `json:"fromAddress"` // 发送者邮箱地址
	Password    string `json:"password"`    // 发送者邮箱密码
	SmtpHost    string `json:"smtpHost"`    // smtp服务器地址
	SmtpPort    string `json:"smtpPort"`    // smtp服务器端口
}

// SendMail 发送邮件
func (m *MailAccount) SendMail(toAddress string, subject string, body string) error {
	// 设置头部信息
	headers := make(map[string]string, 0)
	headers["From"] = m.FromAddress
	headers["To"] = toAddress
	headers["Subject"] = subject

	// 设置邮件内容
	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	servername := strKit.Splicing(m.SmtpHost, ":", m.SmtpPort)
	host, _, _ := net.SplitHostPort(servername)
	// 连接邮箱服务器
	auth := smtp.PlainAuth("", m.FromAddress, m.Password, host)
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}
	conn, err := tls.Dial("tcp", servername, tlsconfig)
	if err != nil {
		return err
	}
	c, err := smtp.NewClient(conn, host)
	if err != nil {
		return err
	}
	if err = c.Auth(auth); err != nil {
		return err
	}
	if err = c.Mail(m.FromAddress); err != nil {
		return err
	}
	if err = c.Rcpt(toAddress); err != nil {
		return err
	}
	w, err := c.Data()
	if err != nil {
		return err
	}
	_, err = w.Write([]byte(message))
	if err != nil {
		return err
	}
	err = w.Close()
	if err != nil {
		return err
	}
	return c.Quit()
}
