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
type BACnetConfirmedServiceACKAtomicReadFile struct {
	*BACnetConfirmedServiceACK
}

// The corresponding interface
type IBACnetConfirmedServiceACKAtomicReadFile interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BACnetConfirmedServiceACKAtomicReadFile) ServiceChoice() uint8 {
	return 0x06
}

func (m *BACnetConfirmedServiceACKAtomicReadFile) InitializeParent(parent *BACnetConfirmedServiceACK) {
}

func NewBACnetConfirmedServiceACKAtomicReadFile() *BACnetConfirmedServiceACK {
	child := &BACnetConfirmedServiceACKAtomicReadFile{
		BACnetConfirmedServiceACK: NewBACnetConfirmedServiceACK(),
	}
	child.Child = child
	return child.BACnetConfirmedServiceACK
}

func CastBACnetConfirmedServiceACKAtomicReadFile(structType interface{}) *BACnetConfirmedServiceACKAtomicReadFile {
	castFunc := func(typ interface{}) *BACnetConfirmedServiceACKAtomicReadFile {
		if casted, ok := typ.(BACnetConfirmedServiceACKAtomicReadFile); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetConfirmedServiceACKAtomicReadFile); ok {
			return casted
		}
		if casted, ok := typ.(BACnetConfirmedServiceACK); ok {
			return CastBACnetConfirmedServiceACKAtomicReadFile(casted.Child)
		}
		if casted, ok := typ.(*BACnetConfirmedServiceACK); ok {
			return CastBACnetConfirmedServiceACKAtomicReadFile(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetConfirmedServiceACKAtomicReadFile) GetTypeName() string {
	return "BACnetConfirmedServiceACKAtomicReadFile"
}

func (m *BACnetConfirmedServiceACKAtomicReadFile) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetConfirmedServiceACKAtomicReadFile) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.ParentLengthInBits())

	return lengthInBits
}

func (m *BACnetConfirmedServiceACKAtomicReadFile) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetConfirmedServiceACKAtomicReadFileParse(readBuffer utils.ReadBuffer) (*BACnetConfirmedServiceACK, error) {
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceACKAtomicReadFile"); pullErr != nil {
		return nil, pullErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceACKAtomicReadFile"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConfirmedServiceACKAtomicReadFile{
		BACnetConfirmedServiceACK: &BACnetConfirmedServiceACK{},
	}
	_child.BACnetConfirmedServiceACK.Child = _child
	return _child.BACnetConfirmedServiceACK, nil
}

func (m *BACnetConfirmedServiceACKAtomicReadFile) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceACKAtomicReadFile"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceACKAtomicReadFile"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConfirmedServiceACKAtomicReadFile) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
