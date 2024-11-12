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

// BACnetFaultParameterFaultExtendedParametersEntryDouble is the corresponding interface of BACnetFaultParameterFaultExtendedParametersEntryDouble
type BACnetFaultParameterFaultExtendedParametersEntryDouble interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetFaultParameterFaultExtendedParametersEntry
	// GetDoubleValue returns DoubleValue (property field)
	GetDoubleValue() BACnetApplicationTagDouble
	// IsBACnetFaultParameterFaultExtendedParametersEntryDouble is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetFaultParameterFaultExtendedParametersEntryDouble()
	// CreateBuilder creates a BACnetFaultParameterFaultExtendedParametersEntryDoubleBuilder
	CreateBACnetFaultParameterFaultExtendedParametersEntryDoubleBuilder() BACnetFaultParameterFaultExtendedParametersEntryDoubleBuilder
}

// _BACnetFaultParameterFaultExtendedParametersEntryDouble is the data-structure of this message
type _BACnetFaultParameterFaultExtendedParametersEntryDouble struct {
	BACnetFaultParameterFaultExtendedParametersEntryContract
	DoubleValue BACnetApplicationTagDouble
}

var _ BACnetFaultParameterFaultExtendedParametersEntryDouble = (*_BACnetFaultParameterFaultExtendedParametersEntryDouble)(nil)
var _ BACnetFaultParameterFaultExtendedParametersEntryRequirements = (*_BACnetFaultParameterFaultExtendedParametersEntryDouble)(nil)

// NewBACnetFaultParameterFaultExtendedParametersEntryDouble factory function for _BACnetFaultParameterFaultExtendedParametersEntryDouble
func NewBACnetFaultParameterFaultExtendedParametersEntryDouble(peekedTagHeader BACnetTagHeader, doubleValue BACnetApplicationTagDouble) *_BACnetFaultParameterFaultExtendedParametersEntryDouble {
	if doubleValue == nil {
		panic("doubleValue of type BACnetApplicationTagDouble for BACnetFaultParameterFaultExtendedParametersEntryDouble must not be nil")
	}
	_result := &_BACnetFaultParameterFaultExtendedParametersEntryDouble{
		BACnetFaultParameterFaultExtendedParametersEntryContract: NewBACnetFaultParameterFaultExtendedParametersEntry(peekedTagHeader),
		DoubleValue: doubleValue,
	}
	_result.BACnetFaultParameterFaultExtendedParametersEntryContract.(*_BACnetFaultParameterFaultExtendedParametersEntry)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetFaultParameterFaultExtendedParametersEntryDoubleBuilder is a builder for BACnetFaultParameterFaultExtendedParametersEntryDouble
type BACnetFaultParameterFaultExtendedParametersEntryDoubleBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(doubleValue BACnetApplicationTagDouble) BACnetFaultParameterFaultExtendedParametersEntryDoubleBuilder
	// WithDoubleValue adds DoubleValue (property field)
	WithDoubleValue(BACnetApplicationTagDouble) BACnetFaultParameterFaultExtendedParametersEntryDoubleBuilder
	// WithDoubleValueBuilder adds DoubleValue (property field) which is build by the builder
	WithDoubleValueBuilder(func(BACnetApplicationTagDoubleBuilder) BACnetApplicationTagDoubleBuilder) BACnetFaultParameterFaultExtendedParametersEntryDoubleBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetFaultParameterFaultExtendedParametersEntryBuilder
	// Build builds the BACnetFaultParameterFaultExtendedParametersEntryDouble or returns an error if something is wrong
	Build() (BACnetFaultParameterFaultExtendedParametersEntryDouble, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetFaultParameterFaultExtendedParametersEntryDouble
}

// NewBACnetFaultParameterFaultExtendedParametersEntryDoubleBuilder() creates a BACnetFaultParameterFaultExtendedParametersEntryDoubleBuilder
func NewBACnetFaultParameterFaultExtendedParametersEntryDoubleBuilder() BACnetFaultParameterFaultExtendedParametersEntryDoubleBuilder {
	return &_BACnetFaultParameterFaultExtendedParametersEntryDoubleBuilder{_BACnetFaultParameterFaultExtendedParametersEntryDouble: new(_BACnetFaultParameterFaultExtendedParametersEntryDouble)}
}

type _BACnetFaultParameterFaultExtendedParametersEntryDoubleBuilder struct {
	*_BACnetFaultParameterFaultExtendedParametersEntryDouble

	parentBuilder *_BACnetFaultParameterFaultExtendedParametersEntryBuilder

	err *utils.MultiError
}

var _ (BACnetFaultParameterFaultExtendedParametersEntryDoubleBuilder) = (*_BACnetFaultParameterFaultExtendedParametersEntryDoubleBuilder)(nil)

func (b *_BACnetFaultParameterFaultExtendedParametersEntryDoubleBuilder) setParent(contract BACnetFaultParameterFaultExtendedParametersEntryContract) {
	b.BACnetFaultParameterFaultExtendedParametersEntryContract = contract
	contract.(*_BACnetFaultParameterFaultExtendedParametersEntry)._SubType = b._BACnetFaultParameterFaultExtendedParametersEntryDouble
}

func (b *_BACnetFaultParameterFaultExtendedParametersEntryDoubleBuilder) WithMandatoryFields(doubleValue BACnetApplicationTagDouble) BACnetFaultParameterFaultExtendedParametersEntryDoubleBuilder {
	return b.WithDoubleValue(doubleValue)
}

func (b *_BACnetFaultParameterFaultExtendedParametersEntryDoubleBuilder) WithDoubleValue(doubleValue BACnetApplicationTagDouble) BACnetFaultParameterFaultExtendedParametersEntryDoubleBuilder {
	b.DoubleValue = doubleValue
	return b
}

func (b *_BACnetFaultParameterFaultExtendedParametersEntryDoubleBuilder) WithDoubleValueBuilder(builderSupplier func(BACnetApplicationTagDoubleBuilder) BACnetApplicationTagDoubleBuilder) BACnetFaultParameterFaultExtendedParametersEntryDoubleBuilder {
	builder := builderSupplier(b.DoubleValue.CreateBACnetApplicationTagDoubleBuilder())
	var err error
	b.DoubleValue, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagDoubleBuilder failed"))
	}
	return b
}

func (b *_BACnetFaultParameterFaultExtendedParametersEntryDoubleBuilder) Build() (BACnetFaultParameterFaultExtendedParametersEntryDouble, error) {
	if b.DoubleValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'doubleValue' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetFaultParameterFaultExtendedParametersEntryDouble.deepCopy(), nil
}

func (b *_BACnetFaultParameterFaultExtendedParametersEntryDoubleBuilder) MustBuild() BACnetFaultParameterFaultExtendedParametersEntryDouble {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetFaultParameterFaultExtendedParametersEntryDoubleBuilder) Done() BACnetFaultParameterFaultExtendedParametersEntryBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetFaultParameterFaultExtendedParametersEntryBuilder().(*_BACnetFaultParameterFaultExtendedParametersEntryBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetFaultParameterFaultExtendedParametersEntryDoubleBuilder) buildForBACnetFaultParameterFaultExtendedParametersEntry() (BACnetFaultParameterFaultExtendedParametersEntry, error) {
	return b.Build()
}

func (b *_BACnetFaultParameterFaultExtendedParametersEntryDoubleBuilder) DeepCopy() any {
	_copy := b.CreateBACnetFaultParameterFaultExtendedParametersEntryDoubleBuilder().(*_BACnetFaultParameterFaultExtendedParametersEntryDoubleBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetFaultParameterFaultExtendedParametersEntryDoubleBuilder creates a BACnetFaultParameterFaultExtendedParametersEntryDoubleBuilder
func (b *_BACnetFaultParameterFaultExtendedParametersEntryDouble) CreateBACnetFaultParameterFaultExtendedParametersEntryDoubleBuilder() BACnetFaultParameterFaultExtendedParametersEntryDoubleBuilder {
	if b == nil {
		return NewBACnetFaultParameterFaultExtendedParametersEntryDoubleBuilder()
	}
	return &_BACnetFaultParameterFaultExtendedParametersEntryDoubleBuilder{_BACnetFaultParameterFaultExtendedParametersEntryDouble: b.deepCopy()}
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

func (m *_BACnetFaultParameterFaultExtendedParametersEntryDouble) GetParent() BACnetFaultParameterFaultExtendedParametersEntryContract {
	return m.BACnetFaultParameterFaultExtendedParametersEntryContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetFaultParameterFaultExtendedParametersEntryDouble) GetDoubleValue() BACnetApplicationTagDouble {
	return m.DoubleValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetFaultParameterFaultExtendedParametersEntryDouble(structType any) BACnetFaultParameterFaultExtendedParametersEntryDouble {
	if casted, ok := structType.(BACnetFaultParameterFaultExtendedParametersEntryDouble); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetFaultParameterFaultExtendedParametersEntryDouble); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryDouble) GetTypeName() string {
	return "BACnetFaultParameterFaultExtendedParametersEntryDouble"
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryDouble) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetFaultParameterFaultExtendedParametersEntryContract.(*_BACnetFaultParameterFaultExtendedParametersEntry).getLengthInBits(ctx))

	// Simple field (doubleValue)
	lengthInBits += m.DoubleValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryDouble) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryDouble) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetFaultParameterFaultExtendedParametersEntry) (__bACnetFaultParameterFaultExtendedParametersEntryDouble BACnetFaultParameterFaultExtendedParametersEntryDouble, err error) {
	m.BACnetFaultParameterFaultExtendedParametersEntryContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetFaultParameterFaultExtendedParametersEntryDouble"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetFaultParameterFaultExtendedParametersEntryDouble")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	doubleValue, err := ReadSimpleField[BACnetApplicationTagDouble](ctx, "doubleValue", ReadComplex[BACnetApplicationTagDouble](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagDouble](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'doubleValue' field"))
	}
	m.DoubleValue = doubleValue

	if closeErr := readBuffer.CloseContext("BACnetFaultParameterFaultExtendedParametersEntryDouble"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetFaultParameterFaultExtendedParametersEntryDouble")
	}

	return m, nil
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryDouble) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryDouble) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetFaultParameterFaultExtendedParametersEntryDouble"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetFaultParameterFaultExtendedParametersEntryDouble")
		}

		if err := WriteSimpleField[BACnetApplicationTagDouble](ctx, "doubleValue", m.GetDoubleValue(), WriteComplex[BACnetApplicationTagDouble](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'doubleValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetFaultParameterFaultExtendedParametersEntryDouble"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetFaultParameterFaultExtendedParametersEntryDouble")
		}
		return nil
	}
	return m.BACnetFaultParameterFaultExtendedParametersEntryContract.(*_BACnetFaultParameterFaultExtendedParametersEntry).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryDouble) IsBACnetFaultParameterFaultExtendedParametersEntryDouble() {
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryDouble) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryDouble) deepCopy() *_BACnetFaultParameterFaultExtendedParametersEntryDouble {
	if m == nil {
		return nil
	}
	_BACnetFaultParameterFaultExtendedParametersEntryDoubleCopy := &_BACnetFaultParameterFaultExtendedParametersEntryDouble{
		m.BACnetFaultParameterFaultExtendedParametersEntryContract.(*_BACnetFaultParameterFaultExtendedParametersEntry).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagDouble](m.DoubleValue),
	}
	m.BACnetFaultParameterFaultExtendedParametersEntryContract.(*_BACnetFaultParameterFaultExtendedParametersEntry)._SubType = m
	return _BACnetFaultParameterFaultExtendedParametersEntryDoubleCopy
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryDouble) String() string {
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
