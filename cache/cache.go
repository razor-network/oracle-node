package cache

import (
	"errors"
	"sync"
	"time"
)

var (
	errDataNotInCache = errors.New("data not present in cache")
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

func (lc *LocalCache) Read(url string) ([]byte, error) {
	lc.mu.RLock()
	defer lc.mu.RUnlock()

	cacheData, ok := lc.URLs[url]
	if !ok {
		return []byte{}, errDataNotInCache
	}

	return cacheData.Result, nil
}

func (lc *LocalCache) Delete(url string) {
	lc.mu.Lock()
	defer lc.mu.Unlock()

	delete(lc.URLs, url)
}
