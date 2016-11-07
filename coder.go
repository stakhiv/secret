package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
)

var label = []byte("password")

type Coder interface {
	Key([]byte) ([]byte, error)
	Encode([]byte) ([]byte, error)
	Decode([]byte) ([]byte, error)
}

type BaseEncoder struct {
	key *rsa.PrivateKey
}

func (e *BaseEncoder) Key(b []byte) ([]byte, error) {
	return sha256.New().Sum(b), nil
}

func (e *BaseEncoder) Encode(b []byte) ([]byte, error) {
	b, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, &e.key.PublicKey, b, label)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (e *BaseEncoder) Decode(b []byte) ([]byte, error) {
	b, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, e.key, b, label)
	if err != nil {
		return nil, err
	}
	return b, nil
}
