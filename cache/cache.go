package cache

import (
	"sync"
	"time"
)

type cachedData struct {
	Result            []byte
	expireAtTimestamp int64
}

type LocalCache struct {
	mu   sync.RWMutex
	URLs map[string]cachedData //URLs
}

// NewLocalCache creates a new LocalCache instance
func NewLocalCache() *LocalCache {
	return &LocalCache{
		URLs: make(map[string]cachedData),
	}
}

func (lc *LocalCache) Update(data []byte, url string, expireAtTimestamp int64) {
	lc.mu.Lock()
	defer lc.mu.Unlock()

	lc.URLs[url] = cachedData{
		Result:            data,
		expireAtTimestamp: expireAtTimestamp,
	}
}

// Read fetches cached data for the given URL and indicates its presence with a boolean.
func (lc *LocalCache) Read(url string) ([]byte, bool) {
	lc.mu.RLock()
	defer lc.mu.RUnlock()

	cacheData, ok := lc.URLs[url]
	if !ok {
		return nil, false // Data not found, return nil and false
	}

	return cacheData.Result, true
}

// ClearAll deletes all entries in the cache
func (lc *LocalCache) ClearAll() {
	lc.mu.Lock()
	defer lc.mu.Unlock()

	for key := range lc.URLs {
		delete(lc.URLs, key)
	}
}

// Cleanup removes expired cache entries
func (lc *LocalCache) Cleanup() {
	lc.mu.Lock()
	defer lc.mu.Unlock()

	for url, data := range lc.URLs {
		// Remove expired data after the expireAtTimestamp is passed
		if data.expireAtTimestamp <= time.Now().Unix() {
			delete(lc.URLs, url)
		}
	}
}
