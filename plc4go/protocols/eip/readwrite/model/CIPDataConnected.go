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

// CIPDataConnected is the corresponding interface of CIPDataConnected
type CIPDataConnected interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetValue returns Value (property field)
	GetValue() uint32
	// GetTagStatus returns TagStatus (property field)
	GetTagStatus() uint16
	// IsCIPDataConnected is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsCIPDataConnected()
}

// _CIPDataConnected is the data-structure of this message
type _CIPDataConnected struct {
	Value     uint32
	TagStatus uint16
}

var _ CIPDataConnected = (*_CIPDataConnected)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CIPDataConnected) GetValue() uint32 {
	return m.Value
}

func (m *_CIPDataConnected) GetTagStatus() uint16 {
	return m.TagStatus
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCIPDataConnected factory function for _CIPDataConnected
func NewCIPDataConnected(value uint32, tagStatus uint16) *_CIPDataConnected {
	return &_CIPDataConnected{Value: value, TagStatus: tagStatus}
}

// Deprecated: use the interface for direct cast
func CastCIPDataConnected(structType any) CIPDataConnected {
	if casted, ok := structType.(CIPDataConnected); ok {
		return casted
	}
	if casted, ok := structType.(*CIPDataConnected); ok {
		return *casted
	}
	return nil
}

func (m *_CIPDataConnected) GetTypeName() string {
	return "CIPDataConnected"
}

func (m *_CIPDataConnected) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (value)
	lengthInBits += 32

	// Simple field (tagStatus)
	lengthInBits += 16

	return lengthInBits
}

func (m *_CIPDataConnected) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func CIPDataConnectedParse(ctx context.Context, theBytes []byte) (CIPDataConnected, error) {
	return CIPDataConnectedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func CIPDataConnectedParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (CIPDataConnected, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (CIPDataConnected, error) {
		return CIPDataConnectedParseWithBuffer(ctx, readBuffer)
	}
}

func CIPDataConnectedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (CIPDataConnected, error) {
	v, err := (&_CIPDataConnected{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *_CIPDataConnected) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__cIPDataConnected CIPDataConnected, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CIPDataConnected"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CIPDataConnected")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	value, err := ReadSimpleField(ctx, "value", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	tagStatus, err := ReadSimpleField(ctx, "tagStatus", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'tagStatus' field"))
	}
	m.TagStatus = tagStatus

	if closeErr := readBuffer.CloseContext("CIPDataConnected"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CIPDataConnected")
	}

	return m, nil
}

func (m *_CIPDataConnected) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CIPDataConnected) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("CIPDataConnected"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for CIPDataConnected")
	}

	if err := WriteSimpleField[uint32](ctx, "value", m.GetValue(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
		return errors.Wrap(err, "Error serializing 'value' field")
	}

	if err := WriteSimpleField[uint16](ctx, "tagStatus", m.GetTagStatus(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
		return errors.Wrap(err, "Error serializing 'tagStatus' field")
	}

	if popErr := writeBuffer.PopContext("CIPDataConnected"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for CIPDataConnected")
	}
	return nil
}

func (m *_CIPDataConnected) IsCIPDataConnected() {}

func (m *_CIPDataConnected) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
