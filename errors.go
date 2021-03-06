package packager

import "errors"

var FileSystemError = errors.New("filesystem error")
var InvalidPackageError = errors.New("invalid package")
var ConfigurationError = errors.New("configuration error")
