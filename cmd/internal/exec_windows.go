package internal

import (
	"errors"
	"os"
	"os/exec"
)

// Pseudo-replaces current process. Use like `Exec(os.Args[0], os.Args[1:]...)`.
func Exec(exePath string, args ...string) error {
	err := exec.Command(exePath, args...).Run()
	if err != nil {
		var exitError *exec.ExitError
		if errors.As(err, &exitError) {
			os.Exit(exitError.ExitCode())
		} else {
			return err
		}
	} else {
		os.Exit(0)
		return nil
	}
}