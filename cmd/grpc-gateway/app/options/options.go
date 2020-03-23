package options

import (
	genericoptions "hal9000/pkg/httpserver/options"
	cliflag "k8s.io/component-base/cli/flag"
)

type ServerRunOptions struct {
	GenericServerRunOptions *genericoptions.ServerRunOptions
	Debug                   bool
	Loglevel                string
}

func NewServerRunOptions() *ServerRunOptions {

	s := ServerRunOptions{
		GenericServerRunOptions: genericoptions.NewServerRunOptions(),
		Loglevel:                "info",
		Debug:                   true,
	}

	return &s
}

func (s *ServerRunOptions) Flags() (fss cliflag.NamedFlagSets) {

	fs := fss.FlagSet("generic")
	s.GenericServerRunOptions.AddFlags(fs)

	fs.StringVar(&s.Loglevel, "loglevel", s.Loglevel, "info server log level, e.g. debug,info")
	fs.BoolVar(&s.Debug, "debug", s.Debug, "server run mod, true for devel")
	return fss
}
