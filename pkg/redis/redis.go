package redis

import (
	"context"
	"myEntry/pkg/log"
	"time"

	"github.com/go-redis/redis/extra/redisotel/v8"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

// Client redis 客户端
var Client *redis.Client

// ErrRedisNotFound not exist in redis
const ErrRedisNotFound = redis.Nil

// Config redis config
type Config struct {
	Addr         string
	Password     string
	DB           int
	MinIdleConn  int
	DialTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	PoolSize     int
	PoolTimeout  time.Duration
	// tracing switch
	EnableTrace bool
}

// Init 实例化一个redis client
func Init(c *Config) *redis.Client {
	Client = redis.NewClient(&redis.Options{
		Addr:         c.Addr,
		Password:     c.Password,
		DB:           c.DB,
		MinIdleConns: c.MinIdleConn,
		DialTimeout:  c.DialTimeout,
		ReadTimeout:  c.ReadTimeout,
		WriteTimeout: c.WriteTimeout,
		PoolSize:     c.PoolSize,
		PoolTimeout:  c.PoolTimeout,
	})

	_, err := Client.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}

	// hook tracing (using open telemetry)
	if c.EnableTrace {
		Client.AddHook(redisotel.NewTracingHook())
	}

	return Client
}

func Get(key string) (string, error) {
	val, err := Client.Get(ctx, key).Result()
	if err == nil {
		return val, nil
	}
	if err == ErrRedisNotFound {
		log.Warn("Key [%s] is not exist: %s", key)
		return val, nil
	}
	log.Error("Get [%s] error: %s", key, err)
	return val, err
}

func Set(key string, val string) (string, error) {
	val, err := Client.Set(ctx, key, val, -1).Result()
	if err != nil {
		log.Error("Set [%s] error: %s", key, err)
	}
	return val, err
}

func SetEx(key string, val string, expireTime time.Duration) error {
	_, err := Client.SetEX(ctx, key, val, expireTime).Result()
	if err != nil {
		log.Error("SetEx [%s] error: %s", key, err)
	}
	return err
}

func Refresh(key string, expireTime time.Duration) (bool, error) {
	b, err := Client.Expire(ctx, key, expireTime).Result()
	if err != nil {
		log.Error("Refresh [%s] error: %s", key, err)
	}
	return b, err
}

func Del(key string) (int64, error) {
	res, err := Client.Del(ctx, key).Result()
	if err != nil {
		log.Error("Delete [%s] error: %s", key, err)
	}
	return res, err
}
