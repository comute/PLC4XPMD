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

// BACnetPropertyStatesNetworkPortCommand is the data-structure of this message
type BACnetPropertyStatesNetworkPortCommand struct {
	*BACnetPropertyStates
	NetworkPortCommand *BACnetNetworkPortCommandTagged
}

// IBACnetPropertyStatesNetworkPortCommand is the corresponding interface of BACnetPropertyStatesNetworkPortCommand
type IBACnetPropertyStatesNetworkPortCommand interface {
	IBACnetPropertyStates
	// GetNetworkPortCommand returns NetworkPortCommand (property field)
	GetNetworkPortCommand() *BACnetNetworkPortCommandTagged
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

func (m *BACnetPropertyStatesNetworkPortCommand) InitializeParent(parent *BACnetPropertyStates, peekedTagHeader *BACnetTagHeader) {
	m.BACnetPropertyStates.PeekedTagHeader = peekedTagHeader
}

func (m *BACnetPropertyStatesNetworkPortCommand) GetParent() *BACnetPropertyStates {
	return m.BACnetPropertyStates
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetPropertyStatesNetworkPortCommand) GetNetworkPortCommand() *BACnetNetworkPortCommandTagged {
	return m.NetworkPortCommand
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStatesNetworkPortCommand factory function for BACnetPropertyStatesNetworkPortCommand
func NewBACnetPropertyStatesNetworkPortCommand(networkPortCommand *BACnetNetworkPortCommandTagged, peekedTagHeader *BACnetTagHeader) *BACnetPropertyStatesNetworkPortCommand {
	_result := &BACnetPropertyStatesNetworkPortCommand{
		NetworkPortCommand:   networkPortCommand,
		BACnetPropertyStates: NewBACnetPropertyStates(peekedTagHeader),
	}
	_result.Child = _result
	return _result
}

func CastBACnetPropertyStatesNetworkPortCommand(structType interface{}) *BACnetPropertyStatesNetworkPortCommand {
	if casted, ok := structType.(BACnetPropertyStatesNetworkPortCommand); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesNetworkPortCommand); ok {
		return casted
	}
	if casted, ok := structType.(BACnetPropertyStates); ok {
		return CastBACnetPropertyStatesNetworkPortCommand(casted.Child)
	}
	if casted, ok := structType.(*BACnetPropertyStates); ok {
		return CastBACnetPropertyStatesNetworkPortCommand(casted.Child)
	}
	return nil
}

func (m *BACnetPropertyStatesNetworkPortCommand) GetTypeName() string {
	return "BACnetPropertyStatesNetworkPortCommand"
}

func (m *BACnetPropertyStatesNetworkPortCommand) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetPropertyStatesNetworkPortCommand) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (networkPortCommand)
	lengthInBits += m.NetworkPortCommand.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetPropertyStatesNetworkPortCommand) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetPropertyStatesNetworkPortCommandParse(readBuffer utils.ReadBuffer, peekedTagNumber uint8) (*BACnetPropertyStatesNetworkPortCommand, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesNetworkPortCommand"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesNetworkPortCommand")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (networkPortCommand)
	if pullErr := readBuffer.PullContext("networkPortCommand"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for networkPortCommand")
	}
	_networkPortCommand, _networkPortCommandErr := BACnetNetworkPortCommandTaggedParse(readBuffer, uint8(peekedTagNumber), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _networkPortCommandErr != nil {
		return nil, errors.Wrap(_networkPortCommandErr, "Error parsing 'networkPortCommand' field")
	}
	networkPortCommand := CastBACnetNetworkPortCommandTagged(_networkPortCommand)
	if closeErr := readBuffer.CloseContext("networkPortCommand"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for networkPortCommand")
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesNetworkPortCommand"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesNetworkPortCommand")
	}

	// Create a partially initialized instance
	_child := &BACnetPropertyStatesNetworkPortCommand{
		NetworkPortCommand:   CastBACnetNetworkPortCommandTagged(networkPortCommand),
		BACnetPropertyStates: &BACnetPropertyStates{},
	}
	_child.BACnetPropertyStates.Child = _child
	return _child, nil
}

func (m *BACnetPropertyStatesNetworkPortCommand) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesNetworkPortCommand"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesNetworkPortCommand")
		}

		// Simple Field (networkPortCommand)
		if pushErr := writeBuffer.PushContext("networkPortCommand"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for networkPortCommand")
		}
		_networkPortCommandErr := m.NetworkPortCommand.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("networkPortCommand"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for networkPortCommand")
		}
		if _networkPortCommandErr != nil {
			return errors.Wrap(_networkPortCommandErr, "Error serializing 'networkPortCommand' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesNetworkPortCommand"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesNetworkPortCommand")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetPropertyStatesNetworkPortCommand) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
