//go:build !windows

package internal

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/jcbhmr/go-swig/cmd/internal/swigtargz"
)

func InstallSwig(dest string, l *log.Logger) error {
	if l == nil {
		l = log.Default()
	}
	err := os.MkdirAll(dest, 0755)
	if err != nil {
		return err
	}
	source := filepath.Join(dest, "source")
	err = os.Mkdir(source, 0755)
	if err != nil {
		return err
	}
	err = ExtractTarGzBytesWithStripComponents(swigtargz.Get(), source, 1)
	if err != nil {
		return err
	}
	cmd := exec.Command(filepath.Join(source, "configure"), "--prefix", dest)
	cmd.Dir = source
	cmd.Stdout = l.Writer()
	cmd.Stderr = l.Writer()
	l.Printf("$ %s\n", cmd.String())
	err = cmd.Run()
	if err != nil {
		return err
	}
	cmd = exec.Command("make")
	cmd.Dir = source
	cmd.Stdout = l.Writer()
	cmd.Stderr = l.Writer()
	l.Printf("$ %s\n", cmd.String())
	err = cmd.Run()
	if err != nil {
		return err
	}
	cmd = exec.Command("make", "install")
	cmd.Dir = source
	cmd.Stdout = l.Writer()
	cmd.Stderr = l.Writer()
	l.Printf("$ %s\n", cmd.String())
	err = cmd.Run()
	if err != nil {
		return err
	}
	err = os.RemoveAll(source)
	if err != nil {
		return err
	}
	return nil
}