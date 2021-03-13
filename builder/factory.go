package builder

import (
	"github.com/fwojciec/packager"
)

type builderFactoryFunc func(lang packager.Language) packager.Builder

func (f builderFactoryFunc) New(lang packager.Language) packager.Builder { return f(lang) }

func NewBuilderFactory() packager.BuilderFactory {
	return builderFactoryFunc(func(lang packager.Language) packager.Builder {
		switch lang {
		case packager.LanguagePython:
			return &pythonBuilder{}
		default:
			noopBuilder := buildFunc(func(project packager.Locator) error { return nil })
			return noopBuilder
		}
	})
}
