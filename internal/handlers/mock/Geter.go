// Code generated by mockery v2.39.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	model "github.com/mmfshirokan/apod/internal/model"
)

// Geter is an autogenerated mock type for the Geter type
type Geter struct {
	mock.Mock
}

type Geter_Expecter struct {
	mock *mock.Mock
}

func (_m *Geter) EXPECT() *Geter_Expecter {
	return &Geter_Expecter{mock: &_m.Mock}
}

// Get provides a mock function with given fields: ctx, date
func (_m *Geter) Get(ctx context.Context, date string) (model.ImageInfo, error) {
	ret := _m.Called(ctx, date)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 model.ImageInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (model.ImageInfo, error)); ok {
		return rf(ctx, date)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) model.ImageInfo); ok {
		r0 = rf(ctx, date)
	} else {
		r0 = ret.Get(0).(model.ImageInfo)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, date)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Geter_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type Geter_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - date string
func (_e *Geter_Expecter) Get(ctx interface{}, date interface{}) *Geter_Get_Call {
	return &Geter_Get_Call{Call: _e.mock.On("Get", ctx, date)}
}

func (_c *Geter_Get_Call) Run(run func(ctx context.Context, date string)) *Geter_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *Geter_Get_Call) Return(_a0 model.ImageInfo, _a1 error) *Geter_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Geter_Get_Call) RunAndReturn(run func(context.Context, string) (model.ImageInfo, error)) *Geter_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetAll provides a mock function with given fields: ctx
func (_m *Geter) GetAll(ctx context.Context) ([]model.ImageInfo, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetAll")
	}

	var r0 []model.ImageInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]model.ImageInfo, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []model.ImageInfo); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.ImageInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Geter_GetAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAll'
type Geter_GetAll_Call struct {
	*mock.Call
}

// GetAll is a helper method to define mock.On call
//   - ctx context.Context
func (_e *Geter_Expecter) GetAll(ctx interface{}) *Geter_GetAll_Call {
	return &Geter_GetAll_Call{Call: _e.mock.On("GetAll", ctx)}
}

func (_c *Geter_GetAll_Call) Run(run func(ctx context.Context)) *Geter_GetAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Geter_GetAll_Call) Return(_a0 []model.ImageInfo, _a1 error) *Geter_GetAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Geter_GetAll_Call) RunAndReturn(run func(context.Context) ([]model.ImageInfo, error)) *Geter_GetAll_Call {
	_c.Call.Return(run)
	return _c
}

// NewGeter creates a new instance of Geter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewGeter(t interface {
	mock.TestingT
	Cleanup(func())
}) *Geter {
	mock := &Geter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
