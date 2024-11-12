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

// BACnetConstructedDataShedDuration is the corresponding interface of BACnetConstructedDataShedDuration
type BACnetConstructedDataShedDuration interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetShedDuration returns ShedDuration (property field)
	GetShedDuration() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetConstructedDataShedDuration is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataShedDuration()
	// CreateBuilder creates a BACnetConstructedDataShedDurationBuilder
	CreateBACnetConstructedDataShedDurationBuilder() BACnetConstructedDataShedDurationBuilder
}

// _BACnetConstructedDataShedDuration is the data-structure of this message
type _BACnetConstructedDataShedDuration struct {
	BACnetConstructedDataContract
	ShedDuration BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataShedDuration = (*_BACnetConstructedDataShedDuration)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataShedDuration)(nil)

// NewBACnetConstructedDataShedDuration factory function for _BACnetConstructedDataShedDuration
func NewBACnetConstructedDataShedDuration(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, shedDuration BACnetApplicationTagUnsignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataShedDuration {
	if shedDuration == nil {
		panic("shedDuration of type BACnetApplicationTagUnsignedInteger for BACnetConstructedDataShedDuration must not be nil")
	}
	_result := &_BACnetConstructedDataShedDuration{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		ShedDuration:                  shedDuration,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataShedDurationBuilder is a builder for BACnetConstructedDataShedDuration
type BACnetConstructedDataShedDurationBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(shedDuration BACnetApplicationTagUnsignedInteger) BACnetConstructedDataShedDurationBuilder
	// WithShedDuration adds ShedDuration (property field)
	WithShedDuration(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataShedDurationBuilder
	// WithShedDurationBuilder adds ShedDuration (property field) which is build by the builder
	WithShedDurationBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataShedDurationBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataShedDuration or returns an error if something is wrong
	Build() (BACnetConstructedDataShedDuration, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataShedDuration
}

// NewBACnetConstructedDataShedDurationBuilder() creates a BACnetConstructedDataShedDurationBuilder
func NewBACnetConstructedDataShedDurationBuilder() BACnetConstructedDataShedDurationBuilder {
	return &_BACnetConstructedDataShedDurationBuilder{_BACnetConstructedDataShedDuration: new(_BACnetConstructedDataShedDuration)}
}

type _BACnetConstructedDataShedDurationBuilder struct {
	*_BACnetConstructedDataShedDuration

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataShedDurationBuilder) = (*_BACnetConstructedDataShedDurationBuilder)(nil)

func (b *_BACnetConstructedDataShedDurationBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataShedDuration
}

func (b *_BACnetConstructedDataShedDurationBuilder) WithMandatoryFields(shedDuration BACnetApplicationTagUnsignedInteger) BACnetConstructedDataShedDurationBuilder {
	return b.WithShedDuration(shedDuration)
}

func (b *_BACnetConstructedDataShedDurationBuilder) WithShedDuration(shedDuration BACnetApplicationTagUnsignedInteger) BACnetConstructedDataShedDurationBuilder {
	b.ShedDuration = shedDuration
	return b
}

func (b *_BACnetConstructedDataShedDurationBuilder) WithShedDurationBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataShedDurationBuilder {
	builder := builderSupplier(b.ShedDuration.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	b.ShedDuration, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataShedDurationBuilder) Build() (BACnetConstructedDataShedDuration, error) {
	if b.ShedDuration == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'shedDuration' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataShedDuration.deepCopy(), nil
}

func (b *_BACnetConstructedDataShedDurationBuilder) MustBuild() BACnetConstructedDataShedDuration {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataShedDurationBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataShedDurationBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataShedDurationBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataShedDurationBuilder().(*_BACnetConstructedDataShedDurationBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataShedDurationBuilder creates a BACnetConstructedDataShedDurationBuilder
func (b *_BACnetConstructedDataShedDuration) CreateBACnetConstructedDataShedDurationBuilder() BACnetConstructedDataShedDurationBuilder {
	if b == nil {
		return NewBACnetConstructedDataShedDurationBuilder()
	}
	return &_BACnetConstructedDataShedDurationBuilder{_BACnetConstructedDataShedDuration: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataShedDuration) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataShedDuration) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_SHED_DURATION
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataShedDuration) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataShedDuration) GetShedDuration() BACnetApplicationTagUnsignedInteger {
	return m.ShedDuration
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataShedDuration) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetShedDuration())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataShedDuration(structType any) BACnetConstructedDataShedDuration {
	if casted, ok := structType.(BACnetConstructedDataShedDuration); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataShedDuration); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataShedDuration) GetTypeName() string {
	return "BACnetConstructedDataShedDuration"
}

func (m *_BACnetConstructedDataShedDuration) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (shedDuration)
	lengthInBits += m.ShedDuration.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataShedDuration) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataShedDuration) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataShedDuration BACnetConstructedDataShedDuration, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataShedDuration"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataShedDuration")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	shedDuration, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "shedDuration", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'shedDuration' field"))
	}
	m.ShedDuration = shedDuration

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), shedDuration)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataShedDuration"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataShedDuration")
	}

	return m, nil
}

func (m *_BACnetConstructedDataShedDuration) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataShedDuration) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataShedDuration"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataShedDuration")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "shedDuration", m.GetShedDuration(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'shedDuration' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataShedDuration"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataShedDuration")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataShedDuration) IsBACnetConstructedDataShedDuration() {}

func (m *_BACnetConstructedDataShedDuration) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataShedDuration) deepCopy() *_BACnetConstructedDataShedDuration {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataShedDurationCopy := &_BACnetConstructedDataShedDuration{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagUnsignedInteger](m.ShedDuration),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataShedDurationCopy
}

func (m *_BACnetConstructedDataShedDuration) String() string {
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
