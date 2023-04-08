package conf

import (
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"time"
)

var Redis *redis.Client
var ExpirationTime = 30 * 24 * 60 * 60 * time.Second

type RedisClient struct {
}

func InitRedis() (*RedisClient, error) {
	Redis = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.url"),
		Password: viper.GetString("redis.password"),
		DB:       0,
	})

	_, err := Redis.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}

	return &RedisClient{}, nil

}

func (rc *RedisClient) Set(key string, value string) error {
	return Redis.Set(context.Background(), key, value, ExpirationTime).Err()
}

func (rc *RedisClient) Get(key string) (any, error) {
	return Redis.Get(context.Background(), key).Result()
}

func (rc *RedisClient) Delete(key ...string) error {
	return Redis.Del(context.Background(), key...).Err()
}
