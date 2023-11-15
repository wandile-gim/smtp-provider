package domain

import "github.com/wandile/smtp-provider/internal/domain/configuration"

type ConfigChangeResponse struct {
	Id configuration.ConfigId
}
