// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	dto "leal.co/orders/internal/dto"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// CreateOrder provides a mock function with given fields: _a0
func (_m *Service) CreateOrder(_a0 dto.CreateOrderDTO) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(dto.CreateOrderDTO) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
