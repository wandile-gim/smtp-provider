package configuration

import (
	"crypto/tls"
	"gopkg.in/gomail.v2"
	"log"
	"net/smtp"
	"source.clobot.co.kr/spot-team/service/smtp-provider/internal/handler/command"
)

type Configs struct {
	configs []ConfigId
}

func (c *Configs) JoinConfig(configId ConfigId) {
	c.configs = append(c.configs, configId)
}

func (c *Configs) CheckAllAvailable() {

}

type ConfigId struct {
	Id string
}

type SmtpConfig struct {
	ConfigId ConfigId
	// Host represents the host of the SMTP server.
	Host string
	// Port represents the port of the SMTP server.
	Port int
	// Username is the username to use to authenticate to the SMTP server.
	Username string
	// Password is the password to use to authenticate to the SMTP server.
	Password string
	// Auth represents the authentication mechanism used to authenticate to the
	// SMTP server.
	Auth smtp.Auth
	// SSL defines whether an SSL connection is used. It should be false in
	// most cases since the authentication mechanism should use the STARTTLS
	// extension instead.
	SSL bool
	// TSLConfig represents the TLS configuration used for the TLS (when the
	// STARTTLS extension is used) or SSL connection.
	TLSConfig *tls.Config
	// LocalName is the hostname sent to the SMTP server with the HELO command.
	// By default, "localhost" is sent.
	LocalName string
	Activate  bool
	Enable    bool
}

func NewSmtpConfig(port int, ssl bool, host, username, password string) *SmtpConfig {
	c := &SmtpConfig{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
		SSL:      ssl,
	}
	return c
}

func RestoreSmtpConfig(id *ConfigId, port int, enable bool, host, username, password string) *SmtpConfig {
	c := &SmtpConfig{
		ConfigId: *id,
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
		Enable:   enable,
	}
	return c
}

func (c *SmtpConfig) AuthenticatesAccount(config *command.SMTPConfig) error {
	dialer := gomail.NewDialer(config.Host, config.Port, config.Username, config.Password)
	log.Println(">>> AuthenticatesAccount dial begins...")
	_, err := dialer.Dial()
	log.Println(">>> AuthenticatesAccount dial ends...")
	if err != nil {
		return err
	}
	return nil
}

// ChangeUser SMTP 유저 아이디와 비밀번호를 변경하고 변경하려는 정보를 SMTP 서버에 검증합니다.
func (c *SmtpConfig) ChangeUser(config *command.UpdateSMTPConfig) error {
	cmd := &command.SMTPConfig{
		Username: config.Username,
		Password: config.Password,
		Port:     c.Port,
		Host:     c.Host,
	}
	err := c.AuthenticatesAccount(cmd)
	if err != nil {
		return err
	}

	c.Username = config.Username
	c.Password = config.Password
	log.Println(c.Username, c.Password)

	return nil
}

func (c *SmtpConfig) EnableConfig(enable bool) {
	c.Enable = enable
}
