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

// BACnetNotificationParametersChangeOfValueNewValue is the corresponding interface of BACnetNotificationParametersChangeOfValueNewValue
type BACnetNotificationParametersChangeOfValueNewValue interface {
	BACnetNotificationParametersChangeOfValueNewValueContract
	BACnetNotificationParametersChangeOfValueNewValueRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// IsBACnetNotificationParametersChangeOfValueNewValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetNotificationParametersChangeOfValueNewValue()
	// CreateBuilder creates a BACnetNotificationParametersChangeOfValueNewValueBuilder
	CreateBACnetNotificationParametersChangeOfValueNewValueBuilder() BACnetNotificationParametersChangeOfValueNewValueBuilder
}

// BACnetNotificationParametersChangeOfValueNewValueContract provides a set of functions which can be overwritten by a sub struct
type BACnetNotificationParametersChangeOfValueNewValueContract interface {
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetPeekedTagHeader returns PeekedTagHeader (property field)
	GetPeekedTagHeader() BACnetTagHeader
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
	// GetPeekedTagNumber returns PeekedTagNumber (virtual field)
	GetPeekedTagNumber() uint8
	// GetTagNumber() returns a parser argument
	GetTagNumber() uint8
	// IsBACnetNotificationParametersChangeOfValueNewValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetNotificationParametersChangeOfValueNewValue()
	// CreateBuilder creates a BACnetNotificationParametersChangeOfValueNewValueBuilder
	CreateBACnetNotificationParametersChangeOfValueNewValueBuilder() BACnetNotificationParametersChangeOfValueNewValueBuilder
}

// BACnetNotificationParametersChangeOfValueNewValueRequirements provides a set of functions which need to be implemented by a sub struct
type BACnetNotificationParametersChangeOfValueNewValueRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetPeekedTagNumber returns PeekedTagNumber (discriminator field)
	GetPeekedTagNumber() uint8
}

// _BACnetNotificationParametersChangeOfValueNewValue is the data-structure of this message
type _BACnetNotificationParametersChangeOfValueNewValue struct {
	_SubType interface {
		BACnetNotificationParametersChangeOfValueNewValueContract
		BACnetNotificationParametersChangeOfValueNewValueRequirements
	}
	OpeningTag      BACnetOpeningTag
	PeekedTagHeader BACnetTagHeader
	ClosingTag      BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

var _ BACnetNotificationParametersChangeOfValueNewValueContract = (*_BACnetNotificationParametersChangeOfValueNewValue)(nil)

// NewBACnetNotificationParametersChangeOfValueNewValue factory function for _BACnetNotificationParametersChangeOfValueNewValue
func NewBACnetNotificationParametersChangeOfValueNewValue(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetNotificationParametersChangeOfValueNewValue {
	if openingTag == nil {
		panic("openingTag of type BACnetOpeningTag for BACnetNotificationParametersChangeOfValueNewValue must not be nil")
	}
	if peekedTagHeader == nil {
		panic("peekedTagHeader of type BACnetTagHeader for BACnetNotificationParametersChangeOfValueNewValue must not be nil")
	}
	if closingTag == nil {
		panic("closingTag of type BACnetClosingTag for BACnetNotificationParametersChangeOfValueNewValue must not be nil")
	}
	return &_BACnetNotificationParametersChangeOfValueNewValue{OpeningTag: openingTag, PeekedTagHeader: peekedTagHeader, ClosingTag: closingTag, TagNumber: tagNumber}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetNotificationParametersChangeOfValueNewValueBuilder is a builder for BACnetNotificationParametersChangeOfValueNewValue
type BACnetNotificationParametersChangeOfValueNewValueBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) BACnetNotificationParametersChangeOfValueNewValueBuilder
	// WithOpeningTag adds OpeningTag (property field)
	WithOpeningTag(BACnetOpeningTag) BACnetNotificationParametersChangeOfValueNewValueBuilder
	// WithOpeningTagBuilder adds OpeningTag (property field) which is build by the builder
	WithOpeningTagBuilder(func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetNotificationParametersChangeOfValueNewValueBuilder
	// WithPeekedTagHeader adds PeekedTagHeader (property field)
	WithPeekedTagHeader(BACnetTagHeader) BACnetNotificationParametersChangeOfValueNewValueBuilder
	// WithPeekedTagHeaderBuilder adds PeekedTagHeader (property field) which is build by the builder
	WithPeekedTagHeaderBuilder(func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetNotificationParametersChangeOfValueNewValueBuilder
	// WithClosingTag adds ClosingTag (property field)
	WithClosingTag(BACnetClosingTag) BACnetNotificationParametersChangeOfValueNewValueBuilder
	// WithClosingTagBuilder adds ClosingTag (property field) which is build by the builder
	WithClosingTagBuilder(func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetNotificationParametersChangeOfValueNewValueBuilder
	// AsBACnetNotificationParametersChangeOfValueNewValueChangedBits converts this build to a subType of BACnetNotificationParametersChangeOfValueNewValue. It is always possible to return to current builder using Done()
	AsBACnetNotificationParametersChangeOfValueNewValueChangedBits() interface {
		BACnetNotificationParametersChangeOfValueNewValueChangedBitsBuilder
		Done() BACnetNotificationParametersChangeOfValueNewValueBuilder
	}
	// AsBACnetNotificationParametersChangeOfValueNewValueChangedValue converts this build to a subType of BACnetNotificationParametersChangeOfValueNewValue. It is always possible to return to current builder using Done()
	AsBACnetNotificationParametersChangeOfValueNewValueChangedValue() interface {
		BACnetNotificationParametersChangeOfValueNewValueChangedValueBuilder
		Done() BACnetNotificationParametersChangeOfValueNewValueBuilder
	}
	// Build builds the BACnetNotificationParametersChangeOfValueNewValue or returns an error if something is wrong
	PartialBuild() (BACnetNotificationParametersChangeOfValueNewValueContract, error)
	// MustBuild does the same as Build but panics on error
	PartialMustBuild() BACnetNotificationParametersChangeOfValueNewValueContract
	// Build builds the BACnetNotificationParametersChangeOfValueNewValue or returns an error if something is wrong
	Build() (BACnetNotificationParametersChangeOfValueNewValue, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetNotificationParametersChangeOfValueNewValue
}

// NewBACnetNotificationParametersChangeOfValueNewValueBuilder() creates a BACnetNotificationParametersChangeOfValueNewValueBuilder
func NewBACnetNotificationParametersChangeOfValueNewValueBuilder() BACnetNotificationParametersChangeOfValueNewValueBuilder {
	return &_BACnetNotificationParametersChangeOfValueNewValueBuilder{_BACnetNotificationParametersChangeOfValueNewValue: new(_BACnetNotificationParametersChangeOfValueNewValue)}
}

type _BACnetNotificationParametersChangeOfValueNewValueChildBuilder interface {
	utils.Copyable
	setParent(BACnetNotificationParametersChangeOfValueNewValueContract)
	buildForBACnetNotificationParametersChangeOfValueNewValue() (BACnetNotificationParametersChangeOfValueNewValue, error)
}

type _BACnetNotificationParametersChangeOfValueNewValueBuilder struct {
	*_BACnetNotificationParametersChangeOfValueNewValue

	childBuilder _BACnetNotificationParametersChangeOfValueNewValueChildBuilder

	err *utils.MultiError
}

var _ (BACnetNotificationParametersChangeOfValueNewValueBuilder) = (*_BACnetNotificationParametersChangeOfValueNewValueBuilder)(nil)

func (b *_BACnetNotificationParametersChangeOfValueNewValueBuilder) WithMandatoryFields(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) BACnetNotificationParametersChangeOfValueNewValueBuilder {
	return b.WithOpeningTag(openingTag).WithPeekedTagHeader(peekedTagHeader).WithClosingTag(closingTag)
}

func (b *_BACnetNotificationParametersChangeOfValueNewValueBuilder) WithOpeningTag(openingTag BACnetOpeningTag) BACnetNotificationParametersChangeOfValueNewValueBuilder {
	b.OpeningTag = openingTag
	return b
}

func (b *_BACnetNotificationParametersChangeOfValueNewValueBuilder) WithOpeningTagBuilder(builderSupplier func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetNotificationParametersChangeOfValueNewValueBuilder {
	builder := builderSupplier(b.OpeningTag.CreateBACnetOpeningTagBuilder())
	var err error
	b.OpeningTag, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetOpeningTagBuilder failed"))
	}
	return b
}

func (b *_BACnetNotificationParametersChangeOfValueNewValueBuilder) WithPeekedTagHeader(peekedTagHeader BACnetTagHeader) BACnetNotificationParametersChangeOfValueNewValueBuilder {
	b.PeekedTagHeader = peekedTagHeader
	return b
}

func (b *_BACnetNotificationParametersChangeOfValueNewValueBuilder) WithPeekedTagHeaderBuilder(builderSupplier func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetNotificationParametersChangeOfValueNewValueBuilder {
	builder := builderSupplier(b.PeekedTagHeader.CreateBACnetTagHeaderBuilder())
	var err error
	b.PeekedTagHeader, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetTagHeaderBuilder failed"))
	}
	return b
}

func (b *_BACnetNotificationParametersChangeOfValueNewValueBuilder) WithClosingTag(closingTag BACnetClosingTag) BACnetNotificationParametersChangeOfValueNewValueBuilder {
	b.ClosingTag = closingTag
	return b
}

func (b *_BACnetNotificationParametersChangeOfValueNewValueBuilder) WithClosingTagBuilder(builderSupplier func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetNotificationParametersChangeOfValueNewValueBuilder {
	builder := builderSupplier(b.ClosingTag.CreateBACnetClosingTagBuilder())
	var err error
	b.ClosingTag, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetClosingTagBuilder failed"))
	}
	return b
}

func (b *_BACnetNotificationParametersChangeOfValueNewValueBuilder) PartialBuild() (BACnetNotificationParametersChangeOfValueNewValueContract, error) {
	if b.OpeningTag == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'openingTag' not set"))
	}
	if b.PeekedTagHeader == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'peekedTagHeader' not set"))
	}
	if b.ClosingTag == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'closingTag' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetNotificationParametersChangeOfValueNewValue.deepCopy(), nil
}

func (b *_BACnetNotificationParametersChangeOfValueNewValueBuilder) PartialMustBuild() BACnetNotificationParametersChangeOfValueNewValueContract {
	build, err := b.PartialBuild()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetNotificationParametersChangeOfValueNewValueBuilder) AsBACnetNotificationParametersChangeOfValueNewValueChangedBits() interface {
	BACnetNotificationParametersChangeOfValueNewValueChangedBitsBuilder
	Done() BACnetNotificationParametersChangeOfValueNewValueBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		BACnetNotificationParametersChangeOfValueNewValueChangedBitsBuilder
		Done() BACnetNotificationParametersChangeOfValueNewValueBuilder
	}); ok {
		return cb
	}
	cb := NewBACnetNotificationParametersChangeOfValueNewValueChangedBitsBuilder().(*_BACnetNotificationParametersChangeOfValueNewValueChangedBitsBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_BACnetNotificationParametersChangeOfValueNewValueBuilder) AsBACnetNotificationParametersChangeOfValueNewValueChangedValue() interface {
	BACnetNotificationParametersChangeOfValueNewValueChangedValueBuilder
	Done() BACnetNotificationParametersChangeOfValueNewValueBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		BACnetNotificationParametersChangeOfValueNewValueChangedValueBuilder
		Done() BACnetNotificationParametersChangeOfValueNewValueBuilder
	}); ok {
		return cb
	}
	cb := NewBACnetNotificationParametersChangeOfValueNewValueChangedValueBuilder().(*_BACnetNotificationParametersChangeOfValueNewValueChangedValueBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_BACnetNotificationParametersChangeOfValueNewValueBuilder) Build() (BACnetNotificationParametersChangeOfValueNewValue, error) {
	v, err := b.PartialBuild()
	if err != nil {
		return nil, errors.Wrap(err, "error occurred during partial build")
	}
	if b.childBuilder == nil {
		return nil, errors.New("no child builder present")
	}
	b.childBuilder.setParent(v)
	return b.childBuilder.buildForBACnetNotificationParametersChangeOfValueNewValue()
}

func (b *_BACnetNotificationParametersChangeOfValueNewValueBuilder) MustBuild() BACnetNotificationParametersChangeOfValueNewValue {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetNotificationParametersChangeOfValueNewValueBuilder) DeepCopy() any {
	_copy := b.CreateBACnetNotificationParametersChangeOfValueNewValueBuilder().(*_BACnetNotificationParametersChangeOfValueNewValueBuilder)
	_copy.childBuilder = b.childBuilder.DeepCopy().(_BACnetNotificationParametersChangeOfValueNewValueChildBuilder)
	_copy.childBuilder.setParent(_copy)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetNotificationParametersChangeOfValueNewValueBuilder creates a BACnetNotificationParametersChangeOfValueNewValueBuilder
func (b *_BACnetNotificationParametersChangeOfValueNewValue) CreateBACnetNotificationParametersChangeOfValueNewValueBuilder() BACnetNotificationParametersChangeOfValueNewValueBuilder {
	if b == nil {
		return NewBACnetNotificationParametersChangeOfValueNewValueBuilder()
	}
	return &_BACnetNotificationParametersChangeOfValueNewValueBuilder{_BACnetNotificationParametersChangeOfValueNewValue: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetNotificationParametersChangeOfValueNewValue) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetNotificationParametersChangeOfValueNewValue) GetPeekedTagHeader() BACnetTagHeader {
	return m.PeekedTagHeader
}

func (m *_BACnetNotificationParametersChangeOfValueNewValue) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (pm *_BACnetNotificationParametersChangeOfValueNewValue) GetPeekedTagNumber() uint8 {
	m := pm._SubType
	ctx := context.Background()
	_ = ctx
	return uint8(m.GetPeekedTagHeader().GetActualTagNumber())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetNotificationParametersChangeOfValueNewValue(structType any) BACnetNotificationParametersChangeOfValueNewValue {
	if casted, ok := structType.(BACnetNotificationParametersChangeOfValueNewValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetNotificationParametersChangeOfValueNewValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetNotificationParametersChangeOfValueNewValue) GetTypeName() string {
	return "BACnetNotificationParametersChangeOfValueNewValue"
}

func (m *_BACnetNotificationParametersChangeOfValueNewValue) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetNotificationParametersChangeOfValueNewValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func BACnetNotificationParametersChangeOfValueNewValueParse[T BACnetNotificationParametersChangeOfValueNewValue](ctx context.Context, theBytes []byte, tagNumber uint8) (T, error) {
	return BACnetNotificationParametersChangeOfValueNewValueParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetNotificationParametersChangeOfValueNewValueParseWithBufferProducer[T BACnetNotificationParametersChangeOfValueNewValue](tagNumber uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := BACnetNotificationParametersChangeOfValueNewValueParseWithBuffer[T](ctx, readBuffer, tagNumber)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, nil
	}
}

func BACnetNotificationParametersChangeOfValueNewValueParseWithBuffer[T BACnetNotificationParametersChangeOfValueNewValue](ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (T, error) {
	v, err := (&_BACnetNotificationParametersChangeOfValueNewValue{TagNumber: tagNumber}).parse(ctx, readBuffer, tagNumber)
	if err != nil {
		var zero T
		return zero, err
	}
	vc, ok := v.(T)
	if !ok {
		var zero T
		return zero, errors.Errorf("Unexpected type %T. Expected type %T", v, *new(T))
	}
	return vc, nil
}

func (m *_BACnetNotificationParametersChangeOfValueNewValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (__bACnetNotificationParametersChangeOfValueNewValue BACnetNotificationParametersChangeOfValueNewValue, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetNotificationParametersChangeOfValueNewValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetNotificationParametersChangeOfValueNewValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	m.OpeningTag = openingTag

	peekedTagHeader, err := ReadPeekField[BACnetTagHeader](ctx, "peekedTagHeader", ReadComplex[BACnetTagHeader](BACnetTagHeaderParseWithBuffer, readBuffer), 0)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'peekedTagHeader' field"))
	}
	m.PeekedTagHeader = peekedTagHeader

	peekedTagNumber, err := ReadVirtualField[uint8](ctx, "peekedTagNumber", (*uint8)(nil), peekedTagHeader.GetActualTagNumber())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'peekedTagNumber' field"))
	}
	_ = peekedTagNumber

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child BACnetNotificationParametersChangeOfValueNewValue
	switch {
	case peekedTagNumber == uint8(0): // BACnetNotificationParametersChangeOfValueNewValueChangedBits
		if _child, err = new(_BACnetNotificationParametersChangeOfValueNewValueChangedBits).parse(ctx, readBuffer, m, peekedTagNumber, tagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetNotificationParametersChangeOfValueNewValueChangedBits for type-switch of BACnetNotificationParametersChangeOfValueNewValue")
		}
	case peekedTagNumber == uint8(1): // BACnetNotificationParametersChangeOfValueNewValueChangedValue
		if _child, err = new(_BACnetNotificationParametersChangeOfValueNewValueChangedValue).parse(ctx, readBuffer, m, peekedTagNumber, tagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetNotificationParametersChangeOfValueNewValueChangedValue for type-switch of BACnetNotificationParametersChangeOfValueNewValue")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [peekedTagNumber=%v]", peekedTagNumber)
	}

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("BACnetNotificationParametersChangeOfValueNewValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetNotificationParametersChangeOfValueNewValue")
	}

	return _child, nil
}

func (pm *_BACnetNotificationParametersChangeOfValueNewValue) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetNotificationParametersChangeOfValueNewValue, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetNotificationParametersChangeOfValueNewValue"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetNotificationParametersChangeOfValueNewValue")
	}

	if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'openingTag' field")
	}
	// Virtual field
	peekedTagNumber := m.GetPeekedTagNumber()
	_ = peekedTagNumber
	if _peekedTagNumberErr := writeBuffer.WriteVirtual(ctx, "peekedTagNumber", m.GetPeekedTagNumber()); _peekedTagNumberErr != nil {
		return errors.Wrap(_peekedTagNumberErr, "Error serializing 'peekedTagNumber' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetNotificationParametersChangeOfValueNewValue"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetNotificationParametersChangeOfValueNewValue")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetNotificationParametersChangeOfValueNewValue) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetNotificationParametersChangeOfValueNewValue) IsBACnetNotificationParametersChangeOfValueNewValue() {
}

func (m *_BACnetNotificationParametersChangeOfValueNewValue) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetNotificationParametersChangeOfValueNewValue) deepCopy() *_BACnetNotificationParametersChangeOfValueNewValue {
	if m == nil {
		return nil
	}
	_BACnetNotificationParametersChangeOfValueNewValueCopy := &_BACnetNotificationParametersChangeOfValueNewValue{
		nil, // will be set by child
		m.OpeningTag.DeepCopy().(BACnetOpeningTag),
		m.PeekedTagHeader.DeepCopy().(BACnetTagHeader),
		m.ClosingTag.DeepCopy().(BACnetClosingTag),
		m.TagNumber,
	}
	return _BACnetNotificationParametersChangeOfValueNewValueCopy
}
