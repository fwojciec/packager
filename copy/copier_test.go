package copy_test

// func TestCopierCopiesFilesToATemporaryDirectory(t *testing.T) {
// 	t.Parallel()

// 	testFileName := "test.txt"
// 	testFileContents := "test file contents\n"

// 	srcDir := createFilesInTemporaryDirectory(t, map[string]string{
// 		testFileName: testFileContents,
// 	})
// 	t.Cleanup(func() { os.RemoveAll(srcDir) })
// 	subject := packager.NewCopier()

// 	destDir, err := subject.CopyDir(srcDir, nil)
// 	ok(t, err)
// 	t.Cleanup(func() { os.RemoveAll(destDir) })

// 	f, err := os.Open(path.Join(destDir, "test.txt"))
// 	ok(t, err)
// 	defer f.Close()
// 	buf := new(bytes.Buffer)
// 	_, err = io.Copy(buf, f)
// 	ok(t, err)

// 	t.Log(srcDir)
// 	t.Log(destDir)

// 	equals(t, testFileContents, buf.String())
// }

// func createFilesInTemporaryDirectory(t *testing.T, config map[string]string) string {
// 	tDir, err := ioutil.TempDir("", "")
// 	if err != nil {
// 		t.Fatal("failed to create temporary directory:", err)
// 	}

// 	for p, c := range config {
// 		dir, file := filepath.Split(p)
// 		if dir != "" {
// 			if err := os.MkdirAll(path.Join(tDir, dir), os.ModePerm); err != nil {
// 				t.Fatal("failed to create directory:", err)
// 			}
// 		}
// 		f, err := os.Create(path.Join(tDir, dir, file))
// 		if err != nil {
// 			t.Fatal("failed to create test file:", err)
// 		}
// 		_, err = f.WriteString(c)
// 		if err != nil {
// 			t.Fatal("failed to write to test file:", err)
// 		}
// 		if err := f.Close(); err != nil {
// 			t.Fatal("failed to close the test file:", err)
// 		}
// 	}

// 	return tDir
// }
