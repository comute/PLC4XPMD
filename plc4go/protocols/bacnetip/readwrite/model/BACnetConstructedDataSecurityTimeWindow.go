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

// BACnetConstructedDataSecurityTimeWindow is the corresponding interface of BACnetConstructedDataSecurityTimeWindow
type BACnetConstructedDataSecurityTimeWindow interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetSecurityTimeWindow returns SecurityTimeWindow (property field)
	GetSecurityTimeWindow() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataSecurityTimeWindowExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataSecurityTimeWindow.
// This is useful for switch cases.
type BACnetConstructedDataSecurityTimeWindowExactly interface {
	BACnetConstructedDataSecurityTimeWindow
	isBACnetConstructedDataSecurityTimeWindow() bool
}

// _BACnetConstructedDataSecurityTimeWindow is the data-structure of this message
type _BACnetConstructedDataSecurityTimeWindow struct {
	*_BACnetConstructedData
	SecurityTimeWindow BACnetApplicationTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataSecurityTimeWindow) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataSecurityTimeWindow) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_SECURITY_TIME_WINDOW
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataSecurityTimeWindow) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataSecurityTimeWindow) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataSecurityTimeWindow) GetSecurityTimeWindow() BACnetApplicationTagUnsignedInteger {
	return m.SecurityTimeWindow
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataSecurityTimeWindow) GetActualValue() BACnetApplicationTagUnsignedInteger {
	return CastBACnetApplicationTagUnsignedInteger(m.GetSecurityTimeWindow())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataSecurityTimeWindow factory function for _BACnetConstructedDataSecurityTimeWindow
func NewBACnetConstructedDataSecurityTimeWindow(securityTimeWindow BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataSecurityTimeWindow {
	_result := &_BACnetConstructedDataSecurityTimeWindow{
		SecurityTimeWindow:     securityTimeWindow,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataSecurityTimeWindow(structType interface{}) BACnetConstructedDataSecurityTimeWindow {
	if casted, ok := structType.(BACnetConstructedDataSecurityTimeWindow); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataSecurityTimeWindow); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataSecurityTimeWindow) GetTypeName() string {
	return "BACnetConstructedDataSecurityTimeWindow"
}

func (m *_BACnetConstructedDataSecurityTimeWindow) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataSecurityTimeWindow) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (securityTimeWindow)
	lengthInBits += m.SecurityTimeWindow.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataSecurityTimeWindow) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataSecurityTimeWindowParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataSecurityTimeWindow, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataSecurityTimeWindow"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataSecurityTimeWindow")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (securityTimeWindow)
	if pullErr := readBuffer.PullContext("securityTimeWindow"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for securityTimeWindow")
	}
	_securityTimeWindow, _securityTimeWindowErr := BACnetApplicationTagParse(readBuffer)
	if _securityTimeWindowErr != nil {
		return nil, errors.Wrap(_securityTimeWindowErr, "Error parsing 'securityTimeWindow' field")
	}
	securityTimeWindow := _securityTimeWindow.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("securityTimeWindow"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for securityTimeWindow")
	}

	// Virtual field
	_actualValue := securityTimeWindow
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataSecurityTimeWindow"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataSecurityTimeWindow")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataSecurityTimeWindow{
		SecurityTimeWindow: securityTimeWindow,
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataSecurityTimeWindow) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataSecurityTimeWindow"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataSecurityTimeWindow")
		}

		// Simple Field (securityTimeWindow)
		if pushErr := writeBuffer.PushContext("securityTimeWindow"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for securityTimeWindow")
		}
		_securityTimeWindowErr := writeBuffer.WriteSerializable(m.GetSecurityTimeWindow())
		if popErr := writeBuffer.PopContext("securityTimeWindow"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for securityTimeWindow")
		}
		if _securityTimeWindowErr != nil {
			return errors.Wrap(_securityTimeWindowErr, "Error serializing 'securityTimeWindow' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataSecurityTimeWindow"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataSecurityTimeWindow")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataSecurityTimeWindow) isBACnetConstructedDataSecurityTimeWindow() bool {
	return true
}

func (m *_BACnetConstructedDataSecurityTimeWindow) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
