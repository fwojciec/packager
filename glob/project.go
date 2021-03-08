package glob

import (
	"bufio"
	"bytes"
	"path"
	"strings"

	"github.com/fwojciec/packager"
	"github.com/gobwas/glob"
)

type project struct {
	fr    packager.FileReader
	globs map[string]glob.Glob
	root  string
}

func NewProject(root string, fr packager.FileReader) (packager.Project, error) {
	p := &project{
		fr:    fr,
		globs: make(map[string]glob.Glob),
		root:  root,
	}
	if err := p.addGlob(packager.IGNORE_FILE); err != nil {
		return nil, err
	}
	patterns, err := p.readIgnoreFile()
	if err != nil {
		return nil, err
	}
	for _, pattern := range patterns {
		if err := p.addGlob(pattern); err != nil {
			return nil, err
		}
	}
	return p, nil
}

func (p *project) Exclude(path string) bool {
	for _, g := range p.globs {
		if g.Match(path) {
			return true
		}
	}
	return false
}

func (p *project) Root() string {
	return p.root
}

func (p *project) addGlob(pattern string) error {
	if _, ok := p.globs[pattern]; ok {
		return nil
	}
	g, err := glob.Compile(pattern)
	if err != nil {
		return err
	}
	p.globs[pattern] = g
	return nil
}

func (p *project) readIgnoreFile() ([]string, error) {
	b, err := p.fr.ReadFile(path.Join(p.Root(), packager.IGNORE_FILE))
	if err != nil {
		// ignore error, since ignoreFile is not required
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