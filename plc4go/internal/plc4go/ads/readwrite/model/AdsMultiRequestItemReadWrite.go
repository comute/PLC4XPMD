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
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type AdsMultiRequestItemReadWrite struct {
	*AdsMultiRequestItem
	ItemIndexGroup  uint32
	ItemIndexOffset uint32
	ItemReadLength  uint32
	ItemWriteLength uint32
}

// The corresponding interface
type IAdsMultiRequestItemReadWrite interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *AdsMultiRequestItemReadWrite) IndexGroup() uint32 {
	return uint32(61570)
}

func (m *AdsMultiRequestItemReadWrite) InitializeParent(parent *AdsMultiRequestItem) {
}

func NewAdsMultiRequestItemReadWrite(itemIndexGroup uint32, itemIndexOffset uint32, itemReadLength uint32, itemWriteLength uint32) *AdsMultiRequestItem {
	child := &AdsMultiRequestItemReadWrite{
		ItemIndexGroup:      itemIndexGroup,
		ItemIndexOffset:     itemIndexOffset,
		ItemReadLength:      itemReadLength,
		ItemWriteLength:     itemWriteLength,
		AdsMultiRequestItem: NewAdsMultiRequestItem(),
	}
	child.Child = child
	return child.AdsMultiRequestItem
}

func CastAdsMultiRequestItemReadWrite(structType interface{}) *AdsMultiRequestItemReadWrite {
	castFunc := func(typ interface{}) *AdsMultiRequestItemReadWrite {
		if casted, ok := typ.(AdsMultiRequestItemReadWrite); ok {
			return &casted
		}
		if casted, ok := typ.(*AdsMultiRequestItemReadWrite); ok {
			return casted
		}
		if casted, ok := typ.(AdsMultiRequestItem); ok {
			return CastAdsMultiRequestItemReadWrite(casted.Child)
		}
		if casted, ok := typ.(*AdsMultiRequestItem); ok {
			return CastAdsMultiRequestItemReadWrite(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *AdsMultiRequestItemReadWrite) GetTypeName() string {
	return "AdsMultiRequestItemReadWrite"
}

func (m *AdsMultiRequestItemReadWrite) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *AdsMultiRequestItemReadWrite) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.ParentLengthInBits())

	// Simple field (itemIndexGroup)
	lengthInBits += 32

	// Simple field (itemIndexOffset)
	lengthInBits += 32

	// Simple field (itemReadLength)
	lengthInBits += 32

	// Simple field (itemWriteLength)
	lengthInBits += 32

	return lengthInBits
}

func (m *AdsMultiRequestItemReadWrite) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func AdsMultiRequestItemReadWriteParse(readBuffer utils.ReadBuffer, indexGroup uint32) (*AdsMultiRequestItem, error) {
	if pullErr := readBuffer.PullContext("AdsMultiRequestItemReadWrite"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (itemIndexGroup)
	_itemIndexGroup, _itemIndexGroupErr := readBuffer.ReadUint32("itemIndexGroup", 32)
	if _itemIndexGroupErr != nil {
		return nil, errors.Wrap(_itemIndexGroupErr, "Error parsing 'itemIndexGroup' field")
	}
	itemIndexGroup := _itemIndexGroup

	// Simple Field (itemIndexOffset)
	_itemIndexOffset, _itemIndexOffsetErr := readBuffer.ReadUint32("itemIndexOffset", 32)
	if _itemIndexOffsetErr != nil {
		return nil, errors.Wrap(_itemIndexOffsetErr, "Error parsing 'itemIndexOffset' field")
	}
	itemIndexOffset := _itemIndexOffset

	// Simple Field (itemReadLength)
	_itemReadLength, _itemReadLengthErr := readBuffer.ReadUint32("itemReadLength", 32)
	if _itemReadLengthErr != nil {
		return nil, errors.Wrap(_itemReadLengthErr, "Error parsing 'itemReadLength' field")
	}
	itemReadLength := _itemReadLength

	// Simple Field (itemWriteLength)
	_itemWriteLength, _itemWriteLengthErr := readBuffer.ReadUint32("itemWriteLength", 32)
	if _itemWriteLengthErr != nil {
		return nil, errors.Wrap(_itemWriteLengthErr, "Error parsing 'itemWriteLength' field")
	}
	itemWriteLength := _itemWriteLength

	if closeErr := readBuffer.CloseContext("AdsMultiRequestItemReadWrite"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &AdsMultiRequestItemReadWrite{
		ItemIndexGroup:      itemIndexGroup,
		ItemIndexOffset:     itemIndexOffset,
		ItemReadLength:      itemReadLength,
		ItemWriteLength:     itemWriteLength,
		AdsMultiRequestItem: &AdsMultiRequestItem{},
	}
	_child.AdsMultiRequestItem.Child = _child
	return _child.AdsMultiRequestItem, nil
}

func (m *AdsMultiRequestItemReadWrite) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AdsMultiRequestItemReadWrite"); pushErr != nil {
			return pushErr
		}

		// Simple Field (itemIndexGroup)
		itemIndexGroup := uint32(m.ItemIndexGroup)
		_itemIndexGroupErr := writeBuffer.WriteUint32("itemIndexGroup", 32, (itemIndexGroup))
		if _itemIndexGroupErr != nil {
			return errors.Wrap(_itemIndexGroupErr, "Error serializing 'itemIndexGroup' field")
		}

		// Simple Field (itemIndexOffset)
		itemIndexOffset := uint32(m.ItemIndexOffset)
		_itemIndexOffsetErr := writeBuffer.WriteUint32("itemIndexOffset", 32, (itemIndexOffset))
		if _itemIndexOffsetErr != nil {
			return errors.Wrap(_itemIndexOffsetErr, "Error serializing 'itemIndexOffset' field")
		}

		// Simple Field (itemReadLength)
		itemReadLength := uint32(m.ItemReadLength)
		_itemReadLengthErr := writeBuffer.WriteUint32("itemReadLength", 32, (itemReadLength))
		if _itemReadLengthErr != nil {
			return errors.Wrap(_itemReadLengthErr, "Error serializing 'itemReadLength' field")
		}

		// Simple Field (itemWriteLength)
		itemWriteLength := uint32(m.ItemWriteLength)
		_itemWriteLengthErr := writeBuffer.WriteUint32("itemWriteLength", 32, (itemWriteLength))
		if _itemWriteLengthErr != nil {
			return errors.Wrap(_itemWriteLengthErr, "Error serializing 'itemWriteLength' field")
		}

		if popErr := writeBuffer.PopContext("AdsMultiRequestItemReadWrite"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *AdsMultiRequestItemReadWrite) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
