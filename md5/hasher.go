package md5

import "github.com/fwojciec/packager"

type hashFunc func(project packager.LocatorExcluder) (string, error)

func (f hashFunc) Hash(project packager.LocatorExcluder) (string, error) {
	return f(project)
}

func New() packager.Hasher {
	return hashFunc(func(project packager.LocatorExcluder) (string, error) {
		return "", nil
	})
}
