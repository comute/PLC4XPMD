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

// BACnetDateTimeEnclosed is the data-structure of this message
type BACnetDateTimeEnclosed struct {
	OpeningTag    *BACnetOpeningTag
	DateTimeValue *BACnetDateTime
	ClosingTag    *BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

// IBACnetDateTimeEnclosed is the corresponding interface of BACnetDateTimeEnclosed
type IBACnetDateTimeEnclosed interface {
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() *BACnetOpeningTag
	// GetDateTimeValue returns DateTimeValue (property field)
	GetDateTimeValue() *BACnetDateTime
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() *BACnetClosingTag
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetDateTimeEnclosed) GetOpeningTag() *BACnetOpeningTag {
	return m.OpeningTag
}

func (m *BACnetDateTimeEnclosed) GetDateTimeValue() *BACnetDateTime {
	return m.DateTimeValue
}

func (m *BACnetDateTimeEnclosed) GetClosingTag() *BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetDateTimeEnclosed factory function for BACnetDateTimeEnclosed
func NewBACnetDateTimeEnclosed(openingTag *BACnetOpeningTag, dateTimeValue *BACnetDateTime, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetDateTimeEnclosed {
	return &BACnetDateTimeEnclosed{OpeningTag: openingTag, DateTimeValue: dateTimeValue, ClosingTag: closingTag, TagNumber: tagNumber}
}

func CastBACnetDateTimeEnclosed(structType interface{}) *BACnetDateTimeEnclosed {
	if casted, ok := structType.(BACnetDateTimeEnclosed); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetDateTimeEnclosed); ok {
		return casted
	}
	return nil
}

func (m *BACnetDateTimeEnclosed) GetTypeName() string {
	return "BACnetDateTimeEnclosed"
}

func (m *BACnetDateTimeEnclosed) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetDateTimeEnclosed) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits()

	// Simple field (dateTimeValue)
	lengthInBits += m.DateTimeValue.GetLengthInBits()

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetDateTimeEnclosed) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetDateTimeEnclosedParse(readBuffer utils.ReadBuffer, tagNumber uint8) (*BACnetDateTimeEnclosed, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetDateTimeEnclosed"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetDateTimeEnclosed")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for openingTag")
	}
	_openingTag, _openingTagErr := BACnetOpeningTagParse(readBuffer, uint8(tagNumber))
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field")
	}
	openingTag := CastBACnetOpeningTag(_openingTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for openingTag")
	}

	// Simple Field (dateTimeValue)
	if pullErr := readBuffer.PullContext("dateTimeValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for dateTimeValue")
	}
	_dateTimeValue, _dateTimeValueErr := BACnetDateTimeParse(readBuffer)
	if _dateTimeValueErr != nil {
		return nil, errors.Wrap(_dateTimeValueErr, "Error parsing 'dateTimeValue' field")
	}
	dateTimeValue := CastBACnetDateTime(_dateTimeValue)
	if closeErr := readBuffer.CloseContext("dateTimeValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for dateTimeValue")
	}

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for closingTag")
	}
	_closingTag, _closingTagErr := BACnetClosingTagParse(readBuffer, uint8(tagNumber))
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field")
	}
	closingTag := CastBACnetClosingTag(_closingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for closingTag")
	}

	if closeErr := readBuffer.CloseContext("BACnetDateTimeEnclosed"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetDateTimeEnclosed")
	}

	// Create the instance
	return NewBACnetDateTimeEnclosed(openingTag, dateTimeValue, closingTag, tagNumber), nil
}

func (m *BACnetDateTimeEnclosed) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetDateTimeEnclosed"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetDateTimeEnclosed")
	}

	// Simple Field (openingTag)
	if pushErr := writeBuffer.PushContext("openingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for openingTag")
	}
	_openingTagErr := m.OpeningTag.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("openingTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for openingTag")
	}
	if _openingTagErr != nil {
		return errors.Wrap(_openingTagErr, "Error serializing 'openingTag' field")
	}

	// Simple Field (dateTimeValue)
	if pushErr := writeBuffer.PushContext("dateTimeValue"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for dateTimeValue")
	}
	_dateTimeValueErr := m.DateTimeValue.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("dateTimeValue"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for dateTimeValue")
	}
	if _dateTimeValueErr != nil {
		return errors.Wrap(_dateTimeValueErr, "Error serializing 'dateTimeValue' field")
	}

	// Simple Field (closingTag)
	if pushErr := writeBuffer.PushContext("closingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for closingTag")
	}
	_closingTagErr := m.ClosingTag.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("closingTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for closingTag")
	}
	if _closingTagErr != nil {
		return errors.Wrap(_closingTagErr, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetDateTimeEnclosed"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetDateTimeEnclosed")
	}
	return nil
}

func (m *BACnetDateTimeEnclosed) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
