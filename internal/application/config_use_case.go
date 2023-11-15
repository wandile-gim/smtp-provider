package application

import (
	"github.com/wandile/smtp-provider/internal/domain/configuration"
	"github.com/wandile/smtp-provider/internal/handler/command"
)

type (
	// SmtpConfigUseCase 서비스 유즈케이를 정의한 인터페이스

	SmtpConfigUseCase interface {
		// ChangeSmtpConfiguration 사용자는 사용할 smtp 설정을 변경할 수 있다.
		ChangeSmtpConfiguration(config *command.UpdateSMTPConfig) (*configuration.ConfigChanged, error)

		// CreateSmtpConfiguration 사용자는 smtp 설정 생성할 수 있다.
		CreateSmtpConfiguration(config *command.SMTPConfig) (*configuration.ConfigCreated, error)

		// EnableSmtpConfiguration 사용자는 smtp 설정을 선택하여 활성화 시킬 수 있다.
		EnableSmtpConfiguration(config *command.EnableSMTPConfig) (*configuration.ConfigChanged, error)
	}
)
