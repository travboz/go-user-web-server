package main

import (
	"errors"
	"sync"
)

type SafeCache struct {
	userCache  map[int]User
	cacheMutex sync.RWMutex
}

func NewSafeCache() *SafeCache {
	return &SafeCache{
		userCache: make(map[int]User),
	}
}

func (s *SafeCache) Insert(u User) User {
	s.cacheMutex.Lock()
	s.userCache[len(s.userCache)+1] = u
	s.cacheMutex.Unlock()

	return u
}

func (s *SafeCache) Get(id int) (User, error) {
	s.cacheMutex.RLock()
	user, exists := s.userCache[id]
	s.cacheMutex.RUnlock()

	if !exists {
		return User{}, errors.New("user doesn't exist")
	}

	return user, nil
}

func (s *SafeCache) Delete(id int) error {
	s.cacheMutex.Lock()
	_, ok := s.userCache[id]

	if ok {
		delete(s.userCache, id)
	}

	s.cacheMutex.Unlock()

	if !ok {
		return errors.New("user doesn't exist")
	}

	return nil
}
