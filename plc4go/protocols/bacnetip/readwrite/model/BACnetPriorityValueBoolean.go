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

// BACnetPriorityValueBoolean is the corresponding interface of BACnetPriorityValueBoolean
type BACnetPriorityValueBoolean interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetPriorityValue
	// GetBooleanValue returns BooleanValue (property field)
	GetBooleanValue() BACnetApplicationTagBoolean
	// IsBACnetPriorityValueBoolean is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetPriorityValueBoolean()
	// CreateBuilder creates a BACnetPriorityValueBooleanBuilder
	CreateBACnetPriorityValueBooleanBuilder() BACnetPriorityValueBooleanBuilder
}

// _BACnetPriorityValueBoolean is the data-structure of this message
type _BACnetPriorityValueBoolean struct {
	BACnetPriorityValueContract
	BooleanValue BACnetApplicationTagBoolean
}

var _ BACnetPriorityValueBoolean = (*_BACnetPriorityValueBoolean)(nil)
var _ BACnetPriorityValueRequirements = (*_BACnetPriorityValueBoolean)(nil)

// NewBACnetPriorityValueBoolean factory function for _BACnetPriorityValueBoolean
func NewBACnetPriorityValueBoolean(peekedTagHeader BACnetTagHeader, booleanValue BACnetApplicationTagBoolean, objectTypeArgument BACnetObjectType) *_BACnetPriorityValueBoolean {
	if booleanValue == nil {
		panic("booleanValue of type BACnetApplicationTagBoolean for BACnetPriorityValueBoolean must not be nil")
	}
	_result := &_BACnetPriorityValueBoolean{
		BACnetPriorityValueContract: NewBACnetPriorityValue(peekedTagHeader, objectTypeArgument),
		BooleanValue:                booleanValue,
	}
	_result.BACnetPriorityValueContract.(*_BACnetPriorityValue)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetPriorityValueBooleanBuilder is a builder for BACnetPriorityValueBoolean
type BACnetPriorityValueBooleanBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(booleanValue BACnetApplicationTagBoolean) BACnetPriorityValueBooleanBuilder
	// WithBooleanValue adds BooleanValue (property field)
	WithBooleanValue(BACnetApplicationTagBoolean) BACnetPriorityValueBooleanBuilder
	// WithBooleanValueBuilder adds BooleanValue (property field) which is build by the builder
	WithBooleanValueBuilder(func(BACnetApplicationTagBooleanBuilder) BACnetApplicationTagBooleanBuilder) BACnetPriorityValueBooleanBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetPriorityValueBuilder
	// Build builds the BACnetPriorityValueBoolean or returns an error if something is wrong
	Build() (BACnetPriorityValueBoolean, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetPriorityValueBoolean
}

// NewBACnetPriorityValueBooleanBuilder() creates a BACnetPriorityValueBooleanBuilder
func NewBACnetPriorityValueBooleanBuilder() BACnetPriorityValueBooleanBuilder {
	return &_BACnetPriorityValueBooleanBuilder{_BACnetPriorityValueBoolean: new(_BACnetPriorityValueBoolean)}
}

type _BACnetPriorityValueBooleanBuilder struct {
	*_BACnetPriorityValueBoolean

	parentBuilder *_BACnetPriorityValueBuilder

	err *utils.MultiError
}

var _ (BACnetPriorityValueBooleanBuilder) = (*_BACnetPriorityValueBooleanBuilder)(nil)

func (b *_BACnetPriorityValueBooleanBuilder) setParent(contract BACnetPriorityValueContract) {
	b.BACnetPriorityValueContract = contract
	contract.(*_BACnetPriorityValue)._SubType = b._BACnetPriorityValueBoolean
}

func (b *_BACnetPriorityValueBooleanBuilder) WithMandatoryFields(booleanValue BACnetApplicationTagBoolean) BACnetPriorityValueBooleanBuilder {
	return b.WithBooleanValue(booleanValue)
}

func (b *_BACnetPriorityValueBooleanBuilder) WithBooleanValue(booleanValue BACnetApplicationTagBoolean) BACnetPriorityValueBooleanBuilder {
	b.BooleanValue = booleanValue
	return b
}

func (b *_BACnetPriorityValueBooleanBuilder) WithBooleanValueBuilder(builderSupplier func(BACnetApplicationTagBooleanBuilder) BACnetApplicationTagBooleanBuilder) BACnetPriorityValueBooleanBuilder {
	builder := builderSupplier(b.BooleanValue.CreateBACnetApplicationTagBooleanBuilder())
	var err error
	b.BooleanValue, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagBooleanBuilder failed"))
	}
	return b
}

func (b *_BACnetPriorityValueBooleanBuilder) Build() (BACnetPriorityValueBoolean, error) {
	if b.BooleanValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'booleanValue' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetPriorityValueBoolean.deepCopy(), nil
}

func (b *_BACnetPriorityValueBooleanBuilder) MustBuild() BACnetPriorityValueBoolean {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetPriorityValueBooleanBuilder) Done() BACnetPriorityValueBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetPriorityValueBuilder().(*_BACnetPriorityValueBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetPriorityValueBooleanBuilder) buildForBACnetPriorityValue() (BACnetPriorityValue, error) {
	return b.Build()
}

func (b *_BACnetPriorityValueBooleanBuilder) DeepCopy() any {
	_copy := b.CreateBACnetPriorityValueBooleanBuilder().(*_BACnetPriorityValueBooleanBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetPriorityValueBooleanBuilder creates a BACnetPriorityValueBooleanBuilder
func (b *_BACnetPriorityValueBoolean) CreateBACnetPriorityValueBooleanBuilder() BACnetPriorityValueBooleanBuilder {
	if b == nil {
		return NewBACnetPriorityValueBooleanBuilder()
	}
	return &_BACnetPriorityValueBooleanBuilder{_BACnetPriorityValueBoolean: b.deepCopy()}
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

func (m *_BACnetPriorityValueBoolean) GetParent() BACnetPriorityValueContract {
	return m.BACnetPriorityValueContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPriorityValueBoolean) GetBooleanValue() BACnetApplicationTagBoolean {
	return m.BooleanValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetPriorityValueBoolean(structType any) BACnetPriorityValueBoolean {
	if casted, ok := structType.(BACnetPriorityValueBoolean); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPriorityValueBoolean); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPriorityValueBoolean) GetTypeName() string {
	return "BACnetPriorityValueBoolean"
}

func (m *_BACnetPriorityValueBoolean) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetPriorityValueContract.(*_BACnetPriorityValue).getLengthInBits(ctx))

	// Simple field (booleanValue)
	lengthInBits += m.BooleanValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPriorityValueBoolean) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetPriorityValueBoolean) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetPriorityValue, objectTypeArgument BACnetObjectType) (__bACnetPriorityValueBoolean BACnetPriorityValueBoolean, err error) {
	m.BACnetPriorityValueContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPriorityValueBoolean"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPriorityValueBoolean")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	booleanValue, err := ReadSimpleField[BACnetApplicationTagBoolean](ctx, "booleanValue", ReadComplex[BACnetApplicationTagBoolean](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagBoolean](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'booleanValue' field"))
	}
	m.BooleanValue = booleanValue

	if closeErr := readBuffer.CloseContext("BACnetPriorityValueBoolean"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPriorityValueBoolean")
	}

	return m, nil
}

func (m *_BACnetPriorityValueBoolean) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPriorityValueBoolean) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPriorityValueBoolean"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPriorityValueBoolean")
		}

		if err := WriteSimpleField[BACnetApplicationTagBoolean](ctx, "booleanValue", m.GetBooleanValue(), WriteComplex[BACnetApplicationTagBoolean](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'booleanValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPriorityValueBoolean"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPriorityValueBoolean")
		}
		return nil
	}
	return m.BACnetPriorityValueContract.(*_BACnetPriorityValue).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPriorityValueBoolean) IsBACnetPriorityValueBoolean() {}

func (m *_BACnetPriorityValueBoolean) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetPriorityValueBoolean) deepCopy() *_BACnetPriorityValueBoolean {
	if m == nil {
		return nil
	}
	_BACnetPriorityValueBooleanCopy := &_BACnetPriorityValueBoolean{
		m.BACnetPriorityValueContract.(*_BACnetPriorityValue).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagBoolean](m.BooleanValue),
	}
	m.BACnetPriorityValueContract.(*_BACnetPriorityValue)._SubType = m
	return _BACnetPriorityValueBooleanCopy
}

func (m *_BACnetPriorityValueBoolean) String() string {
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
