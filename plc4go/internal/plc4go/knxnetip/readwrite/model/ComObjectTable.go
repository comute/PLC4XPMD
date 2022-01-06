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
type ComObjectTable struct {
	Child IComObjectTableChild
}

// The corresponding interface
type IComObjectTable interface {
	FirmwareType() FirmwareType
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

type IComObjectTableParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child IComObjectTable, serializeChildFunction func() error) error
	GetTypeName() string
}

type IComObjectTableChild interface {
	Serialize(writeBuffer utils.WriteBuffer) error
	InitializeParent(parent *ComObjectTable)
	GetTypeName() string
	IComObjectTable
}

func NewComObjectTable() *ComObjectTable {
	return &ComObjectTable{}
}

func CastComObjectTable(structType interface{}) *ComObjectTable {
	castFunc := func(typ interface{}) *ComObjectTable {
		if casted, ok := typ.(ComObjectTable); ok {
			return &casted
		}
		if casted, ok := typ.(*ComObjectTable); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *ComObjectTable) GetTypeName() string {
	return "ComObjectTable"
}

func (m *ComObjectTable) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *ComObjectTable) LengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.LengthInBits()
}

func (m *ComObjectTable) ParentLengthInBits() uint16 {
	lengthInBits := uint16(0)

	return lengthInBits
}

func (m *ComObjectTable) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func ComObjectTableParse(readBuffer utils.ReadBuffer, firmwareType FirmwareType) (*ComObjectTable, error) {
	if pullErr := readBuffer.PullContext("ComObjectTable"); pullErr != nil {
		return nil, pullErr
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _parent *ComObjectTable
	var typeSwitchError error
	switch {
	case firmwareType == FirmwareType_SYSTEM_1: // ComObjectTableRealisationType1
		_parent, typeSwitchError = ComObjectTableRealisationType1Parse(readBuffer, firmwareType)
	case firmwareType == FirmwareType_SYSTEM_2: // ComObjectTableRealisationType2
		_parent, typeSwitchError = ComObjectTableRealisationType2Parse(readBuffer, firmwareType)
	case firmwareType == FirmwareType_SYSTEM_300: // ComObjectTableRealisationType6
		_parent, typeSwitchError = ComObjectTableRealisationType6Parse(readBuffer, firmwareType)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	if closeErr := readBuffer.CloseContext("ComObjectTable"); closeErr != nil {
		return nil, closeErr
	}

	// Finish initializing
	_parent.Child.InitializeParent(_parent)
	return _parent, nil
}

func (m *ComObjectTable) Serialize(writeBuffer utils.WriteBuffer) error {
	return m.Child.Serialize(writeBuffer)
}

func (m *ComObjectTable) SerializeParent(writeBuffer utils.WriteBuffer, child IComObjectTable, serializeChildFunction func() error) error {
	if pushErr := writeBuffer.PushContext("ComObjectTable"); pushErr != nil {
		return pushErr
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("ComObjectTable"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *ComObjectTable) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
