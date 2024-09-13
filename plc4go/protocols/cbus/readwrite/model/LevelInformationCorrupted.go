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

// LevelInformationCorrupted is the corresponding interface of LevelInformationCorrupted
type LevelInformationCorrupted interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	LevelInformation
	// GetCorruptedNibble1 returns CorruptedNibble1 (property field)
	GetCorruptedNibble1() uint8
	// GetCorruptedNibble2 returns CorruptedNibble2 (property field)
	GetCorruptedNibble2() uint8
	// GetCorruptedNibble3 returns CorruptedNibble3 (property field)
	GetCorruptedNibble3() uint8
	// GetCorruptedNibble4 returns CorruptedNibble4 (property field)
	GetCorruptedNibble4() uint8
	// IsLevelInformationCorrupted is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsLevelInformationCorrupted()
}

// _LevelInformationCorrupted is the data-structure of this message
type _LevelInformationCorrupted struct {
	LevelInformationContract
	CorruptedNibble1 uint8
	CorruptedNibble2 uint8
	CorruptedNibble3 uint8
	CorruptedNibble4 uint8
}

var _ LevelInformationCorrupted = (*_LevelInformationCorrupted)(nil)
var _ LevelInformationRequirements = (*_LevelInformationCorrupted)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_LevelInformationCorrupted) GetParent() LevelInformationContract {
	return m.LevelInformationContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_LevelInformationCorrupted) GetCorruptedNibble1() uint8 {
	return m.CorruptedNibble1
}

func (m *_LevelInformationCorrupted) GetCorruptedNibble2() uint8 {
	return m.CorruptedNibble2
}

func (m *_LevelInformationCorrupted) GetCorruptedNibble3() uint8 {
	return m.CorruptedNibble3
}

func (m *_LevelInformationCorrupted) GetCorruptedNibble4() uint8 {
	return m.CorruptedNibble4
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewLevelInformationCorrupted factory function for _LevelInformationCorrupted
func NewLevelInformationCorrupted(corruptedNibble1 uint8, corruptedNibble2 uint8, corruptedNibble3 uint8, corruptedNibble4 uint8, raw uint16) *_LevelInformationCorrupted {
	_result := &_LevelInformationCorrupted{
		LevelInformationContract: NewLevelInformation(raw),
		CorruptedNibble1:         corruptedNibble1,
		CorruptedNibble2:         corruptedNibble2,
		CorruptedNibble3:         corruptedNibble3,
		CorruptedNibble4:         corruptedNibble4,
	}
	_result.LevelInformationContract.(*_LevelInformation)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastLevelInformationCorrupted(structType any) LevelInformationCorrupted {
	if casted, ok := structType.(LevelInformationCorrupted); ok {
		return casted
	}
	if casted, ok := structType.(*LevelInformationCorrupted); ok {
		return *casted
	}
	return nil
}

func (m *_LevelInformationCorrupted) GetTypeName() string {
	return "LevelInformationCorrupted"
}

func (m *_LevelInformationCorrupted) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.LevelInformationContract.(*_LevelInformation).getLengthInBits(ctx))

	// Simple field (corruptedNibble1)
	lengthInBits += 4

	// Simple field (corruptedNibble2)
	lengthInBits += 4

	// Simple field (corruptedNibble3)
	lengthInBits += 4

	// Simple field (corruptedNibble4)
	lengthInBits += 4

	return lengthInBits
}

func (m *_LevelInformationCorrupted) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_LevelInformationCorrupted) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_LevelInformation) (__levelInformationCorrupted LevelInformationCorrupted, err error) {
	m.LevelInformationContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("LevelInformationCorrupted"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for LevelInformationCorrupted")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	corruptedNibble1, err := ReadSimpleField(ctx, "corruptedNibble1", ReadUnsignedByte(readBuffer, uint8(4)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'corruptedNibble1' field"))
	}
	m.CorruptedNibble1 = corruptedNibble1

	corruptedNibble2, err := ReadSimpleField(ctx, "corruptedNibble2", ReadUnsignedByte(readBuffer, uint8(4)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'corruptedNibble2' field"))
	}
	m.CorruptedNibble2 = corruptedNibble2

	corruptedNibble3, err := ReadSimpleField(ctx, "corruptedNibble3", ReadUnsignedByte(readBuffer, uint8(4)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'corruptedNibble3' field"))
	}
	m.CorruptedNibble3 = corruptedNibble3

	corruptedNibble4, err := ReadSimpleField(ctx, "corruptedNibble4", ReadUnsignedByte(readBuffer, uint8(4)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'corruptedNibble4' field"))
	}
	m.CorruptedNibble4 = corruptedNibble4

	if closeErr := readBuffer.CloseContext("LevelInformationCorrupted"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for LevelInformationCorrupted")
	}

	return m, nil
}

func (m *_LevelInformationCorrupted) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_LevelInformationCorrupted) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("LevelInformationCorrupted"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for LevelInformationCorrupted")
		}

		if err := WriteSimpleField[uint8](ctx, "corruptedNibble1", m.GetCorruptedNibble1(), WriteUnsignedByte(writeBuffer, 4)); err != nil {
			return errors.Wrap(err, "Error serializing 'corruptedNibble1' field")
		}

		if err := WriteSimpleField[uint8](ctx, "corruptedNibble2", m.GetCorruptedNibble2(), WriteUnsignedByte(writeBuffer, 4)); err != nil {
			return errors.Wrap(err, "Error serializing 'corruptedNibble2' field")
		}

		if err := WriteSimpleField[uint8](ctx, "corruptedNibble3", m.GetCorruptedNibble3(), WriteUnsignedByte(writeBuffer, 4)); err != nil {
			return errors.Wrap(err, "Error serializing 'corruptedNibble3' field")
		}

		if err := WriteSimpleField[uint8](ctx, "corruptedNibble4", m.GetCorruptedNibble4(), WriteUnsignedByte(writeBuffer, 4)); err != nil {
			return errors.Wrap(err, "Error serializing 'corruptedNibble4' field")
		}

		if popErr := writeBuffer.PopContext("LevelInformationCorrupted"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for LevelInformationCorrupted")
		}
		return nil
	}
	return m.LevelInformationContract.(*_LevelInformation).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_LevelInformationCorrupted) IsLevelInformationCorrupted() {}

func (m *_LevelInformationCorrupted) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
