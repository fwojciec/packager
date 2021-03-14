package md5

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"path/filepath"

	"github.com/fwojciec/packager"
	"golang.org/x/sync/errgroup"
)

type md5Hasher struct {
	dl packager.DirLister
	fr packager.FileReader
}

// New returns a new instance of MD5 Hasher.
func New(dl packager.DirLister, fr packager.FileReader) packager.Hasher {
	return &md5Hasher{dl, fr}
}

func (h *md5Hasher) Hash(project packager.LocatorExcluder) (string, error) {
	root := project.Location()

	projFiles, err := h.dl.ListDir(root, project.Exclude)
	if err != nil {
		return "", err
	}

	md5sums, err := h.getFileHashes(projFiles)
	if err != nil {
		return "", err
	}

	buf := &bytes.Buffer{}
	for i, path := range projFiles {
		relPath, err := filepath.Rel(root, path)
		if err != nil {
			return "", err
		}
		buf.WriteString(filepath.ToSlash(relPath))
		buf.WriteString(md5sums[i])
		buf.WriteString("\n")
	}

	dirHash := md5.Sum(buf.Bytes())
	return hex.EncodeToString(dirHash[:]), nil
}

func (h *md5Hasher) getFileHashes(files []string) ([]string, error) {
	res := make([]string, len(files))

	g := new(errgroup.Group)
	for i, path := range files {
		i, path := i, path
		g.Go(func() error {
			b, err := h.fr.ReadFile(path)
			if err != nil {
				return err
			}
			hash := md5.Sum(b)
			res[i] = hex.EncodeToString(hash[:])
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}

	return res, nil
}
