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

// BACnetProgramStateTagged is the data-structure of this message
type BACnetProgramStateTagged struct {
	Header *BACnetTagHeader
	Value  BACnetProgramState

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

// IBACnetProgramStateTagged is the corresponding interface of BACnetProgramStateTagged
type IBACnetProgramStateTagged interface {
	// GetHeader returns Header (property field)
	GetHeader() *BACnetTagHeader
	// GetValue returns Value (property field)
	GetValue() BACnetProgramState
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

func (m *BACnetProgramStateTagged) GetHeader() *BACnetTagHeader {
	return m.Header
}

func (m *BACnetProgramStateTagged) GetValue() BACnetProgramState {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetProgramStateTagged factory function for BACnetProgramStateTagged
func NewBACnetProgramStateTagged(header *BACnetTagHeader, value BACnetProgramState, tagNumber uint8, tagClass TagClass) *BACnetProgramStateTagged {
	return &BACnetProgramStateTagged{Header: header, Value: value, TagNumber: tagNumber, TagClass: tagClass}
}

func CastBACnetProgramStateTagged(structType interface{}) *BACnetProgramStateTagged {
	if casted, ok := structType.(BACnetProgramStateTagged); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetProgramStateTagged); ok {
		return casted
	}
	return nil
}

func (m *BACnetProgramStateTagged) GetTypeName() string {
	return "BACnetProgramStateTagged"
}

func (m *BACnetProgramStateTagged) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetProgramStateTagged) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits()

	// Manual Field (value)
	lengthInBits += uint16(int32(m.GetHeader().GetActualLength()) * int32(int32(8)))

	return lengthInBits
}

func (m *BACnetProgramStateTagged) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetProgramStateTaggedParse(readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (*BACnetProgramStateTagged, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetProgramStateTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetProgramStateTagged")
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
	_value, _valueErr := ReadEnumGenericFailing(readBuffer, header.GetActualLength(), BACnetProgramState_IDLE)
	if _valueErr != nil {
		return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
	}
	value := _value.(BACnetProgramState)

	if closeErr := readBuffer.CloseContext("BACnetProgramStateTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetProgramStateTagged")
	}

	// Create the instance
	return NewBACnetProgramStateTagged(header, value, tagNumber, tagClass), nil
}

func (m *BACnetProgramStateTagged) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetProgramStateTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetProgramStateTagged")
	}

	// Simple Field (header)
	if pushErr := writeBuffer.PushContext("header"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for header")
	}
	_headerErr := m.Header.Serialize(writeBuffer)
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

	if popErr := writeBuffer.PopContext("BACnetProgramStateTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetProgramStateTagged")
	}
	return nil
}

func (m *BACnetProgramStateTagged) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
