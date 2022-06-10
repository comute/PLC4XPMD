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

// SysexCommandCapabilityQuery is the data-structure of this message
type SysexCommandCapabilityQuery struct {
	*SysexCommand
}

// ISysexCommandCapabilityQuery is the corresponding interface of SysexCommandCapabilityQuery
type ISysexCommandCapabilityQuery interface {
	ISysexCommand
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

func (m *SysexCommandCapabilityQuery) GetCommandType() uint8 {
	return 0x6B
}

func (m *SysexCommandCapabilityQuery) GetResponse() bool {
	return false
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *SysexCommandCapabilityQuery) InitializeParent(parent *SysexCommand) {}

func (m *SysexCommandCapabilityQuery) GetParent() *SysexCommand {
	return m.SysexCommand
}

// NewSysexCommandCapabilityQuery factory function for SysexCommandCapabilityQuery
func NewSysexCommandCapabilityQuery() *SysexCommandCapabilityQuery {
	_result := &SysexCommandCapabilityQuery{
		SysexCommand: NewSysexCommand(),
	}
	_result.Child = _result
	return _result
}

func CastSysexCommandCapabilityQuery(structType interface{}) *SysexCommandCapabilityQuery {
	if casted, ok := structType.(SysexCommandCapabilityQuery); ok {
		return &casted
	}
	if casted, ok := structType.(*SysexCommandCapabilityQuery); ok {
		return casted
	}
	if casted, ok := structType.(SysexCommand); ok {
		return CastSysexCommandCapabilityQuery(casted.Child)
	}
	if casted, ok := structType.(*SysexCommand); ok {
		return CastSysexCommandCapabilityQuery(casted.Child)
	}
	return nil
}

func (m *SysexCommandCapabilityQuery) GetTypeName() string {
	return "SysexCommandCapabilityQuery"
}

func (m *SysexCommandCapabilityQuery) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *SysexCommandCapabilityQuery) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *SysexCommandCapabilityQuery) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func SysexCommandCapabilityQueryParse(readBuffer utils.ReadBuffer, response bool) (*SysexCommandCapabilityQuery, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SysexCommandCapabilityQuery"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SysexCommandCapabilityQuery")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("SysexCommandCapabilityQuery"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SysexCommandCapabilityQuery")
	}

	// Create a partially initialized instance
	_child := &SysexCommandCapabilityQuery{
		SysexCommand: &SysexCommand{},
	}
	_child.SysexCommand.Child = _child
	return _child, nil
}

func (m *SysexCommandCapabilityQuery) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SysexCommandCapabilityQuery"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SysexCommandCapabilityQuery")
		}

		if popErr := writeBuffer.PopContext("SysexCommandCapabilityQuery"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SysexCommandCapabilityQuery")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *SysexCommandCapabilityQuery) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
