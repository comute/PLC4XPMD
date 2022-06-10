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
	"github.com/rs/zerolog/log"
)

// Code generated by code-generation. DO NOT EDIT.

// StatusHeader is the data-structure of this message
type StatusHeader struct {
	NumberOfCharacterPairs uint8
}

// IStatusHeader is the corresponding interface of StatusHeader
type IStatusHeader interface {
	// GetNumberOfCharacterPairs returns NumberOfCharacterPairs (property field)
	GetNumberOfCharacterPairs() uint8
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *StatusHeader) GetNumberOfCharacterPairs() uint8 {
	return m.NumberOfCharacterPairs
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewStatusHeader factory function for StatusHeader
func NewStatusHeader(numberOfCharacterPairs uint8) *StatusHeader {
	return &StatusHeader{NumberOfCharacterPairs: numberOfCharacterPairs}
}

func CastStatusHeader(structType interface{}) *StatusHeader {
	if casted, ok := structType.(StatusHeader); ok {
		return &casted
	}
	if casted, ok := structType.(*StatusHeader); ok {
		return casted
	}
	return nil
}

func (m *StatusHeader) GetTypeName() string {
	return "StatusHeader"
}

func (m *StatusHeader) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *StatusHeader) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Reserved Field (reserved)
	lengthInBits += 2

	// Simple field (numberOfCharacterPairs)
	lengthInBits += 6

	return lengthInBits
}

func (m *StatusHeader) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func StatusHeaderParse(readBuffer utils.ReadBuffer) (*StatusHeader, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("StatusHeader"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for StatusHeader")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 2)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field")
		}
		if reserved != uint8(0x3) {
			log.Info().Fields(map[string]interface{}{
				"expected value": uint8(0x3),
				"got value":      reserved,
			}).Msg("Got unexpected response.")
		}
	}

	// Simple Field (numberOfCharacterPairs)
	_numberOfCharacterPairs, _numberOfCharacterPairsErr := readBuffer.ReadUint8("numberOfCharacterPairs", 6)
	if _numberOfCharacterPairsErr != nil {
		return nil, errors.Wrap(_numberOfCharacterPairsErr, "Error parsing 'numberOfCharacterPairs' field")
	}
	numberOfCharacterPairs := _numberOfCharacterPairs

	if closeErr := readBuffer.CloseContext("StatusHeader"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for StatusHeader")
	}

	// Create the instance
	return NewStatusHeader(numberOfCharacterPairs), nil
}

func (m *StatusHeader) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("StatusHeader"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for StatusHeader")
	}

	// Reserved Field (reserved)
	{
		_err := writeBuffer.WriteUint8("reserved", 2, uint8(0x3))
		if _err != nil {
			return errors.Wrap(_err, "Error serializing 'reserved' field")
		}
	}

	// Simple Field (numberOfCharacterPairs)
	numberOfCharacterPairs := uint8(m.NumberOfCharacterPairs)
	_numberOfCharacterPairsErr := writeBuffer.WriteUint8("numberOfCharacterPairs", 6, (numberOfCharacterPairs))
	if _numberOfCharacterPairsErr != nil {
		return errors.Wrap(_numberOfCharacterPairsErr, "Error serializing 'numberOfCharacterPairs' field")
	}

	if popErr := writeBuffer.PopContext("StatusHeader"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for StatusHeader")
	}
	return nil
}

func (m *StatusHeader) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
