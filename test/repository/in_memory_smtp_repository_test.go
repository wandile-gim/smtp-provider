package repository

import (
	"github.com/stretchr/testify/assert"
	"github.com/wandile/smtp-provider/internal/handler/command"
	"github.com/wandile/smtp-provider/internal/repository"
	"github.com/wandile/smtp-provider/internal/repository/exception"
	"testing"
)

func TestInMemorySmtpRepository_DeleteSmtpConfig(t *testing.T) {
	s := repository.NewStore()
	repo := repository.NewInMemorySmtpRepository(*s)

	config := &command.SMTPConfig{
		Host:     "TestHost",
		Port:     12,
		Username: "test@test.com",
		Password: "newPassword",
		SSL:      false,
		Enable:   false,
	}

	id, err := repo.SaveSmtpConfig(config)

	repo.DeleteSmtpConfig(&command.DeleteSMTPConfig{Id: id.Id})
	_, err2 := repo.FindById(id.Id)

	assert.Nil(t, err)
	assert.IsType(t, err2, &exception.ConfigException{})
}

func TestInMemorySmtpRepository_EditSmtpConfig(t *testing.T) {

}

func TestInMemorySmtpRepository_FindByHost(t *testing.T) {
	s := repository.NewStore()
	repo := repository.NewInMemorySmtpRepository(*s)

	config := &command.SMTPConfig{
		Host:     "TestHost",
		Port:     12,
		Username: "test@test.com",
		Password: "newPassword",
		SSL:      false,
		Enable:   false,
	}
	_, err := repo.SaveSmtpConfig(config)
	if err != nil {
		return
	}

	host, err2 := repo.FindByHost(config.Host)

	assert.Nil(t, err2)
	assert.Equal(t, config.Host, host.Host)

}

func TestInMemorySmtpRepository_FindById(t *testing.T) {
	s := repository.NewStore()
	repo := repository.NewInMemorySmtpRepository(*s)

	config := &command.SMTPConfig{
		Host:     "TestHost",
		Port:     12,
		Username: "test@test.com",
		Password: "newPassword",
		SSL:      false,
		Enable:   false,
	}
	id, err := repo.SaveSmtpConfig(config)
	if err != nil {
		return
	}

	host, err2 := repo.FindById(id.Id)

	assert.Nil(t, err2)
	assert.Equal(t, config.Host, host.Host)
}

func TestInMemorySmtpRepository_SaveSmtpConfig(t *testing.T) {

}
