package util

import (
	"os"

	"gopkg.in/gomail.v2"
)

type EmailConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	From     string
}

type EmailSender struct {
	config EmailConfig
}

func NewEmailSender(config *EmailConfig) *EmailSender {
	if config == nil {
		config = &EmailConfig{
			Host:     os.Getenv("EMAIL_HOST"),
			Port:     587,
			Username: os.Getenv("MAIL_USERNAME"),
			Password: os.Getenv("MAIL_PASSWORD"),
			From:     os.Getenv("MAIL_USERNAME"),
		}
	}
	return &EmailSender{config: *config}
}

func (e *EmailSender) SendEmail(to, subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", e.config.From)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	d := gomail.NewDialer(e.config.Host, e.config.Port, e.config.Username, e.config.Password)
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
