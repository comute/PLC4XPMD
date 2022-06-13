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

// BACnetTimerStateTagged is the data-structure of this message
type BACnetTimerStateTagged struct {
	Header *BACnetTagHeader
	Value  BACnetTimerState

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

// IBACnetTimerStateTagged is the corresponding interface of BACnetTimerStateTagged
type IBACnetTimerStateTagged interface {
	// GetHeader returns Header (property field)
	GetHeader() *BACnetTagHeader
	// GetValue returns Value (property field)
	GetValue() BACnetTimerState
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

func (m *BACnetTimerStateTagged) GetHeader() *BACnetTagHeader {
	return m.Header
}

func (m *BACnetTimerStateTagged) GetValue() BACnetTimerState {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetTimerStateTagged factory function for BACnetTimerStateTagged
func NewBACnetTimerStateTagged(header *BACnetTagHeader, value BACnetTimerState, tagNumber uint8, tagClass TagClass) *BACnetTimerStateTagged {
	return &BACnetTimerStateTagged{Header: header, Value: value, TagNumber: tagNumber, TagClass: tagClass}
}

func CastBACnetTimerStateTagged(structType interface{}) *BACnetTimerStateTagged {
	if casted, ok := structType.(BACnetTimerStateTagged); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetTimerStateTagged); ok {
		return casted
	}
	return nil
}

func (m *BACnetTimerStateTagged) GetTypeName() string {
	return "BACnetTimerStateTagged"
}

func (m *BACnetTimerStateTagged) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetTimerStateTagged) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits()

	// Manual Field (value)
	lengthInBits += uint16(int32(m.GetHeader().GetActualLength()) * int32(int32(8)))

	return lengthInBits
}

func (m *BACnetTimerStateTagged) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetTimerStateTaggedParse(readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (*BACnetTimerStateTagged, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetTimerStateTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetTimerStateTagged")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (header)
	if pullErr := readBuffer.PullContext("header"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for header")
	}
	_header, _headerErr := BACnetTagHeaderParse(readBuffer)
	if _headerErr != nil {
		return nil, errors.Wrap(_headerErr, "Error parsing 'header' field")
	}
	header := CastBACnetTagHeader(_header)
	if closeErr := readBuffer.CloseContext("header"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for header")
	}

	// Validation
	if !(bool((header.GetTagClass()) == (tagClass))) {
		return nil, errors.WithStack(utils.ParseValidationError{"tag class doesn't match"})
	}

	// Validation
	if !(bool(bool(bool((header.GetTagClass()) == (TagClass_APPLICATION_TAGS)))) || bool(bool(bool((header.GetActualTagNumber()) == (tagNumber))))) {
		return nil, errors.WithStack(utils.ParseAssertError{"tagnumber doesn't match"})
	}

	// Manual Field (value)
	_value, _valueErr := ReadEnumGenericFailing(readBuffer, header.GetActualLength(), BACnetTimerState_IDLE)
	if _valueErr != nil {
		return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
	}
	value := _value.(BACnetTimerState)

	if closeErr := readBuffer.CloseContext("BACnetTimerStateTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetTimerStateTagged")
	}

	// Create the instance
	return NewBACnetTimerStateTagged(header, value, tagNumber, tagClass), nil
}

func (m *BACnetTimerStateTagged) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetTimerStateTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetTimerStateTagged")
	}

	// Simple Field (header)
	if pushErr := writeBuffer.PushContext("header"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for header")
	}
	_headerErr := writeBuffer.WriteSerializable(m.Header)
	if popErr := writeBuffer.PopContext("header"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for header")
	}
	if _headerErr != nil {
		return errors.Wrap(_headerErr, "Error serializing 'header' field")
	}

	// Manual Field (value)
	_valueErr := WriteEnumGeneric(writeBuffer, m.GetValue())
	if _valueErr != nil {
		return errors.Wrap(_valueErr, "Error serializing 'value' field")
	}

	if popErr := writeBuffer.PopContext("BACnetTimerStateTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetTimerStateTagged")
	}
	return nil
}

func (m *BACnetTimerStateTagged) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
