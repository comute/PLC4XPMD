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

// BACnetConstructedDataProfileLocation is the data-structure of this message
type BACnetConstructedDataProfileLocation struct {
	*BACnetConstructedData
	ProfileLocation *BACnetApplicationTagCharacterString

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataProfileLocation is the corresponding interface of BACnetConstructedDataProfileLocation
type IBACnetConstructedDataProfileLocation interface {
	IBACnetConstructedData
	// GetProfileLocation returns ProfileLocation (property field)
	GetProfileLocation() *BACnetApplicationTagCharacterString
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

func (m *BACnetConstructedDataProfileLocation) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataProfileLocation) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_PROFILE_LOCATION
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataProfileLocation) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataProfileLocation) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataProfileLocation) GetProfileLocation() *BACnetApplicationTagCharacterString {
	return m.ProfileLocation
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataProfileLocation factory function for BACnetConstructedDataProfileLocation
func NewBACnetConstructedDataProfileLocation(profileLocation *BACnetApplicationTagCharacterString, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataProfileLocation {
	_result := &BACnetConstructedDataProfileLocation{
		ProfileLocation:       profileLocation,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataProfileLocation(structType interface{}) *BACnetConstructedDataProfileLocation {
	if casted, ok := structType.(BACnetConstructedDataProfileLocation); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataProfileLocation); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataProfileLocation(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataProfileLocation(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataProfileLocation) GetTypeName() string {
	return "BACnetConstructedDataProfileLocation"
}

func (m *BACnetConstructedDataProfileLocation) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataProfileLocation) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (profileLocation)
	lengthInBits += m.ProfileLocation.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataProfileLocation) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataProfileLocationParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataProfileLocation, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataProfileLocation"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataProfileLocation")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (profileLocation)
	if pullErr := readBuffer.PullContext("profileLocation"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for profileLocation")
	}
	_profileLocation, _profileLocationErr := BACnetApplicationTagParse(readBuffer)
	if _profileLocationErr != nil {
		return nil, errors.Wrap(_profileLocationErr, "Error parsing 'profileLocation' field")
	}
	profileLocation := CastBACnetApplicationTagCharacterString(_profileLocation)
	if closeErr := readBuffer.CloseContext("profileLocation"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for profileLocation")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataProfileLocation"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataProfileLocation")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataProfileLocation{
		ProfileLocation:       CastBACnetApplicationTagCharacterString(profileLocation),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataProfileLocation) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataProfileLocation"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataProfileLocation")
		}

		// Simple Field (profileLocation)
		if pushErr := writeBuffer.PushContext("profileLocation"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for profileLocation")
		}
		_profileLocationErr := writeBuffer.WriteSerializable(m.ProfileLocation)
		if popErr := writeBuffer.PopContext("profileLocation"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for profileLocation")
		}
		if _profileLocationErr != nil {
			return errors.Wrap(_profileLocationErr, "Error serializing 'profileLocation' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataProfileLocation"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataProfileLocation")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataProfileLocation) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
