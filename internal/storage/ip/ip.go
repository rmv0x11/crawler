package ip

import "sync"

type storageByIP struct {
	mu *sync.RWMutex
	m  map[string]any
}

var (
	Storage = &storageByIP{
		mu: &sync.RWMutex{},
		m:  map[string]any{},
	}
)

func (s *storageByIP) LookupKey(k string) (any, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	val, ok := s.m[k]

	return val, ok
}

func (s *storageByIP) Set(k string, v any) {
	s.mu.Lock()
	s.m[k] = v
	s.mu.Unlock()
}

func (s *storageByIP) Delete(k string) {
	s.mu.Lock()
	delete(s.m, k)
	s.mu.Unlock()
}
