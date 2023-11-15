package repository

import (
	"context"
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"github.com/wandile/smtp-provider/internal/handler/command"
	"github.com/wandile/smtp-provider/internal/repository"
	"github.com/wandile/smtp-provider/test"
	"strconv"

	"testing"
)

func TestEntSmtpRepository_DeleteSmtpConfig(t *testing.T) {
	ctx := context.Background()
	client := test.DbConnection(t)
	defer client.Close()

	repo := repository.NewEntSmtpRepository(client)

	// given
	config := &command.SMTPConfig{
		Host:     "TestHost",
		Port:     12,
		Username: "test@test.com",
		Password: "newPassword",
		SSL:      false,
		Enable:   false,
	}

	id, _ := repo.SaveSmtpConfig(config)
	all, _ := client.Configuration.Query().All(ctx)

	// expected
	assert.Equal(t, len(all), 1)
	repo.DeleteSmtpConfig(&command.DeleteSMTPConfig{Id: id.Id})

	//_, err := repo.FindById(id.Id)
	//assert.NotNil(t, err)
}

func TestEntSmtpRepository_EditSmtpConfig(t *testing.T) {
	client := test.DbConnection(t)
	defer client.Close()

	repo := repository.NewEntSmtpRepository(client)
	config := command.SMTPConfig{
		Host:     "TestHost",
		Port:     12,
		Username: "test@test.com",
		Password: "newPassword",
		SSL:      false,
		Enable:   false,
	}
	id, _ := repo.SaveSmtpConfig(&config)

	updateC := &command.UpdateSMTPConfig{
		Id:       id.Id,
		Username: "newTest@test.com",
		Password: config.Password,
		Enable:   true,
	}
	repo.UpdateSmtpConfig(*id, updateC)
	updatedConfig, _ := repo.FindById(id.Id)
	assert.Equal(t, updatedConfig.Username, updateC.Username)

}

func TestEntSmtpRepository_FindAll(t *testing.T) {
	client := test.DbConnection(t)
	defer client.Close()

	repo := repository.NewEntSmtpRepository(client)
	config := command.SMTPConfig{
		Host:     "TestHost",
		Port:     12,
		Username: "test@test.com",
		Password: "newPassword",
		SSL:      false,
		Enable:   false,
	}

	for i := range [10]int{} {
		conf := config
		conf.Username = conf.Username + " " + strconv.Itoa(i)
		repo.SaveSmtpConfig(&conf)
	}

	assert.Equal(t, len(repo.FindAll()), 10)
}

func TestEntSmtpRepository_FindByHost(t *testing.T) {
	client := test.DbConnection(t)
	defer client.Close()

	repo := repository.NewEntSmtpRepository(client)
	config := &command.SMTPConfig{
		Host:     "TestHost",
		Port:     12,
		Username: "test@test.com",
		Password: "newPassword",
		SSL:      false,
		Enable:   false,
	}
	repo.SaveSmtpConfig(config)

	found, _ := repo.FindByHost(config.Host)

	assert.Equal(t, config.Host, found.Host)
}

func TestEntSmtpRepository_FindById(t *testing.T) {
	client := test.DbConnection(t)
	defer client.Close()

	repo := repository.NewEntSmtpRepository(client)
	config := &command.SMTPConfig{
		Host:     "TestHost",
		Port:     12,
		Username: "test@test.com",
		Password: "newPassword",
		SSL:      false,
		Enable:   false,
	}
	id, _ := repo.SaveSmtpConfig(config)
	uid, _ := uuid.Parse(id.Id)

	found, _ := repo.FindById(uid.String())

	assert.Equal(t, config.Host, found.Host)
}

func TestEntSmtpRepository_SaveSmtpConfig(t *testing.T) {
	client := test.DbConnection(t)
	defer client.Close()

	repo := repository.NewEntSmtpRepository(client)
	config := &command.SMTPConfig{
		Host:     "TestHost",
		Port:     12,
		Username: "test@test.com",
		Password: "newPassword",
		SSL:      false,
		Enable:   false,
	}
	id, _ := repo.SaveSmtpConfig(config)
	uid, _ := uuid.Parse(id.Id)

	get, err := client.Configuration.Get(context.Background(), uid)

	if err != nil {
		log.Error(err.Error())
		return
	}

	assert.Equal(t, get.Username, config.Username)
	assert.NotNil(t, get.ID)

}

func TestEntSmtpRepository_configEntityToDomainModel(t *testing.T) {

}
