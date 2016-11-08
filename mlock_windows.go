// +build windows

package main

import (
	"errors"
)

func Mlock() error {
	return errors.New("mlock not supported")
}
