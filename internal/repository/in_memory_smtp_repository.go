package repository

import (
	"errors"
	"github.com/google/uuid"
	"log"
	"source.clobot.co.kr/spot-team/service/smtp-provider/internal/domain"
	"source.clobot.co.kr/spot-team/service/smtp-provider/internal/handler/command"
)

func NewStore() *map[string]domain.SmtpConfig {
	d := make(map[string]domain.SmtpConfig, 0)
	return &d
}

type InMemorySmtpRepository struct {
	Store map[string]domain.SmtpConfig
}

func NewInMemorySmtpRepository(data map[string]domain.SmtpConfig) SmtpConfigRepository {
	return &InMemorySmtpRepository{Store: data}
}

func (r *InMemorySmtpRepository) FindById(id string) (*domain.SmtpConfig, error) {
	val, ok := r.Store[id]
	if !ok {
		return nil, errors.New("config with ID:" + id + "Does not exist")
	}
	return &val, nil
}

func (r *InMemorySmtpRepository) FindByHost(host string) (*domain.SmtpConfig, error) {

	for k, v := range r.Store {
		log.Println(r.Store[k])
		if v.Host == host {
			return &v, nil
		}
	}
	return nil, errors.New("no such host found: " + host)
}

func (r *InMemorySmtpRepository) SaveSmtpConfig(config *command.SMTPConfig) (*domain.ConfigId, error) {
	// save to DB
	d := &domain.SmtpConfig{
		ConfigId:  domain.ConfigId{Id: uuid.New().String()},
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
func (r *InMemorySmtpRepository) EditSmtpConfig(id domain.ConfigId, config *domain.SmtpConfig) (*domain.ConfigId, error) {
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
