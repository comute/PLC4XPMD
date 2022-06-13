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

// BACnetPriorityValueConstructedValue is the data-structure of this message
type BACnetPriorityValueConstructedValue struct {
	*BACnetPriorityValue
	ConstructedValue *BACnetConstructedData

	// Arguments.
	ObjectTypeArgument BACnetObjectType
}

// IBACnetPriorityValueConstructedValue is the corresponding interface of BACnetPriorityValueConstructedValue
type IBACnetPriorityValueConstructedValue interface {
	IBACnetPriorityValue
	// GetConstructedValue returns ConstructedValue (property field)
	GetConstructedValue() *BACnetConstructedData
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

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetPriorityValueConstructedValue) InitializeParent(parent *BACnetPriorityValue, peekedTagHeader *BACnetTagHeader) {
	m.BACnetPriorityValue.PeekedTagHeader = peekedTagHeader
}

func (m *BACnetPriorityValueConstructedValue) GetParent() *BACnetPriorityValue {
	return m.BACnetPriorityValue
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetPriorityValueConstructedValue) GetConstructedValue() *BACnetConstructedData {
	return m.ConstructedValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPriorityValueConstructedValue factory function for BACnetPriorityValueConstructedValue
func NewBACnetPriorityValueConstructedValue(constructedValue *BACnetConstructedData, peekedTagHeader *BACnetTagHeader, objectTypeArgument BACnetObjectType) *BACnetPriorityValueConstructedValue {
	_result := &BACnetPriorityValueConstructedValue{
		ConstructedValue:    constructedValue,
		BACnetPriorityValue: NewBACnetPriorityValue(peekedTagHeader, objectTypeArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetPriorityValueConstructedValue(structType interface{}) *BACnetPriorityValueConstructedValue {
	if casted, ok := structType.(BACnetPriorityValueConstructedValue); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetPriorityValueConstructedValue); ok {
		return casted
	}
	if casted, ok := structType.(BACnetPriorityValue); ok {
		return CastBACnetPriorityValueConstructedValue(casted.Child)
	}
	if casted, ok := structType.(*BACnetPriorityValue); ok {
		return CastBACnetPriorityValueConstructedValue(casted.Child)
	}
	return nil
}

func (m *BACnetPriorityValueConstructedValue) GetTypeName() string {
	return "BACnetPriorityValueConstructedValue"
}

func (m *BACnetPriorityValueConstructedValue) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetPriorityValueConstructedValue) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (constructedValue)
	lengthInBits += m.ConstructedValue.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetPriorityValueConstructedValue) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetPriorityValueConstructedValueParse(readBuffer utils.ReadBuffer, objectTypeArgument BACnetObjectType) (*BACnetPriorityValueConstructedValue, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPriorityValueConstructedValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPriorityValueConstructedValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (constructedValue)
	if pullErr := readBuffer.PullContext("constructedValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for constructedValue")
	}
	_constructedValue, _constructedValueErr := BACnetConstructedDataParse(readBuffer, uint8(uint8(0)), BACnetObjectType(objectTypeArgument), BACnetPropertyIdentifier(BACnetPropertyIdentifier_VENDOR_PROPRIETARY_VALUE), nil)
	if _constructedValueErr != nil {
		return nil, errors.Wrap(_constructedValueErr, "Error parsing 'constructedValue' field")
	}
	constructedValue := CastBACnetConstructedData(_constructedValue)
	if closeErr := readBuffer.CloseContext("constructedValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for constructedValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetPriorityValueConstructedValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPriorityValueConstructedValue")
	}

	// Create a partially initialized instance
	_child := &BACnetPriorityValueConstructedValue{
		ConstructedValue:    CastBACnetConstructedData(constructedValue),
		BACnetPriorityValue: &BACnetPriorityValue{},
	}
	_child.BACnetPriorityValue.Child = _child
	return _child, nil
}

func (m *BACnetPriorityValueConstructedValue) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPriorityValueConstructedValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPriorityValueConstructedValue")
		}

		// Simple Field (constructedValue)
		if pushErr := writeBuffer.PushContext("constructedValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for constructedValue")
		}
		_constructedValueErr := writeBuffer.WriteSerializable(m.ConstructedValue)
		if popErr := writeBuffer.PopContext("constructedValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for constructedValue")
		}
		if _constructedValueErr != nil {
			return errors.Wrap(_constructedValueErr, "Error serializing 'constructedValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPriorityValueConstructedValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPriorityValueConstructedValue")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetPriorityValueConstructedValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
