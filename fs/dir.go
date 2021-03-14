package fs

import (
	"io/fs"
	"path/filepath"

	"github.com/fwojciec/packager"
)

type dirListFunc func(target string, exclFn func(path string) (bool, error)) ([]string, error)

func (f dirListFunc) ListDir(target string, exclFn func(path string) (bool, error)) ([]string, error) {
	return f(target, exclFn)
}

func NewDirLister() packager.DirLister {
	return dirListFunc(func(target string, exclFn func(path string) (bool, error)) ([]string, error) {
		res := make([]string, 0)
		err := filepath.WalkDir(target, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if d.IsDir() {
				return nil
			}
			excluded, err := exclFn(path)
			if err != nil {
				return err
			}
			if excluded {
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
