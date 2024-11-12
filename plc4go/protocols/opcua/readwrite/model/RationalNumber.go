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

// RationalNumber is the corresponding interface of RationalNumber
type RationalNumber interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetNumerator returns Numerator (property field)
	GetNumerator() int32
	// GetDenominator returns Denominator (property field)
	GetDenominator() uint32
	// IsRationalNumber is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsRationalNumber()
	// CreateBuilder creates a RationalNumberBuilder
	CreateRationalNumberBuilder() RationalNumberBuilder
}

// _RationalNumber is the data-structure of this message
type _RationalNumber struct {
	ExtensionObjectDefinitionContract
	Numerator   int32
	Denominator uint32
}

var _ RationalNumber = (*_RationalNumber)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_RationalNumber)(nil)

// NewRationalNumber factory function for _RationalNumber
func NewRationalNumber(numerator int32, denominator uint32) *_RationalNumber {
	_result := &_RationalNumber{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		Numerator:                         numerator,
		Denominator:                       denominator,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// RationalNumberBuilder is a builder for RationalNumber
type RationalNumberBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(numerator int32, denominator uint32) RationalNumberBuilder
	// WithNumerator adds Numerator (property field)
	WithNumerator(int32) RationalNumberBuilder
	// WithDenominator adds Denominator (property field)
	WithDenominator(uint32) RationalNumberBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ExtensionObjectDefinitionBuilder
	// Build builds the RationalNumber or returns an error if something is wrong
	Build() (RationalNumber, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() RationalNumber
}

// NewRationalNumberBuilder() creates a RationalNumberBuilder
func NewRationalNumberBuilder() RationalNumberBuilder {
	return &_RationalNumberBuilder{_RationalNumber: new(_RationalNumber)}
}

type _RationalNumberBuilder struct {
	*_RationalNumber

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (RationalNumberBuilder) = (*_RationalNumberBuilder)(nil)

func (b *_RationalNumberBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
	contract.(*_ExtensionObjectDefinition)._SubType = b._RationalNumber
}

func (b *_RationalNumberBuilder) WithMandatoryFields(numerator int32, denominator uint32) RationalNumberBuilder {
	return b.WithNumerator(numerator).WithDenominator(denominator)
}

func (b *_RationalNumberBuilder) WithNumerator(numerator int32) RationalNumberBuilder {
	b.Numerator = numerator
	return b
}

func (b *_RationalNumberBuilder) WithDenominator(denominator uint32) RationalNumberBuilder {
	b.Denominator = denominator
	return b
}

func (b *_RationalNumberBuilder) Build() (RationalNumber, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._RationalNumber.deepCopy(), nil
}

func (b *_RationalNumberBuilder) MustBuild() RationalNumber {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_RationalNumberBuilder) Done() ExtensionObjectDefinitionBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewExtensionObjectDefinitionBuilder().(*_ExtensionObjectDefinitionBuilder)
	}
	return b.parentBuilder
}

func (b *_RationalNumberBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_RationalNumberBuilder) DeepCopy() any {
	_copy := b.CreateRationalNumberBuilder().(*_RationalNumberBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateRationalNumberBuilder creates a RationalNumberBuilder
func (b *_RationalNumber) CreateRationalNumberBuilder() RationalNumberBuilder {
	if b == nil {
		return NewRationalNumberBuilder()
	}
	return &_RationalNumberBuilder{_RationalNumber: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_RationalNumber) GetExtensionId() int32 {
	return int32(18808)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_RationalNumber) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_RationalNumber) GetNumerator() int32 {
	return m.Numerator
}

func (m *_RationalNumber) GetDenominator() uint32 {
	return m.Denominator
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastRationalNumber(structType any) RationalNumber {
	if casted, ok := structType.(RationalNumber); ok {
		return casted
	}
	if casted, ok := structType.(*RationalNumber); ok {
		return *casted
	}
	return nil
}

func (m *_RationalNumber) GetTypeName() string {
	return "RationalNumber"
}

func (m *_RationalNumber) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (numerator)
	lengthInBits += 32

	// Simple field (denominator)
	lengthInBits += 32

	return lengthInBits
}

func (m *_RationalNumber) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_RationalNumber) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__rationalNumber RationalNumber, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("RationalNumber"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for RationalNumber")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	numerator, err := ReadSimpleField(ctx, "numerator", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'numerator' field"))
	}
	m.Numerator = numerator

	denominator, err := ReadSimpleField(ctx, "denominator", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'denominator' field"))
	}
	m.Denominator = denominator

	if closeErr := readBuffer.CloseContext("RationalNumber"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for RationalNumber")
	}

	return m, nil
}

func (m *_RationalNumber) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_RationalNumber) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("RationalNumber"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for RationalNumber")
		}

		if err := WriteSimpleField[int32](ctx, "numerator", m.GetNumerator(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'numerator' field")
		}

		if err := WriteSimpleField[uint32](ctx, "denominator", m.GetDenominator(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'denominator' field")
		}

		if popErr := writeBuffer.PopContext("RationalNumber"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for RationalNumber")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_RationalNumber) IsRationalNumber() {}

func (m *_RationalNumber) DeepCopy() any {
	return m.deepCopy()
}

func (m *_RationalNumber) deepCopy() *_RationalNumber {
	if m == nil {
		return nil
	}
	_RationalNumberCopy := &_RationalNumber{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.Numerator,
		m.Denominator,
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _RationalNumberCopy
}

func (m *_RationalNumber) String() string {
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
