package options

import (
	"fmt"
	"github.com/spf13/pflag"
	"hal9000/pkg/addr"
)

type ServerRunOptions struct {
	// server bind address
	BindAddress string

	// insecure port number
	InsecurePort int

	// secure port number
	SecurePort int

	// tls cert file
	TlsCertFile string

	// tls private key file
	TlsPrivateKey string
}

func NewServerRunOptions() *ServerRunOptions {
	// create default server run options
	s := ServerRunOptions{
		BindAddress:   "0.0.0.0",
		InsecurePort:  8080,
		SecurePort:    0,
		TlsCertFile:   "",
		TlsPrivateKey: "",
	}

	return &s
}

func (s *ServerRunOptions) Validate() []error {
	var errs []error

	if s.SecurePort == 0 && s.InsecurePort == 0 {
		errs = append(errs, fmt.Errorf("insecure and secure port can not be disabled at the same time"))
	}

	if addr.IsValidPort(s.SecurePort) {
		if s.TlsCertFile == "" {
			errs = append(errs, fmt.Errorf("tls cert file is empty while secure serving"))
		}

		if s.TlsPrivateKey == "" {
			errs = append(errs, fmt.Errorf("tls private key file is empty while secure serving"))
		}
	}

	return errs
}

func (s *ServerRunOptions) AddFlags(fs *pflag.FlagSet) {

	fs.StringVar(&s.BindAddress, "bind-address", "0.0.0.0", "server bind address")
	fs.IntVar(&s.InsecurePort, "insecure-port", 8080, "insecure port number")
	fs.IntVar(&s.SecurePort, "secure-port", 0, "secure port number")
	fs.StringVar(&s.TlsCertFile, "tls-cert-file", "", "tls cert file")
	fs.StringVar(&s.TlsPrivateKey, "tls-private-key", "", "tls private key")
}
