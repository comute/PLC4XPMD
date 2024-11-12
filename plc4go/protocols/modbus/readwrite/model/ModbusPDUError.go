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

// ModbusPDUError is the corresponding interface of ModbusPDUError
type ModbusPDUError interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ModbusPDU
	// GetExceptionCode returns ExceptionCode (property field)
	GetExceptionCode() ModbusErrorCode
	// IsModbusPDUError is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsModbusPDUError()
	// CreateBuilder creates a ModbusPDUErrorBuilder
	CreateModbusPDUErrorBuilder() ModbusPDUErrorBuilder
}

// _ModbusPDUError is the data-structure of this message
type _ModbusPDUError struct {
	ModbusPDUContract
	ExceptionCode ModbusErrorCode
}

var _ ModbusPDUError = (*_ModbusPDUError)(nil)
var _ ModbusPDURequirements = (*_ModbusPDUError)(nil)

// NewModbusPDUError factory function for _ModbusPDUError
func NewModbusPDUError(exceptionCode ModbusErrorCode) *_ModbusPDUError {
	_result := &_ModbusPDUError{
		ModbusPDUContract: NewModbusPDU(),
		ExceptionCode:     exceptionCode,
	}
	_result.ModbusPDUContract.(*_ModbusPDU)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ModbusPDUErrorBuilder is a builder for ModbusPDUError
type ModbusPDUErrorBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(exceptionCode ModbusErrorCode) ModbusPDUErrorBuilder
	// WithExceptionCode adds ExceptionCode (property field)
	WithExceptionCode(ModbusErrorCode) ModbusPDUErrorBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ModbusPDUBuilder
	// Build builds the ModbusPDUError or returns an error if something is wrong
	Build() (ModbusPDUError, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ModbusPDUError
}

// NewModbusPDUErrorBuilder() creates a ModbusPDUErrorBuilder
func NewModbusPDUErrorBuilder() ModbusPDUErrorBuilder {
	return &_ModbusPDUErrorBuilder{_ModbusPDUError: new(_ModbusPDUError)}
}

type _ModbusPDUErrorBuilder struct {
	*_ModbusPDUError

	parentBuilder *_ModbusPDUBuilder

	err *utils.MultiError
}

var _ (ModbusPDUErrorBuilder) = (*_ModbusPDUErrorBuilder)(nil)

func (b *_ModbusPDUErrorBuilder) setParent(contract ModbusPDUContract) {
	b.ModbusPDUContract = contract
	contract.(*_ModbusPDU)._SubType = b._ModbusPDUError
}

func (b *_ModbusPDUErrorBuilder) WithMandatoryFields(exceptionCode ModbusErrorCode) ModbusPDUErrorBuilder {
	return b.WithExceptionCode(exceptionCode)
}

func (b *_ModbusPDUErrorBuilder) WithExceptionCode(exceptionCode ModbusErrorCode) ModbusPDUErrorBuilder {
	b.ExceptionCode = exceptionCode
	return b
}

func (b *_ModbusPDUErrorBuilder) Build() (ModbusPDUError, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ModbusPDUError.deepCopy(), nil
}

func (b *_ModbusPDUErrorBuilder) MustBuild() ModbusPDUError {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_ModbusPDUErrorBuilder) Done() ModbusPDUBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewModbusPDUBuilder().(*_ModbusPDUBuilder)
	}
	return b.parentBuilder
}

func (b *_ModbusPDUErrorBuilder) buildForModbusPDU() (ModbusPDU, error) {
	return b.Build()
}

func (b *_ModbusPDUErrorBuilder) DeepCopy() any {
	_copy := b.CreateModbusPDUErrorBuilder().(*_ModbusPDUErrorBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateModbusPDUErrorBuilder creates a ModbusPDUErrorBuilder
func (b *_ModbusPDUError) CreateModbusPDUErrorBuilder() ModbusPDUErrorBuilder {
	if b == nil {
		return NewModbusPDUErrorBuilder()
	}
	return &_ModbusPDUErrorBuilder{_ModbusPDUError: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ModbusPDUError) GetErrorFlag() bool {
	return bool(true)
}

func (m *_ModbusPDUError) GetFunctionFlag() uint8 {
	return 0
}

func (m *_ModbusPDUError) GetResponse() bool {
	return false
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ModbusPDUError) GetParent() ModbusPDUContract {
	return m.ModbusPDUContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ModbusPDUError) GetExceptionCode() ModbusErrorCode {
	return m.ExceptionCode
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastModbusPDUError(structType any) ModbusPDUError {
	if casted, ok := structType.(ModbusPDUError); ok {
		return casted
	}
	if casted, ok := structType.(*ModbusPDUError); ok {
		return *casted
	}
	return nil
}

func (m *_ModbusPDUError) GetTypeName() string {
	return "ModbusPDUError"
}

func (m *_ModbusPDUError) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ModbusPDUContract.(*_ModbusPDU).getLengthInBits(ctx))

	// Simple field (exceptionCode)
	lengthInBits += 8

	return lengthInBits
}

func (m *_ModbusPDUError) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ModbusPDUError) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ModbusPDU, response bool) (__modbusPDUError ModbusPDUError, err error) {
	m.ModbusPDUContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModbusPDUError"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusPDUError")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	exceptionCode, err := ReadEnumField[ModbusErrorCode](ctx, "exceptionCode", "ModbusErrorCode", ReadEnum(ModbusErrorCodeByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'exceptionCode' field"))
	}
	m.ExceptionCode = exceptionCode

	if closeErr := readBuffer.CloseContext("ModbusPDUError"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusPDUError")
	}

	return m, nil
}

func (m *_ModbusPDUError) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ModbusPDUError) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUError"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ModbusPDUError")
		}

		if err := WriteSimpleEnumField[ModbusErrorCode](ctx, "exceptionCode", "ModbusErrorCode", m.GetExceptionCode(), WriteEnum[ModbusErrorCode, uint8](ModbusErrorCode.GetValue, ModbusErrorCode.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
			return errors.Wrap(err, "Error serializing 'exceptionCode' field")
		}

		if popErr := writeBuffer.PopContext("ModbusPDUError"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ModbusPDUError")
		}
		return nil
	}
	return m.ModbusPDUContract.(*_ModbusPDU).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ModbusPDUError) IsModbusPDUError() {}

func (m *_ModbusPDUError) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ModbusPDUError) deepCopy() *_ModbusPDUError {
	if m == nil {
		return nil
	}
	_ModbusPDUErrorCopy := &_ModbusPDUError{
		m.ModbusPDUContract.(*_ModbusPDU).deepCopy(),
		m.ExceptionCode,
	}
	m.ModbusPDUContract.(*_ModbusPDU)._SubType = m
	return _ModbusPDUErrorCopy
}

func (m *_ModbusPDUError) String() string {
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
