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

// ReplyNetwork is the corresponding interface of ReplyNetwork
type ReplyNetwork interface {
	utils.LengthAware
	utils.Serializable
	// GetRouteType returns RouteType (property field)
	GetRouteType() RouteType
	// GetAdditionalBridgeAddresses returns AdditionalBridgeAddresses (property field)
	GetAdditionalBridgeAddresses() []BridgeAddress
	// GetUnitAddress returns UnitAddress (property field)
	GetUnitAddress() UnitAddress
}

// ReplyNetworkExactly can be used when we want exactly this type and not a type which fulfills ReplyNetwork.
// This is useful for switch cases.
type ReplyNetworkExactly interface {
	ReplyNetwork
	isReplyNetwork() bool
}

// _ReplyNetwork is the data-structure of this message
type _ReplyNetwork struct {
	RouteType                 RouteType
	AdditionalBridgeAddresses []BridgeAddress
	UnitAddress               UnitAddress
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ReplyNetwork) GetRouteType() RouteType {
	return m.RouteType
}

func (m *_ReplyNetwork) GetAdditionalBridgeAddresses() []BridgeAddress {
	return m.AdditionalBridgeAddresses
}

func (m *_ReplyNetwork) GetUnitAddress() UnitAddress {
	return m.UnitAddress
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewReplyNetwork factory function for _ReplyNetwork
func NewReplyNetwork(routeType RouteType, additionalBridgeAddresses []BridgeAddress, unitAddress UnitAddress) *_ReplyNetwork {
	return &_ReplyNetwork{RouteType: routeType, AdditionalBridgeAddresses: additionalBridgeAddresses, UnitAddress: unitAddress}
}

// Deprecated: use the interface for direct cast
func CastReplyNetwork(structType interface{}) ReplyNetwork {
	if casted, ok := structType.(ReplyNetwork); ok {
		return casted
	}
	if casted, ok := structType.(*ReplyNetwork); ok {
		return *casted
	}
	return nil
}

func (m *_ReplyNetwork) GetTypeName() string {
	return "ReplyNetwork"
}

func (m *_ReplyNetwork) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_ReplyNetwork) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (routeType)
	lengthInBits += 8

	// Array field
	if len(m.AdditionalBridgeAddresses) > 0 {
		for i, element := range m.AdditionalBridgeAddresses {
			last := i == len(m.AdditionalBridgeAddresses)-1
			lengthInBits += element.(interface{ GetLengthInBitsConditional(bool) uint16 }).GetLengthInBitsConditional(last)
		}
	}

	// Simple field (unitAddress)
	lengthInBits += m.UnitAddress.GetLengthInBits()

	return lengthInBits
}

func (m *_ReplyNetwork) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ReplyNetworkParse(readBuffer utils.ReadBuffer) (ReplyNetwork, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ReplyNetwork"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ReplyNetwork")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (routeType)
	if pullErr := readBuffer.PullContext("routeType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for routeType")
	}
	_routeType, _routeTypeErr := RouteTypeParse(readBuffer)
	if _routeTypeErr != nil {
		return nil, errors.Wrap(_routeTypeErr, "Error parsing 'routeType' field")
	}
	routeType := _routeType
	if closeErr := readBuffer.CloseContext("routeType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for routeType")
	}

	// Array field (additionalBridgeAddresses)
	if pullErr := readBuffer.PullContext("additionalBridgeAddresses", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for additionalBridgeAddresses")
	}
	// Count array
	additionalBridgeAddresses := make([]BridgeAddress, routeType.AdditionalBridges())
	{
		for curItem := uint16(0); curItem < uint16(routeType.AdditionalBridges()); curItem++ {
			_item, _err := BridgeAddressParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'additionalBridgeAddresses' field")
			}
			additionalBridgeAddresses[curItem] = _item.(BridgeAddress)
		}
	}
	if closeErr := readBuffer.CloseContext("additionalBridgeAddresses", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for additionalBridgeAddresses")
	}

	// Simple Field (unitAddress)
	if pullErr := readBuffer.PullContext("unitAddress"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for unitAddress")
	}
	_unitAddress, _unitAddressErr := UnitAddressParse(readBuffer)
	if _unitAddressErr != nil {
		return nil, errors.Wrap(_unitAddressErr, "Error parsing 'unitAddress' field")
	}
	unitAddress := _unitAddress.(UnitAddress)
	if closeErr := readBuffer.CloseContext("unitAddress"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for unitAddress")
	}

	if closeErr := readBuffer.CloseContext("ReplyNetwork"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ReplyNetwork")
	}

	// Create the instance
	return NewReplyNetwork(routeType, additionalBridgeAddresses, unitAddress), nil
}

func (m *_ReplyNetwork) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("ReplyNetwork"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ReplyNetwork")
	}

	// Simple Field (routeType)
	if pushErr := writeBuffer.PushContext("routeType"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for routeType")
	}
	_routeTypeErr := writeBuffer.WriteSerializable(m.GetRouteType())
	if popErr := writeBuffer.PopContext("routeType"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for routeType")
	}
	if _routeTypeErr != nil {
		return errors.Wrap(_routeTypeErr, "Error serializing 'routeType' field")
	}

	// Array Field (additionalBridgeAddresses)
	if m.GetAdditionalBridgeAddresses() != nil {
		if pushErr := writeBuffer.PushContext("additionalBridgeAddresses", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for additionalBridgeAddresses")
		}
		for _, _element := range m.GetAdditionalBridgeAddresses() {
			_elementErr := writeBuffer.WriteSerializable(_element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'additionalBridgeAddresses' field")
			}
		}
		if popErr := writeBuffer.PopContext("additionalBridgeAddresses", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for additionalBridgeAddresses")
		}
	}

	// Simple Field (unitAddress)
	if pushErr := writeBuffer.PushContext("unitAddress"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for unitAddress")
	}
	_unitAddressErr := writeBuffer.WriteSerializable(m.GetUnitAddress())
	if popErr := writeBuffer.PopContext("unitAddress"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for unitAddress")
	}
	if _unitAddressErr != nil {
		return errors.Wrap(_unitAddressErr, "Error serializing 'unitAddress' field")
	}

	if popErr := writeBuffer.PopContext("ReplyNetwork"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ReplyNetwork")
	}
	return nil
}

func (m *_ReplyNetwork) isReplyNetwork() bool {
	return true
}

func (m *_ReplyNetwork) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
