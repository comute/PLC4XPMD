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

// BACnetConstructedDataValueBeforeChange is the corresponding interface of BACnetConstructedDataValueBeforeChange
type BACnetConstructedDataValueBeforeChange interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetValuesBeforeChange returns ValuesBeforeChange (property field)
	GetValuesBeforeChange() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetConstructedDataValueBeforeChange is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataValueBeforeChange()
	// CreateBuilder creates a BACnetConstructedDataValueBeforeChangeBuilder
	CreateBACnetConstructedDataValueBeforeChangeBuilder() BACnetConstructedDataValueBeforeChangeBuilder
}

// _BACnetConstructedDataValueBeforeChange is the data-structure of this message
type _BACnetConstructedDataValueBeforeChange struct {
	BACnetConstructedDataContract
	ValuesBeforeChange BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataValueBeforeChange = (*_BACnetConstructedDataValueBeforeChange)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataValueBeforeChange)(nil)

// NewBACnetConstructedDataValueBeforeChange factory function for _BACnetConstructedDataValueBeforeChange
func NewBACnetConstructedDataValueBeforeChange(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, valuesBeforeChange BACnetApplicationTagUnsignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataValueBeforeChange {
	if valuesBeforeChange == nil {
		panic("valuesBeforeChange of type BACnetApplicationTagUnsignedInteger for BACnetConstructedDataValueBeforeChange must not be nil")
	}
	_result := &_BACnetConstructedDataValueBeforeChange{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		ValuesBeforeChange:            valuesBeforeChange,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataValueBeforeChangeBuilder is a builder for BACnetConstructedDataValueBeforeChange
type BACnetConstructedDataValueBeforeChangeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(valuesBeforeChange BACnetApplicationTagUnsignedInteger) BACnetConstructedDataValueBeforeChangeBuilder
	// WithValuesBeforeChange adds ValuesBeforeChange (property field)
	WithValuesBeforeChange(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataValueBeforeChangeBuilder
	// WithValuesBeforeChangeBuilder adds ValuesBeforeChange (property field) which is build by the builder
	WithValuesBeforeChangeBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataValueBeforeChangeBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataValueBeforeChange or returns an error if something is wrong
	Build() (BACnetConstructedDataValueBeforeChange, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataValueBeforeChange
}

// NewBACnetConstructedDataValueBeforeChangeBuilder() creates a BACnetConstructedDataValueBeforeChangeBuilder
func NewBACnetConstructedDataValueBeforeChangeBuilder() BACnetConstructedDataValueBeforeChangeBuilder {
	return &_BACnetConstructedDataValueBeforeChangeBuilder{_BACnetConstructedDataValueBeforeChange: new(_BACnetConstructedDataValueBeforeChange)}
}

type _BACnetConstructedDataValueBeforeChangeBuilder struct {
	*_BACnetConstructedDataValueBeforeChange

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataValueBeforeChangeBuilder) = (*_BACnetConstructedDataValueBeforeChangeBuilder)(nil)

func (b *_BACnetConstructedDataValueBeforeChangeBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataValueBeforeChange
}

func (b *_BACnetConstructedDataValueBeforeChangeBuilder) WithMandatoryFields(valuesBeforeChange BACnetApplicationTagUnsignedInteger) BACnetConstructedDataValueBeforeChangeBuilder {
	return b.WithValuesBeforeChange(valuesBeforeChange)
}

func (b *_BACnetConstructedDataValueBeforeChangeBuilder) WithValuesBeforeChange(valuesBeforeChange BACnetApplicationTagUnsignedInteger) BACnetConstructedDataValueBeforeChangeBuilder {
	b.ValuesBeforeChange = valuesBeforeChange
	return b
}

func (b *_BACnetConstructedDataValueBeforeChangeBuilder) WithValuesBeforeChangeBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataValueBeforeChangeBuilder {
	builder := builderSupplier(b.ValuesBeforeChange.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	b.ValuesBeforeChange, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataValueBeforeChangeBuilder) Build() (BACnetConstructedDataValueBeforeChange, error) {
	if b.ValuesBeforeChange == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'valuesBeforeChange' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataValueBeforeChange.deepCopy(), nil
}

func (b *_BACnetConstructedDataValueBeforeChangeBuilder) MustBuild() BACnetConstructedDataValueBeforeChange {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataValueBeforeChangeBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataValueBeforeChangeBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataValueBeforeChangeBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataValueBeforeChangeBuilder().(*_BACnetConstructedDataValueBeforeChangeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataValueBeforeChangeBuilder creates a BACnetConstructedDataValueBeforeChangeBuilder
func (b *_BACnetConstructedDataValueBeforeChange) CreateBACnetConstructedDataValueBeforeChangeBuilder() BACnetConstructedDataValueBeforeChangeBuilder {
	if b == nil {
		return NewBACnetConstructedDataValueBeforeChangeBuilder()
	}
	return &_BACnetConstructedDataValueBeforeChangeBuilder{_BACnetConstructedDataValueBeforeChange: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataValueBeforeChange) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataValueBeforeChange) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_VALUE_BEFORE_CHANGE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataValueBeforeChange) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataValueBeforeChange) GetValuesBeforeChange() BACnetApplicationTagUnsignedInteger {
	return m.ValuesBeforeChange
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataValueBeforeChange) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetValuesBeforeChange())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataValueBeforeChange(structType any) BACnetConstructedDataValueBeforeChange {
	if casted, ok := structType.(BACnetConstructedDataValueBeforeChange); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataValueBeforeChange); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataValueBeforeChange) GetTypeName() string {
	return "BACnetConstructedDataValueBeforeChange"
}

func (m *_BACnetConstructedDataValueBeforeChange) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (valuesBeforeChange)
	lengthInBits += m.ValuesBeforeChange.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataValueBeforeChange) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataValueBeforeChange) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataValueBeforeChange BACnetConstructedDataValueBeforeChange, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataValueBeforeChange"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataValueBeforeChange")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	valuesBeforeChange, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "valuesBeforeChange", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'valuesBeforeChange' field"))
	}
	m.ValuesBeforeChange = valuesBeforeChange

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), valuesBeforeChange)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataValueBeforeChange"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataValueBeforeChange")
	}

	return m, nil
}

func (m *_BACnetConstructedDataValueBeforeChange) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataValueBeforeChange) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataValueBeforeChange"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataValueBeforeChange")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "valuesBeforeChange", m.GetValuesBeforeChange(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'valuesBeforeChange' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataValueBeforeChange"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataValueBeforeChange")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataValueBeforeChange) IsBACnetConstructedDataValueBeforeChange() {}

func (m *_BACnetConstructedDataValueBeforeChange) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataValueBeforeChange) deepCopy() *_BACnetConstructedDataValueBeforeChange {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataValueBeforeChangeCopy := &_BACnetConstructedDataValueBeforeChange{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagUnsignedInteger](m.ValuesBeforeChange),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataValueBeforeChangeCopy
}

func (m *_BACnetConstructedDataValueBeforeChange) String() string {
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
