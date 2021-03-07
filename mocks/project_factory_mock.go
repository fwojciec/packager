// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mocks

import (
	"github.com/fwojciec/packager"
	"sync"
)

// Ensure, that ProjectFactoryMock does implement packager.ProjectFactory.
// If this is not the case, regenerate this file with moq.
var _ packager.ProjectFactory = &ProjectFactoryMock{}

// ProjectFactoryMock is a mock implementation of packager.ProjectFactory.
//
// 	func TestSomethingThatUsesProjectFactory(t *testing.T) {
//
// 		// make and configure a mocked packager.ProjectFactory
// 		mockedProjectFactory := &ProjectFactoryMock{
// 			NewProjectFunc: func(lang packager.Language, root string) (packager.Project, error) {
// 				panic("mock out the NewProject method")
// 			},
// 		}
//
// 		// use mockedProjectFactory in code that requires packager.ProjectFactory
// 		// and then make assertions.
//
// 	}
type ProjectFactoryMock struct {
	// NewProjectFunc mocks the NewProject method.
	NewProjectFunc func(lang packager.Language, root string) (packager.Project, error)

	// calls tracks calls to the methods.
	calls struct {
		// NewProject holds details about calls to the NewProject method.
		NewProject []struct {
			// Lang is the lang argument value.
			Lang packager.Language
			// Root is the root argument value.
			Root string
		}
	}
	lockNewProject sync.RWMutex
}

// NewProject calls NewProjectFunc.
func (mock *ProjectFactoryMock) NewProject(lang packager.Language, root string) (packager.Project, error) {
	if mock.NewProjectFunc == nil {
		panic("ProjectFactoryMock.NewProjectFunc: method is nil but ProjectFactory.NewProject was just called")
	}
	callInfo := struct {
		Lang packager.Language
		Root string
	}{
		Lang: lang,
		Root: root,
	}
	mock.lockNewProject.Lock()
	mock.calls.NewProject = append(mock.calls.NewProject, callInfo)
	mock.lockNewProject.Unlock()
	return mock.NewProjectFunc(lang, root)
}

// NewProjectCalls gets all the calls that were made to NewProject.
// Check the length with:
//     len(mockedProjectFactory.NewProjectCalls())
func (mock *ProjectFactoryMock) NewProjectCalls() []struct {
	Lang packager.Language
	Root string
} {
	var calls []struct {
		Lang packager.Language
		Root string
	}
	mock.lockNewProject.RLock()
	calls = mock.calls.NewProject
	mock.lockNewProject.RUnlock()
	return calls
}