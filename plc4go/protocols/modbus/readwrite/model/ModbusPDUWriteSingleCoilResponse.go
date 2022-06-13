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

// ModbusPDUWriteSingleCoilResponse is the data-structure of this message
type ModbusPDUWriteSingleCoilResponse struct {
	*ModbusPDU
	Address uint16
	Value   uint16
}

// IModbusPDUWriteSingleCoilResponse is the corresponding interface of ModbusPDUWriteSingleCoilResponse
type IModbusPDUWriteSingleCoilResponse interface {
	IModbusPDU
	// GetAddress returns Address (property field)
	GetAddress() uint16
	// GetValue returns Value (property field)
	GetValue() uint16
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

func (m *ModbusPDUWriteSingleCoilResponse) GetErrorFlag() bool {
	return bool(false)
}

func (m *ModbusPDUWriteSingleCoilResponse) GetFunctionFlag() uint8 {
	return 0x05
}

func (m *ModbusPDUWriteSingleCoilResponse) GetResponse() bool {
	return bool(true)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *ModbusPDUWriteSingleCoilResponse) InitializeParent(parent *ModbusPDU) {}

func (m *ModbusPDUWriteSingleCoilResponse) GetParent() *ModbusPDU {
	return m.ModbusPDU
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *ModbusPDUWriteSingleCoilResponse) GetAddress() uint16 {
	return m.Address
}

func (m *ModbusPDUWriteSingleCoilResponse) GetValue() uint16 {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewModbusPDUWriteSingleCoilResponse factory function for ModbusPDUWriteSingleCoilResponse
func NewModbusPDUWriteSingleCoilResponse(address uint16, value uint16) *ModbusPDUWriteSingleCoilResponse {
	_result := &ModbusPDUWriteSingleCoilResponse{
		Address:   address,
		Value:     value,
		ModbusPDU: NewModbusPDU(),
	}
	_result.Child = _result
	return _result
}

func CastModbusPDUWriteSingleCoilResponse(structType interface{}) *ModbusPDUWriteSingleCoilResponse {
	if casted, ok := structType.(ModbusPDUWriteSingleCoilResponse); ok {
		return &casted
	}
	if casted, ok := structType.(*ModbusPDUWriteSingleCoilResponse); ok {
		return casted
	}
	if casted, ok := structType.(ModbusPDU); ok {
		return CastModbusPDUWriteSingleCoilResponse(casted.Child)
	}
	if casted, ok := structType.(*ModbusPDU); ok {
		return CastModbusPDUWriteSingleCoilResponse(casted.Child)
	}
	return nil
}

func (m *ModbusPDUWriteSingleCoilResponse) GetTypeName() string {
	return "ModbusPDUWriteSingleCoilResponse"
}

func (m *ModbusPDUWriteSingleCoilResponse) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *ModbusPDUWriteSingleCoilResponse) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (address)
	lengthInBits += 16

	// Simple field (value)
	lengthInBits += 16

	return lengthInBits
}

func (m *ModbusPDUWriteSingleCoilResponse) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ModbusPDUWriteSingleCoilResponseParse(readBuffer utils.ReadBuffer, response bool) (*ModbusPDUWriteSingleCoilResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModbusPDUWriteSingleCoilResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusPDUWriteSingleCoilResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (address)
	_address, _addressErr := readBuffer.ReadUint16("address", 16)
	if _addressErr != nil {
		return nil, errors.Wrap(_addressErr, "Error parsing 'address' field")
	}
	address := _address

	// Simple Field (value)
	_value, _valueErr := readBuffer.ReadUint16("value", 16)
	if _valueErr != nil {
		return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
	}
	value := _value

	if closeErr := readBuffer.CloseContext("ModbusPDUWriteSingleCoilResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusPDUWriteSingleCoilResponse")
	}

	// Create a partially initialized instance
	_child := &ModbusPDUWriteSingleCoilResponse{
		Address:   address,
		Value:     value,
		ModbusPDU: &ModbusPDU{},
	}
	_child.ModbusPDU.Child = _child
	return _child, nil
}

func (m *ModbusPDUWriteSingleCoilResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUWriteSingleCoilResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ModbusPDUWriteSingleCoilResponse")
		}

		// Simple Field (address)
		address := uint16(m.Address)
		_addressErr := writeBuffer.WriteUint16("address", 16, (address))
		if _addressErr != nil {
			return errors.Wrap(_addressErr, "Error serializing 'address' field")
		}

		// Simple Field (value)
		value := uint16(m.Value)
		_valueErr := writeBuffer.WriteUint16("value", 16, (value))
		if _valueErr != nil {
			return errors.Wrap(_valueErr, "Error serializing 'value' field")
		}

		if popErr := writeBuffer.PopContext("ModbusPDUWriteSingleCoilResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ModbusPDUWriteSingleCoilResponse")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *ModbusPDUWriteSingleCoilResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
