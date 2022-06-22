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

// BACnetConstructedDataWindowInterval is the corresponding interface of BACnetConstructedDataWindowInterval
type BACnetConstructedDataWindowInterval interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetWindowInterval returns WindowInterval (property field)
	GetWindowInterval() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataWindowIntervalExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataWindowInterval.
// This is useful for switch cases.
type BACnetConstructedDataWindowIntervalExactly interface {
	BACnetConstructedDataWindowInterval
	isBACnetConstructedDataWindowInterval() bool
}

// _BACnetConstructedDataWindowInterval is the data-structure of this message
type _BACnetConstructedDataWindowInterval struct {
	*_BACnetConstructedData
	WindowInterval BACnetApplicationTagUnsignedInteger

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument BACnetTagPayloadUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataWindowInterval) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataWindowInterval) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_WINDOW_INTERVAL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataWindowInterval) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataWindowInterval) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataWindowInterval) GetWindowInterval() BACnetApplicationTagUnsignedInteger {
	return m.WindowInterval
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataWindowInterval) GetActualValue() BACnetApplicationTagUnsignedInteger {
	return CastBACnetApplicationTagUnsignedInteger(m.GetWindowInterval())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataWindowInterval factory function for _BACnetConstructedDataWindowInterval
func NewBACnetConstructedDataWindowInterval(windowInterval BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataWindowInterval {
	_result := &_BACnetConstructedDataWindowInterval{
		WindowInterval:         windowInterval,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataWindowInterval(structType interface{}) BACnetConstructedDataWindowInterval {
	if casted, ok := structType.(BACnetConstructedDataWindowInterval); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataWindowInterval); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataWindowInterval) GetTypeName() string {
	return "BACnetConstructedDataWindowInterval"
}

func (m *_BACnetConstructedDataWindowInterval) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataWindowInterval) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (windowInterval)
	lengthInBits += m.WindowInterval.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataWindowInterval) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataWindowIntervalParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataWindowInterval, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataWindowInterval"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataWindowInterval")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (windowInterval)
	if pullErr := readBuffer.PullContext("windowInterval"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for windowInterval")
	}
	_windowInterval, _windowIntervalErr := BACnetApplicationTagParse(readBuffer)
	if _windowIntervalErr != nil {
		return nil, errors.Wrap(_windowIntervalErr, "Error parsing 'windowInterval' field")
	}
	windowInterval := _windowInterval.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("windowInterval"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for windowInterval")
	}

	// Virtual field
	_actualValue := windowInterval
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataWindowInterval"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataWindowInterval")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataWindowInterval{
		WindowInterval:         windowInterval,
		_BACnetConstructedData: &_BACnetConstructedData{},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataWindowInterval) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataWindowInterval"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataWindowInterval")
		}

		// Simple Field (windowInterval)
		if pushErr := writeBuffer.PushContext("windowInterval"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for windowInterval")
		}
		_windowIntervalErr := writeBuffer.WriteSerializable(m.GetWindowInterval())
		if popErr := writeBuffer.PopContext("windowInterval"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for windowInterval")
		}
		if _windowIntervalErr != nil {
			return errors.Wrap(_windowIntervalErr, "Error serializing 'windowInterval' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataWindowInterval"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataWindowInterval")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataWindowInterval) isBACnetConstructedDataWindowInterval() bool {
	return true
}

func (m *_BACnetConstructedDataWindowInterval) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
