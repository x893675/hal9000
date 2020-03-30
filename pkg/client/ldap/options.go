package ldap

import (
	"github.com/spf13/pflag"
	"hal9000/pkg/utils/reflectutils"
)

type LdapOptions struct {
	Host            string `json:"host,omitempty" yaml:"host"`
	ManagerDN       string `json:"managerDN,omitempty" yaml:"managerDN"`
	ManagerPassword string `json:"managerPassword,omitempty" yaml:"managerPassword"`
	UserSearchBase  string `json:"userSearchBase,omitempty" yaml:"userSearchBase"`
	GroupSearchBase string `json:"groupSearchBase,omitempty" yaml:"groupSearchBase"`
}

func NewLdapOptions() *LdapOptions {
	return &LdapOptions{
		Host:            "",
		ManagerDN:       "cn=admin,dc=example,dc=org",
		UserSearchBase:  "ou=users,dc=example,dc=org",
		GroupSearchBase: "ou=groups,dc=example,dc=org",
	}
}

func (l *LdapOptions) Validate() []error {
	var errors []error

	return errors
}

func (l *LdapOptions) ApplyTo(options *LdapOptions) {
	if l.Host != "" {
		reflectutils.Override(options, l)
	}
}

func (l *LdapOptions) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&l.Host, "ldap-host", l.Host, ""+
		"Ldap service host, if left blank, all of the following ldap options will "+
		"be ignored and ldap will be disabled.")

	fs.StringVar(&l.ManagerDN, "ldap-manager-dn", l.ManagerDN, ""+
		"Ldap manager account domain name.")

	fs.StringVar(&l.ManagerPassword, "ldap-manager-password", l.ManagerPassword, ""+
		"Ldap manager account password.")

	fs.StringVar(&l.UserSearchBase, "ldap-user-search-base", l.UserSearchBase, ""+
		"Ldap user search base.")

	fs.StringVar(&l.GroupSearchBase, "ldap-group-search-base", l.GroupSearchBase, ""+
		"Ldap group search base.")
}
