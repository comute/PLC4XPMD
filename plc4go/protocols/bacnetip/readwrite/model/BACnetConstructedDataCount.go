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

// BACnetConstructedDataCount is the corresponding interface of BACnetConstructedDataCount
type BACnetConstructedDataCount interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetCount returns Count (property field)
	GetCount() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataCountExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataCount.
// This is useful for switch cases.
type BACnetConstructedDataCountExactly interface {
	BACnetConstructedDataCount
	isBACnetConstructedDataCount() bool
}

// _BACnetConstructedDataCount is the data-structure of this message
type _BACnetConstructedDataCount struct {
	*_BACnetConstructedData
	Count BACnetApplicationTagUnsignedInteger

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument BACnetTagPayloadUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataCount) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataCount) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_COUNT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataCount) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataCount) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataCount) GetCount() BACnetApplicationTagUnsignedInteger {
	return m.Count
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataCount) GetActualValue() BACnetApplicationTagUnsignedInteger {
	return CastBACnetApplicationTagUnsignedInteger(m.GetCount())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataCount factory function for _BACnetConstructedDataCount
func NewBACnetConstructedDataCount(count BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataCount {
	_result := &_BACnetConstructedDataCount{
		Count:                  count,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataCount(structType interface{}) BACnetConstructedDataCount {
	if casted, ok := structType.(BACnetConstructedDataCount); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataCount); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataCount) GetTypeName() string {
	return "BACnetConstructedDataCount"
}

func (m *_BACnetConstructedDataCount) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataCount) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (count)
	lengthInBits += m.Count.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataCount) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataCountParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataCount, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataCount"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataCount")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (count)
	if pullErr := readBuffer.PullContext("count"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for count")
	}
	_count, _countErr := BACnetApplicationTagParse(readBuffer)
	if _countErr != nil {
		return nil, errors.Wrap(_countErr, "Error parsing 'count' field")
	}
	count := _count.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("count"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for count")
	}

	// Virtual field
	_actualValue := count
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataCount"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataCount")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataCount{
		Count:                  count,
		_BACnetConstructedData: &_BACnetConstructedData{},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataCount) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataCount"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataCount")
		}

		// Simple Field (count)
		if pushErr := writeBuffer.PushContext("count"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for count")
		}
		_countErr := writeBuffer.WriteSerializable(m.GetCount())
		if popErr := writeBuffer.PopContext("count"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for count")
		}
		if _countErr != nil {
			return errors.Wrap(_countErr, "Error serializing 'count' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataCount"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataCount")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataCount) isBACnetConstructedDataCount() bool {
	return true
}

func (m *_BACnetConstructedDataCount) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
