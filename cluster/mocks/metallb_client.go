// Code generated by mockery v2.32.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	v1beta3 "github.com/akash-network/provider/cluster/types/v1beta3"

	v1beta4 "github.com/akash-network/akash-api/go/node/market/v1beta4"
)

// MetalLBClient is an autogenerated mock type for the Client type
type MetalLBClient struct {
	mock.Mock
}

type MetalLBClient_Expecter struct {
	mock *mock.Mock
}

func (_m *MetalLBClient) EXPECT() *MetalLBClient_Expecter {
	return &MetalLBClient_Expecter{mock: &_m.Mock}
}

// CreateIPPassthrough provides a mock function with given fields: ctx, directive
func (_m *MetalLBClient) CreateIPPassthrough(ctx context.Context, directive v1beta3.ClusterIPPassthroughDirective) error {
	ret := _m.Called(ctx, directive)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, v1beta3.ClusterIPPassthroughDirective) error); ok {
		r0 = rf(ctx, directive)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MetalLBClient_CreateIPPassthrough_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateIPPassthrough'
type MetalLBClient_CreateIPPassthrough_Call struct {
	*mock.Call
}

// CreateIPPassthrough is a helper method to define mock.On call
//   - ctx context.Context
//   - directive v1beta3.ClusterIPPassthroughDirective
func (_e *MetalLBClient_Expecter) CreateIPPassthrough(ctx interface{}, directive interface{}) *MetalLBClient_CreateIPPassthrough_Call {
	return &MetalLBClient_CreateIPPassthrough_Call{Call: _e.mock.On("CreateIPPassthrough", ctx, directive)}
}

func (_c *MetalLBClient_CreateIPPassthrough_Call) Run(run func(ctx context.Context, directive v1beta3.ClusterIPPassthroughDirective)) *MetalLBClient_CreateIPPassthrough_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(v1beta3.ClusterIPPassthroughDirective))
	})
	return _c
}

func (_c *MetalLBClient_CreateIPPassthrough_Call) Return(_a0 error) *MetalLBClient_CreateIPPassthrough_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MetalLBClient_CreateIPPassthrough_Call) RunAndReturn(run func(context.Context, v1beta3.ClusterIPPassthroughDirective) error) *MetalLBClient_CreateIPPassthrough_Call {
	_c.Call.Return(run)
	return _c
}

// DetectPoolChanges provides a mock function with given fields: ctx
func (_m *MetalLBClient) DetectPoolChanges(ctx context.Context) (<-chan struct{}, error) {
	ret := _m.Called(ctx)

	var r0 <-chan struct{}
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (<-chan struct{}, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) <-chan struct{}); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan struct{})
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MetalLBClient_DetectPoolChanges_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DetectPoolChanges'
type MetalLBClient_DetectPoolChanges_Call struct {
	*mock.Call
}

// DetectPoolChanges is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MetalLBClient_Expecter) DetectPoolChanges(ctx interface{}) *MetalLBClient_DetectPoolChanges_Call {
	return &MetalLBClient_DetectPoolChanges_Call{Call: _e.mock.On("DetectPoolChanges", ctx)}
}

func (_c *MetalLBClient_DetectPoolChanges_Call) Run(run func(ctx context.Context)) *MetalLBClient_DetectPoolChanges_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MetalLBClient_DetectPoolChanges_Call) Return(_a0 <-chan struct{}, _a1 error) *MetalLBClient_DetectPoolChanges_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MetalLBClient_DetectPoolChanges_Call) RunAndReturn(run func(context.Context) (<-chan struct{}, error)) *MetalLBClient_DetectPoolChanges_Call {
	_c.Call.Return(run)
	return _c
}

// GetIPAddressStatusForLease provides a mock function with given fields: ctx, leaseID
func (_m *MetalLBClient) GetIPAddressStatusForLease(ctx context.Context, leaseID v1beta4.LeaseID) ([]v1beta3.IPLeaseState, error) {
	ret := _m.Called(ctx, leaseID)

	var r0 []v1beta3.IPLeaseState
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, v1beta4.LeaseID) ([]v1beta3.IPLeaseState, error)); ok {
		return rf(ctx, leaseID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, v1beta4.LeaseID) []v1beta3.IPLeaseState); ok {
		r0 = rf(ctx, leaseID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]v1beta3.IPLeaseState)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, v1beta4.LeaseID) error); ok {
		r1 = rf(ctx, leaseID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MetalLBClient_GetIPAddressStatusForLease_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetIPAddressStatusForLease'
type MetalLBClient_GetIPAddressStatusForLease_Call struct {
	*mock.Call
}

// GetIPAddressStatusForLease is a helper method to define mock.On call
//   - ctx context.Context
//   - leaseID v1beta4.LeaseID
func (_e *MetalLBClient_Expecter) GetIPAddressStatusForLease(ctx interface{}, leaseID interface{}) *MetalLBClient_GetIPAddressStatusForLease_Call {
	return &MetalLBClient_GetIPAddressStatusForLease_Call{Call: _e.mock.On("GetIPAddressStatusForLease", ctx, leaseID)}
}

func (_c *MetalLBClient_GetIPAddressStatusForLease_Call) Run(run func(ctx context.Context, leaseID v1beta4.LeaseID)) *MetalLBClient_GetIPAddressStatusForLease_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(v1beta4.LeaseID))
	})
	return _c
}

func (_c *MetalLBClient_GetIPAddressStatusForLease_Call) Return(_a0 []v1beta3.IPLeaseState, _a1 error) *MetalLBClient_GetIPAddressStatusForLease_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MetalLBClient_GetIPAddressStatusForLease_Call) RunAndReturn(run func(context.Context, v1beta4.LeaseID) ([]v1beta3.IPLeaseState, error)) *MetalLBClient_GetIPAddressStatusForLease_Call {
	_c.Call.Return(run)
	return _c
}

// GetIPAddressUsage provides a mock function with given fields: ctx
func (_m *MetalLBClient) GetIPAddressUsage(ctx context.Context) (uint, uint, error) {
	ret := _m.Called(ctx)

	var r0 uint
	var r1 uint
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context) (uint, uint, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) uint); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(uint)
	}

	if rf, ok := ret.Get(1).(func(context.Context) uint); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Get(1).(uint)
	}

	if rf, ok := ret.Get(2).(func(context.Context) error); ok {
		r2 = rf(ctx)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MetalLBClient_GetIPAddressUsage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetIPAddressUsage'
type MetalLBClient_GetIPAddressUsage_Call struct {
	*mock.Call
}

// GetIPAddressUsage is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MetalLBClient_Expecter) GetIPAddressUsage(ctx interface{}) *MetalLBClient_GetIPAddressUsage_Call {
	return &MetalLBClient_GetIPAddressUsage_Call{Call: _e.mock.On("GetIPAddressUsage", ctx)}
}

func (_c *MetalLBClient_GetIPAddressUsage_Call) Run(run func(ctx context.Context)) *MetalLBClient_GetIPAddressUsage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MetalLBClient_GetIPAddressUsage_Call) Return(_a0 uint, _a1 uint, _a2 error) *MetalLBClient_GetIPAddressUsage_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MetalLBClient_GetIPAddressUsage_Call) RunAndReturn(run func(context.Context) (uint, uint, error)) *MetalLBClient_GetIPAddressUsage_Call {
	_c.Call.Return(run)
	return _c
}

// GetIPPassthroughs provides a mock function with given fields: ctx
func (_m *MetalLBClient) GetIPPassthroughs(ctx context.Context) ([]v1beta3.IPPassthrough, error) {
	ret := _m.Called(ctx)

	var r0 []v1beta3.IPPassthrough
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]v1beta3.IPPassthrough, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []v1beta3.IPPassthrough); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]v1beta3.IPPassthrough)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MetalLBClient_GetIPPassthroughs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetIPPassthroughs'
type MetalLBClient_GetIPPassthroughs_Call struct {
	*mock.Call
}

// GetIPPassthroughs is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MetalLBClient_Expecter) GetIPPassthroughs(ctx interface{}) *MetalLBClient_GetIPPassthroughs_Call {
	return &MetalLBClient_GetIPPassthroughs_Call{Call: _e.mock.On("GetIPPassthroughs", ctx)}
}

func (_c *MetalLBClient_GetIPPassthroughs_Call) Run(run func(ctx context.Context)) *MetalLBClient_GetIPPassthroughs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MetalLBClient_GetIPPassthroughs_Call) Return(_a0 []v1beta3.IPPassthrough, _a1 error) *MetalLBClient_GetIPPassthroughs_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MetalLBClient_GetIPPassthroughs_Call) RunAndReturn(run func(context.Context) ([]v1beta3.IPPassthrough, error)) *MetalLBClient_GetIPPassthroughs_Call {
	_c.Call.Return(run)
	return _c
}

// PurgeIPPassthrough provides a mock function with given fields: ctx, directive
func (_m *MetalLBClient) PurgeIPPassthrough(ctx context.Context, directive v1beta3.ClusterIPPassthroughDirective) error {
	ret := _m.Called(ctx, directive)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, v1beta3.ClusterIPPassthroughDirective) error); ok {
		r0 = rf(ctx, directive)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MetalLBClient_PurgeIPPassthrough_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PurgeIPPassthrough'
type MetalLBClient_PurgeIPPassthrough_Call struct {
	*mock.Call
}

// PurgeIPPassthrough is a helper method to define mock.On call
//   - ctx context.Context
//   - directive v1beta3.ClusterIPPassthroughDirective
func (_e *MetalLBClient_Expecter) PurgeIPPassthrough(ctx interface{}, directive interface{}) *MetalLBClient_PurgeIPPassthrough_Call {
	return &MetalLBClient_PurgeIPPassthrough_Call{Call: _e.mock.On("PurgeIPPassthrough", ctx, directive)}
}

func (_c *MetalLBClient_PurgeIPPassthrough_Call) Run(run func(ctx context.Context, directive v1beta3.ClusterIPPassthroughDirective)) *MetalLBClient_PurgeIPPassthrough_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(v1beta3.ClusterIPPassthroughDirective))
	})
	return _c
}

func (_c *MetalLBClient_PurgeIPPassthrough_Call) Return(_a0 error) *MetalLBClient_PurgeIPPassthrough_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MetalLBClient_PurgeIPPassthrough_Call) RunAndReturn(run func(context.Context, v1beta3.ClusterIPPassthroughDirective) error) *MetalLBClient_PurgeIPPassthrough_Call {
	_c.Call.Return(run)
	return _c
}

// Stop provides a mock function with given fields:
func (_m *MetalLBClient) Stop() {
	_m.Called()
}

// MetalLBClient_Stop_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Stop'
type MetalLBClient_Stop_Call struct {
	*mock.Call
}

// Stop is a helper method to define mock.On call
func (_e *MetalLBClient_Expecter) Stop() *MetalLBClient_Stop_Call {
	return &MetalLBClient_Stop_Call{Call: _e.mock.On("Stop")}
}

func (_c *MetalLBClient_Stop_Call) Run(run func()) *MetalLBClient_Stop_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MetalLBClient_Stop_Call) Return() *MetalLBClient_Stop_Call {
	_c.Call.Return()
	return _c
}

func (_c *MetalLBClient_Stop_Call) RunAndReturn(run func()) *MetalLBClient_Stop_Call {
	_c.Call.Return(run)
	return _c
}

// NewMetalLBClient creates a new instance of MetalLBClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMetalLBClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *MetalLBClient {
	mock := &MetalLBClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
