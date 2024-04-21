package caching

import (
	"encoding/json"
	"fmt"
	"slices"
	"time"
)

// TODO: implement auto cache clean up after a specific duration.

type ICachingService interface {
	Set(key string, value any) error
	Get(key string) ([]byte, error)
	Del(key string)
}

type CustomExpiringCachingService struct {
	keys              []string
	store             *map[string][]byte
	expirationMap     map[string]time.Time
	defaultExpireTime time.Duration
}

var globalCacheStore = map[string][]byte{}

var _ ICachingService = &CustomExpiringCachingService{}

func NewCustomExpiringCachingService(defaultExpireTime time.Duration) *CustomExpiringCachingService {
	cachingService := CustomExpiringCachingService{
		keys:              []string{},
		expirationMap:     map[string]time.Time{},
		defaultExpireTime: defaultExpireTime,
		store:             &globalCacheStore,
	}
	return &cachingService
}

func (c *CustomExpiringCachingService) Set(key string, v any) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}
	(*c.store)[key] = b
	if !slices.Contains(c.keys, key) {
		c.keys = append(c.keys, key)
	}
	c.expirationMap[key] = time.Now().UTC().Add(c.defaultExpireTime)
	return nil
}

func (c *CustomExpiringCachingService) Get(key string) ([]byte, error) {
	v, exists := (*c.store)[key]
	if !exists {
		return nil, fmt.Errorf("No cache found with key: '%s'", key)
	}
	if time.Now().UTC().UnixMilli() > c.expirationMap[key].UnixMilli() {
		return nil, fmt.Errorf("Cache expired")
	}

	return v, nil
}

func (c *CustomExpiringCachingService) Del(key string) {
	delete(*c.store, key)
	delete(c.expirationMap, key)
	c.keys = slices.DeleteFunc(c.keys, func(k string) bool {
		return k == key
	})
}
