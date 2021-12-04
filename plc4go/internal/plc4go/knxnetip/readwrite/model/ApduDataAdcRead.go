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
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type ApduDataAdcRead struct {
	*ApduData
}

// The corresponding interface
type IApduDataAdcRead interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *ApduDataAdcRead) ApciType() uint8 {
	return 0x6
}

func (m *ApduDataAdcRead) InitializeParent(parent *ApduData) {
}

func NewApduDataAdcRead() *ApduData {
	child := &ApduDataAdcRead{
		ApduData: NewApduData(),
	}
	child.Child = child
	return child.ApduData
}

func CastApduDataAdcRead(structType interface{}) *ApduDataAdcRead {
	castFunc := func(typ interface{}) *ApduDataAdcRead {
		if casted, ok := typ.(ApduDataAdcRead); ok {
			return &casted
		}
		if casted, ok := typ.(*ApduDataAdcRead); ok {
			return casted
		}
		if casted, ok := typ.(ApduData); ok {
			return CastApduDataAdcRead(casted.Child)
		}
		if casted, ok := typ.(*ApduData); ok {
			return CastApduDataAdcRead(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *ApduDataAdcRead) GetTypeName() string {
	return "ApduDataAdcRead"
}

func (m *ApduDataAdcRead) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *ApduDataAdcRead) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.ParentLengthInBits())

	return lengthInBits
}

func (m *ApduDataAdcRead) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func ApduDataAdcReadParse(readBuffer utils.ReadBuffer, dataLength uint8) (*ApduData, error) {
	if pullErr := readBuffer.PullContext("ApduDataAdcRead"); pullErr != nil {
		return nil, pullErr
	}

	if closeErr := readBuffer.CloseContext("ApduDataAdcRead"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &ApduDataAdcRead{
		ApduData: &ApduData{},
	}
	_child.ApduData.Child = _child
	return _child.ApduData, nil
}

func (m *ApduDataAdcRead) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataAdcRead"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("ApduDataAdcRead"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *ApduDataAdcRead) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
