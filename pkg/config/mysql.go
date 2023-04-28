package config

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
)

type MySQL struct {
	Host     string `envconfig:"MYSQL_HOST"`
	Port     int    `envconfig:"MYSQL_PORT"`
	Username string `envconfig:"MYSQL_USERNAME"`
	Password string `envconfig:"MYSQL_PASSWORD"`
	Database string `envconfig:"MYSQL_DATABASE"`
}

func NewMySQLConfig() (MySQL, error) {
	var config MySQL
	if err := envconfig.Process("", &config); err != nil {
		return config, err
	}
	return config, nil
}

func (m MySQL) ConnectionString() string {
	return fmt.Sprintf("%s:%v@tcp(%v:%v)/%v", m.Username, m.Password, m.Host, m.Port, m.Database)
}
