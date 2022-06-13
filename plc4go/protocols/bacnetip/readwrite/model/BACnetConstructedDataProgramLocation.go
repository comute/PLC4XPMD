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

// BACnetConstructedDataProgramLocation is the data-structure of this message
type BACnetConstructedDataProgramLocation struct {
	*BACnetConstructedData
	ProgramLocation *BACnetApplicationTagCharacterString

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataProgramLocation is the corresponding interface of BACnetConstructedDataProgramLocation
type IBACnetConstructedDataProgramLocation interface {
	IBACnetConstructedData
	// GetProgramLocation returns ProgramLocation (property field)
	GetProgramLocation() *BACnetApplicationTagCharacterString
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

func (m *BACnetConstructedDataProgramLocation) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataProgramLocation) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_PROGRAM_LOCATION
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataProgramLocation) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataProgramLocation) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataProgramLocation) GetProgramLocation() *BACnetApplicationTagCharacterString {
	return m.ProgramLocation
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataProgramLocation factory function for BACnetConstructedDataProgramLocation
func NewBACnetConstructedDataProgramLocation(programLocation *BACnetApplicationTagCharacterString, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataProgramLocation {
	_result := &BACnetConstructedDataProgramLocation{
		ProgramLocation:       programLocation,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataProgramLocation(structType interface{}) *BACnetConstructedDataProgramLocation {
	if casted, ok := structType.(BACnetConstructedDataProgramLocation); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataProgramLocation); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataProgramLocation(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataProgramLocation(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataProgramLocation) GetTypeName() string {
	return "BACnetConstructedDataProgramLocation"
}

func (m *BACnetConstructedDataProgramLocation) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataProgramLocation) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (programLocation)
	lengthInBits += m.ProgramLocation.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataProgramLocation) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataProgramLocationParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataProgramLocation, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataProgramLocation"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataProgramLocation")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (programLocation)
	if pullErr := readBuffer.PullContext("programLocation"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for programLocation")
	}
	_programLocation, _programLocationErr := BACnetApplicationTagParse(readBuffer)
	if _programLocationErr != nil {
		return nil, errors.Wrap(_programLocationErr, "Error parsing 'programLocation' field")
	}
	programLocation := CastBACnetApplicationTagCharacterString(_programLocation)
	if closeErr := readBuffer.CloseContext("programLocation"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for programLocation")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataProgramLocation"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataProgramLocation")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataProgramLocation{
		ProgramLocation:       CastBACnetApplicationTagCharacterString(programLocation),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataProgramLocation) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataProgramLocation"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataProgramLocation")
		}

		// Simple Field (programLocation)
		if pushErr := writeBuffer.PushContext("programLocation"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for programLocation")
		}
		_programLocationErr := writeBuffer.WriteSerializable(m.ProgramLocation)
		if popErr := writeBuffer.PopContext("programLocation"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for programLocation")
		}
		if _programLocationErr != nil {
			return errors.Wrap(_programLocationErr, "Error serializing 'programLocation' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataProgramLocation"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataProgramLocation")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataProgramLocation) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
