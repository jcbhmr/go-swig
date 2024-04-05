package main

import (
	"bytes"
	"os"
	"os/exec"

	"github.com/inconshreveable/go-update"
)

func main() {
	executable, err := os.Executable()
	if err != nil {
		panic(err)
	}
	reader := bytes.NewReader(swig)
	err = update.Apply(reader, update.Options{})
	if err != nil {
		panic(err)
	}
	cmd := exec.Command(executable, os.Args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err = cmd.Run()
	if exitError, ok := err.(*exec.ExitError); ok {
		os.Exit(exitError.ExitCode())
	}
}
