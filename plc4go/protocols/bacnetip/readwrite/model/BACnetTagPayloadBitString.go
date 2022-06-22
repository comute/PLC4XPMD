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

// BACnetTagPayloadBitString is the corresponding interface of BACnetTagPayloadBitString
type BACnetTagPayloadBitString interface {
	utils.LengthAware
	utils.Serializable
	// GetUnusedBits returns UnusedBits (property field)
	GetUnusedBits() uint8
	// GetData returns Data (property field)
	GetData() []bool
	// GetUnused returns Unused (property field)
	GetUnused() []bool
}

// BACnetTagPayloadBitStringExactly can be used when we want exactly this type and not a type which fulfills BACnetTagPayloadBitString.
// This is useful for switch cases.
type BACnetTagPayloadBitStringExactly interface {
	BACnetTagPayloadBitString
	isBACnetTagPayloadBitString() bool
}

// _BACnetTagPayloadBitString is the data-structure of this message
type _BACnetTagPayloadBitString struct {
	UnusedBits uint8
	Data       []bool
	Unused     []bool

	// Arguments.
	ActualLength uint32
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetTagPayloadBitString) GetUnusedBits() uint8 {
	return m.UnusedBits
}

func (m *_BACnetTagPayloadBitString) GetData() []bool {
	return m.Data
}

func (m *_BACnetTagPayloadBitString) GetUnused() []bool {
	return m.Unused
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetTagPayloadBitString factory function for _BACnetTagPayloadBitString
func NewBACnetTagPayloadBitString(unusedBits uint8, data []bool, unused []bool, actualLength uint32) *_BACnetTagPayloadBitString {
	return &_BACnetTagPayloadBitString{UnusedBits: unusedBits, Data: data, Unused: unused, ActualLength: actualLength}
}

// Deprecated: use the interface for direct cast
func CastBACnetTagPayloadBitString(structType interface{}) BACnetTagPayloadBitString {
	if casted, ok := structType.(BACnetTagPayloadBitString); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetTagPayloadBitString); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetTagPayloadBitString) GetTypeName() string {
	return "BACnetTagPayloadBitString"
}

func (m *_BACnetTagPayloadBitString) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetTagPayloadBitString) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (unusedBits)
	lengthInBits += 8

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 1 * uint16(len(m.Data))
	}

	// Array field
	if len(m.Unused) > 0 {
		lengthInBits += 1 * uint16(len(m.Unused))
	}

	return lengthInBits
}

func (m *_BACnetTagPayloadBitString) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetTagPayloadBitStringParse(readBuffer utils.ReadBuffer, actualLength uint32) (BACnetTagPayloadBitString, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetTagPayloadBitString"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetTagPayloadBitString")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (unusedBits)
	_unusedBits, _unusedBitsErr := readBuffer.ReadUint8("unusedBits", 8)
	if _unusedBitsErr != nil {
		return nil, errors.Wrap(_unusedBitsErr, "Error parsing 'unusedBits' field")
	}
	unusedBits := _unusedBits

	// Array field (data)
	if pullErr := readBuffer.PullContext("data", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for data")
	}
	// Count array
	data := make([]bool, uint16(uint16(uint16(uint16(uint16(actualLength)-uint16(uint16(1))))*uint16(uint16(8))))-uint16(unusedBits))
	{
		for curItem := uint16(0); curItem < uint16(uint16(uint16(uint16(uint16(uint16(actualLength)-uint16(uint16(1))))*uint16(uint16(8))))-uint16(unusedBits)); curItem++ {
			_item, _err := readBuffer.ReadBit("")
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'data' field")
			}
			data[curItem] = _item
		}
	}
	if closeErr := readBuffer.CloseContext("data", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for data")
	}

	// Array field (unused)
	if pullErr := readBuffer.PullContext("unused", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for unused")
	}
	// Count array
	unused := make([]bool, unusedBits)
	{
		for curItem := uint16(0); curItem < uint16(unusedBits); curItem++ {
			_item, _err := readBuffer.ReadBit("")
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'unused' field")
			}
			unused[curItem] = _item
		}
	}
	if closeErr := readBuffer.CloseContext("unused", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for unused")
	}

	if closeErr := readBuffer.CloseContext("BACnetTagPayloadBitString"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetTagPayloadBitString")
	}

	// Create the instance
	return NewBACnetTagPayloadBitString(unusedBits, data, unused, actualLength), nil
}

func (m *_BACnetTagPayloadBitString) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetTagPayloadBitString"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetTagPayloadBitString")
	}

	// Simple Field (unusedBits)
	unusedBits := uint8(m.GetUnusedBits())
	_unusedBitsErr := writeBuffer.WriteUint8("unusedBits", 8, (unusedBits))
	if _unusedBitsErr != nil {
		return errors.Wrap(_unusedBitsErr, "Error serializing 'unusedBits' field")
	}

	// Array Field (data)
	if m.GetData() != nil {
		if pushErr := writeBuffer.PushContext("data", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for data")
		}
		for _, _element := range m.GetData() {
			_elementErr := writeBuffer.WriteBit("", _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'data' field")
			}
		}
		if popErr := writeBuffer.PopContext("data", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for data")
		}
	}

	// Array Field (unused)
	if m.GetUnused() != nil {
		if pushErr := writeBuffer.PushContext("unused", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for unused")
		}
		for _, _element := range m.GetUnused() {
			_elementErr := writeBuffer.WriteBit("", _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'unused' field")
			}
		}
		if popErr := writeBuffer.PopContext("unused", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for unused")
		}
	}

	if popErr := writeBuffer.PopContext("BACnetTagPayloadBitString"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetTagPayloadBitString")
	}
	return nil
}

func (m *_BACnetTagPayloadBitString) isBACnetTagPayloadBitString() bool {
	return true
}

func (m *_BACnetTagPayloadBitString) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
