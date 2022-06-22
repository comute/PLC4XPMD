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

// BACnetConstructedDataExpirationTime is the corresponding interface of BACnetConstructedDataExpirationTime
type BACnetConstructedDataExpirationTime interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetExpirationTime returns ExpirationTime (property field)
	GetExpirationTime() BACnetDateTime
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetDateTime
}

// BACnetConstructedDataExpirationTimeExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataExpirationTime.
// This is useful for switch cases.
type BACnetConstructedDataExpirationTimeExactly interface {
	BACnetConstructedDataExpirationTime
	isBACnetConstructedDataExpirationTime() bool
}

// _BACnetConstructedDataExpirationTime is the data-structure of this message
type _BACnetConstructedDataExpirationTime struct {
	*_BACnetConstructedData
	ExpirationTime BACnetDateTime

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument BACnetTagPayloadUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataExpirationTime) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataExpirationTime) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_EXPIRATION_TIME
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataExpirationTime) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataExpirationTime) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataExpirationTime) GetExpirationTime() BACnetDateTime {
	return m.ExpirationTime
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataExpirationTime) GetActualValue() BACnetDateTime {
	return CastBACnetDateTime(m.GetExpirationTime())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataExpirationTime factory function for _BACnetConstructedDataExpirationTime
func NewBACnetConstructedDataExpirationTime(expirationTime BACnetDateTime, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataExpirationTime {
	_result := &_BACnetConstructedDataExpirationTime{
		ExpirationTime:         expirationTime,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataExpirationTime(structType interface{}) BACnetConstructedDataExpirationTime {
	if casted, ok := structType.(BACnetConstructedDataExpirationTime); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataExpirationTime); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataExpirationTime) GetTypeName() string {
	return "BACnetConstructedDataExpirationTime"
}

func (m *_BACnetConstructedDataExpirationTime) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataExpirationTime) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (expirationTime)
	lengthInBits += m.ExpirationTime.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataExpirationTime) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataExpirationTimeParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataExpirationTime, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataExpirationTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataExpirationTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (expirationTime)
	if pullErr := readBuffer.PullContext("expirationTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for expirationTime")
	}
	_expirationTime, _expirationTimeErr := BACnetDateTimeParse(readBuffer)
	if _expirationTimeErr != nil {
		return nil, errors.Wrap(_expirationTimeErr, "Error parsing 'expirationTime' field")
	}
	expirationTime := _expirationTime.(BACnetDateTime)
	if closeErr := readBuffer.CloseContext("expirationTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for expirationTime")
	}

	// Virtual field
	_actualValue := expirationTime
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataExpirationTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataExpirationTime")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataExpirationTime{
		ExpirationTime:         expirationTime,
		_BACnetConstructedData: &_BACnetConstructedData{},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataExpirationTime) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataExpirationTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataExpirationTime")
		}

		// Simple Field (expirationTime)
		if pushErr := writeBuffer.PushContext("expirationTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for expirationTime")
		}
		_expirationTimeErr := writeBuffer.WriteSerializable(m.GetExpirationTime())
		if popErr := writeBuffer.PopContext("expirationTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for expirationTime")
		}
		if _expirationTimeErr != nil {
			return errors.Wrap(_expirationTimeErr, "Error serializing 'expirationTime' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataExpirationTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataExpirationTime")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataExpirationTime) isBACnetConstructedDataExpirationTime() bool {
	return true
}

func (m *_BACnetConstructedDataExpirationTime) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
