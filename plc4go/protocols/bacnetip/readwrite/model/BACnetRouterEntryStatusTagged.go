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

// BACnetRouterEntryStatusTagged is the data-structure of this message
type BACnetRouterEntryStatusTagged struct {
	Header *BACnetTagHeader
	Value  BACnetRouterEntryStatus

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

// IBACnetRouterEntryStatusTagged is the corresponding interface of BACnetRouterEntryStatusTagged
type IBACnetRouterEntryStatusTagged interface {
	// GetHeader returns Header (property field)
	GetHeader() *BACnetTagHeader
	// GetValue returns Value (property field)
	GetValue() BACnetRouterEntryStatus
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

func (m *BACnetRouterEntryStatusTagged) GetHeader() *BACnetTagHeader {
	return m.Header
}

func (m *BACnetRouterEntryStatusTagged) GetValue() BACnetRouterEntryStatus {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetRouterEntryStatusTagged factory function for BACnetRouterEntryStatusTagged
func NewBACnetRouterEntryStatusTagged(header *BACnetTagHeader, value BACnetRouterEntryStatus, tagNumber uint8, tagClass TagClass) *BACnetRouterEntryStatusTagged {
	return &BACnetRouterEntryStatusTagged{Header: header, Value: value, TagNumber: tagNumber, TagClass: tagClass}
}

func CastBACnetRouterEntryStatusTagged(structType interface{}) *BACnetRouterEntryStatusTagged {
	if casted, ok := structType.(BACnetRouterEntryStatusTagged); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetRouterEntryStatusTagged); ok {
		return casted
	}
	return nil
}

func (m *BACnetRouterEntryStatusTagged) GetTypeName() string {
	return "BACnetRouterEntryStatusTagged"
}

func (m *BACnetRouterEntryStatusTagged) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetRouterEntryStatusTagged) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits()

	// Manual Field (value)
	lengthInBits += uint16(int32(m.GetHeader().GetActualLength()) * int32(int32(8)))

	return lengthInBits
}

func (m *BACnetRouterEntryStatusTagged) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetRouterEntryStatusTaggedParse(readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (*BACnetRouterEntryStatusTagged, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetRouterEntryStatusTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetRouterEntryStatusTagged")
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
	_value, _valueErr := ReadEnumGenericFailing(readBuffer, header.GetActualLength(), BACnetRouterEntryStatus_AVAILABLE)
	if _valueErr != nil {
		return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
	}
	value := _value.(BACnetRouterEntryStatus)

	if closeErr := readBuffer.CloseContext("BACnetRouterEntryStatusTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetRouterEntryStatusTagged")
	}

	// Create the instance
	return NewBACnetRouterEntryStatusTagged(header, value, tagNumber, tagClass), nil
}

func (m *BACnetRouterEntryStatusTagged) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetRouterEntryStatusTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetRouterEntryStatusTagged")
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

	if popErr := writeBuffer.PopContext("BACnetRouterEntryStatusTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetRouterEntryStatusTagged")
	}
	return nil
}

func (m *BACnetRouterEntryStatusTagged) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
