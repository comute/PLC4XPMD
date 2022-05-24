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

// BACnetLightingCommandEnclosed is the data-structure of this message
type BACnetLightingCommandEnclosed struct {
	OpeningTag      *BACnetOpeningTag
	LightingCommand *BACnetLightingCommand
	ClosingTag      *BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

// IBACnetLightingCommandEnclosed is the corresponding interface of BACnetLightingCommandEnclosed
type IBACnetLightingCommandEnclosed interface {
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() *BACnetOpeningTag
	// GetLightingCommand returns LightingCommand (property field)
	GetLightingCommand() *BACnetLightingCommand
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

func (m *BACnetLightingCommandEnclosed) GetOpeningTag() *BACnetOpeningTag {
	return m.OpeningTag
}

func (m *BACnetLightingCommandEnclosed) GetLightingCommand() *BACnetLightingCommand {
	return m.LightingCommand
}

func (m *BACnetLightingCommandEnclosed) GetClosingTag() *BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetLightingCommandEnclosed factory function for BACnetLightingCommandEnclosed
func NewBACnetLightingCommandEnclosed(openingTag *BACnetOpeningTag, lightingCommand *BACnetLightingCommand, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetLightingCommandEnclosed {
	return &BACnetLightingCommandEnclosed{OpeningTag: openingTag, LightingCommand: lightingCommand, ClosingTag: closingTag, TagNumber: tagNumber}
}

func CastBACnetLightingCommandEnclosed(structType interface{}) *BACnetLightingCommandEnclosed {
	if casted, ok := structType.(BACnetLightingCommandEnclosed); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetLightingCommandEnclosed); ok {
		return casted
	}
	return nil
}

func (m *BACnetLightingCommandEnclosed) GetTypeName() string {
	return "BACnetLightingCommandEnclosed"
}

func (m *BACnetLightingCommandEnclosed) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetLightingCommandEnclosed) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits()

	// Simple field (lightingCommand)
	lengthInBits += m.LightingCommand.GetLengthInBits()

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetLightingCommandEnclosed) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetLightingCommandEnclosedParse(readBuffer utils.ReadBuffer, tagNumber uint8) (*BACnetLightingCommandEnclosed, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetLightingCommandEnclosed"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, pullErr
	}
	_openingTag, _openingTagErr := BACnetOpeningTagParse(readBuffer, uint8(tagNumber))
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field")
	}
	openingTag := CastBACnetOpeningTag(_openingTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (lightingCommand)
	if pullErr := readBuffer.PullContext("lightingCommand"); pullErr != nil {
		return nil, pullErr
	}
	_lightingCommand, _lightingCommandErr := BACnetLightingCommandParse(readBuffer)
	if _lightingCommandErr != nil {
		return nil, errors.Wrap(_lightingCommandErr, "Error parsing 'lightingCommand' field")
	}
	lightingCommand := CastBACnetLightingCommand(_lightingCommand)
	if closeErr := readBuffer.CloseContext("lightingCommand"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, pullErr
	}
	_closingTag, _closingTagErr := BACnetClosingTagParse(readBuffer, uint8(tagNumber))
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field")
	}
	closingTag := CastBACnetClosingTag(_closingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetLightingCommandEnclosed"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewBACnetLightingCommandEnclosed(openingTag, lightingCommand, closingTag, tagNumber), nil
}

func (m *BACnetLightingCommandEnclosed) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetLightingCommandEnclosed"); pushErr != nil {
		return pushErr
	}

	// Simple Field (openingTag)
	if pushErr := writeBuffer.PushContext("openingTag"); pushErr != nil {
		return pushErr
	}
	_openingTagErr := m.OpeningTag.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("openingTag"); popErr != nil {
		return popErr
	}
	if _openingTagErr != nil {
		return errors.Wrap(_openingTagErr, "Error serializing 'openingTag' field")
	}

	// Simple Field (lightingCommand)
	if pushErr := writeBuffer.PushContext("lightingCommand"); pushErr != nil {
		return pushErr
	}
	_lightingCommandErr := m.LightingCommand.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("lightingCommand"); popErr != nil {
		return popErr
	}
	if _lightingCommandErr != nil {
		return errors.Wrap(_lightingCommandErr, "Error serializing 'lightingCommand' field")
	}

	// Simple Field (closingTag)
	if pushErr := writeBuffer.PushContext("closingTag"); pushErr != nil {
		return pushErr
	}
	_closingTagErr := m.ClosingTag.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("closingTag"); popErr != nil {
		return popErr
	}
	if _closingTagErr != nil {
		return errors.Wrap(_closingTagErr, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetLightingCommandEnclosed"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *BACnetLightingCommandEnclosed) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
