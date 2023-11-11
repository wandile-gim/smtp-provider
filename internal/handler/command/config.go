package command

type SMTPConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	SSL      bool
	Enable   bool
}

type UpdateSMTPConfig struct {
	Id       string
	Username string
	Password string
	Enable   bool
}

type EnableSMTPConfig struct {
	Id     string
	Enable bool
}
type DeleteSMTPConfig struct {
	Id string
}
