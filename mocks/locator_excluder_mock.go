// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mocks

import (
	"github.com/fwojciec/packager"
	"sync"
)

// Ensure, that LocatorExcluderMock does implement packager.LocatorExcluder.
// If this is not the case, regenerate this file with moq.
var _ packager.LocatorExcluder = &LocatorExcluderMock{}

// LocatorExcluderMock is a mock implementation of packager.LocatorExcluder.
//
// 	func TestSomethingThatUsesLocatorExcluder(t *testing.T) {
//
// 		// make and configure a mocked packager.LocatorExcluder
// 		mockedLocatorExcluder := &LocatorExcluderMock{
// 			ExcludeFunc: func(path string) bool {
// 				panic("mock out the Exclude method")
// 			},
// 			LocationFunc: func() string {
// 				panic("mock out the Location method")
// 			},
// 		}
//
// 		// use mockedLocatorExcluder in code that requires packager.LocatorExcluder
// 		// and then make assertions.
//
// 	}
type LocatorExcluderMock struct {
	// ExcludeFunc mocks the Exclude method.
	ExcludeFunc func(path string) bool

	// LocationFunc mocks the Location method.
	LocationFunc func() string

	// calls tracks calls to the methods.
	calls struct {
		// Exclude holds details about calls to the Exclude method.
		Exclude []struct {
			// Path is the path argument value.
			Path string
		}
		// Location holds details about calls to the Location method.
		Location []struct {
		}
	}
	lockExclude  sync.RWMutex
	lockLocation sync.RWMutex
}

// Exclude calls ExcludeFunc.
func (mock *LocatorExcluderMock) Exclude(path string) bool {
	if mock.ExcludeFunc == nil {
		panic("LocatorExcluderMock.ExcludeFunc: method is nil but LocatorExcluder.Exclude was just called")
	}
	callInfo := struct {
		Path string
	}{
		Path: path,
	}
	mock.lockExclude.Lock()
	mock.calls.Exclude = append(mock.calls.Exclude, callInfo)
	mock.lockExclude.Unlock()
	return mock.ExcludeFunc(path)
}

// ExcludeCalls gets all the calls that were made to Exclude.
// Check the length with:
//     len(mockedLocatorExcluder.ExcludeCalls())
func (mock *LocatorExcluderMock) ExcludeCalls() []struct {
	Path string
} {
	var calls []struct {
		Path string
	}
	mock.lockExclude.RLock()
	calls = mock.calls.Exclude
	mock.lockExclude.RUnlock()
	return calls
}

// Location calls LocationFunc.
func (mock *LocatorExcluderMock) Location() string {
	if mock.LocationFunc == nil {
		panic("LocatorExcluderMock.LocationFunc: method is nil but LocatorExcluder.Location was just called")
	}
	callInfo := struct {
	}{}
	mock.lockLocation.Lock()
	mock.calls.Location = append(mock.calls.Location, callInfo)
	mock.lockLocation.Unlock()
	return mock.LocationFunc()
}

// LocationCalls gets all the calls that were made to Location.
// Check the length with:
//     len(mockedLocatorExcluder.LocationCalls())
func (mock *LocatorExcluderMock) LocationCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockLocation.RLock()
	calls = mock.calls.Location
	mock.lockLocation.RUnlock()
	return calls
}