package options

import (
	"hal9000/pkg/client/mysql"
	cliflag "k8s.io/component-base/cli/flag"
)

type ServiceRunOptions struct {
	MySQLOptions *mysql.MySQLOptions
	Loglevel     string
}

func NewServiceRunOptions() *ServiceRunOptions {

	s := ServiceRunOptions{
		MySQLOptions: mysql.NewMySQLOptions(),
		Loglevel:     "info",
	}

	return &s
}

func (s *ServiceRunOptions) Flags() (fss cliflag.NamedFlagSets) {

	s.MySQLOptions.AddFlags(fss.FlagSet("mysql"))

	return fss
}
