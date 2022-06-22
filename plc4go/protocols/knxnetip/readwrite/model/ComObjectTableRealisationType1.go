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

// ComObjectTableRealisationType1 is the corresponding interface of ComObjectTableRealisationType1
type ComObjectTableRealisationType1 interface {
	utils.LengthAware
	utils.Serializable
	ComObjectTable
	// GetNumEntries returns NumEntries (property field)
	GetNumEntries() uint8
	// GetRamFlagsTablePointer returns RamFlagsTablePointer (property field)
	GetRamFlagsTablePointer() uint8
	// GetComObjectDescriptors returns ComObjectDescriptors (property field)
	GetComObjectDescriptors() []GroupObjectDescriptorRealisationType1
}

// ComObjectTableRealisationType1Exactly can be used when we want exactly this type and not a type which fulfills ComObjectTableRealisationType1.
// This is useful for switch cases.
type ComObjectTableRealisationType1Exactly interface {
	ComObjectTableRealisationType1
	isComObjectTableRealisationType1() bool
}

// _ComObjectTableRealisationType1 is the data-structure of this message
type _ComObjectTableRealisationType1 struct {
	*_ComObjectTable
	NumEntries           uint8
	RamFlagsTablePointer uint8
	ComObjectDescriptors []GroupObjectDescriptorRealisationType1
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ComObjectTableRealisationType1) GetFirmwareType() FirmwareType {
	return FirmwareType_SYSTEM_1
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ComObjectTableRealisationType1) InitializeParent(parent ComObjectTable) {}

func (m *_ComObjectTableRealisationType1) GetParent() ComObjectTable {
	return m._ComObjectTable
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ComObjectTableRealisationType1) GetNumEntries() uint8 {
	return m.NumEntries
}

func (m *_ComObjectTableRealisationType1) GetRamFlagsTablePointer() uint8 {
	return m.RamFlagsTablePointer
}

func (m *_ComObjectTableRealisationType1) GetComObjectDescriptors() []GroupObjectDescriptorRealisationType1 {
	return m.ComObjectDescriptors
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewComObjectTableRealisationType1 factory function for _ComObjectTableRealisationType1
func NewComObjectTableRealisationType1(numEntries uint8, ramFlagsTablePointer uint8, comObjectDescriptors []GroupObjectDescriptorRealisationType1) *_ComObjectTableRealisationType1 {
	_result := &_ComObjectTableRealisationType1{
		NumEntries:           numEntries,
		RamFlagsTablePointer: ramFlagsTablePointer,
		ComObjectDescriptors: comObjectDescriptors,
		_ComObjectTable:      NewComObjectTable(),
	}
	_result._ComObjectTable._ComObjectTableChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastComObjectTableRealisationType1(structType interface{}) ComObjectTableRealisationType1 {
	if casted, ok := structType.(ComObjectTableRealisationType1); ok {
		return casted
	}
	if casted, ok := structType.(*ComObjectTableRealisationType1); ok {
		return *casted
	}
	return nil
}

func (m *_ComObjectTableRealisationType1) GetTypeName() string {
	return "ComObjectTableRealisationType1"
}

func (m *_ComObjectTableRealisationType1) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_ComObjectTableRealisationType1) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (numEntries)
	lengthInBits += 8

	// Simple field (ramFlagsTablePointer)
	lengthInBits += 8

	// Array field
	if len(m.ComObjectDescriptors) > 0 {
		for i, element := range m.ComObjectDescriptors {
			last := i == len(m.ComObjectDescriptors)-1
			lengthInBits += element.(interface{ GetLengthInBitsConditional(bool) uint16 }).GetLengthInBitsConditional(last)
		}
	}

	return lengthInBits
}

func (m *_ComObjectTableRealisationType1) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ComObjectTableRealisationType1Parse(readBuffer utils.ReadBuffer, firmwareType FirmwareType) (ComObjectTableRealisationType1, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ComObjectTableRealisationType1"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ComObjectTableRealisationType1")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (numEntries)
	_numEntries, _numEntriesErr := readBuffer.ReadUint8("numEntries", 8)
	if _numEntriesErr != nil {
		return nil, errors.Wrap(_numEntriesErr, "Error parsing 'numEntries' field")
	}
	numEntries := _numEntries

	// Simple Field (ramFlagsTablePointer)
	_ramFlagsTablePointer, _ramFlagsTablePointerErr := readBuffer.ReadUint8("ramFlagsTablePointer", 8)
	if _ramFlagsTablePointerErr != nil {
		return nil, errors.Wrap(_ramFlagsTablePointerErr, "Error parsing 'ramFlagsTablePointer' field")
	}
	ramFlagsTablePointer := _ramFlagsTablePointer

	// Array field (comObjectDescriptors)
	if pullErr := readBuffer.PullContext("comObjectDescriptors", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for comObjectDescriptors")
	}
	// Count array
	comObjectDescriptors := make([]GroupObjectDescriptorRealisationType1, numEntries)
	{
		for curItem := uint16(0); curItem < uint16(numEntries); curItem++ {
			_item, _err := GroupObjectDescriptorRealisationType1Parse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'comObjectDescriptors' field")
			}
			comObjectDescriptors[curItem] = _item.(GroupObjectDescriptorRealisationType1)
		}
	}
	if closeErr := readBuffer.CloseContext("comObjectDescriptors", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for comObjectDescriptors")
	}

	if closeErr := readBuffer.CloseContext("ComObjectTableRealisationType1"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ComObjectTableRealisationType1")
	}

	// Create a partially initialized instance
	_child := &_ComObjectTableRealisationType1{
		NumEntries:           numEntries,
		RamFlagsTablePointer: ramFlagsTablePointer,
		ComObjectDescriptors: comObjectDescriptors,
		_ComObjectTable:      &_ComObjectTable{},
	}
	_child._ComObjectTable._ComObjectTableChildRequirements = _child
	return _child, nil
}

func (m *_ComObjectTableRealisationType1) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ComObjectTableRealisationType1"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ComObjectTableRealisationType1")
		}

		// Simple Field (numEntries)
		numEntries := uint8(m.GetNumEntries())
		_numEntriesErr := writeBuffer.WriteUint8("numEntries", 8, (numEntries))
		if _numEntriesErr != nil {
			return errors.Wrap(_numEntriesErr, "Error serializing 'numEntries' field")
		}

		// Simple Field (ramFlagsTablePointer)
		ramFlagsTablePointer := uint8(m.GetRamFlagsTablePointer())
		_ramFlagsTablePointerErr := writeBuffer.WriteUint8("ramFlagsTablePointer", 8, (ramFlagsTablePointer))
		if _ramFlagsTablePointerErr != nil {
			return errors.Wrap(_ramFlagsTablePointerErr, "Error serializing 'ramFlagsTablePointer' field")
		}

		// Array Field (comObjectDescriptors)
		if m.GetComObjectDescriptors() != nil {
			if pushErr := writeBuffer.PushContext("comObjectDescriptors", utils.WithRenderAsList(true)); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for comObjectDescriptors")
			}
			for _, _element := range m.GetComObjectDescriptors() {
				_elementErr := writeBuffer.WriteSerializable(_element)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'comObjectDescriptors' field")
				}
			}
			if popErr := writeBuffer.PopContext("comObjectDescriptors", utils.WithRenderAsList(true)); popErr != nil {
				return errors.Wrap(popErr, "Error popping for comObjectDescriptors")
			}
		}

		if popErr := writeBuffer.PopContext("ComObjectTableRealisationType1"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ComObjectTableRealisationType1")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_ComObjectTableRealisationType1) isComObjectTableRealisationType1() bool {
	return true
}

func (m *_ComObjectTableRealisationType1) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
