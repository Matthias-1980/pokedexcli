package pokecache

import (
	"time"
	"sync"
)

type Cache struct {
	mapCache  map[string]cacheEntry
	mu sync.Mutex
}

func (c *Cache) Add(key string, value []byte) {
	caEntry := cacheEntry{time.Now(), value}
	c.mu.Lock()
	c.mapCache[key] = caEntry
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	retVal, ok := c.mapCache[key]
	c.mu.Unlock()
	if ok {
		return retVal.val, true
	} 
	return nil, false		
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		c.mu.Lock()
		for key, value := range c.mapCache {
			if value.createdAt.Add(interval).Before(time.Now()) {
				delete(c.mapCache, key)
			}
		}
		c.mu.Unlock()
	}
}

type cacheEntry struct {
	createdAt time.Time
	val []byte
}

func NewCache(interval time.Duration) *Cache {
	var myCache = Cache{
		mapCache: make(map[string]cacheEntry),
	}
	go myCache.reapLoop(interval)
	return &myCache
}

