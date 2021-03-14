package fs

import (
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
