// Code generated by mockery. DO NOT EDIT.

package telemetry

import (
	pmmv1 "github.com/percona-platform/saas/gen/telemetry/events/pmm"
	mock "github.com/stretchr/testify/mock"

	serverv1 "github.com/percona/pmm/api/server/v1"
)

// mockDistributionUtilService is an autogenerated mock type for the distributionUtilService type
type mockDistributionUtilService struct {
	mock.Mock
}

// GetDistributionMethodAndOS provides a mock function with given fields:
func (_m *mockDistributionUtilService) GetDistributionMethodAndOS() (serverv1.DistributionMethod, pmmv1.DistributionMethod, string) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetDistributionMethodAndOS")
	}

	var r0 serverv1.DistributionMethod
	var r1 pmmv1.DistributionMethod
	var r2 string
	if rf, ok := ret.Get(0).(func() (serverv1.DistributionMethod, pmmv1.DistributionMethod, string)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() serverv1.DistributionMethod); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(serverv1.DistributionMethod)
	}

	if rf, ok := ret.Get(1).(func() pmmv1.DistributionMethod); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(1).(pmmv1.DistributionMethod)
	}

	if rf, ok := ret.Get(2).(func() string); ok {
		r2 = rf()
	} else {
		r2 = ret.Get(2).(string)
	}

	return r0, r1, r2
}

// newMockDistributionUtilService creates a new instance of mockDistributionUtilService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockDistributionUtilService(t interface {
	mock.TestingT
	Cleanup(func())
},
) *mockDistributionUtilService {
	mock := &mockDistributionUtilService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
