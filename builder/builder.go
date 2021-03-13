package builder

import "github.com/fwojciec/packager"

type buildFunc func(project packager.Locator) error

func (f buildFunc) Build(project packager.Locator) error {
	return f(project)
}
