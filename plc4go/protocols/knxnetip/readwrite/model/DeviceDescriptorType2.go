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

// DeviceDescriptorType2 is the corresponding interface of DeviceDescriptorType2
type DeviceDescriptorType2 interface {
	utils.LengthAware
	utils.Serializable
	// GetManufacturerId returns ManufacturerId (property field)
	GetManufacturerId() uint16
	// GetDeviceType returns DeviceType (property field)
	GetDeviceType() uint16
	// GetVersion returns Version (property field)
	GetVersion() uint8
	// GetReadSupported returns ReadSupported (property field)
	GetReadSupported() bool
	// GetWriteSupported returns WriteSupported (property field)
	GetWriteSupported() bool
	// GetLogicalTagBase returns LogicalTagBase (property field)
	GetLogicalTagBase() uint8
	// GetChannelInfo1 returns ChannelInfo1 (property field)
	GetChannelInfo1() ChannelInformation
	// GetChannelInfo2 returns ChannelInfo2 (property field)
	GetChannelInfo2() ChannelInformation
	// GetChannelInfo3 returns ChannelInfo3 (property field)
	GetChannelInfo3() ChannelInformation
	// GetChannelInfo4 returns ChannelInfo4 (property field)
	GetChannelInfo4() ChannelInformation
}

// DeviceDescriptorType2Exactly can be used when we want exactly this type and not a type which fulfills DeviceDescriptorType2.
// This is useful for switch cases.
type DeviceDescriptorType2Exactly interface {
	DeviceDescriptorType2
	isDeviceDescriptorType2() bool
}

// _DeviceDescriptorType2 is the data-structure of this message
type _DeviceDescriptorType2 struct {
	ManufacturerId uint16
	DeviceType     uint16
	Version        uint8
	ReadSupported  bool
	WriteSupported bool
	LogicalTagBase uint8
	ChannelInfo1   ChannelInformation
	ChannelInfo2   ChannelInformation
	ChannelInfo3   ChannelInformation
	ChannelInfo4   ChannelInformation
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_DeviceDescriptorType2) GetManufacturerId() uint16 {
	return m.ManufacturerId
}

func (m *_DeviceDescriptorType2) GetDeviceType() uint16 {
	return m.DeviceType
}

func (m *_DeviceDescriptorType2) GetVersion() uint8 {
	return m.Version
}

func (m *_DeviceDescriptorType2) GetReadSupported() bool {
	return m.ReadSupported
}

func (m *_DeviceDescriptorType2) GetWriteSupported() bool {
	return m.WriteSupported
}

func (m *_DeviceDescriptorType2) GetLogicalTagBase() uint8 {
	return m.LogicalTagBase
}

func (m *_DeviceDescriptorType2) GetChannelInfo1() ChannelInformation {
	return m.ChannelInfo1
}

func (m *_DeviceDescriptorType2) GetChannelInfo2() ChannelInformation {
	return m.ChannelInfo2
}

func (m *_DeviceDescriptorType2) GetChannelInfo3() ChannelInformation {
	return m.ChannelInfo3
}

func (m *_DeviceDescriptorType2) GetChannelInfo4() ChannelInformation {
	return m.ChannelInfo4
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewDeviceDescriptorType2 factory function for _DeviceDescriptorType2
func NewDeviceDescriptorType2(manufacturerId uint16, deviceType uint16, version uint8, readSupported bool, writeSupported bool, logicalTagBase uint8, channelInfo1 ChannelInformation, channelInfo2 ChannelInformation, channelInfo3 ChannelInformation, channelInfo4 ChannelInformation) *_DeviceDescriptorType2 {
	return &_DeviceDescriptorType2{ManufacturerId: manufacturerId, DeviceType: deviceType, Version: version, ReadSupported: readSupported, WriteSupported: writeSupported, LogicalTagBase: logicalTagBase, ChannelInfo1: channelInfo1, ChannelInfo2: channelInfo2, ChannelInfo3: channelInfo3, ChannelInfo4: channelInfo4}
}

// Deprecated: use the interface for direct cast
func CastDeviceDescriptorType2(structType interface{}) DeviceDescriptorType2 {
	if casted, ok := structType.(DeviceDescriptorType2); ok {
		return casted
	}
	if casted, ok := structType.(*DeviceDescriptorType2); ok {
		return *casted
	}
	return nil
}

func (m *_DeviceDescriptorType2) GetTypeName() string {
	return "DeviceDescriptorType2"
}

func (m *_DeviceDescriptorType2) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_DeviceDescriptorType2) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (manufacturerId)
	lengthInBits += 16

	// Simple field (deviceType)
	lengthInBits += 16

	// Simple field (version)
	lengthInBits += 8

	// Simple field (readSupported)
	lengthInBits += 1

	// Simple field (writeSupported)
	lengthInBits += 1

	// Simple field (logicalTagBase)
	lengthInBits += 6

	// Simple field (channelInfo1)
	lengthInBits += m.ChannelInfo1.GetLengthInBits()

	// Simple field (channelInfo2)
	lengthInBits += m.ChannelInfo2.GetLengthInBits()

	// Simple field (channelInfo3)
	lengthInBits += m.ChannelInfo3.GetLengthInBits()

	// Simple field (channelInfo4)
	lengthInBits += m.ChannelInfo4.GetLengthInBits()

	return lengthInBits
}

func (m *_DeviceDescriptorType2) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func DeviceDescriptorType2Parse(readBuffer utils.ReadBuffer) (DeviceDescriptorType2, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("DeviceDescriptorType2"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DeviceDescriptorType2")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (manufacturerId)
	_manufacturerId, _manufacturerIdErr := readBuffer.ReadUint16("manufacturerId", 16)
	if _manufacturerIdErr != nil {
		return nil, errors.Wrap(_manufacturerIdErr, "Error parsing 'manufacturerId' field")
	}
	manufacturerId := _manufacturerId

	// Simple Field (deviceType)
	_deviceType, _deviceTypeErr := readBuffer.ReadUint16("deviceType", 16)
	if _deviceTypeErr != nil {
		return nil, errors.Wrap(_deviceTypeErr, "Error parsing 'deviceType' field")
	}
	deviceType := _deviceType

	// Simple Field (version)
	_version, _versionErr := readBuffer.ReadUint8("version", 8)
	if _versionErr != nil {
		return nil, errors.Wrap(_versionErr, "Error parsing 'version' field")
	}
	version := _version

	// Simple Field (readSupported)
	_readSupported, _readSupportedErr := readBuffer.ReadBit("readSupported")
	if _readSupportedErr != nil {
		return nil, errors.Wrap(_readSupportedErr, "Error parsing 'readSupported' field")
	}
	readSupported := _readSupported

	// Simple Field (writeSupported)
	_writeSupported, _writeSupportedErr := readBuffer.ReadBit("writeSupported")
	if _writeSupportedErr != nil {
		return nil, errors.Wrap(_writeSupportedErr, "Error parsing 'writeSupported' field")
	}
	writeSupported := _writeSupported

	// Simple Field (logicalTagBase)
	_logicalTagBase, _logicalTagBaseErr := readBuffer.ReadUint8("logicalTagBase", 6)
	if _logicalTagBaseErr != nil {
		return nil, errors.Wrap(_logicalTagBaseErr, "Error parsing 'logicalTagBase' field")
	}
	logicalTagBase := _logicalTagBase

	// Simple Field (channelInfo1)
	if pullErr := readBuffer.PullContext("channelInfo1"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for channelInfo1")
	}
	_channelInfo1, _channelInfo1Err := ChannelInformationParse(readBuffer)
	if _channelInfo1Err != nil {
		return nil, errors.Wrap(_channelInfo1Err, "Error parsing 'channelInfo1' field")
	}
	channelInfo1 := _channelInfo1.(ChannelInformation)
	if closeErr := readBuffer.CloseContext("channelInfo1"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for channelInfo1")
	}

	// Simple Field (channelInfo2)
	if pullErr := readBuffer.PullContext("channelInfo2"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for channelInfo2")
	}
	_channelInfo2, _channelInfo2Err := ChannelInformationParse(readBuffer)
	if _channelInfo2Err != nil {
		return nil, errors.Wrap(_channelInfo2Err, "Error parsing 'channelInfo2' field")
	}
	channelInfo2 := _channelInfo2.(ChannelInformation)
	if closeErr := readBuffer.CloseContext("channelInfo2"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for channelInfo2")
	}

	// Simple Field (channelInfo3)
	if pullErr := readBuffer.PullContext("channelInfo3"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for channelInfo3")
	}
	_channelInfo3, _channelInfo3Err := ChannelInformationParse(readBuffer)
	if _channelInfo3Err != nil {
		return nil, errors.Wrap(_channelInfo3Err, "Error parsing 'channelInfo3' field")
	}
	channelInfo3 := _channelInfo3.(ChannelInformation)
	if closeErr := readBuffer.CloseContext("channelInfo3"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for channelInfo3")
	}

	// Simple Field (channelInfo4)
	if pullErr := readBuffer.PullContext("channelInfo4"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for channelInfo4")
	}
	_channelInfo4, _channelInfo4Err := ChannelInformationParse(readBuffer)
	if _channelInfo4Err != nil {
		return nil, errors.Wrap(_channelInfo4Err, "Error parsing 'channelInfo4' field")
	}
	channelInfo4 := _channelInfo4.(ChannelInformation)
	if closeErr := readBuffer.CloseContext("channelInfo4"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for channelInfo4")
	}

	if closeErr := readBuffer.CloseContext("DeviceDescriptorType2"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DeviceDescriptorType2")
	}

	// Create the instance
	return NewDeviceDescriptorType2(manufacturerId, deviceType, version, readSupported, writeSupported, logicalTagBase, channelInfo1, channelInfo2, channelInfo3, channelInfo4), nil
}

func (m *_DeviceDescriptorType2) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("DeviceDescriptorType2"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for DeviceDescriptorType2")
	}

	// Simple Field (manufacturerId)
	manufacturerId := uint16(m.GetManufacturerId())
	_manufacturerIdErr := writeBuffer.WriteUint16("manufacturerId", 16, (manufacturerId))
	if _manufacturerIdErr != nil {
		return errors.Wrap(_manufacturerIdErr, "Error serializing 'manufacturerId' field")
	}

	// Simple Field (deviceType)
	deviceType := uint16(m.GetDeviceType())
	_deviceTypeErr := writeBuffer.WriteUint16("deviceType", 16, (deviceType))
	if _deviceTypeErr != nil {
		return errors.Wrap(_deviceTypeErr, "Error serializing 'deviceType' field")
	}

	// Simple Field (version)
	version := uint8(m.GetVersion())
	_versionErr := writeBuffer.WriteUint8("version", 8, (version))
	if _versionErr != nil {
		return errors.Wrap(_versionErr, "Error serializing 'version' field")
	}

	// Simple Field (readSupported)
	readSupported := bool(m.GetReadSupported())
	_readSupportedErr := writeBuffer.WriteBit("readSupported", (readSupported))
	if _readSupportedErr != nil {
		return errors.Wrap(_readSupportedErr, "Error serializing 'readSupported' field")
	}

	// Simple Field (writeSupported)
	writeSupported := bool(m.GetWriteSupported())
	_writeSupportedErr := writeBuffer.WriteBit("writeSupported", (writeSupported))
	if _writeSupportedErr != nil {
		return errors.Wrap(_writeSupportedErr, "Error serializing 'writeSupported' field")
	}

	// Simple Field (logicalTagBase)
	logicalTagBase := uint8(m.GetLogicalTagBase())
	_logicalTagBaseErr := writeBuffer.WriteUint8("logicalTagBase", 6, (logicalTagBase))
	if _logicalTagBaseErr != nil {
		return errors.Wrap(_logicalTagBaseErr, "Error serializing 'logicalTagBase' field")
	}

	// Simple Field (channelInfo1)
	if pushErr := writeBuffer.PushContext("channelInfo1"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for channelInfo1")
	}
	_channelInfo1Err := writeBuffer.WriteSerializable(m.GetChannelInfo1())
	if popErr := writeBuffer.PopContext("channelInfo1"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for channelInfo1")
	}
	if _channelInfo1Err != nil {
		return errors.Wrap(_channelInfo1Err, "Error serializing 'channelInfo1' field")
	}

	// Simple Field (channelInfo2)
	if pushErr := writeBuffer.PushContext("channelInfo2"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for channelInfo2")
	}
	_channelInfo2Err := writeBuffer.WriteSerializable(m.GetChannelInfo2())
	if popErr := writeBuffer.PopContext("channelInfo2"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for channelInfo2")
	}
	if _channelInfo2Err != nil {
		return errors.Wrap(_channelInfo2Err, "Error serializing 'channelInfo2' field")
	}

	// Simple Field (channelInfo3)
	if pushErr := writeBuffer.PushContext("channelInfo3"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for channelInfo3")
	}
	_channelInfo3Err := writeBuffer.WriteSerializable(m.GetChannelInfo3())
	if popErr := writeBuffer.PopContext("channelInfo3"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for channelInfo3")
	}
	if _channelInfo3Err != nil {
		return errors.Wrap(_channelInfo3Err, "Error serializing 'channelInfo3' field")
	}

	// Simple Field (channelInfo4)
	if pushErr := writeBuffer.PushContext("channelInfo4"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for channelInfo4")
	}
	_channelInfo4Err := writeBuffer.WriteSerializable(m.GetChannelInfo4())
	if popErr := writeBuffer.PopContext("channelInfo4"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for channelInfo4")
	}
	if _channelInfo4Err != nil {
		return errors.Wrap(_channelInfo4Err, "Error serializing 'channelInfo4' field")
	}

	if popErr := writeBuffer.PopContext("DeviceDescriptorType2"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for DeviceDescriptorType2")
	}
	return nil
}

func (m *_DeviceDescriptorType2) isDeviceDescriptorType2() bool {
	return true
}

func (m *_DeviceDescriptorType2) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
