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

// BACnetConstructedDataAccessEventAuthenticationFactor is the corresponding interface of BACnetConstructedDataAccessEventAuthenticationFactor
type BACnetConstructedDataAccessEventAuthenticationFactor interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetAccessEventAuthenticationFactor returns AccessEventAuthenticationFactor (property field)
	GetAccessEventAuthenticationFactor() BACnetAuthenticationFactor
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetAuthenticationFactor
	// IsBACnetConstructedDataAccessEventAuthenticationFactor is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataAccessEventAuthenticationFactor()
	// CreateBuilder creates a BACnetConstructedDataAccessEventAuthenticationFactorBuilder
	CreateBACnetConstructedDataAccessEventAuthenticationFactorBuilder() BACnetConstructedDataAccessEventAuthenticationFactorBuilder
}

// _BACnetConstructedDataAccessEventAuthenticationFactor is the data-structure of this message
type _BACnetConstructedDataAccessEventAuthenticationFactor struct {
	BACnetConstructedDataContract
	AccessEventAuthenticationFactor BACnetAuthenticationFactor
}

var _ BACnetConstructedDataAccessEventAuthenticationFactor = (*_BACnetConstructedDataAccessEventAuthenticationFactor)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataAccessEventAuthenticationFactor)(nil)

// NewBACnetConstructedDataAccessEventAuthenticationFactor factory function for _BACnetConstructedDataAccessEventAuthenticationFactor
func NewBACnetConstructedDataAccessEventAuthenticationFactor(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, accessEventAuthenticationFactor BACnetAuthenticationFactor, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAccessEventAuthenticationFactor {
	if accessEventAuthenticationFactor == nil {
		panic("accessEventAuthenticationFactor of type BACnetAuthenticationFactor for BACnetConstructedDataAccessEventAuthenticationFactor must not be nil")
	}
	_result := &_BACnetConstructedDataAccessEventAuthenticationFactor{
		BACnetConstructedDataContract:   NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		AccessEventAuthenticationFactor: accessEventAuthenticationFactor,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataAccessEventAuthenticationFactorBuilder is a builder for BACnetConstructedDataAccessEventAuthenticationFactor
type BACnetConstructedDataAccessEventAuthenticationFactorBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(accessEventAuthenticationFactor BACnetAuthenticationFactor) BACnetConstructedDataAccessEventAuthenticationFactorBuilder
	// WithAccessEventAuthenticationFactor adds AccessEventAuthenticationFactor (property field)
	WithAccessEventAuthenticationFactor(BACnetAuthenticationFactor) BACnetConstructedDataAccessEventAuthenticationFactorBuilder
	// WithAccessEventAuthenticationFactorBuilder adds AccessEventAuthenticationFactor (property field) which is build by the builder
	WithAccessEventAuthenticationFactorBuilder(func(BACnetAuthenticationFactorBuilder) BACnetAuthenticationFactorBuilder) BACnetConstructedDataAccessEventAuthenticationFactorBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataAccessEventAuthenticationFactor or returns an error if something is wrong
	Build() (BACnetConstructedDataAccessEventAuthenticationFactor, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataAccessEventAuthenticationFactor
}

// NewBACnetConstructedDataAccessEventAuthenticationFactorBuilder() creates a BACnetConstructedDataAccessEventAuthenticationFactorBuilder
func NewBACnetConstructedDataAccessEventAuthenticationFactorBuilder() BACnetConstructedDataAccessEventAuthenticationFactorBuilder {
	return &_BACnetConstructedDataAccessEventAuthenticationFactorBuilder{_BACnetConstructedDataAccessEventAuthenticationFactor: new(_BACnetConstructedDataAccessEventAuthenticationFactor)}
}

type _BACnetConstructedDataAccessEventAuthenticationFactorBuilder struct {
	*_BACnetConstructedDataAccessEventAuthenticationFactor

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataAccessEventAuthenticationFactorBuilder) = (*_BACnetConstructedDataAccessEventAuthenticationFactorBuilder)(nil)

func (b *_BACnetConstructedDataAccessEventAuthenticationFactorBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataAccessEventAuthenticationFactor
}

func (b *_BACnetConstructedDataAccessEventAuthenticationFactorBuilder) WithMandatoryFields(accessEventAuthenticationFactor BACnetAuthenticationFactor) BACnetConstructedDataAccessEventAuthenticationFactorBuilder {
	return b.WithAccessEventAuthenticationFactor(accessEventAuthenticationFactor)
}

func (b *_BACnetConstructedDataAccessEventAuthenticationFactorBuilder) WithAccessEventAuthenticationFactor(accessEventAuthenticationFactor BACnetAuthenticationFactor) BACnetConstructedDataAccessEventAuthenticationFactorBuilder {
	b.AccessEventAuthenticationFactor = accessEventAuthenticationFactor
	return b
}

func (b *_BACnetConstructedDataAccessEventAuthenticationFactorBuilder) WithAccessEventAuthenticationFactorBuilder(builderSupplier func(BACnetAuthenticationFactorBuilder) BACnetAuthenticationFactorBuilder) BACnetConstructedDataAccessEventAuthenticationFactorBuilder {
	builder := builderSupplier(b.AccessEventAuthenticationFactor.CreateBACnetAuthenticationFactorBuilder())
	var err error
	b.AccessEventAuthenticationFactor, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetAuthenticationFactorBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataAccessEventAuthenticationFactorBuilder) Build() (BACnetConstructedDataAccessEventAuthenticationFactor, error) {
	if b.AccessEventAuthenticationFactor == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'accessEventAuthenticationFactor' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataAccessEventAuthenticationFactor.deepCopy(), nil
}

func (b *_BACnetConstructedDataAccessEventAuthenticationFactorBuilder) MustBuild() BACnetConstructedDataAccessEventAuthenticationFactor {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataAccessEventAuthenticationFactorBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataAccessEventAuthenticationFactorBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataAccessEventAuthenticationFactorBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataAccessEventAuthenticationFactorBuilder().(*_BACnetConstructedDataAccessEventAuthenticationFactorBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataAccessEventAuthenticationFactorBuilder creates a BACnetConstructedDataAccessEventAuthenticationFactorBuilder
func (b *_BACnetConstructedDataAccessEventAuthenticationFactor) CreateBACnetConstructedDataAccessEventAuthenticationFactorBuilder() BACnetConstructedDataAccessEventAuthenticationFactorBuilder {
	if b == nil {
		return NewBACnetConstructedDataAccessEventAuthenticationFactorBuilder()
	}
	return &_BACnetConstructedDataAccessEventAuthenticationFactorBuilder{_BACnetConstructedDataAccessEventAuthenticationFactor: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAccessEventAuthenticationFactor) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataAccessEventAuthenticationFactor) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ACCESS_EVENT_AUTHENTICATION_FACTOR
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAccessEventAuthenticationFactor) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataAccessEventAuthenticationFactor) GetAccessEventAuthenticationFactor() BACnetAuthenticationFactor {
	return m.AccessEventAuthenticationFactor
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataAccessEventAuthenticationFactor) GetActualValue() BACnetAuthenticationFactor {
	ctx := context.Background()
	_ = ctx
	return CastBACnetAuthenticationFactor(m.GetAccessEventAuthenticationFactor())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAccessEventAuthenticationFactor(structType any) BACnetConstructedDataAccessEventAuthenticationFactor {
	if casted, ok := structType.(BACnetConstructedDataAccessEventAuthenticationFactor); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAccessEventAuthenticationFactor); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAccessEventAuthenticationFactor) GetTypeName() string {
	return "BACnetConstructedDataAccessEventAuthenticationFactor"
}

func (m *_BACnetConstructedDataAccessEventAuthenticationFactor) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (accessEventAuthenticationFactor)
	lengthInBits += m.AccessEventAuthenticationFactor.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataAccessEventAuthenticationFactor) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataAccessEventAuthenticationFactor) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataAccessEventAuthenticationFactor BACnetConstructedDataAccessEventAuthenticationFactor, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAccessEventAuthenticationFactor"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAccessEventAuthenticationFactor")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	accessEventAuthenticationFactor, err := ReadSimpleField[BACnetAuthenticationFactor](ctx, "accessEventAuthenticationFactor", ReadComplex[BACnetAuthenticationFactor](BACnetAuthenticationFactorParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'accessEventAuthenticationFactor' field"))
	}
	m.AccessEventAuthenticationFactor = accessEventAuthenticationFactor

	actualValue, err := ReadVirtualField[BACnetAuthenticationFactor](ctx, "actualValue", (*BACnetAuthenticationFactor)(nil), accessEventAuthenticationFactor)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAccessEventAuthenticationFactor"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAccessEventAuthenticationFactor")
	}

	return m, nil
}

func (m *_BACnetConstructedDataAccessEventAuthenticationFactor) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataAccessEventAuthenticationFactor) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAccessEventAuthenticationFactor"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAccessEventAuthenticationFactor")
		}

		if err := WriteSimpleField[BACnetAuthenticationFactor](ctx, "accessEventAuthenticationFactor", m.GetAccessEventAuthenticationFactor(), WriteComplex[BACnetAuthenticationFactor](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'accessEventAuthenticationFactor' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAccessEventAuthenticationFactor"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAccessEventAuthenticationFactor")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAccessEventAuthenticationFactor) IsBACnetConstructedDataAccessEventAuthenticationFactor() {
}

func (m *_BACnetConstructedDataAccessEventAuthenticationFactor) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataAccessEventAuthenticationFactor) deepCopy() *_BACnetConstructedDataAccessEventAuthenticationFactor {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataAccessEventAuthenticationFactorCopy := &_BACnetConstructedDataAccessEventAuthenticationFactor{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetAuthenticationFactor](m.AccessEventAuthenticationFactor),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataAccessEventAuthenticationFactorCopy
}

func (m *_BACnetConstructedDataAccessEventAuthenticationFactor) String() string {
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
