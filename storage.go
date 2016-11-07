package main

import (
	"errors"
)

var (
	errNotFound = errors.New("not found")
)

type Storage interface {
	Set(key []byte, data []byte) error
	Get(key []byte) ([]byte, error)
}

type MemStorage struct {
	store map[string][]byte
}

func NewMemStorage() *MemStorage {
	return &MemStorage{
		store: make(map[string][]byte),
	}
}

func (s *MemStorage) Set(key []byte, data []byte) error {
	s.store[string(key)] = data
	return nil
}

func (s *MemStorage) Get(key []byte) ([]byte, error) {
	d, ok := s.store[string(key)]
	if !ok {
		return nil, errNotFound
	}
	return d, nil
}
