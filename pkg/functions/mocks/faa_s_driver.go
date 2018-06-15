// Code generated by mockery v1.0.0

// CLOSE THIS FILE AS QUICKLY AS POSSIBLE

package mocks

import context "context"
import functions "github.com/vmware/dispatch/pkg/functions"
import mock "github.com/stretchr/testify/mock"

// FaaSDriver is an autogenerated mock type for the FaaSDriver type
type FaaSDriver struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, f
func (_m *FaaSDriver) Create(ctx context.Context, f *functions.Function) error {
	ret := _m.Called(ctx, f)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *functions.Function) error); ok {
		r0 = rf(ctx, f)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: ctx, f
func (_m *FaaSDriver) Delete(ctx context.Context, f *functions.Function) error {
	ret := _m.Called(ctx, f)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *functions.Function) error); ok {
		r0 = rf(ctx, f)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetRunnable provides a mock function with given fields: e
func (_m *FaaSDriver) GetRunnable(e *functions.FunctionExecution) functions.Runnable {
	ret := _m.Called(e)

	var r0 functions.Runnable
	if rf, ok := ret.Get(0).(func(*functions.FunctionExecution) functions.Runnable); ok {
		r0 = rf(e)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(functions.Runnable)
		}
	}

	return r0
}