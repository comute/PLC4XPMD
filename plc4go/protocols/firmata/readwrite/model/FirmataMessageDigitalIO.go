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
	"encoding/binary"
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	"github.com/apache/plc4x/plc4go/spi/codegen"
	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// FirmataMessageDigitalIO is the corresponding interface of FirmataMessageDigitalIO
type FirmataMessageDigitalIO interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	FirmataMessage
	// GetPinBlock returns PinBlock (property field)
	GetPinBlock() uint8
	// GetData returns Data (property field)
	GetData() []int8
	// IsFirmataMessageDigitalIO is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsFirmataMessageDigitalIO()
	// CreateBuilder creates a FirmataMessageDigitalIOBuilder
	CreateFirmataMessageDigitalIOBuilder() FirmataMessageDigitalIOBuilder
}

// _FirmataMessageDigitalIO is the data-structure of this message
type _FirmataMessageDigitalIO struct {
	FirmataMessageContract
	PinBlock uint8
	Data     []int8
}

var _ FirmataMessageDigitalIO = (*_FirmataMessageDigitalIO)(nil)
var _ FirmataMessageRequirements = (*_FirmataMessageDigitalIO)(nil)

// NewFirmataMessageDigitalIO factory function for _FirmataMessageDigitalIO
func NewFirmataMessageDigitalIO(pinBlock uint8, data []int8, response bool) *_FirmataMessageDigitalIO {
	_result := &_FirmataMessageDigitalIO{
		FirmataMessageContract: NewFirmataMessage(response),
		PinBlock:               pinBlock,
		Data:                   data,
	}
	_result.FirmataMessageContract.(*_FirmataMessage)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// FirmataMessageDigitalIOBuilder is a builder for FirmataMessageDigitalIO
type FirmataMessageDigitalIOBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(pinBlock uint8, data []int8) FirmataMessageDigitalIOBuilder
	// WithPinBlock adds PinBlock (property field)
	WithPinBlock(uint8) FirmataMessageDigitalIOBuilder
	// WithData adds Data (property field)
	WithData(...int8) FirmataMessageDigitalIOBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() FirmataMessageBuilder
	// Build builds the FirmataMessageDigitalIO or returns an error if something is wrong
	Build() (FirmataMessageDigitalIO, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() FirmataMessageDigitalIO
}

// NewFirmataMessageDigitalIOBuilder() creates a FirmataMessageDigitalIOBuilder
func NewFirmataMessageDigitalIOBuilder() FirmataMessageDigitalIOBuilder {
	return &_FirmataMessageDigitalIOBuilder{_FirmataMessageDigitalIO: new(_FirmataMessageDigitalIO)}
}

type _FirmataMessageDigitalIOBuilder struct {
	*_FirmataMessageDigitalIO

	parentBuilder *_FirmataMessageBuilder

	err *utils.MultiError
}

var _ (FirmataMessageDigitalIOBuilder) = (*_FirmataMessageDigitalIOBuilder)(nil)

func (b *_FirmataMessageDigitalIOBuilder) setParent(contract FirmataMessageContract) {
	b.FirmataMessageContract = contract
	contract.(*_FirmataMessage)._SubType = b._FirmataMessageDigitalIO
}

func (b *_FirmataMessageDigitalIOBuilder) WithMandatoryFields(pinBlock uint8, data []int8) FirmataMessageDigitalIOBuilder {
	return b.WithPinBlock(pinBlock).WithData(data...)
}

func (b *_FirmataMessageDigitalIOBuilder) WithPinBlock(pinBlock uint8) FirmataMessageDigitalIOBuilder {
	b.PinBlock = pinBlock
	return b
}

func (b *_FirmataMessageDigitalIOBuilder) WithData(data ...int8) FirmataMessageDigitalIOBuilder {
	b.Data = data
	return b
}

func (b *_FirmataMessageDigitalIOBuilder) Build() (FirmataMessageDigitalIO, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._FirmataMessageDigitalIO.deepCopy(), nil
}

func (b *_FirmataMessageDigitalIOBuilder) MustBuild() FirmataMessageDigitalIO {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_FirmataMessageDigitalIOBuilder) Done() FirmataMessageBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewFirmataMessageBuilder().(*_FirmataMessageBuilder)
	}
	return b.parentBuilder
}

func (b *_FirmataMessageDigitalIOBuilder) buildForFirmataMessage() (FirmataMessage, error) {
	return b.Build()
}

func (b *_FirmataMessageDigitalIOBuilder) DeepCopy() any {
	_copy := b.CreateFirmataMessageDigitalIOBuilder().(*_FirmataMessageDigitalIOBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateFirmataMessageDigitalIOBuilder creates a FirmataMessageDigitalIOBuilder
func (b *_FirmataMessageDigitalIO) CreateFirmataMessageDigitalIOBuilder() FirmataMessageDigitalIOBuilder {
	if b == nil {
		return NewFirmataMessageDigitalIOBuilder()
	}
	return &_FirmataMessageDigitalIOBuilder{_FirmataMessageDigitalIO: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_FirmataMessageDigitalIO) GetMessageType() uint8 {
	return 0x9
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_FirmataMessageDigitalIO) GetParent() FirmataMessageContract {
	return m.FirmataMessageContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_FirmataMessageDigitalIO) GetPinBlock() uint8 {
	return m.PinBlock
}

func (m *_FirmataMessageDigitalIO) GetData() []int8 {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastFirmataMessageDigitalIO(structType any) FirmataMessageDigitalIO {
	if casted, ok := structType.(FirmataMessageDigitalIO); ok {
		return casted
	}
	if casted, ok := structType.(*FirmataMessageDigitalIO); ok {
		return *casted
	}
	return nil
}

func (m *_FirmataMessageDigitalIO) GetTypeName() string {
	return "FirmataMessageDigitalIO"
}

func (m *_FirmataMessageDigitalIO) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.FirmataMessageContract.(*_FirmataMessage).getLengthInBits(ctx))

	// Simple field (pinBlock)
	lengthInBits += 4

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *_FirmataMessageDigitalIO) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_FirmataMessageDigitalIO) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_FirmataMessage, response bool) (__firmataMessageDigitalIO FirmataMessageDigitalIO, err error) {
	m.FirmataMessageContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("FirmataMessageDigitalIO"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for FirmataMessageDigitalIO")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	pinBlock, err := ReadSimpleField(ctx, "pinBlock", ReadUnsignedByte(readBuffer, uint8(4)), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'pinBlock' field"))
	}
	m.PinBlock = pinBlock

	data, err := ReadCountArrayField[int8](ctx, "data", ReadSignedByte(readBuffer, uint8(8)), uint64(int32(2)), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'data' field"))
	}
	m.Data = data

	if closeErr := readBuffer.CloseContext("FirmataMessageDigitalIO"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for FirmataMessageDigitalIO")
	}

	return m, nil
}

func (m *_FirmataMessageDigitalIO) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))), utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_FirmataMessageDigitalIO) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("FirmataMessageDigitalIO"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for FirmataMessageDigitalIO")
		}

		if err := WriteSimpleField[uint8](ctx, "pinBlock", m.GetPinBlock(), WriteUnsignedByte(writeBuffer, 4), codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'pinBlock' field")
		}

		if err := WriteSimpleTypeArrayField(ctx, "data", m.GetData(), WriteSignedByte(writeBuffer, 8), codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'data' field")
		}

		if popErr := writeBuffer.PopContext("FirmataMessageDigitalIO"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for FirmataMessageDigitalIO")
		}
		return nil
	}
	return m.FirmataMessageContract.(*_FirmataMessage).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_FirmataMessageDigitalIO) IsFirmataMessageDigitalIO() {}

func (m *_FirmataMessageDigitalIO) DeepCopy() any {
	return m.deepCopy()
}

func (m *_FirmataMessageDigitalIO) deepCopy() *_FirmataMessageDigitalIO {
	if m == nil {
		return nil
	}
	_FirmataMessageDigitalIOCopy := &_FirmataMessageDigitalIO{
		m.FirmataMessageContract.(*_FirmataMessage).deepCopy(),
		m.PinBlock,
		utils.DeepCopySlice[int8, int8](m.Data),
	}
	m.FirmataMessageContract.(*_FirmataMessage)._SubType = m
	return _FirmataMessageDigitalIOCopy
}

func (m *_FirmataMessageDigitalIO) String() string {
	if m == nil {
		return "<nil>"
	}
	wb := utils.NewWriteBufferBoxBased(
		utils.WithWriteBufferBoxBasedMergeSingleBoxes(),
		utils.WithWriteBufferBoxBasedOmitEmptyBoxes(),
		utils.WithWriteBufferBoxBasedPrintPosLengthFooter(),
	)
	if err := wb.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return wb.GetBox().String()
}
