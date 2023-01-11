// Code generated by mockery 2.12.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	testing "testing"

	typesv1beta2 "github.com/akash-network/node/x/market/types/v1beta2"

	v1beta2 "github.com/akash-network/provider/cluster/types/v1beta2"
)

// MetalLBClient is an autogenerated mock type for the Client type
type MetalLBClient struct {
	mock.Mock
}

// CreateIPPassthrough provides a mock function with given fields: ctx, directive
func (_m *MetalLBClient) CreateIPPassthrough(ctx context.Context, directive v1beta2.ClusterIPPassthroughDirective) error {
	ret := _m.Called(ctx, directive)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, v1beta2.ClusterIPPassthroughDirective) error); ok {
		r0 = rf(ctx, directive)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DetectPoolChanges provides a mock function with given fields: ctx
func (_m *MetalLBClient) DetectPoolChanges(ctx context.Context) (<-chan struct{}, error) {
	ret := _m.Called(ctx)

	var r0 <-chan struct{}
	if rf, ok := ret.Get(0).(func(context.Context) <-chan struct{}); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan struct{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetIPAddressStatusForLease provides a mock function with given fields: ctx, leaseID
func (_m *MetalLBClient) GetIPAddressStatusForLease(ctx context.Context, leaseID typesv1beta2.LeaseID) ([]v1beta2.IPLeaseState, error) {
	ret := _m.Called(ctx, leaseID)

	var r0 []v1beta2.IPLeaseState
	if rf, ok := ret.Get(0).(func(context.Context, typesv1beta2.LeaseID) []v1beta2.IPLeaseState); ok {
		r0 = rf(ctx, leaseID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]v1beta2.IPLeaseState)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, typesv1beta2.LeaseID) error); ok {
		r1 = rf(ctx, leaseID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetIPAddressUsage provides a mock function with given fields: ctx
func (_m *MetalLBClient) GetIPAddressUsage(ctx context.Context) (uint, uint, error) {
	ret := _m.Called(ctx)

	var r0 uint
	if rf, ok := ret.Get(0).(func(context.Context) uint); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(uint)
	}

	var r1 uint
	if rf, ok := ret.Get(1).(func(context.Context) uint); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Get(1).(uint)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context) error); ok {
		r2 = rf(ctx)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetIPPassthroughs provides a mock function with given fields: ctx
func (_m *MetalLBClient) GetIPPassthroughs(ctx context.Context) ([]v1beta2.IPPassthrough, error) {
	ret := _m.Called(ctx)

	var r0 []v1beta2.IPPassthrough
	if rf, ok := ret.Get(0).(func(context.Context) []v1beta2.IPPassthrough); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]v1beta2.IPPassthrough)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PurgeIPPassthrough provides a mock function with given fields: ctx, directive
func (_m *MetalLBClient) PurgeIPPassthrough(ctx context.Context, directive v1beta2.ClusterIPPassthroughDirective) error {
	ret := _m.Called(ctx, directive)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, v1beta2.ClusterIPPassthroughDirective) error); ok {
		r0 = rf(ctx, directive)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Stop provides a mock function with given fields:
func (_m *MetalLBClient) Stop() {
	_m.Called()
}

// NewMetalLBClient creates a new instance of MetalLBClient. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewMetalLBClient(t testing.TB) *MetalLBClient {
	mock := &MetalLBClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
