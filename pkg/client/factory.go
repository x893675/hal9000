package client

import (
	"fmt"
	goredis "github.com/go-redis/redis"
	"hal9000/pkg/client/database"
	"hal9000/pkg/client/ldap"
	"hal9000/pkg/client/redis"
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
	databaseOptions *database.DatabaseOptions
	redisOptions    *redis.RedisOptions
	ldapOptions     *ldap.LdapOptions
}

func NewClientSetOptions() *ClientSetOptions {
	return &ClientSetOptions{
		databaseOptions: database.NewDatabaseOptions(),
		redisOptions:    redis.NewRedisOptions(),
		ldapOptions:     ldap.NewLdapOptions(),
	}
}

func (c *ClientSetOptions) SetDatabaseOptions(options *database.DatabaseOptions) *ClientSetOptions {
	c.databaseOptions = options
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

	database    *database.Database
	ldapClient  *ldap.LdapClient
	redisClient *redis.RedisClient
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

func (cs *ClientSet) Redis() (*goredis.Client, error) {
	var err error

	if cs.csoptions.redisOptions == nil || cs.csoptions.redisOptions.RedisURL == "" {
		return nil, ClientSetNotEnabledError{}
	}

	if cs.redisClient != nil {
		return cs.redisClient.Redis(), nil
	} else {
		mutex.Lock()
		defer mutex.Unlock()
		if cs.redisClient == nil {
			cs.redisClient, err = redis.NewRedisClient(cs.csoptions.redisOptions, cs.stopCh)
			if err != nil {
				return nil, err
			}
		}

		return cs.redisClient.Redis(), nil
	}
}

func (cs *ClientSet) Database() (*database.Database, error) {
	var err error

	if cs.csoptions.databaseOptions == nil || cs.csoptions.databaseOptions.Host == "" {
		return nil, ClientSetNotEnabledError{}
	}

	if cs.database != nil {
		return cs.database, nil
	} else {
		mutex.Lock()
		defer mutex.Unlock()

		if cs.database == nil {
			cs.database, err = database.NewDataBase(cs.csoptions.databaseOptions, cs.stopCh)
			if err != nil {
				return nil, err
			}
		}
		return cs.database, nil
	}
}
