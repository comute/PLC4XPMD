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

// BACnetConstructedDataDescriptionOfHalt is the data-structure of this message
type BACnetConstructedDataDescriptionOfHalt struct {
	*BACnetConstructedData
	DescriptionForHalt *BACnetApplicationTagCharacterString

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataDescriptionOfHalt is the corresponding interface of BACnetConstructedDataDescriptionOfHalt
type IBACnetConstructedDataDescriptionOfHalt interface {
	IBACnetConstructedData
	// GetDescriptionForHalt returns DescriptionForHalt (property field)
	GetDescriptionForHalt() *BACnetApplicationTagCharacterString
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

func (m *BACnetConstructedDataDescriptionOfHalt) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataDescriptionOfHalt) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_DESCRIPTION_OF_HALT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataDescriptionOfHalt) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataDescriptionOfHalt) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataDescriptionOfHalt) GetDescriptionForHalt() *BACnetApplicationTagCharacterString {
	return m.DescriptionForHalt
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataDescriptionOfHalt factory function for BACnetConstructedDataDescriptionOfHalt
func NewBACnetConstructedDataDescriptionOfHalt(descriptionForHalt *BACnetApplicationTagCharacterString, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataDescriptionOfHalt {
	_result := &BACnetConstructedDataDescriptionOfHalt{
		DescriptionForHalt:    descriptionForHalt,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataDescriptionOfHalt(structType interface{}) *BACnetConstructedDataDescriptionOfHalt {
	if casted, ok := structType.(BACnetConstructedDataDescriptionOfHalt); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataDescriptionOfHalt); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataDescriptionOfHalt(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataDescriptionOfHalt(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataDescriptionOfHalt) GetTypeName() string {
	return "BACnetConstructedDataDescriptionOfHalt"
}

func (m *BACnetConstructedDataDescriptionOfHalt) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataDescriptionOfHalt) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (descriptionForHalt)
	lengthInBits += m.DescriptionForHalt.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataDescriptionOfHalt) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataDescriptionOfHaltParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataDescriptionOfHalt, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataDescriptionOfHalt"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataDescriptionOfHalt")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (descriptionForHalt)
	if pullErr := readBuffer.PullContext("descriptionForHalt"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for descriptionForHalt")
	}
	_descriptionForHalt, _descriptionForHaltErr := BACnetApplicationTagParse(readBuffer)
	if _descriptionForHaltErr != nil {
		return nil, errors.Wrap(_descriptionForHaltErr, "Error parsing 'descriptionForHalt' field")
	}
	descriptionForHalt := CastBACnetApplicationTagCharacterString(_descriptionForHalt)
	if closeErr := readBuffer.CloseContext("descriptionForHalt"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for descriptionForHalt")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataDescriptionOfHalt"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataDescriptionOfHalt")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataDescriptionOfHalt{
		DescriptionForHalt:    CastBACnetApplicationTagCharacterString(descriptionForHalt),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataDescriptionOfHalt) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataDescriptionOfHalt"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataDescriptionOfHalt")
		}

		// Simple Field (descriptionForHalt)
		if pushErr := writeBuffer.PushContext("descriptionForHalt"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for descriptionForHalt")
		}
		_descriptionForHaltErr := m.DescriptionForHalt.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("descriptionForHalt"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for descriptionForHalt")
		}
		if _descriptionForHaltErr != nil {
			return errors.Wrap(_descriptionForHaltErr, "Error serializing 'descriptionForHalt' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataDescriptionOfHalt"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataDescriptionOfHalt")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataDescriptionOfHalt) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
