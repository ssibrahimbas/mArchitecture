// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	events "clean-boilerplate/shared/events"

	mock "github.com/stretchr/testify/mock"
)

// Engine is an autogenerated mock type for the Engine type
type Engine struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *Engine) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Marshal provides a mock function with given fields: data
func (_m *Engine) Marshal(data interface{}) ([]byte, error) {
	ret := _m.Called(data)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(interface{}) []byte); ok {
		r0 = rf(data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Open provides a mock function with given fields:
func (_m *Engine) Open() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Publish provides a mock function with given fields: event, data
func (_m *Engine) Publish(event string, data interface{}) error {
	ret := _m.Called(event, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, interface{}) error); ok {
		r0 = rf(event, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Subscribe provides a mock function with given fields: event, handler
func (_m *Engine) Subscribe(event string, handler events.Handler) error {
	ret := _m.Called(event, handler)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, events.Handler) error); ok {
		r0 = rf(event, handler)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Unmarshal provides a mock function with given fields: data, v
func (_m *Engine) Unmarshal(data []byte, v interface{}) error {
	ret := _m.Called(data, v)

	var r0 error
	if rf, ok := ret.Get(0).(func([]byte, interface{}) error); ok {
		r0 = rf(data, v)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Unsubscribe provides a mock function with given fields: event, handler
func (_m *Engine) Unsubscribe(event string, handler events.Handler) error {
	ret := _m.Called(event, handler)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, events.Handler) error); ok {
		r0 = rf(event, handler)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewEngine interface {
	mock.TestingT
	Cleanup(func())
}

// NewEngine creates a new instance of Engine. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewEngine(t mockConstructorTestingTNewEngine) *Engine {
	mock := &Engine{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}