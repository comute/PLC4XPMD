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

// BACnetConstructedDataCredentialStatus is the corresponding interface of BACnetConstructedDataCredentialStatus
type BACnetConstructedDataCredentialStatus interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetBinaryPv returns BinaryPv (property field)
	GetBinaryPv() BACnetBinaryPVTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetBinaryPVTagged
	// IsBACnetConstructedDataCredentialStatus is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataCredentialStatus()
	// CreateBuilder creates a BACnetConstructedDataCredentialStatusBuilder
	CreateBACnetConstructedDataCredentialStatusBuilder() BACnetConstructedDataCredentialStatusBuilder
}

// _BACnetConstructedDataCredentialStatus is the data-structure of this message
type _BACnetConstructedDataCredentialStatus struct {
	BACnetConstructedDataContract
	BinaryPv BACnetBinaryPVTagged
}

var _ BACnetConstructedDataCredentialStatus = (*_BACnetConstructedDataCredentialStatus)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataCredentialStatus)(nil)

// NewBACnetConstructedDataCredentialStatus factory function for _BACnetConstructedDataCredentialStatus
func NewBACnetConstructedDataCredentialStatus(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, binaryPv BACnetBinaryPVTagged, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataCredentialStatus {
	if binaryPv == nil {
		panic("binaryPv of type BACnetBinaryPVTagged for BACnetConstructedDataCredentialStatus must not be nil")
	}
	_result := &_BACnetConstructedDataCredentialStatus{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		BinaryPv:                      binaryPv,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataCredentialStatusBuilder is a builder for BACnetConstructedDataCredentialStatus
type BACnetConstructedDataCredentialStatusBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(binaryPv BACnetBinaryPVTagged) BACnetConstructedDataCredentialStatusBuilder
	// WithBinaryPv adds BinaryPv (property field)
	WithBinaryPv(BACnetBinaryPVTagged) BACnetConstructedDataCredentialStatusBuilder
	// WithBinaryPvBuilder adds BinaryPv (property field) which is build by the builder
	WithBinaryPvBuilder(func(BACnetBinaryPVTaggedBuilder) BACnetBinaryPVTaggedBuilder) BACnetConstructedDataCredentialStatusBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataCredentialStatus or returns an error if something is wrong
	Build() (BACnetConstructedDataCredentialStatus, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataCredentialStatus
}

// NewBACnetConstructedDataCredentialStatusBuilder() creates a BACnetConstructedDataCredentialStatusBuilder
func NewBACnetConstructedDataCredentialStatusBuilder() BACnetConstructedDataCredentialStatusBuilder {
	return &_BACnetConstructedDataCredentialStatusBuilder{_BACnetConstructedDataCredentialStatus: new(_BACnetConstructedDataCredentialStatus)}
}

type _BACnetConstructedDataCredentialStatusBuilder struct {
	*_BACnetConstructedDataCredentialStatus

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataCredentialStatusBuilder) = (*_BACnetConstructedDataCredentialStatusBuilder)(nil)

func (b *_BACnetConstructedDataCredentialStatusBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataCredentialStatus
}

func (b *_BACnetConstructedDataCredentialStatusBuilder) WithMandatoryFields(binaryPv BACnetBinaryPVTagged) BACnetConstructedDataCredentialStatusBuilder {
	return b.WithBinaryPv(binaryPv)
}

func (b *_BACnetConstructedDataCredentialStatusBuilder) WithBinaryPv(binaryPv BACnetBinaryPVTagged) BACnetConstructedDataCredentialStatusBuilder {
	b.BinaryPv = binaryPv
	return b
}

func (b *_BACnetConstructedDataCredentialStatusBuilder) WithBinaryPvBuilder(builderSupplier func(BACnetBinaryPVTaggedBuilder) BACnetBinaryPVTaggedBuilder) BACnetConstructedDataCredentialStatusBuilder {
	builder := builderSupplier(b.BinaryPv.CreateBACnetBinaryPVTaggedBuilder())
	var err error
	b.BinaryPv, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetBinaryPVTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataCredentialStatusBuilder) Build() (BACnetConstructedDataCredentialStatus, error) {
	if b.BinaryPv == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'binaryPv' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataCredentialStatus.deepCopy(), nil
}

func (b *_BACnetConstructedDataCredentialStatusBuilder) MustBuild() BACnetConstructedDataCredentialStatus {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataCredentialStatusBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataCredentialStatusBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataCredentialStatusBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataCredentialStatusBuilder().(*_BACnetConstructedDataCredentialStatusBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataCredentialStatusBuilder creates a BACnetConstructedDataCredentialStatusBuilder
func (b *_BACnetConstructedDataCredentialStatus) CreateBACnetConstructedDataCredentialStatusBuilder() BACnetConstructedDataCredentialStatusBuilder {
	if b == nil {
		return NewBACnetConstructedDataCredentialStatusBuilder()
	}
	return &_BACnetConstructedDataCredentialStatusBuilder{_BACnetConstructedDataCredentialStatus: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataCredentialStatus) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataCredentialStatus) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_CREDENTIAL_STATUS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataCredentialStatus) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataCredentialStatus) GetBinaryPv() BACnetBinaryPVTagged {
	return m.BinaryPv
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataCredentialStatus) GetActualValue() BACnetBinaryPVTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetBinaryPVTagged(m.GetBinaryPv())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataCredentialStatus(structType any) BACnetConstructedDataCredentialStatus {
	if casted, ok := structType.(BACnetConstructedDataCredentialStatus); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataCredentialStatus); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataCredentialStatus) GetTypeName() string {
	return "BACnetConstructedDataCredentialStatus"
}

func (m *_BACnetConstructedDataCredentialStatus) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (binaryPv)
	lengthInBits += m.BinaryPv.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataCredentialStatus) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataCredentialStatus) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataCredentialStatus BACnetConstructedDataCredentialStatus, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataCredentialStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataCredentialStatus")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	binaryPv, err := ReadSimpleField[BACnetBinaryPVTagged](ctx, "binaryPv", ReadComplex[BACnetBinaryPVTagged](BACnetBinaryPVTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'binaryPv' field"))
	}
	m.BinaryPv = binaryPv

	actualValue, err := ReadVirtualField[BACnetBinaryPVTagged](ctx, "actualValue", (*BACnetBinaryPVTagged)(nil), binaryPv)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataCredentialStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataCredentialStatus")
	}

	return m, nil
}

func (m *_BACnetConstructedDataCredentialStatus) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataCredentialStatus) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataCredentialStatus"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataCredentialStatus")
		}

		if err := WriteSimpleField[BACnetBinaryPVTagged](ctx, "binaryPv", m.GetBinaryPv(), WriteComplex[BACnetBinaryPVTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'binaryPv' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataCredentialStatus"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataCredentialStatus")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataCredentialStatus) IsBACnetConstructedDataCredentialStatus() {}

func (m *_BACnetConstructedDataCredentialStatus) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataCredentialStatus) deepCopy() *_BACnetConstructedDataCredentialStatus {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataCredentialStatusCopy := &_BACnetConstructedDataCredentialStatus{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetBinaryPVTagged](m.BinaryPv),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataCredentialStatusCopy
}

func (m *_BACnetConstructedDataCredentialStatus) String() string {
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
