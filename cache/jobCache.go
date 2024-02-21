package cache

import (
	"razor/pkg/bindings"
	"sync"
)

// JobsCacheStruct struct to hold job cache and associated mutex
type JobsCacheStruct struct {
	Jobs map[uint16]bindings.StructsJob
	Mu   sync.RWMutex
}

// JobsCache Global instance of JobsCacheStruct directly initialized
var JobsCache = JobsCacheStruct{
	Jobs: make(map[uint16]bindings.StructsJob),
	Mu:   sync.RWMutex{},
}

func GetJobFromCache(jobId uint16) (bindings.StructsJob, bool) {
	JobsCache.Mu.RLock() // Use read lock for concurrency safety
	defer JobsCache.Mu.RUnlock()

	job, exists := JobsCache.Jobs[jobId]
	return job, exists
}

func UpdateJobCache(jobId uint16, updatedJob bindings.StructsJob) {
	JobsCache.Mu.Lock()
	defer JobsCache.Mu.Unlock()

	// Update or add the job in the cache with the new data
	JobsCache.Jobs[jobId] = updatedJob
}
