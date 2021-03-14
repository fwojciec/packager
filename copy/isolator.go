package copy

import (
	"os"

	"github.com/fwojciec/packager"
	"github.com/otiai10/copy"
)

type locatorRemover struct {
	root string
}

func (lr *locatorRemover) Location() string { return lr.root }
func (lr *locatorRemover) Remove() error    { return os.RemoveAll(lr.root) }

type isolatorFunc func(packager.LocatorExcluder) (packager.LocatorRemover, error)

func (f isolatorFunc) Isolate(project packager.LocatorExcluder) (packager.LocatorRemover, error) {
	return f(project)
}

func isolate(project packager.LocatorExcluder) (packager.LocatorRemover, error) {
	tmpDir, err := os.MkdirTemp("", "packager_isolator")
	if err != nil {
		return nil, err
	}
	options := copy.Options{
		Skip: func(src string) (bool, error) {
			return project.Exclude(src)
		},
	}
	if err := copy.Copy(project.Location(), tmpDir, options); err != nil {
		return nil, err
	}
	return &locatorRemover{root: tmpDir}, nil
}

func NewIsolator() packager.Isolator {
	return isolatorFunc(isolate)
}
