package config

type MariaDB struct {
	UserName string `env:"USERNAME"`
	Password string `env:"PASSWORD"`
	Host     string `env:"HOST"`
	Port     string `env:"PORT"`
	Schema   string `env:"SCHEMA"`
}
