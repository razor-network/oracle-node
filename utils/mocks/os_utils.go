// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	fs "io/fs"
	os "os"

	mock "github.com/stretchr/testify/mock"
)

// OSUtils is an autogenerated mock type for the OSUtils type
type OSUtils struct {
	mock.Mock
}

// Open provides a mock function with given fields: name
func (_m *OSUtils) Open(name string) (*os.File, error) {
	ret := _m.Called(name)

	var r0 *os.File
	if rf, ok := ret.Get(0).(func(string) *os.File); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*os.File)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OpenFile provides a mock function with given fields: name, flag, perm
func (_m *OSUtils) OpenFile(name string, flag int, perm fs.FileMode) (*os.File, error) {
	ret := _m.Called(name, flag, perm)

	var r0 *os.File
	if rf, ok := ret.Get(0).(func(string, int, fs.FileMode) *os.File); ok {
		r0 = rf(name, flag, perm)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*os.File)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, int, fs.FileMode) error); ok {
		r1 = rf(name, flag, perm)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WriteFile provides a mock function with given fields: name, data, perm
func (_m *OSUtils) WriteFile(name string, data []byte, perm fs.FileMode) error {
	ret := _m.Called(name, data, perm)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, []byte, fs.FileMode) error); ok {
		r0 = rf(name, data, perm)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
