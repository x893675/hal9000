package ldap

import (
	"fmt"
	"github.com/go-ldap/ldap/v3"
	"hal9000/pkg/logger"
)

type LdapClient struct {
	pool    Pool
	options *LdapOptions
}

// panic if cannot connect to ldap service
func NewLdapClient(options *LdapOptions, stopCh <-chan struct{}) (*LdapClient, error) {
	pool, err := NewChannelPool(8, 64, "hal9000", func(s string) (ldap.Client, error) {
		conn, err := ldap.Dial("tcp", options.Host)
		if err != nil {
			return nil, err
		}
		return conn, nil
	}, []uint16{ldap.LDAPResultAdminLimitExceeded, ldap.ErrorNetwork})

	if err != nil {
		logger.Error(nil, err.Error())
		pool.Close()
		return nil, err
	}

	client := &LdapClient{
		pool:    pool,
		options: options,
	}

	go func() {
		<-stopCh
		if client.pool != nil {
			client.pool.Close()
		}
	}()

	return client, nil
}

func (l *LdapClient) NewConn() (ldap.Client, error) {
	if l.pool == nil {
		err := fmt.Errorf("ldap connection pool is not initialized")
		logger.Error(nil, err.Error())
		return nil, err
	}

	conn, err := l.pool.Get()
	// cannot connect to ldap server or pool is closed
	if err != nil {
		logger.Error(nil, err.Error())
		return nil, err
	}
	err = conn.Bind(l.options.ManagerDN, l.options.ManagerPassword)
	if err != nil {
		conn.Close()
		logger.Error(nil, err.Error())
		return nil, err
	}
	return conn, nil
}

func (l *LdapClient) GroupSearchBase() string {
	return l.options.GroupSearchBase
}

func (l *LdapClient) UserSearchBase() string {
	return l.options.UserSearchBase
}