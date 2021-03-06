package notification

import (
	"gitlab.com/bloom42/bloom/server/config"
	"gitlab.com/bloom42/gobox/email"
)

// Init initializes the notification service
func Init(conf config.SMTPConfig) error {
	smtpConfig := email.SMTPConfig{
		Host:     conf.Host,
		Port:     conf.Port,
		Username: conf.Username,
		Password: conf.Password,
	}
	email.InitDefaultMailer(smtpConfig)
	return nil
}
