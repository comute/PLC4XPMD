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
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// LevelInformationAbsent is the corresponding interface of LevelInformationAbsent
type LevelInformationAbsent interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	LevelInformation
	// IsLevelInformationAbsent is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsLevelInformationAbsent()
}

// _LevelInformationAbsent is the data-structure of this message
type _LevelInformationAbsent struct {
	LevelInformationContract
	// Reserved Fields
	reservedField0 *uint16
}

var _ LevelInformationAbsent = (*_LevelInformationAbsent)(nil)
var _ LevelInformationRequirements = (*_LevelInformationAbsent)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_LevelInformationAbsent) GetParent() LevelInformationContract {
	return m.LevelInformationContract
}

// NewLevelInformationAbsent factory function for _LevelInformationAbsent
func NewLevelInformationAbsent(raw uint16) *_LevelInformationAbsent {
	_result := &_LevelInformationAbsent{
		LevelInformationContract: NewLevelInformation(raw),
	}
	_result.LevelInformationContract.(*_LevelInformation)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastLevelInformationAbsent(structType any) LevelInformationAbsent {
	if casted, ok := structType.(LevelInformationAbsent); ok {
		return casted
	}
	if casted, ok := structType.(*LevelInformationAbsent); ok {
		return *casted
	}
	return nil
}

func (m *_LevelInformationAbsent) GetTypeName() string {
	return "LevelInformationAbsent"
}

func (m *_LevelInformationAbsent) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.LevelInformationContract.(*_LevelInformation).getLengthInBits(ctx))

	// Reserved Field (reserved)
	lengthInBits += 16

	return lengthInBits
}

func (m *_LevelInformationAbsent) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_LevelInformationAbsent) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_LevelInformation) (__levelInformationAbsent LevelInformationAbsent, err error) {
	m.LevelInformationContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("LevelInformationAbsent"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for LevelInformationAbsent")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedShort(readBuffer, uint8(16)), uint16(0x0000))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	if closeErr := readBuffer.CloseContext("LevelInformationAbsent"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for LevelInformationAbsent")
	}

	return m, nil
}

func (m *_LevelInformationAbsent) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_LevelInformationAbsent) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("LevelInformationAbsent"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for LevelInformationAbsent")
		}

		if err := WriteReservedField[uint16](ctx, "reserved", uint16(0x0000), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}

		if popErr := writeBuffer.PopContext("LevelInformationAbsent"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for LevelInformationAbsent")
		}
		return nil
	}
	return m.LevelInformationContract.(*_LevelInformation).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_LevelInformationAbsent) IsLevelInformationAbsent() {}

func (m *_LevelInformationAbsent) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
