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
	"github.com/rs/zerolog/log"
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetBDTEntry is the corresponding interface of BACnetBDTEntry
type BACnetBDTEntry interface {
	utils.LengthAware
	utils.Serializable
	// GetBbmdAddress returns BbmdAddress (property field)
	GetBbmdAddress() BACnetHostNPortEnclosed
	// GetBroadcastMask returns BroadcastMask (property field)
	GetBroadcastMask() BACnetContextTagOctetString
}

// BACnetBDTEntryExactly can be used when we want exactly this type and not a type which fulfills BACnetBDTEntry.
// This is useful for switch cases.
type BACnetBDTEntryExactly interface {
	BACnetBDTEntry
	isBACnetBDTEntry() bool
}

// _BACnetBDTEntry is the data-structure of this message
type _BACnetBDTEntry struct {
	BbmdAddress   BACnetHostNPortEnclosed
	BroadcastMask BACnetContextTagOctetString
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetBDTEntry) GetBbmdAddress() BACnetHostNPortEnclosed {
	return m.BbmdAddress
}

func (m *_BACnetBDTEntry) GetBroadcastMask() BACnetContextTagOctetString {
	return m.BroadcastMask
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetBDTEntry factory function for _BACnetBDTEntry
func NewBACnetBDTEntry(bbmdAddress BACnetHostNPortEnclosed, broadcastMask BACnetContextTagOctetString) *_BACnetBDTEntry {
	return &_BACnetBDTEntry{BbmdAddress: bbmdAddress, BroadcastMask: broadcastMask}
}

// Deprecated: use the interface for direct cast
func CastBACnetBDTEntry(structType interface{}) BACnetBDTEntry {
	if casted, ok := structType.(BACnetBDTEntry); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetBDTEntry); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetBDTEntry) GetTypeName() string {
	return "BACnetBDTEntry"
}

func (m *_BACnetBDTEntry) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetBDTEntry) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (bbmdAddress)
	lengthInBits += m.BbmdAddress.GetLengthInBits()

	// Optional Field (broadcastMask)
	if m.BroadcastMask != nil {
		lengthInBits += m.BroadcastMask.GetLengthInBits()
	}

	return lengthInBits
}

func (m *_BACnetBDTEntry) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetBDTEntryParse(readBuffer utils.ReadBuffer) (BACnetBDTEntry, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetBDTEntry"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetBDTEntry")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (bbmdAddress)
	if pullErr := readBuffer.PullContext("bbmdAddress"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for bbmdAddress")
	}
	_bbmdAddress, _bbmdAddressErr := BACnetHostNPortEnclosedParse(readBuffer, uint8(uint8(0)))
	if _bbmdAddressErr != nil {
		return nil, errors.Wrap(_bbmdAddressErr, "Error parsing 'bbmdAddress' field")
	}
	bbmdAddress := _bbmdAddress.(BACnetHostNPortEnclosed)
	if closeErr := readBuffer.CloseContext("bbmdAddress"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for bbmdAddress")
	}

	// Optional Field (broadcastMask) (Can be skipped, if a given expression evaluates to false)
	var broadcastMask BACnetContextTagOctetString = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("broadcastMask"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for broadcastMask")
		}
		_val, _err := BACnetContextTagParse(readBuffer, uint8(1), BACnetDataType_OCTET_STRING)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'broadcastMask' field")
		default:
			broadcastMask = _val.(BACnetContextTagOctetString)
			if closeErr := readBuffer.CloseContext("broadcastMask"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for broadcastMask")
			}
		}
	}

	if closeErr := readBuffer.CloseContext("BACnetBDTEntry"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetBDTEntry")
	}

	// Create the instance
	return NewBACnetBDTEntry(bbmdAddress, broadcastMask), nil
}

func (m *_BACnetBDTEntry) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetBDTEntry"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetBDTEntry")
	}

	// Simple Field (bbmdAddress)
	if pushErr := writeBuffer.PushContext("bbmdAddress"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for bbmdAddress")
	}
	_bbmdAddressErr := writeBuffer.WriteSerializable(m.GetBbmdAddress())
	if popErr := writeBuffer.PopContext("bbmdAddress"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for bbmdAddress")
	}
	if _bbmdAddressErr != nil {
		return errors.Wrap(_bbmdAddressErr, "Error serializing 'bbmdAddress' field")
	}

	// Optional Field (broadcastMask) (Can be skipped, if the value is null)
	var broadcastMask BACnetContextTagOctetString = nil
	if m.GetBroadcastMask() != nil {
		if pushErr := writeBuffer.PushContext("broadcastMask"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for broadcastMask")
		}
		broadcastMask = m.GetBroadcastMask()
		_broadcastMaskErr := writeBuffer.WriteSerializable(broadcastMask)
		if popErr := writeBuffer.PopContext("broadcastMask"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for broadcastMask")
		}
		if _broadcastMaskErr != nil {
			return errors.Wrap(_broadcastMaskErr, "Error serializing 'broadcastMask' field")
		}
	}

	if popErr := writeBuffer.PopContext("BACnetBDTEntry"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetBDTEntry")
	}
	return nil
}

func (m *_BACnetBDTEntry) isBACnetBDTEntry() bool {
	return true
}

func (m *_BACnetBDTEntry) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
