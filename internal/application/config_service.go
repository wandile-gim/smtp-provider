package application

import (
	"errors"
	"github.com/wandile/smtp-provider/internal/domain"
	"github.com/wandile/smtp-provider/internal/handler/command"
	"github.com/wandile/smtp-provider/internal/repository"
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
func (s *ConfigService) ChangeSmtpConfiguration(config *command.UpdateSMTPConfig) (*domain.ConfigChangeResponse, error) {
	confObj, err := s.repository.FindById(config.Id)
	if err != nil {
		return nil, err
	}

	// domain behavior
	// change user and authenticate
	err = confObj.ChangeUser(config)
	if err != nil {
		return nil, err
	}

	// save to DB
	smtpConfig, err := s.repository.EditSmtpConfig(confObj.ConfigId, confObj)
	if err != nil {
		return nil, err
	}

	// response
	response := &domain.ConfigChangeResponse{Id: *smtpConfig}

	return response, nil
}

// CreateSmtpConfiguration 사용자는 smtp 설정 생성할 수 있다.
func (s *ConfigService) CreateSmtpConfiguration(config *command.SMTPConfig) (*domain.ConfigChangeResponse, error) {
	// find duplicate
	found, _ := s.repository.FindByHost(config.Host)
	if found != nil {
		return nil, errors.New("user Already Exist")
	}

	smtpConfig := domain.NewSmtpConfig(config.Port, config.SSL, config.Host, config.Username, config.Password)
	err := smtpConfig.AuthenticatesAccount(config)
	if err != nil {
		return nil, err
	}

	id, err := s.repository.SaveSmtpConfig(config)
	if err != nil {
		return nil, err
	}

	dto := &domain.ConfigChangeResponse{
		Id: *id,
	}

	return dto, nil
}

// EnableSmtpConfiguration 사용자는 smtp 설정을 선택하여 활성화 시킬 수 있다.
func (s *ConfigService) EnableSmtpConfiguration(config *command.EnableSMTPConfig) *domain.ConfigChangeResponse {

	return nil
}
