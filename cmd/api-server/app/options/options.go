package options

import (
	"hal9000/pkg/server/config/mysql"
	genericoptions "hal9000/pkg/server/options"
	cliflag "k8s.io/component-base/cli/flag"
)

type ServerRunOptions struct {
	GenericServerRunOptions *genericoptions.ServerRunOptions

	MySQLOptions *mysql.MySQLOptions
}

func NewServerRunOptions() *ServerRunOptions {

	s := ServerRunOptions{
		GenericServerRunOptions: genericoptions.NewServerRunOptions(),
		MySQLOptions:            mysql.NewMySQLOptions(),
	}

	return &s
}

func (s *ServerRunOptions) Flags() (fss cliflag.NamedFlagSets) {

	s.GenericServerRunOptions.AddFlags(fss.FlagSet("generic"))
	s.MySQLOptions.AddFlags(fss.FlagSet("mysql"))

	return fss
}
