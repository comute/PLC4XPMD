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

// ParameterValueApplicationAddress2 is the corresponding interface of ParameterValueApplicationAddress2
type ParameterValueApplicationAddress2 interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ParameterValue
	// GetValue returns Value (property field)
	GetValue() ApplicationAddress2
	// GetData returns Data (property field)
	GetData() []byte
	// IsParameterValueApplicationAddress2 is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsParameterValueApplicationAddress2()
	// CreateBuilder creates a ParameterValueApplicationAddress2Builder
	CreateParameterValueApplicationAddress2Builder() ParameterValueApplicationAddress2Builder
}

// _ParameterValueApplicationAddress2 is the data-structure of this message
type _ParameterValueApplicationAddress2 struct {
	ParameterValueContract
	Value ApplicationAddress2
	Data  []byte
}

var _ ParameterValueApplicationAddress2 = (*_ParameterValueApplicationAddress2)(nil)
var _ ParameterValueRequirements = (*_ParameterValueApplicationAddress2)(nil)

// NewParameterValueApplicationAddress2 factory function for _ParameterValueApplicationAddress2
func NewParameterValueApplicationAddress2(value ApplicationAddress2, data []byte, numBytes uint8) *_ParameterValueApplicationAddress2 {
	if value == nil {
		panic("value of type ApplicationAddress2 for ParameterValueApplicationAddress2 must not be nil")
	}
	_result := &_ParameterValueApplicationAddress2{
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

// ParameterValueApplicationAddress2Builder is a builder for ParameterValueApplicationAddress2
type ParameterValueApplicationAddress2Builder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(value ApplicationAddress2, data []byte) ParameterValueApplicationAddress2Builder
	// WithValue adds Value (property field)
	WithValue(ApplicationAddress2) ParameterValueApplicationAddress2Builder
	// WithValueBuilder adds Value (property field) which is build by the builder
	WithValueBuilder(func(ApplicationAddress2Builder) ApplicationAddress2Builder) ParameterValueApplicationAddress2Builder
	// WithData adds Data (property field)
	WithData(...byte) ParameterValueApplicationAddress2Builder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ParameterValueBuilder
	// Build builds the ParameterValueApplicationAddress2 or returns an error if something is wrong
	Build() (ParameterValueApplicationAddress2, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ParameterValueApplicationAddress2
}

// NewParameterValueApplicationAddress2Builder() creates a ParameterValueApplicationAddress2Builder
func NewParameterValueApplicationAddress2Builder() ParameterValueApplicationAddress2Builder {
	return &_ParameterValueApplicationAddress2Builder{_ParameterValueApplicationAddress2: new(_ParameterValueApplicationAddress2)}
}

type _ParameterValueApplicationAddress2Builder struct {
	*_ParameterValueApplicationAddress2

	parentBuilder *_ParameterValueBuilder

	err *utils.MultiError
}

var _ (ParameterValueApplicationAddress2Builder) = (*_ParameterValueApplicationAddress2Builder)(nil)

func (b *_ParameterValueApplicationAddress2Builder) setParent(contract ParameterValueContract) {
	b.ParameterValueContract = contract
	contract.(*_ParameterValue)._SubType = b._ParameterValueApplicationAddress2
}

func (b *_ParameterValueApplicationAddress2Builder) WithMandatoryFields(value ApplicationAddress2, data []byte) ParameterValueApplicationAddress2Builder {
	return b.WithValue(value).WithData(data...)
}

func (b *_ParameterValueApplicationAddress2Builder) WithValue(value ApplicationAddress2) ParameterValueApplicationAddress2Builder {
	b.Value = value
	return b
}

func (b *_ParameterValueApplicationAddress2Builder) WithValueBuilder(builderSupplier func(ApplicationAddress2Builder) ApplicationAddress2Builder) ParameterValueApplicationAddress2Builder {
	builder := builderSupplier(b.Value.CreateApplicationAddress2Builder())
	var err error
	b.Value, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "ApplicationAddress2Builder failed"))
	}
	return b
}

func (b *_ParameterValueApplicationAddress2Builder) WithData(data ...byte) ParameterValueApplicationAddress2Builder {
	b.Data = data
	return b
}

func (b *_ParameterValueApplicationAddress2Builder) Build() (ParameterValueApplicationAddress2, error) {
	if b.Value == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'value' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ParameterValueApplicationAddress2.deepCopy(), nil
}

func (b *_ParameterValueApplicationAddress2Builder) MustBuild() ParameterValueApplicationAddress2 {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_ParameterValueApplicationAddress2Builder) Done() ParameterValueBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewParameterValueBuilder().(*_ParameterValueBuilder)
	}
	return b.parentBuilder
}

func (b *_ParameterValueApplicationAddress2Builder) buildForParameterValue() (ParameterValue, error) {
	return b.Build()
}

func (b *_ParameterValueApplicationAddress2Builder) DeepCopy() any {
	_copy := b.CreateParameterValueApplicationAddress2Builder().(*_ParameterValueApplicationAddress2Builder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateParameterValueApplicationAddress2Builder creates a ParameterValueApplicationAddress2Builder
func (b *_ParameterValueApplicationAddress2) CreateParameterValueApplicationAddress2Builder() ParameterValueApplicationAddress2Builder {
	if b == nil {
		return NewParameterValueApplicationAddress2Builder()
	}
	return &_ParameterValueApplicationAddress2Builder{_ParameterValueApplicationAddress2: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ParameterValueApplicationAddress2) GetParameterType() ParameterType {
	return ParameterType_APPLICATION_ADDRESS_2
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ParameterValueApplicationAddress2) GetParent() ParameterValueContract {
	return m.ParameterValueContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ParameterValueApplicationAddress2) GetValue() ApplicationAddress2 {
	return m.Value
}

func (m *_ParameterValueApplicationAddress2) GetData() []byte {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastParameterValueApplicationAddress2(structType any) ParameterValueApplicationAddress2 {
	if casted, ok := structType.(ParameterValueApplicationAddress2); ok {
		return casted
	}
	if casted, ok := structType.(*ParameterValueApplicationAddress2); ok {
		return *casted
	}
	return nil
}

func (m *_ParameterValueApplicationAddress2) GetTypeName() string {
	return "ParameterValueApplicationAddress2"
}

func (m *_ParameterValueApplicationAddress2) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ParameterValueContract.(*_ParameterValue).getLengthInBits(ctx))

	// Simple field (value)
	lengthInBits += m.Value.GetLengthInBits(ctx)

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *_ParameterValueApplicationAddress2) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ParameterValueApplicationAddress2) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ParameterValue, parameterType ParameterType, numBytes uint8) (__parameterValueApplicationAddress2 ParameterValueApplicationAddress2, err error) {
	m.ParameterValueContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ParameterValueApplicationAddress2"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ParameterValueApplicationAddress2")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(bool((numBytes) >= (1))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "ApplicationAddress2 has exactly one byte"})
	}

	value, err := ReadSimpleField[ApplicationAddress2](ctx, "value", ReadComplex[ApplicationAddress2](ApplicationAddress2ParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	data, err := readBuffer.ReadByteArray("data", int(int32(numBytes)-int32(int32(1))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'data' field"))
	}
	m.Data = data

	if closeErr := readBuffer.CloseContext("ParameterValueApplicationAddress2"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ParameterValueApplicationAddress2")
	}

	return m, nil
}

func (m *_ParameterValueApplicationAddress2) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ParameterValueApplicationAddress2) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ParameterValueApplicationAddress2"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ParameterValueApplicationAddress2")
		}

		if err := WriteSimpleField[ApplicationAddress2](ctx, "value", m.GetValue(), WriteComplex[ApplicationAddress2](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'value' field")
		}

		if err := WriteByteArrayField(ctx, "data", m.GetData(), WriteByteArray(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'data' field")
		}

		if popErr := writeBuffer.PopContext("ParameterValueApplicationAddress2"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ParameterValueApplicationAddress2")
		}
		return nil
	}
	return m.ParameterValueContract.(*_ParameterValue).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ParameterValueApplicationAddress2) IsParameterValueApplicationAddress2() {}

func (m *_ParameterValueApplicationAddress2) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ParameterValueApplicationAddress2) deepCopy() *_ParameterValueApplicationAddress2 {
	if m == nil {
		return nil
	}
	_ParameterValueApplicationAddress2Copy := &_ParameterValueApplicationAddress2{
		m.ParameterValueContract.(*_ParameterValue).deepCopy(),
		utils.DeepCopy[ApplicationAddress2](m.Value),
		utils.DeepCopySlice[byte, byte](m.Data),
	}
	m.ParameterValueContract.(*_ParameterValue)._SubType = m
	return _ParameterValueApplicationAddress2Copy
}

func (m *_ParameterValueApplicationAddress2) String() string {
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
