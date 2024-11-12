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

// BACnetConstructedDataManipulatedVariableReference is the corresponding interface of BACnetConstructedDataManipulatedVariableReference
type BACnetConstructedDataManipulatedVariableReference interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetManipulatedVariableReference returns ManipulatedVariableReference (property field)
	GetManipulatedVariableReference() BACnetObjectPropertyReference
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetObjectPropertyReference
	// IsBACnetConstructedDataManipulatedVariableReference is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataManipulatedVariableReference()
	// CreateBuilder creates a BACnetConstructedDataManipulatedVariableReferenceBuilder
	CreateBACnetConstructedDataManipulatedVariableReferenceBuilder() BACnetConstructedDataManipulatedVariableReferenceBuilder
}

// _BACnetConstructedDataManipulatedVariableReference is the data-structure of this message
type _BACnetConstructedDataManipulatedVariableReference struct {
	BACnetConstructedDataContract
	ManipulatedVariableReference BACnetObjectPropertyReference
}

var _ BACnetConstructedDataManipulatedVariableReference = (*_BACnetConstructedDataManipulatedVariableReference)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataManipulatedVariableReference)(nil)

// NewBACnetConstructedDataManipulatedVariableReference factory function for _BACnetConstructedDataManipulatedVariableReference
func NewBACnetConstructedDataManipulatedVariableReference(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, manipulatedVariableReference BACnetObjectPropertyReference, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataManipulatedVariableReference {
	if manipulatedVariableReference == nil {
		panic("manipulatedVariableReference of type BACnetObjectPropertyReference for BACnetConstructedDataManipulatedVariableReference must not be nil")
	}
	_result := &_BACnetConstructedDataManipulatedVariableReference{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		ManipulatedVariableReference:  manipulatedVariableReference,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataManipulatedVariableReferenceBuilder is a builder for BACnetConstructedDataManipulatedVariableReference
type BACnetConstructedDataManipulatedVariableReferenceBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(manipulatedVariableReference BACnetObjectPropertyReference) BACnetConstructedDataManipulatedVariableReferenceBuilder
	// WithManipulatedVariableReference adds ManipulatedVariableReference (property field)
	WithManipulatedVariableReference(BACnetObjectPropertyReference) BACnetConstructedDataManipulatedVariableReferenceBuilder
	// WithManipulatedVariableReferenceBuilder adds ManipulatedVariableReference (property field) which is build by the builder
	WithManipulatedVariableReferenceBuilder(func(BACnetObjectPropertyReferenceBuilder) BACnetObjectPropertyReferenceBuilder) BACnetConstructedDataManipulatedVariableReferenceBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataManipulatedVariableReference or returns an error if something is wrong
	Build() (BACnetConstructedDataManipulatedVariableReference, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataManipulatedVariableReference
}

// NewBACnetConstructedDataManipulatedVariableReferenceBuilder() creates a BACnetConstructedDataManipulatedVariableReferenceBuilder
func NewBACnetConstructedDataManipulatedVariableReferenceBuilder() BACnetConstructedDataManipulatedVariableReferenceBuilder {
	return &_BACnetConstructedDataManipulatedVariableReferenceBuilder{_BACnetConstructedDataManipulatedVariableReference: new(_BACnetConstructedDataManipulatedVariableReference)}
}

type _BACnetConstructedDataManipulatedVariableReferenceBuilder struct {
	*_BACnetConstructedDataManipulatedVariableReference

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataManipulatedVariableReferenceBuilder) = (*_BACnetConstructedDataManipulatedVariableReferenceBuilder)(nil)

func (b *_BACnetConstructedDataManipulatedVariableReferenceBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataManipulatedVariableReference
}

func (b *_BACnetConstructedDataManipulatedVariableReferenceBuilder) WithMandatoryFields(manipulatedVariableReference BACnetObjectPropertyReference) BACnetConstructedDataManipulatedVariableReferenceBuilder {
	return b.WithManipulatedVariableReference(manipulatedVariableReference)
}

func (b *_BACnetConstructedDataManipulatedVariableReferenceBuilder) WithManipulatedVariableReference(manipulatedVariableReference BACnetObjectPropertyReference) BACnetConstructedDataManipulatedVariableReferenceBuilder {
	b.ManipulatedVariableReference = manipulatedVariableReference
	return b
}

func (b *_BACnetConstructedDataManipulatedVariableReferenceBuilder) WithManipulatedVariableReferenceBuilder(builderSupplier func(BACnetObjectPropertyReferenceBuilder) BACnetObjectPropertyReferenceBuilder) BACnetConstructedDataManipulatedVariableReferenceBuilder {
	builder := builderSupplier(b.ManipulatedVariableReference.CreateBACnetObjectPropertyReferenceBuilder())
	var err error
	b.ManipulatedVariableReference, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetObjectPropertyReferenceBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataManipulatedVariableReferenceBuilder) Build() (BACnetConstructedDataManipulatedVariableReference, error) {
	if b.ManipulatedVariableReference == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'manipulatedVariableReference' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataManipulatedVariableReference.deepCopy(), nil
}

func (b *_BACnetConstructedDataManipulatedVariableReferenceBuilder) MustBuild() BACnetConstructedDataManipulatedVariableReference {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataManipulatedVariableReferenceBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataManipulatedVariableReferenceBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataManipulatedVariableReferenceBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataManipulatedVariableReferenceBuilder().(*_BACnetConstructedDataManipulatedVariableReferenceBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataManipulatedVariableReferenceBuilder creates a BACnetConstructedDataManipulatedVariableReferenceBuilder
func (b *_BACnetConstructedDataManipulatedVariableReference) CreateBACnetConstructedDataManipulatedVariableReferenceBuilder() BACnetConstructedDataManipulatedVariableReferenceBuilder {
	if b == nil {
		return NewBACnetConstructedDataManipulatedVariableReferenceBuilder()
	}
	return &_BACnetConstructedDataManipulatedVariableReferenceBuilder{_BACnetConstructedDataManipulatedVariableReference: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataManipulatedVariableReference) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataManipulatedVariableReference) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_MANIPULATED_VARIABLE_REFERENCE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataManipulatedVariableReference) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataManipulatedVariableReference) GetManipulatedVariableReference() BACnetObjectPropertyReference {
	return m.ManipulatedVariableReference
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataManipulatedVariableReference) GetActualValue() BACnetObjectPropertyReference {
	ctx := context.Background()
	_ = ctx
	return CastBACnetObjectPropertyReference(m.GetManipulatedVariableReference())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataManipulatedVariableReference(structType any) BACnetConstructedDataManipulatedVariableReference {
	if casted, ok := structType.(BACnetConstructedDataManipulatedVariableReference); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataManipulatedVariableReference); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataManipulatedVariableReference) GetTypeName() string {
	return "BACnetConstructedDataManipulatedVariableReference"
}

func (m *_BACnetConstructedDataManipulatedVariableReference) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (manipulatedVariableReference)
	lengthInBits += m.ManipulatedVariableReference.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataManipulatedVariableReference) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataManipulatedVariableReference) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataManipulatedVariableReference BACnetConstructedDataManipulatedVariableReference, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataManipulatedVariableReference"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataManipulatedVariableReference")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	manipulatedVariableReference, err := ReadSimpleField[BACnetObjectPropertyReference](ctx, "manipulatedVariableReference", ReadComplex[BACnetObjectPropertyReference](BACnetObjectPropertyReferenceParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'manipulatedVariableReference' field"))
	}
	m.ManipulatedVariableReference = manipulatedVariableReference

	actualValue, err := ReadVirtualField[BACnetObjectPropertyReference](ctx, "actualValue", (*BACnetObjectPropertyReference)(nil), manipulatedVariableReference)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataManipulatedVariableReference"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataManipulatedVariableReference")
	}

	return m, nil
}

func (m *_BACnetConstructedDataManipulatedVariableReference) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataManipulatedVariableReference) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataManipulatedVariableReference"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataManipulatedVariableReference")
		}

		if err := WriteSimpleField[BACnetObjectPropertyReference](ctx, "manipulatedVariableReference", m.GetManipulatedVariableReference(), WriteComplex[BACnetObjectPropertyReference](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'manipulatedVariableReference' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataManipulatedVariableReference"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataManipulatedVariableReference")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataManipulatedVariableReference) IsBACnetConstructedDataManipulatedVariableReference() {
}

func (m *_BACnetConstructedDataManipulatedVariableReference) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataManipulatedVariableReference) deepCopy() *_BACnetConstructedDataManipulatedVariableReference {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataManipulatedVariableReferenceCopy := &_BACnetConstructedDataManipulatedVariableReference{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetObjectPropertyReference](m.ManipulatedVariableReference),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataManipulatedVariableReferenceCopy
}

func (m *_BACnetConstructedDataManipulatedVariableReference) String() string {
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
