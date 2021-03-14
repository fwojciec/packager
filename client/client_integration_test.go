// +build integration

package client_test

import (
	"archive/zip"
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"testing"

	"github.com/fwojciec/packager"
	"github.com/fwojciec/packager/client"
)

func TestPackagesAPythonProjectWithNoDependencies(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	t.Parallel()

	// arrange
	target, err := ioutil.TempDir("", "target")
	ok(t, err)
	t.Cleanup(func() { os.RemoveAll(target) })

	destDir, err := ioutil.TempDir("", "destination")
	ok(t, err)
	t.Cleanup(func() { os.RemoveAll(destDir) })

	destination := path.Join(destDir, "destination.zip")

	testContents := "print(\"hello world\")\n"

	f, err := os.Create(path.Join(target, "main.py"))
	ok(t, err)
	f.WriteString(testContents)
	f.Close()

	t.Cleanup(func() { os.RemoveAll(target) })

	// act
	subject := client.New()
	ok(t, err)
	err = subject.Package("python", target, destination)
	ok(t, err)

	// assert
	r, err := zip.OpenReader(destination)
	ok(t, err)
	defer r.Close()

	foundNames := make([]string, 0)
	foundContents := make([]string, 0)

	for _, zf := range r.File {
		foundNames = append(foundNames, zf.Name)
		rc, err := zf.Open()
		ok(t, err)
		defer rc.Close()
		buf := new(bytes.Buffer)
		_, err = io.Copy(buf, rc)
		ok(t, err)
		foundContents = append(foundContents, buf.String())
	}

	equals(t, []string{"main.py"}, foundNames)
	equals(t, []string{testContents}, foundContents)
}

func TestHashesAPythonProject(t *testing.T) {
	t.Parallel()

	testFiles := []testFile{
		{filepath.Clean(".lambdaignore"), []byte("*_test.py\n")},
		{filepath.Clean("main.py"), []byte("print(\"hello world\")\n")},
		{filepath.Clean("main_test.py"), []byte("")},
		{filepath.Clean("requirements.txt"), []byte("certifi==2020.12.5\nchardet==4.0.0\nidna==2.10\nurllib3==1.26.3\nrequests==2.25.1\n")},
		{filepath.Clean("subpackage/__init__.py"), []byte("")},
		{filepath.Clean("subpackage/subpackage.py"), []byte("# just a comment\n")},
	}

	root := createFilesInTemporaryDirectory(t, testFiles)
	t.Cleanup(func() { os.RemoveAll(root) })

	subject := client.New()

	res, err := subject.Hash(packager.LanguagePython, root)
	ok(t, err)

	equals(t, "9847721f4b50a480344bb1ac4ced4ffd", res)
}
