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

// BACnetConstructedDataIPv6PrefixLength is the corresponding interface of BACnetConstructedDataIPv6PrefixLength
type BACnetConstructedDataIPv6PrefixLength interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetIpv6PrefixLength returns Ipv6PrefixLength (property field)
	GetIpv6PrefixLength() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetConstructedDataIPv6PrefixLength is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataIPv6PrefixLength()
	// CreateBuilder creates a BACnetConstructedDataIPv6PrefixLengthBuilder
	CreateBACnetConstructedDataIPv6PrefixLengthBuilder() BACnetConstructedDataIPv6PrefixLengthBuilder
}

// _BACnetConstructedDataIPv6PrefixLength is the data-structure of this message
type _BACnetConstructedDataIPv6PrefixLength struct {
	BACnetConstructedDataContract
	Ipv6PrefixLength BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataIPv6PrefixLength = (*_BACnetConstructedDataIPv6PrefixLength)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataIPv6PrefixLength)(nil)

// NewBACnetConstructedDataIPv6PrefixLength factory function for _BACnetConstructedDataIPv6PrefixLength
func NewBACnetConstructedDataIPv6PrefixLength(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, ipv6PrefixLength BACnetApplicationTagUnsignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataIPv6PrefixLength {
	if ipv6PrefixLength == nil {
		panic("ipv6PrefixLength of type BACnetApplicationTagUnsignedInteger for BACnetConstructedDataIPv6PrefixLength must not be nil")
	}
	_result := &_BACnetConstructedDataIPv6PrefixLength{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		Ipv6PrefixLength:              ipv6PrefixLength,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataIPv6PrefixLengthBuilder is a builder for BACnetConstructedDataIPv6PrefixLength
type BACnetConstructedDataIPv6PrefixLengthBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(ipv6PrefixLength BACnetApplicationTagUnsignedInteger) BACnetConstructedDataIPv6PrefixLengthBuilder
	// WithIpv6PrefixLength adds Ipv6PrefixLength (property field)
	WithIpv6PrefixLength(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataIPv6PrefixLengthBuilder
	// WithIpv6PrefixLengthBuilder adds Ipv6PrefixLength (property field) which is build by the builder
	WithIpv6PrefixLengthBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataIPv6PrefixLengthBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataIPv6PrefixLength or returns an error if something is wrong
	Build() (BACnetConstructedDataIPv6PrefixLength, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataIPv6PrefixLength
}

// NewBACnetConstructedDataIPv6PrefixLengthBuilder() creates a BACnetConstructedDataIPv6PrefixLengthBuilder
func NewBACnetConstructedDataIPv6PrefixLengthBuilder() BACnetConstructedDataIPv6PrefixLengthBuilder {
	return &_BACnetConstructedDataIPv6PrefixLengthBuilder{_BACnetConstructedDataIPv6PrefixLength: new(_BACnetConstructedDataIPv6PrefixLength)}
}

type _BACnetConstructedDataIPv6PrefixLengthBuilder struct {
	*_BACnetConstructedDataIPv6PrefixLength

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataIPv6PrefixLengthBuilder) = (*_BACnetConstructedDataIPv6PrefixLengthBuilder)(nil)

func (b *_BACnetConstructedDataIPv6PrefixLengthBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataIPv6PrefixLength
}

func (b *_BACnetConstructedDataIPv6PrefixLengthBuilder) WithMandatoryFields(ipv6PrefixLength BACnetApplicationTagUnsignedInteger) BACnetConstructedDataIPv6PrefixLengthBuilder {
	return b.WithIpv6PrefixLength(ipv6PrefixLength)
}

func (b *_BACnetConstructedDataIPv6PrefixLengthBuilder) WithIpv6PrefixLength(ipv6PrefixLength BACnetApplicationTagUnsignedInteger) BACnetConstructedDataIPv6PrefixLengthBuilder {
	b.Ipv6PrefixLength = ipv6PrefixLength
	return b
}

func (b *_BACnetConstructedDataIPv6PrefixLengthBuilder) WithIpv6PrefixLengthBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataIPv6PrefixLengthBuilder {
	builder := builderSupplier(b.Ipv6PrefixLength.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	b.Ipv6PrefixLength, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataIPv6PrefixLengthBuilder) Build() (BACnetConstructedDataIPv6PrefixLength, error) {
	if b.Ipv6PrefixLength == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'ipv6PrefixLength' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataIPv6PrefixLength.deepCopy(), nil
}

func (b *_BACnetConstructedDataIPv6PrefixLengthBuilder) MustBuild() BACnetConstructedDataIPv6PrefixLength {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataIPv6PrefixLengthBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataIPv6PrefixLengthBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataIPv6PrefixLengthBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataIPv6PrefixLengthBuilder().(*_BACnetConstructedDataIPv6PrefixLengthBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataIPv6PrefixLengthBuilder creates a BACnetConstructedDataIPv6PrefixLengthBuilder
func (b *_BACnetConstructedDataIPv6PrefixLength) CreateBACnetConstructedDataIPv6PrefixLengthBuilder() BACnetConstructedDataIPv6PrefixLengthBuilder {
	if b == nil {
		return NewBACnetConstructedDataIPv6PrefixLengthBuilder()
	}
	return &_BACnetConstructedDataIPv6PrefixLengthBuilder{_BACnetConstructedDataIPv6PrefixLength: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataIPv6PrefixLength) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataIPv6PrefixLength) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_IPV6_PREFIX_LENGTH
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataIPv6PrefixLength) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataIPv6PrefixLength) GetIpv6PrefixLength() BACnetApplicationTagUnsignedInteger {
	return m.Ipv6PrefixLength
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataIPv6PrefixLength) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetIpv6PrefixLength())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataIPv6PrefixLength(structType any) BACnetConstructedDataIPv6PrefixLength {
	if casted, ok := structType.(BACnetConstructedDataIPv6PrefixLength); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataIPv6PrefixLength); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataIPv6PrefixLength) GetTypeName() string {
	return "BACnetConstructedDataIPv6PrefixLength"
}

func (m *_BACnetConstructedDataIPv6PrefixLength) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (ipv6PrefixLength)
	lengthInBits += m.Ipv6PrefixLength.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataIPv6PrefixLength) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataIPv6PrefixLength) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataIPv6PrefixLength BACnetConstructedDataIPv6PrefixLength, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataIPv6PrefixLength"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataIPv6PrefixLength")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	ipv6PrefixLength, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "ipv6PrefixLength", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'ipv6PrefixLength' field"))
	}
	m.Ipv6PrefixLength = ipv6PrefixLength

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), ipv6PrefixLength)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataIPv6PrefixLength"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataIPv6PrefixLength")
	}

	return m, nil
}

func (m *_BACnetConstructedDataIPv6PrefixLength) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataIPv6PrefixLength) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataIPv6PrefixLength"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataIPv6PrefixLength")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "ipv6PrefixLength", m.GetIpv6PrefixLength(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'ipv6PrefixLength' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataIPv6PrefixLength"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataIPv6PrefixLength")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataIPv6PrefixLength) IsBACnetConstructedDataIPv6PrefixLength() {}

func (m *_BACnetConstructedDataIPv6PrefixLength) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataIPv6PrefixLength) deepCopy() *_BACnetConstructedDataIPv6PrefixLength {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataIPv6PrefixLengthCopy := &_BACnetConstructedDataIPv6PrefixLength{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagUnsignedInteger](m.Ipv6PrefixLength),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataIPv6PrefixLengthCopy
}

func (m *_BACnetConstructedDataIPv6PrefixLength) String() string {
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
