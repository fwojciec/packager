// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mocks

import (
	"github.com/fwojciec/packager"
	"sync"
)

// Ensure, that IsolatorMock does implement packager.Isolator.
// If this is not the case, regenerate this file with moq.
var _ packager.Isolator = &IsolatorMock{}

// IsolatorMock is a mock implementation of packager.Isolator.
//
// 	func TestSomethingThatUsesIsolator(t *testing.T) {
//
// 		// make and configure a mocked packager.Isolator
// 		mockedIsolator := &IsolatorMock{
// 			IsolateFunc: func(project packager.Project) (packager.IsolatedProject, error) {
// 				panic("mock out the Isolate method")
// 			},
// 		}
//
// 		// use mockedIsolator in code that requires packager.Isolator
// 		// and then make assertions.
//
// 	}
type IsolatorMock struct {
	// IsolateFunc mocks the Isolate method.
	IsolateFunc func(project packager.Project) (packager.IsolatedProject, error)

	// calls tracks calls to the methods.
	calls struct {
		// Isolate holds details about calls to the Isolate method.
		Isolate []struct {
			// Project is the project argument value.
			Project packager.Project
		}
	}
	lockIsolate sync.RWMutex
}

// Isolate calls IsolateFunc.
func (mock *IsolatorMock) Isolate(project packager.Project) (packager.IsolatedProject, error) {
	if mock.IsolateFunc == nil {
		panic("IsolatorMock.IsolateFunc: method is nil but Isolator.Isolate was just called")
	}
	callInfo := struct {
		Project packager.Project
	}{
		Project: project,
	}
	mock.lockIsolate.Lock()
	mock.calls.Isolate = append(mock.calls.Isolate, callInfo)
	mock.lockIsolate.Unlock()
	return mock.IsolateFunc(project)
}

// IsolateCalls gets all the calls that were made to Isolate.
// Check the length with:
//     len(mockedIsolator.IsolateCalls())
func (mock *IsolatorMock) IsolateCalls() []struct {
	Project packager.Project
} {
	var calls []struct {
		Project packager.Project
	}
	mock.lockIsolate.RLock()
	calls = mock.calls.Isolate
	mock.lockIsolate.RUnlock()
	return calls
}
