package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	data     map[string]cacheEntry
	mu       sync.Mutex
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(duration time.Duration) *Cache {
	cache := Cache{data: map[string]cacheEntry{}, interval: duration}
	go cache.reapLoop()
	return &cache
}

func (c *Cache) Add(key string, val []byte) {
	fmt.Println("001")
	c.mu.Lock()
	fmt.Println("002")
	defer c.mu.Unlock()
	fmt.Println(c)

	c.data[key] = cacheEntry{val: val, createdAt: time.Now()}
	fmt.Println("004")
}
func (c *Cache) Get(key string) (cachedData []byte, hasEntry bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, ok := c.data[key]
	if ok {
		return entry.val, ok
	}
	return nil, ok
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()
	for {
		//fmt.Println("tic")
		t := <-ticker.C
		c.mu.Lock()
		for key, etn := range c.data {
			if etn.createdAt.Add(c.interval).Before(t) {
				delete(c.data, key)
				continue
			}

			break

		}
		c.mu.Unlock()
	}

}
