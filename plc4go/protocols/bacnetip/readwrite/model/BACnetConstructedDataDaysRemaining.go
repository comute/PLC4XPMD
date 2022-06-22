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

// BACnetConstructedDataDaysRemaining is the corresponding interface of BACnetConstructedDataDaysRemaining
type BACnetConstructedDataDaysRemaining interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetDaysRemaining returns DaysRemaining (property field)
	GetDaysRemaining() BACnetApplicationTagSignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagSignedInteger
}

// BACnetConstructedDataDaysRemainingExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataDaysRemaining.
// This is useful for switch cases.
type BACnetConstructedDataDaysRemainingExactly interface {
	BACnetConstructedDataDaysRemaining
	isBACnetConstructedDataDaysRemaining() bool
}

// _BACnetConstructedDataDaysRemaining is the data-structure of this message
type _BACnetConstructedDataDaysRemaining struct {
	*_BACnetConstructedData
	DaysRemaining BACnetApplicationTagSignedInteger

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument BACnetTagPayloadUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataDaysRemaining) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataDaysRemaining) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_DAYS_REMAINING
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataDaysRemaining) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataDaysRemaining) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataDaysRemaining) GetDaysRemaining() BACnetApplicationTagSignedInteger {
	return m.DaysRemaining
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataDaysRemaining) GetActualValue() BACnetApplicationTagSignedInteger {
	return CastBACnetApplicationTagSignedInteger(m.GetDaysRemaining())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataDaysRemaining factory function for _BACnetConstructedDataDaysRemaining
func NewBACnetConstructedDataDaysRemaining(daysRemaining BACnetApplicationTagSignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataDaysRemaining {
	_result := &_BACnetConstructedDataDaysRemaining{
		DaysRemaining:          daysRemaining,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataDaysRemaining(structType interface{}) BACnetConstructedDataDaysRemaining {
	if casted, ok := structType.(BACnetConstructedDataDaysRemaining); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataDaysRemaining); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataDaysRemaining) GetTypeName() string {
	return "BACnetConstructedDataDaysRemaining"
}

func (m *_BACnetConstructedDataDaysRemaining) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataDaysRemaining) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (daysRemaining)
	lengthInBits += m.DaysRemaining.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataDaysRemaining) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataDaysRemainingParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataDaysRemaining, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataDaysRemaining"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataDaysRemaining")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (daysRemaining)
	if pullErr := readBuffer.PullContext("daysRemaining"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for daysRemaining")
	}
	_daysRemaining, _daysRemainingErr := BACnetApplicationTagParse(readBuffer)
	if _daysRemainingErr != nil {
		return nil, errors.Wrap(_daysRemainingErr, "Error parsing 'daysRemaining' field")
	}
	daysRemaining := _daysRemaining.(BACnetApplicationTagSignedInteger)
	if closeErr := readBuffer.CloseContext("daysRemaining"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for daysRemaining")
	}

	// Virtual field
	_actualValue := daysRemaining
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataDaysRemaining"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataDaysRemaining")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataDaysRemaining{
		DaysRemaining:          daysRemaining,
		_BACnetConstructedData: &_BACnetConstructedData{},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataDaysRemaining) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataDaysRemaining"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataDaysRemaining")
		}

		// Simple Field (daysRemaining)
		if pushErr := writeBuffer.PushContext("daysRemaining"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for daysRemaining")
		}
		_daysRemainingErr := writeBuffer.WriteSerializable(m.GetDaysRemaining())
		if popErr := writeBuffer.PopContext("daysRemaining"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for daysRemaining")
		}
		if _daysRemainingErr != nil {
			return errors.Wrap(_daysRemainingErr, "Error serializing 'daysRemaining' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataDaysRemaining"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataDaysRemaining")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataDaysRemaining) isBACnetConstructedDataDaysRemaining() bool {
	return true
}

func (m *_BACnetConstructedDataDaysRemaining) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
