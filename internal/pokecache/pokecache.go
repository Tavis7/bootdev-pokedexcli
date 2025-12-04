package pokecache

import (
	"fmt"
	"time"
	"sync"
)

type cacheEntry struct {
	createdAt time.Time
	val []byte
}

type Cache struct {
	entries map[string]cacheEntry
	interval time.Duration
	mu sync.Mutex
}

func (c Cache) reap() {
	cutoff := time.Now().UTC().Add(-c.interval)

	c.mu.Lock()
	defer c.mu.Unlock()

	for key, val := range c.entries {
		if val.createdAt.Before(cutoff) {
			//fmt.Printf("Evicting %v\n", key)
			delete(c.entries, key)
		}
	}
}

func (c Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for _ = range ticker.C {
		c.reap()
	}

	fmt.Println("reapLoop exiting")
}

func NewCache(interval time.Duration) Cache {
	c := Cache {
		interval: interval,
		entries: make(map[string]cacheEntry),
	}
	go c.reapLoop(interval)
	return c
}

func (c Cache) Add(key string, entry []byte) {
	//fmt.Printf("Adding %v\n", key)
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entries[key] = cacheEntry{
		val: entry,
		createdAt: time.Now().UTC(),
	}
}

func (c Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	got,ok := c.entries[key]

	return got.val, ok
}

func Test() {
	fmt.Println("Hello from pokecache")
}
