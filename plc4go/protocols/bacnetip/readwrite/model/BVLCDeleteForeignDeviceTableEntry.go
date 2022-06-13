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

// BVLCDeleteForeignDeviceTableEntry is the data-structure of this message
type BVLCDeleteForeignDeviceTableEntry struct {
	*BVLC
	Ip   []uint8
	Port uint16
}

// IBVLCDeleteForeignDeviceTableEntry is the corresponding interface of BVLCDeleteForeignDeviceTableEntry
type IBVLCDeleteForeignDeviceTableEntry interface {
	IBVLC
	// GetIp returns Ip (property field)
	GetIp() []uint8
	// GetPort returns Port (property field)
	GetPort() uint16
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

func (m *BVLCDeleteForeignDeviceTableEntry) GetBvlcFunction() uint8 {
	return 0x08
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BVLCDeleteForeignDeviceTableEntry) InitializeParent(parent *BVLC) {}

func (m *BVLCDeleteForeignDeviceTableEntry) GetParent() *BVLC {
	return m.BVLC
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BVLCDeleteForeignDeviceTableEntry) GetIp() []uint8 {
	return m.Ip
}

func (m *BVLCDeleteForeignDeviceTableEntry) GetPort() uint16 {
	return m.Port
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBVLCDeleteForeignDeviceTableEntry factory function for BVLCDeleteForeignDeviceTableEntry
func NewBVLCDeleteForeignDeviceTableEntry(ip []uint8, port uint16) *BVLCDeleteForeignDeviceTableEntry {
	_result := &BVLCDeleteForeignDeviceTableEntry{
		Ip:   ip,
		Port: port,
		BVLC: NewBVLC(),
	}
	_result.Child = _result
	return _result
}

func CastBVLCDeleteForeignDeviceTableEntry(structType interface{}) *BVLCDeleteForeignDeviceTableEntry {
	if casted, ok := structType.(BVLCDeleteForeignDeviceTableEntry); ok {
		return &casted
	}
	if casted, ok := structType.(*BVLCDeleteForeignDeviceTableEntry); ok {
		return casted
	}
	if casted, ok := structType.(BVLC); ok {
		return CastBVLCDeleteForeignDeviceTableEntry(casted.Child)
	}
	if casted, ok := structType.(*BVLC); ok {
		return CastBVLCDeleteForeignDeviceTableEntry(casted.Child)
	}
	return nil
}

func (m *BVLCDeleteForeignDeviceTableEntry) GetTypeName() string {
	return "BVLCDeleteForeignDeviceTableEntry"
}

func (m *BVLCDeleteForeignDeviceTableEntry) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BVLCDeleteForeignDeviceTableEntry) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.Ip) > 0 {
		lengthInBits += 8 * uint16(len(m.Ip))
	}

	// Simple field (port)
	lengthInBits += 16

	return lengthInBits
}

func (m *BVLCDeleteForeignDeviceTableEntry) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BVLCDeleteForeignDeviceTableEntryParse(readBuffer utils.ReadBuffer) (*BVLCDeleteForeignDeviceTableEntry, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BVLCDeleteForeignDeviceTableEntry"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BVLCDeleteForeignDeviceTableEntry")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (ip)
	if pullErr := readBuffer.PullContext("ip", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ip")
	}
	// Count array
	ip := make([]uint8, uint16(4))
	{
		for curItem := uint16(0); curItem < uint16(uint16(4)); curItem++ {
			_item, _err := readBuffer.ReadUint8("", 8)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'ip' field")
			}
			ip[curItem] = _item
		}
	}
	if closeErr := readBuffer.CloseContext("ip", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ip")
	}

	// Simple Field (port)
	_port, _portErr := readBuffer.ReadUint16("port", 16)
	if _portErr != nil {
		return nil, errors.Wrap(_portErr, "Error parsing 'port' field")
	}
	port := _port

	if closeErr := readBuffer.CloseContext("BVLCDeleteForeignDeviceTableEntry"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BVLCDeleteForeignDeviceTableEntry")
	}

	// Create a partially initialized instance
	_child := &BVLCDeleteForeignDeviceTableEntry{
		Ip:   ip,
		Port: port,
		BVLC: &BVLC{},
	}
	_child.BVLC.Child = _child
	return _child, nil
}

func (m *BVLCDeleteForeignDeviceTableEntry) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BVLCDeleteForeignDeviceTableEntry"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BVLCDeleteForeignDeviceTableEntry")
		}

		// Array Field (ip)
		if m.Ip != nil {
			if pushErr := writeBuffer.PushContext("ip", utils.WithRenderAsList(true)); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for ip")
			}
			for _, _element := range m.Ip {
				_elementErr := writeBuffer.WriteUint8("", 8, _element)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'ip' field")
				}
			}
			if popErr := writeBuffer.PopContext("ip", utils.WithRenderAsList(true)); popErr != nil {
				return errors.Wrap(popErr, "Error popping for ip")
			}
		}

		// Simple Field (port)
		port := uint16(m.Port)
		_portErr := writeBuffer.WriteUint16("port", 16, (port))
		if _portErr != nil {
			return errors.Wrap(_portErr, "Error serializing 'port' field")
		}

		if popErr := writeBuffer.PopContext("BVLCDeleteForeignDeviceTableEntry"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BVLCDeleteForeignDeviceTableEntry")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BVLCDeleteForeignDeviceTableEntry) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
