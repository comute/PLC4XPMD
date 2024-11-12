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

// BACnetConstructedDataBinaryOutputRelinquishDefault is the corresponding interface of BACnetConstructedDataBinaryOutputRelinquishDefault
type BACnetConstructedDataBinaryOutputRelinquishDefault interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetRelinquishDefault returns RelinquishDefault (property field)
	GetRelinquishDefault() BACnetBinaryPVTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetBinaryPVTagged
	// IsBACnetConstructedDataBinaryOutputRelinquishDefault is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataBinaryOutputRelinquishDefault()
	// CreateBuilder creates a BACnetConstructedDataBinaryOutputRelinquishDefaultBuilder
	CreateBACnetConstructedDataBinaryOutputRelinquishDefaultBuilder() BACnetConstructedDataBinaryOutputRelinquishDefaultBuilder
}

// _BACnetConstructedDataBinaryOutputRelinquishDefault is the data-structure of this message
type _BACnetConstructedDataBinaryOutputRelinquishDefault struct {
	BACnetConstructedDataContract
	RelinquishDefault BACnetBinaryPVTagged
}

var _ BACnetConstructedDataBinaryOutputRelinquishDefault = (*_BACnetConstructedDataBinaryOutputRelinquishDefault)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataBinaryOutputRelinquishDefault)(nil)

// NewBACnetConstructedDataBinaryOutputRelinquishDefault factory function for _BACnetConstructedDataBinaryOutputRelinquishDefault
func NewBACnetConstructedDataBinaryOutputRelinquishDefault(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, relinquishDefault BACnetBinaryPVTagged, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataBinaryOutputRelinquishDefault {
	if relinquishDefault == nil {
		panic("relinquishDefault of type BACnetBinaryPVTagged for BACnetConstructedDataBinaryOutputRelinquishDefault must not be nil")
	}
	_result := &_BACnetConstructedDataBinaryOutputRelinquishDefault{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		RelinquishDefault:             relinquishDefault,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataBinaryOutputRelinquishDefaultBuilder is a builder for BACnetConstructedDataBinaryOutputRelinquishDefault
type BACnetConstructedDataBinaryOutputRelinquishDefaultBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(relinquishDefault BACnetBinaryPVTagged) BACnetConstructedDataBinaryOutputRelinquishDefaultBuilder
	// WithRelinquishDefault adds RelinquishDefault (property field)
	WithRelinquishDefault(BACnetBinaryPVTagged) BACnetConstructedDataBinaryOutputRelinquishDefaultBuilder
	// WithRelinquishDefaultBuilder adds RelinquishDefault (property field) which is build by the builder
	WithRelinquishDefaultBuilder(func(BACnetBinaryPVTaggedBuilder) BACnetBinaryPVTaggedBuilder) BACnetConstructedDataBinaryOutputRelinquishDefaultBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataBinaryOutputRelinquishDefault or returns an error if something is wrong
	Build() (BACnetConstructedDataBinaryOutputRelinquishDefault, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataBinaryOutputRelinquishDefault
}

// NewBACnetConstructedDataBinaryOutputRelinquishDefaultBuilder() creates a BACnetConstructedDataBinaryOutputRelinquishDefaultBuilder
func NewBACnetConstructedDataBinaryOutputRelinquishDefaultBuilder() BACnetConstructedDataBinaryOutputRelinquishDefaultBuilder {
	return &_BACnetConstructedDataBinaryOutputRelinquishDefaultBuilder{_BACnetConstructedDataBinaryOutputRelinquishDefault: new(_BACnetConstructedDataBinaryOutputRelinquishDefault)}
}

type _BACnetConstructedDataBinaryOutputRelinquishDefaultBuilder struct {
	*_BACnetConstructedDataBinaryOutputRelinquishDefault

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataBinaryOutputRelinquishDefaultBuilder) = (*_BACnetConstructedDataBinaryOutputRelinquishDefaultBuilder)(nil)

func (b *_BACnetConstructedDataBinaryOutputRelinquishDefaultBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataBinaryOutputRelinquishDefault
}

func (b *_BACnetConstructedDataBinaryOutputRelinquishDefaultBuilder) WithMandatoryFields(relinquishDefault BACnetBinaryPVTagged) BACnetConstructedDataBinaryOutputRelinquishDefaultBuilder {
	return b.WithRelinquishDefault(relinquishDefault)
}

func (b *_BACnetConstructedDataBinaryOutputRelinquishDefaultBuilder) WithRelinquishDefault(relinquishDefault BACnetBinaryPVTagged) BACnetConstructedDataBinaryOutputRelinquishDefaultBuilder {
	b.RelinquishDefault = relinquishDefault
	return b
}

func (b *_BACnetConstructedDataBinaryOutputRelinquishDefaultBuilder) WithRelinquishDefaultBuilder(builderSupplier func(BACnetBinaryPVTaggedBuilder) BACnetBinaryPVTaggedBuilder) BACnetConstructedDataBinaryOutputRelinquishDefaultBuilder {
	builder := builderSupplier(b.RelinquishDefault.CreateBACnetBinaryPVTaggedBuilder())
	var err error
	b.RelinquishDefault, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetBinaryPVTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataBinaryOutputRelinquishDefaultBuilder) Build() (BACnetConstructedDataBinaryOutputRelinquishDefault, error) {
	if b.RelinquishDefault == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'relinquishDefault' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataBinaryOutputRelinquishDefault.deepCopy(), nil
}

func (b *_BACnetConstructedDataBinaryOutputRelinquishDefaultBuilder) MustBuild() BACnetConstructedDataBinaryOutputRelinquishDefault {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataBinaryOutputRelinquishDefaultBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataBinaryOutputRelinquishDefaultBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataBinaryOutputRelinquishDefaultBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataBinaryOutputRelinquishDefaultBuilder().(*_BACnetConstructedDataBinaryOutputRelinquishDefaultBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataBinaryOutputRelinquishDefaultBuilder creates a BACnetConstructedDataBinaryOutputRelinquishDefaultBuilder
func (b *_BACnetConstructedDataBinaryOutputRelinquishDefault) CreateBACnetConstructedDataBinaryOutputRelinquishDefaultBuilder() BACnetConstructedDataBinaryOutputRelinquishDefaultBuilder {
	if b == nil {
		return NewBACnetConstructedDataBinaryOutputRelinquishDefaultBuilder()
	}
	return &_BACnetConstructedDataBinaryOutputRelinquishDefaultBuilder{_BACnetConstructedDataBinaryOutputRelinquishDefault: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataBinaryOutputRelinquishDefault) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_BINARY_OUTPUT
}

func (m *_BACnetConstructedDataBinaryOutputRelinquishDefault) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_RELINQUISH_DEFAULT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataBinaryOutputRelinquishDefault) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataBinaryOutputRelinquishDefault) GetRelinquishDefault() BACnetBinaryPVTagged {
	return m.RelinquishDefault
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataBinaryOutputRelinquishDefault) GetActualValue() BACnetBinaryPVTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetBinaryPVTagged(m.GetRelinquishDefault())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataBinaryOutputRelinquishDefault(structType any) BACnetConstructedDataBinaryOutputRelinquishDefault {
	if casted, ok := structType.(BACnetConstructedDataBinaryOutputRelinquishDefault); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataBinaryOutputRelinquishDefault); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataBinaryOutputRelinquishDefault) GetTypeName() string {
	return "BACnetConstructedDataBinaryOutputRelinquishDefault"
}

func (m *_BACnetConstructedDataBinaryOutputRelinquishDefault) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (relinquishDefault)
	lengthInBits += m.RelinquishDefault.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataBinaryOutputRelinquishDefault) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataBinaryOutputRelinquishDefault) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataBinaryOutputRelinquishDefault BACnetConstructedDataBinaryOutputRelinquishDefault, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataBinaryOutputRelinquishDefault"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataBinaryOutputRelinquishDefault")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	relinquishDefault, err := ReadSimpleField[BACnetBinaryPVTagged](ctx, "relinquishDefault", ReadComplex[BACnetBinaryPVTagged](BACnetBinaryPVTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'relinquishDefault' field"))
	}
	m.RelinquishDefault = relinquishDefault

	actualValue, err := ReadVirtualField[BACnetBinaryPVTagged](ctx, "actualValue", (*BACnetBinaryPVTagged)(nil), relinquishDefault)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataBinaryOutputRelinquishDefault"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataBinaryOutputRelinquishDefault")
	}

	return m, nil
}

func (m *_BACnetConstructedDataBinaryOutputRelinquishDefault) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataBinaryOutputRelinquishDefault) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataBinaryOutputRelinquishDefault"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataBinaryOutputRelinquishDefault")
		}

		if err := WriteSimpleField[BACnetBinaryPVTagged](ctx, "relinquishDefault", m.GetRelinquishDefault(), WriteComplex[BACnetBinaryPVTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'relinquishDefault' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataBinaryOutputRelinquishDefault"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataBinaryOutputRelinquishDefault")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataBinaryOutputRelinquishDefault) IsBACnetConstructedDataBinaryOutputRelinquishDefault() {
}

func (m *_BACnetConstructedDataBinaryOutputRelinquishDefault) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataBinaryOutputRelinquishDefault) deepCopy() *_BACnetConstructedDataBinaryOutputRelinquishDefault {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataBinaryOutputRelinquishDefaultCopy := &_BACnetConstructedDataBinaryOutputRelinquishDefault{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetBinaryPVTagged](m.RelinquishDefault),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataBinaryOutputRelinquishDefaultCopy
}

func (m *_BACnetConstructedDataBinaryOutputRelinquishDefault) String() string {
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
