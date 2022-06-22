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

// BACnetConstructedDataTimeOfDeviceRestart is the corresponding interface of BACnetConstructedDataTimeOfDeviceRestart
type BACnetConstructedDataTimeOfDeviceRestart interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetTimeOfDeviceRestart returns TimeOfDeviceRestart (property field)
	GetTimeOfDeviceRestart() BACnetTimeStamp
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetTimeStamp
}

// BACnetConstructedDataTimeOfDeviceRestartExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataTimeOfDeviceRestart.
// This is useful for switch cases.
type BACnetConstructedDataTimeOfDeviceRestartExactly interface {
	BACnetConstructedDataTimeOfDeviceRestart
	isBACnetConstructedDataTimeOfDeviceRestart() bool
}

// _BACnetConstructedDataTimeOfDeviceRestart is the data-structure of this message
type _BACnetConstructedDataTimeOfDeviceRestart struct {
	*_BACnetConstructedData
	TimeOfDeviceRestart BACnetTimeStamp

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument BACnetTagPayloadUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataTimeOfDeviceRestart) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataTimeOfDeviceRestart) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_TIME_OF_DEVICE_RESTART
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataTimeOfDeviceRestart) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataTimeOfDeviceRestart) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataTimeOfDeviceRestart) GetTimeOfDeviceRestart() BACnetTimeStamp {
	return m.TimeOfDeviceRestart
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataTimeOfDeviceRestart) GetActualValue() BACnetTimeStamp {
	return CastBACnetTimeStamp(m.GetTimeOfDeviceRestart())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataTimeOfDeviceRestart factory function for _BACnetConstructedDataTimeOfDeviceRestart
func NewBACnetConstructedDataTimeOfDeviceRestart(timeOfDeviceRestart BACnetTimeStamp, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataTimeOfDeviceRestart {
	_result := &_BACnetConstructedDataTimeOfDeviceRestart{
		TimeOfDeviceRestart:    timeOfDeviceRestart,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataTimeOfDeviceRestart(structType interface{}) BACnetConstructedDataTimeOfDeviceRestart {
	if casted, ok := structType.(BACnetConstructedDataTimeOfDeviceRestart); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataTimeOfDeviceRestart); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataTimeOfDeviceRestart) GetTypeName() string {
	return "BACnetConstructedDataTimeOfDeviceRestart"
}

func (m *_BACnetConstructedDataTimeOfDeviceRestart) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataTimeOfDeviceRestart) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (timeOfDeviceRestart)
	lengthInBits += m.TimeOfDeviceRestart.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataTimeOfDeviceRestart) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataTimeOfDeviceRestartParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataTimeOfDeviceRestart, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataTimeOfDeviceRestart"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataTimeOfDeviceRestart")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (timeOfDeviceRestart)
	if pullErr := readBuffer.PullContext("timeOfDeviceRestart"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for timeOfDeviceRestart")
	}
	_timeOfDeviceRestart, _timeOfDeviceRestartErr := BACnetTimeStampParse(readBuffer)
	if _timeOfDeviceRestartErr != nil {
		return nil, errors.Wrap(_timeOfDeviceRestartErr, "Error parsing 'timeOfDeviceRestart' field")
	}
	timeOfDeviceRestart := _timeOfDeviceRestart.(BACnetTimeStamp)
	if closeErr := readBuffer.CloseContext("timeOfDeviceRestart"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for timeOfDeviceRestart")
	}

	// Virtual field
	_actualValue := timeOfDeviceRestart
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataTimeOfDeviceRestart"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataTimeOfDeviceRestart")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataTimeOfDeviceRestart{
		TimeOfDeviceRestart:    timeOfDeviceRestart,
		_BACnetConstructedData: &_BACnetConstructedData{},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataTimeOfDeviceRestart) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataTimeOfDeviceRestart"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataTimeOfDeviceRestart")
		}

		// Simple Field (timeOfDeviceRestart)
		if pushErr := writeBuffer.PushContext("timeOfDeviceRestart"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for timeOfDeviceRestart")
		}
		_timeOfDeviceRestartErr := writeBuffer.WriteSerializable(m.GetTimeOfDeviceRestart())
		if popErr := writeBuffer.PopContext("timeOfDeviceRestart"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for timeOfDeviceRestart")
		}
		if _timeOfDeviceRestartErr != nil {
			return errors.Wrap(_timeOfDeviceRestartErr, "Error serializing 'timeOfDeviceRestart' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataTimeOfDeviceRestart"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataTimeOfDeviceRestart")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataTimeOfDeviceRestart) isBACnetConstructedDataTimeOfDeviceRestart() bool {
	return true
}

func (m *_BACnetConstructedDataTimeOfDeviceRestart) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
