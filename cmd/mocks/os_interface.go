// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// OSInterface is an autogenerated mock type for the OSInterface type
type OSInterface struct {
	mock.Mock
}

// Exit provides a mock function with given fields: code
func (_m *OSInterface) Exit(code int) {
	_m.Called(code)
}

// NewOSInterface creates a new instance of OSInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewOSInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *OSInterface {
	mock := &OSInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
