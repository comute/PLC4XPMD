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
type TunnelingResponseDataBlock struct {
	CommunicationChannelId uint8
	SequenceCounter        uint8
	Status                 Status
}

// The corresponding interface
type ITunnelingResponseDataBlock interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

func NewTunnelingResponseDataBlock(communicationChannelId uint8, sequenceCounter uint8, status Status) *TunnelingResponseDataBlock {
	return &TunnelingResponseDataBlock{CommunicationChannelId: communicationChannelId, SequenceCounter: sequenceCounter, Status: status}
}

func CastTunnelingResponseDataBlock(structType interface{}) *TunnelingResponseDataBlock {
	castFunc := func(typ interface{}) *TunnelingResponseDataBlock {
		if casted, ok := typ.(TunnelingResponseDataBlock); ok {
			return &casted
		}
		if casted, ok := typ.(*TunnelingResponseDataBlock); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *TunnelingResponseDataBlock) GetTypeName() string {
	return "TunnelingResponseDataBlock"
}

func (m *TunnelingResponseDataBlock) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *TunnelingResponseDataBlock) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Implicit Field (structureLength)
	lengthInBits += 8

	// Simple field (communicationChannelId)
	lengthInBits += 8

	// Simple field (sequenceCounter)
	lengthInBits += 8

	// Simple field (status)
	lengthInBits += 8

	return lengthInBits
}

func (m *TunnelingResponseDataBlock) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func TunnelingResponseDataBlockParse(readBuffer utils.ReadBuffer) (*TunnelingResponseDataBlock, error) {
	if pullErr := readBuffer.PullContext("TunnelingResponseDataBlock"); pullErr != nil {
		return nil, pullErr
	}

	// Implicit Field (structureLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	structureLength, _structureLengthErr := readBuffer.ReadUint8("structureLength", 8)
	_ = structureLength
	if _structureLengthErr != nil {
		return nil, errors.Wrap(_structureLengthErr, "Error parsing 'structureLength' field")
	}

	// Simple Field (communicationChannelId)
	_communicationChannelId, _communicationChannelIdErr := readBuffer.ReadUint8("communicationChannelId", 8)
	if _communicationChannelIdErr != nil {
		return nil, errors.Wrap(_communicationChannelIdErr, "Error parsing 'communicationChannelId' field")
	}
	communicationChannelId := _communicationChannelId

	// Simple Field (sequenceCounter)
	_sequenceCounter, _sequenceCounterErr := readBuffer.ReadUint8("sequenceCounter", 8)
	if _sequenceCounterErr != nil {
		return nil, errors.Wrap(_sequenceCounterErr, "Error parsing 'sequenceCounter' field")
	}
	sequenceCounter := _sequenceCounter

	// Simple Field (status)
	if pullErr := readBuffer.PullContext("status"); pullErr != nil {
		return nil, pullErr
	}
	_status, _statusErr := StatusParse(readBuffer)
	if _statusErr != nil {
		return nil, errors.Wrap(_statusErr, "Error parsing 'status' field")
	}
	status := _status
	if closeErr := readBuffer.CloseContext("status"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("TunnelingResponseDataBlock"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewTunnelingResponseDataBlock(communicationChannelId, sequenceCounter, status), nil
}

func (m *TunnelingResponseDataBlock) Serialize(writeBuffer utils.WriteBuffer) error {
	if pushErr := writeBuffer.PushContext("TunnelingResponseDataBlock"); pushErr != nil {
		return pushErr
	}

	// Implicit Field (structureLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	structureLength := uint8(uint8(m.LengthInBytes()))
	_structureLengthErr := writeBuffer.WriteUint8("structureLength", 8, (structureLength))
	if _structureLengthErr != nil {
		return errors.Wrap(_structureLengthErr, "Error serializing 'structureLength' field")
	}

	// Simple Field (communicationChannelId)
	communicationChannelId := uint8(m.CommunicationChannelId)
	_communicationChannelIdErr := writeBuffer.WriteUint8("communicationChannelId", 8, (communicationChannelId))
	if _communicationChannelIdErr != nil {
		return errors.Wrap(_communicationChannelIdErr, "Error serializing 'communicationChannelId' field")
	}

	// Simple Field (sequenceCounter)
	sequenceCounter := uint8(m.SequenceCounter)
	_sequenceCounterErr := writeBuffer.WriteUint8("sequenceCounter", 8, (sequenceCounter))
	if _sequenceCounterErr != nil {
		return errors.Wrap(_sequenceCounterErr, "Error serializing 'sequenceCounter' field")
	}

	// Simple Field (status)
	if pushErr := writeBuffer.PushContext("status"); pushErr != nil {
		return pushErr
	}
	_statusErr := m.Status.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("status"); popErr != nil {
		return popErr
	}
	if _statusErr != nil {
		return errors.Wrap(_statusErr, "Error serializing 'status' field")
	}

	if popErr := writeBuffer.PopContext("TunnelingResponseDataBlock"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *TunnelingResponseDataBlock) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
