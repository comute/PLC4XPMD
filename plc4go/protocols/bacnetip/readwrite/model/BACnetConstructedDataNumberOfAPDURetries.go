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

// BACnetConstructedDataNumberOfAPDURetries is the corresponding interface of BACnetConstructedDataNumberOfAPDURetries
type BACnetConstructedDataNumberOfAPDURetries interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetNumberOfApduRetries returns NumberOfApduRetries (property field)
	GetNumberOfApduRetries() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetConstructedDataNumberOfAPDURetries is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataNumberOfAPDURetries()
	// CreateBuilder creates a BACnetConstructedDataNumberOfAPDURetriesBuilder
	CreateBACnetConstructedDataNumberOfAPDURetriesBuilder() BACnetConstructedDataNumberOfAPDURetriesBuilder
}

// _BACnetConstructedDataNumberOfAPDURetries is the data-structure of this message
type _BACnetConstructedDataNumberOfAPDURetries struct {
	BACnetConstructedDataContract
	NumberOfApduRetries BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataNumberOfAPDURetries = (*_BACnetConstructedDataNumberOfAPDURetries)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataNumberOfAPDURetries)(nil)

// NewBACnetConstructedDataNumberOfAPDURetries factory function for _BACnetConstructedDataNumberOfAPDURetries
func NewBACnetConstructedDataNumberOfAPDURetries(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, numberOfApduRetries BACnetApplicationTagUnsignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataNumberOfAPDURetries {
	if numberOfApduRetries == nil {
		panic("numberOfApduRetries of type BACnetApplicationTagUnsignedInteger for BACnetConstructedDataNumberOfAPDURetries must not be nil")
	}
	_result := &_BACnetConstructedDataNumberOfAPDURetries{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		NumberOfApduRetries:           numberOfApduRetries,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataNumberOfAPDURetriesBuilder is a builder for BACnetConstructedDataNumberOfAPDURetries
type BACnetConstructedDataNumberOfAPDURetriesBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(numberOfApduRetries BACnetApplicationTagUnsignedInteger) BACnetConstructedDataNumberOfAPDURetriesBuilder
	// WithNumberOfApduRetries adds NumberOfApduRetries (property field)
	WithNumberOfApduRetries(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataNumberOfAPDURetriesBuilder
	// WithNumberOfApduRetriesBuilder adds NumberOfApduRetries (property field) which is build by the builder
	WithNumberOfApduRetriesBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataNumberOfAPDURetriesBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataNumberOfAPDURetries or returns an error if something is wrong
	Build() (BACnetConstructedDataNumberOfAPDURetries, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataNumberOfAPDURetries
}

// NewBACnetConstructedDataNumberOfAPDURetriesBuilder() creates a BACnetConstructedDataNumberOfAPDURetriesBuilder
func NewBACnetConstructedDataNumberOfAPDURetriesBuilder() BACnetConstructedDataNumberOfAPDURetriesBuilder {
	return &_BACnetConstructedDataNumberOfAPDURetriesBuilder{_BACnetConstructedDataNumberOfAPDURetries: new(_BACnetConstructedDataNumberOfAPDURetries)}
}

type _BACnetConstructedDataNumberOfAPDURetriesBuilder struct {
	*_BACnetConstructedDataNumberOfAPDURetries

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataNumberOfAPDURetriesBuilder) = (*_BACnetConstructedDataNumberOfAPDURetriesBuilder)(nil)

func (b *_BACnetConstructedDataNumberOfAPDURetriesBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataNumberOfAPDURetries
}

func (b *_BACnetConstructedDataNumberOfAPDURetriesBuilder) WithMandatoryFields(numberOfApduRetries BACnetApplicationTagUnsignedInteger) BACnetConstructedDataNumberOfAPDURetriesBuilder {
	return b.WithNumberOfApduRetries(numberOfApduRetries)
}

func (b *_BACnetConstructedDataNumberOfAPDURetriesBuilder) WithNumberOfApduRetries(numberOfApduRetries BACnetApplicationTagUnsignedInteger) BACnetConstructedDataNumberOfAPDURetriesBuilder {
	b.NumberOfApduRetries = numberOfApduRetries
	return b
}

func (b *_BACnetConstructedDataNumberOfAPDURetriesBuilder) WithNumberOfApduRetriesBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataNumberOfAPDURetriesBuilder {
	builder := builderSupplier(b.NumberOfApduRetries.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	b.NumberOfApduRetries, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataNumberOfAPDURetriesBuilder) Build() (BACnetConstructedDataNumberOfAPDURetries, error) {
	if b.NumberOfApduRetries == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'numberOfApduRetries' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataNumberOfAPDURetries.deepCopy(), nil
}

func (b *_BACnetConstructedDataNumberOfAPDURetriesBuilder) MustBuild() BACnetConstructedDataNumberOfAPDURetries {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataNumberOfAPDURetriesBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataNumberOfAPDURetriesBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataNumberOfAPDURetriesBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataNumberOfAPDURetriesBuilder().(*_BACnetConstructedDataNumberOfAPDURetriesBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataNumberOfAPDURetriesBuilder creates a BACnetConstructedDataNumberOfAPDURetriesBuilder
func (b *_BACnetConstructedDataNumberOfAPDURetries) CreateBACnetConstructedDataNumberOfAPDURetriesBuilder() BACnetConstructedDataNumberOfAPDURetriesBuilder {
	if b == nil {
		return NewBACnetConstructedDataNumberOfAPDURetriesBuilder()
	}
	return &_BACnetConstructedDataNumberOfAPDURetriesBuilder{_BACnetConstructedDataNumberOfAPDURetries: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataNumberOfAPDURetries) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataNumberOfAPDURetries) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_NUMBER_OF_APDU_RETRIES
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataNumberOfAPDURetries) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataNumberOfAPDURetries) GetNumberOfApduRetries() BACnetApplicationTagUnsignedInteger {
	return m.NumberOfApduRetries
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataNumberOfAPDURetries) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetNumberOfApduRetries())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataNumberOfAPDURetries(structType any) BACnetConstructedDataNumberOfAPDURetries {
	if casted, ok := structType.(BACnetConstructedDataNumberOfAPDURetries); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataNumberOfAPDURetries); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataNumberOfAPDURetries) GetTypeName() string {
	return "BACnetConstructedDataNumberOfAPDURetries"
}

func (m *_BACnetConstructedDataNumberOfAPDURetries) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (numberOfApduRetries)
	lengthInBits += m.NumberOfApduRetries.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataNumberOfAPDURetries) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataNumberOfAPDURetries) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataNumberOfAPDURetries BACnetConstructedDataNumberOfAPDURetries, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataNumberOfAPDURetries"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataNumberOfAPDURetries")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	numberOfApduRetries, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "numberOfApduRetries", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'numberOfApduRetries' field"))
	}
	m.NumberOfApduRetries = numberOfApduRetries

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), numberOfApduRetries)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataNumberOfAPDURetries"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataNumberOfAPDURetries")
	}

	return m, nil
}

func (m *_BACnetConstructedDataNumberOfAPDURetries) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataNumberOfAPDURetries) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataNumberOfAPDURetries"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataNumberOfAPDURetries")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "numberOfApduRetries", m.GetNumberOfApduRetries(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'numberOfApduRetries' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataNumberOfAPDURetries"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataNumberOfAPDURetries")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataNumberOfAPDURetries) IsBACnetConstructedDataNumberOfAPDURetries() {}

func (m *_BACnetConstructedDataNumberOfAPDURetries) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataNumberOfAPDURetries) deepCopy() *_BACnetConstructedDataNumberOfAPDURetries {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataNumberOfAPDURetriesCopy := &_BACnetConstructedDataNumberOfAPDURetries{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagUnsignedInteger](m.NumberOfApduRetries),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataNumberOfAPDURetriesCopy
}

func (m *_BACnetConstructedDataNumberOfAPDURetries) String() string {
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
