package packager

import "fmt"

type Client struct {
	ProjectFactory ProjectFactory
	BuilderFactory BuilderFactory
	Isolator       Isolator
	Archiver       Archiver
}

func (c *Client) Package(lang Language, target, destination string) error {
	p, err := c.ProjectFactory.New(target)
	if err != nil {
		return fmt.Errorf("%w: error initializing project: %s", ProjectError, err)
	}

	ip, err := c.Isolator.Isolate(p)
	if err != nil {
		return fmt.Errorf("%w: error isolating project: %s", IsolateError, err)
	}
	defer ip.Remove()

	b, err := c.BuilderFactory.New(lang)
	if err != nil {
		return fmt.Errorf("%w: error initializing builder: %s", BuildError, err)
	}

	if err := b.Build(ip); err != nil {
		return fmt.Errorf("%w: error building project: %s", BuildError, err)
	}

	if err := c.Archiver.Archive(ip, destination); err != nil {
		return fmt.Errorf("%w: error making package archive: %s", BuildError, err)
	}

	return nil
}
