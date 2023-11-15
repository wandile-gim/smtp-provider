package repository

import (
	"errors"
	"github.com/google/uuid"
	"github.com/wandile/smtp-provider/internal/domain/configuration"
	"github.com/wandile/smtp-provider/internal/handler/command"
	"log"
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

func (r *InMemorySmtpRepository) FindById(id string) (*configuration.SmtpConfig, error) {
	val, ok := r.Store[id]
	if !ok {
		return nil, errors.New("config with ID:" + id + "Does not exist")
	}
	return &val, nil
}

func (r *InMemorySmtpRepository) FindByHost(host string) (*configuration.SmtpConfig, error) {

	for k, v := range r.Store {
		log.Println(r.Store[k])
		if v.Host == host {
			return &v, nil
		}
	}
	return nil, errors.New("no such host found: " + host)
}

func (r *InMemorySmtpRepository) SaveSmtpConfig(config *command.SMTPConfig) (*configuration.ConfigId, error) {
	// save to DB
	d := &configuration.SmtpConfig{
		ConfigId:  configuration.ConfigId{Id: uuid.New().String()},
		Host:      config.Host,
		Port:      config.Port,
		Username:  config.Username,
		Password:  config.Password,
		Auth:      nil,
		SSL:       config.SSL,
		TLSConfig: nil,
		LocalName: "",
	}
	r.Store[d.ConfigId.Id] = *d

	return &d.ConfigId, nil

}

// EditSmtpConfig Change username and password
func (r *InMemorySmtpRepository) EditSmtpConfig(id configuration.ConfigId, config *configuration.SmtpConfig) (*configuration.ConfigId, error) {
	confObj, ok := r.Store[id.Id]
	if !ok {
		return nil, errors.New("no such host found: " + id.Id)
	}

	// save to DB
	r.Store[id.Id] = *config

	return &confObj.ConfigId, nil
}
func (r *InMemorySmtpRepository) DeleteSmtpConfig(config *command.DeleteSMTPConfig) {
	delete(r.Store, config.Id)
}
