// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mocks

import (
	"github.com/fwojciec/packager"
	"sync"
)

// Ensure, that IsolatedProjectMock does implement packager.IsolatedProject.
// If this is not the case, regenerate this file with moq.
var _ packager.IsolatedProject = &IsolatedProjectMock{}

// IsolatedProjectMock is a mock implementation of packager.IsolatedProject.
//
// 	func TestSomethingThatUsesIsolatedProject(t *testing.T) {
//
// 		// make and configure a mocked packager.IsolatedProject
// 		mockedIsolatedProject := &IsolatedProjectMock{
// 			RemoveFunc: func() error {
// 				panic("mock out the Remove method")
// 			},
// 			RootFunc: func() string {
// 				panic("mock out the Root method")
// 			},
// 		}
//
// 		// use mockedIsolatedProject in code that requires packager.IsolatedProject
// 		// and then make assertions.
//
// 	}
type IsolatedProjectMock struct {
	// RemoveFunc mocks the Remove method.
	RemoveFunc func() error

	// RootFunc mocks the Root method.
	RootFunc func() string

	// calls tracks calls to the methods.
	calls struct {
		// Remove holds details about calls to the Remove method.
		Remove []struct {
		}
		// Root holds details about calls to the Root method.
		Root []struct {
		}
	}
	lockRemove sync.RWMutex
	lockRoot   sync.RWMutex
}

// Remove calls RemoveFunc.
func (mock *IsolatedProjectMock) Remove() error {
	if mock.RemoveFunc == nil {
		panic("IsolatedProjectMock.RemoveFunc: method is nil but IsolatedProject.Remove was just called")
	}
	callInfo := struct {
	}{}
	mock.lockRemove.Lock()
	mock.calls.Remove = append(mock.calls.Remove, callInfo)
	mock.lockRemove.Unlock()
	return mock.RemoveFunc()
}

// RemoveCalls gets all the calls that were made to Remove.
// Check the length with:
//     len(mockedIsolatedProject.RemoveCalls())
func (mock *IsolatedProjectMock) RemoveCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockRemove.RLock()
	calls = mock.calls.Remove
	mock.lockRemove.RUnlock()
	return calls
}

// Root calls RootFunc.
func (mock *IsolatedProjectMock) Root() string {
	if mock.RootFunc == nil {
		panic("IsolatedProjectMock.RootFunc: method is nil but IsolatedProject.Root was just called")
	}
	callInfo := struct {
	}{}
	mock.lockRoot.Lock()
	mock.calls.Root = append(mock.calls.Root, callInfo)
	mock.lockRoot.Unlock()
	return mock.RootFunc()
}

// RootCalls gets all the calls that were made to Root.
// Check the length with:
//     len(mockedIsolatedProject.RootCalls())
func (mock *IsolatedProjectMock) RootCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockRoot.RLock()
	calls = mock.calls.Root
	mock.lockRoot.RUnlock()
	return calls
}