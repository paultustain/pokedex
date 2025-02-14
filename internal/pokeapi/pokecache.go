package pokeapi

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	entries  map[string]cacheEntry
	mu       sync.Mutex
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {

	cache := Cache{
		entries:  make(map[string]cacheEntry, 1),
		interval: interval,
	}

	go cache.reapLoop(interval)

	return cache

}

func (c *Cache) reapLoop(interval time.Duration) {
	for {
		time.Sleep(interval)
		c.mu.Lock()
		for key, entry := range c.entries {
			if time.Since(entry.createdAt) > interval {

				delete(c.entries, key)

			}
		}
		c.mu.Unlock()
	}

}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	new_entry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}

	c.entries[key] = new_entry
	fmt.Println(c.entries)

}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	val, ok := c.entries[key]
	if !ok {
		return nil, false
	}
	return val.val, true
}
