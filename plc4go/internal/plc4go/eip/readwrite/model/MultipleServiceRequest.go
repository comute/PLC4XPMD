/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package model

import (
	"fmt"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// Constant values.
const MultipleServiceRequest_REQUESTPATHSIZE int8 = 0x02
const MultipleServiceRequest_REQUESTPATH uint32 = 0x01240220

// The data-structure of this message
type MultipleServiceRequest struct {
	*CipService
	Data *Services
}

// The corresponding interface
type IMultipleServiceRequest interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *MultipleServiceRequest) Service() uint8 {
	return 0x0A
}

func (m *MultipleServiceRequest) InitializeParent(parent *CipService) {
}

func NewMultipleServiceRequest(data *Services) *CipService {
	child := &MultipleServiceRequest{
		Data:       data,
		CipService: NewCipService(),
	}
	child.Child = child
	return child.CipService
}

func CastMultipleServiceRequest(structType interface{}) *MultipleServiceRequest {
	castFunc := func(typ interface{}) *MultipleServiceRequest {
		if casted, ok := typ.(MultipleServiceRequest); ok {
			return &casted
		}
		if casted, ok := typ.(*MultipleServiceRequest); ok {
			return casted
		}
		if casted, ok := typ.(CipService); ok {
			return CastMultipleServiceRequest(casted.Child)
		}
		if casted, ok := typ.(*CipService); ok {
			return CastMultipleServiceRequest(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *MultipleServiceRequest) GetTypeName() string {
	return "MultipleServiceRequest"
}

func (m *MultipleServiceRequest) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *MultipleServiceRequest) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.ParentLengthInBits())

	// Const Field (requestPathSize)
	lengthInBits += 8

	// Const Field (requestPath)
	lengthInBits += 32

	// Simple field (data)
	lengthInBits += m.Data.LengthInBits()

	return lengthInBits
}

func (m *MultipleServiceRequest) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func MultipleServiceRequestParse(readBuffer utils.ReadBuffer, serviceLen uint16) (*CipService, error) {
	if pullErr := readBuffer.PullContext("MultipleServiceRequest"); pullErr != nil {
		return nil, pullErr
	}

	// Const Field (requestPathSize)
	requestPathSize, _requestPathSizeErr := readBuffer.ReadInt8("requestPathSize", 8)
	if _requestPathSizeErr != nil {
		return nil, errors.Wrap(_requestPathSizeErr, "Error parsing 'requestPathSize' field")
	}
	if requestPathSize != MultipleServiceRequest_REQUESTPATHSIZE {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", MultipleServiceRequest_REQUESTPATHSIZE) + " but got " + fmt.Sprintf("%d", requestPathSize))
	}

	// Const Field (requestPath)
	requestPath, _requestPathErr := readBuffer.ReadUint32("requestPath", 32)
	if _requestPathErr != nil {
		return nil, errors.Wrap(_requestPathErr, "Error parsing 'requestPath' field")
	}
	if requestPath != MultipleServiceRequest_REQUESTPATH {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", MultipleServiceRequest_REQUESTPATH) + " but got " + fmt.Sprintf("%d", requestPath))
	}

	// Simple Field (data)
	if pullErr := readBuffer.PullContext("data"); pullErr != nil {
		return nil, pullErr
	}
	_data, _dataErr := ServicesParse(readBuffer, uint16(serviceLen)-uint16(uint16(6)))
	if _dataErr != nil {
		return nil, errors.Wrap(_dataErr, "Error parsing 'data' field")
	}
	data := CastServices(_data)
	if closeErr := readBuffer.CloseContext("data"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("MultipleServiceRequest"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &MultipleServiceRequest{
		Data:       CastServices(data),
		CipService: &CipService{},
	}
	_child.CipService.Child = _child
	return _child.CipService, nil
}

func (m *MultipleServiceRequest) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MultipleServiceRequest"); pushErr != nil {
			return pushErr
		}

		// Const Field (requestPathSize)
		_requestPathSizeErr := writeBuffer.WriteInt8("requestPathSize", 8, 0x02)
		if _requestPathSizeErr != nil {
			return errors.Wrap(_requestPathSizeErr, "Error serializing 'requestPathSize' field")
		}

		// Const Field (requestPath)
		_requestPathErr := writeBuffer.WriteUint32("requestPath", 32, 0x01240220)
		if _requestPathErr != nil {
			return errors.Wrap(_requestPathErr, "Error serializing 'requestPath' field")
		}

		// Simple Field (data)
		if pushErr := writeBuffer.PushContext("data"); pushErr != nil {
			return pushErr
		}
		_dataErr := m.Data.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("data"); popErr != nil {
			return popErr
		}
		if _dataErr != nil {
			return errors.Wrap(_dataErr, "Error serializing 'data' field")
		}

		if popErr := writeBuffer.PopContext("MultipleServiceRequest"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *MultipleServiceRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
