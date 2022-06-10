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

// COTPParameterCallingTsap is the data-structure of this message
type COTPParameterCallingTsap struct {
	*COTPParameter
	TsapId uint16

	// Arguments.
	Rest uint8
}

// ICOTPParameterCallingTsap is the corresponding interface of COTPParameterCallingTsap
type ICOTPParameterCallingTsap interface {
	ICOTPParameter
	// GetTsapId returns TsapId (property field)
	GetTsapId() uint16
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

func (m *COTPParameterCallingTsap) GetParameterType() uint8 {
	return 0xC1
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *COTPParameterCallingTsap) InitializeParent(parent *COTPParameter) {}

func (m *COTPParameterCallingTsap) GetParent() *COTPParameter {
	return m.COTPParameter
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *COTPParameterCallingTsap) GetTsapId() uint16 {
	return m.TsapId
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCOTPParameterCallingTsap factory function for COTPParameterCallingTsap
func NewCOTPParameterCallingTsap(tsapId uint16, rest uint8) *COTPParameterCallingTsap {
	_result := &COTPParameterCallingTsap{
		TsapId:        tsapId,
		COTPParameter: NewCOTPParameter(rest),
	}
	_result.Child = _result
	return _result
}

func CastCOTPParameterCallingTsap(structType interface{}) *COTPParameterCallingTsap {
	if casted, ok := structType.(COTPParameterCallingTsap); ok {
		return &casted
	}
	if casted, ok := structType.(*COTPParameterCallingTsap); ok {
		return casted
	}
	if casted, ok := structType.(COTPParameter); ok {
		return CastCOTPParameterCallingTsap(casted.Child)
	}
	if casted, ok := structType.(*COTPParameter); ok {
		return CastCOTPParameterCallingTsap(casted.Child)
	}
	return nil
}

func (m *COTPParameterCallingTsap) GetTypeName() string {
	return "COTPParameterCallingTsap"
}

func (m *COTPParameterCallingTsap) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *COTPParameterCallingTsap) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (tsapId)
	lengthInBits += 16

	return lengthInBits
}

func (m *COTPParameterCallingTsap) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func COTPParameterCallingTsapParse(readBuffer utils.ReadBuffer, rest uint8) (*COTPParameterCallingTsap, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("COTPParameterCallingTsap"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for COTPParameterCallingTsap")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (tsapId)
	_tsapId, _tsapIdErr := readBuffer.ReadUint16("tsapId", 16)
	if _tsapIdErr != nil {
		return nil, errors.Wrap(_tsapIdErr, "Error parsing 'tsapId' field")
	}
	tsapId := _tsapId

	if closeErr := readBuffer.CloseContext("COTPParameterCallingTsap"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for COTPParameterCallingTsap")
	}

	// Create a partially initialized instance
	_child := &COTPParameterCallingTsap{
		TsapId:        tsapId,
		COTPParameter: &COTPParameter{},
	}
	_child.COTPParameter.Child = _child
	return _child, nil
}

func (m *COTPParameterCallingTsap) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("COTPParameterCallingTsap"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for COTPParameterCallingTsap")
		}

		// Simple Field (tsapId)
		tsapId := uint16(m.TsapId)
		_tsapIdErr := writeBuffer.WriteUint16("tsapId", 16, (tsapId))
		if _tsapIdErr != nil {
			return errors.Wrap(_tsapIdErr, "Error serializing 'tsapId' field")
		}

		if popErr := writeBuffer.PopContext("COTPParameterCallingTsap"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for COTPParameterCallingTsap")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *COTPParameterCallingTsap) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
