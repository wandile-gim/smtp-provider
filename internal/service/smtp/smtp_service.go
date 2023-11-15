package smtp

import (
	"github.com/wandile/smtp-provider/internal/domain/configuration"
	"github.com/wandile/smtp-provider/internal/handler/command"
	"github.com/wandile/smtp-provider/internal/repository"
)

type SmtpService struct {
	repository repository.SmtpConfigRepository
}

func NewSmtpService(repo repository.SmtpConfigRepository) *SmtpService {
	return &SmtpService{repository: repo}
}

func (s *SmtpService) GetOnlyEnabled(id *configuration.ConfigId) configuration.ConfigId {
	all := s.repository.FindAll()

	for _, v := range all {
		if v.ConfigId != *id {
			v.EnableConfig(false)
		} else {
			v.EnableConfig(true)
			id = &v.ConfigId
		}
		c := &command.UpdateSMTPConfig{
			Id:       v.ConfigId.Id,
			Username: v.Username,
			Password: v.Password,
			Enable:   v.Enable,
		}

		s.repository.UpdateSmtpConfig(v.ConfigId, c)
	}
	return *id
}
