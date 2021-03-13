// +build integration

package builder_test

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"testing"

	"github.com/fwojciec/packager"
	"github.com/fwojciec/packager/builder"
	"github.com/fwojciec/packager/mocks"
)

func TestPythonBuilderNoRequirementsFile(t *testing.T) {
	t.Parallel()

	srcDir := createFilesInTemporaryDirectory(t, map[string][]byte{
		"main.py": []byte("print(\"hello world\")\n"),
	})
	t.Cleanup(func() { os.RemoveAll(srcDir) })
	project := &mocks.LocatorMock{LocationFunc: func() string { return srcDir }}

	bf := builder.NewBuilderFactory()
	subject := bf.New(packager.LanguagePython)

	err := subject.Build(project)
	ok(t, err) // should be a no-op
}

func TestPythonBuilderWithRequirementsFile(t *testing.T) {
	t.Parallel()

	srcDir := createFilesInTemporaryDirectory(t, map[string][]byte{
		"main.py":          []byte("print(\"hello world\")\n"),
		"requirements.txt": []byte("certifi==2020.12.5\nchardet==4.0.0\nidna==2.10\nurllib3==1.26.3\nrequests==2.25.1\n"),
	})
	t.Cleanup(func() { os.RemoveAll(srcDir) })
	project := &mocks.LocatorMock{LocationFunc: func() string { return srcDir }}

	bf := builder.NewBuilderFactory()
	subject := bf.New(packager.LanguagePython)

	err := subject.Build(project)
	ok(t, err)

	res, err := os.ReadDir(srcDir)
	ok(t, err)

	expected := []struct {
		fileName string
		isDir    bool
	}{
		{"certifi", true},
		{"certifi-2020.12.5.dist-info", true},
		{"chardet", true},
		{"chardet-4.0.0.dist-info", true},
		{"idna", true},
		{"idna-2.10.dist-info", true},
		{"main.py", false},
		{"requests", true},
		{"requests-2.25.1.dist-info", true},
		{"urllib3", true},
		{"urllib3-1.26.3.dist-info", true},
	}

	for i, item := range res {
		equals(t, expected[i].fileName, item.Name())
		assert(t, expected[i].isDir == item.IsDir(), fmt.Sprintf("value of IsDir for %q should be %v", item.Name(), expected[i].isDir))
	}
}

func createFilesInTemporaryDirectory(t *testing.T, config map[string][]byte) string {
	tDir, err := os.MkdirTemp("", "")
	if err != nil {
		t.Fatal("failed to create temporary directory:", err)
	}

	for p, c := range config {
		dir, file := filepath.Split(p)
		if dir != "" {
			if err := os.MkdirAll(path.Join(tDir, dir), os.ModePerm); err != nil {
				t.Fatal("failed to create directory:", err)
			}
		}
		f, err := os.Create(path.Join(tDir, dir, file))
		if err != nil {
			t.Fatal("failed to create test file:", err)
		}
		_, err = f.Write(c)
		if err != nil {
			t.Fatal("failed to write to test file:", err)
		}
		if err := f.Close(); err != nil {
			t.Fatal("failed to close the test file:", err)
		}
	}

	return tDir
}
