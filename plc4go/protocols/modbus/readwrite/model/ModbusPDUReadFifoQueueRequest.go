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

// ModbusPDUReadFifoQueueRequest is the corresponding interface of ModbusPDUReadFifoQueueRequest
type ModbusPDUReadFifoQueueRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ModbusPDU
	// GetFifoPointerAddress returns FifoPointerAddress (property field)
	GetFifoPointerAddress() uint16
	// IsModbusPDUReadFifoQueueRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsModbusPDUReadFifoQueueRequest()
	// CreateBuilder creates a ModbusPDUReadFifoQueueRequestBuilder
	CreateModbusPDUReadFifoQueueRequestBuilder() ModbusPDUReadFifoQueueRequestBuilder
}

// _ModbusPDUReadFifoQueueRequest is the data-structure of this message
type _ModbusPDUReadFifoQueueRequest struct {
	ModbusPDUContract
	FifoPointerAddress uint16
}

var _ ModbusPDUReadFifoQueueRequest = (*_ModbusPDUReadFifoQueueRequest)(nil)
var _ ModbusPDURequirements = (*_ModbusPDUReadFifoQueueRequest)(nil)

// NewModbusPDUReadFifoQueueRequest factory function for _ModbusPDUReadFifoQueueRequest
func NewModbusPDUReadFifoQueueRequest(fifoPointerAddress uint16) *_ModbusPDUReadFifoQueueRequest {
	_result := &_ModbusPDUReadFifoQueueRequest{
		ModbusPDUContract:  NewModbusPDU(),
		FifoPointerAddress: fifoPointerAddress,
	}
	_result.ModbusPDUContract.(*_ModbusPDU)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ModbusPDUReadFifoQueueRequestBuilder is a builder for ModbusPDUReadFifoQueueRequest
type ModbusPDUReadFifoQueueRequestBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(fifoPointerAddress uint16) ModbusPDUReadFifoQueueRequestBuilder
	// WithFifoPointerAddress adds FifoPointerAddress (property field)
	WithFifoPointerAddress(uint16) ModbusPDUReadFifoQueueRequestBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ModbusPDUBuilder
	// Build builds the ModbusPDUReadFifoQueueRequest or returns an error if something is wrong
	Build() (ModbusPDUReadFifoQueueRequest, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ModbusPDUReadFifoQueueRequest
}

// NewModbusPDUReadFifoQueueRequestBuilder() creates a ModbusPDUReadFifoQueueRequestBuilder
func NewModbusPDUReadFifoQueueRequestBuilder() ModbusPDUReadFifoQueueRequestBuilder {
	return &_ModbusPDUReadFifoQueueRequestBuilder{_ModbusPDUReadFifoQueueRequest: new(_ModbusPDUReadFifoQueueRequest)}
}

type _ModbusPDUReadFifoQueueRequestBuilder struct {
	*_ModbusPDUReadFifoQueueRequest

	parentBuilder *_ModbusPDUBuilder

	err *utils.MultiError
}

var _ (ModbusPDUReadFifoQueueRequestBuilder) = (*_ModbusPDUReadFifoQueueRequestBuilder)(nil)

func (b *_ModbusPDUReadFifoQueueRequestBuilder) setParent(contract ModbusPDUContract) {
	b.ModbusPDUContract = contract
	contract.(*_ModbusPDU)._SubType = b._ModbusPDUReadFifoQueueRequest
}

func (b *_ModbusPDUReadFifoQueueRequestBuilder) WithMandatoryFields(fifoPointerAddress uint16) ModbusPDUReadFifoQueueRequestBuilder {
	return b.WithFifoPointerAddress(fifoPointerAddress)
}

func (b *_ModbusPDUReadFifoQueueRequestBuilder) WithFifoPointerAddress(fifoPointerAddress uint16) ModbusPDUReadFifoQueueRequestBuilder {
	b.FifoPointerAddress = fifoPointerAddress
	return b
}

func (b *_ModbusPDUReadFifoQueueRequestBuilder) Build() (ModbusPDUReadFifoQueueRequest, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ModbusPDUReadFifoQueueRequest.deepCopy(), nil
}

func (b *_ModbusPDUReadFifoQueueRequestBuilder) MustBuild() ModbusPDUReadFifoQueueRequest {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_ModbusPDUReadFifoQueueRequestBuilder) Done() ModbusPDUBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewModbusPDUBuilder().(*_ModbusPDUBuilder)
	}
	return b.parentBuilder
}

func (b *_ModbusPDUReadFifoQueueRequestBuilder) buildForModbusPDU() (ModbusPDU, error) {
	return b.Build()
}

func (b *_ModbusPDUReadFifoQueueRequestBuilder) DeepCopy() any {
	_copy := b.CreateModbusPDUReadFifoQueueRequestBuilder().(*_ModbusPDUReadFifoQueueRequestBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateModbusPDUReadFifoQueueRequestBuilder creates a ModbusPDUReadFifoQueueRequestBuilder
func (b *_ModbusPDUReadFifoQueueRequest) CreateModbusPDUReadFifoQueueRequestBuilder() ModbusPDUReadFifoQueueRequestBuilder {
	if b == nil {
		return NewModbusPDUReadFifoQueueRequestBuilder()
	}
	return &_ModbusPDUReadFifoQueueRequestBuilder{_ModbusPDUReadFifoQueueRequest: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ModbusPDUReadFifoQueueRequest) GetErrorFlag() bool {
	return bool(false)
}

func (m *_ModbusPDUReadFifoQueueRequest) GetFunctionFlag() uint8 {
	return 0x18
}

func (m *_ModbusPDUReadFifoQueueRequest) GetResponse() bool {
	return bool(false)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ModbusPDUReadFifoQueueRequest) GetParent() ModbusPDUContract {
	return m.ModbusPDUContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ModbusPDUReadFifoQueueRequest) GetFifoPointerAddress() uint16 {
	return m.FifoPointerAddress
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastModbusPDUReadFifoQueueRequest(structType any) ModbusPDUReadFifoQueueRequest {
	if casted, ok := structType.(ModbusPDUReadFifoQueueRequest); ok {
		return casted
	}
	if casted, ok := structType.(*ModbusPDUReadFifoQueueRequest); ok {
		return *casted
	}
	return nil
}

func (m *_ModbusPDUReadFifoQueueRequest) GetTypeName() string {
	return "ModbusPDUReadFifoQueueRequest"
}

func (m *_ModbusPDUReadFifoQueueRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ModbusPDUContract.(*_ModbusPDU).getLengthInBits(ctx))

	// Simple field (fifoPointerAddress)
	lengthInBits += 16

	return lengthInBits
}

func (m *_ModbusPDUReadFifoQueueRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ModbusPDUReadFifoQueueRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ModbusPDU, response bool) (__modbusPDUReadFifoQueueRequest ModbusPDUReadFifoQueueRequest, err error) {
	m.ModbusPDUContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModbusPDUReadFifoQueueRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusPDUReadFifoQueueRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	fifoPointerAddress, err := ReadSimpleField(ctx, "fifoPointerAddress", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'fifoPointerAddress' field"))
	}
	m.FifoPointerAddress = fifoPointerAddress

	if closeErr := readBuffer.CloseContext("ModbusPDUReadFifoQueueRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusPDUReadFifoQueueRequest")
	}

	return m, nil
}

func (m *_ModbusPDUReadFifoQueueRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ModbusPDUReadFifoQueueRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUReadFifoQueueRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ModbusPDUReadFifoQueueRequest")
		}

		if err := WriteSimpleField[uint16](ctx, "fifoPointerAddress", m.GetFifoPointerAddress(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'fifoPointerAddress' field")
		}

		if popErr := writeBuffer.PopContext("ModbusPDUReadFifoQueueRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ModbusPDUReadFifoQueueRequest")
		}
		return nil
	}
	return m.ModbusPDUContract.(*_ModbusPDU).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ModbusPDUReadFifoQueueRequest) IsModbusPDUReadFifoQueueRequest() {}

func (m *_ModbusPDUReadFifoQueueRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ModbusPDUReadFifoQueueRequest) deepCopy() *_ModbusPDUReadFifoQueueRequest {
	if m == nil {
		return nil
	}
	_ModbusPDUReadFifoQueueRequestCopy := &_ModbusPDUReadFifoQueueRequest{
		m.ModbusPDUContract.(*_ModbusPDU).deepCopy(),
		m.FifoPointerAddress,
	}
	_ModbusPDUReadFifoQueueRequestCopy.ModbusPDUContract.(*_ModbusPDU)._SubType = m
	return _ModbusPDUReadFifoQueueRequestCopy
}

func (m *_ModbusPDUReadFifoQueueRequest) String() string {
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
