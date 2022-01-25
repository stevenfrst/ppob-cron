package email

import (
	"gopkg.in/gomail.v2"
)

type SmtpConfig struct {
	CONFIG_SMTP_HOST       string
	CONFIG_SMTP_PORT       int
	CONFIG_SMTP_AUTH_EMAIL string
	CONFIG_AUTH_PASSWORD   string
	CONFIG_SENDER_NAME     string
}

func NewGmailConfig(g SmtpConfig) *gomail.Dialer {
	return gomail.NewDialer(
		g.CONFIG_SMTP_HOST,
		g.CONFIG_SMTP_PORT,
		g.CONFIG_SMTP_AUTH_EMAIL,
		g.CONFIG_AUTH_PASSWORD,
	)
}
