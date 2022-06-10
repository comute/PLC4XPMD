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

// NLM is the data-structure of this message
type NLM struct {
	VendorId *BACnetVendorId

	// Arguments.
	ApduLength uint16
	Child      INLMChild
}

// INLM is the corresponding interface of NLM
type INLM interface {
	// GetMessageType returns MessageType (discriminator field)
	GetMessageType() uint8
	// GetVendorId returns VendorId (property field)
	GetVendorId() *BACnetVendorId
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

type INLMParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child INLM, serializeChildFunction func() error) error
	GetTypeName() string
}

type INLMChild interface {
	Serialize(writeBuffer utils.WriteBuffer) error
	InitializeParent(parent *NLM, vendorId *BACnetVendorId)
	GetParent() *NLM

	GetTypeName() string
	INLM
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *NLM) GetVendorId() *BACnetVendorId {
	return m.VendorId
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewNLM factory function for NLM
func NewNLM(vendorId *BACnetVendorId, apduLength uint16) *NLM {
	return &NLM{VendorId: vendorId, ApduLength: apduLength}
}

func CastNLM(structType interface{}) *NLM {
	if casted, ok := structType.(NLM); ok {
		return &casted
	}
	if casted, ok := structType.(*NLM); ok {
		return casted
	}
	if casted, ok := structType.(INLMChild); ok {
		return casted.GetParent()
	}
	return nil
}

func (m *NLM) GetTypeName() string {
	return "NLM"
}

func (m *NLM) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *NLM) GetLengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.GetLengthInBits()
}

func (m *NLM) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (messageType)
	lengthInBits += 8

	// Optional Field (vendorId)
	if m.VendorId != nil {
		lengthInBits += 16
	}

	return lengthInBits
}

func (m *NLM) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func NLMParse(readBuffer utils.ReadBuffer, apduLength uint16) (*NLM, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("NLM"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NLM")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Discriminator Field (messageType) (Used as input to a switch field)
	messageType, _messageTypeErr := readBuffer.ReadUint8("messageType", 8)
	if _messageTypeErr != nil {
		return nil, errors.Wrap(_messageTypeErr, "Error parsing 'messageType' field")
	}

	// Optional Field (vendorId) (Can be skipped, if a given expression evaluates to false)
	var vendorId *BACnetVendorId = nil
	if bool(bool(bool((messageType) >= (128)))) && bool(bool(bool((messageType) <= (255)))) {
		if pullErr := readBuffer.PullContext("vendorId"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for vendorId")
		}
		_val, _err := BACnetVendorIdParse(readBuffer)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'vendorId' field")
		}
		vendorId = &_val
		if closeErr := readBuffer.CloseContext("vendorId"); closeErr != nil {
			return nil, errors.Wrap(closeErr, "Error closing for vendorId")
		}
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type NLMChild interface {
		InitializeParent(*NLM, *BACnetVendorId)
		GetParent() *NLM
	}
	var _child NLMChild
	var typeSwitchError error
	switch {
	case messageType == 0x00: // NLMWhoIsRouterToNetwork
		_child, typeSwitchError = NLMWhoIsRouterToNetworkParse(readBuffer, apduLength, messageType)
	case messageType == 0x01: // NLMIAmRouterToNetwork
		_child, typeSwitchError = NLMIAmRouterToNetworkParse(readBuffer, apduLength, messageType)
	case messageType == 0x02: // NLMICouldBeRouterToNetwork
		_child, typeSwitchError = NLMICouldBeRouterToNetworkParse(readBuffer, apduLength, messageType)
	case messageType == 0x03: // NLMRejectRouterToNetwork
		_child, typeSwitchError = NLMRejectRouterToNetworkParse(readBuffer, apduLength, messageType)
	case messageType == 0x04: // NLMRouterBusyToNetwork
		_child, typeSwitchError = NLMRouterBusyToNetworkParse(readBuffer, apduLength, messageType)
	case messageType == 0x05: // NLMRouterAvailableToNetwork
		_child, typeSwitchError = NLMRouterAvailableToNetworkParse(readBuffer, apduLength, messageType)
	case messageType == 0x06: // NLMInitalizeRoutingTable
		_child, typeSwitchError = NLMInitalizeRoutingTableParse(readBuffer, apduLength, messageType)
	case messageType == 0x07: // NLMInitalizeRoutingTableAck
		_child, typeSwitchError = NLMInitalizeRoutingTableAckParse(readBuffer, apduLength, messageType)
	case messageType == 0x08: // NLMEstablishConnectionToNetwork
		_child, typeSwitchError = NLMEstablishConnectionToNetworkParse(readBuffer, apduLength, messageType)
	case messageType == 0x09: // NLMDisconnectConnectionToNetwork
		_child, typeSwitchError = NLMDisconnectConnectionToNetworkParse(readBuffer, apduLength, messageType)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	if closeErr := readBuffer.CloseContext("NLM"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NLM")
	}

	// Finish initializing
	_child.InitializeParent(_child.GetParent(), vendorId)
	return _child.GetParent(), nil
}

func (m *NLM) Serialize(writeBuffer utils.WriteBuffer) error {
	return m.Child.Serialize(writeBuffer)
}

func (m *NLM) SerializeParent(writeBuffer utils.WriteBuffer, child INLM, serializeChildFunction func() error) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("NLM"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for NLM")
	}

	// Discriminator Field (messageType) (Used as input to a switch field)
	messageType := uint8(child.GetMessageType())
	_messageTypeErr := writeBuffer.WriteUint8("messageType", 8, (messageType))

	if _messageTypeErr != nil {
		return errors.Wrap(_messageTypeErr, "Error serializing 'messageType' field")
	}

	// Optional Field (vendorId) (Can be skipped, if the value is null)
	var vendorId *BACnetVendorId = nil
	if m.VendorId != nil {
		if pushErr := writeBuffer.PushContext("vendorId"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for vendorId")
		}
		vendorId = m.VendorId
		_vendorIdErr := vendorId.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("vendorId"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for vendorId")
		}
		if _vendorIdErr != nil {
			return errors.Wrap(_vendorIdErr, "Error serializing 'vendorId' field")
		}
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("NLM"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for NLM")
	}
	return nil
}

func (m *NLM) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
