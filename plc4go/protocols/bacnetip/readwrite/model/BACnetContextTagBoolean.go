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

// BACnetContextTagBoolean is the data-structure of this message
type BACnetContextTagBoolean struct {
	*BACnetContextTag
	Value   uint8
	Payload *BACnetTagPayloadBoolean

	// Arguments.
	TagNumberArgument uint8
}

// IBACnetContextTagBoolean is the corresponding interface of BACnetContextTagBoolean
type IBACnetContextTagBoolean interface {
	IBACnetContextTag
	// GetValue returns Value (property field)
	GetValue() uint8
	// GetPayload returns Payload (property field)
	GetPayload() *BACnetTagPayloadBoolean
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() bool
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

func (m *BACnetContextTagBoolean) GetDataType() BACnetDataType {
	return BACnetDataType_BOOLEAN
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetContextTagBoolean) InitializeParent(parent *BACnetContextTag, header *BACnetTagHeader) {
	m.BACnetContextTag.Header = header
}

func (m *BACnetContextTagBoolean) GetParent() *BACnetContextTag {
	return m.BACnetContextTag
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetContextTagBoolean) GetValue() uint8 {
	return m.Value
}

func (m *BACnetContextTagBoolean) GetPayload() *BACnetTagPayloadBoolean {
	return m.Payload
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *BACnetContextTagBoolean) GetActualValue() bool {
	return bool(m.GetPayload().GetValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetContextTagBoolean factory function for BACnetContextTagBoolean
func NewBACnetContextTagBoolean(value uint8, payload *BACnetTagPayloadBoolean, header *BACnetTagHeader, tagNumberArgument uint8) *BACnetContextTagBoolean {
	_result := &BACnetContextTagBoolean{
		Value:            value,
		Payload:          payload,
		BACnetContextTag: NewBACnetContextTag(header, tagNumberArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetContextTagBoolean(structType interface{}) *BACnetContextTagBoolean {
	if casted, ok := structType.(BACnetContextTagBoolean); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetContextTagBoolean); ok {
		return casted
	}
	if casted, ok := structType.(BACnetContextTag); ok {
		return CastBACnetContextTagBoolean(casted.Child)
	}
	if casted, ok := structType.(*BACnetContextTag); ok {
		return CastBACnetContextTagBoolean(casted.Child)
	}
	return nil
}

func (m *BACnetContextTagBoolean) GetTypeName() string {
	return "BACnetContextTagBoolean"
}

func (m *BACnetContextTagBoolean) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetContextTagBoolean) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (value)
	lengthInBits += 8

	// Simple field (payload)
	lengthInBits += m.Payload.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *BACnetContextTagBoolean) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetContextTagBooleanParse(readBuffer utils.ReadBuffer, tagNumberArgument uint8, dataType BACnetDataType, header *BACnetTagHeader) (*BACnetContextTagBoolean, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetContextTagBoolean"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetContextTagBoolean")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(bool((header.GetActualLength()) == (1))) {
		return nil, errors.WithStack(utils.ParseValidationError{"length field should be 1"})
	}

	// Simple Field (value)
	_value, _valueErr := readBuffer.ReadUint8("value", 8)
	if _valueErr != nil {
		return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
	}
	value := _value

	// Simple Field (payload)
	if pullErr := readBuffer.PullContext("payload"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for payload")
	}
	_payload, _payloadErr := BACnetTagPayloadBooleanParse(readBuffer, uint32(value))
	if _payloadErr != nil {
		return nil, errors.Wrap(_payloadErr, "Error parsing 'payload' field")
	}
	payload := CastBACnetTagPayloadBoolean(_payload)
	if closeErr := readBuffer.CloseContext("payload"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for payload")
	}

	// Virtual field
	_actualValue := payload.GetValue()
	actualValue := bool(_actualValue)
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetContextTagBoolean"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetContextTagBoolean")
	}

	// Create a partially initialized instance
	_child := &BACnetContextTagBoolean{
		Value:            value,
		Payload:          CastBACnetTagPayloadBoolean(payload),
		BACnetContextTag: &BACnetContextTag{},
	}
	_child.BACnetContextTag.Child = _child
	return _child, nil
}

func (m *BACnetContextTagBoolean) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetContextTagBoolean"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetContextTagBoolean")
		}

		// Simple Field (value)
		value := uint8(m.Value)
		_valueErr := writeBuffer.WriteUint8("value", 8, (value))
		if _valueErr != nil {
			return errors.Wrap(_valueErr, "Error serializing 'value' field")
		}

		// Simple Field (payload)
		if pushErr := writeBuffer.PushContext("payload"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for payload")
		}
		_payloadErr := m.Payload.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("payload"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for payload")
		}
		if _payloadErr != nil {
			return errors.Wrap(_payloadErr, "Error serializing 'payload' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetContextTagBoolean"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetContextTagBoolean")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetContextTagBoolean) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
