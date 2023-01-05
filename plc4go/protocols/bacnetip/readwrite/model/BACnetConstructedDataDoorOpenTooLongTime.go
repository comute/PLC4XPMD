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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataDoorOpenTooLongTime is the corresponding interface of BACnetConstructedDataDoorOpenTooLongTime
type BACnetConstructedDataDoorOpenTooLongTime interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetDoorOpenTooLongTime returns DoorOpenTooLongTime (property field)
	GetDoorOpenTooLongTime() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataDoorOpenTooLongTimeExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataDoorOpenTooLongTime.
// This is useful for switch cases.
type BACnetConstructedDataDoorOpenTooLongTimeExactly interface {
	BACnetConstructedDataDoorOpenTooLongTime
	isBACnetConstructedDataDoorOpenTooLongTime() bool
}

// _BACnetConstructedDataDoorOpenTooLongTime is the data-structure of this message
type _BACnetConstructedDataDoorOpenTooLongTime struct {
	*_BACnetConstructedData
	DoorOpenTooLongTime BACnetApplicationTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataDoorOpenTooLongTime) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataDoorOpenTooLongTime) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_DOOR_OPEN_TOO_LONG_TIME
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataDoorOpenTooLongTime) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataDoorOpenTooLongTime) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataDoorOpenTooLongTime) GetDoorOpenTooLongTime() BACnetApplicationTagUnsignedInteger {
	return m.DoorOpenTooLongTime
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataDoorOpenTooLongTime) GetActualValue() BACnetApplicationTagUnsignedInteger {
	return CastBACnetApplicationTagUnsignedInteger(m.GetDoorOpenTooLongTime())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataDoorOpenTooLongTime factory function for _BACnetConstructedDataDoorOpenTooLongTime
func NewBACnetConstructedDataDoorOpenTooLongTime(doorOpenTooLongTime BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataDoorOpenTooLongTime {
	_result := &_BACnetConstructedDataDoorOpenTooLongTime{
		DoorOpenTooLongTime:    doorOpenTooLongTime,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataDoorOpenTooLongTime(structType interface{}) BACnetConstructedDataDoorOpenTooLongTime {
	if casted, ok := structType.(BACnetConstructedDataDoorOpenTooLongTime); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataDoorOpenTooLongTime); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataDoorOpenTooLongTime) GetTypeName() string {
	return "BACnetConstructedDataDoorOpenTooLongTime"
}

func (m *_BACnetConstructedDataDoorOpenTooLongTime) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataDoorOpenTooLongTime) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (doorOpenTooLongTime)
	lengthInBits += m.DoorOpenTooLongTime.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataDoorOpenTooLongTime) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataDoorOpenTooLongTimeParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataDoorOpenTooLongTime, error) {
	return BACnetConstructedDataDoorOpenTooLongTimeParseWithBuffer(utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataDoorOpenTooLongTimeParseWithBuffer(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataDoorOpenTooLongTime, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataDoorOpenTooLongTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataDoorOpenTooLongTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (doorOpenTooLongTime)
	if pullErr := readBuffer.PullContext("doorOpenTooLongTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for doorOpenTooLongTime")
	}
	_doorOpenTooLongTime, _doorOpenTooLongTimeErr := BACnetApplicationTagParseWithBuffer(readBuffer)
	if _doorOpenTooLongTimeErr != nil {
		return nil, errors.Wrap(_doorOpenTooLongTimeErr, "Error parsing 'doorOpenTooLongTime' field of BACnetConstructedDataDoorOpenTooLongTime")
	}
	doorOpenTooLongTime := _doorOpenTooLongTime.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("doorOpenTooLongTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for doorOpenTooLongTime")
	}

	// Virtual field
	_actualValue := doorOpenTooLongTime
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataDoorOpenTooLongTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataDoorOpenTooLongTime")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataDoorOpenTooLongTime{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		DoorOpenTooLongTime: doorOpenTooLongTime,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataDoorOpenTooLongTime) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes())))
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataDoorOpenTooLongTime) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataDoorOpenTooLongTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataDoorOpenTooLongTime")
		}

		// Simple Field (doorOpenTooLongTime)
		if pushErr := writeBuffer.PushContext("doorOpenTooLongTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for doorOpenTooLongTime")
		}
		_doorOpenTooLongTimeErr := writeBuffer.WriteSerializable(m.GetDoorOpenTooLongTime())
		if popErr := writeBuffer.PopContext("doorOpenTooLongTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for doorOpenTooLongTime")
		}
		if _doorOpenTooLongTimeErr != nil {
			return errors.Wrap(_doorOpenTooLongTimeErr, "Error serializing 'doorOpenTooLongTime' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataDoorOpenTooLongTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataDoorOpenTooLongTime")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataDoorOpenTooLongTime) isBACnetConstructedDataDoorOpenTooLongTime() bool {
	return true
}

func (m *_BACnetConstructedDataDoorOpenTooLongTime) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
