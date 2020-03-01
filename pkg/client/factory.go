package client

import (
	"fmt"
	"hal9000/pkg/client/ldap"
	"hal9000/pkg/client/redis"
	"hal9000/pkg/server/config/mysql"
	"sync"
)

type ClientSetNotEnabledError struct {
	err error
}

func (e ClientSetNotEnabledError) Error() string {
	return fmt.Sprintf("client set not enabled: %v", e.err)
}

var mutex sync.Mutex

type ClientSetOptions struct {
	mySQLOptions        *mysql.MySQLOptions
	redisOptions        *redis.RedisOptions
	ldapOptions         *ldap.LdapOptions
}

func NewClientSetOptions() *ClientSetOptions {
	return &ClientSetOptions{
		mySQLOptions:        mysql.NewMySQLOptions(),
		redisOptions:        redis.NewRedisOptions(),
		ldapOptions:         ldap.NewLdapOptions(),
	}
}

func (c *ClientSetOptions) SetMySQLOptions(options *mysql.MySQLOptions) *ClientSetOptions {
	c.mySQLOptions = options
	return c
}

func (c *ClientSetOptions) SetRedisOptions(options *redis.RedisOptions) *ClientSetOptions {
	c.redisOptions = options
	return c
}

func (c *ClientSetOptions) SetLdapOptions(options *ldap.LdapOptions) *ClientSetOptions {
	c.ldapOptions = options
	return c
}

// ClientSet provide best of effort service to initialize clients,
// but there is no guarantee to return a valid client instance,
// so do validity check before use
type ClientSet struct {
	csoptions *ClientSetOptions
	stopCh    <-chan struct{}

	//mySQLClient *mysql.MySQLClient
	ldapClient          *ldap.LdapClient
	//redisClient         *redis.RedisClient

}

// global clientsets instance
var sharedClientSet *ClientSet

func ClientSets() *ClientSet {
	return sharedClientSet
}


func NewClientSetFactory(c *ClientSetOptions, stopCh <-chan struct{}) *ClientSet {
	sharedClientSet = &ClientSet{csoptions: c, stopCh: stopCh}

	return sharedClientSet
}

func (cs *ClientSet) Ldap() (*ldap.LdapClient, error) {
	var err error

	if cs.csoptions.ldapOptions == nil || cs.csoptions.ldapOptions.Host == "" {
		return nil, ClientSetNotEnabledError{}
	}

	if cs.ldapClient != nil {
		return cs.ldapClient, nil
	} else {
		mutex.Lock()
		defer mutex.Unlock()

		if cs.ldapClient == nil {
			cs.ldapClient, err = ldap.NewLdapClient(cs.csoptions.ldapOptions, cs.stopCh)
			if err != nil {
				return nil, err
			}
		}
		return cs.ldapClient, nil
	}
}