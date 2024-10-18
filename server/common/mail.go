package common

import (
	"crypto/tls"
	"fmt"
	"gincrm/global"
	"gincrm/models"
	"gopkg.in/gomail.v2"
	"log"
)

func SendEmail(email, content string) error {
	smtp := global.Config.Mail.Smtp
	secret := global.Config.Mail.Secret
	sender := global.Config.Mail.Sender

	message := gomail.NewMessage()
	message.SetHeader("From", sender)
	message.SetHeader("To", email)
	message.SetHeader("Cc", email)
	message.SetHeader("Bcc", email)
	message.SetHeader("Subject", "TestSubject")
	message.SetBody("text/html", content)
	dialer := gomail.NewDialer(smtp, 465, sender, secret)
	dialer.TLSConfig = &tls.Config{
		InsecureSkipVerify: true,
	}

	if err := dialer.DialAndSend(message); err != nil {
		fmt.Printf("mail send error: %s", err)
		return err
	}
	return nil
}

func SendMailToCustomer(mp models.MailParam) error {
	m := gomail.NewMessage()
	m.SetHeader("From", mp.Sender)
	m.SetHeader("To", mp.Receiver)
	m.SetHeader("Subject", mp.Subject)
	m.SetBody("text/html", mp.Content)
	if mp.Attachment != "" {
		m.Attach(mp.Attachment)
	}
	d := gomail.NewDialer(mp.Smtp, mp.Port, mp.Sender, mp.AuthCode)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := d.DialAndSend(m); err != nil {
		log.Printf("send mail to customer error : %s", err)
		return err
	}
	return nil
}

func DialMail(smtp string, port int, sender, authCode string) error {
	d := gomail.NewDialer(smtp, port, sender, authCode)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	_, err := d.Dial()
	return err
}
