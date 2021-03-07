package lambda

import (
	"fmt"

	"github.com/fwojciec/packager"
)

type lambdaPackager struct {
	pf   packager.ProjectFactory
	bf   packager.BuilderFactory
	arch packager.Archiver
}

// NewPackager returns a new instance of Lambda project packager.
func NewPackager(pf packager.ProjectFactory, bf packager.BuilderFactory, arch packager.Archiver) packager.Packager {
	return &lambdaPackager{pf, bf, arch}
}

func (lp *lambdaPackager) Package(lang packager.Language, target, destination string) error {
	project, err := lp.pf.NewProject(lang, target)
	if err != nil {
		return fmt.Errorf("%w: error initializing project: %s", packager.ProjectError, err)
	}

	builder, err := lp.bf.NewBuilder(project)
	if err != nil {
		return fmt.Errorf("%w: error initializing builder: %s", packager.BuildError, err)
	}
	defer builder.Close()

	buildDir, err := builder.Build()
	if err != nil {
		return fmt.Errorf("%w: error building project: %s", packager.BuildError, err)
	}

	if err := lp.arch.Archive(destination, buildDir); err != nil {
		return fmt.Errorf("%w: error making package archive: %s", packager.BuildError, err)
	}

	return nil
}
