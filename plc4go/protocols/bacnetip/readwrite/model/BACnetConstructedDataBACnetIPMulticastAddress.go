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

// BACnetConstructedDataBACnetIPMulticastAddress is the corresponding interface of BACnetConstructedDataBACnetIPMulticastAddress
type BACnetConstructedDataBACnetIPMulticastAddress interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetIpMulticastAddress returns IpMulticastAddress (property field)
	GetIpMulticastAddress() BACnetApplicationTagOctetString
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagOctetString
	// IsBACnetConstructedDataBACnetIPMulticastAddress is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataBACnetIPMulticastAddress()
	// CreateBuilder creates a BACnetConstructedDataBACnetIPMulticastAddressBuilder
	CreateBACnetConstructedDataBACnetIPMulticastAddressBuilder() BACnetConstructedDataBACnetIPMulticastAddressBuilder
}

// _BACnetConstructedDataBACnetIPMulticastAddress is the data-structure of this message
type _BACnetConstructedDataBACnetIPMulticastAddress struct {
	BACnetConstructedDataContract
	IpMulticastAddress BACnetApplicationTagOctetString
}

var _ BACnetConstructedDataBACnetIPMulticastAddress = (*_BACnetConstructedDataBACnetIPMulticastAddress)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataBACnetIPMulticastAddress)(nil)

// NewBACnetConstructedDataBACnetIPMulticastAddress factory function for _BACnetConstructedDataBACnetIPMulticastAddress
func NewBACnetConstructedDataBACnetIPMulticastAddress(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, ipMulticastAddress BACnetApplicationTagOctetString, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataBACnetIPMulticastAddress {
	if ipMulticastAddress == nil {
		panic("ipMulticastAddress of type BACnetApplicationTagOctetString for BACnetConstructedDataBACnetIPMulticastAddress must not be nil")
	}
	_result := &_BACnetConstructedDataBACnetIPMulticastAddress{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		IpMulticastAddress:            ipMulticastAddress,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataBACnetIPMulticastAddressBuilder is a builder for BACnetConstructedDataBACnetIPMulticastAddress
type BACnetConstructedDataBACnetIPMulticastAddressBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(ipMulticastAddress BACnetApplicationTagOctetString) BACnetConstructedDataBACnetIPMulticastAddressBuilder
	// WithIpMulticastAddress adds IpMulticastAddress (property field)
	WithIpMulticastAddress(BACnetApplicationTagOctetString) BACnetConstructedDataBACnetIPMulticastAddressBuilder
	// WithIpMulticastAddressBuilder adds IpMulticastAddress (property field) which is build by the builder
	WithIpMulticastAddressBuilder(func(BACnetApplicationTagOctetStringBuilder) BACnetApplicationTagOctetStringBuilder) BACnetConstructedDataBACnetIPMulticastAddressBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataBACnetIPMulticastAddress or returns an error if something is wrong
	Build() (BACnetConstructedDataBACnetIPMulticastAddress, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataBACnetIPMulticastAddress
}

// NewBACnetConstructedDataBACnetIPMulticastAddressBuilder() creates a BACnetConstructedDataBACnetIPMulticastAddressBuilder
func NewBACnetConstructedDataBACnetIPMulticastAddressBuilder() BACnetConstructedDataBACnetIPMulticastAddressBuilder {
	return &_BACnetConstructedDataBACnetIPMulticastAddressBuilder{_BACnetConstructedDataBACnetIPMulticastAddress: new(_BACnetConstructedDataBACnetIPMulticastAddress)}
}

type _BACnetConstructedDataBACnetIPMulticastAddressBuilder struct {
	*_BACnetConstructedDataBACnetIPMulticastAddress

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataBACnetIPMulticastAddressBuilder) = (*_BACnetConstructedDataBACnetIPMulticastAddressBuilder)(nil)

func (b *_BACnetConstructedDataBACnetIPMulticastAddressBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataBACnetIPMulticastAddress
}

func (b *_BACnetConstructedDataBACnetIPMulticastAddressBuilder) WithMandatoryFields(ipMulticastAddress BACnetApplicationTagOctetString) BACnetConstructedDataBACnetIPMulticastAddressBuilder {
	return b.WithIpMulticastAddress(ipMulticastAddress)
}

func (b *_BACnetConstructedDataBACnetIPMulticastAddressBuilder) WithIpMulticastAddress(ipMulticastAddress BACnetApplicationTagOctetString) BACnetConstructedDataBACnetIPMulticastAddressBuilder {
	b.IpMulticastAddress = ipMulticastAddress
	return b
}

func (b *_BACnetConstructedDataBACnetIPMulticastAddressBuilder) WithIpMulticastAddressBuilder(builderSupplier func(BACnetApplicationTagOctetStringBuilder) BACnetApplicationTagOctetStringBuilder) BACnetConstructedDataBACnetIPMulticastAddressBuilder {
	builder := builderSupplier(b.IpMulticastAddress.CreateBACnetApplicationTagOctetStringBuilder())
	var err error
	b.IpMulticastAddress, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagOctetStringBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataBACnetIPMulticastAddressBuilder) Build() (BACnetConstructedDataBACnetIPMulticastAddress, error) {
	if b.IpMulticastAddress == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'ipMulticastAddress' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataBACnetIPMulticastAddress.deepCopy(), nil
}

func (b *_BACnetConstructedDataBACnetIPMulticastAddressBuilder) MustBuild() BACnetConstructedDataBACnetIPMulticastAddress {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataBACnetIPMulticastAddressBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataBACnetIPMulticastAddressBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataBACnetIPMulticastAddressBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataBACnetIPMulticastAddressBuilder().(*_BACnetConstructedDataBACnetIPMulticastAddressBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataBACnetIPMulticastAddressBuilder creates a BACnetConstructedDataBACnetIPMulticastAddressBuilder
func (b *_BACnetConstructedDataBACnetIPMulticastAddress) CreateBACnetConstructedDataBACnetIPMulticastAddressBuilder() BACnetConstructedDataBACnetIPMulticastAddressBuilder {
	if b == nil {
		return NewBACnetConstructedDataBACnetIPMulticastAddressBuilder()
	}
	return &_BACnetConstructedDataBACnetIPMulticastAddressBuilder{_BACnetConstructedDataBACnetIPMulticastAddress: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataBACnetIPMulticastAddress) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataBACnetIPMulticastAddress) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_BACNET_IP_MULTICAST_ADDRESS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataBACnetIPMulticastAddress) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataBACnetIPMulticastAddress) GetIpMulticastAddress() BACnetApplicationTagOctetString {
	return m.IpMulticastAddress
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataBACnetIPMulticastAddress) GetActualValue() BACnetApplicationTagOctetString {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagOctetString(m.GetIpMulticastAddress())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataBACnetIPMulticastAddress(structType any) BACnetConstructedDataBACnetIPMulticastAddress {
	if casted, ok := structType.(BACnetConstructedDataBACnetIPMulticastAddress); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataBACnetIPMulticastAddress); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataBACnetIPMulticastAddress) GetTypeName() string {
	return "BACnetConstructedDataBACnetIPMulticastAddress"
}

func (m *_BACnetConstructedDataBACnetIPMulticastAddress) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (ipMulticastAddress)
	lengthInBits += m.IpMulticastAddress.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataBACnetIPMulticastAddress) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataBACnetIPMulticastAddress) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataBACnetIPMulticastAddress BACnetConstructedDataBACnetIPMulticastAddress, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataBACnetIPMulticastAddress"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataBACnetIPMulticastAddress")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	ipMulticastAddress, err := ReadSimpleField[BACnetApplicationTagOctetString](ctx, "ipMulticastAddress", ReadComplex[BACnetApplicationTagOctetString](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagOctetString](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'ipMulticastAddress' field"))
	}
	m.IpMulticastAddress = ipMulticastAddress

	actualValue, err := ReadVirtualField[BACnetApplicationTagOctetString](ctx, "actualValue", (*BACnetApplicationTagOctetString)(nil), ipMulticastAddress)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataBACnetIPMulticastAddress"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataBACnetIPMulticastAddress")
	}

	return m, nil
}

func (m *_BACnetConstructedDataBACnetIPMulticastAddress) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataBACnetIPMulticastAddress) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataBACnetIPMulticastAddress"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataBACnetIPMulticastAddress")
		}

		if err := WriteSimpleField[BACnetApplicationTagOctetString](ctx, "ipMulticastAddress", m.GetIpMulticastAddress(), WriteComplex[BACnetApplicationTagOctetString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'ipMulticastAddress' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataBACnetIPMulticastAddress"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataBACnetIPMulticastAddress")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataBACnetIPMulticastAddress) IsBACnetConstructedDataBACnetIPMulticastAddress() {
}

func (m *_BACnetConstructedDataBACnetIPMulticastAddress) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataBACnetIPMulticastAddress) deepCopy() *_BACnetConstructedDataBACnetIPMulticastAddress {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataBACnetIPMulticastAddressCopy := &_BACnetConstructedDataBACnetIPMulticastAddress{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagOctetString](m.IpMulticastAddress),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataBACnetIPMulticastAddressCopy
}

func (m *_BACnetConstructedDataBACnetIPMulticastAddress) String() string {
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
