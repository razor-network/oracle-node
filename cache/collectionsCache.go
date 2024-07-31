package cache

import (
	"razor/pkg/bindings"
	"sync"
)

// CollectionsCache struct to hold collection cache and associated mutex
type CollectionsCache struct {
	Collections map[uint16]bindings.StructsCollection
	Mu          sync.RWMutex
}

// NewCollectionsCache creates a new instance of CollectionsCache
func NewCollectionsCache() *CollectionsCache {
	return &CollectionsCache{
		Collections: make(map[uint16]bindings.StructsCollection),
		Mu:          sync.RWMutex{},
	}
}

func (c *CollectionsCache) GetCollection(collectionId uint16) (bindings.StructsCollection, bool) {
	c.Mu.RLock()
	defer c.Mu.RUnlock()

	collection, exists := c.Collections[collectionId]
	return collection, exists
}

func (c *CollectionsCache) UpdateCollection(collectionId uint16, updatedCollection bindings.StructsCollection) {
	c.Mu.Lock()
	defer c.Mu.Unlock()

	c.Collections[collectionId] = updatedCollection
}
