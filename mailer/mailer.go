package mailer

import (
	"crypto/tls"

	"github.com/sifatulrabbi/sifatul-api/configs"

	"gopkg.in/mail.v2"
)

type mailer struct {
	To      string `from:"to"`
	Subject string `from:"subject"`
	Message string `from:"message"`
}

func (m *mailer) Send() error {
	configs := configs.GetConfigs()
	d := mail.NewDialer("smtp.google.com", 587, configs.SMTP_USER, configs.SMTP_PASSWORD)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	return nil
}
