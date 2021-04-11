package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"hackathon/config"
)

var client *redis.Client

func Init() (err error) {
	cfg := config.Conf.RedisConfig
	client = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password:     cfg.Password,
		DB:           cfg.DB,
		PoolSize:     cfg.PoolSize,
		MinIdleConns: cfg.MinIdleConns,
	})

	_, err = client.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

func Close() {
	_ = client.Close()
}
