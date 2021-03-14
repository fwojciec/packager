package client

import (
	"fmt"

	"github.com/fwojciec/packager"
	"github.com/fwojciec/packager/builder"
	"github.com/fwojciec/packager/copy"
	"github.com/fwojciec/packager/fs"
	"github.com/fwojciec/packager/glob"
	"github.com/fwojciec/packager/md5"
	"github.com/fwojciec/packager/zip"
)

type Packager struct {
	ProjectFactory packager.ProjectFactory
	Isolator       packager.Isolator
	BuilderFactory packager.BuilderFactory
	Archiver       packager.Archiver
	Hasher         packager.Hasher
}

func (p *Packager) Package(lang packager.Language, target, destination string) error {
	project, err := p.ProjectFactory.New(target, lang)
	if err != nil {
		return fmt.Errorf("%w: error initializing project: %s", packager.ProjectError, err)
	}

	isolatedProject, err := p.Isolator.Isolate(project)
	if err != nil {
		return fmt.Errorf("%w: error isolating project: %s", packager.IsolateError, err)
	}
	defer isolatedProject.Remove()

	builder := p.BuilderFactory.New(lang)
	if err := builder.Build(isolatedProject); err != nil {
		return fmt.Errorf("%w: error building project: %s", packager.BuildError, err)
	}

	if err := p.Archiver.Archive(isolatedProject, destination); err != nil {
		return fmt.Errorf("%w: error making package archive: %s", packager.ArchiverError, err)
	}

	return nil
}

func (p *Packager) Hash(lang packager.Language, target string) (string, error) {
	project, err := p.ProjectFactory.New(target, lang)
	if err != nil {
		return "", fmt.Errorf("%w: error initializing project: %s", packager.ProjectError, err)
	}
	res, err := p.Hasher.Hash(project)
	if err != nil {
		return "", fmt.Errorf("%w: %s", packager.ProjectError, err)
	}
	return res, nil
}

func New() *Packager {
	projectFactory := glob.NewProjectFactory(fs.NewFileReader())
	isolator := copy.NewIsolator()
	builderFactory := builder.NewBuilderFactory()
	archiver := zip.New(fs.NewDirLister())
	hasher := md5.New()

	return &Packager{
		ProjectFactory: projectFactory,
		Isolator:       isolator,
		BuilderFactory: builderFactory,
		Archiver:       archiver,
		Hasher:         hasher,
	}
}
