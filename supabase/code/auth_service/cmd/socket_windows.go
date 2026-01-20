//go:build windows

package cmd

func setReusePort(fd uintptr) error {
	return nil
}
