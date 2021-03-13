package packager_test

import (
	"testing"

	"github.com/fwojciec/packager"
	"github.com/fwojciec/packager/mocks"
)

func TestClientPackagesPythonProjects(t *testing.T) {
	t.Parallel()

	mockProject := &mocks.LocatorExcluderMock{}

	mockProjectFactory := &mocks.ProjectFactoryMock{
		NewFunc: func(root string) (packager.LocatorExcluder, error) { return mockProject, nil },
	}

	mockTempProject := &mocks.LocatorRemoverMock{
		RemoveFunc: func() error { return nil },
	}

	mockIsolator := &mocks.IsolatorMock{
		IsolateFunc: func(project packager.LocatorExcluder) (packager.LocatorRemover, error) { return mockTempProject, nil },
	}

	mockBuilder := &mocks.BuilderMock{
		BuildFunc: func(isolatedProject packager.LocatorRemover) error { return nil },
	}

	mockBuilderFactory := &mocks.BuilderFactoryMock{
		NewFunc: func(lang packager.Language) (packager.Builder, error) { return mockBuilder, nil },
	}

	mockArchiver := &mocks.ArchiverMock{
		ArchiveFunc: func(tempProject packager.LocatorRemover, path string) error { return nil },
	}

	subject := &packager.Client{
		ProjectFactory: mockProjectFactory,
		BuilderFactory: mockBuilderFactory,
		Isolator:       mockIsolator,
		Archiver:       mockArchiver,
	}

	// act
	err := subject.Package(packager.LanguagePython, "./project_dir", "./out/package.zip")

	// assert
	ok(t, err)
	equals(t, 1, len(mockProjectFactory.NewCalls()))
	equals(t, "./project_dir", mockProjectFactory.NewCalls()[0].Root)
	equals(t, 1, len(mockIsolator.IsolateCalls()))
	equals(t, mockProject, mockIsolator.IsolateCalls()[0].Project)
	equals(t, 1, len(mockBuilderFactory.NewCalls()))
	equals(t, packager.LanguagePython, mockBuilderFactory.NewCalls()[0].Lang)
	equals(t, 1, len(mockBuilder.BuildCalls()))
	equals(t, mockTempProject, mockBuilder.BuildCalls()[0].TempProject)
	equals(t, 1, len(mockArchiver.ArchiveCalls()))
	equals(t, mockTempProject, mockArchiver.ArchiveCalls()[0].TempProject)
	equals(t, "./out/package.zip", mockArchiver.ArchiveCalls()[0].Path)
	equals(t, 1, len(mockTempProject.RemoveCalls()))
}
