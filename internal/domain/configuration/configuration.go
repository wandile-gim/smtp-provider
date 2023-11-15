package configuration

type ConfigChanged struct {
	Id      ConfigId
	Enabled bool
}

type ConfigCreated struct {
	Id   ConfigId
	Host string
}
