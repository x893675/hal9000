package database

import (
	"github.com/spf13/pflag"
	"hal9000/pkg/utils/reflectutils"
	"time"
)

type DatabaseOptions struct {
	Host                  string        `json:"host,omitempty" yaml:"host" description:"database service host address"`
	Port                  string        `json:"port,omitempty" yaml:"port" description:"database service port number"`
	Database              string        `json:"database,omitempty" yaml:"database" description:"database name"`
	Username              string        `json:"username,omitempty" yaml:"username"`
	Password              string        `json:"-" yaml:"password"`
	MaxIdleConnections    int           `json:"maxIdleConnections,omitempty" yaml:"maxIdleConnections"`
	MaxOpenConnections    int           `json:"maxOpenConnections,omitempty" yaml:"maxOpenConnections"`
	MaxConnectionLifeTime time.Duration `json:"maxConnectionLifeTime,omitempty" yaml:"maxConnectionLifeTime"`
}


func NewDatabaseOptions() *DatabaseOptions {
	return &DatabaseOptions{
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

func (m *DatabaseOptions) Validate() []error {
	var errors []error

	//TODO: validate database parameter
	return errors
}

func (m *DatabaseOptions) ApplyTo(options *DatabaseOptions) {
	reflectutils.Override(options, m)
}

func (m *DatabaseOptions) AddFlags(fs *pflag.FlagSet) {

	fs.StringVar(&m.Host, "db-host", m.Host, ""+
		"Database service host address. If left blank, the following related database options will be ignored.")

	fs.StringVar(&m.Port, "db-port", m.Host, ""+
		"Database service port number. If left blank, the following related database options will be ignored.")

	fs.StringVar(&m.Database, "db-name", m.Host, ""+
		"Database database name. If left blank, the following related database options will be ignored.")

	fs.StringVar(&m.Username, "db-username", m.Username, ""+
		"Username for access to database service.")

	fs.StringVar(&m.Password, "db-password", m.Password, ""+
		"Password for access to database, should be used pair with password.")

	fs.IntVar(&m.MaxIdleConnections, "db-max-idle-connections", m.MaxOpenConnections, ""+
		"Maximum idle connections allowed to connect to database.")

	fs.IntVar(&m.MaxOpenConnections, "db-max-open-connections", m.MaxOpenConnections, ""+
		"Maximum open connections allowed to connect to database.")

	fs.DurationVar(&m.MaxConnectionLifeTime, "db-max-connection-life-time", m.MaxConnectionLifeTime, ""+
		"Maximum connection life time allowed to connecto to database.")
}