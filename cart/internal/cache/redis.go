package cache

import (
	"OMS/cart/config"
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/rs/zerolog/log"
)

type Redis struct {
	RedisClient *redis.Client
}

func (r *Redis) NewSession(config *config.Config) {
	session, e := newRedisSession(config)
	if e != nil {
		log.Fatal().Err(e).Msg("Redis Client Initialization failed")
	}
	r.RedisClient = session
}

func (r *Redis) Get(key string) (string, error) {
	get := r.RedisClient.Get(context.Background(), key)
	s, err := get.Result()
	return s, err
}

func (r *Redis) Set(key string, value interface{}, exp time.Duration) error {
	err := r.RedisClient.Set(context.Background(), key, value, exp).Err()
	return err
}

func newRedisSession(config *config.Config) (*redis.Client, error) {

	redisDB := redis.NewClient(&redis.Options{
		Addr:         config.RedisAddress,
		Password:     config.RedisPassword,
		DB:           config.RedisDB,
		DialTimeout:  config.RedisDialTimeoutMs * time.Millisecond,
		ReadTimeout:  config.RedisReadTimeoutMs * time.Millisecond,
		WriteTimeout: config.RedisWriteTimeoutMs * time.Millisecond,
	})

	err := redisDB.Ping(context.Background()).Err()
	if err != nil {
		return nil, fmt.Errorf("error while initiating connection for redis: %v", err.Error())
	}

	return redisDB, nil
}
