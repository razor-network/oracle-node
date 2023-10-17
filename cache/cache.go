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
	stop chan struct{}

	wg   sync.WaitGroup
	mu   sync.RWMutex
	URLs map[string]cachedData //URLs
}

func NewLocalCache(cleanupInterval time.Duration) *LocalCache {
	lc := &LocalCache{
		URLs: make(map[string]cachedData),
		stop: make(chan struct{}),
	}

	lc.wg.Add(1)
	go func(cleanupInterval time.Duration) {
		defer lc.wg.Done()
		lc.cleanupLoop(cleanupInterval)
	}(cleanupInterval)

	return lc
}

func (lc *LocalCache) cleanupLoop(interval time.Duration) {
	t := time.NewTicker(interval)
	defer t.Stop()

	for {
		select {
		case <-lc.stop:
			return
		case <-t.C:
			lc.mu.Lock()
			for url, cu := range lc.URLs {
				if cu.expireAtTimestamp <= time.Now().Unix() {
					delete(lc.URLs, url)
				}
			}
			lc.mu.Unlock()
		}
	}
}

func (lc *LocalCache) StopCleanup() {
	close(lc.stop)
	lc.wg.Wait()
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

func (lc *LocalCache) Delete(url string) {
	lc.mu.Lock()
	defer lc.mu.Unlock()

	delete(lc.URLs, url)
}
