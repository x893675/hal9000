package options


import (
	"hal9000/pkg/client/database"
	cliflag "k8s.io/component-base/cli/flag"
)

type AuthServiceOptions struct {
	Loglevel        string
	DatabaseOptions *database.DatabaseOptions
}

func NewAuthServiceOptions() *AuthServiceOptions {
	s := &AuthServiceOptions{
		Loglevel:        "info",
		DatabaseOptions: database.NewDatabaseOptions(),
	}
	return s
}

func (a *AuthServiceOptions) Flags() (fss cliflag.NamedFlagSets) {
	fs := fss.FlagSet("generic")
	fs.StringVar(&a.Loglevel, "loglevel", a.Loglevel, "server log level, e.g. debug,info")
	a.DatabaseOptions.AddFlags(fss.FlagSet("database"))
	return fss
}