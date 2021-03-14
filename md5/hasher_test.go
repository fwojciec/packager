package md5_test

import (
	"testing"

	mymd5 "github.com/fwojciec/packager/md5"
	"github.com/fwojciec/packager/mocks"
)

func TestHasher(t *testing.T) {
	t.Parallel()

	testFiles := []string{"/test/path/file.txt"}

	mockDirLister := &mocks.DirListerMock{
		ListDirFunc: func(target string, exclFn func(path string) (bool, error)) ([]string, error) { return testFiles, nil },
	}
	mockFileReader := &mocks.FileReaderMock{
		ReadFileFunc: func(path string) ([]byte, error) { return []byte("test\n"), nil },
	}
	mockProject := &mocks.LocatorExcluderMock{
		LocationFunc: func() string { return "/test/path" },
		ExcludeFunc:  func(path string) (bool, error) { return false, nil },
	}

	subject := mymd5.New(mockDirLister, mockFileReader)

	res, err := subject.Hash(mockProject)
	ok(t, err)

	equals(t, "f2348ba264125a9a82decb90318a21c0", res)
}
