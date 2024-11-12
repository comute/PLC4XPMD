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

// ParameterValueSerialNumber is the corresponding interface of ParameterValueSerialNumber
type ParameterValueSerialNumber interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ParameterValue
	// GetValue returns Value (property field)
	GetValue() SerialNumber
	// GetData returns Data (property field)
	GetData() []byte
	// IsParameterValueSerialNumber is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsParameterValueSerialNumber()
	// CreateBuilder creates a ParameterValueSerialNumberBuilder
	CreateParameterValueSerialNumberBuilder() ParameterValueSerialNumberBuilder
}

// _ParameterValueSerialNumber is the data-structure of this message
type _ParameterValueSerialNumber struct {
	ParameterValueContract
	Value SerialNumber
	Data  []byte
}

var _ ParameterValueSerialNumber = (*_ParameterValueSerialNumber)(nil)
var _ ParameterValueRequirements = (*_ParameterValueSerialNumber)(nil)

// NewParameterValueSerialNumber factory function for _ParameterValueSerialNumber
func NewParameterValueSerialNumber(value SerialNumber, data []byte, numBytes uint8) *_ParameterValueSerialNumber {
	if value == nil {
		panic("value of type SerialNumber for ParameterValueSerialNumber must not be nil")
	}
	_result := &_ParameterValueSerialNumber{
		ParameterValueContract: NewParameterValue(numBytes),
		Value:                  value,
		Data:                   data,
	}
	_result.ParameterValueContract.(*_ParameterValue)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ParameterValueSerialNumberBuilder is a builder for ParameterValueSerialNumber
type ParameterValueSerialNumberBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(value SerialNumber, data []byte) ParameterValueSerialNumberBuilder
	// WithValue adds Value (property field)
	WithValue(SerialNumber) ParameterValueSerialNumberBuilder
	// WithValueBuilder adds Value (property field) which is build by the builder
	WithValueBuilder(func(SerialNumberBuilder) SerialNumberBuilder) ParameterValueSerialNumberBuilder
	// WithData adds Data (property field)
	WithData(...byte) ParameterValueSerialNumberBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ParameterValueBuilder
	// Build builds the ParameterValueSerialNumber or returns an error if something is wrong
	Build() (ParameterValueSerialNumber, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ParameterValueSerialNumber
}

// NewParameterValueSerialNumberBuilder() creates a ParameterValueSerialNumberBuilder
func NewParameterValueSerialNumberBuilder() ParameterValueSerialNumberBuilder {
	return &_ParameterValueSerialNumberBuilder{_ParameterValueSerialNumber: new(_ParameterValueSerialNumber)}
}

type _ParameterValueSerialNumberBuilder struct {
	*_ParameterValueSerialNumber

	parentBuilder *_ParameterValueBuilder

	err *utils.MultiError
}

var _ (ParameterValueSerialNumberBuilder) = (*_ParameterValueSerialNumberBuilder)(nil)

func (b *_ParameterValueSerialNumberBuilder) setParent(contract ParameterValueContract) {
	b.ParameterValueContract = contract
	contract.(*_ParameterValue)._SubType = b._ParameterValueSerialNumber
}

func (b *_ParameterValueSerialNumberBuilder) WithMandatoryFields(value SerialNumber, data []byte) ParameterValueSerialNumberBuilder {
	return b.WithValue(value).WithData(data...)
}

func (b *_ParameterValueSerialNumberBuilder) WithValue(value SerialNumber) ParameterValueSerialNumberBuilder {
	b.Value = value
	return b
}

func (b *_ParameterValueSerialNumberBuilder) WithValueBuilder(builderSupplier func(SerialNumberBuilder) SerialNumberBuilder) ParameterValueSerialNumberBuilder {
	builder := builderSupplier(b.Value.CreateSerialNumberBuilder())
	var err error
	b.Value, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "SerialNumberBuilder failed"))
	}
	return b
}

func (b *_ParameterValueSerialNumberBuilder) WithData(data ...byte) ParameterValueSerialNumberBuilder {
	b.Data = data
	return b
}

func (b *_ParameterValueSerialNumberBuilder) Build() (ParameterValueSerialNumber, error) {
	if b.Value == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'value' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ParameterValueSerialNumber.deepCopy(), nil
}

func (b *_ParameterValueSerialNumberBuilder) MustBuild() ParameterValueSerialNumber {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_ParameterValueSerialNumberBuilder) Done() ParameterValueBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewParameterValueBuilder().(*_ParameterValueBuilder)
	}
	return b.parentBuilder
}

func (b *_ParameterValueSerialNumberBuilder) buildForParameterValue() (ParameterValue, error) {
	return b.Build()
}

func (b *_ParameterValueSerialNumberBuilder) DeepCopy() any {
	_copy := b.CreateParameterValueSerialNumberBuilder().(*_ParameterValueSerialNumberBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateParameterValueSerialNumberBuilder creates a ParameterValueSerialNumberBuilder
func (b *_ParameterValueSerialNumber) CreateParameterValueSerialNumberBuilder() ParameterValueSerialNumberBuilder {
	if b == nil {
		return NewParameterValueSerialNumberBuilder()
	}
	return &_ParameterValueSerialNumberBuilder{_ParameterValueSerialNumber: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ParameterValueSerialNumber) GetParameterType() ParameterType {
	return ParameterType_SERIAL_NUMBER
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ParameterValueSerialNumber) GetParent() ParameterValueContract {
	return m.ParameterValueContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ParameterValueSerialNumber) GetValue() SerialNumber {
	return m.Value
}

func (m *_ParameterValueSerialNumber) GetData() []byte {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastParameterValueSerialNumber(structType any) ParameterValueSerialNumber {
	if casted, ok := structType.(ParameterValueSerialNumber); ok {
		return casted
	}
	if casted, ok := structType.(*ParameterValueSerialNumber); ok {
		return *casted
	}
	return nil
}

func (m *_ParameterValueSerialNumber) GetTypeName() string {
	return "ParameterValueSerialNumber"
}

func (m *_ParameterValueSerialNumber) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ParameterValueContract.(*_ParameterValue).getLengthInBits(ctx))

	// Simple field (value)
	lengthInBits += m.Value.GetLengthInBits(ctx)

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *_ParameterValueSerialNumber) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ParameterValueSerialNumber) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ParameterValue, parameterType ParameterType, numBytes uint8) (__parameterValueSerialNumber ParameterValueSerialNumber, err error) {
	m.ParameterValueContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ParameterValueSerialNumber"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ParameterValueSerialNumber")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(bool((numBytes) >= (4))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "SerialNumber has exactly four bytes"})
	}

	value, err := ReadSimpleField[SerialNumber](ctx, "value", ReadComplex[SerialNumber](SerialNumberParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	data, err := readBuffer.ReadByteArray("data", int(int32(numBytes)-int32(int32(4))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'data' field"))
	}
	m.Data = data

	if closeErr := readBuffer.CloseContext("ParameterValueSerialNumber"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ParameterValueSerialNumber")
	}

	return m, nil
}

func (m *_ParameterValueSerialNumber) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ParameterValueSerialNumber) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ParameterValueSerialNumber"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ParameterValueSerialNumber")
		}

		if err := WriteSimpleField[SerialNumber](ctx, "value", m.GetValue(), WriteComplex[SerialNumber](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'value' field")
		}

		if err := WriteByteArrayField(ctx, "data", m.GetData(), WriteByteArray(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'data' field")
		}

		if popErr := writeBuffer.PopContext("ParameterValueSerialNumber"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ParameterValueSerialNumber")
		}
		return nil
	}
	return m.ParameterValueContract.(*_ParameterValue).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ParameterValueSerialNumber) IsParameterValueSerialNumber() {}

func (m *_ParameterValueSerialNumber) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ParameterValueSerialNumber) deepCopy() *_ParameterValueSerialNumber {
	if m == nil {
		return nil
	}
	_ParameterValueSerialNumberCopy := &_ParameterValueSerialNumber{
		m.ParameterValueContract.(*_ParameterValue).deepCopy(),
		utils.DeepCopy[SerialNumber](m.Value),
		utils.DeepCopySlice[byte, byte](m.Data),
	}
	m.ParameterValueContract.(*_ParameterValue)._SubType = m
	return _ParameterValueSerialNumberCopy
}

func (m *_ParameterValueSerialNumber) String() string {
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
