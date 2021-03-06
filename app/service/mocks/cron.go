// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"

	cron "github.com/robfig/cron/v3"
	mock "github.com/stretchr/testify/mock"
)

// Cron is an autogenerated mock type for the Cron type
type Cron struct {
	mock.Mock
}

// Entries provides a mock function with given fields:
func (_m *Cron) Entries() []cron.Entry {
	ret := _m.Called()

	var r0 []cron.Entry
	if rf, ok := ret.Get(0).(func() []cron.Entry); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]cron.Entry)
		}
	}

	return r0
}

// Remove provides a mock function with given fields: id
func (_m *Cron) Remove(id cron.EntryID) {
	_m.Called(id)
}

// Schedule provides a mock function with given fields: schedule, cmd
func (_m *Cron) Schedule(schedule cron.Schedule, cmd cron.Job) cron.EntryID {
	ret := _m.Called(schedule, cmd)

	var r0 cron.EntryID
	if rf, ok := ret.Get(0).(func(cron.Schedule, cron.Job) cron.EntryID); ok {
		r0 = rf(schedule, cmd)
	} else {
		r0 = ret.Get(0).(cron.EntryID)
	}

	return r0
}

// Start provides a mock function with given fields:
func (_m *Cron) Start() {
	_m.Called()
}

// Stop provides a mock function with given fields:
func (_m *Cron) Stop() context.Context {
	ret := _m.Called()

	var r0 context.Context
	if rf, ok := ret.Get(0).(func() context.Context); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(context.Context)
		}
	}

	return r0
}
