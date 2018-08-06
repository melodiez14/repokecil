package conn

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

type RedisConfig struct {
	ConnectionString string
	MaxActive        int
	MaxIdle          int
	IdleTimeout      int
	Wait             bool
}

func InitRedis(cfg RedisConfig) *redis.Pool {
	return &redis.Pool{
		MaxActive:   cfg.MaxActive,
		IdleTimeout: time.Duration(cfg.IdleTimeout) * time.Second,
		MaxIdle:     cfg.MaxIdle,
		Wait:        cfg.Wait,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", cfg.ConnectionString, redis.DialConnectTimeout(time.Duration(cfg.IdleTimeout)*time.Second))
			if err != nil {
				return nil, err
			}
			return c, err
		},
	}
}
