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

// AdsData is the data-structure of this message
type AdsData struct {
	Child IAdsDataChild
}

// IAdsData is the corresponding interface of AdsData
type IAdsData interface {
	// GetCommandId returns CommandId (discriminator field)
	GetCommandId() CommandId
	// GetResponse returns Response (discriminator field)
	GetResponse() bool
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

type IAdsDataParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child IAdsData, serializeChildFunction func() error) error
	GetTypeName() string
}

type IAdsDataChild interface {
	Serialize(writeBuffer utils.WriteBuffer) error
	InitializeParent(parent *AdsData)
	GetParent() *AdsData

	GetTypeName() string
	IAdsData
}

// NewAdsData factory function for AdsData
func NewAdsData() *AdsData {
	return &AdsData{}
}

func CastAdsData(structType interface{}) *AdsData {
	if casted, ok := structType.(AdsData); ok {
		return &casted
	}
	if casted, ok := structType.(*AdsData); ok {
		return casted
	}
	if casted, ok := structType.(IAdsDataChild); ok {
		return casted.GetParent()
	}
	return nil
}

func (m *AdsData) GetTypeName() string {
	return "AdsData"
}

func (m *AdsData) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *AdsData) GetLengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.GetLengthInBits()
}

func (m *AdsData) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)

	return lengthInBits
}

func (m *AdsData) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func AdsDataParse(readBuffer utils.ReadBuffer, commandId CommandId, response bool) (*AdsData, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AdsData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AdsData")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type AdsDataChild interface {
		InitializeParent(*AdsData)
		GetParent() *AdsData
	}
	var _child AdsDataChild
	var typeSwitchError error
	switch {
	case commandId == CommandId_INVALID && response == bool(false): // AdsInvalidRequest
		_child, typeSwitchError = AdsInvalidRequestParse(readBuffer, commandId, response)
	case commandId == CommandId_INVALID && response == bool(true): // AdsInvalidResponse
		_child, typeSwitchError = AdsInvalidResponseParse(readBuffer, commandId, response)
	case commandId == CommandId_ADS_READ_DEVICE_INFO && response == bool(false): // AdsReadDeviceInfoRequest
		_child, typeSwitchError = AdsReadDeviceInfoRequestParse(readBuffer, commandId, response)
	case commandId == CommandId_ADS_READ_DEVICE_INFO && response == bool(true): // AdsReadDeviceInfoResponse
		_child, typeSwitchError = AdsReadDeviceInfoResponseParse(readBuffer, commandId, response)
	case commandId == CommandId_ADS_READ && response == bool(false): // AdsReadRequest
		_child, typeSwitchError = AdsReadRequestParse(readBuffer, commandId, response)
	case commandId == CommandId_ADS_READ && response == bool(true): // AdsReadResponse
		_child, typeSwitchError = AdsReadResponseParse(readBuffer, commandId, response)
	case commandId == CommandId_ADS_WRITE && response == bool(false): // AdsWriteRequest
		_child, typeSwitchError = AdsWriteRequestParse(readBuffer, commandId, response)
	case commandId == CommandId_ADS_WRITE && response == bool(true): // AdsWriteResponse
		_child, typeSwitchError = AdsWriteResponseParse(readBuffer, commandId, response)
	case commandId == CommandId_ADS_READ_STATE && response == bool(false): // AdsReadStateRequest
		_child, typeSwitchError = AdsReadStateRequestParse(readBuffer, commandId, response)
	case commandId == CommandId_ADS_READ_STATE && response == bool(true): // AdsReadStateResponse
		_child, typeSwitchError = AdsReadStateResponseParse(readBuffer, commandId, response)
	case commandId == CommandId_ADS_WRITE_CONTROL && response == bool(false): // AdsWriteControlRequest
		_child, typeSwitchError = AdsWriteControlRequestParse(readBuffer, commandId, response)
	case commandId == CommandId_ADS_WRITE_CONTROL && response == bool(true): // AdsWriteControlResponse
		_child, typeSwitchError = AdsWriteControlResponseParse(readBuffer, commandId, response)
	case commandId == CommandId_ADS_ADD_DEVICE_NOTIFICATION && response == bool(false): // AdsAddDeviceNotificationRequest
		_child, typeSwitchError = AdsAddDeviceNotificationRequestParse(readBuffer, commandId, response)
	case commandId == CommandId_ADS_ADD_DEVICE_NOTIFICATION && response == bool(true): // AdsAddDeviceNotificationResponse
		_child, typeSwitchError = AdsAddDeviceNotificationResponseParse(readBuffer, commandId, response)
	case commandId == CommandId_ADS_DELETE_DEVICE_NOTIFICATION && response == bool(false): // AdsDeleteDeviceNotificationRequest
		_child, typeSwitchError = AdsDeleteDeviceNotificationRequestParse(readBuffer, commandId, response)
	case commandId == CommandId_ADS_DELETE_DEVICE_NOTIFICATION && response == bool(true): // AdsDeleteDeviceNotificationResponse
		_child, typeSwitchError = AdsDeleteDeviceNotificationResponseParse(readBuffer, commandId, response)
	case commandId == CommandId_ADS_DEVICE_NOTIFICATION && response == bool(false): // AdsDeviceNotificationRequest
		_child, typeSwitchError = AdsDeviceNotificationRequestParse(readBuffer, commandId, response)
	case commandId == CommandId_ADS_DEVICE_NOTIFICATION && response == bool(true): // AdsDeviceNotificationResponse
		_child, typeSwitchError = AdsDeviceNotificationResponseParse(readBuffer, commandId, response)
	case commandId == CommandId_ADS_READ_WRITE && response == bool(false): // AdsReadWriteRequest
		_child, typeSwitchError = AdsReadWriteRequestParse(readBuffer, commandId, response)
	case commandId == CommandId_ADS_READ_WRITE && response == bool(true): // AdsReadWriteResponse
		_child, typeSwitchError = AdsReadWriteResponseParse(readBuffer, commandId, response)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	if closeErr := readBuffer.CloseContext("AdsData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AdsData")
	}

	// Finish initializing
	_child.InitializeParent(_child.GetParent())
	return _child.GetParent(), nil
}

func (m *AdsData) Serialize(writeBuffer utils.WriteBuffer) error {
	return m.Child.Serialize(writeBuffer)
}

func (m *AdsData) SerializeParent(writeBuffer utils.WriteBuffer, child IAdsData, serializeChildFunction func() error) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("AdsData"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for AdsData")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("AdsData"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for AdsData")
	}
	return nil
}

func (m *AdsData) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
