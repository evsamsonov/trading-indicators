// Code generated by mockery v1.0.0. DO NOT EDIT.

package indicator

import mock "github.com/stretchr/testify/mock"

// MockIndicator is an autogenerated mock type for the Indicator type
type MockIndicator struct {
	mock.Mock
}

// Calculate provides a mock function with given fields: index
func (_m *MockIndicator) Calculate(index int) float64 {
	ret := _m.Called(index)

	var r0 float64
	if rf, ok := ret.Get(0).(func(int) float64); ok {
		r0 = rf(index)
	} else {
		r0 = ret.Get(0).(float64)
	}

	return r0
}
