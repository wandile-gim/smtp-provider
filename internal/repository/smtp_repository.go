package repository

import (
	"github.com/wandile/smtp-provider/internal/domain/configuration"
	"github.com/wandile/smtp-provider/internal/handler/command"
	"github.com/wandile/smtp-provider/internal/repository/exception"
)

type SmtpConfigRepository interface {
	FindAll() []*configuration.SmtpConfig
	FindById(id string) (*configuration.SmtpConfig, *exception.ConfigException)
	FindByHost(host string) (*configuration.SmtpConfig, *exception.ConfigException)
	SaveSmtpConfig(config *command.SMTPConfig) (*configuration.ConfigId, error)
	UpdateSmtpConfig(id configuration.ConfigId, config *command.UpdateSMTPConfig) (*configuration.ConfigId, *exception.ConfigException)
	DeleteSmtpConfig(config *command.DeleteSMTPConfig) *exception.ConfigException
}
