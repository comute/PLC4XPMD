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

// BACnetConstructedDataProgramChange is the data-structure of this message
type BACnetConstructedDataProgramChange struct {
	*BACnetConstructedData
	ProgramChange *BACnetProgramRequestTagged

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataProgramChange is the corresponding interface of BACnetConstructedDataProgramChange
type IBACnetConstructedDataProgramChange interface {
	IBACnetConstructedData
	// GetProgramChange returns ProgramChange (property field)
	GetProgramChange() *BACnetProgramRequestTagged
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

func (m *BACnetConstructedDataProgramChange) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataProgramChange) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_PROGRAM_CHANGE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataProgramChange) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataProgramChange) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataProgramChange) GetProgramChange() *BACnetProgramRequestTagged {
	return m.ProgramChange
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataProgramChange factory function for BACnetConstructedDataProgramChange
func NewBACnetConstructedDataProgramChange(programChange *BACnetProgramRequestTagged, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataProgramChange {
	_result := &BACnetConstructedDataProgramChange{
		ProgramChange:         programChange,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataProgramChange(structType interface{}) *BACnetConstructedDataProgramChange {
	if casted, ok := structType.(BACnetConstructedDataProgramChange); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataProgramChange); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataProgramChange(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataProgramChange(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataProgramChange) GetTypeName() string {
	return "BACnetConstructedDataProgramChange"
}

func (m *BACnetConstructedDataProgramChange) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataProgramChange) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (programChange)
	lengthInBits += m.ProgramChange.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataProgramChange) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataProgramChangeParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataProgramChange, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataProgramChange"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataProgramChange")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (programChange)
	if pullErr := readBuffer.PullContext("programChange"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for programChange")
	}
	_programChange, _programChangeErr := BACnetProgramRequestTaggedParse(readBuffer, uint8(uint8(0)), TagClass(TagClass_APPLICATION_TAGS))
	if _programChangeErr != nil {
		return nil, errors.Wrap(_programChangeErr, "Error parsing 'programChange' field")
	}
	programChange := CastBACnetProgramRequestTagged(_programChange)
	if closeErr := readBuffer.CloseContext("programChange"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for programChange")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataProgramChange"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataProgramChange")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataProgramChange{
		ProgramChange:         CastBACnetProgramRequestTagged(programChange),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataProgramChange) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataProgramChange"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataProgramChange")
		}

		// Simple Field (programChange)
		if pushErr := writeBuffer.PushContext("programChange"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for programChange")
		}
		_programChangeErr := writeBuffer.WriteSerializable(m.ProgramChange)
		if popErr := writeBuffer.PopContext("programChange"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for programChange")
		}
		if _programChangeErr != nil {
			return errors.Wrap(_programChangeErr, "Error serializing 'programChange' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataProgramChange"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataProgramChange")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataProgramChange) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
