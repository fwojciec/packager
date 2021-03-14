package fs

import (
	"io/fs"
	"path/filepath"

	"github.com/fwojciec/packager"
)

type dirListFunc func(target string, excl packager.Excluder) ([]string, error)

func (f dirListFunc) ListDir(target string, excl packager.Excluder) ([]string, error) {
	return f(target, excl)
}

func NewDirLister() packager.DirLister {
	return dirListFunc(func(target string, excl packager.Excluder) ([]string, error) {
		res := make([]string, 0)
		err := filepath.WalkDir(target, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if d.IsDir() {
				return nil
			}
			if excl.Exclude(path) {
				return nil
			}
			res = append(res, path)
			return nil
		})
		if err != nil {
			return nil, err
		}
		return res, nil
	})
}
