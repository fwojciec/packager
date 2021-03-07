package lambda_test

import (
	"testing"

	"github.com/fwojciec/packager"
	"github.com/fwojciec/packager/lambda"
	"github.com/fwojciec/packager/mocks"
)

func TestLambdaPackagerPackagesPythonProject(t *testing.T) {
	t.Parallel()

	mockProjectFactory := &mocks.ProjectFactoryMock{
		NewProjectFunc: func(lang packager.Language, root string) (packager.Project, error) { return nil, nil },
	}

	mockBuilder := &mocks.BuilderMock{
		BuildFunc: func() (string, error) { return "", nil },
		CloseFunc: func() error { return nil },
	}

	mockBuilderFactory := &mocks.BuilderFactoryMock{
		NewBuilderFunc: func(project packager.Project) (packager.Builder, error) { return mockBuilder, nil },
	}

	mockArchiver := &mocks.ArchiverMock{
		ArchiveFunc: func(target, path string) error { return nil },
	}

	subject := lambda.NewPackager(mockProjectFactory, mockBuilderFactory, mockArchiver)

	// act
	err := subject.Package(packager.LanguagePython, "./project_dir", "./out/package.zip")
	ok(t, err)

	// assert
	equals(t, 1, len(mockProjectFactory.NewProjectCalls()))
	equals(t, 1, len(mockBuilderFactory.NewBuilderCalls()))
	equals(t, 1, len(mockBuilder.BuildCalls()))
	equals(t, 1, len(mockArchiver.ArchiveCalls()))
	equals(t, 1, len(mockBuilder.CloseCalls()))
}
