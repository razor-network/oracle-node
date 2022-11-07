package cache

import (
	"errors"
	"sync"
	"time"
)

var (
	errUserNotInCache = errors.New("the user isn't in cache")
)

type Data struct {
	Result []byte
}

type cachedData struct {
	Data
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

func (lc *LocalCache) Update(u Data, url string, expireAtTimestamp int64) {
	lc.mu.Lock()
	defer lc.mu.Unlock()

	lc.URLs[url] = cachedData{
		Data:              u,
		expireAtTimestamp: expireAtTimestamp,
	}
}

func (lc *LocalCache) Read(url string) (Data, error) {
	lc.mu.RLock()
	defer lc.mu.RUnlock()

	cu, ok := lc.URLs[url]
	if !ok {
		return Data{}, errUserNotInCache
	}

	return cu.Data, nil
}

func (lc *LocalCache) Delete(url string) {
	lc.mu.Lock()
	defer lc.mu.Unlock()

	delete(lc.URLs, url)
}
