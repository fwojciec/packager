package builder

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path"

	"github.com/fwojciec/packager"
)

type pythonBuilder struct{}

func (pb *pythonBuilder) Build(project packager.Locator) error {
	projectRoot := project.Location()

	if pb.hasNoRequirementsFile(projectRoot) {
		return nil
	}

	if err := pb.install(projectRoot); err != nil {
		return err
	}

	if err := pb.postInstall(projectRoot); err != nil {
		return err
	}

	return nil
}

func (pb *pythonBuilder) hasNoRequirementsFile(projectRoot string) bool {
	_, err := os.Stat(path.Join(projectRoot, "requirements.txt"))
	return os.IsNotExist(err)
}

func (pb *pythonBuilder) install(projectRoot string) error {
	pip, err := exec.LookPath("pip")
	if err != nil {
		return fmt.Errorf("pip command not found")
	}
	cmd := exec.Command(
		pip,
		"--no-color",
		"--disable-pip-version-check",
		"--no-python-version-warning",
		"install",
		"-r",
		"requirements.txt",
		"--target",
		".",
	)
	cmd.Dir = projectRoot
	var errOut bytes.Buffer
	cmd.Stderr = &errOut

	if err := cmd.Run(); err != nil {
		if _, ok := err.(*exec.ExitError); ok {
			return fmt.Errorf("pip command error: %q", errOut.String())
		}
		return err
	}

	return nil
}

func (pb *pythonBuilder) postInstall(projectRoot string) error {
	if err := os.Remove(path.Join(projectRoot, "requirements.txt")); err != nil {
		return err
	}
	if err := os.RemoveAll(path.Join(projectRoot, "bin")); err != nil {
		return err
	}
	return nil
}
