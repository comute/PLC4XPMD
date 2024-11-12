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

// APDUError is the corresponding interface of APDUError
type APDUError interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	APDU
	// GetOriginalInvokeId returns OriginalInvokeId (property field)
	GetOriginalInvokeId() uint8
	// GetErrorChoice returns ErrorChoice (property field)
	GetErrorChoice() BACnetConfirmedServiceChoice
	// GetError returns Error (property field)
	GetError() BACnetError
	// IsAPDUError is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsAPDUError()
	// CreateBuilder creates a APDUErrorBuilder
	CreateAPDUErrorBuilder() APDUErrorBuilder
}

// _APDUError is the data-structure of this message
type _APDUError struct {
	APDUContract
	OriginalInvokeId uint8
	ErrorChoice      BACnetConfirmedServiceChoice
	Error            BACnetError
	// Reserved Fields
	reservedField0 *uint8
}

var _ APDUError = (*_APDUError)(nil)
var _ APDURequirements = (*_APDUError)(nil)

// NewAPDUError factory function for _APDUError
func NewAPDUError(originalInvokeId uint8, errorChoice BACnetConfirmedServiceChoice, error BACnetError, apduLength uint16) *_APDUError {
	if error == nil {
		panic("error of type BACnetError for APDUError must not be nil")
	}
	_result := &_APDUError{
		APDUContract:     NewAPDU(apduLength),
		OriginalInvokeId: originalInvokeId,
		ErrorChoice:      errorChoice,
		Error:            error,
	}
	_result.APDUContract.(*_APDU)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// APDUErrorBuilder is a builder for APDUError
type APDUErrorBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(originalInvokeId uint8, errorChoice BACnetConfirmedServiceChoice, error BACnetError) APDUErrorBuilder
	// WithOriginalInvokeId adds OriginalInvokeId (property field)
	WithOriginalInvokeId(uint8) APDUErrorBuilder
	// WithErrorChoice adds ErrorChoice (property field)
	WithErrorChoice(BACnetConfirmedServiceChoice) APDUErrorBuilder
	// WithError adds Error (property field)
	WithError(BACnetError) APDUErrorBuilder
	// WithErrorBuilder adds Error (property field) which is build by the builder
	WithErrorBuilder(func(BACnetErrorBuilder) BACnetErrorBuilder) APDUErrorBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() APDUBuilder
	// Build builds the APDUError or returns an error if something is wrong
	Build() (APDUError, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() APDUError
}

// NewAPDUErrorBuilder() creates a APDUErrorBuilder
func NewAPDUErrorBuilder() APDUErrorBuilder {
	return &_APDUErrorBuilder{_APDUError: new(_APDUError)}
}

type _APDUErrorBuilder struct {
	*_APDUError

	parentBuilder *_APDUBuilder

	err *utils.MultiError
}

var _ (APDUErrorBuilder) = (*_APDUErrorBuilder)(nil)

func (b *_APDUErrorBuilder) setParent(contract APDUContract) {
	b.APDUContract = contract
	contract.(*_APDU)._SubType = b._APDUError
}

func (b *_APDUErrorBuilder) WithMandatoryFields(originalInvokeId uint8, errorChoice BACnetConfirmedServiceChoice, error BACnetError) APDUErrorBuilder {
	return b.WithOriginalInvokeId(originalInvokeId).WithErrorChoice(errorChoice).WithError(error)
}

func (b *_APDUErrorBuilder) WithOriginalInvokeId(originalInvokeId uint8) APDUErrorBuilder {
	b.OriginalInvokeId = originalInvokeId
	return b
}

func (b *_APDUErrorBuilder) WithErrorChoice(errorChoice BACnetConfirmedServiceChoice) APDUErrorBuilder {
	b.ErrorChoice = errorChoice
	return b
}

func (b *_APDUErrorBuilder) WithError(error BACnetError) APDUErrorBuilder {
	b.Error = error
	return b
}

func (b *_APDUErrorBuilder) WithErrorBuilder(builderSupplier func(BACnetErrorBuilder) BACnetErrorBuilder) APDUErrorBuilder {
	builder := builderSupplier(b.Error.CreateBACnetErrorBuilder())
	var err error
	b.Error, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetErrorBuilder failed"))
	}
	return b
}

func (b *_APDUErrorBuilder) Build() (APDUError, error) {
	if b.Error == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'error' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._APDUError.deepCopy(), nil
}

func (b *_APDUErrorBuilder) MustBuild() APDUError {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_APDUErrorBuilder) Done() APDUBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewAPDUBuilder().(*_APDUBuilder)
	}
	return b.parentBuilder
}

func (b *_APDUErrorBuilder) buildForAPDU() (APDU, error) {
	return b.Build()
}

func (b *_APDUErrorBuilder) DeepCopy() any {
	_copy := b.CreateAPDUErrorBuilder().(*_APDUErrorBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateAPDUErrorBuilder creates a APDUErrorBuilder
func (b *_APDUError) CreateAPDUErrorBuilder() APDUErrorBuilder {
	if b == nil {
		return NewAPDUErrorBuilder()
	}
	return &_APDUErrorBuilder{_APDUError: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_APDUError) GetApduType() ApduType {
	return ApduType_ERROR_PDU
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_APDUError) GetParent() APDUContract {
	return m.APDUContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_APDUError) GetOriginalInvokeId() uint8 {
	return m.OriginalInvokeId
}

func (m *_APDUError) GetErrorChoice() BACnetConfirmedServiceChoice {
	return m.ErrorChoice
}

func (m *_APDUError) GetError() BACnetError {
	return m.Error
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastAPDUError(structType any) APDUError {
	if casted, ok := structType.(APDUError); ok {
		return casted
	}
	if casted, ok := structType.(*APDUError); ok {
		return *casted
	}
	return nil
}

func (m *_APDUError) GetTypeName() string {
	return "APDUError"
}

func (m *_APDUError) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.APDUContract.(*_APDU).getLengthInBits(ctx))

	// Reserved Field (reserved)
	lengthInBits += 4

	// Simple field (originalInvokeId)
	lengthInBits += 8

	// Simple field (errorChoice)
	lengthInBits += 8

	// Simple field (error)
	lengthInBits += m.Error.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_APDUError) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_APDUError) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_APDU, apduLength uint16) (__aPDUError APDUError, err error) {
	m.APDUContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("APDUError"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for APDUError")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(4)), uint8(0x00))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	originalInvokeId, err := ReadSimpleField(ctx, "originalInvokeId", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'originalInvokeId' field"))
	}
	m.OriginalInvokeId = originalInvokeId

	errorChoice, err := ReadEnumField[BACnetConfirmedServiceChoice](ctx, "errorChoice", "BACnetConfirmedServiceChoice", ReadEnum(BACnetConfirmedServiceChoiceByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'errorChoice' field"))
	}
	m.ErrorChoice = errorChoice

	error, err := ReadSimpleField[BACnetError](ctx, "error", ReadComplex[BACnetError](BACnetErrorParseWithBufferProducer[BACnetError]((BACnetConfirmedServiceChoice)(errorChoice)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'error' field"))
	}
	m.Error = error

	if closeErr := readBuffer.CloseContext("APDUError"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for APDUError")
	}

	return m, nil
}

func (m *_APDUError) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_APDUError) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("APDUError"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for APDUError")
		}

		if err := WriteReservedField[uint8](ctx, "reserved", uint8(0x00), WriteUnsignedByte(writeBuffer, 4)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}

		if err := WriteSimpleField[uint8](ctx, "originalInvokeId", m.GetOriginalInvokeId(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'originalInvokeId' field")
		}

		if err := WriteSimpleEnumField[BACnetConfirmedServiceChoice](ctx, "errorChoice", "BACnetConfirmedServiceChoice", m.GetErrorChoice(), WriteEnum[BACnetConfirmedServiceChoice, uint8](BACnetConfirmedServiceChoice.GetValue, BACnetConfirmedServiceChoice.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
			return errors.Wrap(err, "Error serializing 'errorChoice' field")
		}

		if err := WriteSimpleField[BACnetError](ctx, "error", m.GetError(), WriteComplex[BACnetError](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'error' field")
		}

		if popErr := writeBuffer.PopContext("APDUError"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for APDUError")
		}
		return nil
	}
	return m.APDUContract.(*_APDU).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_APDUError) IsAPDUError() {}

func (m *_APDUError) DeepCopy() any {
	return m.deepCopy()
}

func (m *_APDUError) deepCopy() *_APDUError {
	if m == nil {
		return nil
	}
	_APDUErrorCopy := &_APDUError{
		m.APDUContract.(*_APDU).deepCopy(),
		m.OriginalInvokeId,
		m.ErrorChoice,
		utils.DeepCopy[BACnetError](m.Error),
		m.reservedField0,
	}
	m.APDUContract.(*_APDU)._SubType = m
	return _APDUErrorCopy
}

func (m *_APDUError) String() string {
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
