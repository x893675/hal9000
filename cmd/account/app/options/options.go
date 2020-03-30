package options

import (
	"hal9000/pkg/client/database"
	cliflag "k8s.io/component-base/cli/flag"
)

type AccountServiceOptions struct {
	Loglevel        string
	DatabaseOptions *database.DatabaseOptions
}

func NewAccountServiceOptions() *AccountServiceOptions {
	s := &AccountServiceOptions{
		Loglevel:        "info",
		DatabaseOptions: database.NewDatabaseOptions(),
	}
	return s
}

func (a *AccountServiceOptions) Flags() (fss cliflag.NamedFlagSets) {
	fs := fss.FlagSet("generic")
	fs.StringVar(&a.Loglevel, "loglevel", a.Loglevel, "server log level, e.g. debug,info")
	a.DatabaseOptions.AddFlags(fss.FlagSet("database"))
	return fss
}
