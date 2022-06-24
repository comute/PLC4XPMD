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

// NLMRouterBusyToNetwork is the corresponding interface of NLMRouterBusyToNetwork
type NLMRouterBusyToNetwork interface {
	utils.LengthAware
	utils.Serializable
	NLM
	// GetDestinationNetworkAddress returns DestinationNetworkAddress (property field)
	GetDestinationNetworkAddress() []uint16
}

// NLMRouterBusyToNetworkExactly can be used when we want exactly this type and not a type which fulfills NLMRouterBusyToNetwork.
// This is useful for switch cases.
type NLMRouterBusyToNetworkExactly interface {
	NLMRouterBusyToNetwork
	isNLMRouterBusyToNetwork() bool
}

// _NLMRouterBusyToNetwork is the data-structure of this message
type _NLMRouterBusyToNetwork struct {
	*_NLM
	DestinationNetworkAddress []uint16
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_NLMRouterBusyToNetwork) GetMessageType() uint8 {
	return 0x04
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_NLMRouterBusyToNetwork) InitializeParent(parent NLM, vendorId *BACnetVendorId) {
	m.VendorId = vendorId
}

func (m *_NLMRouterBusyToNetwork) GetParent() NLM {
	return m._NLM
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_NLMRouterBusyToNetwork) GetDestinationNetworkAddress() []uint16 {
	return m.DestinationNetworkAddress
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewNLMRouterBusyToNetwork factory function for _NLMRouterBusyToNetwork
func NewNLMRouterBusyToNetwork(destinationNetworkAddress []uint16, vendorId *BACnetVendorId, apduLength uint16) *_NLMRouterBusyToNetwork {
	_result := &_NLMRouterBusyToNetwork{
		DestinationNetworkAddress: destinationNetworkAddress,
		_NLM:                      NewNLM(vendorId, apduLength),
	}
	_result._NLM._NLMChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastNLMRouterBusyToNetwork(structType interface{}) NLMRouterBusyToNetwork {
	if casted, ok := structType.(NLMRouterBusyToNetwork); ok {
		return casted
	}
	if casted, ok := structType.(*NLMRouterBusyToNetwork); ok {
		return *casted
	}
	return nil
}

func (m *_NLMRouterBusyToNetwork) GetTypeName() string {
	return "NLMRouterBusyToNetwork"
}

func (m *_NLMRouterBusyToNetwork) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_NLMRouterBusyToNetwork) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.DestinationNetworkAddress) > 0 {
		lengthInBits += 16 * uint16(len(m.DestinationNetworkAddress))
	}

	return lengthInBits
}

func (m *_NLMRouterBusyToNetwork) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func NLMRouterBusyToNetworkParse(readBuffer utils.ReadBuffer, apduLength uint16, messageType uint8) (NLMRouterBusyToNetwork, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("NLMRouterBusyToNetwork"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NLMRouterBusyToNetwork")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (destinationNetworkAddress)
	if pullErr := readBuffer.PullContext("destinationNetworkAddress", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for destinationNetworkAddress")
	}
	// Length array
	var destinationNetworkAddress []uint16
	{
		_destinationNetworkAddressLength := uint16(apduLength) - uint16(uint16(utils.InlineIf(bool(bool(bool(bool((messageType) >= (128)))) && bool(bool(bool((messageType) <= (255))))), func() interface{} { return uint16(uint16(3)) }, func() interface{} { return uint16(uint16(1)) }).(uint16)))
		_destinationNetworkAddressEndPos := positionAware.GetPos() + uint16(_destinationNetworkAddressLength)
		for positionAware.GetPos() < _destinationNetworkAddressEndPos {
			_item, _err := readBuffer.ReadUint16("", 16)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'destinationNetworkAddress' field")
			}
			destinationNetworkAddress = append(destinationNetworkAddress, _item)
		}
	}
	if closeErr := readBuffer.CloseContext("destinationNetworkAddress", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for destinationNetworkAddress")
	}

	if closeErr := readBuffer.CloseContext("NLMRouterBusyToNetwork"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NLMRouterBusyToNetwork")
	}

	// Create a partially initialized instance
	_child := &_NLMRouterBusyToNetwork{
		DestinationNetworkAddress: destinationNetworkAddress,
		_NLM: &_NLM{
			ApduLength: apduLength,
		},
	}
	_child._NLM._NLMChildRequirements = _child
	return _child, nil
}

func (m *_NLMRouterBusyToNetwork) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("NLMRouterBusyToNetwork"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for NLMRouterBusyToNetwork")
		}

		// Array Field (destinationNetworkAddress)
		if pushErr := writeBuffer.PushContext("destinationNetworkAddress", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for destinationNetworkAddress")
		}
		for _, _element := range m.GetDestinationNetworkAddress() {
			_elementErr := writeBuffer.WriteUint16("", 16, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'destinationNetworkAddress' field")
			}
		}
		if popErr := writeBuffer.PopContext("destinationNetworkAddress", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for destinationNetworkAddress")
		}

		if popErr := writeBuffer.PopContext("NLMRouterBusyToNetwork"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for NLMRouterBusyToNetwork")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_NLMRouterBusyToNetwork) isNLMRouterBusyToNetwork() bool {
	return true
}

func (m *_NLMRouterBusyToNetwork) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
