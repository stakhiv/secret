// +build windows

package main

import (
	"errors"
)

func Mlock() {
	return errors.New("mlock not supported")
}
