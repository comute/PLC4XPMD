/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
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

// BACnetConstructedDataMaxActualValue is the data-structure of this message
type BACnetConstructedDataMaxActualValue struct {
	*BACnetConstructedData
	MaxActualValue *BACnetApplicationTagReal

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataMaxActualValue is the corresponding interface of BACnetConstructedDataMaxActualValue
type IBACnetConstructedDataMaxActualValue interface {
	IBACnetConstructedData
	// GetMaxActualValue returns MaxActualValue (property field)
	GetMaxActualValue() *BACnetApplicationTagReal
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *BACnetConstructedDataMaxActualValue) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataMaxActualValue) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_MAX_ACTUAL_VALUE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataMaxActualValue) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataMaxActualValue) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataMaxActualValue) GetMaxActualValue() *BACnetApplicationTagReal {
	return m.MaxActualValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataMaxActualValue factory function for BACnetConstructedDataMaxActualValue
func NewBACnetConstructedDataMaxActualValue(maxActualValue *BACnetApplicationTagReal, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataMaxActualValue {
	_result := &BACnetConstructedDataMaxActualValue{
		MaxActualValue:        maxActualValue,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataMaxActualValue(structType interface{}) *BACnetConstructedDataMaxActualValue {
	if casted, ok := structType.(BACnetConstructedDataMaxActualValue); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataMaxActualValue); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataMaxActualValue(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataMaxActualValue(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataMaxActualValue) GetTypeName() string {
	return "BACnetConstructedDataMaxActualValue"
}

func (m *BACnetConstructedDataMaxActualValue) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataMaxActualValue) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (maxActualValue)
	lengthInBits += m.MaxActualValue.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataMaxActualValue) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataMaxActualValueParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataMaxActualValue, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataMaxActualValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataMaxActualValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (maxActualValue)
	if pullErr := readBuffer.PullContext("maxActualValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for maxActualValue")
	}
	_maxActualValue, _maxActualValueErr := BACnetApplicationTagParse(readBuffer)
	if _maxActualValueErr != nil {
		return nil, errors.Wrap(_maxActualValueErr, "Error parsing 'maxActualValue' field")
	}
	maxActualValue := CastBACnetApplicationTagReal(_maxActualValue)
	if closeErr := readBuffer.CloseContext("maxActualValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for maxActualValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataMaxActualValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataMaxActualValue")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataMaxActualValue{
		MaxActualValue:        CastBACnetApplicationTagReal(maxActualValue),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataMaxActualValue) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataMaxActualValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataMaxActualValue")
		}

		// Simple Field (maxActualValue)
		if pushErr := writeBuffer.PushContext("maxActualValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for maxActualValue")
		}
		_maxActualValueErr := m.MaxActualValue.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("maxActualValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for maxActualValue")
		}
		if _maxActualValueErr != nil {
			return errors.Wrap(_maxActualValueErr, "Error serializing 'maxActualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataMaxActualValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataMaxActualValue")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataMaxActualValue) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
