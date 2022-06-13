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
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConfirmedServiceRequestReinitializeDevice is the data-structure of this message
type BACnetConfirmedServiceRequestReinitializeDevice struct {
	*BACnetConfirmedServiceRequest
	ReinitializedStateOfDevice *BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged
	Password                   *BACnetContextTagCharacterString

	// Arguments.
	ServiceRequestLength uint16
}

// IBACnetConfirmedServiceRequestReinitializeDevice is the corresponding interface of BACnetConfirmedServiceRequestReinitializeDevice
type IBACnetConfirmedServiceRequestReinitializeDevice interface {
	IBACnetConfirmedServiceRequest
	// GetReinitializedStateOfDevice returns ReinitializedStateOfDevice (property field)
	GetReinitializedStateOfDevice() *BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged
	// GetPassword returns Password (property field)
	GetPassword() *BACnetContextTagCharacterString
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

func (m *BACnetConfirmedServiceRequestReinitializeDevice) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_REINITIALIZE_DEVICE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConfirmedServiceRequestReinitializeDevice) InitializeParent(parent *BACnetConfirmedServiceRequest) {
}

func (m *BACnetConfirmedServiceRequestReinitializeDevice) GetParent() *BACnetConfirmedServiceRequest {
	return m.BACnetConfirmedServiceRequest
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConfirmedServiceRequestReinitializeDevice) GetReinitializedStateOfDevice() *BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged {
	return m.ReinitializedStateOfDevice
}

func (m *BACnetConfirmedServiceRequestReinitializeDevice) GetPassword() *BACnetContextTagCharacterString {
	return m.Password
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConfirmedServiceRequestReinitializeDevice factory function for BACnetConfirmedServiceRequestReinitializeDevice
func NewBACnetConfirmedServiceRequestReinitializeDevice(reinitializedStateOfDevice *BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged, password *BACnetContextTagCharacterString, serviceRequestLength uint16) *BACnetConfirmedServiceRequestReinitializeDevice {
	_result := &BACnetConfirmedServiceRequestReinitializeDevice{
		ReinitializedStateOfDevice:    reinitializedStateOfDevice,
		Password:                      password,
		BACnetConfirmedServiceRequest: NewBACnetConfirmedServiceRequest(serviceRequestLength),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConfirmedServiceRequestReinitializeDevice(structType interface{}) *BACnetConfirmedServiceRequestReinitializeDevice {
	if casted, ok := structType.(BACnetConfirmedServiceRequestReinitializeDevice); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestReinitializeDevice); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConfirmedServiceRequest); ok {
		return CastBACnetConfirmedServiceRequestReinitializeDevice(casted.Child)
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequest); ok {
		return CastBACnetConfirmedServiceRequestReinitializeDevice(casted.Child)
	}
	return nil
}

func (m *BACnetConfirmedServiceRequestReinitializeDevice) GetTypeName() string {
	return "BACnetConfirmedServiceRequestReinitializeDevice"
}

func (m *BACnetConfirmedServiceRequestReinitializeDevice) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConfirmedServiceRequestReinitializeDevice) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (reinitializedStateOfDevice)
	lengthInBits += m.ReinitializedStateOfDevice.GetLengthInBits()

	// Optional Field (password)
	if m.Password != nil {
		lengthInBits += (*m.Password).GetLengthInBits()
	}

	return lengthInBits
}

func (m *BACnetConfirmedServiceRequestReinitializeDevice) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConfirmedServiceRequestReinitializeDeviceParse(readBuffer utils.ReadBuffer, serviceRequestLength uint16) (*BACnetConfirmedServiceRequestReinitializeDevice, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestReinitializeDevice"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestReinitializeDevice")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (reinitializedStateOfDevice)
	if pullErr := readBuffer.PullContext("reinitializedStateOfDevice"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for reinitializedStateOfDevice")
	}
	_reinitializedStateOfDevice, _reinitializedStateOfDeviceErr := BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTaggedParse(readBuffer, uint8(uint8(0)), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _reinitializedStateOfDeviceErr != nil {
		return nil, errors.Wrap(_reinitializedStateOfDeviceErr, "Error parsing 'reinitializedStateOfDevice' field")
	}
	reinitializedStateOfDevice := CastBACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged(_reinitializedStateOfDevice)
	if closeErr := readBuffer.CloseContext("reinitializedStateOfDevice"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for reinitializedStateOfDevice")
	}

	// Optional Field (password) (Can be skipped, if a given expression evaluates to false)
	var password *BACnetContextTagCharacterString = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("password"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for password")
		}
		_val, _err := BACnetContextTagParse(readBuffer, uint8(1), BACnetDataType_CHARACTER_STRING)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'password' field")
		default:
			password = CastBACnetContextTagCharacterString(_val)
			if closeErr := readBuffer.CloseContext("password"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for password")
			}
		}
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestReinitializeDevice"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestReinitializeDevice")
	}

	// Create a partially initialized instance
	_child := &BACnetConfirmedServiceRequestReinitializeDevice{
		ReinitializedStateOfDevice:    CastBACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged(reinitializedStateOfDevice),
		Password:                      CastBACnetContextTagCharacterString(password),
		BACnetConfirmedServiceRequest: &BACnetConfirmedServiceRequest{},
	}
	_child.BACnetConfirmedServiceRequest.Child = _child
	return _child, nil
}

func (m *BACnetConfirmedServiceRequestReinitializeDevice) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestReinitializeDevice"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestReinitializeDevice")
		}

		// Simple Field (reinitializedStateOfDevice)
		if pushErr := writeBuffer.PushContext("reinitializedStateOfDevice"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for reinitializedStateOfDevice")
		}
		_reinitializedStateOfDeviceErr := writeBuffer.WriteSerializable(m.ReinitializedStateOfDevice)
		if popErr := writeBuffer.PopContext("reinitializedStateOfDevice"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for reinitializedStateOfDevice")
		}
		if _reinitializedStateOfDeviceErr != nil {
			return errors.Wrap(_reinitializedStateOfDeviceErr, "Error serializing 'reinitializedStateOfDevice' field")
		}

		// Optional Field (password) (Can be skipped, if the value is null)
		var password *BACnetContextTagCharacterString = nil
		if m.Password != nil {
			if pushErr := writeBuffer.PushContext("password"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for password")
			}
			password = m.Password
			_passwordErr := writeBuffer.WriteSerializable(password)
			if popErr := writeBuffer.PopContext("password"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for password")
			}
			if _passwordErr != nil {
				return errors.Wrap(_passwordErr, "Error serializing 'password' field")
			}
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestReinitializeDevice"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestReinitializeDevice")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConfirmedServiceRequestReinitializeDevice) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
