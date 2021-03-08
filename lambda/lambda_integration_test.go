// +build integration

package lambda_test

import (
	"archive/zip"
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"path"
	"testing"

	"github.com/fwojciec/packager/lambda"
)

func TestIntegrationPackageAPythonPackageWithNoDependencies(t *testing.T) {
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
	subject, err := lambda.NewPackager()
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
