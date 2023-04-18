package cache

import "time"

type Cache interface {
	Set(key string, value interface{}, expiration ...time.Duration) error
	Get(key string, target interface{}) error
	Delete(key string) error
}
