package store

import "sync"

type Store struct {
	mu   sync.RWMutex
	data map[string]string // will begin with simple string as a value, but later i'll use generic []bytes and serialization/de-serialization.
}

func NewMemoryStore() *Store {
	return &Store{
		data: make(map[string]string),
	}
}

func (s *Store) Put(key, value string) error {
	s.mu.Lock()
	s.data[key] = value
	s.mu.Unlock()

	return nil
}

func (s *Store) Get(key string) (string, error) {
	return s.data[key], nil
}

func (s *Store) Remove(key string) error {
	s.mu.Lock()
	delete(s.data, key)
	s.mu.Unlock()

	return nil
}
