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

// BACnetTimeStampDateTime is the data-structure of this message
type BACnetTimeStampDateTime struct {
	*BACnetTimeStamp
	DateTimeValue *BACnetDateTimeEnclosed
}

// IBACnetTimeStampDateTime is the corresponding interface of BACnetTimeStampDateTime
type IBACnetTimeStampDateTime interface {
	IBACnetTimeStamp
	// GetDateTimeValue returns DateTimeValue (property field)
	GetDateTimeValue() *BACnetDateTimeEnclosed
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

func (m *BACnetTimeStampDateTime) InitializeParent(parent *BACnetTimeStamp, peekedTagHeader *BACnetTagHeader) {
	m.BACnetTimeStamp.PeekedTagHeader = peekedTagHeader
}

func (m *BACnetTimeStampDateTime) GetParent() *BACnetTimeStamp {
	return m.BACnetTimeStamp
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetTimeStampDateTime) GetDateTimeValue() *BACnetDateTimeEnclosed {
	return m.DateTimeValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetTimeStampDateTime factory function for BACnetTimeStampDateTime
func NewBACnetTimeStampDateTime(dateTimeValue *BACnetDateTimeEnclosed, peekedTagHeader *BACnetTagHeader) *BACnetTimeStampDateTime {
	_result := &BACnetTimeStampDateTime{
		DateTimeValue:   dateTimeValue,
		BACnetTimeStamp: NewBACnetTimeStamp(peekedTagHeader),
	}
	_result.Child = _result
	return _result
}

func CastBACnetTimeStampDateTime(structType interface{}) *BACnetTimeStampDateTime {
	if casted, ok := structType.(BACnetTimeStampDateTime); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetTimeStampDateTime); ok {
		return casted
	}
	if casted, ok := structType.(BACnetTimeStamp); ok {
		return CastBACnetTimeStampDateTime(casted.Child)
	}
	if casted, ok := structType.(*BACnetTimeStamp); ok {
		return CastBACnetTimeStampDateTime(casted.Child)
	}
	return nil
}

func (m *BACnetTimeStampDateTime) GetTypeName() string {
	return "BACnetTimeStampDateTime"
}

func (m *BACnetTimeStampDateTime) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetTimeStampDateTime) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (dateTimeValue)
	lengthInBits += m.DateTimeValue.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetTimeStampDateTime) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetTimeStampDateTimeParse(readBuffer utils.ReadBuffer) (*BACnetTimeStampDateTime, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetTimeStampDateTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetTimeStampDateTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (dateTimeValue)
	if pullErr := readBuffer.PullContext("dateTimeValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for dateTimeValue")
	}
	_dateTimeValue, _dateTimeValueErr := BACnetDateTimeEnclosedParse(readBuffer, uint8(uint8(2)))
	if _dateTimeValueErr != nil {
		return nil, errors.Wrap(_dateTimeValueErr, "Error parsing 'dateTimeValue' field")
	}
	dateTimeValue := CastBACnetDateTimeEnclosed(_dateTimeValue)
	if closeErr := readBuffer.CloseContext("dateTimeValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for dateTimeValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetTimeStampDateTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetTimeStampDateTime")
	}

	// Create a partially initialized instance
	_child := &BACnetTimeStampDateTime{
		DateTimeValue:   CastBACnetDateTimeEnclosed(dateTimeValue),
		BACnetTimeStamp: &BACnetTimeStamp{},
	}
	_child.BACnetTimeStamp.Child = _child
	return _child, nil
}

func (m *BACnetTimeStampDateTime) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetTimeStampDateTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetTimeStampDateTime")
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

		if popErr := writeBuffer.PopContext("BACnetTimeStampDateTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetTimeStampDateTime")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetTimeStampDateTime) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
