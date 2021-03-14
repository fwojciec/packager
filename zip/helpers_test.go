package zip_test

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

type testWriteCloser struct {
	buf *bytes.Buffer
}

func (wc *testWriteCloser) Write(p []byte) (int, error) {
	return wc.buf.Write(p)
}

func (wc *testWriteCloser) Close() error {
	return nil
}

func newWriteCloser() io.WriteCloser {
	var buf bytes.Buffer
	return &testWriteCloser{&buf}
}

type testFile struct {
	name     string
	contents []byte
}

func readContents(t *testing.T, zipFile *zip.File) []byte {
	t.Helper()

	f, err := zipFile.Open()
	ok(t, err)
	defer f.Close()
	var buf bytes.Buffer
	_, err = io.Copy(&buf, f)
	ok(t, err)
	return buf.Bytes()
}

func createFilesInTemporaryDirectory(t *testing.T, testFiles []testFile) string {
	t.Helper()

	tDir, err := os.MkdirTemp("", "")
	ok(t, err)

	for _, tf := range testFiles {
		dir, file := filepath.Split(tf.name)
		if dir != "" {
			err := os.MkdirAll(path.Join(tDir, dir), os.ModePerm)
			ok(t, err)
		}
		f, err := os.Create(path.Join(tDir, dir, file))
		ok(t, err)
		defer f.Close()
		_, err = f.Write(tf.contents)
		ok(t, err)
	}

	return tDir
}

// assert fails the test if the condition is false.
func assert(tb testing.TB, condition bool, msg string, v ...interface{}) {
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: "+msg+"\033[39m\n\n", append([]interface{}{filepath.Base(file), line}, v...)...)
		tb.FailNow()
	}
}

// ok fails the test if an err is not nil.
func ok(tb testing.TB, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: unexpected error: %s\033[39m\n\n", filepath.Base(file), line, err.Error())
		tb.FailNow()
	}
}

// equals fails the test if exp is not equal to act.
func equals(tb testing.TB, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, exp, act)
		tb.FailNow()
	}
}
