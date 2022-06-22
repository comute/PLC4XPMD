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

// BACnetConstructedDataCredentialDataInputUpdateTime is the corresponding interface of BACnetConstructedDataCredentialDataInputUpdateTime
type BACnetConstructedDataCredentialDataInputUpdateTime interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetUpdateTime returns UpdateTime (property field)
	GetUpdateTime() BACnetTimeStamp
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetTimeStamp
}

// BACnetConstructedDataCredentialDataInputUpdateTimeExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataCredentialDataInputUpdateTime.
// This is useful for switch cases.
type BACnetConstructedDataCredentialDataInputUpdateTimeExactly interface {
	BACnetConstructedDataCredentialDataInputUpdateTime
	isBACnetConstructedDataCredentialDataInputUpdateTime() bool
}

// _BACnetConstructedDataCredentialDataInputUpdateTime is the data-structure of this message
type _BACnetConstructedDataCredentialDataInputUpdateTime struct {
	*_BACnetConstructedData
	UpdateTime BACnetTimeStamp

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument BACnetTagPayloadUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataCredentialDataInputUpdateTime) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_CREDENTIAL_DATA_INPUT
}

func (m *_BACnetConstructedDataCredentialDataInputUpdateTime) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_UPDATE_TIME
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataCredentialDataInputUpdateTime) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataCredentialDataInputUpdateTime) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataCredentialDataInputUpdateTime) GetUpdateTime() BACnetTimeStamp {
	return m.UpdateTime
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataCredentialDataInputUpdateTime) GetActualValue() BACnetTimeStamp {
	return CastBACnetTimeStamp(m.GetUpdateTime())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataCredentialDataInputUpdateTime factory function for _BACnetConstructedDataCredentialDataInputUpdateTime
func NewBACnetConstructedDataCredentialDataInputUpdateTime(updateTime BACnetTimeStamp, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataCredentialDataInputUpdateTime {
	_result := &_BACnetConstructedDataCredentialDataInputUpdateTime{
		UpdateTime:             updateTime,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataCredentialDataInputUpdateTime(structType interface{}) BACnetConstructedDataCredentialDataInputUpdateTime {
	if casted, ok := structType.(BACnetConstructedDataCredentialDataInputUpdateTime); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataCredentialDataInputUpdateTime); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataCredentialDataInputUpdateTime) GetTypeName() string {
	return "BACnetConstructedDataCredentialDataInputUpdateTime"
}

func (m *_BACnetConstructedDataCredentialDataInputUpdateTime) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataCredentialDataInputUpdateTime) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (updateTime)
	lengthInBits += m.UpdateTime.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataCredentialDataInputUpdateTime) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataCredentialDataInputUpdateTimeParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataCredentialDataInputUpdateTime, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataCredentialDataInputUpdateTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataCredentialDataInputUpdateTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (updateTime)
	if pullErr := readBuffer.PullContext("updateTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for updateTime")
	}
	_updateTime, _updateTimeErr := BACnetTimeStampParse(readBuffer)
	if _updateTimeErr != nil {
		return nil, errors.Wrap(_updateTimeErr, "Error parsing 'updateTime' field")
	}
	updateTime := _updateTime.(BACnetTimeStamp)
	if closeErr := readBuffer.CloseContext("updateTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for updateTime")
	}

	// Virtual field
	_actualValue := updateTime
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataCredentialDataInputUpdateTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataCredentialDataInputUpdateTime")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataCredentialDataInputUpdateTime{
		UpdateTime:             updateTime,
		_BACnetConstructedData: &_BACnetConstructedData{},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataCredentialDataInputUpdateTime) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataCredentialDataInputUpdateTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataCredentialDataInputUpdateTime")
		}

		// Simple Field (updateTime)
		if pushErr := writeBuffer.PushContext("updateTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for updateTime")
		}
		_updateTimeErr := writeBuffer.WriteSerializable(m.GetUpdateTime())
		if popErr := writeBuffer.PopContext("updateTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for updateTime")
		}
		if _updateTimeErr != nil {
			return errors.Wrap(_updateTimeErr, "Error serializing 'updateTime' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataCredentialDataInputUpdateTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataCredentialDataInputUpdateTime")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataCredentialDataInputUpdateTime) isBACnetConstructedDataCredentialDataInputUpdateTime() bool {
	return true
}

func (m *_BACnetConstructedDataCredentialDataInputUpdateTime) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
