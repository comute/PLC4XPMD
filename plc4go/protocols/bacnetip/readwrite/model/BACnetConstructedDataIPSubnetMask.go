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

// BACnetConstructedDataIPSubnetMask is the corresponding interface of BACnetConstructedDataIPSubnetMask
type BACnetConstructedDataIPSubnetMask interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetIpSubnetMask returns IpSubnetMask (property field)
	GetIpSubnetMask() BACnetApplicationTagOctetString
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagOctetString
}

// BACnetConstructedDataIPSubnetMaskExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataIPSubnetMask.
// This is useful for switch cases.
type BACnetConstructedDataIPSubnetMaskExactly interface {
	BACnetConstructedDataIPSubnetMask
	isBACnetConstructedDataIPSubnetMask() bool
}

// _BACnetConstructedDataIPSubnetMask is the data-structure of this message
type _BACnetConstructedDataIPSubnetMask struct {
	*_BACnetConstructedData
	IpSubnetMask BACnetApplicationTagOctetString

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument BACnetTagPayloadUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataIPSubnetMask) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataIPSubnetMask) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_IP_SUBNET_MASK
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataIPSubnetMask) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataIPSubnetMask) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataIPSubnetMask) GetIpSubnetMask() BACnetApplicationTagOctetString {
	return m.IpSubnetMask
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataIPSubnetMask) GetActualValue() BACnetApplicationTagOctetString {
	return CastBACnetApplicationTagOctetString(m.GetIpSubnetMask())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataIPSubnetMask factory function for _BACnetConstructedDataIPSubnetMask
func NewBACnetConstructedDataIPSubnetMask(ipSubnetMask BACnetApplicationTagOctetString, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataIPSubnetMask {
	_result := &_BACnetConstructedDataIPSubnetMask{
		IpSubnetMask:           ipSubnetMask,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataIPSubnetMask(structType interface{}) BACnetConstructedDataIPSubnetMask {
	if casted, ok := structType.(BACnetConstructedDataIPSubnetMask); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataIPSubnetMask); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataIPSubnetMask) GetTypeName() string {
	return "BACnetConstructedDataIPSubnetMask"
}

func (m *_BACnetConstructedDataIPSubnetMask) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataIPSubnetMask) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (ipSubnetMask)
	lengthInBits += m.IpSubnetMask.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataIPSubnetMask) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataIPSubnetMaskParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataIPSubnetMask, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataIPSubnetMask"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataIPSubnetMask")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (ipSubnetMask)
	if pullErr := readBuffer.PullContext("ipSubnetMask"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ipSubnetMask")
	}
	_ipSubnetMask, _ipSubnetMaskErr := BACnetApplicationTagParse(readBuffer)
	if _ipSubnetMaskErr != nil {
		return nil, errors.Wrap(_ipSubnetMaskErr, "Error parsing 'ipSubnetMask' field")
	}
	ipSubnetMask := _ipSubnetMask.(BACnetApplicationTagOctetString)
	if closeErr := readBuffer.CloseContext("ipSubnetMask"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ipSubnetMask")
	}

	// Virtual field
	_actualValue := ipSubnetMask
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataIPSubnetMask"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataIPSubnetMask")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataIPSubnetMask{
		IpSubnetMask:           ipSubnetMask,
		_BACnetConstructedData: &_BACnetConstructedData{},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataIPSubnetMask) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataIPSubnetMask"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataIPSubnetMask")
		}

		// Simple Field (ipSubnetMask)
		if pushErr := writeBuffer.PushContext("ipSubnetMask"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ipSubnetMask")
		}
		_ipSubnetMaskErr := writeBuffer.WriteSerializable(m.GetIpSubnetMask())
		if popErr := writeBuffer.PopContext("ipSubnetMask"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ipSubnetMask")
		}
		if _ipSubnetMaskErr != nil {
			return errors.Wrap(_ipSubnetMaskErr, "Error serializing 'ipSubnetMask' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataIPSubnetMask"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataIPSubnetMask")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataIPSubnetMask) isBACnetConstructedDataIPSubnetMask() bool {
	return true
}

func (m *_BACnetConstructedDataIPSubnetMask) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
