# packager

`packager` provides a client for packaging local directories as deployment ready AWS Lambda zip packages. It will eventually be wrapped in a Terraform provider, but can be also used as a library.

## Features

- creates isolated builds (i.e. doesn't touch/modify original repositories)
- respects a `.lambdaignore` file in the root of a repository, if present (works similar to `.gitignore`)
- installs dependencies/builds
- supports Python packages

## Planned features

- support for JavaScript repositories
- support for TypeScript repositories
- ability to generate a unique hash for a source directory (for Terraform to be able to determine whether the project has changed since the previous deployment)
