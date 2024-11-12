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

// BACnetFaultParameterFaultExtendedParametersEntryInteger is the corresponding interface of BACnetFaultParameterFaultExtendedParametersEntryInteger
type BACnetFaultParameterFaultExtendedParametersEntryInteger interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetFaultParameterFaultExtendedParametersEntry
	// GetIntegerValue returns IntegerValue (property field)
	GetIntegerValue() BACnetApplicationTagSignedInteger
	// IsBACnetFaultParameterFaultExtendedParametersEntryInteger is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetFaultParameterFaultExtendedParametersEntryInteger()
	// CreateBuilder creates a BACnetFaultParameterFaultExtendedParametersEntryIntegerBuilder
	CreateBACnetFaultParameterFaultExtendedParametersEntryIntegerBuilder() BACnetFaultParameterFaultExtendedParametersEntryIntegerBuilder
}

// _BACnetFaultParameterFaultExtendedParametersEntryInteger is the data-structure of this message
type _BACnetFaultParameterFaultExtendedParametersEntryInteger struct {
	BACnetFaultParameterFaultExtendedParametersEntryContract
	IntegerValue BACnetApplicationTagSignedInteger
}

var _ BACnetFaultParameterFaultExtendedParametersEntryInteger = (*_BACnetFaultParameterFaultExtendedParametersEntryInteger)(nil)
var _ BACnetFaultParameterFaultExtendedParametersEntryRequirements = (*_BACnetFaultParameterFaultExtendedParametersEntryInteger)(nil)

// NewBACnetFaultParameterFaultExtendedParametersEntryInteger factory function for _BACnetFaultParameterFaultExtendedParametersEntryInteger
func NewBACnetFaultParameterFaultExtendedParametersEntryInteger(peekedTagHeader BACnetTagHeader, integerValue BACnetApplicationTagSignedInteger) *_BACnetFaultParameterFaultExtendedParametersEntryInteger {
	if integerValue == nil {
		panic("integerValue of type BACnetApplicationTagSignedInteger for BACnetFaultParameterFaultExtendedParametersEntryInteger must not be nil")
	}
	_result := &_BACnetFaultParameterFaultExtendedParametersEntryInteger{
		BACnetFaultParameterFaultExtendedParametersEntryContract: NewBACnetFaultParameterFaultExtendedParametersEntry(peekedTagHeader),
		IntegerValue: integerValue,
	}
	_result.BACnetFaultParameterFaultExtendedParametersEntryContract.(*_BACnetFaultParameterFaultExtendedParametersEntry)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetFaultParameterFaultExtendedParametersEntryIntegerBuilder is a builder for BACnetFaultParameterFaultExtendedParametersEntryInteger
type BACnetFaultParameterFaultExtendedParametersEntryIntegerBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(integerValue BACnetApplicationTagSignedInteger) BACnetFaultParameterFaultExtendedParametersEntryIntegerBuilder
	// WithIntegerValue adds IntegerValue (property field)
	WithIntegerValue(BACnetApplicationTagSignedInteger) BACnetFaultParameterFaultExtendedParametersEntryIntegerBuilder
	// WithIntegerValueBuilder adds IntegerValue (property field) which is build by the builder
	WithIntegerValueBuilder(func(BACnetApplicationTagSignedIntegerBuilder) BACnetApplicationTagSignedIntegerBuilder) BACnetFaultParameterFaultExtendedParametersEntryIntegerBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetFaultParameterFaultExtendedParametersEntryBuilder
	// Build builds the BACnetFaultParameterFaultExtendedParametersEntryInteger or returns an error if something is wrong
	Build() (BACnetFaultParameterFaultExtendedParametersEntryInteger, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetFaultParameterFaultExtendedParametersEntryInteger
}

// NewBACnetFaultParameterFaultExtendedParametersEntryIntegerBuilder() creates a BACnetFaultParameterFaultExtendedParametersEntryIntegerBuilder
func NewBACnetFaultParameterFaultExtendedParametersEntryIntegerBuilder() BACnetFaultParameterFaultExtendedParametersEntryIntegerBuilder {
	return &_BACnetFaultParameterFaultExtendedParametersEntryIntegerBuilder{_BACnetFaultParameterFaultExtendedParametersEntryInteger: new(_BACnetFaultParameterFaultExtendedParametersEntryInteger)}
}

type _BACnetFaultParameterFaultExtendedParametersEntryIntegerBuilder struct {
	*_BACnetFaultParameterFaultExtendedParametersEntryInteger

	parentBuilder *_BACnetFaultParameterFaultExtendedParametersEntryBuilder

	err *utils.MultiError
}

var _ (BACnetFaultParameterFaultExtendedParametersEntryIntegerBuilder) = (*_BACnetFaultParameterFaultExtendedParametersEntryIntegerBuilder)(nil)

func (b *_BACnetFaultParameterFaultExtendedParametersEntryIntegerBuilder) setParent(contract BACnetFaultParameterFaultExtendedParametersEntryContract) {
	b.BACnetFaultParameterFaultExtendedParametersEntryContract = contract
	contract.(*_BACnetFaultParameterFaultExtendedParametersEntry)._SubType = b._BACnetFaultParameterFaultExtendedParametersEntryInteger
}

func (b *_BACnetFaultParameterFaultExtendedParametersEntryIntegerBuilder) WithMandatoryFields(integerValue BACnetApplicationTagSignedInteger) BACnetFaultParameterFaultExtendedParametersEntryIntegerBuilder {
	return b.WithIntegerValue(integerValue)
}

func (b *_BACnetFaultParameterFaultExtendedParametersEntryIntegerBuilder) WithIntegerValue(integerValue BACnetApplicationTagSignedInteger) BACnetFaultParameterFaultExtendedParametersEntryIntegerBuilder {
	b.IntegerValue = integerValue
	return b
}

func (b *_BACnetFaultParameterFaultExtendedParametersEntryIntegerBuilder) WithIntegerValueBuilder(builderSupplier func(BACnetApplicationTagSignedIntegerBuilder) BACnetApplicationTagSignedIntegerBuilder) BACnetFaultParameterFaultExtendedParametersEntryIntegerBuilder {
	builder := builderSupplier(b.IntegerValue.CreateBACnetApplicationTagSignedIntegerBuilder())
	var err error
	b.IntegerValue, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagSignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetFaultParameterFaultExtendedParametersEntryIntegerBuilder) Build() (BACnetFaultParameterFaultExtendedParametersEntryInteger, error) {
	if b.IntegerValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'integerValue' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetFaultParameterFaultExtendedParametersEntryInteger.deepCopy(), nil
}

func (b *_BACnetFaultParameterFaultExtendedParametersEntryIntegerBuilder) MustBuild() BACnetFaultParameterFaultExtendedParametersEntryInteger {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetFaultParameterFaultExtendedParametersEntryIntegerBuilder) Done() BACnetFaultParameterFaultExtendedParametersEntryBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetFaultParameterFaultExtendedParametersEntryBuilder().(*_BACnetFaultParameterFaultExtendedParametersEntryBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetFaultParameterFaultExtendedParametersEntryIntegerBuilder) buildForBACnetFaultParameterFaultExtendedParametersEntry() (BACnetFaultParameterFaultExtendedParametersEntry, error) {
	return b.Build()
}

func (b *_BACnetFaultParameterFaultExtendedParametersEntryIntegerBuilder) DeepCopy() any {
	_copy := b.CreateBACnetFaultParameterFaultExtendedParametersEntryIntegerBuilder().(*_BACnetFaultParameterFaultExtendedParametersEntryIntegerBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetFaultParameterFaultExtendedParametersEntryIntegerBuilder creates a BACnetFaultParameterFaultExtendedParametersEntryIntegerBuilder
func (b *_BACnetFaultParameterFaultExtendedParametersEntryInteger) CreateBACnetFaultParameterFaultExtendedParametersEntryIntegerBuilder() BACnetFaultParameterFaultExtendedParametersEntryIntegerBuilder {
	if b == nil {
		return NewBACnetFaultParameterFaultExtendedParametersEntryIntegerBuilder()
	}
	return &_BACnetFaultParameterFaultExtendedParametersEntryIntegerBuilder{_BACnetFaultParameterFaultExtendedParametersEntryInteger: b.deepCopy()}
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

func (m *_BACnetFaultParameterFaultExtendedParametersEntryInteger) GetParent() BACnetFaultParameterFaultExtendedParametersEntryContract {
	return m.BACnetFaultParameterFaultExtendedParametersEntryContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetFaultParameterFaultExtendedParametersEntryInteger) GetIntegerValue() BACnetApplicationTagSignedInteger {
	return m.IntegerValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetFaultParameterFaultExtendedParametersEntryInteger(structType any) BACnetFaultParameterFaultExtendedParametersEntryInteger {
	if casted, ok := structType.(BACnetFaultParameterFaultExtendedParametersEntryInteger); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetFaultParameterFaultExtendedParametersEntryInteger); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryInteger) GetTypeName() string {
	return "BACnetFaultParameterFaultExtendedParametersEntryInteger"
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryInteger) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetFaultParameterFaultExtendedParametersEntryContract.(*_BACnetFaultParameterFaultExtendedParametersEntry).getLengthInBits(ctx))

	// Simple field (integerValue)
	lengthInBits += m.IntegerValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryInteger) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryInteger) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetFaultParameterFaultExtendedParametersEntry) (__bACnetFaultParameterFaultExtendedParametersEntryInteger BACnetFaultParameterFaultExtendedParametersEntryInteger, err error) {
	m.BACnetFaultParameterFaultExtendedParametersEntryContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetFaultParameterFaultExtendedParametersEntryInteger"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetFaultParameterFaultExtendedParametersEntryInteger")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	integerValue, err := ReadSimpleField[BACnetApplicationTagSignedInteger](ctx, "integerValue", ReadComplex[BACnetApplicationTagSignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagSignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'integerValue' field"))
	}
	m.IntegerValue = integerValue

	if closeErr := readBuffer.CloseContext("BACnetFaultParameterFaultExtendedParametersEntryInteger"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetFaultParameterFaultExtendedParametersEntryInteger")
	}

	return m, nil
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryInteger) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryInteger) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetFaultParameterFaultExtendedParametersEntryInteger"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetFaultParameterFaultExtendedParametersEntryInteger")
		}

		if err := WriteSimpleField[BACnetApplicationTagSignedInteger](ctx, "integerValue", m.GetIntegerValue(), WriteComplex[BACnetApplicationTagSignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'integerValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetFaultParameterFaultExtendedParametersEntryInteger"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetFaultParameterFaultExtendedParametersEntryInteger")
		}
		return nil
	}
	return m.BACnetFaultParameterFaultExtendedParametersEntryContract.(*_BACnetFaultParameterFaultExtendedParametersEntry).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryInteger) IsBACnetFaultParameterFaultExtendedParametersEntryInteger() {
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryInteger) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryInteger) deepCopy() *_BACnetFaultParameterFaultExtendedParametersEntryInteger {
	if m == nil {
		return nil
	}
	_BACnetFaultParameterFaultExtendedParametersEntryIntegerCopy := &_BACnetFaultParameterFaultExtendedParametersEntryInteger{
		m.BACnetFaultParameterFaultExtendedParametersEntryContract.(*_BACnetFaultParameterFaultExtendedParametersEntry).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagSignedInteger](m.IntegerValue),
	}
	m.BACnetFaultParameterFaultExtendedParametersEntryContract.(*_BACnetFaultParameterFaultExtendedParametersEntry)._SubType = m
	return _BACnetFaultParameterFaultExtendedParametersEntryIntegerCopy
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryInteger) String() string {
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
