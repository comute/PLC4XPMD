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

// BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger is the data-structure of this message
type BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger struct {
	*BACnetFaultParameterFaultOutOfRangeMinNormalValue
	IntegerValue *BACnetApplicationTagSignedInteger

	// Arguments.
	TagNumber uint8
}

// IBACnetFaultParameterFaultOutOfRangeMinNormalValueInteger is the corresponding interface of BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger
type IBACnetFaultParameterFaultOutOfRangeMinNormalValueInteger interface {
	IBACnetFaultParameterFaultOutOfRangeMinNormalValue
	// GetIntegerValue returns IntegerValue (property field)
	GetIntegerValue() *BACnetApplicationTagSignedInteger
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

func (m *BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger) InitializeParent(parent *BACnetFaultParameterFaultOutOfRangeMinNormalValue, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetFaultParameterFaultOutOfRangeMinNormalValue.OpeningTag = openingTag
	m.BACnetFaultParameterFaultOutOfRangeMinNormalValue.PeekedTagHeader = peekedTagHeader
	m.BACnetFaultParameterFaultOutOfRangeMinNormalValue.ClosingTag = closingTag
}

func (m *BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger) GetParent() *BACnetFaultParameterFaultOutOfRangeMinNormalValue {
	return m.BACnetFaultParameterFaultOutOfRangeMinNormalValue
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger) GetIntegerValue() *BACnetApplicationTagSignedInteger {
	return m.IntegerValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetFaultParameterFaultOutOfRangeMinNormalValueInteger factory function for BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger
func NewBACnetFaultParameterFaultOutOfRangeMinNormalValueInteger(integerValue *BACnetApplicationTagSignedInteger, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger {
	_result := &BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger{
		IntegerValue: integerValue,
		BACnetFaultParameterFaultOutOfRangeMinNormalValue: NewBACnetFaultParameterFaultOutOfRangeMinNormalValue(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetFaultParameterFaultOutOfRangeMinNormalValueInteger(structType interface{}) *BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger {
	if casted, ok := structType.(BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger); ok {
		return casted
	}
	if casted, ok := structType.(BACnetFaultParameterFaultOutOfRangeMinNormalValue); ok {
		return CastBACnetFaultParameterFaultOutOfRangeMinNormalValueInteger(casted.Child)
	}
	if casted, ok := structType.(*BACnetFaultParameterFaultOutOfRangeMinNormalValue); ok {
		return CastBACnetFaultParameterFaultOutOfRangeMinNormalValueInteger(casted.Child)
	}
	return nil
}

func (m *BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger) GetTypeName() string {
	return "BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger"
}

func (m *BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (integerValue)
	lengthInBits += m.IntegerValue.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetFaultParameterFaultOutOfRangeMinNormalValueIntegerParse(readBuffer utils.ReadBuffer, tagNumber uint8) (*BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (integerValue)
	if pullErr := readBuffer.PullContext("integerValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for integerValue")
	}
	_integerValue, _integerValueErr := BACnetApplicationTagParse(readBuffer)
	if _integerValueErr != nil {
		return nil, errors.Wrap(_integerValueErr, "Error parsing 'integerValue' field")
	}
	integerValue := CastBACnetApplicationTagSignedInteger(_integerValue)
	if closeErr := readBuffer.CloseContext("integerValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for integerValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger")
	}

	// Create a partially initialized instance
	_child := &BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger{
		IntegerValue: CastBACnetApplicationTagSignedInteger(integerValue),
		BACnetFaultParameterFaultOutOfRangeMinNormalValue: &BACnetFaultParameterFaultOutOfRangeMinNormalValue{},
	}
	_child.BACnetFaultParameterFaultOutOfRangeMinNormalValue.Child = _child
	return _child, nil
}

func (m *BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger")
		}

		// Simple Field (integerValue)
		if pushErr := writeBuffer.PushContext("integerValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for integerValue")
		}
		_integerValueErr := writeBuffer.WriteSerializable(m.IntegerValue)
		if popErr := writeBuffer.PopContext("integerValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for integerValue")
		}
		if _integerValueErr != nil {
			return errors.Wrap(_integerValueErr, "Error serializing 'integerValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
