// Code generated by mockery v1.0.0. DO NOT EDIT.

package automock

import mock "github.com/stretchr/testify/mock"
import resource "github.com/kyma-project/kyma/components/console-backend-service/pkg/resource"

// addonsCfgService is an autogenerated mock type for the addonsCfgService type
type addonsCfgService struct {
	mock.Mock
}

// Subscribe provides a mock function with given fields: listener
func (_m *addonsCfgService) Subscribe(listener resource.Listener) {
	_m.Called(listener)
}

// Unsubscribe provides a mock function with given fields: listener
func (_m *addonsCfgService) Unsubscribe(listener resource.Listener) {
	_m.Called(listener)
}
