// +build integration

package copy_test

import (
	"os"
	"path"
	"strings"
	"testing"

	"github.com/fwojciec/packager/copy"
	"github.com/fwojciec/packager/mocks"
)

func TestIsolatorCreatesAnIsolatedProject(t *testing.T) {
	t.Parallel()

	testFileName := "test.txt"
	testFileContents := []byte("test file contents\n")

	srcDir := createFilesInTemporaryDirectory(t, map[string][]byte{
		testFileName: testFileContents,
	})
	t.Cleanup(func() { os.RemoveAll(srcDir) })

	mockProject := &mocks.LocatorExcluderMock{
		LocationFunc: func() string { return srcDir },
		ExcludeFunc:  func(path string) bool { return false },
	}

	subject := copy.NewIsolator()

	res, err := subject.Isolate(mockProject)
	ok(t, err)
	t.Cleanup(func() { res.Remove() })

	contents, err := os.ReadFile(path.Join(res.Location(), testFileName))
	ok(t, err)

	equals(t, testFileContents, contents)
}

func TestIsolatorExcludesExcludedFiles(t *testing.T) {
	t.Parallel()

	testIncludedFileName := "included.txt"
	testExcludedFileName := "excluded.txt"

	srcDir := createFilesInTemporaryDirectory(t, map[string][]byte{
		testIncludedFileName: {},
		testExcludedFileName: {},
	})
	t.Cleanup(func() { os.RemoveAll(srcDir) })

	mockProject := &mocks.LocatorExcluderMock{
		LocationFunc: func() string { return srcDir },
		ExcludeFunc: func(path string) bool {
			if strings.HasSuffix(path, testExcludedFileName) {
				return true
			}
			return false
		},
	}

	subject := copy.NewIsolator()

	res, err := subject.Isolate(mockProject)
	ok(t, err)
	t.Cleanup(func() { res.Remove() })

	// included file should exist
	_, err = os.Stat(path.Join(res.Location(), testIncludedFileName))
	assert(t, err == nil, "included file should exist")

	// excluded file should not exist
	_, err = os.Stat(path.Join(res.Location(), testExcludedFileName))
	assert(t, os.IsNotExist(err), "excluded file should not exist")
}
