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

// BACnetConstructedDataLinkSpeedAutonegotiate is the corresponding interface of BACnetConstructedDataLinkSpeedAutonegotiate
type BACnetConstructedDataLinkSpeedAutonegotiate interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetLinkSpeedAutonegotiate returns LinkSpeedAutonegotiate (property field)
	GetLinkSpeedAutonegotiate() BACnetApplicationTagBoolean
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagBoolean
	// IsBACnetConstructedDataLinkSpeedAutonegotiate is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataLinkSpeedAutonegotiate()
	// CreateBuilder creates a BACnetConstructedDataLinkSpeedAutonegotiateBuilder
	CreateBACnetConstructedDataLinkSpeedAutonegotiateBuilder() BACnetConstructedDataLinkSpeedAutonegotiateBuilder
}

// _BACnetConstructedDataLinkSpeedAutonegotiate is the data-structure of this message
type _BACnetConstructedDataLinkSpeedAutonegotiate struct {
	BACnetConstructedDataContract
	LinkSpeedAutonegotiate BACnetApplicationTagBoolean
}

var _ BACnetConstructedDataLinkSpeedAutonegotiate = (*_BACnetConstructedDataLinkSpeedAutonegotiate)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataLinkSpeedAutonegotiate)(nil)

// NewBACnetConstructedDataLinkSpeedAutonegotiate factory function for _BACnetConstructedDataLinkSpeedAutonegotiate
func NewBACnetConstructedDataLinkSpeedAutonegotiate(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, linkSpeedAutonegotiate BACnetApplicationTagBoolean, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLinkSpeedAutonegotiate {
	if linkSpeedAutonegotiate == nil {
		panic("linkSpeedAutonegotiate of type BACnetApplicationTagBoolean for BACnetConstructedDataLinkSpeedAutonegotiate must not be nil")
	}
	_result := &_BACnetConstructedDataLinkSpeedAutonegotiate{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		LinkSpeedAutonegotiate:        linkSpeedAutonegotiate,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataLinkSpeedAutonegotiateBuilder is a builder for BACnetConstructedDataLinkSpeedAutonegotiate
type BACnetConstructedDataLinkSpeedAutonegotiateBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(linkSpeedAutonegotiate BACnetApplicationTagBoolean) BACnetConstructedDataLinkSpeedAutonegotiateBuilder
	// WithLinkSpeedAutonegotiate adds LinkSpeedAutonegotiate (property field)
	WithLinkSpeedAutonegotiate(BACnetApplicationTagBoolean) BACnetConstructedDataLinkSpeedAutonegotiateBuilder
	// WithLinkSpeedAutonegotiateBuilder adds LinkSpeedAutonegotiate (property field) which is build by the builder
	WithLinkSpeedAutonegotiateBuilder(func(BACnetApplicationTagBooleanBuilder) BACnetApplicationTagBooleanBuilder) BACnetConstructedDataLinkSpeedAutonegotiateBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataLinkSpeedAutonegotiate or returns an error if something is wrong
	Build() (BACnetConstructedDataLinkSpeedAutonegotiate, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataLinkSpeedAutonegotiate
}

// NewBACnetConstructedDataLinkSpeedAutonegotiateBuilder() creates a BACnetConstructedDataLinkSpeedAutonegotiateBuilder
func NewBACnetConstructedDataLinkSpeedAutonegotiateBuilder() BACnetConstructedDataLinkSpeedAutonegotiateBuilder {
	return &_BACnetConstructedDataLinkSpeedAutonegotiateBuilder{_BACnetConstructedDataLinkSpeedAutonegotiate: new(_BACnetConstructedDataLinkSpeedAutonegotiate)}
}

type _BACnetConstructedDataLinkSpeedAutonegotiateBuilder struct {
	*_BACnetConstructedDataLinkSpeedAutonegotiate

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataLinkSpeedAutonegotiateBuilder) = (*_BACnetConstructedDataLinkSpeedAutonegotiateBuilder)(nil)

func (b *_BACnetConstructedDataLinkSpeedAutonegotiateBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataLinkSpeedAutonegotiate
}

func (b *_BACnetConstructedDataLinkSpeedAutonegotiateBuilder) WithMandatoryFields(linkSpeedAutonegotiate BACnetApplicationTagBoolean) BACnetConstructedDataLinkSpeedAutonegotiateBuilder {
	return b.WithLinkSpeedAutonegotiate(linkSpeedAutonegotiate)
}

func (b *_BACnetConstructedDataLinkSpeedAutonegotiateBuilder) WithLinkSpeedAutonegotiate(linkSpeedAutonegotiate BACnetApplicationTagBoolean) BACnetConstructedDataLinkSpeedAutonegotiateBuilder {
	b.LinkSpeedAutonegotiate = linkSpeedAutonegotiate
	return b
}

func (b *_BACnetConstructedDataLinkSpeedAutonegotiateBuilder) WithLinkSpeedAutonegotiateBuilder(builderSupplier func(BACnetApplicationTagBooleanBuilder) BACnetApplicationTagBooleanBuilder) BACnetConstructedDataLinkSpeedAutonegotiateBuilder {
	builder := builderSupplier(b.LinkSpeedAutonegotiate.CreateBACnetApplicationTagBooleanBuilder())
	var err error
	b.LinkSpeedAutonegotiate, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagBooleanBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataLinkSpeedAutonegotiateBuilder) Build() (BACnetConstructedDataLinkSpeedAutonegotiate, error) {
	if b.LinkSpeedAutonegotiate == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'linkSpeedAutonegotiate' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataLinkSpeedAutonegotiate.deepCopy(), nil
}

func (b *_BACnetConstructedDataLinkSpeedAutonegotiateBuilder) MustBuild() BACnetConstructedDataLinkSpeedAutonegotiate {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataLinkSpeedAutonegotiateBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataLinkSpeedAutonegotiateBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataLinkSpeedAutonegotiateBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataLinkSpeedAutonegotiateBuilder().(*_BACnetConstructedDataLinkSpeedAutonegotiateBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataLinkSpeedAutonegotiateBuilder creates a BACnetConstructedDataLinkSpeedAutonegotiateBuilder
func (b *_BACnetConstructedDataLinkSpeedAutonegotiate) CreateBACnetConstructedDataLinkSpeedAutonegotiateBuilder() BACnetConstructedDataLinkSpeedAutonegotiateBuilder {
	if b == nil {
		return NewBACnetConstructedDataLinkSpeedAutonegotiateBuilder()
	}
	return &_BACnetConstructedDataLinkSpeedAutonegotiateBuilder{_BACnetConstructedDataLinkSpeedAutonegotiate: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLinkSpeedAutonegotiate) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataLinkSpeedAutonegotiate) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LINK_SPEED_AUTONEGOTIATE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLinkSpeedAutonegotiate) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLinkSpeedAutonegotiate) GetLinkSpeedAutonegotiate() BACnetApplicationTagBoolean {
	return m.LinkSpeedAutonegotiate
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataLinkSpeedAutonegotiate) GetActualValue() BACnetApplicationTagBoolean {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagBoolean(m.GetLinkSpeedAutonegotiate())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLinkSpeedAutonegotiate(structType any) BACnetConstructedDataLinkSpeedAutonegotiate {
	if casted, ok := structType.(BACnetConstructedDataLinkSpeedAutonegotiate); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLinkSpeedAutonegotiate); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLinkSpeedAutonegotiate) GetTypeName() string {
	return "BACnetConstructedDataLinkSpeedAutonegotiate"
}

func (m *_BACnetConstructedDataLinkSpeedAutonegotiate) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (linkSpeedAutonegotiate)
	lengthInBits += m.LinkSpeedAutonegotiate.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataLinkSpeedAutonegotiate) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataLinkSpeedAutonegotiate) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataLinkSpeedAutonegotiate BACnetConstructedDataLinkSpeedAutonegotiate, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLinkSpeedAutonegotiate"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLinkSpeedAutonegotiate")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	linkSpeedAutonegotiate, err := ReadSimpleField[BACnetApplicationTagBoolean](ctx, "linkSpeedAutonegotiate", ReadComplex[BACnetApplicationTagBoolean](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagBoolean](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'linkSpeedAutonegotiate' field"))
	}
	m.LinkSpeedAutonegotiate = linkSpeedAutonegotiate

	actualValue, err := ReadVirtualField[BACnetApplicationTagBoolean](ctx, "actualValue", (*BACnetApplicationTagBoolean)(nil), linkSpeedAutonegotiate)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLinkSpeedAutonegotiate"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLinkSpeedAutonegotiate")
	}

	return m, nil
}

func (m *_BACnetConstructedDataLinkSpeedAutonegotiate) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataLinkSpeedAutonegotiate) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLinkSpeedAutonegotiate"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLinkSpeedAutonegotiate")
		}

		if err := WriteSimpleField[BACnetApplicationTagBoolean](ctx, "linkSpeedAutonegotiate", m.GetLinkSpeedAutonegotiate(), WriteComplex[BACnetApplicationTagBoolean](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'linkSpeedAutonegotiate' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLinkSpeedAutonegotiate"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLinkSpeedAutonegotiate")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLinkSpeedAutonegotiate) IsBACnetConstructedDataLinkSpeedAutonegotiate() {
}

func (m *_BACnetConstructedDataLinkSpeedAutonegotiate) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataLinkSpeedAutonegotiate) deepCopy() *_BACnetConstructedDataLinkSpeedAutonegotiate {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataLinkSpeedAutonegotiateCopy := &_BACnetConstructedDataLinkSpeedAutonegotiate{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagBoolean](m.LinkSpeedAutonegotiate),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataLinkSpeedAutonegotiateCopy
}

func (m *_BACnetConstructedDataLinkSpeedAutonegotiate) String() string {
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
