package redis_client

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var DB *redis.Client

func InitClient(url, password string, database int) {
	DB = redis.NewClient(&redis.Options{
		Addr:     url,
		Password: password,
		DB:       database,
	})
}

func GetMaxSeq(key string) (int64, error) {
	seq, err := DB.Get(context.Background(), key).Int64()
	if err != nil {
		return -1, err
	}
	return seq, nil
}
