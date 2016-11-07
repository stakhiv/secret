package main

import (
	"crypto/rsa"
	"encoding/base64"
	"errors"
	"flag"
	"fmt"
	"github.com/howeyc/gopass"
	"os"
	"os/user"
	"path/filepath"
)

var (
	errMissingValue  = errors.New("missing value")
	errInvalidAction = errors.New("invalid action")
)

var (
	escape = flag.Bool("e", false, "Escape response")
)

func usage() {
	fmt.Println(`secret <store|get> name [value]
		-e - Denotes if data is binary and needs escaping
	`)
	os.Exit(0)
}

const workDir = ".secret"

func getKey(path string) (*rsa.PrivateKey, error) {
	var key *rsa.PrivateKey
	keyfile, err := os.Open(path)
	if err != nil {
		if !os.IsNotExist(err) {
			exit(err)
		}
		fmt.Print("Generating new key.\nEnter password: ")
		passwd, err := gopass.GetPasswdMasked()
		if err != nil {
			return nil, err
		}

		key, err = generateKey(path, passwd)
		if err != nil {
			exit(err)
		}
	} else {
		fmt.Print("Reading existing key.\nEnter password: ")
		passwd, err := gopass.GetPasswdMasked()
		if err != nil {
			return nil, err
		}
		defer keyfile.Close()
		key, err = readKey(keyfile, passwd)
		if err != nil {
			return nil, err
		}
	}
	return key, nil
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 2 {
		usage()
	}

	usr, err := user.Current()
	if err != nil {
		exit(err)
	}
	workDir := filepath.Join(usr.HomeDir, workDir)
	err = os.MkdirAll(workDir, 0700)
	if err != nil {
		exit(err)
	}
	keyPath := filepath.Join(workDir, ".key")
	key, err := getKey(keyPath)
	if err != nil {
		exit(err)
	}

	storagePath := filepath.Join(workDir, ".storage")
	storage := NewBoltStorage()
	err = storage.Open(storagePath)
	if err != nil {
		exit(err)
	}
	defer storage.Close()

	coder := &BaseEncoder{key}
	secret := NewSecret(coder, storage)

	switch args[0] {
	case "store":
		if len(args) < 3 {
			exit(errMissingValue)
		}
		// TODO: Add stdin reader
		err := secret.Store(args[1], []byte(args[2]))
		if err != nil {
			exit(err)
		}
		fmt.Println(args[1])
	case "get":
		b, err := secret.Get(args[1])
		if err != nil {
			exit(err)
		}
		if *escape {
			res := make([]byte, base64.StdEncoding.EncodedLen(len(b)))
			base64.StdEncoding.Encode(res, b)
			b = res
		}
		fmt.Println(string(b))
	default:
		exit(errInvalidAction)
	}
}

func exit(err error) {
	fmt.Println(err)
	os.Exit(1)
}
