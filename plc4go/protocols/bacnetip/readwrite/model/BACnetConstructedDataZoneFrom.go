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

// BACnetConstructedDataZoneFrom is the data-structure of this message
type BACnetConstructedDataZoneFrom struct {
	*BACnetConstructedData
	ZoneFrom *BACnetDeviceObjectReference

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataZoneFrom is the corresponding interface of BACnetConstructedDataZoneFrom
type IBACnetConstructedDataZoneFrom interface {
	IBACnetConstructedData
	// GetZoneFrom returns ZoneFrom (property field)
	GetZoneFrom() *BACnetDeviceObjectReference
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

func (m *BACnetConstructedDataZoneFrom) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataZoneFrom) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ZONE_FROM
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataZoneFrom) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataZoneFrom) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataZoneFrom) GetZoneFrom() *BACnetDeviceObjectReference {
	return m.ZoneFrom
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataZoneFrom factory function for BACnetConstructedDataZoneFrom
func NewBACnetConstructedDataZoneFrom(zoneFrom *BACnetDeviceObjectReference, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataZoneFrom {
	_result := &BACnetConstructedDataZoneFrom{
		ZoneFrom:              zoneFrom,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataZoneFrom(structType interface{}) *BACnetConstructedDataZoneFrom {
	if casted, ok := structType.(BACnetConstructedDataZoneFrom); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataZoneFrom); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataZoneFrom(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataZoneFrom(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataZoneFrom) GetTypeName() string {
	return "BACnetConstructedDataZoneFrom"
}

func (m *BACnetConstructedDataZoneFrom) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataZoneFrom) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (zoneFrom)
	lengthInBits += m.ZoneFrom.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataZoneFrom) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataZoneFromParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataZoneFrom, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataZoneFrom"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataZoneFrom")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (zoneFrom)
	if pullErr := readBuffer.PullContext("zoneFrom"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for zoneFrom")
	}
	_zoneFrom, _zoneFromErr := BACnetDeviceObjectReferenceParse(readBuffer)
	if _zoneFromErr != nil {
		return nil, errors.Wrap(_zoneFromErr, "Error parsing 'zoneFrom' field")
	}
	zoneFrom := CastBACnetDeviceObjectReference(_zoneFrom)
	if closeErr := readBuffer.CloseContext("zoneFrom"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for zoneFrom")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataZoneFrom"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataZoneFrom")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataZoneFrom{
		ZoneFrom:              CastBACnetDeviceObjectReference(zoneFrom),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataZoneFrom) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataZoneFrom"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataZoneFrom")
		}

		// Simple Field (zoneFrom)
		if pushErr := writeBuffer.PushContext("zoneFrom"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for zoneFrom")
		}
		_zoneFromErr := m.ZoneFrom.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("zoneFrom"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for zoneFrom")
		}
		if _zoneFromErr != nil {
			return errors.Wrap(_zoneFromErr, "Error serializing 'zoneFrom' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataZoneFrom"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataZoneFrom")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataZoneFrom) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
