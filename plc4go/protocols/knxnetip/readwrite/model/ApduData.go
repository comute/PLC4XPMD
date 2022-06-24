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

// ApduData is the corresponding interface of ApduData
type ApduData interface {
	utils.LengthAware
	utils.Serializable
	// GetApciType returns ApciType (discriminator field)
	GetApciType() uint8
}

// ApduDataExactly can be used when we want exactly this type and not a type which fulfills ApduData.
// This is useful for switch cases.
type ApduDataExactly interface {
	ApduData
	isApduData() bool
}

// _ApduData is the data-structure of this message
type _ApduData struct {
	_ApduDataChildRequirements

	// Arguments.
	DataLength uint8
}

type _ApduDataChildRequirements interface {
	utils.Serializable
	GetLengthInBits() uint16
	GetLengthInBitsConditional(lastItem bool) uint16
	GetApciType() uint8
}

type ApduDataParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child ApduData, serializeChildFunction func() error) error
	GetTypeName() string
}

type ApduDataChild interface {
	utils.Serializable
	InitializeParent(parent ApduData)
	GetParent() *ApduData

	GetTypeName() string
	ApduData
}

// NewApduData factory function for _ApduData
func NewApduData(dataLength uint8) *_ApduData {
	return &_ApduData{DataLength: dataLength}
}

// Deprecated: use the interface for direct cast
func CastApduData(structType interface{}) ApduData {
	if casted, ok := structType.(ApduData); ok {
		return casted
	}
	if casted, ok := structType.(*ApduData); ok {
		return *casted
	}
	return nil
}

func (m *_ApduData) GetTypeName() string {
	return "ApduData"
}

func (m *_ApduData) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (apciType)
	lengthInBits += 4

	return lengthInBits
}

func (m *_ApduData) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ApduDataParse(readBuffer utils.ReadBuffer, dataLength uint8) (ApduData, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduData")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Discriminator Field (apciType) (Used as input to a switch field)
	apciType, _apciTypeErr := readBuffer.ReadUint8("apciType", 4)
	if _apciTypeErr != nil {
		return nil, errors.Wrap(_apciTypeErr, "Error parsing 'apciType' field")
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type ApduDataChildSerializeRequirement interface {
		ApduData
		InitializeParent(ApduData)
		GetParent() ApduData
	}
	var _childTemp interface{}
	var _child ApduDataChildSerializeRequirement
	var typeSwitchError error
	switch {
	case apciType == 0x0: // ApduDataGroupValueRead
		_childTemp, typeSwitchError = ApduDataGroupValueReadParse(readBuffer, dataLength)
	case apciType == 0x1: // ApduDataGroupValueResponse
		_childTemp, typeSwitchError = ApduDataGroupValueResponseParse(readBuffer, dataLength)
	case apciType == 0x2: // ApduDataGroupValueWrite
		_childTemp, typeSwitchError = ApduDataGroupValueWriteParse(readBuffer, dataLength)
	case apciType == 0x3: // ApduDataIndividualAddressWrite
		_childTemp, typeSwitchError = ApduDataIndividualAddressWriteParse(readBuffer, dataLength)
	case apciType == 0x4: // ApduDataIndividualAddressRead
		_childTemp, typeSwitchError = ApduDataIndividualAddressReadParse(readBuffer, dataLength)
	case apciType == 0x5: // ApduDataIndividualAddressResponse
		_childTemp, typeSwitchError = ApduDataIndividualAddressResponseParse(readBuffer, dataLength)
	case apciType == 0x6: // ApduDataAdcRead
		_childTemp, typeSwitchError = ApduDataAdcReadParse(readBuffer, dataLength)
	case apciType == 0x7: // ApduDataAdcResponse
		_childTemp, typeSwitchError = ApduDataAdcResponseParse(readBuffer, dataLength)
	case apciType == 0x8: // ApduDataMemoryRead
		_childTemp, typeSwitchError = ApduDataMemoryReadParse(readBuffer, dataLength)
	case apciType == 0x9: // ApduDataMemoryResponse
		_childTemp, typeSwitchError = ApduDataMemoryResponseParse(readBuffer, dataLength)
	case apciType == 0xA: // ApduDataMemoryWrite
		_childTemp, typeSwitchError = ApduDataMemoryWriteParse(readBuffer, dataLength)
	case apciType == 0xB: // ApduDataUserMessage
		_childTemp, typeSwitchError = ApduDataUserMessageParse(readBuffer, dataLength)
	case apciType == 0xC: // ApduDataDeviceDescriptorRead
		_childTemp, typeSwitchError = ApduDataDeviceDescriptorReadParse(readBuffer, dataLength)
	case apciType == 0xD: // ApduDataDeviceDescriptorResponse
		_childTemp, typeSwitchError = ApduDataDeviceDescriptorResponseParse(readBuffer, dataLength)
	case apciType == 0xE: // ApduDataRestart
		_childTemp, typeSwitchError = ApduDataRestartParse(readBuffer, dataLength)
	case apciType == 0xF: // ApduDataOther
		_childTemp, typeSwitchError = ApduDataOtherParse(readBuffer, dataLength)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}
	_child = _childTemp.(ApduDataChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("ApduData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduData")
	}

	// Finish initializing
	_child.InitializeParent(_child)
	return _child, nil
}

func (pm *_ApduData) SerializeParent(writeBuffer utils.WriteBuffer, child ApduData, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("ApduData"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ApduData")
	}

	// Discriminator Field (apciType) (Used as input to a switch field)
	apciType := uint8(child.GetApciType())
	_apciTypeErr := writeBuffer.WriteUint8("apciType", 4, (apciType))

	if _apciTypeErr != nil {
		return errors.Wrap(_apciTypeErr, "Error serializing 'apciType' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("ApduData"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ApduData")
	}
	return nil
}

func (m *_ApduData) isApduData() bool {
	return true
}

func (m *_ApduData) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
