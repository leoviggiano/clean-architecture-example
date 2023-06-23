package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"

	"clean/lib/log"
)

type Redis struct {
	Client  *redis.Client
	Context context.Context
	Logger  log.Logger
}

func StartRedis(ctx context.Context, address string, logger log.Logger) *Redis {
	return &Redis{
		Context: ctx,
		Logger:  logger,
		Client: redis.NewClient(&redis.Options{
			Addr: address,
		}),
	}
}

func (r *Redis) log(log string) {
	if r.Logger != nil {
		r.Logger.Infof("[REDIS]: %s", log)
	}
}

func (r *Redis) Set(key string, value interface{}, expiration ...time.Duration) error {
	expirationTime := time.Hour * 24 // 1 DAY DEFAULT

	if expiration != nil {
		expirationTime = expiration[0]
	}

	r.log(fmt.Sprintf("[Set] Key: %s\nValue:\n %#v", key, value))

	jsonValue, _ := json.Marshal(value)
	return r.Client.Set(r.Context, key, jsonValue, expirationTime).Err()
}

func (r *Redis) Get(key string, target interface{}) error {
	redisValue, err := r.Client.Get(r.Context, key).Result()
	if err != nil {
		return err
	}

	jsonValue := []byte(redisValue)
	err = json.Unmarshal(jsonValue, &target)
	if err != nil {
		return err
	}

	r.log(fmt.Sprintf("[Get] Key: %s\nValue:\n %#v", key, target))
	return nil
}

func (r *Redis) Delete(key string) error {
	r.log(fmt.Sprintf("[Delete] Key: %s", key))
	return r.Client.Del(r.Context, key).Err()
}
