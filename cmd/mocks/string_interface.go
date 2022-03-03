// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// StringInterface is an autogenerated mock type for the StringInterface type
type StringInterface struct {
	mock.Mock
}

// ParseBool provides a mock function with given fields: str
func (_m *StringInterface) ParseBool(str string) (bool, error) {
	ret := _m.Called(str)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(str)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(str)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}