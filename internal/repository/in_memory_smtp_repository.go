package repository

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
	"github.com/wandile/smtp-provider/internal/domain/configuration"
	"github.com/wandile/smtp-provider/internal/handler/command"
	"github.com/wandile/smtp-provider/internal/repository/exception"
)

func NewStore() *map[string]configuration.SmtpConfig {
	d := make(map[string]configuration.SmtpConfig, 0)
	return &d
}

type InMemorySmtpRepository struct {
	Store map[string]configuration.SmtpConfig
}

func NewInMemorySmtpRepository(data map[string]configuration.SmtpConfig) SmtpConfigRepository {
	return &InMemorySmtpRepository{Store: data}
}

func (r *InMemorySmtpRepository) FindAll() []*configuration.SmtpConfig {
	var all []*configuration.SmtpConfig
	for _, v := range r.Store {
		all = append(all, &v)
	}

	return all
}

func (r *InMemorySmtpRepository) FindById(id string) (*configuration.SmtpConfig, *exception.ConfigException) {
	val, ok := r.Store[id]
	if !ok {
		return nil, exception.DoesNotExist(nil, "존재하지 않는 config 입니다.")
	}
	return &val, nil
}

func (r *InMemorySmtpRepository) FindByHost(host string) (*configuration.SmtpConfig, *exception.ConfigException) {

	for _, v := range r.Store {
		if v.Host == host {
			return &v, nil
		}
	}
	return nil, exception.DoesNotExist(nil, "존재하지 않는 config 입니다.")
}

func (r *InMemorySmtpRepository) SaveSmtpConfig(conf *command.SMTPConfig) (*configuration.ConfigId, error) {
	// save to DB
	d := &configuration.SmtpConfig{
		ConfigId:  configuration.ConfigId{Id: uuid.New().String()},
		Host:      conf.Host,
		Port:      conf.Port,
		Username:  conf.Username,
		Password:  conf.Password,
		Auth:      nil,
		SSL:       conf.SSL,
		TLSConfig: nil,
		LocalName: "",
	}
	r.Store[d.ConfigId.Id] = *d
	log.Info("saved: ", r.Store[d.ConfigId.Id])
	return &d.ConfigId, nil

}

// EditSmtpConfig Change username and password
func (r *InMemorySmtpRepository) UpdateSmtpConfig(id configuration.ConfigId, config *command.UpdateSMTPConfig) (*configuration.ConfigId, *exception.ConfigException) {
	confObj, ok := r.Store[id.Id]
	if !ok {
		return nil, exception.DoesNotExist(nil, "존재하지 않는 config 입니다.")
	}

	// save to DB
	u := &configuration.SmtpConfig{
		ConfigId: confObj.ConfigId,
		Host:     confObj.Host,
		Port:     confObj.Port,
		Username: config.Username,
		Password: config.Password,
		Enable:   config.Enable,
	}
	log.Info(u.Enable)
	r.Store[id.Id] = *u

	return &confObj.ConfigId, nil
}
func (r *InMemorySmtpRepository) DeleteSmtpConfig(config *command.DeleteSMTPConfig) *exception.ConfigException {
	delete(r.Store, config.Id)
	return nil
}
