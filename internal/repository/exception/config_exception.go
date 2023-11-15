package exception

type ConfigException struct {
	Error   error
	Message string
}

func (c ConfigException) ToMessage() string {
	return c.Message
}

func DoesNotExist(err error, message string) *ConfigException {
	return &ConfigException{Message: message, Error: err}
}

func SaveError(err error, message string) *ConfigException {
	return &ConfigException{Message: message, Error: err}
}
