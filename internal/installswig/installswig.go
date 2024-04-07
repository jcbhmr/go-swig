package installswig

import (
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/jcbhmr/go-swig/internal/swigtargz"
)

const SwigVersion = swigtargz.Version

func InstallSwigToTheCacheDir() (err error) {
	userCacheDir, err := os.UserCacheDir()
	if err != nil {
		return fmt.Errorf("failed to get user cache dir: %w", err)
	}
	myCacheDir := filepath.Join(userCacheDir, "go-swig", SwigVersion)
	defer func(){
		if err != nil {
			_ = os.RemoveAll(myCacheDir)
		}
	}()

	if _, err := os.Stat(myCacheDir); err != nil {
		if !errors.Is(err, fs.ErrNotExist) {
			return fmt.Errorf("failed to stat %s: %w", myCacheDir, err)
		}

		if err := os.MkdirAll(myCacheDir, 0755); err != nil {
			return fmt.Errorf("failed to create %s: %w", myCacheDir, err)
		}

		sourceDir := filepath.Join(myCacheDir, "swig-"+SwigVersion)
		if _, err := os.Stat(sourceDir); err != nil {
			if !errors.Is(err, fs.ErrNotExist) {
				return fmt.Errorf("failed to stat %s: %w", sourceDir, err)
			}
			if err := swigtargz.ExtractTo(myCacheDir); err != nil {
				return fmt.Errorf("failed to extract swig-%s to %s: %w", SwigVersion, myCacheDir, err)
			}
			if _, err := os.Stat(sourceDir); err != nil {
				return fmt.Errorf("failed to stat %s: %w", sourceDir, err)
			}
		}

		var stderr strings.Builder
		cmd := exec.Command(filepath.Join(sourceDir, "configure"), "--prefix", myCacheDir)
		cmd.Dir = sourceDir
		cmd.Stderr = &stderr
		s := spinner.New(spinner.CharSets[21], 100*time.Millisecond)
		s.Suffix = "  " + cmd.String()
		s.Restart()
		err = cmd.Run()
		s.Stop()
		if err != nil {
			log.Print(stderr.String())
			return fmt.Errorf("%s failed: %w", cmd.String(), err)
		}

		stderr = strings.Builder{}
		cmd = exec.Command("make")
		cmd.Dir = sourceDir
		cmd.Stderr = &stderr
		s = spinner.New(spinner.CharSets[21], 100*time.Millisecond)
		s.Suffix = "  " + cmd.String()
		s.Restart()
		err = cmd.Run()
		s.Stop()
		if err != nil {
			log.Print(stderr.String())
			return fmt.Errorf("%s failed: %w", cmd.String(), err)
		}

		stderr = strings.Builder{}
		cmd = exec.Command("make", "install")
		cmd.Dir = sourceDir
		cmd.Stderr = &stderr
		s = spinner.New(spinner.CharSets[21], 100*time.Millisecond)
		s.Suffix = "  " + cmd.String()
		s.Restart()
		err = cmd.Run()
		s.Stop()
		if err != nil {
			log.Print(stderr.String())
			return fmt.Errorf("%s failed: %w", cmd.String(), err)
		}

		log.Printf("Installed swig-%s to %s", SwigVersion, myCacheDir)
	} else {
		log.Printf("swig-%s is already installed in %s", SwigVersion, myCacheDir)
	}
	return nil
}
