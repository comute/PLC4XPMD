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

// BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble is the data-structure of this message
type BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble struct {
	*BACnetFaultParameterFaultOutOfRangeMaxNormalValue
	DoubleValue *BACnetApplicationTagDouble

	// Arguments.
	TagNumber uint8
}

// IBACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble is the corresponding interface of BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble
type IBACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble interface {
	IBACnetFaultParameterFaultOutOfRangeMaxNormalValue
	// GetDoubleValue returns DoubleValue (property field)
	GetDoubleValue() *BACnetApplicationTagDouble
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

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble) InitializeParent(parent *BACnetFaultParameterFaultOutOfRangeMaxNormalValue, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetFaultParameterFaultOutOfRangeMaxNormalValue.OpeningTag = openingTag
	m.BACnetFaultParameterFaultOutOfRangeMaxNormalValue.PeekedTagHeader = peekedTagHeader
	m.BACnetFaultParameterFaultOutOfRangeMaxNormalValue.ClosingTag = closingTag
}

func (m *BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble) GetParent() *BACnetFaultParameterFaultOutOfRangeMaxNormalValue {
	return m.BACnetFaultParameterFaultOutOfRangeMaxNormalValue
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble) GetDoubleValue() *BACnetApplicationTagDouble {
	return m.DoubleValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble factory function for BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble
func NewBACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble(doubleValue *BACnetApplicationTagDouble, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble {
	_result := &BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble{
		DoubleValue: doubleValue,
		BACnetFaultParameterFaultOutOfRangeMaxNormalValue: NewBACnetFaultParameterFaultOutOfRangeMaxNormalValue(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble(structType interface{}) *BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble {
	if casted, ok := structType.(BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble); ok {
		return casted
	}
	if casted, ok := structType.(BACnetFaultParameterFaultOutOfRangeMaxNormalValue); ok {
		return CastBACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble(casted.Child)
	}
	if casted, ok := structType.(*BACnetFaultParameterFaultOutOfRangeMaxNormalValue); ok {
		return CastBACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble(casted.Child)
	}
	return nil
}

func (m *BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble) GetTypeName() string {
	return "BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble"
}

func (m *BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (doubleValue)
	lengthInBits += m.DoubleValue.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetFaultParameterFaultOutOfRangeMaxNormalValueDoubleParse(readBuffer utils.ReadBuffer, tagNumber uint8) (*BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (doubleValue)
	if pullErr := readBuffer.PullContext("doubleValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for doubleValue")
	}
	_doubleValue, _doubleValueErr := BACnetApplicationTagParse(readBuffer)
	if _doubleValueErr != nil {
		return nil, errors.Wrap(_doubleValueErr, "Error parsing 'doubleValue' field")
	}
	doubleValue := CastBACnetApplicationTagDouble(_doubleValue)
	if closeErr := readBuffer.CloseContext("doubleValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for doubleValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble")
	}

	// Create a partially initialized instance
	_child := &BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble{
		DoubleValue: CastBACnetApplicationTagDouble(doubleValue),
		BACnetFaultParameterFaultOutOfRangeMaxNormalValue: &BACnetFaultParameterFaultOutOfRangeMaxNormalValue{},
	}
	_child.BACnetFaultParameterFaultOutOfRangeMaxNormalValue.Child = _child
	return _child, nil
}

func (m *BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble")
		}

		// Simple Field (doubleValue)
		if pushErr := writeBuffer.PushContext("doubleValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for doubleValue")
		}
		_doubleValueErr := writeBuffer.WriteSerializable(m.DoubleValue)
		if popErr := writeBuffer.PopContext("doubleValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for doubleValue")
		}
		if _doubleValueErr != nil {
			return errors.Wrap(_doubleValueErr, "Error serializing 'doubleValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
