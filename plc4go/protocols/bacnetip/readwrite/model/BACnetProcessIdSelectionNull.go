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

// BACnetProcessIdSelectionNull is the corresponding interface of BACnetProcessIdSelectionNull
type BACnetProcessIdSelectionNull interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetProcessIdSelection
	// GetNullValue returns NullValue (property field)
	GetNullValue() BACnetApplicationTagNull
	// IsBACnetProcessIdSelectionNull is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetProcessIdSelectionNull()
	// CreateBuilder creates a BACnetProcessIdSelectionNullBuilder
	CreateBACnetProcessIdSelectionNullBuilder() BACnetProcessIdSelectionNullBuilder
}

// _BACnetProcessIdSelectionNull is the data-structure of this message
type _BACnetProcessIdSelectionNull struct {
	BACnetProcessIdSelectionContract
	NullValue BACnetApplicationTagNull
}

var _ BACnetProcessIdSelectionNull = (*_BACnetProcessIdSelectionNull)(nil)
var _ BACnetProcessIdSelectionRequirements = (*_BACnetProcessIdSelectionNull)(nil)

// NewBACnetProcessIdSelectionNull factory function for _BACnetProcessIdSelectionNull
func NewBACnetProcessIdSelectionNull(peekedTagHeader BACnetTagHeader, nullValue BACnetApplicationTagNull) *_BACnetProcessIdSelectionNull {
	if nullValue == nil {
		panic("nullValue of type BACnetApplicationTagNull for BACnetProcessIdSelectionNull must not be nil")
	}
	_result := &_BACnetProcessIdSelectionNull{
		BACnetProcessIdSelectionContract: NewBACnetProcessIdSelection(peekedTagHeader),
		NullValue:                        nullValue,
	}
	_result.BACnetProcessIdSelectionContract.(*_BACnetProcessIdSelection)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetProcessIdSelectionNullBuilder is a builder for BACnetProcessIdSelectionNull
type BACnetProcessIdSelectionNullBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(nullValue BACnetApplicationTagNull) BACnetProcessIdSelectionNullBuilder
	// WithNullValue adds NullValue (property field)
	WithNullValue(BACnetApplicationTagNull) BACnetProcessIdSelectionNullBuilder
	// WithNullValueBuilder adds NullValue (property field) which is build by the builder
	WithNullValueBuilder(func(BACnetApplicationTagNullBuilder) BACnetApplicationTagNullBuilder) BACnetProcessIdSelectionNullBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetProcessIdSelectionBuilder
	// Build builds the BACnetProcessIdSelectionNull or returns an error if something is wrong
	Build() (BACnetProcessIdSelectionNull, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetProcessIdSelectionNull
}

// NewBACnetProcessIdSelectionNullBuilder() creates a BACnetProcessIdSelectionNullBuilder
func NewBACnetProcessIdSelectionNullBuilder() BACnetProcessIdSelectionNullBuilder {
	return &_BACnetProcessIdSelectionNullBuilder{_BACnetProcessIdSelectionNull: new(_BACnetProcessIdSelectionNull)}
}

type _BACnetProcessIdSelectionNullBuilder struct {
	*_BACnetProcessIdSelectionNull

	parentBuilder *_BACnetProcessIdSelectionBuilder

	err *utils.MultiError
}

var _ (BACnetProcessIdSelectionNullBuilder) = (*_BACnetProcessIdSelectionNullBuilder)(nil)

func (b *_BACnetProcessIdSelectionNullBuilder) setParent(contract BACnetProcessIdSelectionContract) {
	b.BACnetProcessIdSelectionContract = contract
	contract.(*_BACnetProcessIdSelection)._SubType = b._BACnetProcessIdSelectionNull
}

func (b *_BACnetProcessIdSelectionNullBuilder) WithMandatoryFields(nullValue BACnetApplicationTagNull) BACnetProcessIdSelectionNullBuilder {
	return b.WithNullValue(nullValue)
}

func (b *_BACnetProcessIdSelectionNullBuilder) WithNullValue(nullValue BACnetApplicationTagNull) BACnetProcessIdSelectionNullBuilder {
	b.NullValue = nullValue
	return b
}

func (b *_BACnetProcessIdSelectionNullBuilder) WithNullValueBuilder(builderSupplier func(BACnetApplicationTagNullBuilder) BACnetApplicationTagNullBuilder) BACnetProcessIdSelectionNullBuilder {
	builder := builderSupplier(b.NullValue.CreateBACnetApplicationTagNullBuilder())
	var err error
	b.NullValue, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagNullBuilder failed"))
	}
	return b
}

func (b *_BACnetProcessIdSelectionNullBuilder) Build() (BACnetProcessIdSelectionNull, error) {
	if b.NullValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'nullValue' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetProcessIdSelectionNull.deepCopy(), nil
}

func (b *_BACnetProcessIdSelectionNullBuilder) MustBuild() BACnetProcessIdSelectionNull {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetProcessIdSelectionNullBuilder) Done() BACnetProcessIdSelectionBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetProcessIdSelectionBuilder().(*_BACnetProcessIdSelectionBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetProcessIdSelectionNullBuilder) buildForBACnetProcessIdSelection() (BACnetProcessIdSelection, error) {
	return b.Build()
}

func (b *_BACnetProcessIdSelectionNullBuilder) DeepCopy() any {
	_copy := b.CreateBACnetProcessIdSelectionNullBuilder().(*_BACnetProcessIdSelectionNullBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetProcessIdSelectionNullBuilder creates a BACnetProcessIdSelectionNullBuilder
func (b *_BACnetProcessIdSelectionNull) CreateBACnetProcessIdSelectionNullBuilder() BACnetProcessIdSelectionNullBuilder {
	if b == nil {
		return NewBACnetProcessIdSelectionNullBuilder()
	}
	return &_BACnetProcessIdSelectionNullBuilder{_BACnetProcessIdSelectionNull: b.deepCopy()}
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

func (m *_BACnetProcessIdSelectionNull) GetParent() BACnetProcessIdSelectionContract {
	return m.BACnetProcessIdSelectionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetProcessIdSelectionNull) GetNullValue() BACnetApplicationTagNull {
	return m.NullValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetProcessIdSelectionNull(structType any) BACnetProcessIdSelectionNull {
	if casted, ok := structType.(BACnetProcessIdSelectionNull); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetProcessIdSelectionNull); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetProcessIdSelectionNull) GetTypeName() string {
	return "BACnetProcessIdSelectionNull"
}

func (m *_BACnetProcessIdSelectionNull) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetProcessIdSelectionContract.(*_BACnetProcessIdSelection).getLengthInBits(ctx))

	// Simple field (nullValue)
	lengthInBits += m.NullValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetProcessIdSelectionNull) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetProcessIdSelectionNull) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetProcessIdSelection) (__bACnetProcessIdSelectionNull BACnetProcessIdSelectionNull, err error) {
	m.BACnetProcessIdSelectionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetProcessIdSelectionNull"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetProcessIdSelectionNull")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	nullValue, err := ReadSimpleField[BACnetApplicationTagNull](ctx, "nullValue", ReadComplex[BACnetApplicationTagNull](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagNull](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'nullValue' field"))
	}
	m.NullValue = nullValue

	if closeErr := readBuffer.CloseContext("BACnetProcessIdSelectionNull"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetProcessIdSelectionNull")
	}

	return m, nil
}

func (m *_BACnetProcessIdSelectionNull) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetProcessIdSelectionNull) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetProcessIdSelectionNull"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetProcessIdSelectionNull")
		}

		if err := WriteSimpleField[BACnetApplicationTagNull](ctx, "nullValue", m.GetNullValue(), WriteComplex[BACnetApplicationTagNull](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'nullValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetProcessIdSelectionNull"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetProcessIdSelectionNull")
		}
		return nil
	}
	return m.BACnetProcessIdSelectionContract.(*_BACnetProcessIdSelection).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetProcessIdSelectionNull) IsBACnetProcessIdSelectionNull() {}

func (m *_BACnetProcessIdSelectionNull) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetProcessIdSelectionNull) deepCopy() *_BACnetProcessIdSelectionNull {
	if m == nil {
		return nil
	}
	_BACnetProcessIdSelectionNullCopy := &_BACnetProcessIdSelectionNull{
		m.BACnetProcessIdSelectionContract.(*_BACnetProcessIdSelection).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagNull](m.NullValue),
	}
	m.BACnetProcessIdSelectionContract.(*_BACnetProcessIdSelection)._SubType = m
	return _BACnetProcessIdSelectionNullCopy
}

func (m *_BACnetProcessIdSelectionNull) String() string {
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
