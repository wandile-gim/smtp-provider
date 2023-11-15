package application

import (
	"errors"
	"github.com/wandile/smtp-provider/internal/domain/configuration"
	"github.com/wandile/smtp-provider/internal/handler/command"
	"github.com/wandile/smtp-provider/internal/repository"
	"github.com/wandile/smtp-provider/internal/service/smtp"
)

type ConfigService struct {
	repository repository.SmtpConfigRepository
}

func NewConfigService(repo repository.SmtpConfigRepository) SmtpConfigUseCase {
	return &ConfigService{
		repository: repo,
	}
}

// ChangeSmtpConfiguration 사용자는 사용할 smtp 설정을 변경할 수 있다.
func (s *ConfigService) ChangeSmtpConfiguration(config *command.UpdateSMTPConfig) (*configuration.ConfigChanged, error) {
	confObj, err := s.repository.FindById(config.Id)
	if err != nil {
		return nil, err.Error
	}

	// domain behavior
	// change user and authenticate
	err2 := confObj.ChangeUser(config)
	if err2 != nil {
		return nil, err2
	}

	// save to DB
	smtpConfig, err3 := s.repository.UpdateSmtpConfig(confObj.ConfigId, config)
	if err3 != nil {
		return nil, err3.Error
	}

	// response
	response := &configuration.ConfigChanged{Id: *smtpConfig}

	return response, nil
}

// CreateSmtpConfiguration 사용자는 smtp 설정 생성할 수 있다.
func (s *ConfigService) CreateSmtpConfiguration(config *command.SMTPConfig) (*configuration.ConfigCreated, error) {
	// find duplicate
	found, _ := s.repository.FindByHost(config.Host)
	if found != nil {
		return nil, errors.New("user Already Exist")
	}

	smtpConfig := configuration.NewSmtpConfig(config.Port, config.SSL, config.Host, config.Username, config.Password)
	err := smtpConfig.AuthenticatesAccount(config)
	if err != nil {
		return nil, err
	}

	id, err := s.repository.SaveSmtpConfig(config)
	if err != nil {
		return nil, err
	}

	dto := &configuration.ConfigCreated{
		Id:   *id,
		Host: config.Host,
	}

	return dto, nil
}

// EnableSmtpConfiguration 사용자는 smtp 설정을 선택하여 활성화 시킬 수 있다.
func (s *ConfigService) EnableSmtpConfiguration(config *command.EnableSMTPConfig) (*configuration.ConfigChanged, error) {
	ds := smtp.NewSmtpService(s.repository)
	enabled := ds.GetOnlyEnabled(&configuration.ConfigId{Id: config.Id})
	response := &configuration.ConfigChanged{
		Id: enabled,
	}

	return response, nil
}
