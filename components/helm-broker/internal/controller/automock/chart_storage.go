// Code generated by mockery v1.0.0
package automock

import chart "k8s.io/helm/pkg/proto/hapi/chart"

import internal "github.com/kyma-project/kyma/components/helm-broker/internal"
import mock "github.com/stretchr/testify/mock"
import semver "github.com/Masterminds/semver"

// ChartStorage is an autogenerated mock type for the ChartStorage type
type ChartStorage struct {
	mock.Mock
}

// Remove provides a mock function with given fields: _a0, _a1, _a2
func (_m *ChartStorage) Remove(_a0 internal.Namespace, _a1 internal.ChartName, _a2 semver.Version) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(internal.Namespace, internal.ChartName, semver.Version) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Upsert provides a mock function with given fields: _a0, _a1
func (_m *ChartStorage) Upsert(_a0 internal.Namespace, _a1 *chart.Chart) (bool, error) {
	ret := _m.Called(_a0, _a1)

	var r0 bool
	if rf, ok := ret.Get(0).(func(internal.Namespace, *chart.Chart) bool); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(internal.Namespace, *chart.Chart) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}