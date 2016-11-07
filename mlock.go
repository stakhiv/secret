// +build darwin linux freebsd

package main

import (
	"syscall"
)

func Mlock() error {
	return syscall.Mlockall(syscall.MCL_CURRENT | syscall.MCL_FUTURE)
}
