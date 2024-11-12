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

// COTPPacketTpduError is the corresponding interface of COTPPacketTpduError
type COTPPacketTpduError interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	COTPPacket
	// GetDestinationReference returns DestinationReference (property field)
	GetDestinationReference() uint16
	// GetRejectCause returns RejectCause (property field)
	GetRejectCause() uint8
	// IsCOTPPacketTpduError is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsCOTPPacketTpduError()
	// CreateBuilder creates a COTPPacketTpduErrorBuilder
	CreateCOTPPacketTpduErrorBuilder() COTPPacketTpduErrorBuilder
}

// _COTPPacketTpduError is the data-structure of this message
type _COTPPacketTpduError struct {
	COTPPacketContract
	DestinationReference uint16
	RejectCause          uint8
}

var _ COTPPacketTpduError = (*_COTPPacketTpduError)(nil)
var _ COTPPacketRequirements = (*_COTPPacketTpduError)(nil)

// NewCOTPPacketTpduError factory function for _COTPPacketTpduError
func NewCOTPPacketTpduError(parameters []COTPParameter, payload S7Message, destinationReference uint16, rejectCause uint8, cotpLen uint16) *_COTPPacketTpduError {
	_result := &_COTPPacketTpduError{
		COTPPacketContract:   NewCOTPPacket(parameters, payload, cotpLen),
		DestinationReference: destinationReference,
		RejectCause:          rejectCause,
	}
	_result.COTPPacketContract.(*_COTPPacket)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// COTPPacketTpduErrorBuilder is a builder for COTPPacketTpduError
type COTPPacketTpduErrorBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(destinationReference uint16, rejectCause uint8) COTPPacketTpduErrorBuilder
	// WithDestinationReference adds DestinationReference (property field)
	WithDestinationReference(uint16) COTPPacketTpduErrorBuilder
	// WithRejectCause adds RejectCause (property field)
	WithRejectCause(uint8) COTPPacketTpduErrorBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() COTPPacketBuilder
	// Build builds the COTPPacketTpduError or returns an error if something is wrong
	Build() (COTPPacketTpduError, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() COTPPacketTpduError
}

// NewCOTPPacketTpduErrorBuilder() creates a COTPPacketTpduErrorBuilder
func NewCOTPPacketTpduErrorBuilder() COTPPacketTpduErrorBuilder {
	return &_COTPPacketTpduErrorBuilder{_COTPPacketTpduError: new(_COTPPacketTpduError)}
}

type _COTPPacketTpduErrorBuilder struct {
	*_COTPPacketTpduError

	parentBuilder *_COTPPacketBuilder

	err *utils.MultiError
}

var _ (COTPPacketTpduErrorBuilder) = (*_COTPPacketTpduErrorBuilder)(nil)

func (b *_COTPPacketTpduErrorBuilder) setParent(contract COTPPacketContract) {
	b.COTPPacketContract = contract
	contract.(*_COTPPacket)._SubType = b._COTPPacketTpduError
}

func (b *_COTPPacketTpduErrorBuilder) WithMandatoryFields(destinationReference uint16, rejectCause uint8) COTPPacketTpduErrorBuilder {
	return b.WithDestinationReference(destinationReference).WithRejectCause(rejectCause)
}

func (b *_COTPPacketTpduErrorBuilder) WithDestinationReference(destinationReference uint16) COTPPacketTpduErrorBuilder {
	b.DestinationReference = destinationReference
	return b
}

func (b *_COTPPacketTpduErrorBuilder) WithRejectCause(rejectCause uint8) COTPPacketTpduErrorBuilder {
	b.RejectCause = rejectCause
	return b
}

func (b *_COTPPacketTpduErrorBuilder) Build() (COTPPacketTpduError, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._COTPPacketTpduError.deepCopy(), nil
}

func (b *_COTPPacketTpduErrorBuilder) MustBuild() COTPPacketTpduError {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_COTPPacketTpduErrorBuilder) Done() COTPPacketBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewCOTPPacketBuilder().(*_COTPPacketBuilder)
	}
	return b.parentBuilder
}

func (b *_COTPPacketTpduErrorBuilder) buildForCOTPPacket() (COTPPacket, error) {
	return b.Build()
}

func (b *_COTPPacketTpduErrorBuilder) DeepCopy() any {
	_copy := b.CreateCOTPPacketTpduErrorBuilder().(*_COTPPacketTpduErrorBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateCOTPPacketTpduErrorBuilder creates a COTPPacketTpduErrorBuilder
func (b *_COTPPacketTpduError) CreateCOTPPacketTpduErrorBuilder() COTPPacketTpduErrorBuilder {
	if b == nil {
		return NewCOTPPacketTpduErrorBuilder()
	}
	return &_COTPPacketTpduErrorBuilder{_COTPPacketTpduError: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_COTPPacketTpduError) GetTpduCode() uint8 {
	return 0x70
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_COTPPacketTpduError) GetParent() COTPPacketContract {
	return m.COTPPacketContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_COTPPacketTpduError) GetDestinationReference() uint16 {
	return m.DestinationReference
}

func (m *_COTPPacketTpduError) GetRejectCause() uint8 {
	return m.RejectCause
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastCOTPPacketTpduError(structType any) COTPPacketTpduError {
	if casted, ok := structType.(COTPPacketTpduError); ok {
		return casted
	}
	if casted, ok := structType.(*COTPPacketTpduError); ok {
		return *casted
	}
	return nil
}

func (m *_COTPPacketTpduError) GetTypeName() string {
	return "COTPPacketTpduError"
}

func (m *_COTPPacketTpduError) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.COTPPacketContract.(*_COTPPacket).getLengthInBits(ctx))

	// Simple field (destinationReference)
	lengthInBits += 16

	// Simple field (rejectCause)
	lengthInBits += 8

	return lengthInBits
}

func (m *_COTPPacketTpduError) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_COTPPacketTpduError) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_COTPPacket, cotpLen uint16) (__cOTPPacketTpduError COTPPacketTpduError, err error) {
	m.COTPPacketContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("COTPPacketTpduError"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for COTPPacketTpduError")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	destinationReference, err := ReadSimpleField(ctx, "destinationReference", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'destinationReference' field"))
	}
	m.DestinationReference = destinationReference

	rejectCause, err := ReadSimpleField(ctx, "rejectCause", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'rejectCause' field"))
	}
	m.RejectCause = rejectCause

	if closeErr := readBuffer.CloseContext("COTPPacketTpduError"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for COTPPacketTpduError")
	}

	return m, nil
}

func (m *_COTPPacketTpduError) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_COTPPacketTpduError) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("COTPPacketTpduError"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for COTPPacketTpduError")
		}

		if err := WriteSimpleField[uint16](ctx, "destinationReference", m.GetDestinationReference(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'destinationReference' field")
		}

		if err := WriteSimpleField[uint8](ctx, "rejectCause", m.GetRejectCause(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'rejectCause' field")
		}

		if popErr := writeBuffer.PopContext("COTPPacketTpduError"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for COTPPacketTpduError")
		}
		return nil
	}
	return m.COTPPacketContract.(*_COTPPacket).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_COTPPacketTpduError) IsCOTPPacketTpduError() {}

func (m *_COTPPacketTpduError) DeepCopy() any {
	return m.deepCopy()
}

func (m *_COTPPacketTpduError) deepCopy() *_COTPPacketTpduError {
	if m == nil {
		return nil
	}
	_COTPPacketTpduErrorCopy := &_COTPPacketTpduError{
		m.COTPPacketContract.(*_COTPPacket).deepCopy(),
		m.DestinationReference,
		m.RejectCause,
	}
	m.COTPPacketContract.(*_COTPPacket)._SubType = m
	return _COTPPacketTpduErrorCopy
}

func (m *_COTPPacketTpduError) String() string {
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
