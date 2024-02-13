package cache

import (
	"razor/pkg/bindings"
	"sync"
)

// CollectionsCacheStruct struct to hold collection cache and associated mutex
type CollectionsCacheStruct struct {
	Collections map[uint16]bindings.StructsCollection
	Mu          sync.RWMutex
}

// CollectionsCache Global instances of CollectionsCacheStruct directly initialized
var CollectionsCache = CollectionsCacheStruct{
	Collections: make(map[uint16]bindings.StructsCollection),
	Mu:          sync.RWMutex{},
}

func GetCollectionFromCache(collectionId uint16) (bindings.StructsCollection, bool) {
	CollectionsCache.Mu.RLock() // Use read lock for concurrency safety
	defer CollectionsCache.Mu.RUnlock()

	collection, exists := CollectionsCache.Collections[collectionId]
	return collection, exists
}

func UpdateCollectionCache(collectionId uint16, updatedCollection bindings.StructsCollection) {
	CollectionsCache.Mu.Lock()
	defer CollectionsCache.Mu.Unlock()

	// Update or add the collection in the cache with the new data
	CollectionsCache.Collections[collectionId] = updatedCollection
}
