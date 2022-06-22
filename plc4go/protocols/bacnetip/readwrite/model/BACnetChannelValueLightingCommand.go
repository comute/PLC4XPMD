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

// BACnetChannelValueLightingCommand is the corresponding interface of BACnetChannelValueLightingCommand
type BACnetChannelValueLightingCommand interface {
	utils.LengthAware
	utils.Serializable
	BACnetChannelValue
	// GetLigthingCommandValue returns LigthingCommandValue (property field)
	GetLigthingCommandValue() BACnetLightingCommandEnclosed
}

// BACnetChannelValueLightingCommandExactly can be used when we want exactly this type and not a type which fulfills BACnetChannelValueLightingCommand.
// This is useful for switch cases.
type BACnetChannelValueLightingCommandExactly interface {
	BACnetChannelValueLightingCommand
	isBACnetChannelValueLightingCommand() bool
}

// _BACnetChannelValueLightingCommand is the data-structure of this message
type _BACnetChannelValueLightingCommand struct {
	*_BACnetChannelValue
	LigthingCommandValue BACnetLightingCommandEnclosed
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetChannelValueLightingCommand) InitializeParent(parent BACnetChannelValue, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetChannelValueLightingCommand) GetParent() BACnetChannelValue {
	return m._BACnetChannelValue
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetChannelValueLightingCommand) GetLigthingCommandValue() BACnetLightingCommandEnclosed {
	return m.LigthingCommandValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetChannelValueLightingCommand factory function for _BACnetChannelValueLightingCommand
func NewBACnetChannelValueLightingCommand(ligthingCommandValue BACnetLightingCommandEnclosed, peekedTagHeader BACnetTagHeader) *_BACnetChannelValueLightingCommand {
	_result := &_BACnetChannelValueLightingCommand{
		LigthingCommandValue: ligthingCommandValue,
		_BACnetChannelValue:  NewBACnetChannelValue(peekedTagHeader),
	}
	_result._BACnetChannelValue._BACnetChannelValueChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetChannelValueLightingCommand(structType interface{}) BACnetChannelValueLightingCommand {
	if casted, ok := structType.(BACnetChannelValueLightingCommand); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetChannelValueLightingCommand); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetChannelValueLightingCommand) GetTypeName() string {
	return "BACnetChannelValueLightingCommand"
}

func (m *_BACnetChannelValueLightingCommand) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetChannelValueLightingCommand) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (ligthingCommandValue)
	lengthInBits += m.LigthingCommandValue.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetChannelValueLightingCommand) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetChannelValueLightingCommandParse(readBuffer utils.ReadBuffer) (BACnetChannelValueLightingCommand, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetChannelValueLightingCommand"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetChannelValueLightingCommand")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (ligthingCommandValue)
	if pullErr := readBuffer.PullContext("ligthingCommandValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ligthingCommandValue")
	}
	_ligthingCommandValue, _ligthingCommandValueErr := BACnetLightingCommandEnclosedParse(readBuffer, uint8(uint8(0)))
	if _ligthingCommandValueErr != nil {
		return nil, errors.Wrap(_ligthingCommandValueErr, "Error parsing 'ligthingCommandValue' field")
	}
	ligthingCommandValue := _ligthingCommandValue.(BACnetLightingCommandEnclosed)
	if closeErr := readBuffer.CloseContext("ligthingCommandValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ligthingCommandValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetChannelValueLightingCommand"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetChannelValueLightingCommand")
	}

	// Create a partially initialized instance
	_child := &_BACnetChannelValueLightingCommand{
		LigthingCommandValue: ligthingCommandValue,
		_BACnetChannelValue:  &_BACnetChannelValue{},
	}
	_child._BACnetChannelValue._BACnetChannelValueChildRequirements = _child
	return _child, nil
}

func (m *_BACnetChannelValueLightingCommand) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetChannelValueLightingCommand"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetChannelValueLightingCommand")
		}

		// Simple Field (ligthingCommandValue)
		if pushErr := writeBuffer.PushContext("ligthingCommandValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ligthingCommandValue")
		}
		_ligthingCommandValueErr := writeBuffer.WriteSerializable(m.GetLigthingCommandValue())
		if popErr := writeBuffer.PopContext("ligthingCommandValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ligthingCommandValue")
		}
		if _ligthingCommandValueErr != nil {
			return errors.Wrap(_ligthingCommandValueErr, "Error serializing 'ligthingCommandValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetChannelValueLightingCommand"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetChannelValueLightingCommand")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetChannelValueLightingCommand) isBACnetChannelValueLightingCommand() bool {
	return true
}

func (m *_BACnetChannelValueLightingCommand) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
