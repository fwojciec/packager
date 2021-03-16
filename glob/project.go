package glob

import (
	"bufio"
	"bytes"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/fwojciec/packager"
	"github.com/gobwas/glob"
)

type project struct {
	fr    packager.FileReader
	globs map[string]glob.Glob
	root  string
}

func NewProject(root string, lang packager.Language, fr packager.FileReader) (packager.LocatorExcluder, error) {
	p := &project{
		fr:    fr,
		root:  filepath.Clean(root),
		globs: make(map[string]glob.Glob),
	}
	if err := p.addIgnoreFileGlobs(); err != nil {
		return nil, err
	}
	if err := p.addUniversalGlobs(); err != nil {
		return nil, err
	}
	if err := p.addLanguageSpecificGlobs(lang); err != nil {
		return nil, err
	}
	return p, nil
}

func (p *project) Exclude(path string) (bool, error) {
	for _, g := range p.globs {
		relPath, err := filepath.Rel(p.root, path)
		if err != nil {
			return false, err
		}
		if g.Match(relPath) {
			return true, nil
		}
	}
	return false, nil
}

func (p *project) Location() string {
	return p.root
}

func (p *project) addGlobs(patterns []string) error {
	for _, pattern := range patterns {
		if _, ok := p.globs[pattern]; ok {
			return nil
		}
		g, err := glob.Compile(pattern)
		if err != nil {
			return err
		}
		p.globs[pattern] = g
	}
	return nil
}

func (p *project) addIgnoreFileGlobs() error {
	b, err := p.fr.ReadFile(path.Join(p.root, packager.IGNORE_FILE))
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	patterns := make([]string, 0)
	scanner := bufio.NewScanner(bytes.NewBuffer(b))
	for scanner.Scan() {
		patterns = append(patterns, strings.TrimSpace(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	if err := p.addGlobs(patterns); err != nil {
		return err
	}
	return nil
}

var DEFAULT_UNIVERSAL_GLOBS = []string{
	packager.IGNORE_FILE,
	".DS_Store",
}

func (p *project) addUniversalGlobs() error {
	return p.addGlobs(DEFAULT_UNIVERSAL_GLOBS)
}

var DEFAULT_PYTHON_GLOBS = []string{
	"__pycache__/",
	"*.py[cod]",
	".pytest_cache/",
}

func (p *project) addLanguageSpecificGlobs(lang packager.Language) error {
	switch lang {
	case packager.LanguagePython:
		return p.addGlobs(DEFAULT_PYTHON_GLOBS)
	}
	return nil
}
