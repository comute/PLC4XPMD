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

// DeviceStatus is the data-structure of this message
type DeviceStatus struct {
	ProgramMode bool
}

// IDeviceStatus is the corresponding interface of DeviceStatus
type IDeviceStatus interface {
	// GetProgramMode returns ProgramMode (property field)
	GetProgramMode() bool
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

func (m *DeviceStatus) GetProgramMode() bool {
	return m.ProgramMode
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewDeviceStatus factory function for DeviceStatus
func NewDeviceStatus(programMode bool) *DeviceStatus {
	return &DeviceStatus{ProgramMode: programMode}
}

func CastDeviceStatus(structType interface{}) *DeviceStatus {
	if casted, ok := structType.(DeviceStatus); ok {
		return &casted
	}
	if casted, ok := structType.(*DeviceStatus); ok {
		return casted
	}
	return nil
}

func (m *DeviceStatus) GetTypeName() string {
	return "DeviceStatus"
}

func (m *DeviceStatus) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *DeviceStatus) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Reserved Field (reserved)
	lengthInBits += 7

	// Simple field (programMode)
	lengthInBits += 1

	return lengthInBits
}

func (m *DeviceStatus) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func DeviceStatusParse(readBuffer utils.ReadBuffer) (*DeviceStatus, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("DeviceStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DeviceStatus")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 7)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field")
		}
		if reserved != uint8(0x00) {
			log.Info().Fields(map[string]interface{}{
				"expected value": uint8(0x00),
				"got value":      reserved,
			}).Msg("Got unexpected response.")
		}
	}

	// Simple Field (programMode)
	_programMode, _programModeErr := readBuffer.ReadBit("programMode")
	if _programModeErr != nil {
		return nil, errors.Wrap(_programModeErr, "Error parsing 'programMode' field")
	}
	programMode := _programMode

	if closeErr := readBuffer.CloseContext("DeviceStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DeviceStatus")
	}

	// Create the instance
	return NewDeviceStatus(programMode), nil
}

func (m *DeviceStatus) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("DeviceStatus"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for DeviceStatus")
	}

	// Reserved Field (reserved)
	{
		_err := writeBuffer.WriteUint8("reserved", 7, uint8(0x00))
		if _err != nil {
			return errors.Wrap(_err, "Error serializing 'reserved' field")
		}
	}

	// Simple Field (programMode)
	programMode := bool(m.ProgramMode)
	_programModeErr := writeBuffer.WriteBit("programMode", (programMode))
	if _programModeErr != nil {
		return errors.Wrap(_programModeErr, "Error serializing 'programMode' field")
	}

	if popErr := writeBuffer.PopContext("DeviceStatus"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for DeviceStatus")
	}
	return nil
}

func (m *DeviceStatus) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
