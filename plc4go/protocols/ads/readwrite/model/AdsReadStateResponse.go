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

// AdsReadStateResponse is the corresponding interface of AdsReadStateResponse
type AdsReadStateResponse interface {
	utils.LengthAware
	utils.Serializable
	AdsData
	// GetResult returns Result (property field)
	GetResult() ReturnCode
	// GetAdsState returns AdsState (property field)
	GetAdsState() uint16
	// GetDeviceState returns DeviceState (property field)
	GetDeviceState() uint16
}

// AdsReadStateResponseExactly can be used when we want exactly this type and not a type which fulfills AdsReadStateResponse.
// This is useful for switch cases.
type AdsReadStateResponseExactly interface {
	AdsReadStateResponse
	isAdsReadStateResponse() bool
}

// _AdsReadStateResponse is the data-structure of this message
type _AdsReadStateResponse struct {
	*_AdsData
	Result      ReturnCode
	AdsState    uint16
	DeviceState uint16
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_AdsReadStateResponse) GetCommandId() CommandId {
	return CommandId_ADS_READ_STATE
}

func (m *_AdsReadStateResponse) GetResponse() bool {
	return bool(true)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AdsReadStateResponse) InitializeParent(parent AdsData) {}

func (m *_AdsReadStateResponse) GetParent() AdsData {
	return m._AdsData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AdsReadStateResponse) GetResult() ReturnCode {
	return m.Result
}

func (m *_AdsReadStateResponse) GetAdsState() uint16 {
	return m.AdsState
}

func (m *_AdsReadStateResponse) GetDeviceState() uint16 {
	return m.DeviceState
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAdsReadStateResponse factory function for _AdsReadStateResponse
func NewAdsReadStateResponse(result ReturnCode, adsState uint16, deviceState uint16) *_AdsReadStateResponse {
	_result := &_AdsReadStateResponse{
		Result:      result,
		AdsState:    adsState,
		DeviceState: deviceState,
		_AdsData:    NewAdsData(),
	}
	_result._AdsData._AdsDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastAdsReadStateResponse(structType interface{}) AdsReadStateResponse {
	if casted, ok := structType.(AdsReadStateResponse); ok {
		return casted
	}
	if casted, ok := structType.(*AdsReadStateResponse); ok {
		return *casted
	}
	return nil
}

func (m *_AdsReadStateResponse) GetTypeName() string {
	return "AdsReadStateResponse"
}

func (m *_AdsReadStateResponse) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_AdsReadStateResponse) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (result)
	lengthInBits += 32

	// Simple field (adsState)
	lengthInBits += 16

	// Simple field (deviceState)
	lengthInBits += 16

	return lengthInBits
}

func (m *_AdsReadStateResponse) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func AdsReadStateResponseParse(readBuffer utils.ReadBuffer, commandId CommandId, response bool) (AdsReadStateResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AdsReadStateResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AdsReadStateResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (result)
	if pullErr := readBuffer.PullContext("result"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for result")
	}
	_result, _resultErr := ReturnCodeParse(readBuffer)
	if _resultErr != nil {
		return nil, errors.Wrap(_resultErr, "Error parsing 'result' field")
	}
	result := _result
	if closeErr := readBuffer.CloseContext("result"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for result")
	}

	// Simple Field (adsState)
	_adsState, _adsStateErr := readBuffer.ReadUint16("adsState", 16)
	if _adsStateErr != nil {
		return nil, errors.Wrap(_adsStateErr, "Error parsing 'adsState' field")
	}
	adsState := _adsState

	// Simple Field (deviceState)
	_deviceState, _deviceStateErr := readBuffer.ReadUint16("deviceState", 16)
	if _deviceStateErr != nil {
		return nil, errors.Wrap(_deviceStateErr, "Error parsing 'deviceState' field")
	}
	deviceState := _deviceState

	if closeErr := readBuffer.CloseContext("AdsReadStateResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AdsReadStateResponse")
	}

	// Create a partially initialized instance
	_child := &_AdsReadStateResponse{
		Result:      result,
		AdsState:    adsState,
		DeviceState: deviceState,
		_AdsData:    &_AdsData{},
	}
	_child._AdsData._AdsDataChildRequirements = _child
	return _child, nil
}

func (m *_AdsReadStateResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AdsReadStateResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AdsReadStateResponse")
		}

		// Simple Field (result)
		if pushErr := writeBuffer.PushContext("result"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for result")
		}
		_resultErr := writeBuffer.WriteSerializable(m.GetResult())
		if popErr := writeBuffer.PopContext("result"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for result")
		}
		if _resultErr != nil {
			return errors.Wrap(_resultErr, "Error serializing 'result' field")
		}

		// Simple Field (adsState)
		adsState := uint16(m.GetAdsState())
		_adsStateErr := writeBuffer.WriteUint16("adsState", 16, (adsState))
		if _adsStateErr != nil {
			return errors.Wrap(_adsStateErr, "Error serializing 'adsState' field")
		}

		// Simple Field (deviceState)
		deviceState := uint16(m.GetDeviceState())
		_deviceStateErr := writeBuffer.WriteUint16("deviceState", 16, (deviceState))
		if _deviceStateErr != nil {
			return errors.Wrap(_deviceStateErr, "Error serializing 'deviceState' field")
		}

		if popErr := writeBuffer.PopContext("AdsReadStateResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AdsReadStateResponse")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_AdsReadStateResponse) isAdsReadStateResponse() bool {
	return true
}

func (m *_AdsReadStateResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
