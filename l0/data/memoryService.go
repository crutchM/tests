package data

import "tests/l0/data/cache"

type MemoryService struct {
	*cache.Cache
	Repository
}

func NewMemoryService(repo Repository) *MemoryService {
	return &MemoryService{cache.NewCache(0, 0), repo}
}

func (s *MemoryService) getAllItems() {
	if len(s.GetAll()) == 0 {

	}
}
