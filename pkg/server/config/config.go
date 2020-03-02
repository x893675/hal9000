package config

import (
	"fmt"
	"github.com/spf13/viper"
	"hal9000/pkg/client/ldap"
	"hal9000/pkg/client/mysql"
	"hal9000/pkg/client/redis"
	"hal9000/pkg/logger"
)

const (
	DefaultConfigurationName = "config"
	DefaultConfigurationPath = "./"
)

var (
	// sharedConfig holds configuration
	sharedConfig *Config

	// shadowConfig contains options from commandline options
	shadowConfig = &Config{}
)

type Config struct {
	MySQLOptions *mysql.MySQLOptions `json:"mysql,omitempty" yaml:"mysql,omitempty" mapstructure:"mysql"`
	LdapOptions  *ldap.LdapOptions    `json:"ldap,omitempty" yaml:"ldap,omitempty" mapstructure:"ldap"`
	RedisOptions *redis.RedisOptions  `json:"redis,omitempty" yaml:"redis,omitempty" mapstructure:"redis"`
}

func newConfig() *Config {
	return &Config{
		MySQLOptions: mysql.NewMySQLOptions(),
		LdapOptions:  ldap.NewLdapOptions(),
		RedisOptions: redis.NewRedisOptions(),
	}
}

func Get() *Config {
	return sharedConfig
}

func (c *Config) Apply(conf *Config) {
	shadowConfig = conf

	if conf.RedisOptions != nil {
		conf.RedisOptions.ApplyTo(c.RedisOptions)
	}

	if conf.LdapOptions != nil {
		conf.LdapOptions.ApplyTo(c.LdapOptions)
	}

	if conf.MySQLOptions != nil {
		conf.MySQLOptions.ApplyTo(c.MySQLOptions)
	}
}

func (c *Config) stripEmptyOptions() {
	if c.MySQLOptions != nil && c.MySQLOptions.Host == "" {
		c.MySQLOptions = nil
	}

	if c.RedisOptions != nil && c.RedisOptions.RedisURL == "" {
		c.RedisOptions = nil
	}

	if c.LdapOptions != nil && c.LdapOptions.Host == "" {
		c.LdapOptions = nil
	}

}

// Load loads configuration after setup
func Load() error {
	sharedConfig = newConfig()

	viper.SetConfigName(DefaultConfigurationName)
	viper.AddConfigPath(DefaultConfigurationPath)
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			logger.Warn(nil, "configuration file not found")
			return nil
		} else {
			panic(fmt.Errorf("error parsing configuration file %s", err))
		}
	}

	conf := newConfig()
	if err := viper.Unmarshal(conf); err != nil {
		logger.Error(nil, "error unmarshal configuration %v", err)
		return err
	} else {
		conf.Apply(shadowConfig)
		sharedConfig = conf
	}

	return nil
}
