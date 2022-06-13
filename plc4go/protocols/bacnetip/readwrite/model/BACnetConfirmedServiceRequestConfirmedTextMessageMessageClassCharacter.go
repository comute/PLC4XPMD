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

// BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter is the data-structure of this message
type BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter struct {
	*BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass
	CharacterValue *BACnetContextTagCharacterString

	// Arguments.
	TagNumber uint8
}

// IBACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter is the corresponding interface of BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter
type IBACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter interface {
	IBACnetConfirmedServiceRequestConfirmedTextMessageMessageClass
	// GetCharacterValue returns CharacterValue (property field)
	GetCharacterValue() *BACnetContextTagCharacterString
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

func (m *BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter) InitializeParent(parent *BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass.OpeningTag = openingTag
	m.BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass.PeekedTagHeader = peekedTagHeader
	m.BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass.ClosingTag = closingTag
}

func (m *BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter) GetParent() *BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass {
	return m.BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter) GetCharacterValue() *BACnetContextTagCharacterString {
	return m.CharacterValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter factory function for BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter
func NewBACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter(characterValue *BACnetContextTagCharacterString, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter {
	_result := &BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter{
		CharacterValue: characterValue,
		BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass: NewBACnetConfirmedServiceRequestConfirmedTextMessageMessageClass(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter(structType interface{}) *BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter {
	if casted, ok := structType.(BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass); ok {
		return CastBACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter(casted.Child)
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass); ok {
		return CastBACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter(casted.Child)
	}
	return nil
}

func (m *BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter) GetTypeName() string {
	return "BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter"
}

func (m *BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (characterValue)
	lengthInBits += m.CharacterValue.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacterParse(readBuffer utils.ReadBuffer, tagNumber uint8) (*BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (characterValue)
	if pullErr := readBuffer.PullContext("characterValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for characterValue")
	}
	_characterValue, _characterValueErr := BACnetContextTagParse(readBuffer, uint8(uint8(1)), BACnetDataType(BACnetDataType_CHARACTER_STRING))
	if _characterValueErr != nil {
		return nil, errors.Wrap(_characterValueErr, "Error parsing 'characterValue' field")
	}
	characterValue := CastBACnetContextTagCharacterString(_characterValue)
	if closeErr := readBuffer.CloseContext("characterValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for characterValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter")
	}

	// Create a partially initialized instance
	_child := &BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter{
		CharacterValue: CastBACnetContextTagCharacterString(characterValue),
		BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass: &BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass{},
	}
	_child.BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass.Child = _child
	return _child, nil
}

func (m *BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter")
		}

		// Simple Field (characterValue)
		if pushErr := writeBuffer.PushContext("characterValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for characterValue")
		}
		_characterValueErr := writeBuffer.WriteSerializable(m.CharacterValue)
		if popErr := writeBuffer.PopContext("characterValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for characterValue")
		}
		if _characterValueErr != nil {
			return errors.Wrap(_characterValueErr, "Error serializing 'characterValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
