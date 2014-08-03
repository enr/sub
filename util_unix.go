// +build linux darwin

package main

import (
	"syscall"
	"unsafe"
)

// Copied from code.google.com/p/go.crypto/ssh/terminal.
func isatty(fd uintptr) bool {
	termios := syscall.Termios{}
	_, _, err := syscall.Syscall6(syscall.SYS_IOCTL, fd, ioctlReadTermios, uintptr(unsafe.Pointer(&termios)), 0, 0, 0)
	return err == 0
}
