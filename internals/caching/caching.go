package caching

import (
	"fmt"
	"log"
	"slices"
	"time"
)

type ICachingService interface {
	Set(key string, v any) error
	Get(key string) (any, error)
	Del(key string)
}

type CustomExpiringCachingService struct {
	keys              []string
	store             *map[string]any
	expirationMap     *map[string]time.Time
	defaultExpireTime time.Duration
}

var (
	globalCacheStore     = map[string]any{}
	globalExpirationMap  = map[string]time.Time{}
	cacheCleanUpInterval = time.Second * 1
)

func init() {
	fmt.Printf("Starting cache cleanup cycle.\nCurrent cache cleanup interval: %s\n", cacheCleanUpInterval.String())
	go func() {
		for {
			itemsToRemove := []string{}
			for k, exp := range globalExpirationMap {
				if exp.UTC().UnixMilli() < time.Now().UTC().UnixMilli() {
					itemsToRemove = append(itemsToRemove, k)
				}
			}
			for _, k := range itemsToRemove {
				delete(globalCacheStore, k)
				delete(globalExpirationMap, k)
			}
			time.Sleep(cacheCleanUpInterval)
		}
	}()
}

var _ ICachingService = &CustomExpiringCachingService{}

func NewCustomExpiringCachingService(defaultExpireTime time.Duration) *CustomExpiringCachingService {
	cachingService := CustomExpiringCachingService{
		keys:              []string{},
		defaultExpireTime: defaultExpireTime,
		expirationMap:     &globalExpirationMap,
		store:             &globalCacheStore,
	}
	return &cachingService
}

func RetrieveCachedData[T any](service ICachingService, key string) (T, error) {
	var data T
	v, err := service.Get(key)
	if err != nil {
		log.Println("unable to get retrieve data due to:", err)
		return data, err
	}
	if d, ok := v.(T); !ok {
		log.Println("cached data does not match the expected type:", v)
		return data, fmt.Errorf("Unable to convert the cached data. Corrupted cache found.")
	} else {
		data = d
	}
	return data, nil
}

func (c *CustomExpiringCachingService) Set(key string, v any) error {
	(*c.store)[key] = v
	if !slices.Contains(c.keys, key) {
		c.keys = append(c.keys, key)
	}
	(*c.expirationMap)[key] = time.Now().Add(c.defaultExpireTime)
	return nil
}

func (c *CustomExpiringCachingService) Get(key string) (any, error) {
	v, exists := (*c.store)[key]
	if !exists {
		return nil, fmt.Errorf("No cache found or has expired. Key: '%s'", key)
	}
	return v, nil
}

func (c *CustomExpiringCachingService) Del(key string) {
	delete(*c.store, key)
	delete(*c.expirationMap, key)
	c.keys = slices.DeleteFunc(c.keys, func(k string) bool {
		return k == key
	})
}
