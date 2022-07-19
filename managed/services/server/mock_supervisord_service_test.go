// Code generated by mockery v1.0.0. DO NOT EDIT.

package server

import (
	context "context"
	time "time"

	mock "github.com/stretchr/testify/mock"

	models "github.com/percona/pmm/managed/models"
	version "github.com/percona/pmm/version"
)

// mockSupervisordService is an autogenerated mock type for the supervisordService type
type mockSupervisordService struct {
	mock.Mock
}

// ForceCheckUpdates provides a mock function with given fields: ctx
func (_m *mockSupervisordService) ForceCheckUpdates(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InstalledPMMVersion provides a mock function with given fields: ctx
func (_m *mockSupervisordService) InstalledPMMVersion(ctx context.Context) *version.PackageInfo {
	ret := _m.Called(ctx)

	var r0 *version.PackageInfo
	if rf, ok := ret.Get(0).(func(context.Context) *version.PackageInfo); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*version.PackageInfo)
		}
	}

	return r0
}

// LastCheckUpdatesResult provides a mock function with given fields: ctx
func (_m *mockSupervisordService) LastCheckUpdatesResult(ctx context.Context) (*version.UpdateCheckResult, time.Time) {
	ret := _m.Called(ctx)

	var r0 *version.UpdateCheckResult
	if rf, ok := ret.Get(0).(func(context.Context) *version.UpdateCheckResult); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*version.UpdateCheckResult)
		}
	}

	var r1 time.Time
	if rf, ok := ret.Get(1).(func(context.Context) time.Time); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Get(1).(time.Time)
	}

	return r0, r1
}

// StartUpdate provides a mock function with given fields:
func (_m *mockSupervisordService) StartUpdate() (uint32, error) {
	ret := _m.Called()

	var r0 uint32
	if rf, ok := ret.Get(0).(func() uint32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateConfiguration provides a mock function with given fields: settings, ssoDetails
func (_m *mockSupervisordService) UpdateConfiguration(settings *models.Settings, ssoDetails *models.PerconaSSODetails) error {
	ret := _m.Called(settings, ssoDetails)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Settings, *models.PerconaSSODetails) error); ok {
		r0 = rf(settings, ssoDetails)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateLog provides a mock function with given fields: offset
func (_m *mockSupervisordService) UpdateLog(offset uint32) ([]string, uint32, error) {
	ret := _m.Called(offset)

	var r0 []string
	if rf, ok := ret.Get(0).(func(uint32) []string); ok {
		r0 = rf(offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 uint32
	if rf, ok := ret.Get(1).(func(uint32) uint32); ok {
		r1 = rf(offset)
	} else {
		r1 = ret.Get(1).(uint32)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(uint32) error); ok {
		r2 = rf(offset)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// UpdateRunning provides a mock function with given fields:
func (_m *mockSupervisordService) UpdateRunning() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}
