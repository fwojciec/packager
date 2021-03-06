// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mocks

import (
	"github.com/fwojciec/packager"
	"sync"
)

// Ensure, that LocatorMock does implement packager.Locator.
// If this is not the case, regenerate this file with moq.
var _ packager.Locator = &LocatorMock{}

// LocatorMock is a mock implementation of packager.Locator.
//
// 	func TestSomethingThatUsesLocator(t *testing.T) {
//
// 		// make and configure a mocked packager.Locator
// 		mockedLocator := &LocatorMock{
// 			LocationFunc: func() string {
// 				panic("mock out the Location method")
// 			},
// 		}
//
// 		// use mockedLocator in code that requires packager.Locator
// 		// and then make assertions.
//
// 	}
type LocatorMock struct {
	// LocationFunc mocks the Location method.
	LocationFunc func() string

	// calls tracks calls to the methods.
	calls struct {
		// Location holds details about calls to the Location method.
		Location []struct {
		}
	}
	lockLocation sync.RWMutex
}

// Location calls LocationFunc.
func (mock *LocatorMock) Location() string {
	if mock.LocationFunc == nil {
		panic("LocatorMock.LocationFunc: method is nil but Locator.Location was just called")
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
//     len(mockedLocator.LocationCalls())
func (mock *LocatorMock) LocationCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockLocation.RLock()
	calls = mock.calls.Location
	mock.lockLocation.RUnlock()
	return calls
}
