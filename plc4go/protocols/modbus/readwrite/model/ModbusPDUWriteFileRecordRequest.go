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

// ModbusPDUWriteFileRecordRequest is the data-structure of this message
type ModbusPDUWriteFileRecordRequest struct {
	*ModbusPDU
	Items []*ModbusPDUWriteFileRecordRequestItem
}

// IModbusPDUWriteFileRecordRequest is the corresponding interface of ModbusPDUWriteFileRecordRequest
type IModbusPDUWriteFileRecordRequest interface {
	IModbusPDU
	// GetItems returns Items (property field)
	GetItems() []*ModbusPDUWriteFileRecordRequestItem
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

func (m *ModbusPDUWriteFileRecordRequest) GetErrorFlag() bool {
	return bool(false)
}

func (m *ModbusPDUWriteFileRecordRequest) GetFunctionFlag() uint8 {
	return 0x15
}

func (m *ModbusPDUWriteFileRecordRequest) GetResponse() bool {
	return bool(false)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *ModbusPDUWriteFileRecordRequest) InitializeParent(parent *ModbusPDU) {}

func (m *ModbusPDUWriteFileRecordRequest) GetParent() *ModbusPDU {
	return m.ModbusPDU
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *ModbusPDUWriteFileRecordRequest) GetItems() []*ModbusPDUWriteFileRecordRequestItem {
	return m.Items
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewModbusPDUWriteFileRecordRequest factory function for ModbusPDUWriteFileRecordRequest
func NewModbusPDUWriteFileRecordRequest(items []*ModbusPDUWriteFileRecordRequestItem) *ModbusPDUWriteFileRecordRequest {
	_result := &ModbusPDUWriteFileRecordRequest{
		Items:     items,
		ModbusPDU: NewModbusPDU(),
	}
	_result.Child = _result
	return _result
}

func CastModbusPDUWriteFileRecordRequest(structType interface{}) *ModbusPDUWriteFileRecordRequest {
	if casted, ok := structType.(ModbusPDUWriteFileRecordRequest); ok {
		return &casted
	}
	if casted, ok := structType.(*ModbusPDUWriteFileRecordRequest); ok {
		return casted
	}
	if casted, ok := structType.(ModbusPDU); ok {
		return CastModbusPDUWriteFileRecordRequest(casted.Child)
	}
	if casted, ok := structType.(*ModbusPDU); ok {
		return CastModbusPDUWriteFileRecordRequest(casted.Child)
	}
	return nil
}

func (m *ModbusPDUWriteFileRecordRequest) GetTypeName() string {
	return "ModbusPDUWriteFileRecordRequest"
}

func (m *ModbusPDUWriteFileRecordRequest) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *ModbusPDUWriteFileRecordRequest) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Implicit Field (byteCount)
	lengthInBits += 8

	// Array field
	if len(m.Items) > 0 {
		for _, element := range m.Items {
			lengthInBits += element.GetLengthInBits()
		}
	}

	return lengthInBits
}

func (m *ModbusPDUWriteFileRecordRequest) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ModbusPDUWriteFileRecordRequestParse(readBuffer utils.ReadBuffer, response bool) (*ModbusPDUWriteFileRecordRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModbusPDUWriteFileRecordRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusPDUWriteFileRecordRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Implicit Field (byteCount) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	byteCount, _byteCountErr := readBuffer.ReadUint8("byteCount", 8)
	_ = byteCount
	if _byteCountErr != nil {
		return nil, errors.Wrap(_byteCountErr, "Error parsing 'byteCount' field")
	}

	// Array field (items)
	if pullErr := readBuffer.PullContext("items", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for items")
	}
	// Length array
	items := make([]*ModbusPDUWriteFileRecordRequestItem, 0)
	{
		_itemsLength := byteCount
		_itemsEndPos := positionAware.GetPos() + uint16(_itemsLength)
		for positionAware.GetPos() < _itemsEndPos {
			_item, _err := ModbusPDUWriteFileRecordRequestItemParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'items' field")
			}
			items = append(items, CastModbusPDUWriteFileRecordRequestItem(_item))
		}
	}
	if closeErr := readBuffer.CloseContext("items", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for items")
	}

	if closeErr := readBuffer.CloseContext("ModbusPDUWriteFileRecordRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusPDUWriteFileRecordRequest")
	}

	// Create a partially initialized instance
	_child := &ModbusPDUWriteFileRecordRequest{
		Items:     items,
		ModbusPDU: &ModbusPDU{},
	}
	_child.ModbusPDU.Child = _child
	return _child, nil
}

func (m *ModbusPDUWriteFileRecordRequest) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	itemsArraySizeInBytes := func(items []*ModbusPDUWriteFileRecordRequestItem) uint32 {
		var sizeInBytes uint32 = 0
		for _, v := range items {
			sizeInBytes += uint32(v.GetLengthInBytes())
		}
		return sizeInBytes
	}
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUWriteFileRecordRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ModbusPDUWriteFileRecordRequest")
		}

		// Implicit Field (byteCount) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
		byteCount := uint8(uint8(itemsArraySizeInBytes(m.GetItems())))
		_byteCountErr := writeBuffer.WriteUint8("byteCount", 8, (byteCount))
		if _byteCountErr != nil {
			return errors.Wrap(_byteCountErr, "Error serializing 'byteCount' field")
		}

		// Array Field (items)
		if m.Items != nil {
			if pushErr := writeBuffer.PushContext("items", utils.WithRenderAsList(true)); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for items")
			}
			for _, _element := range m.Items {
				_elementErr := _element.Serialize(writeBuffer)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'items' field")
				}
			}
			if popErr := writeBuffer.PopContext("items", utils.WithRenderAsList(true)); popErr != nil {
				return errors.Wrap(popErr, "Error popping for items")
			}
		}

		if popErr := writeBuffer.PopContext("ModbusPDUWriteFileRecordRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ModbusPDUWriteFileRecordRequest")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *ModbusPDUWriteFileRecordRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
