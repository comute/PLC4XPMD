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

// BACnetLogRecordLogDatumIntegerValue is the corresponding interface of BACnetLogRecordLogDatumIntegerValue
type BACnetLogRecordLogDatumIntegerValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetLogRecordLogDatum
	// GetIntegerValue returns IntegerValue (property field)
	GetIntegerValue() BACnetContextTagSignedInteger
	// IsBACnetLogRecordLogDatumIntegerValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetLogRecordLogDatumIntegerValue()
	// CreateBuilder creates a BACnetLogRecordLogDatumIntegerValueBuilder
	CreateBACnetLogRecordLogDatumIntegerValueBuilder() BACnetLogRecordLogDatumIntegerValueBuilder
}

// _BACnetLogRecordLogDatumIntegerValue is the data-structure of this message
type _BACnetLogRecordLogDatumIntegerValue struct {
	BACnetLogRecordLogDatumContract
	IntegerValue BACnetContextTagSignedInteger
}

var _ BACnetLogRecordLogDatumIntegerValue = (*_BACnetLogRecordLogDatumIntegerValue)(nil)
var _ BACnetLogRecordLogDatumRequirements = (*_BACnetLogRecordLogDatumIntegerValue)(nil)

// NewBACnetLogRecordLogDatumIntegerValue factory function for _BACnetLogRecordLogDatumIntegerValue
func NewBACnetLogRecordLogDatumIntegerValue(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, integerValue BACnetContextTagSignedInteger, tagNumber uint8) *_BACnetLogRecordLogDatumIntegerValue {
	if integerValue == nil {
		panic("integerValue of type BACnetContextTagSignedInteger for BACnetLogRecordLogDatumIntegerValue must not be nil")
	}
	_result := &_BACnetLogRecordLogDatumIntegerValue{
		BACnetLogRecordLogDatumContract: NewBACnetLogRecordLogDatum(openingTag, peekedTagHeader, closingTag, tagNumber),
		IntegerValue:                    integerValue,
	}
	_result.BACnetLogRecordLogDatumContract.(*_BACnetLogRecordLogDatum)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetLogRecordLogDatumIntegerValueBuilder is a builder for BACnetLogRecordLogDatumIntegerValue
type BACnetLogRecordLogDatumIntegerValueBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(integerValue BACnetContextTagSignedInteger) BACnetLogRecordLogDatumIntegerValueBuilder
	// WithIntegerValue adds IntegerValue (property field)
	WithIntegerValue(BACnetContextTagSignedInteger) BACnetLogRecordLogDatumIntegerValueBuilder
	// WithIntegerValueBuilder adds IntegerValue (property field) which is build by the builder
	WithIntegerValueBuilder(func(BACnetContextTagSignedIntegerBuilder) BACnetContextTagSignedIntegerBuilder) BACnetLogRecordLogDatumIntegerValueBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetLogRecordLogDatumBuilder
	// Build builds the BACnetLogRecordLogDatumIntegerValue or returns an error if something is wrong
	Build() (BACnetLogRecordLogDatumIntegerValue, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetLogRecordLogDatumIntegerValue
}

// NewBACnetLogRecordLogDatumIntegerValueBuilder() creates a BACnetLogRecordLogDatumIntegerValueBuilder
func NewBACnetLogRecordLogDatumIntegerValueBuilder() BACnetLogRecordLogDatumIntegerValueBuilder {
	return &_BACnetLogRecordLogDatumIntegerValueBuilder{_BACnetLogRecordLogDatumIntegerValue: new(_BACnetLogRecordLogDatumIntegerValue)}
}

type _BACnetLogRecordLogDatumIntegerValueBuilder struct {
	*_BACnetLogRecordLogDatumIntegerValue

	parentBuilder *_BACnetLogRecordLogDatumBuilder

	err *utils.MultiError
}

var _ (BACnetLogRecordLogDatumIntegerValueBuilder) = (*_BACnetLogRecordLogDatumIntegerValueBuilder)(nil)

func (b *_BACnetLogRecordLogDatumIntegerValueBuilder) setParent(contract BACnetLogRecordLogDatumContract) {
	b.BACnetLogRecordLogDatumContract = contract
	contract.(*_BACnetLogRecordLogDatum)._SubType = b._BACnetLogRecordLogDatumIntegerValue
}

func (b *_BACnetLogRecordLogDatumIntegerValueBuilder) WithMandatoryFields(integerValue BACnetContextTagSignedInteger) BACnetLogRecordLogDatumIntegerValueBuilder {
	return b.WithIntegerValue(integerValue)
}

func (b *_BACnetLogRecordLogDatumIntegerValueBuilder) WithIntegerValue(integerValue BACnetContextTagSignedInteger) BACnetLogRecordLogDatumIntegerValueBuilder {
	b.IntegerValue = integerValue
	return b
}

func (b *_BACnetLogRecordLogDatumIntegerValueBuilder) WithIntegerValueBuilder(builderSupplier func(BACnetContextTagSignedIntegerBuilder) BACnetContextTagSignedIntegerBuilder) BACnetLogRecordLogDatumIntegerValueBuilder {
	builder := builderSupplier(b.IntegerValue.CreateBACnetContextTagSignedIntegerBuilder())
	var err error
	b.IntegerValue, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagSignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetLogRecordLogDatumIntegerValueBuilder) Build() (BACnetLogRecordLogDatumIntegerValue, error) {
	if b.IntegerValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'integerValue' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetLogRecordLogDatumIntegerValue.deepCopy(), nil
}

func (b *_BACnetLogRecordLogDatumIntegerValueBuilder) MustBuild() BACnetLogRecordLogDatumIntegerValue {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetLogRecordLogDatumIntegerValueBuilder) Done() BACnetLogRecordLogDatumBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetLogRecordLogDatumBuilder().(*_BACnetLogRecordLogDatumBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetLogRecordLogDatumIntegerValueBuilder) buildForBACnetLogRecordLogDatum() (BACnetLogRecordLogDatum, error) {
	return b.Build()
}

func (b *_BACnetLogRecordLogDatumIntegerValueBuilder) DeepCopy() any {
	_copy := b.CreateBACnetLogRecordLogDatumIntegerValueBuilder().(*_BACnetLogRecordLogDatumIntegerValueBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetLogRecordLogDatumIntegerValueBuilder creates a BACnetLogRecordLogDatumIntegerValueBuilder
func (b *_BACnetLogRecordLogDatumIntegerValue) CreateBACnetLogRecordLogDatumIntegerValueBuilder() BACnetLogRecordLogDatumIntegerValueBuilder {
	if b == nil {
		return NewBACnetLogRecordLogDatumIntegerValueBuilder()
	}
	return &_BACnetLogRecordLogDatumIntegerValueBuilder{_BACnetLogRecordLogDatumIntegerValue: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetLogRecordLogDatumIntegerValue) GetParent() BACnetLogRecordLogDatumContract {
	return m.BACnetLogRecordLogDatumContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetLogRecordLogDatumIntegerValue) GetIntegerValue() BACnetContextTagSignedInteger {
	return m.IntegerValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetLogRecordLogDatumIntegerValue(structType any) BACnetLogRecordLogDatumIntegerValue {
	if casted, ok := structType.(BACnetLogRecordLogDatumIntegerValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetLogRecordLogDatumIntegerValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetLogRecordLogDatumIntegerValue) GetTypeName() string {
	return "BACnetLogRecordLogDatumIntegerValue"
}

func (m *_BACnetLogRecordLogDatumIntegerValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetLogRecordLogDatumContract.(*_BACnetLogRecordLogDatum).getLengthInBits(ctx))

	// Simple field (integerValue)
	lengthInBits += m.IntegerValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetLogRecordLogDatumIntegerValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetLogRecordLogDatumIntegerValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetLogRecordLogDatum, tagNumber uint8) (__bACnetLogRecordLogDatumIntegerValue BACnetLogRecordLogDatumIntegerValue, err error) {
	m.BACnetLogRecordLogDatumContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetLogRecordLogDatumIntegerValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetLogRecordLogDatumIntegerValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	integerValue, err := ReadSimpleField[BACnetContextTagSignedInteger](ctx, "integerValue", ReadComplex[BACnetContextTagSignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagSignedInteger]((uint8)(uint8(5)), (BACnetDataType)(BACnetDataType_SIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'integerValue' field"))
	}
	m.IntegerValue = integerValue

	if closeErr := readBuffer.CloseContext("BACnetLogRecordLogDatumIntegerValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetLogRecordLogDatumIntegerValue")
	}

	return m, nil
}

func (m *_BACnetLogRecordLogDatumIntegerValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetLogRecordLogDatumIntegerValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetLogRecordLogDatumIntegerValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetLogRecordLogDatumIntegerValue")
		}

		if err := WriteSimpleField[BACnetContextTagSignedInteger](ctx, "integerValue", m.GetIntegerValue(), WriteComplex[BACnetContextTagSignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'integerValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetLogRecordLogDatumIntegerValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetLogRecordLogDatumIntegerValue")
		}
		return nil
	}
	return m.BACnetLogRecordLogDatumContract.(*_BACnetLogRecordLogDatum).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetLogRecordLogDatumIntegerValue) IsBACnetLogRecordLogDatumIntegerValue() {}

func (m *_BACnetLogRecordLogDatumIntegerValue) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetLogRecordLogDatumIntegerValue) deepCopy() *_BACnetLogRecordLogDatumIntegerValue {
	if m == nil {
		return nil
	}
	_BACnetLogRecordLogDatumIntegerValueCopy := &_BACnetLogRecordLogDatumIntegerValue{
		m.BACnetLogRecordLogDatumContract.(*_BACnetLogRecordLogDatum).deepCopy(),
		utils.DeepCopy[BACnetContextTagSignedInteger](m.IntegerValue),
	}
	m.BACnetLogRecordLogDatumContract.(*_BACnetLogRecordLogDatum)._SubType = m
	return _BACnetLogRecordLogDatumIntegerValueCopy
}

func (m *_BACnetLogRecordLogDatumIntegerValue) String() string {
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
