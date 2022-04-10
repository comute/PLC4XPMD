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
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type CBusCommandPointToPoint struct {
	*CBusCommand
	Command *CBusPointToPointCommand

	// Arguments.
	Srchk bool
}

// The corresponding interface
type ICBusCommandPointToPoint interface {
	ICBusCommand
	// GetCommand returns Command (property field)
	GetCommand() *CBusPointToPointCommand
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

func (m *CBusCommandPointToPoint) InitializeParent(parent *CBusCommand, header *CBusHeader) {
	m.CBusCommand.Header = header
}

func (m *CBusCommandPointToPoint) GetParent() *CBusCommand {
	return m.CBusCommand
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////
func (m *CBusCommandPointToPoint) GetCommand() *CBusPointToPointCommand {
	return m.Command
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCBusCommandPointToPoint factory function for CBusCommandPointToPoint
func NewCBusCommandPointToPoint(command *CBusPointToPointCommand, header *CBusHeader, srchk bool) *CBusCommandPointToPoint {
	_result := &CBusCommandPointToPoint{
		Command:     command,
		CBusCommand: NewCBusCommand(header, srchk),
	}
	_result.Child = _result
	return _result
}

func CastCBusCommandPointToPoint(structType interface{}) *CBusCommandPointToPoint {
	if casted, ok := structType.(CBusCommandPointToPoint); ok {
		return &casted
	}
	if casted, ok := structType.(*CBusCommandPointToPoint); ok {
		return casted
	}
	if casted, ok := structType.(CBusCommand); ok {
		return CastCBusCommandPointToPoint(casted.Child)
	}
	if casted, ok := structType.(*CBusCommand); ok {
		return CastCBusCommandPointToPoint(casted.Child)
	}
	return nil
}

func (m *CBusCommandPointToPoint) GetTypeName() string {
	return "CBusCommandPointToPoint"
}

func (m *CBusCommandPointToPoint) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *CBusCommandPointToPoint) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (command)
	lengthInBits += m.Command.GetLengthInBits()

	return lengthInBits
}

func (m *CBusCommandPointToPoint) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func CBusCommandPointToPointParse(readBuffer utils.ReadBuffer, srchk bool) (*CBusCommandPointToPoint, error) {
	if pullErr := readBuffer.PullContext("CBusCommandPointToPoint"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (command)
	if pullErr := readBuffer.PullContext("command"); pullErr != nil {
		return nil, pullErr
	}
	_command, _commandErr := CBusPointToPointCommandParse(readBuffer, bool(srchk))
	if _commandErr != nil {
		return nil, errors.Wrap(_commandErr, "Error parsing 'command' field")
	}
	command := CastCBusPointToPointCommand(_command)
	if closeErr := readBuffer.CloseContext("command"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("CBusCommandPointToPoint"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &CBusCommandPointToPoint{
		Command:     CastCBusPointToPointCommand(command),
		CBusCommand: &CBusCommand{},
	}
	_child.CBusCommand.Child = _child
	return _child, nil
}

func (m *CBusCommandPointToPoint) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CBusCommandPointToPoint"); pushErr != nil {
			return pushErr
		}

		// Simple Field (command)
		if pushErr := writeBuffer.PushContext("command"); pushErr != nil {
			return pushErr
		}
		_commandErr := m.Command.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("command"); popErr != nil {
			return popErr
		}
		if _commandErr != nil {
			return errors.Wrap(_commandErr, "Error serializing 'command' field")
		}

		if popErr := writeBuffer.PopContext("CBusCommandPointToPoint"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *CBusCommandPointToPoint) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
