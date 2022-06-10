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

// BACnetApplicationTagDate is the data-structure of this message
type BACnetApplicationTagDate struct {
	*BACnetApplicationTag
	Payload *BACnetTagPayloadDate
}

// IBACnetApplicationTagDate is the corresponding interface of BACnetApplicationTagDate
type IBACnetApplicationTagDate interface {
	IBACnetApplicationTag
	// GetPayload returns Payload (property field)
	GetPayload() *BACnetTagPayloadDate
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

func (m *BACnetApplicationTagDate) InitializeParent(parent *BACnetApplicationTag, header *BACnetTagHeader) {
	m.BACnetApplicationTag.Header = header
}

func (m *BACnetApplicationTagDate) GetParent() *BACnetApplicationTag {
	return m.BACnetApplicationTag
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetApplicationTagDate) GetPayload() *BACnetTagPayloadDate {
	return m.Payload
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetApplicationTagDate factory function for BACnetApplicationTagDate
func NewBACnetApplicationTagDate(payload *BACnetTagPayloadDate, header *BACnetTagHeader) *BACnetApplicationTagDate {
	_result := &BACnetApplicationTagDate{
		Payload:              payload,
		BACnetApplicationTag: NewBACnetApplicationTag(header),
	}
	_result.Child = _result
	return _result
}

func CastBACnetApplicationTagDate(structType interface{}) *BACnetApplicationTagDate {
	if casted, ok := structType.(BACnetApplicationTagDate); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetApplicationTagDate); ok {
		return casted
	}
	if casted, ok := structType.(BACnetApplicationTag); ok {
		return CastBACnetApplicationTagDate(casted.Child)
	}
	if casted, ok := structType.(*BACnetApplicationTag); ok {
		return CastBACnetApplicationTagDate(casted.Child)
	}
	return nil
}

func (m *BACnetApplicationTagDate) GetTypeName() string {
	return "BACnetApplicationTagDate"
}

func (m *BACnetApplicationTagDate) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetApplicationTagDate) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (payload)
	lengthInBits += m.Payload.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetApplicationTagDate) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetApplicationTagDateParse(readBuffer utils.ReadBuffer) (*BACnetApplicationTagDate, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetApplicationTagDate"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetApplicationTagDate")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (payload)
	if pullErr := readBuffer.PullContext("payload"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for payload")
	}
	_payload, _payloadErr := BACnetTagPayloadDateParse(readBuffer)
	if _payloadErr != nil {
		return nil, errors.Wrap(_payloadErr, "Error parsing 'payload' field")
	}
	payload := CastBACnetTagPayloadDate(_payload)
	if closeErr := readBuffer.CloseContext("payload"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for payload")
	}

	if closeErr := readBuffer.CloseContext("BACnetApplicationTagDate"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetApplicationTagDate")
	}

	// Create a partially initialized instance
	_child := &BACnetApplicationTagDate{
		Payload:              CastBACnetTagPayloadDate(payload),
		BACnetApplicationTag: &BACnetApplicationTag{},
	}
	_child.BACnetApplicationTag.Child = _child
	return _child, nil
}

func (m *BACnetApplicationTagDate) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetApplicationTagDate"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetApplicationTagDate")
		}

		// Simple Field (payload)
		if pushErr := writeBuffer.PushContext("payload"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for payload")
		}
		_payloadErr := m.Payload.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("payload"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for payload")
		}
		if _payloadErr != nil {
			return errors.Wrap(_payloadErr, "Error serializing 'payload' field")
		}

		if popErr := writeBuffer.PopContext("BACnetApplicationTagDate"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetApplicationTagDate")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetApplicationTagDate) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
