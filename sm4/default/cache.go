package _default

import (
	"sync"
	"time"
)

type Cache interface {
	Get(k string) ([]byte, bool)
	Set(k string, v []byte)
	GetWriteCache(clear bool) map[string][]byte
	SetWriteCache(k string, v []byte)
}

func newCacheImpl(t time.Duration) Cache {
	c := &cacheImpl{
		read:  make(map[string]val),
		write: make(map[string]val),
		t:     time.NewTicker(t),
	}

	go c.clearAfterTime(t)

	return c
}

type cacheImpl struct {
	read  map[string]val
	write map[string]val
	t     *time.Ticker

	mu sync.RWMutex
}

// clearAfterTime чистит значения кэша которые лежат там больше t и не запрашиваются
func (c *cacheImpl) clearAfterTime(t time.Duration) {
	defer c.t.Stop()
	for range c.t.C {
		func() {
			c.mu.Lock()
			defer c.mu.Unlock()
			for k, v := range c.read {
				if time.Now().Sub(v.lastUsedAt) > t {
					delete(c.read, k)
				}
			}
		}()
	}
}

// SetWriteCache кэширует запись
func (c *cacheImpl) SetWriteCache(k string, v []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.write[k] = val{
		val: v,
	}
}

// GetWriteCache получает все значения для записи и если clr == true то чистит кэш, иначе просто достает значения
func (c *cacheImpl) GetWriteCache(clr bool) map[string][]byte {
	vals := make(map[string][]byte)

	c.mu.Lock()
	for k, v := range c.write {
		vals[k] = v.val
	}
	c.mu.Unlock()

	if clr {
		clear(c.write)
	}

	return vals
}

// Get получет значение и апдейтит время когда оно было в последний раз использовано
func (c *cacheImpl) Get(k string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if v, ok := c.read[k]; ok {
		v.lastUsedAt = time.Now()
		return v.val, true
	}

	return nil, false
}

func (c *cacheImpl) Set(k string, v []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.read[k] = val{
		lastUsedAt: time.Now(),
		insertedAt: time.Now(),
		val:        v,
	}
}

type val struct {
	lastUsedAt time.Time
	insertedAt time.Time
	val        []byte
}
