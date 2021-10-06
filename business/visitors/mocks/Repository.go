// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	visitors "backend/business/visitors"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// AddVisitor provides a mock function with given fields: ctx, domain
func (_m *Repository) AddVisitor(ctx context.Context, domain visitors.Domain) (visitors.Domain, error) {
	ret := _m.Called(ctx, domain)

	var r0 visitors.Domain
	if rf, ok := ret.Get(0).(func(context.Context, visitors.Domain) visitors.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		r0 = ret.Get(0).(visitors.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, visitors.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CancelVisitor provides a mock function with given fields: ctx, domain
func (_m *Repository) CancelVisitor(ctx context.Context, domain visitors.Domain) (visitors.Domain, error) {
	ret := _m.Called(ctx, domain)

	var r0 visitors.Domain
	if rf, ok := ret.Get(0).(func(context.Context, visitors.Domain) visitors.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		r0 = ret.Get(0).(visitors.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, visitors.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DontCome provides a mock function with given fields: ctx, log
func (_m *Repository) DontCome(ctx context.Context, log visitors.Log) (visitors.Log, error) {
	ret := _m.Called(ctx, log)

	var r0 visitors.Log
	if rf, ok := ret.Get(0).(func(context.Context, visitors.Log) visitors.Log); ok {
		r0 = rf(ctx, log)
	} else {
		r0 = ret.Get(0).(visitors.Log)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, visitors.Log) error); ok {
		r1 = rf(ctx, log)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDetailSchedule provides a mock function with given fields: ctx, domain
func (_m *Repository) GetDetailSchedule(ctx context.Context, domain visitors.Domain) (uint, uint, error) {
	ret := _m.Called(ctx, domain)

	var r0 uint
	if rf, ok := ret.Get(0).(func(context.Context, visitors.Domain) uint); ok {
		r0 = rf(ctx, domain)
	} else {
		r0 = ret.Get(0).(uint)
	}

	var r1 uint
	if rf, ok := ret.Get(1).(func(context.Context, visitors.Domain) uint); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Get(1).(uint)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, visitors.Domain) error); ok {
		r2 = rf(ctx, domain)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ModificateVisitor provides a mock function with given fields: ctx, domain
func (_m *Repository) ModificateVisitor(ctx context.Context, domain visitors.Domain) (visitors.Domain, error) {
	ret := _m.Called(ctx, domain)

	var r0 visitors.Domain
	if rf, ok := ret.Get(0).(func(context.Context, visitors.Domain) visitors.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		r0 = ret.Get(0).(visitors.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, visitors.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveVisitorToLog provides a mock function with given fields: ctx, log
func (_m *Repository) RemoveVisitorToLog(ctx context.Context, log visitors.Log) (visitors.Log, error) {
	ret := _m.Called(ctx, log)

	var r0 visitors.Log
	if rf, ok := ret.Get(0).(func(context.Context, visitors.Log) visitors.Log); ok {
		r0 = rf(ctx, log)
	} else {
		r0 = ret.Get(0).(visitors.Log)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, visitors.Log) error); ok {
		r1 = rf(ctx, log)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ShowAllPatient provides a mock function with given fields: ctx, domain
func (_m *Repository) ShowAllPatient(ctx context.Context, domain visitors.Domain) ([]visitors.Domain, error) {
	ret := _m.Called(ctx, domain)

	var r0 []visitors.Domain
	if rf, ok := ret.Get(0).(func(context.Context, visitors.Domain) []visitors.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]visitors.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, visitors.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ShowLogOfPatient provides a mock function with given fields: ctx, log
func (_m *Repository) ShowLogOfPatient(ctx context.Context, log visitors.Log) ([]visitors.Log, error) {
	ret := _m.Called(ctx, log)

	var r0 []visitors.Log
	if rf, ok := ret.Get(0).(func(context.Context, visitors.Log) []visitors.Log); ok {
		r0 = rf(ctx, log)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]visitors.Log)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, visitors.Log) error); ok {
		r1 = rf(ctx, log)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ShowVisitor provides a mock function with given fields: ctx, domain
func (_m *Repository) ShowVisitor(ctx context.Context, domain visitors.Domain) (visitors.Domain, error) {
	ret := _m.Called(ctx, domain)

	var r0 visitors.Domain
	if rf, ok := ret.Get(0).(func(context.Context, visitors.Domain) visitors.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		r0 = ret.Get(0).(visitors.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, visitors.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
