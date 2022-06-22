/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
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

// BACnetConstructedDataIPv6ZoneIndex is the corresponding interface of BACnetConstructedDataIPv6ZoneIndex
type BACnetConstructedDataIPv6ZoneIndex interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetIpv6ZoneIndex returns Ipv6ZoneIndex (property field)
	GetIpv6ZoneIndex() BACnetApplicationTagCharacterString
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagCharacterString
}

// BACnetConstructedDataIPv6ZoneIndexExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataIPv6ZoneIndex.
// This is useful for switch cases.
type BACnetConstructedDataIPv6ZoneIndexExactly interface {
	BACnetConstructedDataIPv6ZoneIndex
	isBACnetConstructedDataIPv6ZoneIndex() bool
}

// _BACnetConstructedDataIPv6ZoneIndex is the data-structure of this message
type _BACnetConstructedDataIPv6ZoneIndex struct {
	*_BACnetConstructedData
	Ipv6ZoneIndex BACnetApplicationTagCharacterString

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument BACnetTagPayloadUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataIPv6ZoneIndex) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataIPv6ZoneIndex) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_IPV6_ZONE_INDEX
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataIPv6ZoneIndex) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataIPv6ZoneIndex) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataIPv6ZoneIndex) GetIpv6ZoneIndex() BACnetApplicationTagCharacterString {
	return m.Ipv6ZoneIndex
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataIPv6ZoneIndex) GetActualValue() BACnetApplicationTagCharacterString {
	return CastBACnetApplicationTagCharacterString(m.GetIpv6ZoneIndex())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataIPv6ZoneIndex factory function for _BACnetConstructedDataIPv6ZoneIndex
func NewBACnetConstructedDataIPv6ZoneIndex(ipv6ZoneIndex BACnetApplicationTagCharacterString, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataIPv6ZoneIndex {
	_result := &_BACnetConstructedDataIPv6ZoneIndex{
		Ipv6ZoneIndex:          ipv6ZoneIndex,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataIPv6ZoneIndex(structType interface{}) BACnetConstructedDataIPv6ZoneIndex {
	if casted, ok := structType.(BACnetConstructedDataIPv6ZoneIndex); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataIPv6ZoneIndex); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataIPv6ZoneIndex) GetTypeName() string {
	return "BACnetConstructedDataIPv6ZoneIndex"
}

func (m *_BACnetConstructedDataIPv6ZoneIndex) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataIPv6ZoneIndex) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (ipv6ZoneIndex)
	lengthInBits += m.Ipv6ZoneIndex.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataIPv6ZoneIndex) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataIPv6ZoneIndexParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataIPv6ZoneIndex, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataIPv6ZoneIndex"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataIPv6ZoneIndex")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (ipv6ZoneIndex)
	if pullErr := readBuffer.PullContext("ipv6ZoneIndex"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ipv6ZoneIndex")
	}
	_ipv6ZoneIndex, _ipv6ZoneIndexErr := BACnetApplicationTagParse(readBuffer)
	if _ipv6ZoneIndexErr != nil {
		return nil, errors.Wrap(_ipv6ZoneIndexErr, "Error parsing 'ipv6ZoneIndex' field")
	}
	ipv6ZoneIndex := _ipv6ZoneIndex.(BACnetApplicationTagCharacterString)
	if closeErr := readBuffer.CloseContext("ipv6ZoneIndex"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ipv6ZoneIndex")
	}

	// Virtual field
	_actualValue := ipv6ZoneIndex
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataIPv6ZoneIndex"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataIPv6ZoneIndex")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataIPv6ZoneIndex{
		Ipv6ZoneIndex:          ipv6ZoneIndex,
		_BACnetConstructedData: &_BACnetConstructedData{},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataIPv6ZoneIndex) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataIPv6ZoneIndex"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataIPv6ZoneIndex")
		}

		// Simple Field (ipv6ZoneIndex)
		if pushErr := writeBuffer.PushContext("ipv6ZoneIndex"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ipv6ZoneIndex")
		}
		_ipv6ZoneIndexErr := writeBuffer.WriteSerializable(m.GetIpv6ZoneIndex())
		if popErr := writeBuffer.PopContext("ipv6ZoneIndex"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ipv6ZoneIndex")
		}
		if _ipv6ZoneIndexErr != nil {
			return errors.Wrap(_ipv6ZoneIndexErr, "Error serializing 'ipv6ZoneIndex' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataIPv6ZoneIndex"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataIPv6ZoneIndex")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataIPv6ZoneIndex) isBACnetConstructedDataIPv6ZoneIndex() bool {
	return true
}

func (m *_BACnetConstructedDataIPv6ZoneIndex) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
