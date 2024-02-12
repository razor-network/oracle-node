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

func UpdateCollectionCache(collectionId uint16, collectionData bindings.StructsCollection) {
	CollectionsCache.Mu.Lock()
	defer CollectionsCache.Mu.Unlock()

	// Check if the collection already exists in the cache
	existingCollection, exists := CollectionsCache.Collections[collectionId]
	if exists {
		// If the collection exists, keep the existing name, active status and update other fields
		existingCollection.Power = collectionData.Power
		existingCollection.Tolerance = collectionData.Tolerance
		existingCollection.AggregationMethod = collectionData.AggregationMethod
		existingCollection.JobIDs = collectionData.JobIDs

		// Update the cache with the modified collection entry
		CollectionsCache.Collections[collectionId] = existingCollection
	}
}
