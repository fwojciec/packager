package mocks

//go:generate moq -out archiver_mock.go -pkg mocks .. Archiver
//go:generate moq -out builder_factory_mock.go -pkg mocks .. BuilderFactory
//go:generate moq -out builder_mock.go -pkg mocks .. Builder
//go:generate moq -out file_reader_mock.go -pkg mocks .. FileReader
//go:generate moq -out file_system_mock.go -pkg mocks .. FileSystem
//go:generate moq -out isolator_mock.go -pkg mocks .. Isolator
//go:generate moq -out isolated_project_mock.go -pkg mocks .. IsolatedProject
//go:generate moq -out project_factory_mock.go -pkg mocks .. ProjectFactory
//go:generate moq -out project_mock.go -pkg mocks .. Project
