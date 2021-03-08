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
// 			NewFunc: func(root string, lang packager.Language) (packager.Project, error) {
// 				panic("mock out the New method")
// 			},
// 		}
//
// 		// use mockedProjectFactory in code that requires packager.ProjectFactory
// 		// and then make assertions.
//
// 	}
type ProjectFactoryMock struct {
	// NewFunc mocks the New method.
	NewFunc func(root string, lang packager.Language) (packager.Project, error)

	// calls tracks calls to the methods.
	calls struct {
		// New holds details about calls to the New method.
		New []struct {
			// Root is the root argument value.
			Root string
			// Lang is the lang argument value.
			Lang packager.Language
		}
	}
	lockNew sync.RWMutex
}

// New calls NewFunc.
func (mock *ProjectFactoryMock) New(root string, lang packager.Language) (packager.Project, error) {
	if mock.NewFunc == nil {
		panic("ProjectFactoryMock.NewFunc: method is nil but ProjectFactory.New was just called")
	}
	callInfo := struct {
		Root string
		Lang packager.Language
	}{
		Root: root,
		Lang: lang,
	}
	mock.lockNew.Lock()
	mock.calls.New = append(mock.calls.New, callInfo)
	mock.lockNew.Unlock()
	return mock.NewFunc(root, lang)
}

// NewCalls gets all the calls that were made to New.
// Check the length with:
//     len(mockedProjectFactory.NewCalls())
func (mock *ProjectFactoryMock) NewCalls() []struct {
	Root string
	Lang packager.Language
} {
	var calls []struct {
		Root string
		Lang packager.Language
	}
	mock.lockNew.RLock()
	calls = mock.calls.New
	mock.lockNew.RUnlock()
	return calls
}
