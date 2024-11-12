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

// BACnetEventParameterUnsignedOutOfRange is the corresponding interface of BACnetEventParameterUnsignedOutOfRange
type BACnetEventParameterUnsignedOutOfRange interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetEventParameter
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetTimeDelay returns TimeDelay (property field)
	GetTimeDelay() BACnetContextTagUnsignedInteger
	// GetLowLimit returns LowLimit (property field)
	GetLowLimit() BACnetContextTagUnsignedInteger
	// GetHighLimit returns HighLimit (property field)
	GetHighLimit() BACnetContextTagUnsignedInteger
	// GetDeadband returns Deadband (property field)
	GetDeadband() BACnetContextTagUnsignedInteger
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
	// IsBACnetEventParameterUnsignedOutOfRange is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetEventParameterUnsignedOutOfRange()
	// CreateBuilder creates a BACnetEventParameterUnsignedOutOfRangeBuilder
	CreateBACnetEventParameterUnsignedOutOfRangeBuilder() BACnetEventParameterUnsignedOutOfRangeBuilder
}

// _BACnetEventParameterUnsignedOutOfRange is the data-structure of this message
type _BACnetEventParameterUnsignedOutOfRange struct {
	BACnetEventParameterContract
	OpeningTag BACnetOpeningTag
	TimeDelay  BACnetContextTagUnsignedInteger
	LowLimit   BACnetContextTagUnsignedInteger
	HighLimit  BACnetContextTagUnsignedInteger
	Deadband   BACnetContextTagUnsignedInteger
	ClosingTag BACnetClosingTag
}

var _ BACnetEventParameterUnsignedOutOfRange = (*_BACnetEventParameterUnsignedOutOfRange)(nil)
var _ BACnetEventParameterRequirements = (*_BACnetEventParameterUnsignedOutOfRange)(nil)

// NewBACnetEventParameterUnsignedOutOfRange factory function for _BACnetEventParameterUnsignedOutOfRange
func NewBACnetEventParameterUnsignedOutOfRange(peekedTagHeader BACnetTagHeader, openingTag BACnetOpeningTag, timeDelay BACnetContextTagUnsignedInteger, lowLimit BACnetContextTagUnsignedInteger, highLimit BACnetContextTagUnsignedInteger, deadband BACnetContextTagUnsignedInteger, closingTag BACnetClosingTag) *_BACnetEventParameterUnsignedOutOfRange {
	if openingTag == nil {
		panic("openingTag of type BACnetOpeningTag for BACnetEventParameterUnsignedOutOfRange must not be nil")
	}
	if timeDelay == nil {
		panic("timeDelay of type BACnetContextTagUnsignedInteger for BACnetEventParameterUnsignedOutOfRange must not be nil")
	}
	if lowLimit == nil {
		panic("lowLimit of type BACnetContextTagUnsignedInteger for BACnetEventParameterUnsignedOutOfRange must not be nil")
	}
	if highLimit == nil {
		panic("highLimit of type BACnetContextTagUnsignedInteger for BACnetEventParameterUnsignedOutOfRange must not be nil")
	}
	if deadband == nil {
		panic("deadband of type BACnetContextTagUnsignedInteger for BACnetEventParameterUnsignedOutOfRange must not be nil")
	}
	if closingTag == nil {
		panic("closingTag of type BACnetClosingTag for BACnetEventParameterUnsignedOutOfRange must not be nil")
	}
	_result := &_BACnetEventParameterUnsignedOutOfRange{
		BACnetEventParameterContract: NewBACnetEventParameter(peekedTagHeader),
		OpeningTag:                   openingTag,
		TimeDelay:                    timeDelay,
		LowLimit:                     lowLimit,
		HighLimit:                    highLimit,
		Deadband:                     deadband,
		ClosingTag:                   closingTag,
	}
	_result.BACnetEventParameterContract.(*_BACnetEventParameter)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetEventParameterUnsignedOutOfRangeBuilder is a builder for BACnetEventParameterUnsignedOutOfRange
type BACnetEventParameterUnsignedOutOfRangeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(openingTag BACnetOpeningTag, timeDelay BACnetContextTagUnsignedInteger, lowLimit BACnetContextTagUnsignedInteger, highLimit BACnetContextTagUnsignedInteger, deadband BACnetContextTagUnsignedInteger, closingTag BACnetClosingTag) BACnetEventParameterUnsignedOutOfRangeBuilder
	// WithOpeningTag adds OpeningTag (property field)
	WithOpeningTag(BACnetOpeningTag) BACnetEventParameterUnsignedOutOfRangeBuilder
	// WithOpeningTagBuilder adds OpeningTag (property field) which is build by the builder
	WithOpeningTagBuilder(func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetEventParameterUnsignedOutOfRangeBuilder
	// WithTimeDelay adds TimeDelay (property field)
	WithTimeDelay(BACnetContextTagUnsignedInteger) BACnetEventParameterUnsignedOutOfRangeBuilder
	// WithTimeDelayBuilder adds TimeDelay (property field) which is build by the builder
	WithTimeDelayBuilder(func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetEventParameterUnsignedOutOfRangeBuilder
	// WithLowLimit adds LowLimit (property field)
	WithLowLimit(BACnetContextTagUnsignedInteger) BACnetEventParameterUnsignedOutOfRangeBuilder
	// WithLowLimitBuilder adds LowLimit (property field) which is build by the builder
	WithLowLimitBuilder(func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetEventParameterUnsignedOutOfRangeBuilder
	// WithHighLimit adds HighLimit (property field)
	WithHighLimit(BACnetContextTagUnsignedInteger) BACnetEventParameterUnsignedOutOfRangeBuilder
	// WithHighLimitBuilder adds HighLimit (property field) which is build by the builder
	WithHighLimitBuilder(func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetEventParameterUnsignedOutOfRangeBuilder
	// WithDeadband adds Deadband (property field)
	WithDeadband(BACnetContextTagUnsignedInteger) BACnetEventParameterUnsignedOutOfRangeBuilder
	// WithDeadbandBuilder adds Deadband (property field) which is build by the builder
	WithDeadbandBuilder(func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetEventParameterUnsignedOutOfRangeBuilder
	// WithClosingTag adds ClosingTag (property field)
	WithClosingTag(BACnetClosingTag) BACnetEventParameterUnsignedOutOfRangeBuilder
	// WithClosingTagBuilder adds ClosingTag (property field) which is build by the builder
	WithClosingTagBuilder(func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetEventParameterUnsignedOutOfRangeBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetEventParameterBuilder
	// Build builds the BACnetEventParameterUnsignedOutOfRange or returns an error if something is wrong
	Build() (BACnetEventParameterUnsignedOutOfRange, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetEventParameterUnsignedOutOfRange
}

// NewBACnetEventParameterUnsignedOutOfRangeBuilder() creates a BACnetEventParameterUnsignedOutOfRangeBuilder
func NewBACnetEventParameterUnsignedOutOfRangeBuilder() BACnetEventParameterUnsignedOutOfRangeBuilder {
	return &_BACnetEventParameterUnsignedOutOfRangeBuilder{_BACnetEventParameterUnsignedOutOfRange: new(_BACnetEventParameterUnsignedOutOfRange)}
}

type _BACnetEventParameterUnsignedOutOfRangeBuilder struct {
	*_BACnetEventParameterUnsignedOutOfRange

	parentBuilder *_BACnetEventParameterBuilder

	err *utils.MultiError
}

var _ (BACnetEventParameterUnsignedOutOfRangeBuilder) = (*_BACnetEventParameterUnsignedOutOfRangeBuilder)(nil)

func (b *_BACnetEventParameterUnsignedOutOfRangeBuilder) setParent(contract BACnetEventParameterContract) {
	b.BACnetEventParameterContract = contract
	contract.(*_BACnetEventParameter)._SubType = b._BACnetEventParameterUnsignedOutOfRange
}

func (b *_BACnetEventParameterUnsignedOutOfRangeBuilder) WithMandatoryFields(openingTag BACnetOpeningTag, timeDelay BACnetContextTagUnsignedInteger, lowLimit BACnetContextTagUnsignedInteger, highLimit BACnetContextTagUnsignedInteger, deadband BACnetContextTagUnsignedInteger, closingTag BACnetClosingTag) BACnetEventParameterUnsignedOutOfRangeBuilder {
	return b.WithOpeningTag(openingTag).WithTimeDelay(timeDelay).WithLowLimit(lowLimit).WithHighLimit(highLimit).WithDeadband(deadband).WithClosingTag(closingTag)
}

func (b *_BACnetEventParameterUnsignedOutOfRangeBuilder) WithOpeningTag(openingTag BACnetOpeningTag) BACnetEventParameterUnsignedOutOfRangeBuilder {
	b.OpeningTag = openingTag
	return b
}

func (b *_BACnetEventParameterUnsignedOutOfRangeBuilder) WithOpeningTagBuilder(builderSupplier func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetEventParameterUnsignedOutOfRangeBuilder {
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

func (b *_BACnetEventParameterUnsignedOutOfRangeBuilder) WithTimeDelay(timeDelay BACnetContextTagUnsignedInteger) BACnetEventParameterUnsignedOutOfRangeBuilder {
	b.TimeDelay = timeDelay
	return b
}

func (b *_BACnetEventParameterUnsignedOutOfRangeBuilder) WithTimeDelayBuilder(builderSupplier func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetEventParameterUnsignedOutOfRangeBuilder {
	builder := builderSupplier(b.TimeDelay.CreateBACnetContextTagUnsignedIntegerBuilder())
	var err error
	b.TimeDelay, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetEventParameterUnsignedOutOfRangeBuilder) WithLowLimit(lowLimit BACnetContextTagUnsignedInteger) BACnetEventParameterUnsignedOutOfRangeBuilder {
	b.LowLimit = lowLimit
	return b
}

func (b *_BACnetEventParameterUnsignedOutOfRangeBuilder) WithLowLimitBuilder(builderSupplier func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetEventParameterUnsignedOutOfRangeBuilder {
	builder := builderSupplier(b.LowLimit.CreateBACnetContextTagUnsignedIntegerBuilder())
	var err error
	b.LowLimit, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetEventParameterUnsignedOutOfRangeBuilder) WithHighLimit(highLimit BACnetContextTagUnsignedInteger) BACnetEventParameterUnsignedOutOfRangeBuilder {
	b.HighLimit = highLimit
	return b
}

func (b *_BACnetEventParameterUnsignedOutOfRangeBuilder) WithHighLimitBuilder(builderSupplier func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetEventParameterUnsignedOutOfRangeBuilder {
	builder := builderSupplier(b.HighLimit.CreateBACnetContextTagUnsignedIntegerBuilder())
	var err error
	b.HighLimit, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetEventParameterUnsignedOutOfRangeBuilder) WithDeadband(deadband BACnetContextTagUnsignedInteger) BACnetEventParameterUnsignedOutOfRangeBuilder {
	b.Deadband = deadband
	return b
}

func (b *_BACnetEventParameterUnsignedOutOfRangeBuilder) WithDeadbandBuilder(builderSupplier func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetEventParameterUnsignedOutOfRangeBuilder {
	builder := builderSupplier(b.Deadband.CreateBACnetContextTagUnsignedIntegerBuilder())
	var err error
	b.Deadband, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetEventParameterUnsignedOutOfRangeBuilder) WithClosingTag(closingTag BACnetClosingTag) BACnetEventParameterUnsignedOutOfRangeBuilder {
	b.ClosingTag = closingTag
	return b
}

func (b *_BACnetEventParameterUnsignedOutOfRangeBuilder) WithClosingTagBuilder(builderSupplier func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetEventParameterUnsignedOutOfRangeBuilder {
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

func (b *_BACnetEventParameterUnsignedOutOfRangeBuilder) Build() (BACnetEventParameterUnsignedOutOfRange, error) {
	if b.OpeningTag == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'openingTag' not set"))
	}
	if b.TimeDelay == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'timeDelay' not set"))
	}
	if b.LowLimit == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'lowLimit' not set"))
	}
	if b.HighLimit == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'highLimit' not set"))
	}
	if b.Deadband == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'deadband' not set"))
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
	return b._BACnetEventParameterUnsignedOutOfRange.deepCopy(), nil
}

func (b *_BACnetEventParameterUnsignedOutOfRangeBuilder) MustBuild() BACnetEventParameterUnsignedOutOfRange {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetEventParameterUnsignedOutOfRangeBuilder) Done() BACnetEventParameterBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetEventParameterBuilder().(*_BACnetEventParameterBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetEventParameterUnsignedOutOfRangeBuilder) buildForBACnetEventParameter() (BACnetEventParameter, error) {
	return b.Build()
}

func (b *_BACnetEventParameterUnsignedOutOfRangeBuilder) DeepCopy() any {
	_copy := b.CreateBACnetEventParameterUnsignedOutOfRangeBuilder().(*_BACnetEventParameterUnsignedOutOfRangeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetEventParameterUnsignedOutOfRangeBuilder creates a BACnetEventParameterUnsignedOutOfRangeBuilder
func (b *_BACnetEventParameterUnsignedOutOfRange) CreateBACnetEventParameterUnsignedOutOfRangeBuilder() BACnetEventParameterUnsignedOutOfRangeBuilder {
	if b == nil {
		return NewBACnetEventParameterUnsignedOutOfRangeBuilder()
	}
	return &_BACnetEventParameterUnsignedOutOfRangeBuilder{_BACnetEventParameterUnsignedOutOfRange: b.deepCopy()}
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

func (m *_BACnetEventParameterUnsignedOutOfRange) GetParent() BACnetEventParameterContract {
	return m.BACnetEventParameterContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetEventParameterUnsignedOutOfRange) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetEventParameterUnsignedOutOfRange) GetTimeDelay() BACnetContextTagUnsignedInteger {
	return m.TimeDelay
}

func (m *_BACnetEventParameterUnsignedOutOfRange) GetLowLimit() BACnetContextTagUnsignedInteger {
	return m.LowLimit
}

func (m *_BACnetEventParameterUnsignedOutOfRange) GetHighLimit() BACnetContextTagUnsignedInteger {
	return m.HighLimit
}

func (m *_BACnetEventParameterUnsignedOutOfRange) GetDeadband() BACnetContextTagUnsignedInteger {
	return m.Deadband
}

func (m *_BACnetEventParameterUnsignedOutOfRange) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetEventParameterUnsignedOutOfRange(structType any) BACnetEventParameterUnsignedOutOfRange {
	if casted, ok := structType.(BACnetEventParameterUnsignedOutOfRange); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetEventParameterUnsignedOutOfRange); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetEventParameterUnsignedOutOfRange) GetTypeName() string {
	return "BACnetEventParameterUnsignedOutOfRange"
}

func (m *_BACnetEventParameterUnsignedOutOfRange) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetEventParameterContract.(*_BACnetEventParameter).getLengthInBits(ctx))

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Simple field (timeDelay)
	lengthInBits += m.TimeDelay.GetLengthInBits(ctx)

	// Simple field (lowLimit)
	lengthInBits += m.LowLimit.GetLengthInBits(ctx)

	// Simple field (highLimit)
	lengthInBits += m.HighLimit.GetLengthInBits(ctx)

	// Simple field (deadband)
	lengthInBits += m.Deadband.GetLengthInBits(ctx)

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetEventParameterUnsignedOutOfRange) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetEventParameterUnsignedOutOfRange) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetEventParameter) (__bACnetEventParameterUnsignedOutOfRange BACnetEventParameterUnsignedOutOfRange, err error) {
	m.BACnetEventParameterContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetEventParameterUnsignedOutOfRange"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetEventParameterUnsignedOutOfRange")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(uint8(16))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	m.OpeningTag = openingTag

	timeDelay, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "timeDelay", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timeDelay' field"))
	}
	m.TimeDelay = timeDelay

	lowLimit, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "lowLimit", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(1)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lowLimit' field"))
	}
	m.LowLimit = lowLimit

	highLimit, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "highLimit", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(2)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'highLimit' field"))
	}
	m.HighLimit = highLimit

	deadband, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "deadband", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(3)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'deadband' field"))
	}
	m.Deadband = deadband

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(uint8(16))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("BACnetEventParameterUnsignedOutOfRange"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetEventParameterUnsignedOutOfRange")
	}

	return m, nil
}

func (m *_BACnetEventParameterUnsignedOutOfRange) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetEventParameterUnsignedOutOfRange) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetEventParameterUnsignedOutOfRange"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetEventParameterUnsignedOutOfRange")
		}

		if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'openingTag' field")
		}

		if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "timeDelay", m.GetTimeDelay(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'timeDelay' field")
		}

		if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "lowLimit", m.GetLowLimit(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'lowLimit' field")
		}

		if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "highLimit", m.GetHighLimit(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'highLimit' field")
		}

		if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "deadband", m.GetDeadband(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'deadband' field")
		}

		if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'closingTag' field")
		}

		if popErr := writeBuffer.PopContext("BACnetEventParameterUnsignedOutOfRange"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetEventParameterUnsignedOutOfRange")
		}
		return nil
	}
	return m.BACnetEventParameterContract.(*_BACnetEventParameter).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetEventParameterUnsignedOutOfRange) IsBACnetEventParameterUnsignedOutOfRange() {}

func (m *_BACnetEventParameterUnsignedOutOfRange) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetEventParameterUnsignedOutOfRange) deepCopy() *_BACnetEventParameterUnsignedOutOfRange {
	if m == nil {
		return nil
	}
	_BACnetEventParameterUnsignedOutOfRangeCopy := &_BACnetEventParameterUnsignedOutOfRange{
		m.BACnetEventParameterContract.(*_BACnetEventParameter).deepCopy(),
		utils.DeepCopy[BACnetOpeningTag](m.OpeningTag),
		utils.DeepCopy[BACnetContextTagUnsignedInteger](m.TimeDelay),
		utils.DeepCopy[BACnetContextTagUnsignedInteger](m.LowLimit),
		utils.DeepCopy[BACnetContextTagUnsignedInteger](m.HighLimit),
		utils.DeepCopy[BACnetContextTagUnsignedInteger](m.Deadband),
		utils.DeepCopy[BACnetClosingTag](m.ClosingTag),
	}
	m.BACnetEventParameterContract.(*_BACnetEventParameter)._SubType = m
	return _BACnetEventParameterUnsignedOutOfRangeCopy
}

func (m *_BACnetEventParameterUnsignedOutOfRange) String() string {
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
