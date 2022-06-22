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

// BACnetConstructedDataUTCOffset is the corresponding interface of BACnetConstructedDataUTCOffset
type BACnetConstructedDataUTCOffset interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetUtcOffset returns UtcOffset (property field)
	GetUtcOffset() BACnetApplicationTagSignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagSignedInteger
}

// BACnetConstructedDataUTCOffsetExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataUTCOffset.
// This is useful for switch cases.
type BACnetConstructedDataUTCOffsetExactly interface {
	BACnetConstructedDataUTCOffset
	isBACnetConstructedDataUTCOffset() bool
}

// _BACnetConstructedDataUTCOffset is the data-structure of this message
type _BACnetConstructedDataUTCOffset struct {
	*_BACnetConstructedData
	UtcOffset BACnetApplicationTagSignedInteger

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument BACnetTagPayloadUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataUTCOffset) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataUTCOffset) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_UTC_OFFSET
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataUTCOffset) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataUTCOffset) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataUTCOffset) GetUtcOffset() BACnetApplicationTagSignedInteger {
	return m.UtcOffset
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataUTCOffset) GetActualValue() BACnetApplicationTagSignedInteger {
	return CastBACnetApplicationTagSignedInteger(m.GetUtcOffset())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataUTCOffset factory function for _BACnetConstructedDataUTCOffset
func NewBACnetConstructedDataUTCOffset(utcOffset BACnetApplicationTagSignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataUTCOffset {
	_result := &_BACnetConstructedDataUTCOffset{
		UtcOffset:              utcOffset,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataUTCOffset(structType interface{}) BACnetConstructedDataUTCOffset {
	if casted, ok := structType.(BACnetConstructedDataUTCOffset); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataUTCOffset); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataUTCOffset) GetTypeName() string {
	return "BACnetConstructedDataUTCOffset"
}

func (m *_BACnetConstructedDataUTCOffset) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataUTCOffset) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (utcOffset)
	lengthInBits += m.UtcOffset.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataUTCOffset) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataUTCOffsetParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataUTCOffset, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataUTCOffset"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataUTCOffset")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (utcOffset)
	if pullErr := readBuffer.PullContext("utcOffset"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for utcOffset")
	}
	_utcOffset, _utcOffsetErr := BACnetApplicationTagParse(readBuffer)
	if _utcOffsetErr != nil {
		return nil, errors.Wrap(_utcOffsetErr, "Error parsing 'utcOffset' field")
	}
	utcOffset := _utcOffset.(BACnetApplicationTagSignedInteger)
	if closeErr := readBuffer.CloseContext("utcOffset"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for utcOffset")
	}

	// Virtual field
	_actualValue := utcOffset
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataUTCOffset"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataUTCOffset")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataUTCOffset{
		UtcOffset:              utcOffset,
		_BACnetConstructedData: &_BACnetConstructedData{},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataUTCOffset) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataUTCOffset"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataUTCOffset")
		}

		// Simple Field (utcOffset)
		if pushErr := writeBuffer.PushContext("utcOffset"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for utcOffset")
		}
		_utcOffsetErr := writeBuffer.WriteSerializable(m.GetUtcOffset())
		if popErr := writeBuffer.PopContext("utcOffset"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for utcOffset")
		}
		if _utcOffsetErr != nil {
			return errors.Wrap(_utcOffsetErr, "Error serializing 'utcOffset' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataUTCOffset"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataUTCOffset")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataUTCOffset) isBACnetConstructedDataUTCOffset() bool {
	return true
}

func (m *_BACnetConstructedDataUTCOffset) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
