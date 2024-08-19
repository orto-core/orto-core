package service

import (
	"bytes"
	"fmt"
	"os"
	"text/template"

	"gopkg.in/gomail.v2"
)

type MailService interface {
	SendMail(string, string, string) error
}

type mailService struct {
	Dialer       *gomail.Dialer
	TemplatePath string
}

func NewMailService(port int, templatePath, host, username, password string) MailService {
	return &mailService{
		TemplatePath: templatePath,
		Dialer:       gomail.NewDialer(host, port, username, password),
	}
}

func (m *mailService) SendMail(to, subject, otp string) error {
	body, err := m.parseTemplate(m.TemplatePath, otp)
	if err != nil {
		return err
	}

	from := os.Getenv("EMAIL_USERNAME")
	if from == "" {
		return fmt.Errorf("EMAIL_USERNAME environment variable not set")
	}

	mail := gomail.NewMessage()
	mail.SetHeader("From", os.Getenv("EMAIL_USERNAME"))
	mail.SetHeader("To", to)
	mail.SetHeader("Subject", subject)
	mail.SetBody("text/html", body)

	if err := m.Dialer.DialAndSend(mail); err != nil {
		return err
	}

	return nil
}

func (m *mailService) parseTemplate(templatePath string, otp string) (string, error) {
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		return "", err
	}
	buff := new(bytes.Buffer)
	if err = t.Execute(buff, otp); err != nil {
		return "", err
	}
	return buff.String(), nil
}
