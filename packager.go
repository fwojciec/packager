package packager

import (
	"fmt"
)

// FileSystem abstracts various file system operations.
type FileSystem interface {
	ReadDir(target string) ([]string, error)
	CopyDirToTemp(target string) (string, error)
}

type Validator interface {
	DirIsValid(language Language, fileList []string) bool
}

// Builder creates a clean build of a deployment package.
type Builder interface {
	Build(language Language, target string) (string, error)
}

// Zipper zips a build directory.
type Zipper interface {
	Zip(buildDir string) (string, error)
}

type Language string

const (
	LanguagePython     Language = "python"
	LanguageJavaScript Language = "javascript"
	LanguageTypeScript Language = "typescript"
)

type Packager struct {
	Lang    Language
	Builder Builder
	Zipper  Zipper
}

func (p *Packager) Package(target string) (string, error) {
	buildDir, err := p.Builder.Build(p.Lang, target)
	if err != nil {
		return "", err
	}
	zipPath, err := p.Zipper.Zip(buildDir)
	return zipPath, nil
}

func New(lang string) (*Packager, error) {
	pkg := &Packager{}
	switch lang {
	case "python":
		pkg.Lang = LanguagePython
	default:
		return nil, fmt.Errorf("%w: unrecognized language value: %s", ConfigurationError, lang)
	}

	return pkg, nil
}
