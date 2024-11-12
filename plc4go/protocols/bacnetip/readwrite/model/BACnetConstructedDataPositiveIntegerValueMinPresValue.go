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

// BACnetConstructedDataPositiveIntegerValueMinPresValue is the corresponding interface of BACnetConstructedDataPositiveIntegerValueMinPresValue
type BACnetConstructedDataPositiveIntegerValueMinPresValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetMinPresValue returns MinPresValue (property field)
	GetMinPresValue() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetConstructedDataPositiveIntegerValueMinPresValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataPositiveIntegerValueMinPresValue()
	// CreateBuilder creates a BACnetConstructedDataPositiveIntegerValueMinPresValueBuilder
	CreateBACnetConstructedDataPositiveIntegerValueMinPresValueBuilder() BACnetConstructedDataPositiveIntegerValueMinPresValueBuilder
}

// _BACnetConstructedDataPositiveIntegerValueMinPresValue is the data-structure of this message
type _BACnetConstructedDataPositiveIntegerValueMinPresValue struct {
	BACnetConstructedDataContract
	MinPresValue BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataPositiveIntegerValueMinPresValue = (*_BACnetConstructedDataPositiveIntegerValueMinPresValue)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataPositiveIntegerValueMinPresValue)(nil)

// NewBACnetConstructedDataPositiveIntegerValueMinPresValue factory function for _BACnetConstructedDataPositiveIntegerValueMinPresValue
func NewBACnetConstructedDataPositiveIntegerValueMinPresValue(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, minPresValue BACnetApplicationTagUnsignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataPositiveIntegerValueMinPresValue {
	if minPresValue == nil {
		panic("minPresValue of type BACnetApplicationTagUnsignedInteger for BACnetConstructedDataPositiveIntegerValueMinPresValue must not be nil")
	}
	_result := &_BACnetConstructedDataPositiveIntegerValueMinPresValue{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		MinPresValue:                  minPresValue,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataPositiveIntegerValueMinPresValueBuilder is a builder for BACnetConstructedDataPositiveIntegerValueMinPresValue
type BACnetConstructedDataPositiveIntegerValueMinPresValueBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(minPresValue BACnetApplicationTagUnsignedInteger) BACnetConstructedDataPositiveIntegerValueMinPresValueBuilder
	// WithMinPresValue adds MinPresValue (property field)
	WithMinPresValue(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataPositiveIntegerValueMinPresValueBuilder
	// WithMinPresValueBuilder adds MinPresValue (property field) which is build by the builder
	WithMinPresValueBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataPositiveIntegerValueMinPresValueBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataPositiveIntegerValueMinPresValue or returns an error if something is wrong
	Build() (BACnetConstructedDataPositiveIntegerValueMinPresValue, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataPositiveIntegerValueMinPresValue
}

// NewBACnetConstructedDataPositiveIntegerValueMinPresValueBuilder() creates a BACnetConstructedDataPositiveIntegerValueMinPresValueBuilder
func NewBACnetConstructedDataPositiveIntegerValueMinPresValueBuilder() BACnetConstructedDataPositiveIntegerValueMinPresValueBuilder {
	return &_BACnetConstructedDataPositiveIntegerValueMinPresValueBuilder{_BACnetConstructedDataPositiveIntegerValueMinPresValue: new(_BACnetConstructedDataPositiveIntegerValueMinPresValue)}
}

type _BACnetConstructedDataPositiveIntegerValueMinPresValueBuilder struct {
	*_BACnetConstructedDataPositiveIntegerValueMinPresValue

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataPositiveIntegerValueMinPresValueBuilder) = (*_BACnetConstructedDataPositiveIntegerValueMinPresValueBuilder)(nil)

func (b *_BACnetConstructedDataPositiveIntegerValueMinPresValueBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataPositiveIntegerValueMinPresValue
}

func (b *_BACnetConstructedDataPositiveIntegerValueMinPresValueBuilder) WithMandatoryFields(minPresValue BACnetApplicationTagUnsignedInteger) BACnetConstructedDataPositiveIntegerValueMinPresValueBuilder {
	return b.WithMinPresValue(minPresValue)
}

func (b *_BACnetConstructedDataPositiveIntegerValueMinPresValueBuilder) WithMinPresValue(minPresValue BACnetApplicationTagUnsignedInteger) BACnetConstructedDataPositiveIntegerValueMinPresValueBuilder {
	b.MinPresValue = minPresValue
	return b
}

func (b *_BACnetConstructedDataPositiveIntegerValueMinPresValueBuilder) WithMinPresValueBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataPositiveIntegerValueMinPresValueBuilder {
	builder := builderSupplier(b.MinPresValue.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	b.MinPresValue, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataPositiveIntegerValueMinPresValueBuilder) Build() (BACnetConstructedDataPositiveIntegerValueMinPresValue, error) {
	if b.MinPresValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'minPresValue' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataPositiveIntegerValueMinPresValue.deepCopy(), nil
}

func (b *_BACnetConstructedDataPositiveIntegerValueMinPresValueBuilder) MustBuild() BACnetConstructedDataPositiveIntegerValueMinPresValue {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataPositiveIntegerValueMinPresValueBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataPositiveIntegerValueMinPresValueBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataPositiveIntegerValueMinPresValueBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataPositiveIntegerValueMinPresValueBuilder().(*_BACnetConstructedDataPositiveIntegerValueMinPresValueBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataPositiveIntegerValueMinPresValueBuilder creates a BACnetConstructedDataPositiveIntegerValueMinPresValueBuilder
func (b *_BACnetConstructedDataPositiveIntegerValueMinPresValue) CreateBACnetConstructedDataPositiveIntegerValueMinPresValueBuilder() BACnetConstructedDataPositiveIntegerValueMinPresValueBuilder {
	if b == nil {
		return NewBACnetConstructedDataPositiveIntegerValueMinPresValueBuilder()
	}
	return &_BACnetConstructedDataPositiveIntegerValueMinPresValueBuilder{_BACnetConstructedDataPositiveIntegerValueMinPresValue: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataPositiveIntegerValueMinPresValue) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_POSITIVE_INTEGER_VALUE
}

func (m *_BACnetConstructedDataPositiveIntegerValueMinPresValue) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_MIN_PRES_VALUE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataPositiveIntegerValueMinPresValue) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataPositiveIntegerValueMinPresValue) GetMinPresValue() BACnetApplicationTagUnsignedInteger {
	return m.MinPresValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataPositiveIntegerValueMinPresValue) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetMinPresValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataPositiveIntegerValueMinPresValue(structType any) BACnetConstructedDataPositiveIntegerValueMinPresValue {
	if casted, ok := structType.(BACnetConstructedDataPositiveIntegerValueMinPresValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataPositiveIntegerValueMinPresValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataPositiveIntegerValueMinPresValue) GetTypeName() string {
	return "BACnetConstructedDataPositiveIntegerValueMinPresValue"
}

func (m *_BACnetConstructedDataPositiveIntegerValueMinPresValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (minPresValue)
	lengthInBits += m.MinPresValue.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataPositiveIntegerValueMinPresValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataPositiveIntegerValueMinPresValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataPositiveIntegerValueMinPresValue BACnetConstructedDataPositiveIntegerValueMinPresValue, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataPositiveIntegerValueMinPresValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataPositiveIntegerValueMinPresValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	minPresValue, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "minPresValue", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'minPresValue' field"))
	}
	m.MinPresValue = minPresValue

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), minPresValue)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataPositiveIntegerValueMinPresValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataPositiveIntegerValueMinPresValue")
	}

	return m, nil
}

func (m *_BACnetConstructedDataPositiveIntegerValueMinPresValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataPositiveIntegerValueMinPresValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataPositiveIntegerValueMinPresValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataPositiveIntegerValueMinPresValue")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "minPresValue", m.GetMinPresValue(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'minPresValue' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataPositiveIntegerValueMinPresValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataPositiveIntegerValueMinPresValue")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataPositiveIntegerValueMinPresValue) IsBACnetConstructedDataPositiveIntegerValueMinPresValue() {
}

func (m *_BACnetConstructedDataPositiveIntegerValueMinPresValue) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataPositiveIntegerValueMinPresValue) deepCopy() *_BACnetConstructedDataPositiveIntegerValueMinPresValue {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataPositiveIntegerValueMinPresValueCopy := &_BACnetConstructedDataPositiveIntegerValueMinPresValue{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagUnsignedInteger](m.MinPresValue),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataPositiveIntegerValueMinPresValueCopy
}

func (m *_BACnetConstructedDataPositiveIntegerValueMinPresValue) String() string {
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
