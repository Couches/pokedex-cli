package pokecache

import (
	"time"
)

type Cache struct {
	contents map[string]cacheEntry
  interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
  cache := Cache {
    contents: map[string]cacheEntry{},
    interval: interval,
  }

  go cache.reapLoop()

  return cache
}

func (c Cache) Add(key string, val []byte) {
  c.contents[key] = cacheEntry{
    createdAt: time.Now(),
    val: val,
  }
}

func (c Cache) Get(key string) ([]byte, bool) {
  entry, ok := c.contents[key]
  return entry.val, ok
}

func (c Cache) reapLoop() {
  ticker := time.NewTicker(1 * time.Second)

  defer ticker.Stop()

  for {
    select {
    case <-ticker.C:
      c.cullExpired()
    }
  }
}

func (c Cache) cullExpired() {
  for key, value := range c.contents {
    if float64(time.Now().Second()) > float64(value.createdAt.Second()) + c.interval.Seconds() {
      delete(c.contents, key)
    }
  }
}
