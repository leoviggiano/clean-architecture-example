package cache

type Cache interface {
	Set(key string, value interface{}, expiration ...int) error
	Get(key string, target interface{}) error
	Delete(key string) error
}
