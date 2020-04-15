package webcache

import (
	"time"
)

type Cache interface {
	Get(key string) *CacheEntry
	Save(key string, data []byte, duration time.Duration)
	Invalidate(key string)
	Name() string
}

type CacheEntry struct {
	Data       []byte
	Expiration time.Time
}
