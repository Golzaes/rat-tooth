package redis

import (
	"fmt"
	"sync"
	"time"

	"github.com/go-redis/redis"
)

//RedisStorage implements the redis storage backend
type RedisStorage struct {
	// Address is the redis server address
	Address string
	// Password is the password for the redis server
	Password string
	// DB is the redis database. Default is 0
	DB int
	// Prefix is an optional string in the keys. It can be used
	// to use one redis database for independent scraping tasks.
	Prefix string
	// Client is the redis connection
	Client *redis.Client

	// Expiration time for Visited keys. After expiration pages
	// are to be visited again.
	Expires time.Duration

	mu sync.RWMutex // Only used for cookie methods.
}

// Init initializes the redis storage
func (s *RedisStorage) Init() error {
	if s.Client == nil {
		s.Client = redis.NewClient(&redis.Options{
			Addr:     s.Address,
			Password: s.Password,
			DB:       s.DB,
		})
	}
	if _, err := s.Client.Ping().Result(); err != nil {
		return fmt.Errorf("redis connection error: %v\n", err)
	}
	return nil
}

// Clear removes all entries from the storage
func (s *RedisStorage) Clear() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	r2 := s.Client.Keys(s.Prefix + ":*")
	keys, err := r2.Result()
	if err != nil {
		return err
	}
	return s.Client.Del(keys...).Err()
}
