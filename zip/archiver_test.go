// +build integration

package zip_test

import (
	"archive/zip"
	"bytes"
	"io"
	"os"
	"path"
	"path/filepath"
	"testing"

	"github.com/fwojciec/packager/fs"
	"github.com/fwojciec/packager/mocks"
	myzip "github.com/fwojciec/packager/zip"
)

type testFile struct {
	name     string
	contents []byte
}

func TestZipArchiver(t *testing.T) {
	t.Parallel()

	testFiles := []testFile{
		{"main.py", []byte("main.py contents\n")},
		{"submodule/__init__.py", []byte{}},
		{"submodule/submodule.py", []byte("submodule.py contents\n")},
	}
	srcDir := createFilesInTemporaryDirectory(t, testFiles)
	t.Cleanup(func() { os.RemoveAll(srcDir) })
	destDir, err := os.MkdirTemp("", "")
	ok(t, err)
	t.Cleanup(func() { os.RemoveAll(destDir) })
	project := &mocks.LocatorMock{LocationFunc: func() string { return srcDir }}

	dirLister := fs.NewDirLister()
	subject := myzip.New(dirLister)
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

func readContents(t *testing.T, zipFile *zip.File) []byte {
	f, err := zipFile.Open()
	if err != nil {
		t.Fatal("error reading zipped file")
	}
	defer f.Close()
	var buf bytes.Buffer
	_, err = io.Copy(&buf, f)
	if err != nil {
		t.Fatal("error reading zipped file contents")
	}
	return buf.Bytes()
}

func createFilesInTemporaryDirectory(t *testing.T, testFiles []testFile) string {
	tDir, err := os.MkdirTemp("", "")
	if err != nil {
		t.Fatal("failed to create temporary directory:", err)
	}

	for _, tf := range testFiles {
		dir, file := filepath.Split(tf.name)
		if dir != "" {
			if err := os.MkdirAll(path.Join(tDir, dir), os.ModePerm); err != nil {
				t.Fatal("failed to create directory:", err)
			}
		}
		f, err := os.Create(path.Join(tDir, dir, file))
		if err != nil {
			t.Fatal("failed to create test file:", err)
		}
		_, err = f.Write(tf.contents)
		if err != nil {
			t.Fatal("failed to write to test file:", err)
		}
		if err := f.Close(); err != nil {
			t.Fatal("failed to close the test file:", err)
		}
	}

	return tDir
}
