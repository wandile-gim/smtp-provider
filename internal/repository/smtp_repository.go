package repository

import (
	"source.clobot.co.kr/spot-team/service/smtp-provider/internal/domain"
	"source.clobot.co.kr/spot-team/service/smtp-provider/internal/handler/command"
)

type SmtpConfigRepository interface {
	FindById(id string) (*domain.SmtpConfig, error)
	FindByHost(host string) (*domain.SmtpConfig, error)
	SaveSmtpConfig(config *command.SMTPConfig) (*domain.ConfigId, error)
	EditSmtpConfig(id domain.ConfigId, config *domain.SmtpConfig) (*domain.ConfigId, error)
	DeleteSmtpConfig(config *command.DeleteSMTPConfig)
}
