package mocks

//go:generate moq -out archiver_mock.go -pkg mocks .. Archiver
//go:generate moq -out builder_mock.go -pkg mocks .. Builder
//go:generate moq -out builder_factory_mock.go -pkg mocks .. BuilderFactory
//go:generate moq -out file_reader_mock.go -pkg mocks .. FileReader
//go:generate moq -out isolator_mock.go -pkg mocks .. Isolator
//go:generate moq -out locator_excluder_mock.go -pkg mocks .. LocatorExcluder
//go:generate moq -out locator_remover_mock.go -pkg mocks .. LocatorRemover
//go:generate moq -out project_factory_mock.go -pkg mocks .. ProjectFactory
