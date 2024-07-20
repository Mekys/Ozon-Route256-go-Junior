package cache

import (
	"sync"
	"time"
)

// NewTTLClient creates imdb client with expired value
func NewTTLClient[K comparable, V any](ttl time.Duration) *TTLClient[K, V] {
	return &TTLClient[K, V]{
		ttl:  ttl,
		data: make(map[K]*Cached[V]),
	}
}

// TTLClient ...
type TTLClient[K comparable, V any] struct {
	ttl  time.Duration
	lock sync.RWMutex
	data map[K]*Cached[V]
}

// Get ...
func (c *TTLClient[K, V]) Get(key K) (V, bool) {
	c.lock.RLock()
	v, ok := c.data[key]
	c.lock.RUnlock()

	if ok && !v.Expired(time.Now()) {
		return v.Value(), true
	}

	return (&Cached[V]{}).Value(), false
}

// Clear
func (c *TTLClient[K, V]) Clear() {
	c.data = make(map[K]*Cached[V])
}

// Set ...
func (c *TTLClient[K, V]) Set(key K, value V, now time.Time) {
	wrapped := NewCached[V](now.Add(c.ttl), value)
	c.lock.Lock()
	c.data[key] = wrapped
	c.lock.Unlock()
}

// NewCached ...
func NewCached[V any](expiredAt time.Time, value V) *Cached[V] {
	return &Cached[V]{
		expiredAt: expiredAt,
		value:     value,
	}
}

// Cached ...
type Cached[V any] struct {
	expiredAt time.Time
	value     V
}

// Expired ...
func (c *Cached[V]) Expired(now time.Time) bool {
	return c.expiredAt.Before(now)
}

// Value ...
func (c *Cached[V]) Value() V {
	return c.value
}
