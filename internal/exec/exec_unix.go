//go:build aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris

package exec

import (
	"os"

	"golang.org/x/sys/unix"
)

func Exec(cmd string, args ...string) error {
	argv := []string{cmd}
	argv = append(argv, args...)
	return unix.Exec(cmd, argv, os.Environ())
}
