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

func UpdateJobCache(jobId uint16, updatedJobData bindings.StructsJob) {
	JobsCache.Mu.Lock()
	defer JobsCache.Mu.Unlock()

	// Check if the job already exists in the cache
	existingJob, exists := JobsCache.Jobs[jobId]
	if exists {
		// If the job exists, keep the existing name and update other fields
		existingJob.SelectorType = updatedJobData.SelectorType
		existingJob.Weight = updatedJobData.Weight
		existingJob.Power = updatedJobData.Power
		existingJob.Selector = updatedJobData.Selector
		existingJob.Url = updatedJobData.Url

		// Update the cache with the updated job data
		JobsCache.Jobs[jobId] = existingJob
	}
}
