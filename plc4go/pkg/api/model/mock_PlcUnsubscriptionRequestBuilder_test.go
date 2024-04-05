/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

// Code generated by mockery v2.42.2. DO NOT EDIT.

package model

import mock "github.com/stretchr/testify/mock"

// MockPlcUnsubscriptionRequestBuilder is an autogenerated mock type for the PlcUnsubscriptionRequestBuilder type
type MockPlcUnsubscriptionRequestBuilder struct {
	mock.Mock
}

type MockPlcUnsubscriptionRequestBuilder_Expecter struct {
	mock *mock.Mock
}

func (_m *MockPlcUnsubscriptionRequestBuilder) EXPECT() *MockPlcUnsubscriptionRequestBuilder_Expecter {
	return &MockPlcUnsubscriptionRequestBuilder_Expecter{mock: &_m.Mock}
}

// AddHandles provides a mock function with given fields: handles
func (_m *MockPlcUnsubscriptionRequestBuilder) AddHandles(handles ...PlcSubscriptionHandle) PlcUnsubscriptionRequestBuilder {
	_va := make([]interface{}, len(handles))
	for _i := range handles {
		_va[_i] = handles[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for AddHandles")
	}

	var r0 PlcUnsubscriptionRequestBuilder
	if rf, ok := ret.Get(0).(func(...PlcSubscriptionHandle) PlcUnsubscriptionRequestBuilder); ok {
		r0 = rf(handles...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(PlcUnsubscriptionRequestBuilder)
		}
	}

	return r0
}

// MockPlcUnsubscriptionRequestBuilder_AddHandles_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddHandles'
type MockPlcUnsubscriptionRequestBuilder_AddHandles_Call struct {
	*mock.Call
}

// AddHandles is a helper method to define mock.On call
//   - handles ...PlcSubscriptionHandle
func (_e *MockPlcUnsubscriptionRequestBuilder_Expecter) AddHandles(handles ...interface{}) *MockPlcUnsubscriptionRequestBuilder_AddHandles_Call {
	return &MockPlcUnsubscriptionRequestBuilder_AddHandles_Call{Call: _e.mock.On("AddHandles",
		append([]interface{}{}, handles...)...)}
}

func (_c *MockPlcUnsubscriptionRequestBuilder_AddHandles_Call) Run(run func(handles ...PlcSubscriptionHandle)) *MockPlcUnsubscriptionRequestBuilder_AddHandles_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]PlcSubscriptionHandle, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(PlcSubscriptionHandle)
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *MockPlcUnsubscriptionRequestBuilder_AddHandles_Call) Return(_a0 PlcUnsubscriptionRequestBuilder) *MockPlcUnsubscriptionRequestBuilder_AddHandles_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcUnsubscriptionRequestBuilder_AddHandles_Call) RunAndReturn(run func(...PlcSubscriptionHandle) PlcUnsubscriptionRequestBuilder) *MockPlcUnsubscriptionRequestBuilder_AddHandles_Call {
	_c.Call.Return(run)
	return _c
}

// Build provides a mock function with given fields:
func (_m *MockPlcUnsubscriptionRequestBuilder) Build() (PlcUnsubscriptionRequest, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Build")
	}

	var r0 PlcUnsubscriptionRequest
	var r1 error
	if rf, ok := ret.Get(0).(func() (PlcUnsubscriptionRequest, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() PlcUnsubscriptionRequest); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(PlcUnsubscriptionRequest)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPlcUnsubscriptionRequestBuilder_Build_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Build'
type MockPlcUnsubscriptionRequestBuilder_Build_Call struct {
	*mock.Call
}

// Build is a helper method to define mock.On call
func (_e *MockPlcUnsubscriptionRequestBuilder_Expecter) Build() *MockPlcUnsubscriptionRequestBuilder_Build_Call {
	return &MockPlcUnsubscriptionRequestBuilder_Build_Call{Call: _e.mock.On("Build")}
}

func (_c *MockPlcUnsubscriptionRequestBuilder_Build_Call) Run(run func()) *MockPlcUnsubscriptionRequestBuilder_Build_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcUnsubscriptionRequestBuilder_Build_Call) Return(_a0 PlcUnsubscriptionRequest, _a1 error) *MockPlcUnsubscriptionRequestBuilder_Build_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPlcUnsubscriptionRequestBuilder_Build_Call) RunAndReturn(run func() (PlcUnsubscriptionRequest, error)) *MockPlcUnsubscriptionRequestBuilder_Build_Call {
	_c.Call.Return(run)
	return _c
}

// String provides a mock function with given fields:
func (_m *MockPlcUnsubscriptionRequestBuilder) String() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for String")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockPlcUnsubscriptionRequestBuilder_String_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'String'
type MockPlcUnsubscriptionRequestBuilder_String_Call struct {
	*mock.Call
}

// String is a helper method to define mock.On call
func (_e *MockPlcUnsubscriptionRequestBuilder_Expecter) String() *MockPlcUnsubscriptionRequestBuilder_String_Call {
	return &MockPlcUnsubscriptionRequestBuilder_String_Call{Call: _e.mock.On("String")}
}

func (_c *MockPlcUnsubscriptionRequestBuilder_String_Call) Run(run func()) *MockPlcUnsubscriptionRequestBuilder_String_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcUnsubscriptionRequestBuilder_String_Call) Return(_a0 string) *MockPlcUnsubscriptionRequestBuilder_String_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcUnsubscriptionRequestBuilder_String_Call) RunAndReturn(run func() string) *MockPlcUnsubscriptionRequestBuilder_String_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockPlcUnsubscriptionRequestBuilder creates a new instance of MockPlcUnsubscriptionRequestBuilder. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockPlcUnsubscriptionRequestBuilder(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockPlcUnsubscriptionRequestBuilder {
	mock := &MockPlcUnsubscriptionRequestBuilder{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
