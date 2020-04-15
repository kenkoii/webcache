package webcache

import (
	"sync"
	"time"
)

type InMemoryCache struct {
	cache map[string]*CacheEntry
	mutex sync.RWMutex
}

func NewInMemoryCache() Cache {
	return &InMemoryCache{
		cache: make(map[string]*CacheEntry),
		mutex: sync.RWMutex{},
	}
}

func (c *InMemoryCache) Save(key string, data []byte, duration time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	expiration := time.Now().Add(time.Duration(duration) * time.Second)

	entry := CacheEntry{
		Data:       data,
		Expiration: expiration,
	}

	c.cache[key] = &entry
}

func (c *InMemoryCache) Invalidate(key string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	delete(c.cache, key)
}

func (c *InMemoryCache) Name() string {
	return "inmemory"
}

func (c *InMemoryCache) Get(key string) *CacheEntry {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	if entry, ok := c.cache[key]; ok {
		return entry
	}
	return nil
}
