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

// BACnetChannelValueBoolean is the data-structure of this message
type BACnetChannelValueBoolean struct {
	*BACnetChannelValue
	BooleanValue *BACnetApplicationTagBoolean
}

// IBACnetChannelValueBoolean is the corresponding interface of BACnetChannelValueBoolean
type IBACnetChannelValueBoolean interface {
	IBACnetChannelValue
	// GetBooleanValue returns BooleanValue (property field)
	GetBooleanValue() *BACnetApplicationTagBoolean
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

func (m *BACnetChannelValueBoolean) InitializeParent(parent *BACnetChannelValue, peekedTagHeader *BACnetTagHeader) {
	m.BACnetChannelValue.PeekedTagHeader = peekedTagHeader
}

func (m *BACnetChannelValueBoolean) GetParent() *BACnetChannelValue {
	return m.BACnetChannelValue
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetChannelValueBoolean) GetBooleanValue() *BACnetApplicationTagBoolean {
	return m.BooleanValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetChannelValueBoolean factory function for BACnetChannelValueBoolean
func NewBACnetChannelValueBoolean(booleanValue *BACnetApplicationTagBoolean, peekedTagHeader *BACnetTagHeader) *BACnetChannelValueBoolean {
	_result := &BACnetChannelValueBoolean{
		BooleanValue:       booleanValue,
		BACnetChannelValue: NewBACnetChannelValue(peekedTagHeader),
	}
	_result.Child = _result
	return _result
}

func CastBACnetChannelValueBoolean(structType interface{}) *BACnetChannelValueBoolean {
	if casted, ok := structType.(BACnetChannelValueBoolean); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetChannelValueBoolean); ok {
		return casted
	}
	if casted, ok := structType.(BACnetChannelValue); ok {
		return CastBACnetChannelValueBoolean(casted.Child)
	}
	if casted, ok := structType.(*BACnetChannelValue); ok {
		return CastBACnetChannelValueBoolean(casted.Child)
	}
	return nil
}

func (m *BACnetChannelValueBoolean) GetTypeName() string {
	return "BACnetChannelValueBoolean"
}

func (m *BACnetChannelValueBoolean) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetChannelValueBoolean) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (booleanValue)
	lengthInBits += m.BooleanValue.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetChannelValueBoolean) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetChannelValueBooleanParse(readBuffer utils.ReadBuffer) (*BACnetChannelValueBoolean, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetChannelValueBoolean"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetChannelValueBoolean")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (booleanValue)
	if pullErr := readBuffer.PullContext("booleanValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for booleanValue")
	}
	_booleanValue, _booleanValueErr := BACnetApplicationTagParse(readBuffer)
	if _booleanValueErr != nil {
		return nil, errors.Wrap(_booleanValueErr, "Error parsing 'booleanValue' field")
	}
	booleanValue := CastBACnetApplicationTagBoolean(_booleanValue)
	if closeErr := readBuffer.CloseContext("booleanValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for booleanValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetChannelValueBoolean"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetChannelValueBoolean")
	}

	// Create a partially initialized instance
	_child := &BACnetChannelValueBoolean{
		BooleanValue:       CastBACnetApplicationTagBoolean(booleanValue),
		BACnetChannelValue: &BACnetChannelValue{},
	}
	_child.BACnetChannelValue.Child = _child
	return _child, nil
}

func (m *BACnetChannelValueBoolean) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetChannelValueBoolean"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetChannelValueBoolean")
		}

		// Simple Field (booleanValue)
		if pushErr := writeBuffer.PushContext("booleanValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for booleanValue")
		}
		_booleanValueErr := writeBuffer.WriteSerializable(m.BooleanValue)
		if popErr := writeBuffer.PopContext("booleanValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for booleanValue")
		}
		if _booleanValueErr != nil {
			return errors.Wrap(_booleanValueErr, "Error serializing 'booleanValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetChannelValueBoolean"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetChannelValueBoolean")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetChannelValueBoolean) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
