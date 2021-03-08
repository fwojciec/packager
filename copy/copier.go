package copy

import (
	"fmt"
	"io/ioutil"

	"github.com/fwojciec/packager"
	"github.com/otiai10/copy"
)

type copier struct {
}

func (c *copier) CopyDir(source string, ignorePatterns []string) (string, error) {
	destDir, err := ioutil.TempDir("", "packager")
	if err != nil {
		return "", fmt.Errorf("%w: failed to create build directory: %s", packager.BuildError, err)
	}
	if err := copy.Copy(source, destDir); err != nil {
		return "", fmt.Errorf("%w: error copying files: %s", packager.BuildError, err)
	}
	return destDir, nil
}
