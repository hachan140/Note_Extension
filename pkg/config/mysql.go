package config

import "fmt"

type MySQL struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

func (m MySQL) ConnectionString() string {
	return fmt.Sprintf("%s:%v@tcp(%v:%v)/%v", m.Username, m.Password, m.Host, m.Port, m.Database)
}
