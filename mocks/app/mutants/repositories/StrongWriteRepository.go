// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"
	mutants "mutants/app/mutants"

	mock "github.com/stretchr/testify/mock"
)

// StrongWriteRepository is an autogenerated mock type for the StrongWriteRepository type
type StrongWriteRepository struct {
	mock.Mock
}

// SaveDna provides a mock function with given fields: ctx, human
func (_m *StrongWriteRepository) SaveDna(ctx context.Context, human mutants.Human) error {
	ret := _m.Called(ctx, human)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, mutants.Human) error); ok {
		r0 = rf(ctx, human)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}