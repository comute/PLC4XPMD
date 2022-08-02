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

package model

import (
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// ParameterValueSerialNumber is the corresponding interface of ParameterValueSerialNumber
type ParameterValueSerialNumber interface {
	utils.LengthAware
	utils.Serializable
	ParameterValue
	// GetValue returns Value (property field)
	GetValue() SerialNumber
}

// ParameterValueSerialNumberExactly can be used when we want exactly this type and not a type which fulfills ParameterValueSerialNumber.
// This is useful for switch cases.
type ParameterValueSerialNumberExactly interface {
	ParameterValueSerialNumber
	isParameterValueSerialNumber() bool
}

// _ParameterValueSerialNumber is the data-structure of this message
type _ParameterValueSerialNumber struct {
	*_ParameterValue
	Value SerialNumber
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ParameterValueSerialNumber) GetParameterType() ParameterType {
	return ParameterType_SERIAL_NUMBER
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ParameterValueSerialNumber) InitializeParent(parent ParameterValue) {}

func (m *_ParameterValueSerialNumber) GetParent() ParameterValue {
	return m._ParameterValue
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ParameterValueSerialNumber) GetValue() SerialNumber {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewParameterValueSerialNumber factory function for _ParameterValueSerialNumber
func NewParameterValueSerialNumber(value SerialNumber, numBytes uint8) *_ParameterValueSerialNumber {
	_result := &_ParameterValueSerialNumber{
		Value:           value,
		_ParameterValue: NewParameterValue(numBytes),
	}
	_result._ParameterValue._ParameterValueChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastParameterValueSerialNumber(structType interface{}) ParameterValueSerialNumber {
	if casted, ok := structType.(ParameterValueSerialNumber); ok {
		return casted
	}
	if casted, ok := structType.(*ParameterValueSerialNumber); ok {
		return *casted
	}
	return nil
}

func (m *_ParameterValueSerialNumber) GetTypeName() string {
	return "ParameterValueSerialNumber"
}

func (m *_ParameterValueSerialNumber) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_ParameterValueSerialNumber) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (value)
	lengthInBits += m.Value.GetLengthInBits()

	return lengthInBits
}

func (m *_ParameterValueSerialNumber) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ParameterValueSerialNumberParse(readBuffer utils.ReadBuffer, parameterType ParameterType, numBytes uint8) (ParameterValueSerialNumber, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ParameterValueSerialNumber"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ParameterValueSerialNumber")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(bool((numBytes) == (4))) {
		return nil, errors.WithStack(utils.ParseValidationError{"SerialNumber has exactly four bytes"})
	}

	// Simple Field (value)
	if pullErr := readBuffer.PullContext("value"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for value")
	}
	_value, _valueErr := SerialNumberParse(readBuffer)
	if _valueErr != nil {
		return nil, errors.Wrap(_valueErr, "Error parsing 'value' field of ParameterValueSerialNumber")
	}
	value := _value.(SerialNumber)
	if closeErr := readBuffer.CloseContext("value"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for value")
	}

	if closeErr := readBuffer.CloseContext("ParameterValueSerialNumber"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ParameterValueSerialNumber")
	}

	// Create a partially initialized instance
	_child := &_ParameterValueSerialNumber{
		Value: value,
		_ParameterValue: &_ParameterValue{
			NumBytes: numBytes,
		},
	}
	_child._ParameterValue._ParameterValueChildRequirements = _child
	return _child, nil
}

func (m *_ParameterValueSerialNumber) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ParameterValueSerialNumber"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ParameterValueSerialNumber")
		}

		// Simple Field (value)
		if pushErr := writeBuffer.PushContext("value"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for value")
		}
		_valueErr := writeBuffer.WriteSerializable(m.GetValue())
		if popErr := writeBuffer.PopContext("value"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for value")
		}
		if _valueErr != nil {
			return errors.Wrap(_valueErr, "Error serializing 'value' field")
		}

		if popErr := writeBuffer.PopContext("ParameterValueSerialNumber"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ParameterValueSerialNumber")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_ParameterValueSerialNumber) isParameterValueSerialNumber() bool {
	return true
}

func (m *_ParameterValueSerialNumber) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
