package webcache

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewInMemoryCache(t *testing.T) {
	t.Run("Create new inmemory cache", func(t *testing.T) {
		c := NewInMemoryCache()
		name := c.Name()
		assert.NotNil(t, c)
		assert.Equal(t, name, "inmemory")
	})
}

func TestCacheOperations(t *testing.T) {
	t.Run("Set Get Invalidate", func(t *testing.T) {
		c := NewInMemoryCache()

		key := "top-secret"
		value := []byte("test test test")
		durationInSecs := time.Duration(60)
		c.Save(key, value, durationInSecs)

		ce := c.Get(key)
		assert.Equal(t, ce.Data, value)
		assert.NotNil(t, ce.Expiration)

		c.Invalidate(key)
		assert.Nil(t, c.Get(key))
	})

	t.Run("Get something that doesnt exist", func(t *testing.T) {
		c := NewInMemoryCache()

		ce := c.Get("some-random-key")

		assert.Nil(t, ce)
	})
}
