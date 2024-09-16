package cache

import (
	"razor/pkg/bindings"
	"sync"
)

// JobsCache struct to hold job cache and associated mutex
type JobsCache struct {
	Jobs map[uint16]bindings.StructsJob
	Mu   sync.RWMutex
}

// NewJobsCache creates a new instance of JobsCache
func NewJobsCache() *JobsCache {
	return &JobsCache{
		Jobs: make(map[uint16]bindings.StructsJob),
		Mu:   sync.RWMutex{},
	}
}

func (j *JobsCache) GetJob(jobId uint16) (bindings.StructsJob, bool) {
	j.Mu.RLock()
	defer j.Mu.RUnlock()

	job, exists := j.Jobs[jobId]
	return job, exists
}

func (j *JobsCache) UpdateJob(jobId uint16, updatedJob bindings.StructsJob) {
	j.Mu.Lock()
	defer j.Mu.Unlock()

	j.Jobs[jobId] = updatedJob
}
