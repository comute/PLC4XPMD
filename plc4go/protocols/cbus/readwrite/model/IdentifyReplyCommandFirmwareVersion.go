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

// IdentifyReplyCommandFirmwareVersion is the corresponding interface of IdentifyReplyCommandFirmwareVersion
type IdentifyReplyCommandFirmwareVersion interface {
	utils.LengthAware
	utils.Serializable
	IdentifyReplyCommand
	// GetFirmwareVersion returns FirmwareVersion (property field)
	GetFirmwareVersion() string
}

// IdentifyReplyCommandFirmwareVersionExactly can be used when we want exactly this type and not a type which fulfills IdentifyReplyCommandFirmwareVersion.
// This is useful for switch cases.
type IdentifyReplyCommandFirmwareVersionExactly interface {
	IdentifyReplyCommandFirmwareVersion
	isIdentifyReplyCommandFirmwareVersion() bool
}

// _IdentifyReplyCommandFirmwareVersion is the data-structure of this message
type _IdentifyReplyCommandFirmwareVersion struct {
	*_IdentifyReplyCommand
	FirmwareVersion string
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_IdentifyReplyCommandFirmwareVersion) GetAttribute() Attribute {
	return Attribute_FirmwareVersion
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_IdentifyReplyCommandFirmwareVersion) InitializeParent(parent IdentifyReplyCommand) {}

func (m *_IdentifyReplyCommandFirmwareVersion) GetParent() IdentifyReplyCommand {
	return m._IdentifyReplyCommand
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_IdentifyReplyCommandFirmwareVersion) GetFirmwareVersion() string {
	return m.FirmwareVersion
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewIdentifyReplyCommandFirmwareVersion factory function for _IdentifyReplyCommandFirmwareVersion
func NewIdentifyReplyCommandFirmwareVersion(firmwareVersion string) *_IdentifyReplyCommandFirmwareVersion {
	_result := &_IdentifyReplyCommandFirmwareVersion{
		FirmwareVersion:       firmwareVersion,
		_IdentifyReplyCommand: NewIdentifyReplyCommand(),
	}
	_result._IdentifyReplyCommand._IdentifyReplyCommandChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastIdentifyReplyCommandFirmwareVersion(structType interface{}) IdentifyReplyCommandFirmwareVersion {
	if casted, ok := structType.(IdentifyReplyCommandFirmwareVersion); ok {
		return casted
	}
	if casted, ok := structType.(*IdentifyReplyCommandFirmwareVersion); ok {
		return *casted
	}
	return nil
}

func (m *_IdentifyReplyCommandFirmwareVersion) GetTypeName() string {
	return "IdentifyReplyCommandFirmwareVersion"
}

func (m *_IdentifyReplyCommandFirmwareVersion) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_IdentifyReplyCommandFirmwareVersion) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (firmwareVersion)
	lengthInBits += 64

	return lengthInBits
}

func (m *_IdentifyReplyCommandFirmwareVersion) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func IdentifyReplyCommandFirmwareVersionParse(readBuffer utils.ReadBuffer, attribute Attribute) (IdentifyReplyCommandFirmwareVersion, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("IdentifyReplyCommandFirmwareVersion"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for IdentifyReplyCommandFirmwareVersion")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (firmwareVersion)
	_firmwareVersion, _firmwareVersionErr := readBuffer.ReadString("firmwareVersion", uint32(64))
	if _firmwareVersionErr != nil {
		return nil, errors.Wrap(_firmwareVersionErr, "Error parsing 'firmwareVersion' field")
	}
	firmwareVersion := _firmwareVersion

	if closeErr := readBuffer.CloseContext("IdentifyReplyCommandFirmwareVersion"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for IdentifyReplyCommandFirmwareVersion")
	}

	// Create a partially initialized instance
	_child := &_IdentifyReplyCommandFirmwareVersion{
		FirmwareVersion:       firmwareVersion,
		_IdentifyReplyCommand: &_IdentifyReplyCommand{},
	}
	_child._IdentifyReplyCommand._IdentifyReplyCommandChildRequirements = _child
	return _child, nil
}

func (m *_IdentifyReplyCommandFirmwareVersion) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("IdentifyReplyCommandFirmwareVersion"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for IdentifyReplyCommandFirmwareVersion")
		}

		// Simple Field (firmwareVersion)
		firmwareVersion := string(m.GetFirmwareVersion())
		_firmwareVersionErr := writeBuffer.WriteString("firmwareVersion", uint32(64), "UTF-8", (firmwareVersion))
		if _firmwareVersionErr != nil {
			return errors.Wrap(_firmwareVersionErr, "Error serializing 'firmwareVersion' field")
		}

		if popErr := writeBuffer.PopContext("IdentifyReplyCommandFirmwareVersion"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for IdentifyReplyCommandFirmwareVersion")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_IdentifyReplyCommandFirmwareVersion) isIdentifyReplyCommandFirmwareVersion() bool {
	return true
}

func (m *_IdentifyReplyCommandFirmwareVersion) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
