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
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// ApduDataDeviceDescriptorResponse is the data-structure of this message
type ApduDataDeviceDescriptorResponse struct {
	*ApduData
	DescriptorType uint8
	Data           []byte

	// Arguments.
	DataLength uint8
}

// IApduDataDeviceDescriptorResponse is the corresponding interface of ApduDataDeviceDescriptorResponse
type IApduDataDeviceDescriptorResponse interface {
	IApduData
	// GetDescriptorType returns DescriptorType (property field)
	GetDescriptorType() uint8
	// GetData returns Data (property field)
	GetData() []byte
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *ApduDataDeviceDescriptorResponse) GetApciType() uint8 {
	return 0xD
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *ApduDataDeviceDescriptorResponse) InitializeParent(parent *ApduData) {}

func (m *ApduDataDeviceDescriptorResponse) GetParent() *ApduData {
	return m.ApduData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *ApduDataDeviceDescriptorResponse) GetDescriptorType() uint8 {
	return m.DescriptorType
}

func (m *ApduDataDeviceDescriptorResponse) GetData() []byte {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewApduDataDeviceDescriptorResponse factory function for ApduDataDeviceDescriptorResponse
func NewApduDataDeviceDescriptorResponse(descriptorType uint8, data []byte, dataLength uint8) *ApduDataDeviceDescriptorResponse {
	_result := &ApduDataDeviceDescriptorResponse{
		DescriptorType: descriptorType,
		Data:           data,
		ApduData:       NewApduData(dataLength),
	}
	_result.Child = _result
	return _result
}

func CastApduDataDeviceDescriptorResponse(structType interface{}) *ApduDataDeviceDescriptorResponse {
	if casted, ok := structType.(ApduDataDeviceDescriptorResponse); ok {
		return &casted
	}
	if casted, ok := structType.(*ApduDataDeviceDescriptorResponse); ok {
		return casted
	}
	if casted, ok := structType.(ApduData); ok {
		return CastApduDataDeviceDescriptorResponse(casted.Child)
	}
	if casted, ok := structType.(*ApduData); ok {
		return CastApduDataDeviceDescriptorResponse(casted.Child)
	}
	return nil
}

func (m *ApduDataDeviceDescriptorResponse) GetTypeName() string {
	return "ApduDataDeviceDescriptorResponse"
}

func (m *ApduDataDeviceDescriptorResponse) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *ApduDataDeviceDescriptorResponse) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (descriptorType)
	lengthInBits += 6

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *ApduDataDeviceDescriptorResponse) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ApduDataDeviceDescriptorResponseParse(readBuffer utils.ReadBuffer, dataLength uint8) (*ApduDataDeviceDescriptorResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataDeviceDescriptorResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataDeviceDescriptorResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (descriptorType)
	_descriptorType, _descriptorTypeErr := readBuffer.ReadUint8("descriptorType", 6)
	if _descriptorTypeErr != nil {
		return nil, errors.Wrap(_descriptorTypeErr, "Error parsing 'descriptorType' field")
	}
	descriptorType := _descriptorType
	// Byte Array field (data)
	numberOfBytesdata := int(utils.InlineIf(bool(bool((dataLength) < (1))), func() interface{} { return uint16(uint16(0)) }, func() interface{} { return uint16(uint16(dataLength) - uint16(uint16(1))) }).(uint16))
	data, _readArrayErr := readBuffer.ReadByteArray("data", numberOfBytesdata)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'data' field")
	}

	if closeErr := readBuffer.CloseContext("ApduDataDeviceDescriptorResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataDeviceDescriptorResponse")
	}

	// Create a partially initialized instance
	_child := &ApduDataDeviceDescriptorResponse{
		DescriptorType: descriptorType,
		Data:           data,
		ApduData:       &ApduData{},
	}
	_child.ApduData.Child = _child
	return _child, nil
}

func (m *ApduDataDeviceDescriptorResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataDeviceDescriptorResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataDeviceDescriptorResponse")
		}

		// Simple Field (descriptorType)
		descriptorType := uint8(m.DescriptorType)
		_descriptorTypeErr := writeBuffer.WriteUint8("descriptorType", 6, (descriptorType))
		if _descriptorTypeErr != nil {
			return errors.Wrap(_descriptorTypeErr, "Error serializing 'descriptorType' field")
		}

		// Array Field (data)
		if m.Data != nil {
			// Byte Array field (data)
			_writeArrayErr := writeBuffer.WriteByteArray("data", m.Data)
			if _writeArrayErr != nil {
				return errors.Wrap(_writeArrayErr, "Error serializing 'data' field")
			}
		}

		if popErr := writeBuffer.PopContext("ApduDataDeviceDescriptorResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataDeviceDescriptorResponse")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *ApduDataDeviceDescriptorResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
