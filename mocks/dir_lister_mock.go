// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mocks

import (
	"github.com/fwojciec/packager"
	"sync"
)

// Ensure, that DirListerMock does implement packager.DirLister.
// If this is not the case, regenerate this file with moq.
var _ packager.DirLister = &DirListerMock{}

// DirListerMock is a mock implementation of packager.DirLister.
//
// 	func TestSomethingThatUsesDirLister(t *testing.T) {
//
// 		// make and configure a mocked packager.DirLister
// 		mockedDirLister := &DirListerMock{
// 			ListDirFunc: func(target string, excl packager.Excluder) ([]string, error) {
// 				panic("mock out the ListDir method")
// 			},
// 		}
//
// 		// use mockedDirLister in code that requires packager.DirLister
// 		// and then make assertions.
//
// 	}
type DirListerMock struct {
	// ListDirFunc mocks the ListDir method.
	ListDirFunc func(target string, excl packager.Excluder) ([]string, error)

	// calls tracks calls to the methods.
	calls struct {
		// ListDir holds details about calls to the ListDir method.
		ListDir []struct {
			// Target is the target argument value.
			Target string
			// Excl is the excl argument value.
			Excl packager.Excluder
		}
	}
	lockListDir sync.RWMutex
}

// ListDir calls ListDirFunc.
func (mock *DirListerMock) ListDir(target string, excl packager.Excluder) ([]string, error) {
	if mock.ListDirFunc == nil {
		panic("DirListerMock.ListDirFunc: method is nil but DirLister.ListDir was just called")
	}
	callInfo := struct {
		Target string
		Excl   packager.Excluder
	}{
		Target: target,
		Excl:   excl,
	}
	mock.lockListDir.Lock()
	mock.calls.ListDir = append(mock.calls.ListDir, callInfo)
	mock.lockListDir.Unlock()
	return mock.ListDirFunc(target, excl)
}

// ListDirCalls gets all the calls that were made to ListDir.
// Check the length with:
//     len(mockedDirLister.ListDirCalls())
func (mock *DirListerMock) ListDirCalls() []struct {
	Target string
	Excl   packager.Excluder
} {
	var calls []struct {
		Target string
		Excl   packager.Excluder
	}
	mock.lockListDir.RLock()
	calls = mock.calls.ListDir
	mock.lockListDir.RUnlock()
	return calls
}
