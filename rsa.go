package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io/ioutil"
	"os"
)

var (
	errPasswordShort = errors.New("password too short")
)

func encryptRSAKey(key *rsa.PrivateKey, password []byte) ([]byte, error) {
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	}
	b := pem.EncodeToMemory(block)
	block, err := x509.EncryptPEMBlock(rand.Reader, "XXX", b, password, x509.PEMCipherAES256)
	if err != nil {
		return nil, err
	}
	b = pem.EncodeToMemory(block)

	return b, nil
}

func generateKey(path string, passwd []byte) (*rsa.PrivateKey, error) {
	if len(passwd) < 8 {
		return nil, errPasswordShort
	}
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, err
	}

	b, err := encryptRSAKey(key, passwd)
	if err != nil {
		return nil, err
	}
	f, err := os.OpenFile(path, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0600)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	_, err = f.Write(b)
	if err != nil {
		return nil, err
	}
	return key, nil
}

func readKey(file *os.File, passwd []byte) (*rsa.PrivateKey, error) {
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(data)

	b, err := x509.DecryptPEMBlock(block, passwd)
	if err != nil {
		return nil, err
	}
	block, _ = pem.Decode(b)
	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return key, nil
}
