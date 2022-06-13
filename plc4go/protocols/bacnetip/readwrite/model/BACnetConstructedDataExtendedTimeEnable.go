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

// BACnetConstructedDataExtendedTimeEnable is the data-structure of this message
type BACnetConstructedDataExtendedTimeEnable struct {
	*BACnetConstructedData
	ExtendedTimeEnable *BACnetApplicationTagBoolean

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataExtendedTimeEnable is the corresponding interface of BACnetConstructedDataExtendedTimeEnable
type IBACnetConstructedDataExtendedTimeEnable interface {
	IBACnetConstructedData
	// GetExtendedTimeEnable returns ExtendedTimeEnable (property field)
	GetExtendedTimeEnable() *BACnetApplicationTagBoolean
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

func (m *BACnetConstructedDataExtendedTimeEnable) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataExtendedTimeEnable) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_EXTENDED_TIME_ENABLE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataExtendedTimeEnable) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataExtendedTimeEnable) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataExtendedTimeEnable) GetExtendedTimeEnable() *BACnetApplicationTagBoolean {
	return m.ExtendedTimeEnable
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataExtendedTimeEnable factory function for BACnetConstructedDataExtendedTimeEnable
func NewBACnetConstructedDataExtendedTimeEnable(extendedTimeEnable *BACnetApplicationTagBoolean, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataExtendedTimeEnable {
	_result := &BACnetConstructedDataExtendedTimeEnable{
		ExtendedTimeEnable:    extendedTimeEnable,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataExtendedTimeEnable(structType interface{}) *BACnetConstructedDataExtendedTimeEnable {
	if casted, ok := structType.(BACnetConstructedDataExtendedTimeEnable); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataExtendedTimeEnable); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataExtendedTimeEnable(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataExtendedTimeEnable(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataExtendedTimeEnable) GetTypeName() string {
	return "BACnetConstructedDataExtendedTimeEnable"
}

func (m *BACnetConstructedDataExtendedTimeEnable) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataExtendedTimeEnable) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (extendedTimeEnable)
	lengthInBits += m.ExtendedTimeEnable.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataExtendedTimeEnable) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataExtendedTimeEnableParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataExtendedTimeEnable, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataExtendedTimeEnable"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataExtendedTimeEnable")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (extendedTimeEnable)
	if pullErr := readBuffer.PullContext("extendedTimeEnable"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for extendedTimeEnable")
	}
	_extendedTimeEnable, _extendedTimeEnableErr := BACnetApplicationTagParse(readBuffer)
	if _extendedTimeEnableErr != nil {
		return nil, errors.Wrap(_extendedTimeEnableErr, "Error parsing 'extendedTimeEnable' field")
	}
	extendedTimeEnable := CastBACnetApplicationTagBoolean(_extendedTimeEnable)
	if closeErr := readBuffer.CloseContext("extendedTimeEnable"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for extendedTimeEnable")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataExtendedTimeEnable"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataExtendedTimeEnable")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataExtendedTimeEnable{
		ExtendedTimeEnable:    CastBACnetApplicationTagBoolean(extendedTimeEnable),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataExtendedTimeEnable) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataExtendedTimeEnable"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataExtendedTimeEnable")
		}

		// Simple Field (extendedTimeEnable)
		if pushErr := writeBuffer.PushContext("extendedTimeEnable"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for extendedTimeEnable")
		}
		_extendedTimeEnableErr := writeBuffer.WriteSerializable(m.ExtendedTimeEnable)
		if popErr := writeBuffer.PopContext("extendedTimeEnable"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for extendedTimeEnable")
		}
		if _extendedTimeEnableErr != nil {
			return errors.Wrap(_extendedTimeEnableErr, "Error serializing 'extendedTimeEnable' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataExtendedTimeEnable"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataExtendedTimeEnable")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataExtendedTimeEnable) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
