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

// BACnetConstructedDataCountChangeTime is the data-structure of this message
type BACnetConstructedDataCountChangeTime struct {
	*BACnetConstructedData
	CountChangeTime *BACnetDateTime

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataCountChangeTime is the corresponding interface of BACnetConstructedDataCountChangeTime
type IBACnetConstructedDataCountChangeTime interface {
	IBACnetConstructedData
	// GetCountChangeTime returns CountChangeTime (property field)
	GetCountChangeTime() *BACnetDateTime
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

func (m *BACnetConstructedDataCountChangeTime) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataCountChangeTime) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_COUNT_CHANGE_TIME
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataCountChangeTime) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataCountChangeTime) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataCountChangeTime) GetCountChangeTime() *BACnetDateTime {
	return m.CountChangeTime
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataCountChangeTime factory function for BACnetConstructedDataCountChangeTime
func NewBACnetConstructedDataCountChangeTime(countChangeTime *BACnetDateTime, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataCountChangeTime {
	_result := &BACnetConstructedDataCountChangeTime{
		CountChangeTime:       countChangeTime,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataCountChangeTime(structType interface{}) *BACnetConstructedDataCountChangeTime {
	if casted, ok := structType.(BACnetConstructedDataCountChangeTime); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataCountChangeTime); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataCountChangeTime(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataCountChangeTime(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataCountChangeTime) GetTypeName() string {
	return "BACnetConstructedDataCountChangeTime"
}

func (m *BACnetConstructedDataCountChangeTime) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataCountChangeTime) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (countChangeTime)
	lengthInBits += m.CountChangeTime.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataCountChangeTime) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataCountChangeTimeParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataCountChangeTime, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataCountChangeTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataCountChangeTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (countChangeTime)
	if pullErr := readBuffer.PullContext("countChangeTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for countChangeTime")
	}
	_countChangeTime, _countChangeTimeErr := BACnetDateTimeParse(readBuffer)
	if _countChangeTimeErr != nil {
		return nil, errors.Wrap(_countChangeTimeErr, "Error parsing 'countChangeTime' field")
	}
	countChangeTime := CastBACnetDateTime(_countChangeTime)
	if closeErr := readBuffer.CloseContext("countChangeTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for countChangeTime")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataCountChangeTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataCountChangeTime")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataCountChangeTime{
		CountChangeTime:       CastBACnetDateTime(countChangeTime),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataCountChangeTime) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataCountChangeTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataCountChangeTime")
		}

		// Simple Field (countChangeTime)
		if pushErr := writeBuffer.PushContext("countChangeTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for countChangeTime")
		}
		_countChangeTimeErr := m.CountChangeTime.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("countChangeTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for countChangeTime")
		}
		if _countChangeTimeErr != nil {
			return errors.Wrap(_countChangeTimeErr, "Error serializing 'countChangeTime' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataCountChangeTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataCountChangeTime")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataCountChangeTime) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
