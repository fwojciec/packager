package packager

const IGNORE_FILE = "./lambdaignore"

// Archiver creates an archive at path with target directory contents.
type Archiver interface {
	Archive(isolatedProject IsolatedProject, path string) error
}

// Builder builds a package at target.
type Builder interface {
	Build(isolatedProject IsolatedProject) error
}

// BuilderFactory creates builder instances.
type BuilderFactory interface {
	New(lang Language) (Builder, error)
}

// FileReader reads file contents as byte slice.
type FileReader interface {
	ReadFile(path string) ([]byte, error)
}

// FileSystem abstracts file system operations.
type FileSystem interface {
	ListDir(path string) ([]string, error)
	MakeTempDir() (string, error)
}

// Isolator creates a temporary copy of the project to enable safe and clean
// build.  Close removes the isolated copy.
type Isolator interface {
	Isolate(project Project) (IsolatedProject, error)
}

// IsolatedProject is a temporary copy of a source directory.
type IsolatedProject interface {
	Remove() error
	Root() string
}

// Project represents a source code repository.
type Project interface {
	Exclude(path string) bool
	Root() string
}

// ProjectFactory creates project instances.
type ProjectFactory interface {
	New(root string, lang Language) (Project, error)
}

// Language is a programming language.
type Language string

const (
	LanguagePython     Language = "python"
	LanguageJavaScript Language = "javascript"
	LanguageTypeScript Language = "typescript"
)
