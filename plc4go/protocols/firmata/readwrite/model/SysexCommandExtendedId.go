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

// SysexCommandExtendedId is the corresponding interface of SysexCommandExtendedId
type SysexCommandExtendedId interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	SysexCommand
	// GetId returns Id (property field)
	GetId() []int8
}

// SysexCommandExtendedIdExactly can be used when we want exactly this type and not a type which fulfills SysexCommandExtendedId.
// This is useful for switch cases.
type SysexCommandExtendedIdExactly interface {
	SysexCommandExtendedId
	isSysexCommandExtendedId() bool
}

// _SysexCommandExtendedId is the data-structure of this message
type _SysexCommandExtendedId struct {
	*_SysexCommand
	Id []int8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SysexCommandExtendedId) GetCommandType() uint8 {
	return 0x00
}

func (m *_SysexCommandExtendedId) GetResponse() bool {
	return false
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SysexCommandExtendedId) InitializeParent(parent SysexCommand) {}

func (m *_SysexCommandExtendedId) GetParent() SysexCommand {
	return m._SysexCommand
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SysexCommandExtendedId) GetId() []int8 {
	return m.Id
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewSysexCommandExtendedId factory function for _SysexCommandExtendedId
func NewSysexCommandExtendedId(id []int8) *_SysexCommandExtendedId {
	_result := &_SysexCommandExtendedId{
		Id:            id,
		_SysexCommand: NewSysexCommand(),
	}
	_result._SysexCommand._SysexCommandChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSysexCommandExtendedId(structType any) SysexCommandExtendedId {
	if casted, ok := structType.(SysexCommandExtendedId); ok {
		return casted
	}
	if casted, ok := structType.(*SysexCommandExtendedId); ok {
		return *casted
	}
	return nil
}

func (m *_SysexCommandExtendedId) GetTypeName() string {
	return "SysexCommandExtendedId"
}

func (m *_SysexCommandExtendedId) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Array field
	if len(m.Id) > 0 {
		lengthInBits += 8 * uint16(len(m.Id))
	}

	return lengthInBits
}

func (m *_SysexCommandExtendedId) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func SysexCommandExtendedIdParse(ctx context.Context, theBytes []byte, response bool) (SysexCommandExtendedId, error) {
	return SysexCommandExtendedIdParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), response)
}

func SysexCommandExtendedIdParseWithBufferProducer(response bool) func(ctx context.Context, readBuffer utils.ReadBuffer) (SysexCommandExtendedId, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (SysexCommandExtendedId, error) {
		return SysexCommandExtendedIdParseWithBuffer(ctx, readBuffer, response)
	}
}

func SysexCommandExtendedIdParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, response bool) (SysexCommandExtendedId, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SysexCommandExtendedId"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SysexCommandExtendedId")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	id, err := ReadCountArrayField[int8](ctx, "id", ReadSignedByte(readBuffer, uint8(8)), uint64(int32(2)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'id' field"))
	}

	if closeErr := readBuffer.CloseContext("SysexCommandExtendedId"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SysexCommandExtendedId")
	}

	// Create a partially initialized instance
	_child := &_SysexCommandExtendedId{
		_SysexCommand: &_SysexCommand{},
		Id:            id,
	}
	_child._SysexCommand._SysexCommandChildRequirements = _child
	return _child, nil
}

func (m *_SysexCommandExtendedId) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SysexCommandExtendedId) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SysexCommandExtendedId"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SysexCommandExtendedId")
		}

		if err := WriteSimpleTypeArrayField(ctx, "id", m.GetId(), WriteSignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'id' field")
		}

		if popErr := writeBuffer.PopContext("SysexCommandExtendedId"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SysexCommandExtendedId")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SysexCommandExtendedId) isSysexCommandExtendedId() bool {
	return true
}

func (m *_SysexCommandExtendedId) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
