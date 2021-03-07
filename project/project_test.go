package project_test

import (
	"testing"

	"github.com/fwojciec/packager"
	"github.com/fwojciec/packager/mocks"
	"github.com/fwojciec/packager/project"
)

func TestProjectLanguage(t *testing.T) {
	t.Parallel()
	mockFileSystem := &mocks.FileSystemMock{
		ReadFileFunc: func(path string) ([]byte, error) { return nil, nil },
	}
	subject, err := project.New(packager.LanguagePython, "", mockFileSystem)
	ok(t, err)

	result := subject.Language()

	equals(t, packager.LanguagePython, result)
}

func TestProjectFilesIncludesProjectFiles(t *testing.T) {
	t.Parallel()
	mockFileSystem := &mocks.FileSystemMock{
		DirFunc: func(root string) ([]string, error) {
			return []string{"handler.py"}, nil
		},
		ReadFileFunc: func(path string) ([]byte, error) { return []byte{}, nil },
	}
	subject, err := project.New(packager.LanguagePython, "", mockFileSystem)
	ok(t, err)

	result, err := subject.Files()

	ok(t, err)
	equals(t, []string{"handler.py"}, result)
}

func TestProjectFilesExcludesIgnoreFile(t *testing.T) {
	t.Parallel()
	mockFileSystem := &mocks.FileSystemMock{
		DirFunc: func(root string) ([]string, error) {
			return []string{"handler.py", ".lambdaignore"}, nil
		},
		ReadFileFunc: func(path string) ([]byte, error) { return []byte{}, nil },
	}
	subject, err := project.New(packager.LanguagePython, "", mockFileSystem)

	result, err := subject.Files()

	ok(t, err)
	equals(t, []string{"handler.py"}, result)
}

func TestProjectFilesRespectsGlobsInIgnoreFile(t *testing.T) {
	t.Parallel()
	mockFileSystem := &mocks.FileSystemMock{
		DirFunc: func(root string) ([]string, error) {
			return []string{"handler.py", ".lambdaignore", "handler_test.py", "test_handler.py"}, nil
		},
		ReadFileFunc: func(path string) ([]byte, error) {
			return []byte("*_test.py\ntest_*.py\n"), nil
		},
	}
	subject, err := project.New(packager.LanguagePython, "", mockFileSystem)

	result, err := subject.Files()

	ok(t, err)
	equals(t, []string{"handler.py"}, result)
}
