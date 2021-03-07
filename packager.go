package packager

// Archiver creates an archive at path with target directory contents.
type Archiver interface {
	Archive(target, path string) error
}

// Copier copies source to destination.
type Copier interface {
	Copy(source, destination string) error
}

// Builder builds a package at target for a given language and returns path to
// the build directory.
type Builder interface {
	Build() (string, error)
	Close() error
}

// BuilderFactory creates builder instances.
type BuilderFactory interface {
	NewBuilder(project Project) (Builder, error)
}

// Packager creates a deployable package at destination for the source code at target.
type Packager interface {
	Package(lang Language, target, destination string) error
}

// Project represents a source code repository to be packaged.
type Project interface {
	// Hash returns a unique hash of project snapshot.
	Hash() string
	// Files returns a list of project files.
	Files() []string
	// Language returns the project language.
	Language() Language
}

// ProjectFactory creates project instances.
type ProjectFactory interface {
	NewProject(lang Language, root string) (Project, error)
}

// Language is a programming language.
type Language string

const (
	LanguagePython     Language = "python"
	LanguageJavaScript Language = "javascript"
	LanguageTypeScript Language = "typescript"
)
