package zip

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"

	"github.com/fwojciec/packager"
)

type archiver struct {
	dl packager.DirLister
}

func (a *archiver) Archive(project packager.Locator, dest string) error {
	src := project.Location()

	noopExclFn := func(path string) (bool, error) { return false, nil }
	projFiles, err := a.dl.ListDir(src, noopExclFn)
	if err != nil {
		return err
	}

	f, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer f.Close()

	w := zip.NewWriter(f)
	for _, file := range projFiles {
		relPath, err := filepath.Rel(src, file)
		if err != nil {
			return err
		}
		zf, err := w.Create(relPath)
		if err != nil {
			return err
		}
		if err := copyFile(file, zf); err != nil {
			return err
		}
	}

	if err := w.Close(); err != nil {
		return err
	}

	return nil
}

func copyFile(srcPath string, dest io.Writer) error {
	src, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer src.Close()
	if _, err := io.Copy(dest, src); err != nil {
		return err
	}
	return nil
}

func New(dl packager.DirLister) packager.Archiver {
	return &archiver{dl}
}
