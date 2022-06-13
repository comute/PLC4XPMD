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

// CBusPointToPointCommandIndirect is the data-structure of this message
type CBusPointToPointCommandIndirect struct {
	*CBusPointToPointCommand
	BridgeAddress *BridgeAddress
	NetworkRoute  *NetworkRoute
	UnitAddress   *UnitAddress

	// Arguments.
	Srchk bool
}

// ICBusPointToPointCommandIndirect is the corresponding interface of CBusPointToPointCommandIndirect
type ICBusPointToPointCommandIndirect interface {
	ICBusPointToPointCommand
	// GetBridgeAddress returns BridgeAddress (property field)
	GetBridgeAddress() *BridgeAddress
	// GetNetworkRoute returns NetworkRoute (property field)
	GetNetworkRoute() *NetworkRoute
	// GetUnitAddress returns UnitAddress (property field)
	GetUnitAddress() *UnitAddress
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

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *CBusPointToPointCommandIndirect) InitializeParent(parent *CBusPointToPointCommand, bridgeAddressCountPeek uint16, calData *CALData, crc *Checksum, peekAlpha byte, alpha *Alpha) {
	m.CBusPointToPointCommand.BridgeAddressCountPeek = bridgeAddressCountPeek
	m.CBusPointToPointCommand.CalData = calData
	m.CBusPointToPointCommand.Crc = crc
	m.CBusPointToPointCommand.PeekAlpha = peekAlpha
	m.CBusPointToPointCommand.Alpha = alpha
}

func (m *CBusPointToPointCommandIndirect) GetParent() *CBusPointToPointCommand {
	return m.CBusPointToPointCommand
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *CBusPointToPointCommandIndirect) GetBridgeAddress() *BridgeAddress {
	return m.BridgeAddress
}

func (m *CBusPointToPointCommandIndirect) GetNetworkRoute() *NetworkRoute {
	return m.NetworkRoute
}

func (m *CBusPointToPointCommandIndirect) GetUnitAddress() *UnitAddress {
	return m.UnitAddress
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCBusPointToPointCommandIndirect factory function for CBusPointToPointCommandIndirect
func NewCBusPointToPointCommandIndirect(bridgeAddress *BridgeAddress, networkRoute *NetworkRoute, unitAddress *UnitAddress, bridgeAddressCountPeek uint16, calData *CALData, crc *Checksum, peekAlpha byte, alpha *Alpha, srchk bool) *CBusPointToPointCommandIndirect {
	_result := &CBusPointToPointCommandIndirect{
		BridgeAddress:           bridgeAddress,
		NetworkRoute:            networkRoute,
		UnitAddress:             unitAddress,
		CBusPointToPointCommand: NewCBusPointToPointCommand(bridgeAddressCountPeek, calData, crc, peekAlpha, alpha, srchk),
	}
	_result.Child = _result
	return _result
}

func CastCBusPointToPointCommandIndirect(structType interface{}) *CBusPointToPointCommandIndirect {
	if casted, ok := structType.(CBusPointToPointCommandIndirect); ok {
		return &casted
	}
	if casted, ok := structType.(*CBusPointToPointCommandIndirect); ok {
		return casted
	}
	if casted, ok := structType.(CBusPointToPointCommand); ok {
		return CastCBusPointToPointCommandIndirect(casted.Child)
	}
	if casted, ok := structType.(*CBusPointToPointCommand); ok {
		return CastCBusPointToPointCommandIndirect(casted.Child)
	}
	return nil
}

func (m *CBusPointToPointCommandIndirect) GetTypeName() string {
	return "CBusPointToPointCommandIndirect"
}

func (m *CBusPointToPointCommandIndirect) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *CBusPointToPointCommandIndirect) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (bridgeAddress)
	lengthInBits += m.BridgeAddress.GetLengthInBits()

	// Simple field (networkRoute)
	lengthInBits += m.NetworkRoute.GetLengthInBits()

	// Simple field (unitAddress)
	lengthInBits += m.UnitAddress.GetLengthInBits()

	return lengthInBits
}

func (m *CBusPointToPointCommandIndirect) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func CBusPointToPointCommandIndirectParse(readBuffer utils.ReadBuffer, srchk bool) (*CBusPointToPointCommandIndirect, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CBusPointToPointCommandIndirect"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CBusPointToPointCommandIndirect")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (bridgeAddress)
	if pullErr := readBuffer.PullContext("bridgeAddress"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for bridgeAddress")
	}
	_bridgeAddress, _bridgeAddressErr := BridgeAddressParse(readBuffer)
	if _bridgeAddressErr != nil {
		return nil, errors.Wrap(_bridgeAddressErr, "Error parsing 'bridgeAddress' field")
	}
	bridgeAddress := CastBridgeAddress(_bridgeAddress)
	if closeErr := readBuffer.CloseContext("bridgeAddress"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for bridgeAddress")
	}

	// Simple Field (networkRoute)
	if pullErr := readBuffer.PullContext("networkRoute"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for networkRoute")
	}
	_networkRoute, _networkRouteErr := NetworkRouteParse(readBuffer)
	if _networkRouteErr != nil {
		return nil, errors.Wrap(_networkRouteErr, "Error parsing 'networkRoute' field")
	}
	networkRoute := CastNetworkRoute(_networkRoute)
	if closeErr := readBuffer.CloseContext("networkRoute"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for networkRoute")
	}

	// Simple Field (unitAddress)
	if pullErr := readBuffer.PullContext("unitAddress"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for unitAddress")
	}
	_unitAddress, _unitAddressErr := UnitAddressParse(readBuffer)
	if _unitAddressErr != nil {
		return nil, errors.Wrap(_unitAddressErr, "Error parsing 'unitAddress' field")
	}
	unitAddress := CastUnitAddress(_unitAddress)
	if closeErr := readBuffer.CloseContext("unitAddress"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for unitAddress")
	}

	if closeErr := readBuffer.CloseContext("CBusPointToPointCommandIndirect"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CBusPointToPointCommandIndirect")
	}

	// Create a partially initialized instance
	_child := &CBusPointToPointCommandIndirect{
		BridgeAddress:           CastBridgeAddress(bridgeAddress),
		NetworkRoute:            CastNetworkRoute(networkRoute),
		UnitAddress:             CastUnitAddress(unitAddress),
		CBusPointToPointCommand: &CBusPointToPointCommand{},
	}
	_child.CBusPointToPointCommand.Child = _child
	return _child, nil
}

func (m *CBusPointToPointCommandIndirect) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CBusPointToPointCommandIndirect"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CBusPointToPointCommandIndirect")
		}

		// Simple Field (bridgeAddress)
		if pushErr := writeBuffer.PushContext("bridgeAddress"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for bridgeAddress")
		}
		_bridgeAddressErr := writeBuffer.WriteSerializable(m.BridgeAddress)
		if popErr := writeBuffer.PopContext("bridgeAddress"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for bridgeAddress")
		}
		if _bridgeAddressErr != nil {
			return errors.Wrap(_bridgeAddressErr, "Error serializing 'bridgeAddress' field")
		}

		// Simple Field (networkRoute)
		if pushErr := writeBuffer.PushContext("networkRoute"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for networkRoute")
		}
		_networkRouteErr := writeBuffer.WriteSerializable(m.NetworkRoute)
		if popErr := writeBuffer.PopContext("networkRoute"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for networkRoute")
		}
		if _networkRouteErr != nil {
			return errors.Wrap(_networkRouteErr, "Error serializing 'networkRoute' field")
		}

		// Simple Field (unitAddress)
		if pushErr := writeBuffer.PushContext("unitAddress"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for unitAddress")
		}
		_unitAddressErr := writeBuffer.WriteSerializable(m.UnitAddress)
		if popErr := writeBuffer.PopContext("unitAddress"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for unitAddress")
		}
		if _unitAddressErr != nil {
			return errors.Wrap(_unitAddressErr, "Error serializing 'unitAddress' field")
		}

		if popErr := writeBuffer.PopContext("CBusPointToPointCommandIndirect"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CBusPointToPointCommandIndirect")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *CBusPointToPointCommandIndirect) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
