package options

import (
	"hal9000/pkg/client/database"
	genericoptions "hal9000/pkg/httpserver/options"
	cliflag "k8s.io/component-base/cli/flag"
)

type ServerRunOptions struct {
	GenericServerRunOptions *genericoptions.ServerRunOptions

	MySQLOptions *database.DatabaseOptions

	Loglevel string
}

func NewServerRunOptions() *ServerRunOptions {

	s := ServerRunOptions{
		GenericServerRunOptions: genericoptions.NewServerRunOptions(),
		MySQLOptions:            database.NewDatabaseOptions(),
		Loglevel:                "info",
	}

	return &s
}

func (s *ServerRunOptions) Flags() (fss cliflag.NamedFlagSets) {

	s.GenericServerRunOptions.AddFlags(fss.FlagSet("generic"))
	s.MySQLOptions.AddFlags(fss.FlagSet("mysql"))

	return fss
}
