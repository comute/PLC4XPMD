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
type Dummy struct {
	Dummy uint16
}

// The corresponding interface
type IDummy interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

func NewDummy(dummy uint16) *Dummy {
	return &Dummy{Dummy: dummy}
}

func CastDummy(structType interface{}) *Dummy {
	castFunc := func(typ interface{}) *Dummy {
		if casted, ok := typ.(Dummy); ok {
			return &casted
		}
		if casted, ok := typ.(*Dummy); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *Dummy) GetTypeName() string {
	return "Dummy"
}

func (m *Dummy) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *Dummy) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (dummy)
	lengthInBits += 16

	return lengthInBits
}

func (m *Dummy) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func DummyParse(readBuffer utils.ReadBuffer) (*Dummy, error) {
	if pullErr := readBuffer.PullContext("Dummy"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (dummy)
	_dummy, _dummyErr := readBuffer.ReadUint16("dummy", 16)
	if _dummyErr != nil {
		return nil, errors.Wrap(_dummyErr, "Error parsing 'dummy' field")
	}
	dummy := _dummy

	if closeErr := readBuffer.CloseContext("Dummy"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewDummy(dummy), nil
}

func (m *Dummy) Serialize(writeBuffer utils.WriteBuffer) error {
	if pushErr := writeBuffer.PushContext("Dummy"); pushErr != nil {
		return pushErr
	}

	// Simple Field (dummy)
	dummy := uint16(m.Dummy)
	_dummyErr := writeBuffer.WriteUint16("dummy", 16, (dummy))
	if _dummyErr != nil {
		return errors.Wrap(_dummyErr, "Error serializing 'dummy' field")
	}

	if popErr := writeBuffer.PopContext("Dummy"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *Dummy) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
