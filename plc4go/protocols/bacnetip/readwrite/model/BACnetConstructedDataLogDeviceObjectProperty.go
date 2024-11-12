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

// BACnetConstructedDataLogDeviceObjectProperty is the corresponding interface of BACnetConstructedDataLogDeviceObjectProperty
type BACnetConstructedDataLogDeviceObjectProperty interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetLogDeviceObjectProperty returns LogDeviceObjectProperty (property field)
	GetLogDeviceObjectProperty() BACnetDeviceObjectPropertyReference
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetDeviceObjectPropertyReference
	// IsBACnetConstructedDataLogDeviceObjectProperty is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataLogDeviceObjectProperty()
	// CreateBuilder creates a BACnetConstructedDataLogDeviceObjectPropertyBuilder
	CreateBACnetConstructedDataLogDeviceObjectPropertyBuilder() BACnetConstructedDataLogDeviceObjectPropertyBuilder
}

// _BACnetConstructedDataLogDeviceObjectProperty is the data-structure of this message
type _BACnetConstructedDataLogDeviceObjectProperty struct {
	BACnetConstructedDataContract
	LogDeviceObjectProperty BACnetDeviceObjectPropertyReference
}

var _ BACnetConstructedDataLogDeviceObjectProperty = (*_BACnetConstructedDataLogDeviceObjectProperty)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataLogDeviceObjectProperty)(nil)

// NewBACnetConstructedDataLogDeviceObjectProperty factory function for _BACnetConstructedDataLogDeviceObjectProperty
func NewBACnetConstructedDataLogDeviceObjectProperty(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, logDeviceObjectProperty BACnetDeviceObjectPropertyReference, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLogDeviceObjectProperty {
	if logDeviceObjectProperty == nil {
		panic("logDeviceObjectProperty of type BACnetDeviceObjectPropertyReference for BACnetConstructedDataLogDeviceObjectProperty must not be nil")
	}
	_result := &_BACnetConstructedDataLogDeviceObjectProperty{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		LogDeviceObjectProperty:       logDeviceObjectProperty,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataLogDeviceObjectPropertyBuilder is a builder for BACnetConstructedDataLogDeviceObjectProperty
type BACnetConstructedDataLogDeviceObjectPropertyBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(logDeviceObjectProperty BACnetDeviceObjectPropertyReference) BACnetConstructedDataLogDeviceObjectPropertyBuilder
	// WithLogDeviceObjectProperty adds LogDeviceObjectProperty (property field)
	WithLogDeviceObjectProperty(BACnetDeviceObjectPropertyReference) BACnetConstructedDataLogDeviceObjectPropertyBuilder
	// WithLogDeviceObjectPropertyBuilder adds LogDeviceObjectProperty (property field) which is build by the builder
	WithLogDeviceObjectPropertyBuilder(func(BACnetDeviceObjectPropertyReferenceBuilder) BACnetDeviceObjectPropertyReferenceBuilder) BACnetConstructedDataLogDeviceObjectPropertyBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataLogDeviceObjectProperty or returns an error if something is wrong
	Build() (BACnetConstructedDataLogDeviceObjectProperty, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataLogDeviceObjectProperty
}

// NewBACnetConstructedDataLogDeviceObjectPropertyBuilder() creates a BACnetConstructedDataLogDeviceObjectPropertyBuilder
func NewBACnetConstructedDataLogDeviceObjectPropertyBuilder() BACnetConstructedDataLogDeviceObjectPropertyBuilder {
	return &_BACnetConstructedDataLogDeviceObjectPropertyBuilder{_BACnetConstructedDataLogDeviceObjectProperty: new(_BACnetConstructedDataLogDeviceObjectProperty)}
}

type _BACnetConstructedDataLogDeviceObjectPropertyBuilder struct {
	*_BACnetConstructedDataLogDeviceObjectProperty

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataLogDeviceObjectPropertyBuilder) = (*_BACnetConstructedDataLogDeviceObjectPropertyBuilder)(nil)

func (b *_BACnetConstructedDataLogDeviceObjectPropertyBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataLogDeviceObjectProperty
}

func (b *_BACnetConstructedDataLogDeviceObjectPropertyBuilder) WithMandatoryFields(logDeviceObjectProperty BACnetDeviceObjectPropertyReference) BACnetConstructedDataLogDeviceObjectPropertyBuilder {
	return b.WithLogDeviceObjectProperty(logDeviceObjectProperty)
}

func (b *_BACnetConstructedDataLogDeviceObjectPropertyBuilder) WithLogDeviceObjectProperty(logDeviceObjectProperty BACnetDeviceObjectPropertyReference) BACnetConstructedDataLogDeviceObjectPropertyBuilder {
	b.LogDeviceObjectProperty = logDeviceObjectProperty
	return b
}

func (b *_BACnetConstructedDataLogDeviceObjectPropertyBuilder) WithLogDeviceObjectPropertyBuilder(builderSupplier func(BACnetDeviceObjectPropertyReferenceBuilder) BACnetDeviceObjectPropertyReferenceBuilder) BACnetConstructedDataLogDeviceObjectPropertyBuilder {
	builder := builderSupplier(b.LogDeviceObjectProperty.CreateBACnetDeviceObjectPropertyReferenceBuilder())
	var err error
	b.LogDeviceObjectProperty, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetDeviceObjectPropertyReferenceBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataLogDeviceObjectPropertyBuilder) Build() (BACnetConstructedDataLogDeviceObjectProperty, error) {
	if b.LogDeviceObjectProperty == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'logDeviceObjectProperty' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataLogDeviceObjectProperty.deepCopy(), nil
}

func (b *_BACnetConstructedDataLogDeviceObjectPropertyBuilder) MustBuild() BACnetConstructedDataLogDeviceObjectProperty {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataLogDeviceObjectPropertyBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataLogDeviceObjectPropertyBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataLogDeviceObjectPropertyBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataLogDeviceObjectPropertyBuilder().(*_BACnetConstructedDataLogDeviceObjectPropertyBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataLogDeviceObjectPropertyBuilder creates a BACnetConstructedDataLogDeviceObjectPropertyBuilder
func (b *_BACnetConstructedDataLogDeviceObjectProperty) CreateBACnetConstructedDataLogDeviceObjectPropertyBuilder() BACnetConstructedDataLogDeviceObjectPropertyBuilder {
	if b == nil {
		return NewBACnetConstructedDataLogDeviceObjectPropertyBuilder()
	}
	return &_BACnetConstructedDataLogDeviceObjectPropertyBuilder{_BACnetConstructedDataLogDeviceObjectProperty: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLogDeviceObjectProperty) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataLogDeviceObjectProperty) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LOG_DEVICE_OBJECT_PROPERTY
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLogDeviceObjectProperty) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLogDeviceObjectProperty) GetLogDeviceObjectProperty() BACnetDeviceObjectPropertyReference {
	return m.LogDeviceObjectProperty
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataLogDeviceObjectProperty) GetActualValue() BACnetDeviceObjectPropertyReference {
	ctx := context.Background()
	_ = ctx
	return CastBACnetDeviceObjectPropertyReference(m.GetLogDeviceObjectProperty())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLogDeviceObjectProperty(structType any) BACnetConstructedDataLogDeviceObjectProperty {
	if casted, ok := structType.(BACnetConstructedDataLogDeviceObjectProperty); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLogDeviceObjectProperty); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLogDeviceObjectProperty) GetTypeName() string {
	return "BACnetConstructedDataLogDeviceObjectProperty"
}

func (m *_BACnetConstructedDataLogDeviceObjectProperty) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (logDeviceObjectProperty)
	lengthInBits += m.LogDeviceObjectProperty.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataLogDeviceObjectProperty) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataLogDeviceObjectProperty) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataLogDeviceObjectProperty BACnetConstructedDataLogDeviceObjectProperty, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLogDeviceObjectProperty"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLogDeviceObjectProperty")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	logDeviceObjectProperty, err := ReadSimpleField[BACnetDeviceObjectPropertyReference](ctx, "logDeviceObjectProperty", ReadComplex[BACnetDeviceObjectPropertyReference](BACnetDeviceObjectPropertyReferenceParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'logDeviceObjectProperty' field"))
	}
	m.LogDeviceObjectProperty = logDeviceObjectProperty

	actualValue, err := ReadVirtualField[BACnetDeviceObjectPropertyReference](ctx, "actualValue", (*BACnetDeviceObjectPropertyReference)(nil), logDeviceObjectProperty)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLogDeviceObjectProperty"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLogDeviceObjectProperty")
	}

	return m, nil
}

func (m *_BACnetConstructedDataLogDeviceObjectProperty) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataLogDeviceObjectProperty) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLogDeviceObjectProperty"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLogDeviceObjectProperty")
		}

		if err := WriteSimpleField[BACnetDeviceObjectPropertyReference](ctx, "logDeviceObjectProperty", m.GetLogDeviceObjectProperty(), WriteComplex[BACnetDeviceObjectPropertyReference](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'logDeviceObjectProperty' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLogDeviceObjectProperty"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLogDeviceObjectProperty")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLogDeviceObjectProperty) IsBACnetConstructedDataLogDeviceObjectProperty() {
}

func (m *_BACnetConstructedDataLogDeviceObjectProperty) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataLogDeviceObjectProperty) deepCopy() *_BACnetConstructedDataLogDeviceObjectProperty {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataLogDeviceObjectPropertyCopy := &_BACnetConstructedDataLogDeviceObjectProperty{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetDeviceObjectPropertyReference](m.LogDeviceObjectProperty),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataLogDeviceObjectPropertyCopy
}

func (m *_BACnetConstructedDataLogDeviceObjectProperty) String() string {
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
