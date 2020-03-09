package mail

import (
	"fmt"
	"gopkg.in/gomail.v2"
	"strings"
)

type Option struct {
	Host     string
	Port     int
	User     string // 发件人
	Password string // 发件人密码
	To       string // 收件人 多个用,分割
	Subject  string // 邮件主题
	Body     string // 邮件内容
}

func Send(o *Option) error {
	m := gomail.NewMessage()

	//设置发件人
	m.SetHeader("From", o.User)

	//设置发送给多个用户
	mailArrTo := strings.Split(o.To, ",")
	m.SetHeader("To", mailArrTo...)

	//设置邮件主题
	m.SetHeader("Subject", o.Subject)

	//设置邮件正文
	m.SetBody("text/html", o.Body)

	d := gomail.NewDialer(o.Host, o.Port, o.User, o.Password)

	err := d.DialAndSend(m)

	if err != nil {
		fmt.Println(err)
	}

	return err
}
