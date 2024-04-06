package internal

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	_ "embed"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"

	"golang.org/x/sys/unix"
)

//go:embed swig-4.2.1.tar.gz
var swigTarGz []byte



// 1. Extract the tarball to "./.go-swig" relative to this executable. The tarball contains a single "swig-4.2.1" directory with the SWIG code in it.
// 2. Run "./configure --prefix $(realpath ..) && make && make install" in the extracted directory. This will install the SWIG project to the "./.go-swig" directory relative to this executable. Inside that directory, there should be a "bin" folder with the "swig" and "ccache-swig" executables.
// 3. Remove the extracted directory.
// 4. Move this executable to an adjacent "${name}.old" file.
// 5. Create new symlinks for "swig" and "ccache-swig" that point to "./.go-swig/bin/swig" and "./.go-swig/bin/ccache-swig" relative to this executable respectively.
// 6. Delete (if possible) this executable (which now has a ".old" suffix).
// 7. Replace this process with a re-execution of the original executable path (which is now a symlink to "./.go-swig/bin/swig" or "./.go-swig/bin/ccache-swig") with the original arguments and environment.
func Main() {
	originalExecutablePath, err := os.Executable()
	if err != nil {
		panic(err)
	}
	goSwigPath := filepath.Join(filepath.Dir(originalExecutablePath), ".go-swig")
	err = extractTarGzBytes(swigTarGz, goSwigPath)
	if err != nil {
		panic(err)
	}
	swigSrcPath := filepath.Join(goSwigPath, "swig-4.2.1")
	cmd := exec.Command(filepath.Join(swigSrcPath, "configure"), "--prefix", goSwigPath)
	cmd.Dir = swigSrcPath
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stderr
	cmd.Stderr = os.Stderr
	fmt.Println(cmd.String())
	err = cmd.Run()
	if err != nil {
		panic(err)
	}
	cmd = exec.Command("make")
	cmd.Dir = swigSrcPath
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stderr
	cmd.Stderr = os.Stderr
	fmt.Println(cmd.String())
	err = cmd.Run()
	if err != nil {
		panic(err)
	}
	cmd = exec.Command("make", "install")
	cmd.Dir = swigSrcPath
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stderr
	cmd.Stderr = os.Stderr
	fmt.Println(cmd.String())
	err = cmd.Run()
	if err != nil {
		panic(err)
	}
	err = os.RemoveAll(swigSrcPath)
	if err != nil {
		panic(err)
	}
	oldExecutablePath := originalExecutablePath + ".old"
	err = os.Rename(originalExecutablePath, oldExecutablePath)
	if err != nil {
		panic(err)
	}
	swigPath := filepath.Join(goSwigPath, "bin", "swig")
	err = os.Symlink(swigPath, filepath.Join(filepath.Dir(originalExecutablePath), "swig"))
	if err != nil {
		panic(err)
	}
	_ = os.Remove(oldExecutablePath)
	err = unix.Exec(originalExecutablePath, os.Args, os.Environ())
	if err != nil {
		panic(err)
	}
}
