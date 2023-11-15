package application

import (
	"github.com/stretchr/testify/assert"
	"github.com/wandile/smtp-provider/internal/application"
	"github.com/wandile/smtp-provider/internal/handler/command"
	"github.com/wandile/smtp-provider/internal/repository"
	"log"
	"os"
	"testing"
)

var smtpConfig = &command.SMTPConfig{
	Host:     "smtp.gmail.com",
	Port:     587,
	Username: os.Getenv("PLATFORM_EMAIL"),
	Password: os.Getenv("PLATFORM_PASSWORD"),
	SSL:      false,
	Enable:   false,
}

func TestConfigService_ChangeSmtpConfiguration(t *testing.T) {
	s := repository.NewStore()
	repo := repository.NewInMemorySmtpRepository(*s)
	service := application.NewConfigService(repo)

	// given
	config, _ := repo.SaveSmtpConfig(smtpConfig)
	updateSMTPConfig := &command.UpdateSMTPConfig{
		Id:       config.Id,
		Username: os.Getenv("WANDILE_EMAIL"),
		Password: os.Getenv("WANDILE_PASSWORD"),
		Enable:   false,
	}

	// when
	_, err := service.ChangeSmtpConfiguration(updateSMTPConfig)
	log.Println(">>>err", err)
	id, _ := repo.FindById(config.Id)

	// then
	assert.Nil(t, err)
	assert.Equal(t, updateSMTPConfig.Username, id.Username)
}

func TestConfigService_CreateSmtpConfiguration(t *testing.T) {
	s := repository.NewStore()
	repo := repository.NewInMemorySmtpRepository(*s)
	service := application.NewConfigService(repo)

	// given
	config := smtpConfig

	// when
	configuration, err := service.CreateSmtpConfiguration(config)
	if err != nil {
		log.Println(err)
	}

	id, err := repo.FindById(configuration.Id.Id)

	// then
	assert.Nil(t, err)
	assert.Equal(t, id.ConfigId, configuration.Id)
}

//func TestConfigService_EnableSmtpConfiguration(t *testing.T) {
//	s := repository.NewStore()
//	repo := repository.NewInMemorySmtpRepository(*s)
//	service := application.NewConfigService(repo)
//
//	// given
//	var dummy = smtpConfig
//	var i = 10
//	var targetId domain.ConfigId
//
//	for idx := 0; idx <= i; idx++ {
//		dummy.Username += strconv.Itoa(idx)
//		config, _ := repo.SaveSmtpConfig(dummy)
//		if idx == 10 {
//			targetId = *config
//		}
//	}
//
//	// when
//	configuration, _ := service.EnableSmtpConfiguration(&command.EnableSMTPConfig{
//		Id:     targetId.Id,
//		Enable: true,
//	})
//
//	// then
//	id, _ := repo.FindById(configuration.Id.Id)
//	id2 := repo.FindAll()[0]
//
//	assert.Equal(t, id.Enable, true)
//	assert.Equal(t, id2.Enable, false)
//}

func TestNewConfigService(t *testing.T) {

}
