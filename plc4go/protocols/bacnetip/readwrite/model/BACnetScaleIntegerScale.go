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

// BACnetScaleIntegerScale is the data-structure of this message
type BACnetScaleIntegerScale struct {
	*BACnetScale
	IntegerScale *BACnetContextTagSignedInteger
}

// IBACnetScaleIntegerScale is the corresponding interface of BACnetScaleIntegerScale
type IBACnetScaleIntegerScale interface {
	IBACnetScale
	// GetIntegerScale returns IntegerScale (property field)
	GetIntegerScale() *BACnetContextTagSignedInteger
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

func (m *BACnetScaleIntegerScale) InitializeParent(parent *BACnetScale, peekedTagHeader *BACnetTagHeader) {
	m.BACnetScale.PeekedTagHeader = peekedTagHeader
}

func (m *BACnetScaleIntegerScale) GetParent() *BACnetScale {
	return m.BACnetScale
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetScaleIntegerScale) GetIntegerScale() *BACnetContextTagSignedInteger {
	return m.IntegerScale
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetScaleIntegerScale factory function for BACnetScaleIntegerScale
func NewBACnetScaleIntegerScale(integerScale *BACnetContextTagSignedInteger, peekedTagHeader *BACnetTagHeader) *BACnetScaleIntegerScale {
	_result := &BACnetScaleIntegerScale{
		IntegerScale: integerScale,
		BACnetScale:  NewBACnetScale(peekedTagHeader),
	}
	_result.Child = _result
	return _result
}

func CastBACnetScaleIntegerScale(structType interface{}) *BACnetScaleIntegerScale {
	if casted, ok := structType.(BACnetScaleIntegerScale); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetScaleIntegerScale); ok {
		return casted
	}
	if casted, ok := structType.(BACnetScale); ok {
		return CastBACnetScaleIntegerScale(casted.Child)
	}
	if casted, ok := structType.(*BACnetScale); ok {
		return CastBACnetScaleIntegerScale(casted.Child)
	}
	return nil
}

func (m *BACnetScaleIntegerScale) GetTypeName() string {
	return "BACnetScaleIntegerScale"
}

func (m *BACnetScaleIntegerScale) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetScaleIntegerScale) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (integerScale)
	lengthInBits += m.IntegerScale.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetScaleIntegerScale) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetScaleIntegerScaleParse(readBuffer utils.ReadBuffer) (*BACnetScaleIntegerScale, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetScaleIntegerScale"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetScaleIntegerScale")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (integerScale)
	if pullErr := readBuffer.PullContext("integerScale"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for integerScale")
	}
	_integerScale, _integerScaleErr := BACnetContextTagParse(readBuffer, uint8(uint8(1)), BACnetDataType(BACnetDataType_SIGNED_INTEGER))
	if _integerScaleErr != nil {
		return nil, errors.Wrap(_integerScaleErr, "Error parsing 'integerScale' field")
	}
	integerScale := CastBACnetContextTagSignedInteger(_integerScale)
	if closeErr := readBuffer.CloseContext("integerScale"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for integerScale")
	}

	if closeErr := readBuffer.CloseContext("BACnetScaleIntegerScale"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetScaleIntegerScale")
	}

	// Create a partially initialized instance
	_child := &BACnetScaleIntegerScale{
		IntegerScale: CastBACnetContextTagSignedInteger(integerScale),
		BACnetScale:  &BACnetScale{},
	}
	_child.BACnetScale.Child = _child
	return _child, nil
}

func (m *BACnetScaleIntegerScale) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetScaleIntegerScale"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetScaleIntegerScale")
		}

		// Simple Field (integerScale)
		if pushErr := writeBuffer.PushContext("integerScale"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for integerScale")
		}
		_integerScaleErr := writeBuffer.WriteSerializable(m.IntegerScale)
		if popErr := writeBuffer.PopContext("integerScale"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for integerScale")
		}
		if _integerScaleErr != nil {
			return errors.Wrap(_integerScaleErr, "Error serializing 'integerScale' field")
		}

		if popErr := writeBuffer.PopContext("BACnetScaleIntegerScale"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetScaleIntegerScale")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetScaleIntegerScale) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
