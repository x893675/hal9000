package mysql

import (
	"github.com/spf13/pflag"
	"hal9000/pkg/utils/reflectutils"
	"time"
)

type MySQLOptions struct {
	Host                  string        `json:"host,omitempty" yaml:"host" description:"MySQL service host address"`
	Port                  string        `json:"port,omitempty" yaml:"port" description:"MySQL service port number"`
	Database              string        `json:"database,omitempty" yaml:"database" description:"MySQL database name"`
	Username              string        `json:"username,omitempty" yaml:"username"`
	Password              string        `json:"-" yaml:"password"`
	MaxIdleConnections    int           `json:"maxIdleConnections,omitempty" yaml:"maxIdleConnections"`
	MaxOpenConnections    int           `json:"maxOpenConnections,omitempty" yaml:"maxOpenConnections"`
	MaxConnectionLifeTime time.Duration `json:"maxConnectionLifeTime,omitempty" yaml:"maxConnectionLifeTime"`
}

func NewMySQLOptions() *MySQLOptions {
	return &MySQLOptions{
		Host:                  "",
		Port:                  "",
		Database:              "",
		Username:              "",
		Password:              "",
		MaxIdleConnections:    100,
		MaxOpenConnections:    100,
		MaxConnectionLifeTime: time.Duration(10) * time.Second,
	}
}

func (m *MySQLOptions) Validate() []error {
	var errors []error

	return errors
}

func (m *MySQLOptions) ApplyTo(options *MySQLOptions) {
	reflectutils.Override(options, m)
}

func (m *MySQLOptions) AddFlags(fs *pflag.FlagSet) {

	fs.StringVar(&m.Host, "mysql-host", m.Host, ""+
		"MySQL service host address. If left blank, the following related mysql options will be ignored.")

	fs.StringVar(&m.Port, "mysql-port", m.Host, ""+
		"MySQL service port number. If left blank, the following related mysql options will be ignored.")

	fs.StringVar(&m.Database, "mysql-db", m.Host, ""+
		"MySQL database name. If left blank, the following related mysql options will be ignored.")

	fs.StringVar(&m.Username, "mysql-username", m.Username, ""+
		"Username for access to mysql service.")

	fs.StringVar(&m.Password, "mysql-password", m.Password, ""+
		"Password for access to mysql, should be used pair with password.")

	fs.IntVar(&m.MaxIdleConnections, "mysql-max-idle-connections", m.MaxOpenConnections, ""+
		"Maximum idle connections allowed to connect to mysql.")

	fs.IntVar(&m.MaxOpenConnections, "mysql-max-open-connections", m.MaxOpenConnections, ""+
		"Maximum open connections allowed to connect to mysql.")

	fs.DurationVar(&m.MaxConnectionLifeTime, "mysql-max-connection-life-time", m.MaxConnectionLifeTime, ""+
		"Maximum connection life time allowed to connecto to mysql.")
}
