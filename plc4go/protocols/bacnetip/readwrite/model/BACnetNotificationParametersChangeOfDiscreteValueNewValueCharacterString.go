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

// BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterString is the corresponding interface of BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterString
type BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterString interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetNotificationParametersChangeOfDiscreteValueNewValue
	// GetCharacterStringValue returns CharacterStringValue (property field)
	GetCharacterStringValue() BACnetApplicationTagCharacterString
	// IsBACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterString is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterString()
	// CreateBuilder creates a BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterStringBuilder
	CreateBACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterStringBuilder() BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterStringBuilder
}

// _BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterString is the data-structure of this message
type _BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterString struct {
	BACnetNotificationParametersChangeOfDiscreteValueNewValueContract
	CharacterStringValue BACnetApplicationTagCharacterString
}

var _ BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterString = (*_BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterString)(nil)
var _ BACnetNotificationParametersChangeOfDiscreteValueNewValueRequirements = (*_BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterString)(nil)

// NewBACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterString factory function for _BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterString
func NewBACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterString(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, characterStringValue BACnetApplicationTagCharacterString, tagNumber uint8) *_BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterString {
	if characterStringValue == nil {
		panic("characterStringValue of type BACnetApplicationTagCharacterString for BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterString must not be nil")
	}
	_result := &_BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterString{
		BACnetNotificationParametersChangeOfDiscreteValueNewValueContract: NewBACnetNotificationParametersChangeOfDiscreteValueNewValue(openingTag, peekedTagHeader, closingTag, tagNumber),
		CharacterStringValue: characterStringValue,
	}
	_result.BACnetNotificationParametersChangeOfDiscreteValueNewValueContract.(*_BACnetNotificationParametersChangeOfDiscreteValueNewValue)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterStringBuilder is a builder for BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterString
type BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterStringBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(characterStringValue BACnetApplicationTagCharacterString) BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterStringBuilder
	// WithCharacterStringValue adds CharacterStringValue (property field)
	WithCharacterStringValue(BACnetApplicationTagCharacterString) BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterStringBuilder
	// WithCharacterStringValueBuilder adds CharacterStringValue (property field) which is build by the builder
	WithCharacterStringValueBuilder(func(BACnetApplicationTagCharacterStringBuilder) BACnetApplicationTagCharacterStringBuilder) BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterStringBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetNotificationParametersChangeOfDiscreteValueNewValueBuilder
	// Build builds the BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterString or returns an error if something is wrong
	Build() (BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterString, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterString
}

// NewBACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterStringBuilder() creates a BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterStringBuilder
func NewBACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterStringBuilder() BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterStringBuilder {
	return &_BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterStringBuilder{_BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterString: new(_BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterString)}
}

type _BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterStringBuilder struct {
	*_BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterString

	parentBuilder *_BACnetNotificationParametersChangeOfDiscreteValueNewValueBuilder

	err *utils.MultiError
}

var _ (BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterStringBuilder) = (*_BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterStringBuilder)(nil)

func (b *_BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterStringBuilder) setParent(contract BACnetNotificationParametersChangeOfDiscreteValueNewValueContract) {
	b.BACnetNotificationParametersChangeOfDiscreteValueNewValueContract = contract
	contract.(*_BACnetNotificationParametersChangeOfDiscreteValueNewValue)._SubType = b._BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterString
}

func (b *_BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterStringBuilder) WithMandatoryFields(characterStringValue BACnetApplicationTagCharacterString) BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterStringBuilder {
	return b.WithCharacterStringValue(characterStringValue)
}

func (b *_BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterStringBuilder) WithCharacterStringValue(characterStringValue BACnetApplicationTagCharacterString) BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterStringBuilder {
	b.CharacterStringValue = characterStringValue
	return b
}

func (b *_BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterStringBuilder) WithCharacterStringValueBuilder(builderSupplier func(BACnetApplicationTagCharacterStringBuilder) BACnetApplicationTagCharacterStringBuilder) BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterStringBuilder {
	builder := builderSupplier(b.CharacterStringValue.CreateBACnetApplicationTagCharacterStringBuilder())
	var err error
	b.CharacterStringValue, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagCharacterStringBuilder failed"))
	}
	return b
}

func (b *_BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterStringBuilder) Build() (BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterString, error) {
	if b.CharacterStringValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'characterStringValue' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterString.deepCopy(), nil
}

func (b *_BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterStringBuilder) MustBuild() BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterString {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterStringBuilder) Done() BACnetNotificationParametersChangeOfDiscreteValueNewValueBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetNotificationParametersChangeOfDiscreteValueNewValueBuilder().(*_BACnetNotificationParametersChangeOfDiscreteValueNewValueBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterStringBuilder) buildForBACnetNotificationParametersChangeOfDiscreteValueNewValue() (BACnetNotificationParametersChangeOfDiscreteValueNewValue, error) {
	return b.Build()
}

func (b *_BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterStringBuilder) DeepCopy() any {
	_copy := b.CreateBACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterStringBuilder().(*_BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterStringBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterStringBuilder creates a BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterStringBuilder
func (b *_BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterString) CreateBACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterStringBuilder() BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterStringBuilder {
	if b == nil {
		return NewBACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterStringBuilder()
	}
	return &_BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterStringBuilder{_BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterString: b.deepCopy()}
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

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterString) GetParent() BACnetNotificationParametersChangeOfDiscreteValueNewValueContract {
	return m.BACnetNotificationParametersChangeOfDiscreteValueNewValueContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterString) GetCharacterStringValue() BACnetApplicationTagCharacterString {
	return m.CharacterStringValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterString(structType any) BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterString {
	if casted, ok := structType.(BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterString); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterString); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterString) GetTypeName() string {
	return "BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterString"
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterString) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetNotificationParametersChangeOfDiscreteValueNewValueContract.(*_BACnetNotificationParametersChangeOfDiscreteValueNewValue).getLengthInBits(ctx))

	// Simple field (characterStringValue)
	lengthInBits += m.CharacterStringValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterString) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterString) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetNotificationParametersChangeOfDiscreteValueNewValue, tagNumber uint8) (__bACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterString BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterString, err error) {
	m.BACnetNotificationParametersChangeOfDiscreteValueNewValueContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterString"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterString")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	characterStringValue, err := ReadSimpleField[BACnetApplicationTagCharacterString](ctx, "characterStringValue", ReadComplex[BACnetApplicationTagCharacterString](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagCharacterString](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'characterStringValue' field"))
	}
	m.CharacterStringValue = characterStringValue

	if closeErr := readBuffer.CloseContext("BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterString"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterString")
	}

	return m, nil
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterString) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterString) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterString"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterString")
		}

		if err := WriteSimpleField[BACnetApplicationTagCharacterString](ctx, "characterStringValue", m.GetCharacterStringValue(), WriteComplex[BACnetApplicationTagCharacterString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'characterStringValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterString"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterString")
		}
		return nil
	}
	return m.BACnetNotificationParametersChangeOfDiscreteValueNewValueContract.(*_BACnetNotificationParametersChangeOfDiscreteValueNewValue).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterString) IsBACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterString() {
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterString) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterString) deepCopy() *_BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterString {
	if m == nil {
		return nil
	}
	_BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterStringCopy := &_BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterString{
		m.BACnetNotificationParametersChangeOfDiscreteValueNewValueContract.(*_BACnetNotificationParametersChangeOfDiscreteValueNewValue).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagCharacterString](m.CharacterStringValue),
	}
	m.BACnetNotificationParametersChangeOfDiscreteValueNewValueContract.(*_BACnetNotificationParametersChangeOfDiscreteValueNewValue)._SubType = m
	return _BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterStringCopy
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterString) String() string {
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
