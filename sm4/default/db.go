package _default

import (
	"sync"
	"time"
)

func newDbImpl(cache Cache) *dbImpl {
	return &dbImpl{cache: cache, dbs: make(map[string][]byte)}
}

type dbImpl struct {
	cache Cache
	dbs   map[string][]byte

	mu sync.Mutex
}

func (d *dbImpl) Get(k string) ([]byte, bool) {
	v, ok := d.cache.Get(k)
	if ok {
		return v, ok
	}

	d.mu.Lock()
	defer d.mu.Unlock()
	v, ok = d.dbs[k]
	if !ok {
		return v, false
	}

	d.cache.Set(k, v)
	return v, true
}

// Insert обычный инсерт
func (d *dbImpl) Insert(k string, val []byte) {
	d.dbs[k] = val
}

// CachedInsert пишет в кэш, где потом записывается в 'бд' по стратегии (ScheduleInsert или LimitInsert)
func (d *dbImpl) CachedInsert(k string, val []byte) {
	d.cache.SetWriteCache(k, val)
}

// Сделал кэширование на запись, где-то читал что в том же постгресе можно делать bunch инсертов и это будет
// эффективнее.

// ScheduleInsert пишет по тикам тикера значения из write cache
func (d *dbImpl) ScheduleInsert(t *time.Ticker) {
	defer t.Stop()
	for range t.C {
		for k, v := range d.cache.GetWriteCache(true) {
			d.Insert(k, v)
		}
	}
}

// LimitInsert пишет когда write cache достигает лимита
func (d *dbImpl) LimitInsert(lim int) {
	t := time.NewTicker(time.Second)

	defer t.Stop()
	for range t.C {
		if len(d.cache.GetWriteCache(false)) == lim {
			for k, v := range d.cache.GetWriteCache(true) {
				d.Insert(k, v)
			}
		}
	}
}
