package cache

import (
	"log"
	"os"

	"bareksa-take-home-test-michael-koh/config"

	"github.com/gomodule/redigo/redis"
)

func CreatePool(cfg config.Config) *redis.Pool {
	return &redis.Pool{
		MaxIdle:   cfg.RedisMaxIdleConnection,
		MaxActive: cfg.RedisMaxActiveConnection,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial(cfg.RedisNetwork, cfg.RedisPort)
			if err != nil {
				log.Printf("ERROR: fail initializing the redis pool: %s", err.Error())
				os.Exit(1)
			}
			return conn, err
		},
	}
}
