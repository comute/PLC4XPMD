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

// BACnetConstructedDataDatePatternValuePresentValue is the corresponding interface of BACnetConstructedDataDatePatternValuePresentValue
type BACnetConstructedDataDatePatternValuePresentValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetPresentValue returns PresentValue (property field)
	GetPresentValue() BACnetApplicationTagDate
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagDate
	// IsBACnetConstructedDataDatePatternValuePresentValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataDatePatternValuePresentValue()
	// CreateBuilder creates a BACnetConstructedDataDatePatternValuePresentValueBuilder
	CreateBACnetConstructedDataDatePatternValuePresentValueBuilder() BACnetConstructedDataDatePatternValuePresentValueBuilder
}

// _BACnetConstructedDataDatePatternValuePresentValue is the data-structure of this message
type _BACnetConstructedDataDatePatternValuePresentValue struct {
	BACnetConstructedDataContract
	PresentValue BACnetApplicationTagDate
}

var _ BACnetConstructedDataDatePatternValuePresentValue = (*_BACnetConstructedDataDatePatternValuePresentValue)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataDatePatternValuePresentValue)(nil)

// NewBACnetConstructedDataDatePatternValuePresentValue factory function for _BACnetConstructedDataDatePatternValuePresentValue
func NewBACnetConstructedDataDatePatternValuePresentValue(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, presentValue BACnetApplicationTagDate, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataDatePatternValuePresentValue {
	if presentValue == nil {
		panic("presentValue of type BACnetApplicationTagDate for BACnetConstructedDataDatePatternValuePresentValue must not be nil")
	}
	_result := &_BACnetConstructedDataDatePatternValuePresentValue{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		PresentValue:                  presentValue,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataDatePatternValuePresentValueBuilder is a builder for BACnetConstructedDataDatePatternValuePresentValue
type BACnetConstructedDataDatePatternValuePresentValueBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(presentValue BACnetApplicationTagDate) BACnetConstructedDataDatePatternValuePresentValueBuilder
	// WithPresentValue adds PresentValue (property field)
	WithPresentValue(BACnetApplicationTagDate) BACnetConstructedDataDatePatternValuePresentValueBuilder
	// WithPresentValueBuilder adds PresentValue (property field) which is build by the builder
	WithPresentValueBuilder(func(BACnetApplicationTagDateBuilder) BACnetApplicationTagDateBuilder) BACnetConstructedDataDatePatternValuePresentValueBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataDatePatternValuePresentValue or returns an error if something is wrong
	Build() (BACnetConstructedDataDatePatternValuePresentValue, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataDatePatternValuePresentValue
}

// NewBACnetConstructedDataDatePatternValuePresentValueBuilder() creates a BACnetConstructedDataDatePatternValuePresentValueBuilder
func NewBACnetConstructedDataDatePatternValuePresentValueBuilder() BACnetConstructedDataDatePatternValuePresentValueBuilder {
	return &_BACnetConstructedDataDatePatternValuePresentValueBuilder{_BACnetConstructedDataDatePatternValuePresentValue: new(_BACnetConstructedDataDatePatternValuePresentValue)}
}

type _BACnetConstructedDataDatePatternValuePresentValueBuilder struct {
	*_BACnetConstructedDataDatePatternValuePresentValue

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataDatePatternValuePresentValueBuilder) = (*_BACnetConstructedDataDatePatternValuePresentValueBuilder)(nil)

func (b *_BACnetConstructedDataDatePatternValuePresentValueBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataDatePatternValuePresentValue
}

func (b *_BACnetConstructedDataDatePatternValuePresentValueBuilder) WithMandatoryFields(presentValue BACnetApplicationTagDate) BACnetConstructedDataDatePatternValuePresentValueBuilder {
	return b.WithPresentValue(presentValue)
}

func (b *_BACnetConstructedDataDatePatternValuePresentValueBuilder) WithPresentValue(presentValue BACnetApplicationTagDate) BACnetConstructedDataDatePatternValuePresentValueBuilder {
	b.PresentValue = presentValue
	return b
}

func (b *_BACnetConstructedDataDatePatternValuePresentValueBuilder) WithPresentValueBuilder(builderSupplier func(BACnetApplicationTagDateBuilder) BACnetApplicationTagDateBuilder) BACnetConstructedDataDatePatternValuePresentValueBuilder {
	builder := builderSupplier(b.PresentValue.CreateBACnetApplicationTagDateBuilder())
	var err error
	b.PresentValue, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagDateBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataDatePatternValuePresentValueBuilder) Build() (BACnetConstructedDataDatePatternValuePresentValue, error) {
	if b.PresentValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'presentValue' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataDatePatternValuePresentValue.deepCopy(), nil
}

func (b *_BACnetConstructedDataDatePatternValuePresentValueBuilder) MustBuild() BACnetConstructedDataDatePatternValuePresentValue {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataDatePatternValuePresentValueBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataDatePatternValuePresentValueBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataDatePatternValuePresentValueBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataDatePatternValuePresentValueBuilder().(*_BACnetConstructedDataDatePatternValuePresentValueBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataDatePatternValuePresentValueBuilder creates a BACnetConstructedDataDatePatternValuePresentValueBuilder
func (b *_BACnetConstructedDataDatePatternValuePresentValue) CreateBACnetConstructedDataDatePatternValuePresentValueBuilder() BACnetConstructedDataDatePatternValuePresentValueBuilder {
	if b == nil {
		return NewBACnetConstructedDataDatePatternValuePresentValueBuilder()
	}
	return &_BACnetConstructedDataDatePatternValuePresentValueBuilder{_BACnetConstructedDataDatePatternValuePresentValue: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataDatePatternValuePresentValue) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_DATEPATTERN_VALUE
}

func (m *_BACnetConstructedDataDatePatternValuePresentValue) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_PRESENT_VALUE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataDatePatternValuePresentValue) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataDatePatternValuePresentValue) GetPresentValue() BACnetApplicationTagDate {
	return m.PresentValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataDatePatternValuePresentValue) GetActualValue() BACnetApplicationTagDate {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagDate(m.GetPresentValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataDatePatternValuePresentValue(structType any) BACnetConstructedDataDatePatternValuePresentValue {
	if casted, ok := structType.(BACnetConstructedDataDatePatternValuePresentValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataDatePatternValuePresentValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataDatePatternValuePresentValue) GetTypeName() string {
	return "BACnetConstructedDataDatePatternValuePresentValue"
}

func (m *_BACnetConstructedDataDatePatternValuePresentValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (presentValue)
	lengthInBits += m.PresentValue.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataDatePatternValuePresentValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataDatePatternValuePresentValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataDatePatternValuePresentValue BACnetConstructedDataDatePatternValuePresentValue, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataDatePatternValuePresentValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataDatePatternValuePresentValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	presentValue, err := ReadSimpleField[BACnetApplicationTagDate](ctx, "presentValue", ReadComplex[BACnetApplicationTagDate](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagDate](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'presentValue' field"))
	}
	m.PresentValue = presentValue

	actualValue, err := ReadVirtualField[BACnetApplicationTagDate](ctx, "actualValue", (*BACnetApplicationTagDate)(nil), presentValue)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataDatePatternValuePresentValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataDatePatternValuePresentValue")
	}

	return m, nil
}

func (m *_BACnetConstructedDataDatePatternValuePresentValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataDatePatternValuePresentValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataDatePatternValuePresentValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataDatePatternValuePresentValue")
		}

		if err := WriteSimpleField[BACnetApplicationTagDate](ctx, "presentValue", m.GetPresentValue(), WriteComplex[BACnetApplicationTagDate](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'presentValue' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataDatePatternValuePresentValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataDatePatternValuePresentValue")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataDatePatternValuePresentValue) IsBACnetConstructedDataDatePatternValuePresentValue() {
}

func (m *_BACnetConstructedDataDatePatternValuePresentValue) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataDatePatternValuePresentValue) deepCopy() *_BACnetConstructedDataDatePatternValuePresentValue {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataDatePatternValuePresentValueCopy := &_BACnetConstructedDataDatePatternValuePresentValue{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagDate](m.PresentValue),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataDatePatternValuePresentValueCopy
}

func (m *_BACnetConstructedDataDatePatternValuePresentValue) String() string {
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
