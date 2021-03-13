package packager

import "io"

const IGNORE_FILE = "./lambdaignore"

// Archiver creates an archive at path with target directory contents.
type Archiver interface {
	Archive(project Locator, dest string) error
}

// Builder builds a package at target.
type Builder interface {
	Build(project Locator) error
}

// BuilderFactory creates builder instances.
type BuilderFactory interface {
	New(lang Language) Builder
}

// DirLister returns a list of all files in a directory.
type DirLister interface {
	ListDir(target string) ([]string, error)
}

// Excluder knows how to exclude paths.
type Excluder interface {
	Exclude(path string) bool
}

// FileCopier writes a file at srcPath to dest.
type FileCopier interface {
	Copy(srcPath string, dest io.Writer) error
}

// FileReader reads file contents as byte slice.
type FileReader interface {
	ReadFile(path string) ([]byte, error)
}

// Isolator creates a temporary copy of the project to enable safe and clean
// build.  Close removes the isolated copy.
type Isolator interface {
	Isolate(project LocatorExcluder) (LocatorRemover, error)
}

// Locator can return it's location.
type Locator interface {
	Location() string
}

// LocatorRemover knows it's location and knows how to remove itself.
type LocatorRemover interface {
	Locator
	Remover
}

// LocatorExcluder knows it's location and knows how to exclude files.
type LocatorExcluder interface {
	Locator
	Excluder
}

// ProjectFactory creates project instances.
type ProjectFactory interface {
	New(root string) (LocatorExcluder, error)
}

// Remover knows how to remove itself.
type Remover interface {
	Remove() error
}

// Language is a programming language.
type Language string

const (
	LanguagePython     Language = "python"
	LanguageJavaScript Language = "javascript"
	LanguageTypeScript Language = "typescript"
)
