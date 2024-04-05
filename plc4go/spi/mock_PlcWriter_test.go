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

package spi

import (
	context "context"

	model "github.com/apache/plc4x/plc4go/pkg/api/model"
	mock "github.com/stretchr/testify/mock"
)

// MockPlcWriter is an autogenerated mock type for the PlcWriter type
type MockPlcWriter struct {
	mock.Mock
}

type MockPlcWriter_Expecter struct {
	mock *mock.Mock
}

func (_m *MockPlcWriter) EXPECT() *MockPlcWriter_Expecter {
	return &MockPlcWriter_Expecter{mock: &_m.Mock}
}

// Write provides a mock function with given fields: ctx, writeRequest
func (_m *MockPlcWriter) Write(ctx context.Context, writeRequest model.PlcWriteRequest) <-chan model.PlcWriteRequestResult {
	ret := _m.Called(ctx, writeRequest)

	if len(ret) == 0 {
		panic("no return value specified for Write")
	}

	var r0 <-chan model.PlcWriteRequestResult
	if rf, ok := ret.Get(0).(func(context.Context, model.PlcWriteRequest) <-chan model.PlcWriteRequestResult); ok {
		r0 = rf(ctx, writeRequest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan model.PlcWriteRequestResult)
		}
	}

	return r0
}

// MockPlcWriter_Write_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Write'
type MockPlcWriter_Write_Call struct {
	*mock.Call
}

// Write is a helper method to define mock.On call
//   - ctx context.Context
//   - writeRequest model.PlcWriteRequest
func (_e *MockPlcWriter_Expecter) Write(ctx interface{}, writeRequest interface{}) *MockPlcWriter_Write_Call {
	return &MockPlcWriter_Write_Call{Call: _e.mock.On("Write", ctx, writeRequest)}
}

func (_c *MockPlcWriter_Write_Call) Run(run func(ctx context.Context, writeRequest model.PlcWriteRequest)) *MockPlcWriter_Write_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.PlcWriteRequest))
	})
	return _c
}

func (_c *MockPlcWriter_Write_Call) Return(_a0 <-chan model.PlcWriteRequestResult) *MockPlcWriter_Write_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcWriter_Write_Call) RunAndReturn(run func(context.Context, model.PlcWriteRequest) <-chan model.PlcWriteRequestResult) *MockPlcWriter_Write_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockPlcWriter creates a new instance of MockPlcWriter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockPlcWriter(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockPlcWriter {
	mock := &MockPlcWriter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
