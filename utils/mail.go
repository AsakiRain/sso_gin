package utils

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"sso_gin/config"

	"gopkg.in/gomail.v2"
)

func ParseTemplate(tpl string, data interface{}) (string, error) {
	filePath := fmt.Sprintf("%s/%s", config.TemplatePath, tpl)
	t, err := template.ParseFiles(filePath)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	err = t.Execute(buf, data)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

func SendMail(to, subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", fmt.Sprintf("%s<%s>", config.MailAlias, config.MailUser))
	m.SetHeader("To", to)
	// m.SetAddressHeader("Cc", config.MailFrom, "System")		//抄送，可多个
	// m.SetAddressHeader("Bcc", config.MailFrom, "System")		//暗送，可多个
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	d := gomail.NewDialer(config.MailHost, int(config.MailPort), config.MailUser, config.MailPass)
	err := d.DialAndSend(m)
	if err != nil {
		log.Printf("发送邮件失败：%v", err)
	}

	return err
}
