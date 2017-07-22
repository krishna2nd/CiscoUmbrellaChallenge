package ucache

import (
	"config"
	"github.com/go-redis/redis"
)

var client *redis.Client

func init() {
	config := config.Cache()
	client = redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.PassWord,
		DB:       config.Database,
	})
}

// Set bad urls will be added to cache
func Set(url string) error {
	return client.Set(url, true, 0).Err()
}

// Get will fetch the client requested url if present from cache
func Get(url string) bool {
	_, err := client.Get(url).Result()
	if err == redis.Nil {
		return false
	}
	return true
}
