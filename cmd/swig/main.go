package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	iexec "github.com/jcbhmr/go-swig/internal/exec"
	"github.com/jcbhmr/go-swig/internal/installswig"
)

func main() {
	log.Default().SetFlags(0)
	err := installswig.InstallSwigToTheCacheDir()
	if err != nil {
		log.Fatal(fmt.Errorf("failed to install swig to cache dir: %w", err))
	}
	userCacheDir, err := os.UserCacheDir()
	if err != nil {
		log.Fatal(fmt.Errorf("no cache dir in main: %w", err))
	}
	myCacheDir := filepath.Join(userCacheDir, "go-swig", installswig.SwigVersion)
	var swigExePath string
	if runtime.GOOS == "windows" {
		swigExePath = filepath.Join(myCacheDir, "swig.exe")
	} else {
		swigExePath = filepath.Join(myCacheDir, "bin", "swig")
	}
	originalExePath, err := os.Executable()
	if err != nil {
		log.Fatal(fmt.Errorf("failed to get executable path: %w", err))
	}
	oldExePath := originalExePath + ".old"
	err = os.Rename(originalExePath, oldExePath)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to rename %s to %s: %w", originalExePath, oldExePath, err))
	}
	err = os.Symlink(swigExePath, originalExePath)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to symlink %s to %s: %w", swigExePath, originalExePath, err))
	}
	_ = os.Remove(oldExePath)
	err = iexec.Exec(swigExePath, os.Args[1:]...)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to exec swig: %w", err))
	}
}
