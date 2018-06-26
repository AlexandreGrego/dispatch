// Code generated by mockery v1.0.0

package mocks

import context "context"
import mock "github.com/stretchr/testify/mock"
import v1 "github.com/vmware/dispatch/pkg/api/v1"

// ServicesClient is an autogenerated mock type for the ServicesClient type
type ServicesClient struct {
	mock.Mock
}

// CreateServiceInstance provides a mock function with given fields: ctx, organizationID, serviceInstance
func (_m *ServicesClient) CreateServiceInstance(ctx context.Context, organizationID string, serviceInstance *v1.ServiceInstance) (*v1.ServiceInstance, error) {
	ret := _m.Called(ctx, organizationID, serviceInstance)

	var r0 *v1.ServiceInstance
	if rf, ok := ret.Get(0).(func(context.Context, string, *v1.ServiceInstance) *v1.ServiceInstance); ok {
		r0 = rf(ctx, organizationID, serviceInstance)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.ServiceInstance)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, *v1.ServiceInstance) error); ok {
		r1 = rf(ctx, organizationID, serviceInstance)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteServiceInstance provides a mock function with given fields: ctx, organizationID, serviceInstanceName
func (_m *ServicesClient) DeleteServiceInstance(ctx context.Context, organizationID string, serviceInstanceName string) error {
	ret := _m.Called(ctx, organizationID, serviceInstanceName)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, organizationID, serviceInstanceName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetServiceClass provides a mock function with given fields: ctx, organizationID, serviceClassName
func (_m *ServicesClient) GetServiceClass(ctx context.Context, organizationID string, serviceClassName string) (*v1.ServiceClass, error) {
	ret := _m.Called(ctx, organizationID, serviceClassName)

	var r0 *v1.ServiceClass
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *v1.ServiceClass); ok {
		r0 = rf(ctx, organizationID, serviceClassName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.ServiceClass)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, organizationID, serviceClassName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetServiceInstance provides a mock function with given fields: ctx, organizationID, serviceInstanceName
func (_m *ServicesClient) GetServiceInstance(ctx context.Context, organizationID string, serviceInstanceName string) (*v1.ServiceInstance, error) {
	ret := _m.Called(ctx, organizationID, serviceInstanceName)

	var r0 *v1.ServiceInstance
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *v1.ServiceInstance); ok {
		r0 = rf(ctx, organizationID, serviceInstanceName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.ServiceInstance)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, organizationID, serviceInstanceName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListServiceClasses provides a mock function with given fields: ctx, organizationID
func (_m *ServicesClient) ListServiceClasses(ctx context.Context, organizationID string) ([]v1.ServiceClass, error) {
	ret := _m.Called(ctx, organizationID)

	var r0 []v1.ServiceClass
	if rf, ok := ret.Get(0).(func(context.Context, string) []v1.ServiceClass); ok {
		r0 = rf(ctx, organizationID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]v1.ServiceClass)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, organizationID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListServiceInstances provides a mock function with given fields: ctx, organizationID
func (_m *ServicesClient) ListServiceInstances(ctx context.Context, organizationID string) ([]v1.ServiceInstance, error) {
	ret := _m.Called(ctx, organizationID)

	var r0 []v1.ServiceInstance
	if rf, ok := ret.Get(0).(func(context.Context, string) []v1.ServiceInstance); ok {
		r0 = rf(ctx, organizationID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]v1.ServiceInstance)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, organizationID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
