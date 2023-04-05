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
	"context"
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// StatusRequestBinaryStateDeprecated is the corresponding interface of StatusRequestBinaryStateDeprecated
type StatusRequestBinaryStateDeprecated interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	StatusRequest
	// GetApplication returns Application (property field)
	GetApplication() ApplicationIdContainer
}

// StatusRequestBinaryStateDeprecatedExactly can be used when we want exactly this type and not a type which fulfills StatusRequestBinaryStateDeprecated.
// This is useful for switch cases.
type StatusRequestBinaryStateDeprecatedExactly interface {
	StatusRequestBinaryStateDeprecated
	isStatusRequestBinaryStateDeprecated() bool
}

// _StatusRequestBinaryStateDeprecated is the data-structure of this message
type _StatusRequestBinaryStateDeprecated struct {
	*_StatusRequest
	Application ApplicationIdContainer
	// Reserved Fields
	reservedField0 *byte
	reservedField1 *byte
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_StatusRequestBinaryStateDeprecated) InitializeParent(parent StatusRequest, statusType byte) {
	m.StatusType = statusType
}

func (m *_StatusRequestBinaryStateDeprecated) GetParent() StatusRequest {
	return m._StatusRequest
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_StatusRequestBinaryStateDeprecated) GetApplication() ApplicationIdContainer {
	return m.Application
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewStatusRequestBinaryStateDeprecated factory function for _StatusRequestBinaryStateDeprecated
func NewStatusRequestBinaryStateDeprecated(application ApplicationIdContainer, statusType byte) *_StatusRequestBinaryStateDeprecated {
	_result := &_StatusRequestBinaryStateDeprecated{
		Application:    application,
		_StatusRequest: NewStatusRequest(statusType),
	}
	_result._StatusRequest._StatusRequestChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastStatusRequestBinaryStateDeprecated(structType interface{}) StatusRequestBinaryStateDeprecated {
	if casted, ok := structType.(StatusRequestBinaryStateDeprecated); ok {
		return casted
	}
	if casted, ok := structType.(*StatusRequestBinaryStateDeprecated); ok {
		return *casted
	}
	return nil
}

func (m *_StatusRequestBinaryStateDeprecated) GetTypeName() string {
	return "StatusRequestBinaryStateDeprecated"
}

func (m *_StatusRequestBinaryStateDeprecated) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Reserved Field (reserved)
	lengthInBits += 8

	// Simple field (application)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 8

	return lengthInBits
}

func (m *_StatusRequestBinaryStateDeprecated) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func StatusRequestBinaryStateDeprecatedParse(theBytes []byte) (StatusRequestBinaryStateDeprecated, error) {
	return StatusRequestBinaryStateDeprecatedParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func StatusRequestBinaryStateDeprecatedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (StatusRequestBinaryStateDeprecated, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("StatusRequestBinaryStateDeprecated"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for StatusRequestBinaryStateDeprecated")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	var reservedField0 *byte
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadByte("reserved")
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of StatusRequestBinaryStateDeprecated")
		}
		if reserved != byte(0xFA) {
			Plc4xModelLog.Info().Fields(map[string]interface{}{
				"expected value": byte(0xFA),
				"got value":      reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField0 = &reserved
		}
	}

	// Simple Field (application)
	if pullErr := readBuffer.PullContext("application"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for application")
	}
	_application, _applicationErr := ApplicationIdContainerParseWithBuffer(ctx, readBuffer)
	if _applicationErr != nil {
		return nil, errors.Wrap(_applicationErr, "Error parsing 'application' field of StatusRequestBinaryStateDeprecated")
	}
	application := _application
	if closeErr := readBuffer.CloseContext("application"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for application")
	}

	var reservedField1 *byte
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadByte("reserved")
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of StatusRequestBinaryStateDeprecated")
		}
		if reserved != byte(0x00) {
			Plc4xModelLog.Info().Fields(map[string]interface{}{
				"expected value": byte(0x00),
				"got value":      reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField1 = &reserved
		}
	}

	if closeErr := readBuffer.CloseContext("StatusRequestBinaryStateDeprecated"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for StatusRequestBinaryStateDeprecated")
	}

	// Create a partially initialized instance
	_child := &_StatusRequestBinaryStateDeprecated{
		_StatusRequest: &_StatusRequest{},
		Application:    application,
		reservedField0: reservedField0,
		reservedField1: reservedField1,
	}
	_child._StatusRequest._StatusRequestChildRequirements = _child
	return _child, nil
}

func (m *_StatusRequestBinaryStateDeprecated) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_StatusRequestBinaryStateDeprecated) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("StatusRequestBinaryStateDeprecated"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for StatusRequestBinaryStateDeprecated")
		}

		// Reserved Field (reserved)
		{
			var reserved byte = byte(0xFA)
			if m.reservedField0 != nil {
				Plc4xModelLog.Info().Fields(map[string]interface{}{
					"expected value": byte(0xFA),
					"got value":      reserved,
				}).Msg("Overriding reserved field with unexpected value.")
				reserved = *m.reservedField0
			}
			_err := writeBuffer.WriteByte("reserved", reserved)
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Simple Field (application)
		if pushErr := writeBuffer.PushContext("application"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for application")
		}
		_applicationErr := writeBuffer.WriteSerializable(ctx, m.GetApplication())
		if popErr := writeBuffer.PopContext("application"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for application")
		}
		if _applicationErr != nil {
			return errors.Wrap(_applicationErr, "Error serializing 'application' field")
		}

		// Reserved Field (reserved)
		{
			var reserved byte = byte(0x00)
			if m.reservedField1 != nil {
				Plc4xModelLog.Info().Fields(map[string]interface{}{
					"expected value": byte(0x00),
					"got value":      reserved,
				}).Msg("Overriding reserved field with unexpected value.")
				reserved = *m.reservedField1
			}
			_err := writeBuffer.WriteByte("reserved", reserved)
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		if popErr := writeBuffer.PopContext("StatusRequestBinaryStateDeprecated"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for StatusRequestBinaryStateDeprecated")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_StatusRequestBinaryStateDeprecated) isStatusRequestBinaryStateDeprecated() bool {
	return true
}

func (m *_StatusRequestBinaryStateDeprecated) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
