package packager_test

import (
	"archive/zip"
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"path"
	"testing"

	"github.com/fwojciec/packager"
	"github.com/fwojciec/packager/mocks"
)

func TestPackageCallsBuilderWithLanguageAndTargetDirectory(t *testing.T) {
	t.Parallel()

	// arrange
	mockBuilder := &mocks.BuilderMock{
		BuildFunc: func(language packager.Language, target string) (string, error) { return "", nil },
	}
	mockZipper := &mocks.ZipperMock{
		ZipFunc: func(target string) (string, error) { return "", nil },
	}
	subject := &packager.Packager{
		Lang:    packager.LanguagePython,
		Builder: mockBuilder,
		Zipper:  mockZipper,
	}

	// act
	_, err := subject.Package("./test_target")
	ok(t, err)

	// assert
	equals(t, 1, len(mockBuilder.BuildCalls()))
	equals(t, packager.LanguagePython, mockBuilder.BuildCalls()[0].Language)
	equals(t, "./test_target", mockBuilder.BuildCalls()[0].Target)
}

func TestPackageCallsZipperWithBuildDirectory(t *testing.T) {
	t.Parallel()

	// arrange
	mockBuilder := &mocks.BuilderMock{
		BuildFunc: func(language packager.Language, target string) (string, error) { return "./build_dir", nil },
	}
	mockZipper := &mocks.ZipperMock{
		ZipFunc: func(target string) (string, error) { return "", nil },
	}
	subject := &packager.Packager{
		Lang:    packager.LanguagePython,
		Builder: mockBuilder,
		Zipper:  mockZipper,
	}

	// act
	_, err := subject.Package("")
	ok(t, err)

	// assert
	equals(t, 1, len(mockZipper.ZipCalls()))
	equals(t, "./build_dir", mockZipper.ZipCalls()[0].BuildDir)
}

func TestPackageReturnsPathReturnedByZipperZip(t *testing.T) {
	t.Parallel()

	// arrange
	mockBuilder := &mocks.BuilderMock{
		BuildFunc: func(language packager.Language, target string) (string, error) { return "", nil },
	}
	mockZipper := &mocks.ZipperMock{
		ZipFunc: func(target string) (string, error) { return "./path/to/file.zip", nil },
	}
	subject := &packager.Packager{
		Lang:    packager.LanguagePython,
		Builder: mockBuilder,
		Zipper:  mockZipper,
	}

	// act
	res, err := subject.Package("")
	ok(t, err)

	// assert
	equals(t, "./path/to/file.zip", res)
}

func TestIntegrationPackageAPythonPackageWithNoDependencies(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	t.Parallel()

	// arrange
	target, err := ioutil.TempDir("", "target")
	ok(t, err)
	t.Log(target)

	testContents := "print(\"hello world\")\n"

	f, err := os.Create(path.Join(target, "main.py"))
	ok(t, err)
	f.WriteString(testContents)
	f.Close()

	t.Cleanup(func() {
		os.RemoveAll(target)
	})

	// act
	subject, err := packager.New("python")
	ok(t, err)
	res, err := subject.Package(target)
	ok(t, err)

	// assert
	r, err := zip.OpenReader(res)
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
