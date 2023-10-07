// Code generated by mockery v2.35.1. DO NOT EDIT.

package client

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	agentpb "github.com/percona/pmm/api/agentpb"
)

// mockConnectionChecker is an autogenerated mock type for the connectionChecker type
type mockConnectionChecker struct {
	mock.Mock
}

// Check provides a mock function with given fields: ctx, req, id
func (_m *mockConnectionChecker) Check(ctx context.Context, req *agentpb.CheckConnectionRequest, id uint32) *agentpb.CheckConnectionResponse {
	ret := _m.Called(ctx, req, id)

	var r0 *agentpb.CheckConnectionResponse
	if rf, ok := ret.Get(0).(func(context.Context, *agentpb.CheckConnectionRequest, uint32) *agentpb.CheckConnectionResponse); ok {
		r0 = rf(ctx, req, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*agentpb.CheckConnectionResponse)
		}
	}

	return r0
}

// newMockConnectionChecker creates a new instance of mockConnectionChecker. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockConnectionChecker(t interface {
	mock.TestingT
	Cleanup(func())
},
) *mockConnectionChecker {
	mock := &mockConnectionChecker{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
