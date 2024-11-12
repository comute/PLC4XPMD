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

// BACnetFaultParameterFaultExtendedParametersEntryReal is the corresponding interface of BACnetFaultParameterFaultExtendedParametersEntryReal
type BACnetFaultParameterFaultExtendedParametersEntryReal interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetFaultParameterFaultExtendedParametersEntry
	// GetRealValue returns RealValue (property field)
	GetRealValue() BACnetApplicationTagReal
	// IsBACnetFaultParameterFaultExtendedParametersEntryReal is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetFaultParameterFaultExtendedParametersEntryReal()
	// CreateBuilder creates a BACnetFaultParameterFaultExtendedParametersEntryRealBuilder
	CreateBACnetFaultParameterFaultExtendedParametersEntryRealBuilder() BACnetFaultParameterFaultExtendedParametersEntryRealBuilder
}

// _BACnetFaultParameterFaultExtendedParametersEntryReal is the data-structure of this message
type _BACnetFaultParameterFaultExtendedParametersEntryReal struct {
	BACnetFaultParameterFaultExtendedParametersEntryContract
	RealValue BACnetApplicationTagReal
}

var _ BACnetFaultParameterFaultExtendedParametersEntryReal = (*_BACnetFaultParameterFaultExtendedParametersEntryReal)(nil)
var _ BACnetFaultParameterFaultExtendedParametersEntryRequirements = (*_BACnetFaultParameterFaultExtendedParametersEntryReal)(nil)

// NewBACnetFaultParameterFaultExtendedParametersEntryReal factory function for _BACnetFaultParameterFaultExtendedParametersEntryReal
func NewBACnetFaultParameterFaultExtendedParametersEntryReal(peekedTagHeader BACnetTagHeader, realValue BACnetApplicationTagReal) *_BACnetFaultParameterFaultExtendedParametersEntryReal {
	if realValue == nil {
		panic("realValue of type BACnetApplicationTagReal for BACnetFaultParameterFaultExtendedParametersEntryReal must not be nil")
	}
	_result := &_BACnetFaultParameterFaultExtendedParametersEntryReal{
		BACnetFaultParameterFaultExtendedParametersEntryContract: NewBACnetFaultParameterFaultExtendedParametersEntry(peekedTagHeader),
		RealValue: realValue,
	}
	_result.BACnetFaultParameterFaultExtendedParametersEntryContract.(*_BACnetFaultParameterFaultExtendedParametersEntry)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetFaultParameterFaultExtendedParametersEntryRealBuilder is a builder for BACnetFaultParameterFaultExtendedParametersEntryReal
type BACnetFaultParameterFaultExtendedParametersEntryRealBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(realValue BACnetApplicationTagReal) BACnetFaultParameterFaultExtendedParametersEntryRealBuilder
	// WithRealValue adds RealValue (property field)
	WithRealValue(BACnetApplicationTagReal) BACnetFaultParameterFaultExtendedParametersEntryRealBuilder
	// WithRealValueBuilder adds RealValue (property field) which is build by the builder
	WithRealValueBuilder(func(BACnetApplicationTagRealBuilder) BACnetApplicationTagRealBuilder) BACnetFaultParameterFaultExtendedParametersEntryRealBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetFaultParameterFaultExtendedParametersEntryBuilder
	// Build builds the BACnetFaultParameterFaultExtendedParametersEntryReal or returns an error if something is wrong
	Build() (BACnetFaultParameterFaultExtendedParametersEntryReal, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetFaultParameterFaultExtendedParametersEntryReal
}

// NewBACnetFaultParameterFaultExtendedParametersEntryRealBuilder() creates a BACnetFaultParameterFaultExtendedParametersEntryRealBuilder
func NewBACnetFaultParameterFaultExtendedParametersEntryRealBuilder() BACnetFaultParameterFaultExtendedParametersEntryRealBuilder {
	return &_BACnetFaultParameterFaultExtendedParametersEntryRealBuilder{_BACnetFaultParameterFaultExtendedParametersEntryReal: new(_BACnetFaultParameterFaultExtendedParametersEntryReal)}
}

type _BACnetFaultParameterFaultExtendedParametersEntryRealBuilder struct {
	*_BACnetFaultParameterFaultExtendedParametersEntryReal

	parentBuilder *_BACnetFaultParameterFaultExtendedParametersEntryBuilder

	err *utils.MultiError
}

var _ (BACnetFaultParameterFaultExtendedParametersEntryRealBuilder) = (*_BACnetFaultParameterFaultExtendedParametersEntryRealBuilder)(nil)

func (b *_BACnetFaultParameterFaultExtendedParametersEntryRealBuilder) setParent(contract BACnetFaultParameterFaultExtendedParametersEntryContract) {
	b.BACnetFaultParameterFaultExtendedParametersEntryContract = contract
	contract.(*_BACnetFaultParameterFaultExtendedParametersEntry)._SubType = b._BACnetFaultParameterFaultExtendedParametersEntryReal
}

func (b *_BACnetFaultParameterFaultExtendedParametersEntryRealBuilder) WithMandatoryFields(realValue BACnetApplicationTagReal) BACnetFaultParameterFaultExtendedParametersEntryRealBuilder {
	return b.WithRealValue(realValue)
}

func (b *_BACnetFaultParameterFaultExtendedParametersEntryRealBuilder) WithRealValue(realValue BACnetApplicationTagReal) BACnetFaultParameterFaultExtendedParametersEntryRealBuilder {
	b.RealValue = realValue
	return b
}

func (b *_BACnetFaultParameterFaultExtendedParametersEntryRealBuilder) WithRealValueBuilder(builderSupplier func(BACnetApplicationTagRealBuilder) BACnetApplicationTagRealBuilder) BACnetFaultParameterFaultExtendedParametersEntryRealBuilder {
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

func (b *_BACnetFaultParameterFaultExtendedParametersEntryRealBuilder) Build() (BACnetFaultParameterFaultExtendedParametersEntryReal, error) {
	if b.RealValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'realValue' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetFaultParameterFaultExtendedParametersEntryReal.deepCopy(), nil
}

func (b *_BACnetFaultParameterFaultExtendedParametersEntryRealBuilder) MustBuild() BACnetFaultParameterFaultExtendedParametersEntryReal {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetFaultParameterFaultExtendedParametersEntryRealBuilder) Done() BACnetFaultParameterFaultExtendedParametersEntryBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetFaultParameterFaultExtendedParametersEntryBuilder().(*_BACnetFaultParameterFaultExtendedParametersEntryBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetFaultParameterFaultExtendedParametersEntryRealBuilder) buildForBACnetFaultParameterFaultExtendedParametersEntry() (BACnetFaultParameterFaultExtendedParametersEntry, error) {
	return b.Build()
}

func (b *_BACnetFaultParameterFaultExtendedParametersEntryRealBuilder) DeepCopy() any {
	_copy := b.CreateBACnetFaultParameterFaultExtendedParametersEntryRealBuilder().(*_BACnetFaultParameterFaultExtendedParametersEntryRealBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetFaultParameterFaultExtendedParametersEntryRealBuilder creates a BACnetFaultParameterFaultExtendedParametersEntryRealBuilder
func (b *_BACnetFaultParameterFaultExtendedParametersEntryReal) CreateBACnetFaultParameterFaultExtendedParametersEntryRealBuilder() BACnetFaultParameterFaultExtendedParametersEntryRealBuilder {
	if b == nil {
		return NewBACnetFaultParameterFaultExtendedParametersEntryRealBuilder()
	}
	return &_BACnetFaultParameterFaultExtendedParametersEntryRealBuilder{_BACnetFaultParameterFaultExtendedParametersEntryReal: b.deepCopy()}
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

func (m *_BACnetFaultParameterFaultExtendedParametersEntryReal) GetParent() BACnetFaultParameterFaultExtendedParametersEntryContract {
	return m.BACnetFaultParameterFaultExtendedParametersEntryContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetFaultParameterFaultExtendedParametersEntryReal) GetRealValue() BACnetApplicationTagReal {
	return m.RealValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetFaultParameterFaultExtendedParametersEntryReal(structType any) BACnetFaultParameterFaultExtendedParametersEntryReal {
	if casted, ok := structType.(BACnetFaultParameterFaultExtendedParametersEntryReal); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetFaultParameterFaultExtendedParametersEntryReal); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryReal) GetTypeName() string {
	return "BACnetFaultParameterFaultExtendedParametersEntryReal"
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryReal) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetFaultParameterFaultExtendedParametersEntryContract.(*_BACnetFaultParameterFaultExtendedParametersEntry).getLengthInBits(ctx))

	// Simple field (realValue)
	lengthInBits += m.RealValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryReal) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryReal) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetFaultParameterFaultExtendedParametersEntry) (__bACnetFaultParameterFaultExtendedParametersEntryReal BACnetFaultParameterFaultExtendedParametersEntryReal, err error) {
	m.BACnetFaultParameterFaultExtendedParametersEntryContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetFaultParameterFaultExtendedParametersEntryReal"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetFaultParameterFaultExtendedParametersEntryReal")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	realValue, err := ReadSimpleField[BACnetApplicationTagReal](ctx, "realValue", ReadComplex[BACnetApplicationTagReal](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagReal](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'realValue' field"))
	}
	m.RealValue = realValue

	if closeErr := readBuffer.CloseContext("BACnetFaultParameterFaultExtendedParametersEntryReal"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetFaultParameterFaultExtendedParametersEntryReal")
	}

	return m, nil
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryReal) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryReal) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetFaultParameterFaultExtendedParametersEntryReal"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetFaultParameterFaultExtendedParametersEntryReal")
		}

		if err := WriteSimpleField[BACnetApplicationTagReal](ctx, "realValue", m.GetRealValue(), WriteComplex[BACnetApplicationTagReal](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'realValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetFaultParameterFaultExtendedParametersEntryReal"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetFaultParameterFaultExtendedParametersEntryReal")
		}
		return nil
	}
	return m.BACnetFaultParameterFaultExtendedParametersEntryContract.(*_BACnetFaultParameterFaultExtendedParametersEntry).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryReal) IsBACnetFaultParameterFaultExtendedParametersEntryReal() {
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryReal) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryReal) deepCopy() *_BACnetFaultParameterFaultExtendedParametersEntryReal {
	if m == nil {
		return nil
	}
	_BACnetFaultParameterFaultExtendedParametersEntryRealCopy := &_BACnetFaultParameterFaultExtendedParametersEntryReal{
		m.BACnetFaultParameterFaultExtendedParametersEntryContract.(*_BACnetFaultParameterFaultExtendedParametersEntry).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagReal](m.RealValue),
	}
	m.BACnetFaultParameterFaultExtendedParametersEntryContract.(*_BACnetFaultParameterFaultExtendedParametersEntry)._SubType = m
	return _BACnetFaultParameterFaultExtendedParametersEntryRealCopy
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryReal) String() string {
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
