package fs

import (
	"io/fs"
	"path/filepath"

	"github.com/fwojciec/packager"
)

type dirListFunc func(target string) ([]string, error)

func (f dirListFunc) ListDir(target string) ([]string, error) {
	return f(target)
}

func NewDirLister() packager.DirLister {
	return dirListFunc(func(target string) ([]string, error) {
		res := make([]string, 0)
		err := filepath.WalkDir(target, func(path string, d fs.DirEntry, err error) error {
			if d.IsDir() {
				return nil
			}
			if err != nil {
				return err
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
