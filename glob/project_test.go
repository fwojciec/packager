package glob_test

import (
	"path/filepath"
	"testing"

	"github.com/fwojciec/packager/glob"
	"github.com/fwojciec/packager/mocks"
)

func TestProjectReturnsItsAbsoluteLocation(t *testing.T) {
	t.Parallel()
	mockFileReader := &mocks.FileReaderMock{
		ReadFileFunc: func(path string) ([]byte, error) { return nil, nil },
	}
	subject, err := glob.NewProject("/root", "", mockFileReader)
	ok(t, err)

	result := subject.Location()

	expected, _ := filepath.Abs("/root")
	equals(t, expected, result)
}

func TestProjectDoesntExcludeProjectFiles(t *testing.T) {
	t.Parallel()
	mockFileReader := &mocks.FileReaderMock{
		ReadFileFunc: func(path string) ([]byte, error) { return nil, nil },
	}
	subject, err := glob.NewProject("/project/root", "", mockFileReader)
	ok(t, err)

	result, err := subject.Exclude(filepath.Clean("/project/root/handler.py"))
	ok(t, err)

	assert(t, !result, "regular project files shouldn't be excluded")
}

func TestProjectExcludesIgnoreFile(t *testing.T) {
	t.Parallel()
	mockFileReader := &mocks.FileReaderMock{
		ReadFileFunc: func(path string) ([]byte, error) { return nil, nil },
	}
	subject, err := glob.NewProject("/project/root", "", mockFileReader)
	ok(t, err)

	result, err := subject.Exclude(filepath.Clean("/project/root/.lambdaignore"))
	ok(t, err)

	assert(t, result, "ignore file should be excluded")
}

func TestProjectExcludesIgnoreFileGlobMatches(t *testing.T) {
	t.Parallel()
	mockFileReader := &mocks.FileReaderMock{
		ReadFileFunc: func(path string) ([]byte, error) { return []byte("*_test.py"), nil },
	}
	subject, err := glob.NewProject("/project/root", "", mockFileReader)
	ok(t, err)

	result, err := subject.Exclude(filepath.Clean("/project/root/handler_test.py"))
	ok(t, err)

	assert(t, result, "files matching ignore globs should be excluded")
}
