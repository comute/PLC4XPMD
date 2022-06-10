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

// BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged is the data-structure of this message
type BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged struct {
	Header *BACnetTagHeader
	Value  BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriority

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

// IBACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged is the corresponding interface of BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged
type IBACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged interface {
	// GetHeader returns Header (property field)
	GetHeader() *BACnetTagHeader
	// GetValue returns Value (property field)
	GetValue() BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriority
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

func (m *BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged) GetHeader() *BACnetTagHeader {
	return m.Header
}

func (m *BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged) GetValue() BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriority {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged factory function for BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged
func NewBACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged(header *BACnetTagHeader, value BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriority, tagNumber uint8, tagClass TagClass) *BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged {
	return &BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged{Header: header, Value: value, TagNumber: tagNumber, TagClass: tagClass}
}

func CastBACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged(structType interface{}) *BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged {
	if casted, ok := structType.(BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged); ok {
		return casted
	}
	return nil
}

func (m *BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged) GetTypeName() string {
	return "BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged"
}

func (m *BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits()

	// Manual Field (value)
	lengthInBits += uint16(int32(m.GetHeader().GetActualLength()) * int32(int32(8)))

	return lengthInBits
}

func (m *BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTaggedParse(readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (*BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged")
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
	_value, _valueErr := ReadEnumGenericFailing(readBuffer, header.GetActualLength(), BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriority_NORMAL)
	if _valueErr != nil {
		return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
	}
	value := _value.(BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriority)

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged")
	}

	// Create the instance
	return NewBACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged(header, value, tagNumber, tagClass), nil
}

func (m *BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged")
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

	if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged")
	}
	return nil
}

func (m *BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
