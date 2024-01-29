// Code generated by mockery. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// GRPCServerConfigurer is an autogenerated mock type for the GRPCServerConfigurer type
type GRPCServerConfigurer struct {
	mock.Mock
}

type GRPCServerConfigurer_Expecter struct {
	mock *mock.Mock
}

func (_m *GRPCServerConfigurer) EXPECT() *GRPCServerConfigurer_Expecter {
	return &GRPCServerConfigurer_Expecter{mock: &_m.Mock}
}

// GRPCServerPort provides a mock function with given fields:
func (_m *GRPCServerConfigurer) GRPCServerPort() int {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GRPCServerPort")
	}

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// GRPCServerConfigurer_GRPCServerPort_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GRPCServerPort'
type GRPCServerConfigurer_GRPCServerPort_Call struct {
	*mock.Call
}

// GRPCServerPort is a helper method to define mock.On call
func (_e *GRPCServerConfigurer_Expecter) GRPCServerPort() *GRPCServerConfigurer_GRPCServerPort_Call {
	return &GRPCServerConfigurer_GRPCServerPort_Call{Call: _e.mock.On("GRPCServerPort")}
}

func (_c *GRPCServerConfigurer_GRPCServerPort_Call) Run(run func()) *GRPCServerConfigurer_GRPCServerPort_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *GRPCServerConfigurer_GRPCServerPort_Call) Return(_a0 int) *GRPCServerConfigurer_GRPCServerPort_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *GRPCServerConfigurer_GRPCServerPort_Call) RunAndReturn(run func() int) *GRPCServerConfigurer_GRPCServerPort_Call {
	_c.Call.Return(run)
	return _c
}

// GRPCServerReflection provides a mock function with given fields:
func (_m *GRPCServerConfigurer) GRPCServerReflection() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GRPCServerReflection")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// GRPCServerConfigurer_GRPCServerReflection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GRPCServerReflection'
type GRPCServerConfigurer_GRPCServerReflection_Call struct {
	*mock.Call
}

// GRPCServerReflection is a helper method to define mock.On call
func (_e *GRPCServerConfigurer_Expecter) GRPCServerReflection() *GRPCServerConfigurer_GRPCServerReflection_Call {
	return &GRPCServerConfigurer_GRPCServerReflection_Call{Call: _e.mock.On("GRPCServerReflection")}
}

func (_c *GRPCServerConfigurer_GRPCServerReflection_Call) Run(run func()) *GRPCServerConfigurer_GRPCServerReflection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *GRPCServerConfigurer_GRPCServerReflection_Call) Return(_a0 bool) *GRPCServerConfigurer_GRPCServerReflection_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *GRPCServerConfigurer_GRPCServerReflection_Call) RunAndReturn(run func() bool) *GRPCServerConfigurer_GRPCServerReflection_Call {
	_c.Call.Return(run)
	return _c
}

// NewGRPCServerConfigurer creates a new instance of GRPCServerConfigurer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewGRPCServerConfigurer(t interface {
	mock.TestingT
	Cleanup(func())
}) *GRPCServerConfigurer {
	mock := &GRPCServerConfigurer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
