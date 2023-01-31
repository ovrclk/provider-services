// Code generated by mockery 2.12.1. DO NOT EDIT.

package mocks

import (
	context "context"
	testing "testing"

	mock "github.com/stretchr/testify/mock"

	types "github.com/cosmos/cosmos-sdk/types"

	v1beta3 "github.com/akash-network/akash-api/go/node/market/v1beta3"
)

// HostnameServiceClient is an autogenerated mock type for the HostnameServiceClient type
type HostnameServiceClient struct {
	mock.Mock
}

// CanReserveHostnames provides a mock function with given fields: hostnames, ownerAddr
func (_m *HostnameServiceClient) CanReserveHostnames(hostnames []string, ownerAddr types.Address) error {
	ret := _m.Called(hostnames, ownerAddr)

	var r0 error
	if rf, ok := ret.Get(0).(func([]string, types.Address) error); ok {
		r0 = rf(hostnames, ownerAddr)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PrepareHostnamesForTransfer provides a mock function with given fields: ctx, hostnames, leaseID
func (_m *HostnameServiceClient) PrepareHostnamesForTransfer(ctx context.Context, hostnames []string, leaseID v1beta3.LeaseID) error {
	ret := _m.Called(ctx, hostnames, leaseID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []string, v1beta3.LeaseID) error); ok {
		r0 = rf(ctx, hostnames, leaseID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ReleaseHostnames provides a mock function with given fields: leaseID
func (_m *HostnameServiceClient) ReleaseHostnames(leaseID v1beta3.LeaseID) error {
	ret := _m.Called(leaseID)

	var r0 error
	if rf, ok := ret.Get(0).(func(v1beta3.LeaseID) error); ok {
		r0 = rf(leaseID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ReserveHostnames provides a mock function with given fields: ctx, hostnames, leaseID
func (_m *HostnameServiceClient) ReserveHostnames(ctx context.Context, hostnames []string, leaseID v1beta3.LeaseID) ([]string, error) {
	ret := _m.Called(ctx, hostnames, leaseID)

	var r0 []string
	if rf, ok := ret.Get(0).(func(context.Context, []string, v1beta3.LeaseID) []string); ok {
		r0 = rf(ctx, hostnames, leaseID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []string, v1beta3.LeaseID) error); ok {
		r1 = rf(ctx, hostnames, leaseID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewHostnameServiceClient creates a new instance of HostnameServiceClient. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewHostnameServiceClient(t testing.TB) *HostnameServiceClient {
	mock := &HostnameServiceClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
