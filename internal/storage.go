package internal

import (
	"sync"
	"sync/atomic"
)

type Storage struct {
	data sync.Map
}

func NewStorage() *Storage {
	return &Storage{}
}

func (s *Storage) Store(key string) {
	value, _ := s.data.LoadOrStore(key, &atomic.Int64{})

	counter, ok := value.(*atomic.Int64)
	if ok {
		counter.Add(1)
	}
}

func (s *Storage) Dump() map[string]int64 {
	result := make(map[string]int64)

	s.data.Range(func(k, v any) bool {
		//nolint:forcetypeassert
		result[k.(string)] = v.(*atomic.Int64).Load()

		return true
	})

	return result
}
