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

// BACnetConstructedDataThreatLevel is the corresponding interface of BACnetConstructedDataThreatLevel
type BACnetConstructedDataThreatLevel interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetThreatLevel returns ThreatLevel (property field)
	GetThreatLevel() BACnetAccessThreatLevel
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetAccessThreatLevel
}

// BACnetConstructedDataThreatLevelExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataThreatLevel.
// This is useful for switch cases.
type BACnetConstructedDataThreatLevelExactly interface {
	BACnetConstructedDataThreatLevel
	isBACnetConstructedDataThreatLevel() bool
}

// _BACnetConstructedDataThreatLevel is the data-structure of this message
type _BACnetConstructedDataThreatLevel struct {
	*_BACnetConstructedData
	ThreatLevel BACnetAccessThreatLevel

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument BACnetTagPayloadUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataThreatLevel) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataThreatLevel) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_THREAT_LEVEL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataThreatLevel) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataThreatLevel) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataThreatLevel) GetThreatLevel() BACnetAccessThreatLevel {
	return m.ThreatLevel
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataThreatLevel) GetActualValue() BACnetAccessThreatLevel {
	return CastBACnetAccessThreatLevel(m.GetThreatLevel())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataThreatLevel factory function for _BACnetConstructedDataThreatLevel
func NewBACnetConstructedDataThreatLevel(threatLevel BACnetAccessThreatLevel, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataThreatLevel {
	_result := &_BACnetConstructedDataThreatLevel{
		ThreatLevel:            threatLevel,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataThreatLevel(structType interface{}) BACnetConstructedDataThreatLevel {
	if casted, ok := structType.(BACnetConstructedDataThreatLevel); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataThreatLevel); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataThreatLevel) GetTypeName() string {
	return "BACnetConstructedDataThreatLevel"
}

func (m *_BACnetConstructedDataThreatLevel) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataThreatLevel) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (threatLevel)
	lengthInBits += m.ThreatLevel.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataThreatLevel) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataThreatLevelParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataThreatLevel, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataThreatLevel"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataThreatLevel")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (threatLevel)
	if pullErr := readBuffer.PullContext("threatLevel"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for threatLevel")
	}
	_threatLevel, _threatLevelErr := BACnetAccessThreatLevelParse(readBuffer)
	if _threatLevelErr != nil {
		return nil, errors.Wrap(_threatLevelErr, "Error parsing 'threatLevel' field")
	}
	threatLevel := _threatLevel.(BACnetAccessThreatLevel)
	if closeErr := readBuffer.CloseContext("threatLevel"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for threatLevel")
	}

	// Virtual field
	_actualValue := threatLevel
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataThreatLevel"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataThreatLevel")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataThreatLevel{
		ThreatLevel:            threatLevel,
		_BACnetConstructedData: &_BACnetConstructedData{},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataThreatLevel) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataThreatLevel"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataThreatLevel")
		}

		// Simple Field (threatLevel)
		if pushErr := writeBuffer.PushContext("threatLevel"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for threatLevel")
		}
		_threatLevelErr := writeBuffer.WriteSerializable(m.GetThreatLevel())
		if popErr := writeBuffer.PopContext("threatLevel"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for threatLevel")
		}
		if _threatLevelErr != nil {
			return errors.Wrap(_threatLevelErr, "Error serializing 'threatLevel' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataThreatLevel"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataThreatLevel")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataThreatLevel) isBACnetConstructedDataThreatLevel() bool {
	return true
}

func (m *_BACnetConstructedDataThreatLevel) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
