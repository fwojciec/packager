package project

import (
	"github.com/fwojciec/packager"
	"github.com/gobwas/glob"
)

type excluder struct {
	globs map[string]glob.Glob
}

var generalExcludePatterns []string = []string{
	".vscode/",
	".vim/",
	".DS_Store",
}

var pythonExcludePatterns []string = []string{
	"__pycache__/",
	"*.py[cod]",
	"*$py.class",
	".pytest_cache/",
	"docs/_build/",
}

func (e *excluder) match(path string) bool {
	for _, g := range e.globs {
		if g.Match(path) {
			return true
		}
	}
	return false
}

func (e *excluder) add(patterns []string) error {
	for _, p := range patterns {
		if _, ok := e.globs[p]; ok {
			continue
		}
		g, err := glob.Compile(p)
		if err != nil {
			return err
		}
		e.globs[p] = g
	}
	return nil
}

func newExcluder(lang packager.Language, patterns ...string) (*excluder, error) {
	ex := &excluder{
		globs: make(map[string]glob.Glob),
	}
	patterns = append(patterns, generalExcludePatterns...)
	switch lang {
	case packager.LanguagePython:
		patterns = append(patterns, pythonExcludePatterns...)
	}
	if err := ex.add(patterns); err != nil {
		return nil, err
	}
	return ex, nil
}
