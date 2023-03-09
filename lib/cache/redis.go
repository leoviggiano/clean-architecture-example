package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"

	"clean/lib/log"
)

type Redis struct {
	Client  *redis.Client
	Context context.Context
}

func StartRedis(ctx context.Context, address string) *Redis {
	return &Redis{
		Context: ctx,
		Client: redis.NewClient(&redis.Options{
			Addr: address,
		}),
	}
}

func (r *Redis) Set(key string, value interface{}, expiration ...int) error {
	expirationTime := time.Minute * time.Duration(24*60) // 1 DAY DEFAULT

	if expiration != nil {
		expirationTime = time.Minute * time.Duration(expiration[0])
	}

	jsonValue, _ := json.Marshal(value)
	err := r.Client.Set(r.Context, key, jsonValue, expirationTime).Err()
	if err != nil {
		log.Errorf("[Set Redis] Error in set redis [%s]:[%s]\n%s", key, value, err)
	}
	return err
}

func (r *Redis) Get(key string, target interface{}) (err error) {
	redisValue, err := r.Client.Get(r.Context, key).Result()
	if err != nil {
		return
	}

	jsonValue := []byte(redisValue)
	err = json.Unmarshal(jsonValue, &target)
	return
}

func (r *Redis) Delete(key string) error {
	err := r.Client.Del(r.Context, key).Err()
	return err
}
