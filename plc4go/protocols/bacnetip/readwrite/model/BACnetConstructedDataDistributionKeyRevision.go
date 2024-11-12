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

// BACnetConstructedDataDistributionKeyRevision is the corresponding interface of BACnetConstructedDataDistributionKeyRevision
type BACnetConstructedDataDistributionKeyRevision interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetDistributionKeyRevision returns DistributionKeyRevision (property field)
	GetDistributionKeyRevision() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetConstructedDataDistributionKeyRevision is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataDistributionKeyRevision()
	// CreateBuilder creates a BACnetConstructedDataDistributionKeyRevisionBuilder
	CreateBACnetConstructedDataDistributionKeyRevisionBuilder() BACnetConstructedDataDistributionKeyRevisionBuilder
}

// _BACnetConstructedDataDistributionKeyRevision is the data-structure of this message
type _BACnetConstructedDataDistributionKeyRevision struct {
	BACnetConstructedDataContract
	DistributionKeyRevision BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataDistributionKeyRevision = (*_BACnetConstructedDataDistributionKeyRevision)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataDistributionKeyRevision)(nil)

// NewBACnetConstructedDataDistributionKeyRevision factory function for _BACnetConstructedDataDistributionKeyRevision
func NewBACnetConstructedDataDistributionKeyRevision(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, distributionKeyRevision BACnetApplicationTagUnsignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataDistributionKeyRevision {
	if distributionKeyRevision == nil {
		panic("distributionKeyRevision of type BACnetApplicationTagUnsignedInteger for BACnetConstructedDataDistributionKeyRevision must not be nil")
	}
	_result := &_BACnetConstructedDataDistributionKeyRevision{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		DistributionKeyRevision:       distributionKeyRevision,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataDistributionKeyRevisionBuilder is a builder for BACnetConstructedDataDistributionKeyRevision
type BACnetConstructedDataDistributionKeyRevisionBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(distributionKeyRevision BACnetApplicationTagUnsignedInteger) BACnetConstructedDataDistributionKeyRevisionBuilder
	// WithDistributionKeyRevision adds DistributionKeyRevision (property field)
	WithDistributionKeyRevision(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataDistributionKeyRevisionBuilder
	// WithDistributionKeyRevisionBuilder adds DistributionKeyRevision (property field) which is build by the builder
	WithDistributionKeyRevisionBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataDistributionKeyRevisionBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataDistributionKeyRevision or returns an error if something is wrong
	Build() (BACnetConstructedDataDistributionKeyRevision, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataDistributionKeyRevision
}

// NewBACnetConstructedDataDistributionKeyRevisionBuilder() creates a BACnetConstructedDataDistributionKeyRevisionBuilder
func NewBACnetConstructedDataDistributionKeyRevisionBuilder() BACnetConstructedDataDistributionKeyRevisionBuilder {
	return &_BACnetConstructedDataDistributionKeyRevisionBuilder{_BACnetConstructedDataDistributionKeyRevision: new(_BACnetConstructedDataDistributionKeyRevision)}
}

type _BACnetConstructedDataDistributionKeyRevisionBuilder struct {
	*_BACnetConstructedDataDistributionKeyRevision

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataDistributionKeyRevisionBuilder) = (*_BACnetConstructedDataDistributionKeyRevisionBuilder)(nil)

func (b *_BACnetConstructedDataDistributionKeyRevisionBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataDistributionKeyRevision
}

func (b *_BACnetConstructedDataDistributionKeyRevisionBuilder) WithMandatoryFields(distributionKeyRevision BACnetApplicationTagUnsignedInteger) BACnetConstructedDataDistributionKeyRevisionBuilder {
	return b.WithDistributionKeyRevision(distributionKeyRevision)
}

func (b *_BACnetConstructedDataDistributionKeyRevisionBuilder) WithDistributionKeyRevision(distributionKeyRevision BACnetApplicationTagUnsignedInteger) BACnetConstructedDataDistributionKeyRevisionBuilder {
	b.DistributionKeyRevision = distributionKeyRevision
	return b
}

func (b *_BACnetConstructedDataDistributionKeyRevisionBuilder) WithDistributionKeyRevisionBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataDistributionKeyRevisionBuilder {
	builder := builderSupplier(b.DistributionKeyRevision.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	b.DistributionKeyRevision, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataDistributionKeyRevisionBuilder) Build() (BACnetConstructedDataDistributionKeyRevision, error) {
	if b.DistributionKeyRevision == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'distributionKeyRevision' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataDistributionKeyRevision.deepCopy(), nil
}

func (b *_BACnetConstructedDataDistributionKeyRevisionBuilder) MustBuild() BACnetConstructedDataDistributionKeyRevision {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataDistributionKeyRevisionBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataDistributionKeyRevisionBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataDistributionKeyRevisionBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataDistributionKeyRevisionBuilder().(*_BACnetConstructedDataDistributionKeyRevisionBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataDistributionKeyRevisionBuilder creates a BACnetConstructedDataDistributionKeyRevisionBuilder
func (b *_BACnetConstructedDataDistributionKeyRevision) CreateBACnetConstructedDataDistributionKeyRevisionBuilder() BACnetConstructedDataDistributionKeyRevisionBuilder {
	if b == nil {
		return NewBACnetConstructedDataDistributionKeyRevisionBuilder()
	}
	return &_BACnetConstructedDataDistributionKeyRevisionBuilder{_BACnetConstructedDataDistributionKeyRevision: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataDistributionKeyRevision) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataDistributionKeyRevision) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_DISTRIBUTION_KEY_REVISION
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataDistributionKeyRevision) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataDistributionKeyRevision) GetDistributionKeyRevision() BACnetApplicationTagUnsignedInteger {
	return m.DistributionKeyRevision
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataDistributionKeyRevision) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetDistributionKeyRevision())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataDistributionKeyRevision(structType any) BACnetConstructedDataDistributionKeyRevision {
	if casted, ok := structType.(BACnetConstructedDataDistributionKeyRevision); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataDistributionKeyRevision); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataDistributionKeyRevision) GetTypeName() string {
	return "BACnetConstructedDataDistributionKeyRevision"
}

func (m *_BACnetConstructedDataDistributionKeyRevision) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (distributionKeyRevision)
	lengthInBits += m.DistributionKeyRevision.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataDistributionKeyRevision) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataDistributionKeyRevision) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataDistributionKeyRevision BACnetConstructedDataDistributionKeyRevision, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataDistributionKeyRevision"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataDistributionKeyRevision")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	distributionKeyRevision, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "distributionKeyRevision", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'distributionKeyRevision' field"))
	}
	m.DistributionKeyRevision = distributionKeyRevision

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), distributionKeyRevision)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataDistributionKeyRevision"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataDistributionKeyRevision")
	}

	return m, nil
}

func (m *_BACnetConstructedDataDistributionKeyRevision) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataDistributionKeyRevision) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataDistributionKeyRevision"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataDistributionKeyRevision")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "distributionKeyRevision", m.GetDistributionKeyRevision(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'distributionKeyRevision' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataDistributionKeyRevision"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataDistributionKeyRevision")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataDistributionKeyRevision) IsBACnetConstructedDataDistributionKeyRevision() {
}

func (m *_BACnetConstructedDataDistributionKeyRevision) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataDistributionKeyRevision) deepCopy() *_BACnetConstructedDataDistributionKeyRevision {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataDistributionKeyRevisionCopy := &_BACnetConstructedDataDistributionKeyRevision{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagUnsignedInteger](m.DistributionKeyRevision),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataDistributionKeyRevisionCopy
}

func (m *_BACnetConstructedDataDistributionKeyRevision) String() string {
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
