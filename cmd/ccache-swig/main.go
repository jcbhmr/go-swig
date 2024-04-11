package main

import (
	"errors"
	"fmt"
	"io/fs"
	"log"
	"log/slog"
	"os"
	"path/filepath"
	"runtime/debug"

	"github.com/jcbhmr/go-swig/internal/dotlock"
	"github.com/jcbhmr/go-swig/internal/swigtargz"
)

var version string
var cacheDir string

func init() {
	log.SetFlags(0)
	bi, ok := debug.ReadBuildInfo()
	if !ok {
		err := errors.New("not a go module")
		panic(fmt.Errorf("could not read build info: %w", err))
	}
	version = bi.Main.Version

	userCacheDir, err := os.UserCacheDir()
	if err != nil {
		userCacheDir = os.TempDir()
	}
	cacheDir = filepath.Join(userCacheDir, ".go-swig")
}

func main() {
	err := os.MkdirAll(cacheDir, 0700)
	if err != nil {
		log.Fatalln(err)
	}
	err = os.MkdirAll(filepath.Join(cacheDir, "source"), 0700)
	if err != nil {
		log.Fatalln(err)
	}

	// 1. Extract
	extracted := dotlock.NewMutex(filepath.Join(cacheDir, "extracted-"+version))
	if locked, _ := extracted.IsLocked(); !locked {
		extracting := dotlock.NewMutex(filepath.Join(cacheDir, "extracting-"+version))
		if ok, _ := extracting.TryLock(); ok {
			err := swigtargz.ExtractTo(filepath.Join(cacheDir, "source", "swig-"+version))
			if err != nil {
				extracting.Unlock()
				log.Fatalln(err)
			}
			extracting.Unlock()
		}
	}
	if ok, _ := extracted.TryLock(); ok {
		err := errors.New("failed to lock extracted")
		log.Fatalln(err)
	}

	// 2. Build
	built := dotlock.NewMutex(filepath.Join(cacheDir, "built-"+version))
	if locked, _ := built.IsLocked(); !locked {
		building := dotlock.NewMutex(filepath.Join(cacheDir, "building-"+version))
		if ok, _ := building.TryLock(); ok {
			err := build(filepath.Join(cacheDir, "source", "swig-"+version), filepath.Join(cacheDir, "swig-"+version))
			if err != nil {
				building.Unlock()
				log.Fatalln(err)
			}
			building.Unlock()
		}
	}
}
