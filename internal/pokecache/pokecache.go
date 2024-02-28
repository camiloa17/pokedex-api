package pokecache

import "time"

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for {
		currTime := <-ticker.C
		keysToDelete := []string{}

		for key, cache := range c.cacheEntry {
			if currTime.After(cache.createdAt) {
				keysToDelete = append(keysToDelete, key)
			}
		}
		c.deleteCache(keysToDelete)

	}
}

func (c *Cache) deleteCache(keysToDelete []string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	for _, key := range keysToDelete {
		delete(c.cacheEntry, key)
	}
}

// Get an entry from the cache
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	val, ok := c.cacheEntry[key]
	if !ok {
		return nil, ok
	}
	return val.val, ok
}

// Add entry to the cache
func (c *Cache) Add(key string, val []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.cacheEntry[key] = cacheEntry{
		val:       val,
		createdAt: time.Now(),
	}

}
