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

// BACnetFaultParameterFaultOutOfRangeMinNormalValueReal is the corresponding interface of BACnetFaultParameterFaultOutOfRangeMinNormalValueReal
type BACnetFaultParameterFaultOutOfRangeMinNormalValueReal interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetFaultParameterFaultOutOfRangeMinNormalValue
	// GetRealValue returns RealValue (property field)
	GetRealValue() BACnetApplicationTagReal
	// IsBACnetFaultParameterFaultOutOfRangeMinNormalValueReal is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetFaultParameterFaultOutOfRangeMinNormalValueReal()
	// CreateBuilder creates a BACnetFaultParameterFaultOutOfRangeMinNormalValueRealBuilder
	CreateBACnetFaultParameterFaultOutOfRangeMinNormalValueRealBuilder() BACnetFaultParameterFaultOutOfRangeMinNormalValueRealBuilder
}

// _BACnetFaultParameterFaultOutOfRangeMinNormalValueReal is the data-structure of this message
type _BACnetFaultParameterFaultOutOfRangeMinNormalValueReal struct {
	BACnetFaultParameterFaultOutOfRangeMinNormalValueContract
	RealValue BACnetApplicationTagReal
}

var _ BACnetFaultParameterFaultOutOfRangeMinNormalValueReal = (*_BACnetFaultParameterFaultOutOfRangeMinNormalValueReal)(nil)
var _ BACnetFaultParameterFaultOutOfRangeMinNormalValueRequirements = (*_BACnetFaultParameterFaultOutOfRangeMinNormalValueReal)(nil)

// NewBACnetFaultParameterFaultOutOfRangeMinNormalValueReal factory function for _BACnetFaultParameterFaultOutOfRangeMinNormalValueReal
func NewBACnetFaultParameterFaultOutOfRangeMinNormalValueReal(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, realValue BACnetApplicationTagReal, tagNumber uint8) *_BACnetFaultParameterFaultOutOfRangeMinNormalValueReal {
	if realValue == nil {
		panic("realValue of type BACnetApplicationTagReal for BACnetFaultParameterFaultOutOfRangeMinNormalValueReal must not be nil")
	}
	_result := &_BACnetFaultParameterFaultOutOfRangeMinNormalValueReal{
		BACnetFaultParameterFaultOutOfRangeMinNormalValueContract: NewBACnetFaultParameterFaultOutOfRangeMinNormalValue(openingTag, peekedTagHeader, closingTag, tagNumber),
		RealValue: realValue,
	}
	_result.BACnetFaultParameterFaultOutOfRangeMinNormalValueContract.(*_BACnetFaultParameterFaultOutOfRangeMinNormalValue)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetFaultParameterFaultOutOfRangeMinNormalValueRealBuilder is a builder for BACnetFaultParameterFaultOutOfRangeMinNormalValueReal
type BACnetFaultParameterFaultOutOfRangeMinNormalValueRealBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(realValue BACnetApplicationTagReal) BACnetFaultParameterFaultOutOfRangeMinNormalValueRealBuilder
	// WithRealValue adds RealValue (property field)
	WithRealValue(BACnetApplicationTagReal) BACnetFaultParameterFaultOutOfRangeMinNormalValueRealBuilder
	// WithRealValueBuilder adds RealValue (property field) which is build by the builder
	WithRealValueBuilder(func(BACnetApplicationTagRealBuilder) BACnetApplicationTagRealBuilder) BACnetFaultParameterFaultOutOfRangeMinNormalValueRealBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetFaultParameterFaultOutOfRangeMinNormalValueBuilder
	// Build builds the BACnetFaultParameterFaultOutOfRangeMinNormalValueReal or returns an error if something is wrong
	Build() (BACnetFaultParameterFaultOutOfRangeMinNormalValueReal, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetFaultParameterFaultOutOfRangeMinNormalValueReal
}

// NewBACnetFaultParameterFaultOutOfRangeMinNormalValueRealBuilder() creates a BACnetFaultParameterFaultOutOfRangeMinNormalValueRealBuilder
func NewBACnetFaultParameterFaultOutOfRangeMinNormalValueRealBuilder() BACnetFaultParameterFaultOutOfRangeMinNormalValueRealBuilder {
	return &_BACnetFaultParameterFaultOutOfRangeMinNormalValueRealBuilder{_BACnetFaultParameterFaultOutOfRangeMinNormalValueReal: new(_BACnetFaultParameterFaultOutOfRangeMinNormalValueReal)}
}

type _BACnetFaultParameterFaultOutOfRangeMinNormalValueRealBuilder struct {
	*_BACnetFaultParameterFaultOutOfRangeMinNormalValueReal

	parentBuilder *_BACnetFaultParameterFaultOutOfRangeMinNormalValueBuilder

	err *utils.MultiError
}

var _ (BACnetFaultParameterFaultOutOfRangeMinNormalValueRealBuilder) = (*_BACnetFaultParameterFaultOutOfRangeMinNormalValueRealBuilder)(nil)

func (b *_BACnetFaultParameterFaultOutOfRangeMinNormalValueRealBuilder) setParent(contract BACnetFaultParameterFaultOutOfRangeMinNormalValueContract) {
	b.BACnetFaultParameterFaultOutOfRangeMinNormalValueContract = contract
	contract.(*_BACnetFaultParameterFaultOutOfRangeMinNormalValue)._SubType = b._BACnetFaultParameterFaultOutOfRangeMinNormalValueReal
}

func (b *_BACnetFaultParameterFaultOutOfRangeMinNormalValueRealBuilder) WithMandatoryFields(realValue BACnetApplicationTagReal) BACnetFaultParameterFaultOutOfRangeMinNormalValueRealBuilder {
	return b.WithRealValue(realValue)
}

func (b *_BACnetFaultParameterFaultOutOfRangeMinNormalValueRealBuilder) WithRealValue(realValue BACnetApplicationTagReal) BACnetFaultParameterFaultOutOfRangeMinNormalValueRealBuilder {
	b.RealValue = realValue
	return b
}

func (b *_BACnetFaultParameterFaultOutOfRangeMinNormalValueRealBuilder) WithRealValueBuilder(builderSupplier func(BACnetApplicationTagRealBuilder) BACnetApplicationTagRealBuilder) BACnetFaultParameterFaultOutOfRangeMinNormalValueRealBuilder {
	builder := builderSupplier(b.RealValue.CreateBACnetApplicationTagRealBuilder())
	var err error
	b.RealValue, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagRealBuilder failed"))
	}
	return b
}

func (b *_BACnetFaultParameterFaultOutOfRangeMinNormalValueRealBuilder) Build() (BACnetFaultParameterFaultOutOfRangeMinNormalValueReal, error) {
	if b.RealValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'realValue' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetFaultParameterFaultOutOfRangeMinNormalValueReal.deepCopy(), nil
}

func (b *_BACnetFaultParameterFaultOutOfRangeMinNormalValueRealBuilder) MustBuild() BACnetFaultParameterFaultOutOfRangeMinNormalValueReal {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetFaultParameterFaultOutOfRangeMinNormalValueRealBuilder) Done() BACnetFaultParameterFaultOutOfRangeMinNormalValueBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetFaultParameterFaultOutOfRangeMinNormalValueBuilder().(*_BACnetFaultParameterFaultOutOfRangeMinNormalValueBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetFaultParameterFaultOutOfRangeMinNormalValueRealBuilder) buildForBACnetFaultParameterFaultOutOfRangeMinNormalValue() (BACnetFaultParameterFaultOutOfRangeMinNormalValue, error) {
	return b.Build()
}

func (b *_BACnetFaultParameterFaultOutOfRangeMinNormalValueRealBuilder) DeepCopy() any {
	_copy := b.CreateBACnetFaultParameterFaultOutOfRangeMinNormalValueRealBuilder().(*_BACnetFaultParameterFaultOutOfRangeMinNormalValueRealBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetFaultParameterFaultOutOfRangeMinNormalValueRealBuilder creates a BACnetFaultParameterFaultOutOfRangeMinNormalValueRealBuilder
func (b *_BACnetFaultParameterFaultOutOfRangeMinNormalValueReal) CreateBACnetFaultParameterFaultOutOfRangeMinNormalValueRealBuilder() BACnetFaultParameterFaultOutOfRangeMinNormalValueRealBuilder {
	if b == nil {
		return NewBACnetFaultParameterFaultOutOfRangeMinNormalValueRealBuilder()
	}
	return &_BACnetFaultParameterFaultOutOfRangeMinNormalValueRealBuilder{_BACnetFaultParameterFaultOutOfRangeMinNormalValueReal: b.deepCopy()}
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

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueReal) GetParent() BACnetFaultParameterFaultOutOfRangeMinNormalValueContract {
	return m.BACnetFaultParameterFaultOutOfRangeMinNormalValueContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueReal) GetRealValue() BACnetApplicationTagReal {
	return m.RealValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetFaultParameterFaultOutOfRangeMinNormalValueReal(structType any) BACnetFaultParameterFaultOutOfRangeMinNormalValueReal {
	if casted, ok := structType.(BACnetFaultParameterFaultOutOfRangeMinNormalValueReal); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetFaultParameterFaultOutOfRangeMinNormalValueReal); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueReal) GetTypeName() string {
	return "BACnetFaultParameterFaultOutOfRangeMinNormalValueReal"
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueReal) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetFaultParameterFaultOutOfRangeMinNormalValueContract.(*_BACnetFaultParameterFaultOutOfRangeMinNormalValue).getLengthInBits(ctx))

	// Simple field (realValue)
	lengthInBits += m.RealValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueReal) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueReal) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetFaultParameterFaultOutOfRangeMinNormalValue, tagNumber uint8) (__bACnetFaultParameterFaultOutOfRangeMinNormalValueReal BACnetFaultParameterFaultOutOfRangeMinNormalValueReal, err error) {
	m.BACnetFaultParameterFaultOutOfRangeMinNormalValueContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetFaultParameterFaultOutOfRangeMinNormalValueReal"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetFaultParameterFaultOutOfRangeMinNormalValueReal")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	realValue, err := ReadSimpleField[BACnetApplicationTagReal](ctx, "realValue", ReadComplex[BACnetApplicationTagReal](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagReal](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'realValue' field"))
	}
	m.RealValue = realValue

	if closeErr := readBuffer.CloseContext("BACnetFaultParameterFaultOutOfRangeMinNormalValueReal"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetFaultParameterFaultOutOfRangeMinNormalValueReal")
	}

	return m, nil
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueReal) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueReal) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetFaultParameterFaultOutOfRangeMinNormalValueReal"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetFaultParameterFaultOutOfRangeMinNormalValueReal")
		}

		if err := WriteSimpleField[BACnetApplicationTagReal](ctx, "realValue", m.GetRealValue(), WriteComplex[BACnetApplicationTagReal](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'realValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetFaultParameterFaultOutOfRangeMinNormalValueReal"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetFaultParameterFaultOutOfRangeMinNormalValueReal")
		}
		return nil
	}
	return m.BACnetFaultParameterFaultOutOfRangeMinNormalValueContract.(*_BACnetFaultParameterFaultOutOfRangeMinNormalValue).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueReal) IsBACnetFaultParameterFaultOutOfRangeMinNormalValueReal() {
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueReal) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueReal) deepCopy() *_BACnetFaultParameterFaultOutOfRangeMinNormalValueReal {
	if m == nil {
		return nil
	}
	_BACnetFaultParameterFaultOutOfRangeMinNormalValueRealCopy := &_BACnetFaultParameterFaultOutOfRangeMinNormalValueReal{
		m.BACnetFaultParameterFaultOutOfRangeMinNormalValueContract.(*_BACnetFaultParameterFaultOutOfRangeMinNormalValue).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagReal](m.RealValue),
	}
	m.BACnetFaultParameterFaultOutOfRangeMinNormalValueContract.(*_BACnetFaultParameterFaultOutOfRangeMinNormalValue)._SubType = m
	return _BACnetFaultParameterFaultOutOfRangeMinNormalValueRealCopy
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueReal) String() string {
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
