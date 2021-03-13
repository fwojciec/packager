package glob

import "github.com/fwojciec/packager"

type projectFactory struct {
	fr packager.FileReader
}

// New returns a new instance of project (LocatorExcluder)
func (pf *projectFactory) New(root string) (packager.LocatorExcluder, error) {
	return NewProject(root, pf.fr)
}

func NewProjectFactory(fr packager.FileReader) packager.ProjectFactory {
	return &projectFactory{fr}
}
