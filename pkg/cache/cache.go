package cache

import (
	"context"
	"encoding/json"
	"log"

	"github.com/gomodule/redigo/redis"
)

type (
	CacheHelper interface {
		UnsetCache(ctx context.Context, key string) error
		Get(key string) ([]byte, error)
		Set(key string, data []byte, ttl int) error
		GetAndUnmarshal(key string, v interface{}) (bool, error)
		MarshalAndSet(key string, ttl int, v interface{}) error
	}

	cacheHelper struct {
		redisPool *redis.Pool
	}

	CacheData struct {
		Key   string
		Field string
		Data  []byte
		TTL   int
	}
)

func NewCacheHelper(ds *redis.Pool) CacheHelper {
	return &cacheHelper{ds}
}

func (c *cacheHelper) UnsetCache(ctx context.Context, key string) error {
	conn := c.redisPool.Get()
	defer conn.Close()

	_, err := conn.Do("DEL", key)
	if err != nil {
		log.Printf("[CACHE-ERROR] unable to invalidate cache %s : %s", key, err.Error())
	}

	return err
}

func (c *cacheHelper) Get(key string) ([]byte, error) {
	conn := c.redisPool.Get()
	defer conn.Close()

	data, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		if err == redis.ErrNil {
			return nil, nil
		}

		return nil, err
	}

	return data, nil
}

func (c *cacheHelper) Set(key string, data []byte, ttl int) error {
	conn := c.redisPool.Get()
	defer conn.Close()

	_, err := conn.Do("SETEX", key, ttl, data)
	return err
}

func (c *cacheHelper) GetAndUnmarshal(key string, v interface{}) (exists bool, err error) {
	cachedBytes, err := c.Get(key)
	if err != nil {
		log.Printf("[CACHE-ERROR] unable to retrieving cache %s : %s", key, err.Error())
	} else {
		if cachedBytes == nil {
			return false, nil
		}

		err = json.Unmarshal(cachedBytes, v)
		if err == nil {
			return true, nil
		}

		log.Printf("[CACHE-ERROR] unable to retrieving cache %s : %s", key, err.Error())
	}

	return false, err
}

func (c *cacheHelper) MarshalAndSet(key string, ttl int, v interface{}) error {
	cachedBytes, err := json.Marshal(v)
	if err != nil {
		log.Printf("[CACHE-ERROR] unable to cache data %s : %s", key, err.Error())
	} else {
		err = c.Set(key, cachedBytes, ttl)
		if err != nil {
			log.Printf("[CACHE-ERROR] unable to cache data %s : %s", key, err.Error())
		}
	}

	return err
}
