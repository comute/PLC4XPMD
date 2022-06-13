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

// BACnetConstructedDataMinimumOffTime is the data-structure of this message
type BACnetConstructedDataMinimumOffTime struct {
	*BACnetConstructedData
	MinimumOffTime *BACnetApplicationTagUnsignedInteger

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataMinimumOffTime is the corresponding interface of BACnetConstructedDataMinimumOffTime
type IBACnetConstructedDataMinimumOffTime interface {
	IBACnetConstructedData
	// GetMinimumOffTime returns MinimumOffTime (property field)
	GetMinimumOffTime() *BACnetApplicationTagUnsignedInteger
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

func (m *BACnetConstructedDataMinimumOffTime) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataMinimumOffTime) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_MINIMUM_OFF_TIME
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataMinimumOffTime) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataMinimumOffTime) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataMinimumOffTime) GetMinimumOffTime() *BACnetApplicationTagUnsignedInteger {
	return m.MinimumOffTime
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataMinimumOffTime factory function for BACnetConstructedDataMinimumOffTime
func NewBACnetConstructedDataMinimumOffTime(minimumOffTime *BACnetApplicationTagUnsignedInteger, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataMinimumOffTime {
	_result := &BACnetConstructedDataMinimumOffTime{
		MinimumOffTime:        minimumOffTime,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataMinimumOffTime(structType interface{}) *BACnetConstructedDataMinimumOffTime {
	if casted, ok := structType.(BACnetConstructedDataMinimumOffTime); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataMinimumOffTime); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataMinimumOffTime(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataMinimumOffTime(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataMinimumOffTime) GetTypeName() string {
	return "BACnetConstructedDataMinimumOffTime"
}

func (m *BACnetConstructedDataMinimumOffTime) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataMinimumOffTime) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (minimumOffTime)
	lengthInBits += m.MinimumOffTime.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataMinimumOffTime) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataMinimumOffTimeParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataMinimumOffTime, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataMinimumOffTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataMinimumOffTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (minimumOffTime)
	if pullErr := readBuffer.PullContext("minimumOffTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for minimumOffTime")
	}
	_minimumOffTime, _minimumOffTimeErr := BACnetApplicationTagParse(readBuffer)
	if _minimumOffTimeErr != nil {
		return nil, errors.Wrap(_minimumOffTimeErr, "Error parsing 'minimumOffTime' field")
	}
	minimumOffTime := CastBACnetApplicationTagUnsignedInteger(_minimumOffTime)
	if closeErr := readBuffer.CloseContext("minimumOffTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for minimumOffTime")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataMinimumOffTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataMinimumOffTime")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataMinimumOffTime{
		MinimumOffTime:        CastBACnetApplicationTagUnsignedInteger(minimumOffTime),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataMinimumOffTime) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataMinimumOffTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataMinimumOffTime")
		}

		// Simple Field (minimumOffTime)
		if pushErr := writeBuffer.PushContext("minimumOffTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for minimumOffTime")
		}
		_minimumOffTimeErr := writeBuffer.WriteSerializable(m.MinimumOffTime)
		if popErr := writeBuffer.PopContext("minimumOffTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for minimumOffTime")
		}
		if _minimumOffTimeErr != nil {
			return errors.Wrap(_minimumOffTimeErr, "Error serializing 'minimumOffTime' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataMinimumOffTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataMinimumOffTime")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataMinimumOffTime) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
