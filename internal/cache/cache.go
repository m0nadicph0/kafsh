package cache

import "sync"

type cache struct {
	cache map[string][]string
	mu    sync.RWMutex
}

func NewCache() *cache {
	return &cache{
		cache: make(map[string][]string),
	}
}

func (lc *cache) Set(key string, value []string) {
	lc.mu.Lock()
	defer lc.mu.Unlock()

	lc.cache[key] = value
}

func (lc *cache) Get(key string) ([]string, bool) {
	lc.mu.RLock()
	defer lc.mu.RUnlock()

	value, ok := lc.cache[key]
	return value, ok
}

var instance *cache = NewCache()

func Set(key string, value []string) {
	instance.Set(key, value)
}

func Get(key string) ([]string, bool) {
	return instance.Get(key)
}
