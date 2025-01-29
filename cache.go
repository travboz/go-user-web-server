package main

import (
	"sync"
)

type SafeCache struct {
	userCache  map[int]User
	cacheMutex sync.RWMutex
	nextID     int
}

// NewSafeCache - init and return a new SafeCache instance.
func NewSafeCache() *SafeCache {
	return &SafeCache{
		userCache: make(map[int]User),
	}
}

// Insert - adds a new user to the cache with a unique ID.
// It returns the inserted user.
func (s *SafeCache) Insert(u User) User {
	s.cacheMutex.Lock()
	defer s.cacheMutex.Unlock()

	s.nextID++
	s.userCache[s.nextID] = u

	return u
}

// Get - retrieves a user by ID from the cache.
// It returns an error if the user does not exist.
func (s *SafeCache) Get(id int) (User, error) {
	s.cacheMutex.RLock()
	user, exists := s.userCache[id]
	s.cacheMutex.RUnlock()

	if !exists {
		return User{}, ErrUserDoesNotExist
	}

	return user, nil
}

// GetAll - retrieves all users from the cache.
// It returns a slice of User objects.
func (s *SafeCache) GetAll() []User {
	s.cacheMutex.RLock()
	defer s.cacheMutex.RUnlock()

	users := make([]User, 0, len(s.userCache))
	for _, user := range s.userCache {
		users = append(users, user)
	}

	return users
}

// Delete - removes a user from the cache by ID.
// It returns an error if the user does not exist.
func (s *SafeCache) Delete(id int) error {
	s.cacheMutex.Lock()
	defer s.cacheMutex.Unlock()

	if _, exists := s.userCache[id]; !exists {
		return ErrUserDoesNotExist
	}

	delete(s.userCache, id)
	return nil
}
