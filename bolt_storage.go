package main

import (
	"github.com/boltdb/bolt"
)

var defaultBucket = []byte("secrets")

type BoltStorage struct {
	db *bolt.DB
}

func NewBoltStorage() *BoltStorage {
	return &BoltStorage{}
}

func (s *BoltStorage) Open(path string) error {
	db, err := bolt.Open(path, 0600, nil)
	if err != nil {
		return err
	}
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(defaultBucket)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		db.Close()
		return err
	}
	s.db = db
	return nil
}

func (s *BoltStorage) Close() error {
	return s.db.Close()
}

func (s *BoltStorage) Get(key []byte) ([]byte, error) {
	var res []byte
	err := s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(defaultBucket)
		if b == nil {
			return errNotFound
		}
		res = b.Get(key)
		if res == nil {
			return errNotFound
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *BoltStorage) Set(key []byte, data []byte) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(defaultBucket)
		if b == nil {
			return errNotFound
		}
		return b.Put(key, data)
	})
}
