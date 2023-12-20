// Code generated by mockery v2.39.0. DO NOT EDIT.

package mocks

import (
	entity "wisdom/internal/domain/entity"

	mock "github.com/stretchr/testify/mock"
)

// ServerService is an autogenerated mock type for the ServerService type
type ServerService struct {
	mock.Mock
}

// SendMessage provides a mock function with given fields: message
func (_m *ServerService) SendMessage(message entity.ServerMessage) error {
	ret := _m.Called(message)

	if len(ret) == 0 {
		panic("no return value specified for SendMessage")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(entity.ServerMessage) error); ok {
		r0 = rf(message)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ServeAndListen provides a mock function with given fields: readChannel, errorChannel
func (_m *ServerService) ServeAndListen(readChannel chan entity.ServerMessage, errorChannel chan error) error {
	ret := _m.Called(readChannel, errorChannel)

	if len(ret) == 0 {
		panic("no return value specified for ServeAndListen")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(chan entity.ServerMessage, chan error) error); ok {
		r0 = rf(readChannel, errorChannel)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewServerService creates a new instance of ServerService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewServerService(t interface {
	mock.TestingT
	Cleanup(func())
}) *ServerService {
	mock := &ServerService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}