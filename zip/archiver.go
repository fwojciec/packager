package zip

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"

	"github.com/fwojciec/packager"
)

type archiver struct {
	dirLister packager.DirLister
}

func (a *archiver) Archive(project packager.Locator, dest string) error {
	root := project.Location()
	projFiles, err := a.dirLister.ListDir(root)
	if err != nil {
		return err
	}

	absDest, err := filepath.Abs(dest)
	if err != nil {
		return err
	}

	f, err := os.Create(absDest)
	defer f.Close()

	w := zip.NewWriter(f)
	for _, file := range projFiles {
		relPath, err := filepath.Rel(root, file)
		if err != nil {
			return err
		}
		zf, err := w.Create(relPath)
		if err != nil {
			return err
		}
		sf, err := os.Open(file)
		if err != nil {
			return err
		}
		defer sf.Close()

		_, err = io.Copy(zf, sf)
		if err != nil {
			return err
		}
	}

	return w.Close()
}

func New(dirLister packager.DirLister) packager.Archiver {
	return &archiver{
		dirLister: dirLister,
	}
}
