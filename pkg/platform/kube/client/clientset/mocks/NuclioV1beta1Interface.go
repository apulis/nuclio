/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by mockery v2.4.0-beta. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	rest "k8s.io/client-go/rest"

	v1beta1 "github.com/nuclio/nuclio/pkg/platform/kube/client/clientset/versioned/typed/nuclio.io/v1beta1"
)

// NuclioV1beta1Interface is an autogenerated mock type for the NuclioV1beta1Interface type
type NuclioV1beta1Interface struct {
	mock.Mock
}

// NuclioAPIGateways provides a mock function with given fields: namespace
func (_m *NuclioV1beta1Interface) NuclioAPIGateways(namespace string) v1beta1.NuclioAPIGatewayInterface {
	ret := _m.Called(namespace)

	var r0 v1beta1.NuclioAPIGatewayInterface
	if rf, ok := ret.Get(0).(func(string) v1beta1.NuclioAPIGatewayInterface); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1beta1.NuclioAPIGatewayInterface)
		}
	}

	return r0
}

// NuclioFunctionEvents provides a mock function with given fields: namespace
func (_m *NuclioV1beta1Interface) NuclioFunctionEvents(namespace string) v1beta1.NuclioFunctionEventInterface {
	ret := _m.Called(namespace)

	var r0 v1beta1.NuclioFunctionEventInterface
	if rf, ok := ret.Get(0).(func(string) v1beta1.NuclioFunctionEventInterface); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1beta1.NuclioFunctionEventInterface)
		}
	}

	return r0
}

// NuclioFunctions provides a mock function with given fields: namespace
func (_m *NuclioV1beta1Interface) NuclioFunctions(namespace string) v1beta1.NuclioFunctionInterface {
	ret := _m.Called(namespace)

	var r0 v1beta1.NuclioFunctionInterface
	if rf, ok := ret.Get(0).(func(string) v1beta1.NuclioFunctionInterface); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1beta1.NuclioFunctionInterface)
		}
	}

	return r0
}

// NuclioProjects provides a mock function with given fields: namespace
func (_m *NuclioV1beta1Interface) NuclioProjects(namespace string) v1beta1.NuclioProjectInterface {
	ret := _m.Called(namespace)

	var r0 v1beta1.NuclioProjectInterface
	if rf, ok := ret.Get(0).(func(string) v1beta1.NuclioProjectInterface); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1beta1.NuclioProjectInterface)
		}
	}

	return r0
}

// RESTClient provides a mock function with given fields:
func (_m *NuclioV1beta1Interface) RESTClient() rest.Interface {
	ret := _m.Called()

	var r0 rest.Interface
	if rf, ok := ret.Get(0).(func() rest.Interface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(rest.Interface)
		}
	}

	return r0
}