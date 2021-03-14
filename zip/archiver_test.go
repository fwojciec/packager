package zip_test

import (
	"archive/zip"
	"os"
	"path"
	"path/filepath"
	"testing"

	"github.com/fwojciec/packager/fs"
	"github.com/fwojciec/packager/mocks"
	myzip "github.com/fwojciec/packager/zip"
)

func TestZipArchiver(t *testing.T) {
	t.Parallel()

	testFiles := []testFile{
		{filepath.Clean("main.py"), []byte("main.py contents\n")},
		{filepath.Clean("submodule/__init__.py"), []byte{}},
		{filepath.Clean("submodule/submodule.py"), []byte("submodule.py contents\n")},
	}
	srcDir := createFilesInTemporaryDirectory(t, testFiles)
	t.Cleanup(func() { os.RemoveAll(srcDir) })

	destDir, err := os.MkdirTemp("", "")
	ok(t, err)
	t.Cleanup(func() { os.RemoveAll(destDir) })

	project := &mocks.LocatorMock{LocationFunc: func() string { return srcDir }}

	subject := myzip.New(fs.NewDirLister())
	err = subject.Archive(project, path.Join(destDir, "test.zip"))
	ok(t, err)

	zrc, err := zip.OpenReader(path.Join(destDir, "test.zip"))
	ok(t, err)
	defer zrc.Close()
	for i, expFile := range testFiles {
		equals(t, expFile.name, zrc.File[i].Name)
		equals(t, expFile.contents, readContents(t, zrc.File[i]))
	}
}
