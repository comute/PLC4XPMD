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

// BACnetConstructedDataLinkSpeedAutonegotiate is the data-structure of this message
type BACnetConstructedDataLinkSpeedAutonegotiate struct {
	*BACnetConstructedData
	LinkSpeedAutonegotiate *BACnetApplicationTagBoolean

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataLinkSpeedAutonegotiate is the corresponding interface of BACnetConstructedDataLinkSpeedAutonegotiate
type IBACnetConstructedDataLinkSpeedAutonegotiate interface {
	IBACnetConstructedData
	// GetLinkSpeedAutonegotiate returns LinkSpeedAutonegotiate (property field)
	GetLinkSpeedAutonegotiate() *BACnetApplicationTagBoolean
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

func (m *BACnetConstructedDataLinkSpeedAutonegotiate) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataLinkSpeedAutonegotiate) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LINK_SPEED_AUTONEGOTIATE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataLinkSpeedAutonegotiate) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataLinkSpeedAutonegotiate) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataLinkSpeedAutonegotiate) GetLinkSpeedAutonegotiate() *BACnetApplicationTagBoolean {
	return m.LinkSpeedAutonegotiate
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataLinkSpeedAutonegotiate factory function for BACnetConstructedDataLinkSpeedAutonegotiate
func NewBACnetConstructedDataLinkSpeedAutonegotiate(linkSpeedAutonegotiate *BACnetApplicationTagBoolean, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataLinkSpeedAutonegotiate {
	_result := &BACnetConstructedDataLinkSpeedAutonegotiate{
		LinkSpeedAutonegotiate: linkSpeedAutonegotiate,
		BACnetConstructedData:  NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataLinkSpeedAutonegotiate(structType interface{}) *BACnetConstructedDataLinkSpeedAutonegotiate {
	if casted, ok := structType.(BACnetConstructedDataLinkSpeedAutonegotiate); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLinkSpeedAutonegotiate); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataLinkSpeedAutonegotiate(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataLinkSpeedAutonegotiate(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataLinkSpeedAutonegotiate) GetTypeName() string {
	return "BACnetConstructedDataLinkSpeedAutonegotiate"
}

func (m *BACnetConstructedDataLinkSpeedAutonegotiate) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataLinkSpeedAutonegotiate) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (linkSpeedAutonegotiate)
	lengthInBits += m.LinkSpeedAutonegotiate.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataLinkSpeedAutonegotiate) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataLinkSpeedAutonegotiateParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataLinkSpeedAutonegotiate, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLinkSpeedAutonegotiate"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLinkSpeedAutonegotiate")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (linkSpeedAutonegotiate)
	if pullErr := readBuffer.PullContext("linkSpeedAutonegotiate"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for linkSpeedAutonegotiate")
	}
	_linkSpeedAutonegotiate, _linkSpeedAutonegotiateErr := BACnetApplicationTagParse(readBuffer)
	if _linkSpeedAutonegotiateErr != nil {
		return nil, errors.Wrap(_linkSpeedAutonegotiateErr, "Error parsing 'linkSpeedAutonegotiate' field")
	}
	linkSpeedAutonegotiate := CastBACnetApplicationTagBoolean(_linkSpeedAutonegotiate)
	if closeErr := readBuffer.CloseContext("linkSpeedAutonegotiate"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for linkSpeedAutonegotiate")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLinkSpeedAutonegotiate"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLinkSpeedAutonegotiate")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataLinkSpeedAutonegotiate{
		LinkSpeedAutonegotiate: CastBACnetApplicationTagBoolean(linkSpeedAutonegotiate),
		BACnetConstructedData:  &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataLinkSpeedAutonegotiate) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLinkSpeedAutonegotiate"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLinkSpeedAutonegotiate")
		}

		// Simple Field (linkSpeedAutonegotiate)
		if pushErr := writeBuffer.PushContext("linkSpeedAutonegotiate"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for linkSpeedAutonegotiate")
		}
		_linkSpeedAutonegotiateErr := m.LinkSpeedAutonegotiate.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("linkSpeedAutonegotiate"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for linkSpeedAutonegotiate")
		}
		if _linkSpeedAutonegotiateErr != nil {
			return errors.Wrap(_linkSpeedAutonegotiateErr, "Error serializing 'linkSpeedAutonegotiate' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLinkSpeedAutonegotiate"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLinkSpeedAutonegotiate")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataLinkSpeedAutonegotiate) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
