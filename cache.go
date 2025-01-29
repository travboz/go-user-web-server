package main

import (
	"sync"
)

type SafeCache struct {
	userCache  map[int]User
	cacheMutex sync.RWMutex
	nextID     int
}

func NewSafeCache() *SafeCache {
	return &SafeCache{
		userCache: make(map[int]User),
	}
}

func (s *SafeCache) Insert(u User) User {
	s.cacheMutex.Lock()
	defer s.cacheMutex.Unlock()

	s.nextID++
	s.userCache[s.nextID] = u

	return u
}

func (s *SafeCache) Get(id int) (User, error) {
	s.cacheMutex.RLock()
	user, exists := s.userCache[id]
	s.cacheMutex.RUnlock()

	if !exists {
		return User{}, ErrUserDoesNotExist
	}

	return user, nil
}

func (s *SafeCache) GetAll() []User {
	s.cacheMutex.RLock()
	defer s.cacheMutex.RUnlock()

	users := make([]User, 0, len(s.userCache))
	for _, user := range s.userCache {
		users = append(users, user)
	}

	return users
}

func (s *SafeCache) Delete(id int) error {
	s.cacheMutex.Lock()
	defer s.cacheMutex.Unlock()

	if _, exists := s.userCache[id]; !exists {
		return ErrUserDoesNotExist
	}

	delete(s.userCache, id)
	return nil
}
