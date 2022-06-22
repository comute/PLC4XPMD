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

// BACnetConstructedDataDaylightSavingsStatus is the corresponding interface of BACnetConstructedDataDaylightSavingsStatus
type BACnetConstructedDataDaylightSavingsStatus interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetDaylightSavingsStatus returns DaylightSavingsStatus (property field)
	GetDaylightSavingsStatus() BACnetApplicationTagBoolean
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagBoolean
}

// BACnetConstructedDataDaylightSavingsStatusExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataDaylightSavingsStatus.
// This is useful for switch cases.
type BACnetConstructedDataDaylightSavingsStatusExactly interface {
	BACnetConstructedDataDaylightSavingsStatus
	isBACnetConstructedDataDaylightSavingsStatus() bool
}

// _BACnetConstructedDataDaylightSavingsStatus is the data-structure of this message
type _BACnetConstructedDataDaylightSavingsStatus struct {
	*_BACnetConstructedData
	DaylightSavingsStatus BACnetApplicationTagBoolean

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument BACnetTagPayloadUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataDaylightSavingsStatus) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataDaylightSavingsStatus) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_DAYLIGHT_SAVINGS_STATUS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataDaylightSavingsStatus) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataDaylightSavingsStatus) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataDaylightSavingsStatus) GetDaylightSavingsStatus() BACnetApplicationTagBoolean {
	return m.DaylightSavingsStatus
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataDaylightSavingsStatus) GetActualValue() BACnetApplicationTagBoolean {
	return CastBACnetApplicationTagBoolean(m.GetDaylightSavingsStatus())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataDaylightSavingsStatus factory function for _BACnetConstructedDataDaylightSavingsStatus
func NewBACnetConstructedDataDaylightSavingsStatus(daylightSavingsStatus BACnetApplicationTagBoolean, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataDaylightSavingsStatus {
	_result := &_BACnetConstructedDataDaylightSavingsStatus{
		DaylightSavingsStatus:  daylightSavingsStatus,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataDaylightSavingsStatus(structType interface{}) BACnetConstructedDataDaylightSavingsStatus {
	if casted, ok := structType.(BACnetConstructedDataDaylightSavingsStatus); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataDaylightSavingsStatus); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataDaylightSavingsStatus) GetTypeName() string {
	return "BACnetConstructedDataDaylightSavingsStatus"
}

func (m *_BACnetConstructedDataDaylightSavingsStatus) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataDaylightSavingsStatus) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (daylightSavingsStatus)
	lengthInBits += m.DaylightSavingsStatus.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataDaylightSavingsStatus) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataDaylightSavingsStatusParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataDaylightSavingsStatus, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataDaylightSavingsStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataDaylightSavingsStatus")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (daylightSavingsStatus)
	if pullErr := readBuffer.PullContext("daylightSavingsStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for daylightSavingsStatus")
	}
	_daylightSavingsStatus, _daylightSavingsStatusErr := BACnetApplicationTagParse(readBuffer)
	if _daylightSavingsStatusErr != nil {
		return nil, errors.Wrap(_daylightSavingsStatusErr, "Error parsing 'daylightSavingsStatus' field")
	}
	daylightSavingsStatus := _daylightSavingsStatus.(BACnetApplicationTagBoolean)
	if closeErr := readBuffer.CloseContext("daylightSavingsStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for daylightSavingsStatus")
	}

	// Virtual field
	_actualValue := daylightSavingsStatus
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataDaylightSavingsStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataDaylightSavingsStatus")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataDaylightSavingsStatus{
		DaylightSavingsStatus:  daylightSavingsStatus,
		_BACnetConstructedData: &_BACnetConstructedData{},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataDaylightSavingsStatus) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataDaylightSavingsStatus"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataDaylightSavingsStatus")
		}

		// Simple Field (daylightSavingsStatus)
		if pushErr := writeBuffer.PushContext("daylightSavingsStatus"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for daylightSavingsStatus")
		}
		_daylightSavingsStatusErr := writeBuffer.WriteSerializable(m.GetDaylightSavingsStatus())
		if popErr := writeBuffer.PopContext("daylightSavingsStatus"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for daylightSavingsStatus")
		}
		if _daylightSavingsStatusErr != nil {
			return errors.Wrap(_daylightSavingsStatusErr, "Error serializing 'daylightSavingsStatus' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataDaylightSavingsStatus"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataDaylightSavingsStatus")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataDaylightSavingsStatus) isBACnetConstructedDataDaylightSavingsStatus() bool {
	return true
}

func (m *_BACnetConstructedDataDaylightSavingsStatus) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
