package packager_test

import (
	"testing"

	"github.com/fwojciec/packager"
	"github.com/fwojciec/packager/mocks"
)

func TestPackagerClientPackagesPythonProject(t *testing.T) {
	t.Parallel()

	mockProjectFactory := &mocks.ProjectFactoryMock{
		NewFunc: func(root string, lang packager.Language) (packager.Project, error) { return nil, nil },
	}

	mockIsolatedProject := &mocks.IsolatedProjectMock{
		RemoveFunc: func() error { return nil },
	}

	mockIsolator := &mocks.IsolatorMock{
		IsolateFunc: func(project packager.Project) (packager.IsolatedProject, error) { return mockIsolatedProject, nil },
	}

	mockBuilder := &mocks.BuilderMock{
		BuildFunc: func(isolatedProject packager.IsolatedProject) error { return nil },
	}

	mockBuilderFactory := &mocks.BuilderFactoryMock{
		NewFunc: func(lang packager.Language) (packager.Builder, error) { return mockBuilder, nil },
	}

	mockArchiver := &mocks.ArchiverMock{
		ArchiveFunc: func(isolatedProject packager.IsolatedProject, path string) error { return nil },
	}

	subject := &packager.Client{
		ProjectFactory: mockProjectFactory,
		BuilderFactory: mockBuilderFactory,
		Isolator:       mockIsolator,
		Archiver:       mockArchiver,
	}

	// act
	err := subject.Package(packager.LanguagePython, "./project_dir", "./out/package.zip")
	ok(t, err)

	// assert
	equals(t, 1, len(mockProjectFactory.NewCalls()))
	equals(t, 1, len(mockIsolator.IsolateCalls()))
	equals(t, 1, len(mockBuilderFactory.NewCalls()))
	equals(t, 1, len(mockBuilder.BuildCalls()))
	equals(t, 1, len(mockArchiver.ArchiveCalls()))
	equals(t, 1, len(mockIsolatedProject.RemoveCalls()))
}
