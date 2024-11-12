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

// BACnetOptionalCharacterStringValue is the corresponding interface of BACnetOptionalCharacterStringValue
type BACnetOptionalCharacterStringValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetOptionalCharacterString
	// GetCharacterstring returns Characterstring (property field)
	GetCharacterstring() BACnetApplicationTagCharacterString
	// IsBACnetOptionalCharacterStringValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetOptionalCharacterStringValue()
	// CreateBuilder creates a BACnetOptionalCharacterStringValueBuilder
	CreateBACnetOptionalCharacterStringValueBuilder() BACnetOptionalCharacterStringValueBuilder
}

// _BACnetOptionalCharacterStringValue is the data-structure of this message
type _BACnetOptionalCharacterStringValue struct {
	BACnetOptionalCharacterStringContract
	Characterstring BACnetApplicationTagCharacterString
}

var _ BACnetOptionalCharacterStringValue = (*_BACnetOptionalCharacterStringValue)(nil)
var _ BACnetOptionalCharacterStringRequirements = (*_BACnetOptionalCharacterStringValue)(nil)

// NewBACnetOptionalCharacterStringValue factory function for _BACnetOptionalCharacterStringValue
func NewBACnetOptionalCharacterStringValue(peekedTagHeader BACnetTagHeader, characterstring BACnetApplicationTagCharacterString) *_BACnetOptionalCharacterStringValue {
	if characterstring == nil {
		panic("characterstring of type BACnetApplicationTagCharacterString for BACnetOptionalCharacterStringValue must not be nil")
	}
	_result := &_BACnetOptionalCharacterStringValue{
		BACnetOptionalCharacterStringContract: NewBACnetOptionalCharacterString(peekedTagHeader),
		Characterstring:                       characterstring,
	}
	_result.BACnetOptionalCharacterStringContract.(*_BACnetOptionalCharacterString)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetOptionalCharacterStringValueBuilder is a builder for BACnetOptionalCharacterStringValue
type BACnetOptionalCharacterStringValueBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(characterstring BACnetApplicationTagCharacterString) BACnetOptionalCharacterStringValueBuilder
	// WithCharacterstring adds Characterstring (property field)
	WithCharacterstring(BACnetApplicationTagCharacterString) BACnetOptionalCharacterStringValueBuilder
	// WithCharacterstringBuilder adds Characterstring (property field) which is build by the builder
	WithCharacterstringBuilder(func(BACnetApplicationTagCharacterStringBuilder) BACnetApplicationTagCharacterStringBuilder) BACnetOptionalCharacterStringValueBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetOptionalCharacterStringBuilder
	// Build builds the BACnetOptionalCharacterStringValue or returns an error if something is wrong
	Build() (BACnetOptionalCharacterStringValue, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetOptionalCharacterStringValue
}

// NewBACnetOptionalCharacterStringValueBuilder() creates a BACnetOptionalCharacterStringValueBuilder
func NewBACnetOptionalCharacterStringValueBuilder() BACnetOptionalCharacterStringValueBuilder {
	return &_BACnetOptionalCharacterStringValueBuilder{_BACnetOptionalCharacterStringValue: new(_BACnetOptionalCharacterStringValue)}
}

type _BACnetOptionalCharacterStringValueBuilder struct {
	*_BACnetOptionalCharacterStringValue

	parentBuilder *_BACnetOptionalCharacterStringBuilder

	err *utils.MultiError
}

var _ (BACnetOptionalCharacterStringValueBuilder) = (*_BACnetOptionalCharacterStringValueBuilder)(nil)

func (b *_BACnetOptionalCharacterStringValueBuilder) setParent(contract BACnetOptionalCharacterStringContract) {
	b.BACnetOptionalCharacterStringContract = contract
	contract.(*_BACnetOptionalCharacterString)._SubType = b._BACnetOptionalCharacterStringValue
}

func (b *_BACnetOptionalCharacterStringValueBuilder) WithMandatoryFields(characterstring BACnetApplicationTagCharacterString) BACnetOptionalCharacterStringValueBuilder {
	return b.WithCharacterstring(characterstring)
}

func (b *_BACnetOptionalCharacterStringValueBuilder) WithCharacterstring(characterstring BACnetApplicationTagCharacterString) BACnetOptionalCharacterStringValueBuilder {
	b.Characterstring = characterstring
	return b
}

func (b *_BACnetOptionalCharacterStringValueBuilder) WithCharacterstringBuilder(builderSupplier func(BACnetApplicationTagCharacterStringBuilder) BACnetApplicationTagCharacterStringBuilder) BACnetOptionalCharacterStringValueBuilder {
	builder := builderSupplier(b.Characterstring.CreateBACnetApplicationTagCharacterStringBuilder())
	var err error
	b.Characterstring, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagCharacterStringBuilder failed"))
	}
	return b
}

func (b *_BACnetOptionalCharacterStringValueBuilder) Build() (BACnetOptionalCharacterStringValue, error) {
	if b.Characterstring == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'characterstring' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetOptionalCharacterStringValue.deepCopy(), nil
}

func (b *_BACnetOptionalCharacterStringValueBuilder) MustBuild() BACnetOptionalCharacterStringValue {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetOptionalCharacterStringValueBuilder) Done() BACnetOptionalCharacterStringBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetOptionalCharacterStringBuilder().(*_BACnetOptionalCharacterStringBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetOptionalCharacterStringValueBuilder) buildForBACnetOptionalCharacterString() (BACnetOptionalCharacterString, error) {
	return b.Build()
}

func (b *_BACnetOptionalCharacterStringValueBuilder) DeepCopy() any {
	_copy := b.CreateBACnetOptionalCharacterStringValueBuilder().(*_BACnetOptionalCharacterStringValueBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetOptionalCharacterStringValueBuilder creates a BACnetOptionalCharacterStringValueBuilder
func (b *_BACnetOptionalCharacterStringValue) CreateBACnetOptionalCharacterStringValueBuilder() BACnetOptionalCharacterStringValueBuilder {
	if b == nil {
		return NewBACnetOptionalCharacterStringValueBuilder()
	}
	return &_BACnetOptionalCharacterStringValueBuilder{_BACnetOptionalCharacterStringValue: b.deepCopy()}
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

func (m *_BACnetOptionalCharacterStringValue) GetParent() BACnetOptionalCharacterStringContract {
	return m.BACnetOptionalCharacterStringContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetOptionalCharacterStringValue) GetCharacterstring() BACnetApplicationTagCharacterString {
	return m.Characterstring
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetOptionalCharacterStringValue(structType any) BACnetOptionalCharacterStringValue {
	if casted, ok := structType.(BACnetOptionalCharacterStringValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetOptionalCharacterStringValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetOptionalCharacterStringValue) GetTypeName() string {
	return "BACnetOptionalCharacterStringValue"
}

func (m *_BACnetOptionalCharacterStringValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetOptionalCharacterStringContract.(*_BACnetOptionalCharacterString).getLengthInBits(ctx))

	// Simple field (characterstring)
	lengthInBits += m.Characterstring.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetOptionalCharacterStringValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetOptionalCharacterStringValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetOptionalCharacterString) (__bACnetOptionalCharacterStringValue BACnetOptionalCharacterStringValue, err error) {
	m.BACnetOptionalCharacterStringContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetOptionalCharacterStringValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetOptionalCharacterStringValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	characterstring, err := ReadSimpleField[BACnetApplicationTagCharacterString](ctx, "characterstring", ReadComplex[BACnetApplicationTagCharacterString](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagCharacterString](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'characterstring' field"))
	}
	m.Characterstring = characterstring

	if closeErr := readBuffer.CloseContext("BACnetOptionalCharacterStringValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetOptionalCharacterStringValue")
	}

	return m, nil
}

func (m *_BACnetOptionalCharacterStringValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetOptionalCharacterStringValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetOptionalCharacterStringValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetOptionalCharacterStringValue")
		}

		if err := WriteSimpleField[BACnetApplicationTagCharacterString](ctx, "characterstring", m.GetCharacterstring(), WriteComplex[BACnetApplicationTagCharacterString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'characterstring' field")
		}

		if popErr := writeBuffer.PopContext("BACnetOptionalCharacterStringValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetOptionalCharacterStringValue")
		}
		return nil
	}
	return m.BACnetOptionalCharacterStringContract.(*_BACnetOptionalCharacterString).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetOptionalCharacterStringValue) IsBACnetOptionalCharacterStringValue() {}

func (m *_BACnetOptionalCharacterStringValue) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetOptionalCharacterStringValue) deepCopy() *_BACnetOptionalCharacterStringValue {
	if m == nil {
		return nil
	}
	_BACnetOptionalCharacterStringValueCopy := &_BACnetOptionalCharacterStringValue{
		m.BACnetOptionalCharacterStringContract.(*_BACnetOptionalCharacterString).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagCharacterString](m.Characterstring),
	}
	m.BACnetOptionalCharacterStringContract.(*_BACnetOptionalCharacterString)._SubType = m
	return _BACnetOptionalCharacterStringValueCopy
}

func (m *_BACnetOptionalCharacterStringValue) String() string {
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
