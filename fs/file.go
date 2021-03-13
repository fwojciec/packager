package fs

import (
	"io"
	"os"

	"github.com/fwojciec/packager"
)

type readFileFunc func(path string) ([]byte, error)

func (f readFileFunc) ReadFile(path string) ([]byte, error) {
	return f(path)
}

func NewFileReader() packager.FileReader {
	return readFileFunc(os.ReadFile)
}

type copyFileFunc func(srcPath string, dest io.Writer) error

func (f copyFileFunc) Copy(srcPath string, dest io.Writer) error {
	return f(srcPath, dest)
}

func NewFileCopier() packager.FileCopier {
	return copyFileFunc(func(srcPath string, dest io.Writer) error {
		f, err := os.Open(srcPath)
		if err != nil {
			return err
		}
		defer f.Close()
		_, err = io.Copy(dest, f)
		if err != nil {
			return err
		}
		return nil
	})
}
