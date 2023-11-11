package application

import (
	"source.clobot.co.kr/spot-team/service/smtp-provider/internal/domain"
	"source.clobot.co.kr/spot-team/service/smtp-provider/internal/handler/command"
)

type (
	// SmtpConfigUseCase 서비스 유즈케이를 정의한 인터페이스

	SmtpConfigUseCase interface {
		// ChangeSmtpConfiguration 사용자는 사용할 smtp 설정을 변경할 수 있다.
		ChangeSmtpConfiguration(config *command.UpdateSMTPConfig) (*domain.ConfigChangeResponse, error)

		// CreateSmtpConfiguration 사용자는 smtp 설정 생성할 수 있다.
		CreateSmtpConfiguration(config *command.SMTPConfig) (*domain.ConfigChangeResponse, error)

		// EnableSmtpConfiguration 사용자는 smtp 설정을 선택하여 활성화 시킬 수 있다.
		EnableSmtpConfiguration(config *command.EnableSMTPConfig) *domain.ConfigChangeResponse
	}
)