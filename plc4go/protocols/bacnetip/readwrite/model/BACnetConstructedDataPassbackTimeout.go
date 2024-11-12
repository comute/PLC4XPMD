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

// BACnetConstructedDataPassbackTimeout is the corresponding interface of BACnetConstructedDataPassbackTimeout
type BACnetConstructedDataPassbackTimeout interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetPassbackTimeout returns PassbackTimeout (property field)
	GetPassbackTimeout() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetConstructedDataPassbackTimeout is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataPassbackTimeout()
	// CreateBuilder creates a BACnetConstructedDataPassbackTimeoutBuilder
	CreateBACnetConstructedDataPassbackTimeoutBuilder() BACnetConstructedDataPassbackTimeoutBuilder
}

// _BACnetConstructedDataPassbackTimeout is the data-structure of this message
type _BACnetConstructedDataPassbackTimeout struct {
	BACnetConstructedDataContract
	PassbackTimeout BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataPassbackTimeout = (*_BACnetConstructedDataPassbackTimeout)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataPassbackTimeout)(nil)

// NewBACnetConstructedDataPassbackTimeout factory function for _BACnetConstructedDataPassbackTimeout
func NewBACnetConstructedDataPassbackTimeout(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, passbackTimeout BACnetApplicationTagUnsignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataPassbackTimeout {
	if passbackTimeout == nil {
		panic("passbackTimeout of type BACnetApplicationTagUnsignedInteger for BACnetConstructedDataPassbackTimeout must not be nil")
	}
	_result := &_BACnetConstructedDataPassbackTimeout{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		PassbackTimeout:               passbackTimeout,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataPassbackTimeoutBuilder is a builder for BACnetConstructedDataPassbackTimeout
type BACnetConstructedDataPassbackTimeoutBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(passbackTimeout BACnetApplicationTagUnsignedInteger) BACnetConstructedDataPassbackTimeoutBuilder
	// WithPassbackTimeout adds PassbackTimeout (property field)
	WithPassbackTimeout(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataPassbackTimeoutBuilder
	// WithPassbackTimeoutBuilder adds PassbackTimeout (property field) which is build by the builder
	WithPassbackTimeoutBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataPassbackTimeoutBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataPassbackTimeout or returns an error if something is wrong
	Build() (BACnetConstructedDataPassbackTimeout, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataPassbackTimeout
}

// NewBACnetConstructedDataPassbackTimeoutBuilder() creates a BACnetConstructedDataPassbackTimeoutBuilder
func NewBACnetConstructedDataPassbackTimeoutBuilder() BACnetConstructedDataPassbackTimeoutBuilder {
	return &_BACnetConstructedDataPassbackTimeoutBuilder{_BACnetConstructedDataPassbackTimeout: new(_BACnetConstructedDataPassbackTimeout)}
}

type _BACnetConstructedDataPassbackTimeoutBuilder struct {
	*_BACnetConstructedDataPassbackTimeout

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataPassbackTimeoutBuilder) = (*_BACnetConstructedDataPassbackTimeoutBuilder)(nil)

func (b *_BACnetConstructedDataPassbackTimeoutBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataPassbackTimeout
}

func (b *_BACnetConstructedDataPassbackTimeoutBuilder) WithMandatoryFields(passbackTimeout BACnetApplicationTagUnsignedInteger) BACnetConstructedDataPassbackTimeoutBuilder {
	return b.WithPassbackTimeout(passbackTimeout)
}

func (b *_BACnetConstructedDataPassbackTimeoutBuilder) WithPassbackTimeout(passbackTimeout BACnetApplicationTagUnsignedInteger) BACnetConstructedDataPassbackTimeoutBuilder {
	b.PassbackTimeout = passbackTimeout
	return b
}

func (b *_BACnetConstructedDataPassbackTimeoutBuilder) WithPassbackTimeoutBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataPassbackTimeoutBuilder {
	builder := builderSupplier(b.PassbackTimeout.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	b.PassbackTimeout, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataPassbackTimeoutBuilder) Build() (BACnetConstructedDataPassbackTimeout, error) {
	if b.PassbackTimeout == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'passbackTimeout' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataPassbackTimeout.deepCopy(), nil
}

func (b *_BACnetConstructedDataPassbackTimeoutBuilder) MustBuild() BACnetConstructedDataPassbackTimeout {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataPassbackTimeoutBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataPassbackTimeoutBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataPassbackTimeoutBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataPassbackTimeoutBuilder().(*_BACnetConstructedDataPassbackTimeoutBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataPassbackTimeoutBuilder creates a BACnetConstructedDataPassbackTimeoutBuilder
func (b *_BACnetConstructedDataPassbackTimeout) CreateBACnetConstructedDataPassbackTimeoutBuilder() BACnetConstructedDataPassbackTimeoutBuilder {
	if b == nil {
		return NewBACnetConstructedDataPassbackTimeoutBuilder()
	}
	return &_BACnetConstructedDataPassbackTimeoutBuilder{_BACnetConstructedDataPassbackTimeout: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataPassbackTimeout) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataPassbackTimeout) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_PASSBACK_TIMEOUT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataPassbackTimeout) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataPassbackTimeout) GetPassbackTimeout() BACnetApplicationTagUnsignedInteger {
	return m.PassbackTimeout
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataPassbackTimeout) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetPassbackTimeout())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataPassbackTimeout(structType any) BACnetConstructedDataPassbackTimeout {
	if casted, ok := structType.(BACnetConstructedDataPassbackTimeout); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataPassbackTimeout); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataPassbackTimeout) GetTypeName() string {
	return "BACnetConstructedDataPassbackTimeout"
}

func (m *_BACnetConstructedDataPassbackTimeout) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (passbackTimeout)
	lengthInBits += m.PassbackTimeout.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataPassbackTimeout) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataPassbackTimeout) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataPassbackTimeout BACnetConstructedDataPassbackTimeout, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataPassbackTimeout"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataPassbackTimeout")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	passbackTimeout, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "passbackTimeout", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'passbackTimeout' field"))
	}
	m.PassbackTimeout = passbackTimeout

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), passbackTimeout)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataPassbackTimeout"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataPassbackTimeout")
	}

	return m, nil
}

func (m *_BACnetConstructedDataPassbackTimeout) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataPassbackTimeout) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataPassbackTimeout"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataPassbackTimeout")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "passbackTimeout", m.GetPassbackTimeout(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'passbackTimeout' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataPassbackTimeout"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataPassbackTimeout")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataPassbackTimeout) IsBACnetConstructedDataPassbackTimeout() {}

func (m *_BACnetConstructedDataPassbackTimeout) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataPassbackTimeout) deepCopy() *_BACnetConstructedDataPassbackTimeout {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataPassbackTimeoutCopy := &_BACnetConstructedDataPassbackTimeout{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagUnsignedInteger](m.PassbackTimeout),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataPassbackTimeoutCopy
}

func (m *_BACnetConstructedDataPassbackTimeout) String() string {
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
