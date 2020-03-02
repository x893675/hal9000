package redis

import (
	"github.com/go-redis/redis"
	"hal9000/pkg/logger"
)

type RedisClient struct {
	client *redis.Client
}

func NewRedisClientOrDie(options *RedisOptions, stopCh <-chan struct{}) *RedisClient {
	client, err := NewRedisClient(options, stopCh)
	if err != nil {
		panic(err)
	}

	return client
}

func NewRedisClient(option *RedisOptions, stopCh <-chan struct{}) (*RedisClient, error) {
	var r RedisClient

	options, err := redis.ParseURL(option.RedisURL)

	if err != nil {
		logger.Error(nil, err.Error())
		return nil, err
	}

	r.client = redis.NewClient(options)

	if err := r.client.Ping().Err(); err != nil {
		logger.Error(nil, "unable to reach redis host", err)
		r.client.Close()
		return nil, err
	}

	if stopCh != nil {
		go func() {
			<-stopCh
			if err := r.client.Close(); err != nil {
				logger.Error(nil, err.Error())
			}
		}()
	}

	return &r, nil
}

func (r *RedisClient) Redis() *redis.Client {
	return r.client
}
