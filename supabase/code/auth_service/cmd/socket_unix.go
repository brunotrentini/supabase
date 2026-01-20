//go:build !windows

package cmd

import "golang.org/x/sys/unix"

func setReusePort(fd uintptr) error {
	return unix.SetsockoptInt(int(fd), unix.SOL_SOCKET, unix.SO_REUSEPORT, 1)
}
