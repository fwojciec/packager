package project

import (
	"bufio"
	"bytes"
	"fmt"
	"path"
	"strings"

	"github.com/fwojciec/packager"
)

const IGNORE_FILE = ".lambdaignore"

type ignoringProject struct {
	root string
	lang packager.Language
	fs   packager.FileSystem
	ex   *excluder
}

func New(lang packager.Language, root string, fs packager.FileSystem) (packager.Project, error) {
	ip := &ignoringProject{
		root: root,
		lang: lang,
		fs:   fs,
	}
	patterns, err := ip.getIgnorePatterns()
	if err != nil {
		return nil, err
	}
	patterns = append(patterns, IGNORE_FILE)
	ex, err := newExcluder(lang, patterns...)
	if err != nil {
		return nil, err
	}
	ip.ex = ex
	return ip, nil
}

// Hash returns a unique hash of project snapshot.
func (ip *ignoringProject) Hash() (string, error) {
	panic("not implemented") // TODO: Implement
}

// Files returns a list of project files.
// NOTE: this is not the most efficient implementation possible, but it makes life easy.
func (ip *ignoringProject) Files() ([]string, error) {
	allFiles, err := ip.fs.Dir(ip.root)
	if err != nil {
		return nil, fmt.Errorf("error listing project directory files: %s", err)
	}
	files := make([]string, 0)
	for _, file := range allFiles {
		if ip.ex.match(file) {
			continue
		}
		files = append(files, file)
	}
	return files, nil
}

// Language returns the project language.
func (ip *ignoringProject) Language() packager.Language {
	return ip.lang
}

func (ip *ignoringProject) getIgnorePatterns() ([]string, error) {
	b, err := ip.fs.ReadFile(path.Join(ip.root, IGNORE_FILE))
	if err != nil {
		// we don't care, there's no ignore file
		return nil, nil
	}
	res := make([]string, 0)
	scanner := bufio.NewScanner(bytes.NewBuffer(b))
	for scanner.Scan() {
		res = append(res, strings.TrimSpace(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return res, nil
}
