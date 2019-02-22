// Code generated by mockery v1.0.0. DO NOT EDIT.

package automock

import gqlschema "github.com/kyma-project/kyma/components/ui-api-layer/internal/gqlschema"

import mock "github.com/stretchr/testify/mock"
import v1 "k8s.io/api/apps/v1"

// gqlReplicaSetConverter is an autogenerated mock type for the gqlReplicaSetConverter type
type gqlReplicaSetConverter struct {
	mock.Mock
}

// GQLJSONToReplicaSet provides a mock function with given fields: in
func (_m *gqlReplicaSetConverter) GQLJSONToReplicaSet(in gqlschema.JSON) (v1.ReplicaSet, error) {
	ret := _m.Called(in)

	var r0 v1.ReplicaSet
	if rf, ok := ret.Get(0).(func(gqlschema.JSON) v1.ReplicaSet); ok {
		r0 = rf(in)
	} else {
		r0 = ret.Get(0).(v1.ReplicaSet)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(gqlschema.JSON) error); ok {
		r1 = rf(in)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ToGQL provides a mock function with given fields: in
func (_m *gqlReplicaSetConverter) ToGQL(in *v1.ReplicaSet) (*gqlschema.ReplicaSet, error) {
	ret := _m.Called(in)

	var r0 *gqlschema.ReplicaSet
	if rf, ok := ret.Get(0).(func(*v1.ReplicaSet) *gqlschema.ReplicaSet); ok {
		r0 = rf(in)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gqlschema.ReplicaSet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*v1.ReplicaSet) error); ok {
		r1 = rf(in)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ToGQLs provides a mock function with given fields: in
func (_m *gqlReplicaSetConverter) ToGQLs(in []*v1.ReplicaSet) ([]gqlschema.ReplicaSet, error) {
	ret := _m.Called(in)

	var r0 []gqlschema.ReplicaSet
	if rf, ok := ret.Get(0).(func([]*v1.ReplicaSet) []gqlschema.ReplicaSet); ok {
		r0 = rf(in)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]gqlschema.ReplicaSet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]*v1.ReplicaSet) error); ok {
		r1 = rf(in)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
