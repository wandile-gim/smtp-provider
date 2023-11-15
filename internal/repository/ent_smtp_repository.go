package repository

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
	"github.com/wandile/smtp-provider/ent"
	"github.com/wandile/smtp-provider/internal/domain/configuration"
	"github.com/wandile/smtp-provider/internal/handler/command"
	"github.com/wandile/smtp-provider/internal/repository/exception"
)

type EntSmtpRepository struct {
	repository *ent.Client
}

func NewEntSmtpRepository(entGo *ent.Client) SmtpConfigRepository {
	return &EntSmtpRepository{repository: entGo}
}

func (r *EntSmtpRepository) configEntityToDomainModel(conf *ent.Configuration) *configuration.SmtpConfig {
	return configuration.RestoreSmtpConfig(&configuration.ConfigId{Id: conf.ID.String()}, int(conf.Port), conf.Enable, conf.Host, conf.Username, conf.Password)
}

func (r *EntSmtpRepository) FindById(id string) (*configuration.SmtpConfig, *exception.ConfigException) {
	uid, _ := uuid.Parse(id)
	got, err := r.repository.Configuration.Get(context.Background(), uid)

	if err != nil {
		return nil, exception.DoesNotExist(err, "존재하지 않는 config 입니다.")
	}
	return r.configEntityToDomainModel(got), nil
}

func (r *EntSmtpRepository) FindAll() []*configuration.SmtpConfig {
	all, err := r.repository.Configuration.Query().All(context.Background())
	if err != nil {
		return nil
	}

	var dl []*configuration.SmtpConfig
	for _, conf := range all {
		dl = append(dl, r.configEntityToDomainModel(conf))
	}
	return dl
}

func (r *EntSmtpRepository) FindByHost(host string) (*configuration.SmtpConfig, *exception.ConfigException) {
	first, err := r.repository.Configuration.Query().Where(func(selector *sql.Selector) {
		selector.Where(sql.EQ("host", host))
	}).First(context.Background())
	if err != nil {
		log.Error(err)
		return nil, exception.DoesNotExist(err, "존재하지 않는 config 입니다.")
	}

	return r.configEntityToDomainModel(first), nil
}

func (r *EntSmtpRepository) SaveSmtpConfig(config *command.SMTPConfig) (*configuration.ConfigId, error) {
	save, err := r.repository.Configuration.Create().
		SetHost(config.Host).
		SetPort(int32(int(config.Port))).
		SetUsername(config.Username).
		SetPassword(config.Password).
		SetEnable(false).
		Save(context.Background())

	if err != nil {
		log.Error(err)
		return nil, err
	}
	id := &configuration.ConfigId{Id: save.ID.String()}
	return id, nil
}

func (r *EntSmtpRepository) UpdateSmtpConfig(id configuration.ConfigId, config *command.UpdateSMTPConfig) (*configuration.ConfigId, *exception.ConfigException) {
	found, configException := r.FindById(id.Id)
	if configException != nil {
		log.Error(configException.ToMessage())
		return nil, configException
	}
	foundId, _ := uuid.Parse(found.ConfigId.Id)
	updateID, err := r.repository.Configuration.UpdateOneID(foundId).
		SetUsername(config.Username).
		SetPassword(config.Password).
		SetEnable(config.Enable).
		Save(context.Background())

	if err != nil {
		log.Error(err)
		return nil, exception.SaveError(err, err.Error())
	}

	return &configuration.ConfigId{Id: updateID.ID.String()}, nil
}

func (r *EntSmtpRepository) DeleteSmtpConfig(config *command.DeleteSMTPConfig) *exception.ConfigException {
	id, err := r.FindById(config.Id)
	if err != nil {
		return exception.DoesNotExist(err.Error, "존재하지 않는 config 입니다.")
	}

	uid, _ := uuid.Parse(id.ConfigId.Id)
	r.repository.Configuration.DeleteOneID(uid)
	return nil
}
