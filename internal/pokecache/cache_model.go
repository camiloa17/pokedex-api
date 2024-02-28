package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	cacheEntry map[string]cacheEntry
	interval   time.Duration
	mutex      sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		cacheEntry: make(map[string]cacheEntry),
		interval:   interval,
	}
	go c.reapLoop()

	return c
}
