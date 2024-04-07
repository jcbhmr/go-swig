package internal

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

// 1. Extract the tarball to "./.go-swig" relative to this executable. The tarball contains a single "swig-4.2.1" directory with the SWIG code in it.
// 2. Run "./configure --prefix $(realpath ..) && make && make install" in the extracted directory. This will install the SWIG project to the "./.go-swig" directory relative to this executable. Inside that directory, there should be a "bin" folder with the "swig" and "ccache-swig" executables.
// 3. Remove the extracted directory.
// 4. Move this executable to an adjacent "${name}.old" file.
// 5. Create new symlinks for "swig" and "ccache-swig" that point to "./.go-swig/bin/swig" and "./.go-swig/bin/ccache-swig" relative to this executable respectively.
// 6. Delete (if possible) this executable (which now has a ".old" suffix).
// 7. Replace this process with a re-execution of the original executable path (which is now a symlink to "./.go-swig/bin/swig" or "./.go-swig/bin/ccache-swig") with the original arguments and environment.
func Run(exeName string) {
	originalExePath, err := os.Executable()
	if err != nil {
		panic(err)
	}
	goSwigPath := filepath.Join(filepath.Dir(originalExePath), ".go-swig")
	_, err = os.Stat(goSwigPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			err = InstallSwig(goSwigPath, nil)
			if err != nil {
				panic(err)
			}
		} else {
			panic(err)
		}
	}
	oldExePath := originalExePath + ".old"
	err = os.Rename(originalExePath, oldExePath)
	if err != nil {
		panic(err)
	}
	var swigPath string
	if runtime.GOOS == "windows" {
		swigPath = filepath.Join(goSwigPath, "swig.exe")
	} else {
		swigPath = filepath.Join(goSwigPath, "bin", "swig")
	}
	err = os.Symlink(swigPath, originalExePath)
	if err != nil {
		_ = os.Rename(oldExePath, originalExePath)
		panic(err)
	}
	_ = os.Remove(oldExePath)
	err = Exec(originalExePath, os.Args[1:]...)
	if err != nil {
		panic(err)
	}
}
