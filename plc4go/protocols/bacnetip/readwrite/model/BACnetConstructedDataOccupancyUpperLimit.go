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

// BACnetConstructedDataOccupancyUpperLimit is the corresponding interface of BACnetConstructedDataOccupancyUpperLimit
type BACnetConstructedDataOccupancyUpperLimit interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetOccupancyUpperLimit returns OccupancyUpperLimit (property field)
	GetOccupancyUpperLimit() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataOccupancyUpperLimitExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataOccupancyUpperLimit.
// This is useful for switch cases.
type BACnetConstructedDataOccupancyUpperLimitExactly interface {
	BACnetConstructedDataOccupancyUpperLimit
	isBACnetConstructedDataOccupancyUpperLimit() bool
}

// _BACnetConstructedDataOccupancyUpperLimit is the data-structure of this message
type _BACnetConstructedDataOccupancyUpperLimit struct {
	*_BACnetConstructedData
	OccupancyUpperLimit BACnetApplicationTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataOccupancyUpperLimit) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataOccupancyUpperLimit) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_OCCUPANCY_UPPER_LIMIT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataOccupancyUpperLimit) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataOccupancyUpperLimit) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataOccupancyUpperLimit) GetOccupancyUpperLimit() BACnetApplicationTagUnsignedInteger {
	return m.OccupancyUpperLimit
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataOccupancyUpperLimit) GetActualValue() BACnetApplicationTagUnsignedInteger {
	return CastBACnetApplicationTagUnsignedInteger(m.GetOccupancyUpperLimit())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataOccupancyUpperLimit factory function for _BACnetConstructedDataOccupancyUpperLimit
func NewBACnetConstructedDataOccupancyUpperLimit(occupancyUpperLimit BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataOccupancyUpperLimit {
	_result := &_BACnetConstructedDataOccupancyUpperLimit{
		OccupancyUpperLimit:    occupancyUpperLimit,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataOccupancyUpperLimit(structType interface{}) BACnetConstructedDataOccupancyUpperLimit {
	if casted, ok := structType.(BACnetConstructedDataOccupancyUpperLimit); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataOccupancyUpperLimit); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataOccupancyUpperLimit) GetTypeName() string {
	return "BACnetConstructedDataOccupancyUpperLimit"
}

func (m *_BACnetConstructedDataOccupancyUpperLimit) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataOccupancyUpperLimit) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (occupancyUpperLimit)
	lengthInBits += m.OccupancyUpperLimit.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataOccupancyUpperLimit) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataOccupancyUpperLimitParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataOccupancyUpperLimit, error) {
	return BACnetConstructedDataOccupancyUpperLimitParseWithBuffer(utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataOccupancyUpperLimitParseWithBuffer(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataOccupancyUpperLimit, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataOccupancyUpperLimit"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataOccupancyUpperLimit")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (occupancyUpperLimit)
	if pullErr := readBuffer.PullContext("occupancyUpperLimit"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for occupancyUpperLimit")
	}
	_occupancyUpperLimit, _occupancyUpperLimitErr := BACnetApplicationTagParseWithBuffer(readBuffer)
	if _occupancyUpperLimitErr != nil {
		return nil, errors.Wrap(_occupancyUpperLimitErr, "Error parsing 'occupancyUpperLimit' field of BACnetConstructedDataOccupancyUpperLimit")
	}
	occupancyUpperLimit := _occupancyUpperLimit.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("occupancyUpperLimit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for occupancyUpperLimit")
	}

	// Virtual field
	_actualValue := occupancyUpperLimit
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataOccupancyUpperLimit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataOccupancyUpperLimit")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataOccupancyUpperLimit{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		OccupancyUpperLimit: occupancyUpperLimit,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataOccupancyUpperLimit) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes())))
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataOccupancyUpperLimit) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataOccupancyUpperLimit"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataOccupancyUpperLimit")
		}

		// Simple Field (occupancyUpperLimit)
		if pushErr := writeBuffer.PushContext("occupancyUpperLimit"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for occupancyUpperLimit")
		}
		_occupancyUpperLimitErr := writeBuffer.WriteSerializable(m.GetOccupancyUpperLimit())
		if popErr := writeBuffer.PopContext("occupancyUpperLimit"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for occupancyUpperLimit")
		}
		if _occupancyUpperLimitErr != nil {
			return errors.Wrap(_occupancyUpperLimitErr, "Error serializing 'occupancyUpperLimit' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataOccupancyUpperLimit"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataOccupancyUpperLimit")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataOccupancyUpperLimit) isBACnetConstructedDataOccupancyUpperLimit() bool {
	return true
}

func (m *_BACnetConstructedDataOccupancyUpperLimit) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
