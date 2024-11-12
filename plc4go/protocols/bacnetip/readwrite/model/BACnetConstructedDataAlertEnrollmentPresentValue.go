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

// BACnetConstructedDataAlertEnrollmentPresentValue is the corresponding interface of BACnetConstructedDataAlertEnrollmentPresentValue
type BACnetConstructedDataAlertEnrollmentPresentValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetPresentValue returns PresentValue (property field)
	GetPresentValue() BACnetApplicationTagObjectIdentifier
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagObjectIdentifier
	// IsBACnetConstructedDataAlertEnrollmentPresentValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataAlertEnrollmentPresentValue()
	// CreateBuilder creates a BACnetConstructedDataAlertEnrollmentPresentValueBuilder
	CreateBACnetConstructedDataAlertEnrollmentPresentValueBuilder() BACnetConstructedDataAlertEnrollmentPresentValueBuilder
}

// _BACnetConstructedDataAlertEnrollmentPresentValue is the data-structure of this message
type _BACnetConstructedDataAlertEnrollmentPresentValue struct {
	BACnetConstructedDataContract
	PresentValue BACnetApplicationTagObjectIdentifier
}

var _ BACnetConstructedDataAlertEnrollmentPresentValue = (*_BACnetConstructedDataAlertEnrollmentPresentValue)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataAlertEnrollmentPresentValue)(nil)

// NewBACnetConstructedDataAlertEnrollmentPresentValue factory function for _BACnetConstructedDataAlertEnrollmentPresentValue
func NewBACnetConstructedDataAlertEnrollmentPresentValue(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, presentValue BACnetApplicationTagObjectIdentifier, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAlertEnrollmentPresentValue {
	if presentValue == nil {
		panic("presentValue of type BACnetApplicationTagObjectIdentifier for BACnetConstructedDataAlertEnrollmentPresentValue must not be nil")
	}
	_result := &_BACnetConstructedDataAlertEnrollmentPresentValue{
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

// BACnetConstructedDataAlertEnrollmentPresentValueBuilder is a builder for BACnetConstructedDataAlertEnrollmentPresentValue
type BACnetConstructedDataAlertEnrollmentPresentValueBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(presentValue BACnetApplicationTagObjectIdentifier) BACnetConstructedDataAlertEnrollmentPresentValueBuilder
	// WithPresentValue adds PresentValue (property field)
	WithPresentValue(BACnetApplicationTagObjectIdentifier) BACnetConstructedDataAlertEnrollmentPresentValueBuilder
	// WithPresentValueBuilder adds PresentValue (property field) which is build by the builder
	WithPresentValueBuilder(func(BACnetApplicationTagObjectIdentifierBuilder) BACnetApplicationTagObjectIdentifierBuilder) BACnetConstructedDataAlertEnrollmentPresentValueBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataAlertEnrollmentPresentValue or returns an error if something is wrong
	Build() (BACnetConstructedDataAlertEnrollmentPresentValue, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataAlertEnrollmentPresentValue
}

// NewBACnetConstructedDataAlertEnrollmentPresentValueBuilder() creates a BACnetConstructedDataAlertEnrollmentPresentValueBuilder
func NewBACnetConstructedDataAlertEnrollmentPresentValueBuilder() BACnetConstructedDataAlertEnrollmentPresentValueBuilder {
	return &_BACnetConstructedDataAlertEnrollmentPresentValueBuilder{_BACnetConstructedDataAlertEnrollmentPresentValue: new(_BACnetConstructedDataAlertEnrollmentPresentValue)}
}

type _BACnetConstructedDataAlertEnrollmentPresentValueBuilder struct {
	*_BACnetConstructedDataAlertEnrollmentPresentValue

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataAlertEnrollmentPresentValueBuilder) = (*_BACnetConstructedDataAlertEnrollmentPresentValueBuilder)(nil)

func (b *_BACnetConstructedDataAlertEnrollmentPresentValueBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataAlertEnrollmentPresentValue
}

func (b *_BACnetConstructedDataAlertEnrollmentPresentValueBuilder) WithMandatoryFields(presentValue BACnetApplicationTagObjectIdentifier) BACnetConstructedDataAlertEnrollmentPresentValueBuilder {
	return b.WithPresentValue(presentValue)
}

func (b *_BACnetConstructedDataAlertEnrollmentPresentValueBuilder) WithPresentValue(presentValue BACnetApplicationTagObjectIdentifier) BACnetConstructedDataAlertEnrollmentPresentValueBuilder {
	b.PresentValue = presentValue
	return b
}

func (b *_BACnetConstructedDataAlertEnrollmentPresentValueBuilder) WithPresentValueBuilder(builderSupplier func(BACnetApplicationTagObjectIdentifierBuilder) BACnetApplicationTagObjectIdentifierBuilder) BACnetConstructedDataAlertEnrollmentPresentValueBuilder {
	builder := builderSupplier(b.PresentValue.CreateBACnetApplicationTagObjectIdentifierBuilder())
	var err error
	b.PresentValue, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagObjectIdentifierBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataAlertEnrollmentPresentValueBuilder) Build() (BACnetConstructedDataAlertEnrollmentPresentValue, error) {
	if b.PresentValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'presentValue' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataAlertEnrollmentPresentValue.deepCopy(), nil
}

func (b *_BACnetConstructedDataAlertEnrollmentPresentValueBuilder) MustBuild() BACnetConstructedDataAlertEnrollmentPresentValue {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataAlertEnrollmentPresentValueBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataAlertEnrollmentPresentValueBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataAlertEnrollmentPresentValueBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataAlertEnrollmentPresentValueBuilder().(*_BACnetConstructedDataAlertEnrollmentPresentValueBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataAlertEnrollmentPresentValueBuilder creates a BACnetConstructedDataAlertEnrollmentPresentValueBuilder
func (b *_BACnetConstructedDataAlertEnrollmentPresentValue) CreateBACnetConstructedDataAlertEnrollmentPresentValueBuilder() BACnetConstructedDataAlertEnrollmentPresentValueBuilder {
	if b == nil {
		return NewBACnetConstructedDataAlertEnrollmentPresentValueBuilder()
	}
	return &_BACnetConstructedDataAlertEnrollmentPresentValueBuilder{_BACnetConstructedDataAlertEnrollmentPresentValue: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAlertEnrollmentPresentValue) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_ALERT_ENROLLMENT
}

func (m *_BACnetConstructedDataAlertEnrollmentPresentValue) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_PRESENT_VALUE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAlertEnrollmentPresentValue) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataAlertEnrollmentPresentValue) GetPresentValue() BACnetApplicationTagObjectIdentifier {
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

func (m *_BACnetConstructedDataAlertEnrollmentPresentValue) GetActualValue() BACnetApplicationTagObjectIdentifier {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagObjectIdentifier(m.GetPresentValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAlertEnrollmentPresentValue(structType any) BACnetConstructedDataAlertEnrollmentPresentValue {
	if casted, ok := structType.(BACnetConstructedDataAlertEnrollmentPresentValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAlertEnrollmentPresentValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAlertEnrollmentPresentValue) GetTypeName() string {
	return "BACnetConstructedDataAlertEnrollmentPresentValue"
}

func (m *_BACnetConstructedDataAlertEnrollmentPresentValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (presentValue)
	lengthInBits += m.PresentValue.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataAlertEnrollmentPresentValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataAlertEnrollmentPresentValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataAlertEnrollmentPresentValue BACnetConstructedDataAlertEnrollmentPresentValue, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAlertEnrollmentPresentValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAlertEnrollmentPresentValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	presentValue, err := ReadSimpleField[BACnetApplicationTagObjectIdentifier](ctx, "presentValue", ReadComplex[BACnetApplicationTagObjectIdentifier](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagObjectIdentifier](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'presentValue' field"))
	}
	m.PresentValue = presentValue

	actualValue, err := ReadVirtualField[BACnetApplicationTagObjectIdentifier](ctx, "actualValue", (*BACnetApplicationTagObjectIdentifier)(nil), presentValue)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAlertEnrollmentPresentValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAlertEnrollmentPresentValue")
	}

	return m, nil
}

func (m *_BACnetConstructedDataAlertEnrollmentPresentValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataAlertEnrollmentPresentValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAlertEnrollmentPresentValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAlertEnrollmentPresentValue")
		}

		if err := WriteSimpleField[BACnetApplicationTagObjectIdentifier](ctx, "presentValue", m.GetPresentValue(), WriteComplex[BACnetApplicationTagObjectIdentifier](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'presentValue' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAlertEnrollmentPresentValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAlertEnrollmentPresentValue")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAlertEnrollmentPresentValue) IsBACnetConstructedDataAlertEnrollmentPresentValue() {
}

func (m *_BACnetConstructedDataAlertEnrollmentPresentValue) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataAlertEnrollmentPresentValue) deepCopy() *_BACnetConstructedDataAlertEnrollmentPresentValue {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataAlertEnrollmentPresentValueCopy := &_BACnetConstructedDataAlertEnrollmentPresentValue{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagObjectIdentifier](m.PresentValue),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataAlertEnrollmentPresentValueCopy
}

func (m *_BACnetConstructedDataAlertEnrollmentPresentValue) String() string {
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
