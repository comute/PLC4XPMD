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

// DescriptionResponse is the corresponding interface of DescriptionResponse
type DescriptionResponse interface {
	utils.LengthAware
	utils.Serializable
	KnxNetIpMessage
	// GetDibDeviceInfo returns DibDeviceInfo (property field)
	GetDibDeviceInfo() DIBDeviceInfo
	// GetDibSuppSvcFamilies returns DibSuppSvcFamilies (property field)
	GetDibSuppSvcFamilies() DIBSuppSvcFamilies
}

// DescriptionResponseExactly can be used when we want exactly this type and not a type which fulfills DescriptionResponse.
// This is useful for switch cases.
type DescriptionResponseExactly interface {
	DescriptionResponse
	isDescriptionResponse() bool
}

// _DescriptionResponse is the data-structure of this message
type _DescriptionResponse struct {
	*_KnxNetIpMessage
	DibDeviceInfo      DIBDeviceInfo
	DibSuppSvcFamilies DIBSuppSvcFamilies
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_DescriptionResponse) GetMsgType() uint16 {
	return 0x0204
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_DescriptionResponse) InitializeParent(parent KnxNetIpMessage) {}

func (m *_DescriptionResponse) GetParent() KnxNetIpMessage {
	return m._KnxNetIpMessage
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_DescriptionResponse) GetDibDeviceInfo() DIBDeviceInfo {
	return m.DibDeviceInfo
}

func (m *_DescriptionResponse) GetDibSuppSvcFamilies() DIBSuppSvcFamilies {
	return m.DibSuppSvcFamilies
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewDescriptionResponse factory function for _DescriptionResponse
func NewDescriptionResponse(dibDeviceInfo DIBDeviceInfo, dibSuppSvcFamilies DIBSuppSvcFamilies) *_DescriptionResponse {
	_result := &_DescriptionResponse{
		DibDeviceInfo:      dibDeviceInfo,
		DibSuppSvcFamilies: dibSuppSvcFamilies,
		_KnxNetIpMessage:   NewKnxNetIpMessage(),
	}
	_result._KnxNetIpMessage._KnxNetIpMessageChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastDescriptionResponse(structType interface{}) DescriptionResponse {
	if casted, ok := structType.(DescriptionResponse); ok {
		return casted
	}
	if casted, ok := structType.(*DescriptionResponse); ok {
		return *casted
	}
	return nil
}

func (m *_DescriptionResponse) GetTypeName() string {
	return "DescriptionResponse"
}

func (m *_DescriptionResponse) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_DescriptionResponse) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (dibDeviceInfo)
	lengthInBits += m.DibDeviceInfo.GetLengthInBits()

	// Simple field (dibSuppSvcFamilies)
	lengthInBits += m.DibSuppSvcFamilies.GetLengthInBits()

	return lengthInBits
}

func (m *_DescriptionResponse) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func DescriptionResponseParse(readBuffer utils.ReadBuffer) (DescriptionResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("DescriptionResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DescriptionResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (dibDeviceInfo)
	if pullErr := readBuffer.PullContext("dibDeviceInfo"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for dibDeviceInfo")
	}
	_dibDeviceInfo, _dibDeviceInfoErr := DIBDeviceInfoParse(readBuffer)
	if _dibDeviceInfoErr != nil {
		return nil, errors.Wrap(_dibDeviceInfoErr, "Error parsing 'dibDeviceInfo' field")
	}
	dibDeviceInfo := _dibDeviceInfo.(DIBDeviceInfo)
	if closeErr := readBuffer.CloseContext("dibDeviceInfo"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for dibDeviceInfo")
	}

	// Simple Field (dibSuppSvcFamilies)
	if pullErr := readBuffer.PullContext("dibSuppSvcFamilies"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for dibSuppSvcFamilies")
	}
	_dibSuppSvcFamilies, _dibSuppSvcFamiliesErr := DIBSuppSvcFamiliesParse(readBuffer)
	if _dibSuppSvcFamiliesErr != nil {
		return nil, errors.Wrap(_dibSuppSvcFamiliesErr, "Error parsing 'dibSuppSvcFamilies' field")
	}
	dibSuppSvcFamilies := _dibSuppSvcFamilies.(DIBSuppSvcFamilies)
	if closeErr := readBuffer.CloseContext("dibSuppSvcFamilies"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for dibSuppSvcFamilies")
	}

	if closeErr := readBuffer.CloseContext("DescriptionResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DescriptionResponse")
	}

	// Create a partially initialized instance
	_child := &_DescriptionResponse{
		DibDeviceInfo:      dibDeviceInfo,
		DibSuppSvcFamilies: dibSuppSvcFamilies,
		_KnxNetIpMessage:   &_KnxNetIpMessage{},
	}
	_child._KnxNetIpMessage._KnxNetIpMessageChildRequirements = _child
	return _child, nil
}

func (m *_DescriptionResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("DescriptionResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for DescriptionResponse")
		}

		// Simple Field (dibDeviceInfo)
		if pushErr := writeBuffer.PushContext("dibDeviceInfo"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for dibDeviceInfo")
		}
		_dibDeviceInfoErr := writeBuffer.WriteSerializable(m.GetDibDeviceInfo())
		if popErr := writeBuffer.PopContext("dibDeviceInfo"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for dibDeviceInfo")
		}
		if _dibDeviceInfoErr != nil {
			return errors.Wrap(_dibDeviceInfoErr, "Error serializing 'dibDeviceInfo' field")
		}

		// Simple Field (dibSuppSvcFamilies)
		if pushErr := writeBuffer.PushContext("dibSuppSvcFamilies"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for dibSuppSvcFamilies")
		}
		_dibSuppSvcFamiliesErr := writeBuffer.WriteSerializable(m.GetDibSuppSvcFamilies())
		if popErr := writeBuffer.PopContext("dibSuppSvcFamilies"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for dibSuppSvcFamilies")
		}
		if _dibSuppSvcFamiliesErr != nil {
			return errors.Wrap(_dibSuppSvcFamiliesErr, "Error serializing 'dibSuppSvcFamilies' field")
		}

		if popErr := writeBuffer.PopContext("DescriptionResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for DescriptionResponse")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_DescriptionResponse) isDescriptionResponse() bool {
	return true
}

func (m *_DescriptionResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
