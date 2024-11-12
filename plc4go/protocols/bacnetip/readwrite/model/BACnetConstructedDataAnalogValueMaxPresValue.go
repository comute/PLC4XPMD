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

// BACnetConstructedDataAnalogValueMaxPresValue is the corresponding interface of BACnetConstructedDataAnalogValueMaxPresValue
type BACnetConstructedDataAnalogValueMaxPresValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetMaxPresValue returns MaxPresValue (property field)
	GetMaxPresValue() BACnetApplicationTagReal
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagReal
	// IsBACnetConstructedDataAnalogValueMaxPresValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataAnalogValueMaxPresValue()
	// CreateBuilder creates a BACnetConstructedDataAnalogValueMaxPresValueBuilder
	CreateBACnetConstructedDataAnalogValueMaxPresValueBuilder() BACnetConstructedDataAnalogValueMaxPresValueBuilder
}

// _BACnetConstructedDataAnalogValueMaxPresValue is the data-structure of this message
type _BACnetConstructedDataAnalogValueMaxPresValue struct {
	BACnetConstructedDataContract
	MaxPresValue BACnetApplicationTagReal
}

var _ BACnetConstructedDataAnalogValueMaxPresValue = (*_BACnetConstructedDataAnalogValueMaxPresValue)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataAnalogValueMaxPresValue)(nil)

// NewBACnetConstructedDataAnalogValueMaxPresValue factory function for _BACnetConstructedDataAnalogValueMaxPresValue
func NewBACnetConstructedDataAnalogValueMaxPresValue(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, maxPresValue BACnetApplicationTagReal, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAnalogValueMaxPresValue {
	if maxPresValue == nil {
		panic("maxPresValue of type BACnetApplicationTagReal for BACnetConstructedDataAnalogValueMaxPresValue must not be nil")
	}
	_result := &_BACnetConstructedDataAnalogValueMaxPresValue{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		MaxPresValue:                  maxPresValue,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataAnalogValueMaxPresValueBuilder is a builder for BACnetConstructedDataAnalogValueMaxPresValue
type BACnetConstructedDataAnalogValueMaxPresValueBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(maxPresValue BACnetApplicationTagReal) BACnetConstructedDataAnalogValueMaxPresValueBuilder
	// WithMaxPresValue adds MaxPresValue (property field)
	WithMaxPresValue(BACnetApplicationTagReal) BACnetConstructedDataAnalogValueMaxPresValueBuilder
	// WithMaxPresValueBuilder adds MaxPresValue (property field) which is build by the builder
	WithMaxPresValueBuilder(func(BACnetApplicationTagRealBuilder) BACnetApplicationTagRealBuilder) BACnetConstructedDataAnalogValueMaxPresValueBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataAnalogValueMaxPresValue or returns an error if something is wrong
	Build() (BACnetConstructedDataAnalogValueMaxPresValue, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataAnalogValueMaxPresValue
}

// NewBACnetConstructedDataAnalogValueMaxPresValueBuilder() creates a BACnetConstructedDataAnalogValueMaxPresValueBuilder
func NewBACnetConstructedDataAnalogValueMaxPresValueBuilder() BACnetConstructedDataAnalogValueMaxPresValueBuilder {
	return &_BACnetConstructedDataAnalogValueMaxPresValueBuilder{_BACnetConstructedDataAnalogValueMaxPresValue: new(_BACnetConstructedDataAnalogValueMaxPresValue)}
}

type _BACnetConstructedDataAnalogValueMaxPresValueBuilder struct {
	*_BACnetConstructedDataAnalogValueMaxPresValue

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataAnalogValueMaxPresValueBuilder) = (*_BACnetConstructedDataAnalogValueMaxPresValueBuilder)(nil)

func (b *_BACnetConstructedDataAnalogValueMaxPresValueBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataAnalogValueMaxPresValue
}

func (b *_BACnetConstructedDataAnalogValueMaxPresValueBuilder) WithMandatoryFields(maxPresValue BACnetApplicationTagReal) BACnetConstructedDataAnalogValueMaxPresValueBuilder {
	return b.WithMaxPresValue(maxPresValue)
}

func (b *_BACnetConstructedDataAnalogValueMaxPresValueBuilder) WithMaxPresValue(maxPresValue BACnetApplicationTagReal) BACnetConstructedDataAnalogValueMaxPresValueBuilder {
	b.MaxPresValue = maxPresValue
	return b
}

func (b *_BACnetConstructedDataAnalogValueMaxPresValueBuilder) WithMaxPresValueBuilder(builderSupplier func(BACnetApplicationTagRealBuilder) BACnetApplicationTagRealBuilder) BACnetConstructedDataAnalogValueMaxPresValueBuilder {
	builder := builderSupplier(b.MaxPresValue.CreateBACnetApplicationTagRealBuilder())
	var err error
	b.MaxPresValue, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagRealBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataAnalogValueMaxPresValueBuilder) Build() (BACnetConstructedDataAnalogValueMaxPresValue, error) {
	if b.MaxPresValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'maxPresValue' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataAnalogValueMaxPresValue.deepCopy(), nil
}

func (b *_BACnetConstructedDataAnalogValueMaxPresValueBuilder) MustBuild() BACnetConstructedDataAnalogValueMaxPresValue {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataAnalogValueMaxPresValueBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataAnalogValueMaxPresValueBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataAnalogValueMaxPresValueBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataAnalogValueMaxPresValueBuilder().(*_BACnetConstructedDataAnalogValueMaxPresValueBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataAnalogValueMaxPresValueBuilder creates a BACnetConstructedDataAnalogValueMaxPresValueBuilder
func (b *_BACnetConstructedDataAnalogValueMaxPresValue) CreateBACnetConstructedDataAnalogValueMaxPresValueBuilder() BACnetConstructedDataAnalogValueMaxPresValueBuilder {
	if b == nil {
		return NewBACnetConstructedDataAnalogValueMaxPresValueBuilder()
	}
	return &_BACnetConstructedDataAnalogValueMaxPresValueBuilder{_BACnetConstructedDataAnalogValueMaxPresValue: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAnalogValueMaxPresValue) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_ANALOG_VALUE
}

func (m *_BACnetConstructedDataAnalogValueMaxPresValue) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_MAX_PRES_VALUE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAnalogValueMaxPresValue) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataAnalogValueMaxPresValue) GetMaxPresValue() BACnetApplicationTagReal {
	return m.MaxPresValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataAnalogValueMaxPresValue) GetActualValue() BACnetApplicationTagReal {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagReal(m.GetMaxPresValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAnalogValueMaxPresValue(structType any) BACnetConstructedDataAnalogValueMaxPresValue {
	if casted, ok := structType.(BACnetConstructedDataAnalogValueMaxPresValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAnalogValueMaxPresValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAnalogValueMaxPresValue) GetTypeName() string {
	return "BACnetConstructedDataAnalogValueMaxPresValue"
}

func (m *_BACnetConstructedDataAnalogValueMaxPresValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (maxPresValue)
	lengthInBits += m.MaxPresValue.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataAnalogValueMaxPresValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataAnalogValueMaxPresValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataAnalogValueMaxPresValue BACnetConstructedDataAnalogValueMaxPresValue, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAnalogValueMaxPresValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAnalogValueMaxPresValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	maxPresValue, err := ReadSimpleField[BACnetApplicationTagReal](ctx, "maxPresValue", ReadComplex[BACnetApplicationTagReal](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagReal](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'maxPresValue' field"))
	}
	m.MaxPresValue = maxPresValue

	actualValue, err := ReadVirtualField[BACnetApplicationTagReal](ctx, "actualValue", (*BACnetApplicationTagReal)(nil), maxPresValue)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAnalogValueMaxPresValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAnalogValueMaxPresValue")
	}

	return m, nil
}

func (m *_BACnetConstructedDataAnalogValueMaxPresValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataAnalogValueMaxPresValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAnalogValueMaxPresValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAnalogValueMaxPresValue")
		}

		if err := WriteSimpleField[BACnetApplicationTagReal](ctx, "maxPresValue", m.GetMaxPresValue(), WriteComplex[BACnetApplicationTagReal](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'maxPresValue' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAnalogValueMaxPresValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAnalogValueMaxPresValue")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAnalogValueMaxPresValue) IsBACnetConstructedDataAnalogValueMaxPresValue() {
}

func (m *_BACnetConstructedDataAnalogValueMaxPresValue) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataAnalogValueMaxPresValue) deepCopy() *_BACnetConstructedDataAnalogValueMaxPresValue {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataAnalogValueMaxPresValueCopy := &_BACnetConstructedDataAnalogValueMaxPresValue{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagReal](m.MaxPresValue),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataAnalogValueMaxPresValueCopy
}

func (m *_BACnetConstructedDataAnalogValueMaxPresValue) String() string {
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
