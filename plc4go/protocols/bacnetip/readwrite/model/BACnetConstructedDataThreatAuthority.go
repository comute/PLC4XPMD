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

// BACnetConstructedDataThreatAuthority is the corresponding interface of BACnetConstructedDataThreatAuthority
type BACnetConstructedDataThreatAuthority interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetThreatAuthority returns ThreatAuthority (property field)
	GetThreatAuthority() BACnetAccessThreatLevel
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetAccessThreatLevel
	// IsBACnetConstructedDataThreatAuthority is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataThreatAuthority()
	// CreateBuilder creates a BACnetConstructedDataThreatAuthorityBuilder
	CreateBACnetConstructedDataThreatAuthorityBuilder() BACnetConstructedDataThreatAuthorityBuilder
}

// _BACnetConstructedDataThreatAuthority is the data-structure of this message
type _BACnetConstructedDataThreatAuthority struct {
	BACnetConstructedDataContract
	ThreatAuthority BACnetAccessThreatLevel
}

var _ BACnetConstructedDataThreatAuthority = (*_BACnetConstructedDataThreatAuthority)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataThreatAuthority)(nil)

// NewBACnetConstructedDataThreatAuthority factory function for _BACnetConstructedDataThreatAuthority
func NewBACnetConstructedDataThreatAuthority(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, threatAuthority BACnetAccessThreatLevel, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataThreatAuthority {
	if threatAuthority == nil {
		panic("threatAuthority of type BACnetAccessThreatLevel for BACnetConstructedDataThreatAuthority must not be nil")
	}
	_result := &_BACnetConstructedDataThreatAuthority{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		ThreatAuthority:               threatAuthority,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataThreatAuthorityBuilder is a builder for BACnetConstructedDataThreatAuthority
type BACnetConstructedDataThreatAuthorityBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(threatAuthority BACnetAccessThreatLevel) BACnetConstructedDataThreatAuthorityBuilder
	// WithThreatAuthority adds ThreatAuthority (property field)
	WithThreatAuthority(BACnetAccessThreatLevel) BACnetConstructedDataThreatAuthorityBuilder
	// WithThreatAuthorityBuilder adds ThreatAuthority (property field) which is build by the builder
	WithThreatAuthorityBuilder(func(BACnetAccessThreatLevelBuilder) BACnetAccessThreatLevelBuilder) BACnetConstructedDataThreatAuthorityBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataThreatAuthority or returns an error if something is wrong
	Build() (BACnetConstructedDataThreatAuthority, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataThreatAuthority
}

// NewBACnetConstructedDataThreatAuthorityBuilder() creates a BACnetConstructedDataThreatAuthorityBuilder
func NewBACnetConstructedDataThreatAuthorityBuilder() BACnetConstructedDataThreatAuthorityBuilder {
	return &_BACnetConstructedDataThreatAuthorityBuilder{_BACnetConstructedDataThreatAuthority: new(_BACnetConstructedDataThreatAuthority)}
}

type _BACnetConstructedDataThreatAuthorityBuilder struct {
	*_BACnetConstructedDataThreatAuthority

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataThreatAuthorityBuilder) = (*_BACnetConstructedDataThreatAuthorityBuilder)(nil)

func (b *_BACnetConstructedDataThreatAuthorityBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataThreatAuthority
}

func (b *_BACnetConstructedDataThreatAuthorityBuilder) WithMandatoryFields(threatAuthority BACnetAccessThreatLevel) BACnetConstructedDataThreatAuthorityBuilder {
	return b.WithThreatAuthority(threatAuthority)
}

func (b *_BACnetConstructedDataThreatAuthorityBuilder) WithThreatAuthority(threatAuthority BACnetAccessThreatLevel) BACnetConstructedDataThreatAuthorityBuilder {
	b.ThreatAuthority = threatAuthority
	return b
}

func (b *_BACnetConstructedDataThreatAuthorityBuilder) WithThreatAuthorityBuilder(builderSupplier func(BACnetAccessThreatLevelBuilder) BACnetAccessThreatLevelBuilder) BACnetConstructedDataThreatAuthorityBuilder {
	builder := builderSupplier(b.ThreatAuthority.CreateBACnetAccessThreatLevelBuilder())
	var err error
	b.ThreatAuthority, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetAccessThreatLevelBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataThreatAuthorityBuilder) Build() (BACnetConstructedDataThreatAuthority, error) {
	if b.ThreatAuthority == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'threatAuthority' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataThreatAuthority.deepCopy(), nil
}

func (b *_BACnetConstructedDataThreatAuthorityBuilder) MustBuild() BACnetConstructedDataThreatAuthority {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataThreatAuthorityBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataThreatAuthorityBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataThreatAuthorityBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataThreatAuthorityBuilder().(*_BACnetConstructedDataThreatAuthorityBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataThreatAuthorityBuilder creates a BACnetConstructedDataThreatAuthorityBuilder
func (b *_BACnetConstructedDataThreatAuthority) CreateBACnetConstructedDataThreatAuthorityBuilder() BACnetConstructedDataThreatAuthorityBuilder {
	if b == nil {
		return NewBACnetConstructedDataThreatAuthorityBuilder()
	}
	return &_BACnetConstructedDataThreatAuthorityBuilder{_BACnetConstructedDataThreatAuthority: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataThreatAuthority) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataThreatAuthority) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_THREAT_AUTHORITY
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataThreatAuthority) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataThreatAuthority) GetThreatAuthority() BACnetAccessThreatLevel {
	return m.ThreatAuthority
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataThreatAuthority) GetActualValue() BACnetAccessThreatLevel {
	ctx := context.Background()
	_ = ctx
	return CastBACnetAccessThreatLevel(m.GetThreatAuthority())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataThreatAuthority(structType any) BACnetConstructedDataThreatAuthority {
	if casted, ok := structType.(BACnetConstructedDataThreatAuthority); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataThreatAuthority); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataThreatAuthority) GetTypeName() string {
	return "BACnetConstructedDataThreatAuthority"
}

func (m *_BACnetConstructedDataThreatAuthority) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (threatAuthority)
	lengthInBits += m.ThreatAuthority.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataThreatAuthority) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataThreatAuthority) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataThreatAuthority BACnetConstructedDataThreatAuthority, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataThreatAuthority"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataThreatAuthority")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	threatAuthority, err := ReadSimpleField[BACnetAccessThreatLevel](ctx, "threatAuthority", ReadComplex[BACnetAccessThreatLevel](BACnetAccessThreatLevelParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'threatAuthority' field"))
	}
	m.ThreatAuthority = threatAuthority

	actualValue, err := ReadVirtualField[BACnetAccessThreatLevel](ctx, "actualValue", (*BACnetAccessThreatLevel)(nil), threatAuthority)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataThreatAuthority"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataThreatAuthority")
	}

	return m, nil
}

func (m *_BACnetConstructedDataThreatAuthority) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataThreatAuthority) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataThreatAuthority"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataThreatAuthority")
		}

		if err := WriteSimpleField[BACnetAccessThreatLevel](ctx, "threatAuthority", m.GetThreatAuthority(), WriteComplex[BACnetAccessThreatLevel](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'threatAuthority' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataThreatAuthority"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataThreatAuthority")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataThreatAuthority) IsBACnetConstructedDataThreatAuthority() {}

func (m *_BACnetConstructedDataThreatAuthority) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataThreatAuthority) deepCopy() *_BACnetConstructedDataThreatAuthority {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataThreatAuthorityCopy := &_BACnetConstructedDataThreatAuthority{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetAccessThreatLevel](m.ThreatAuthority),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataThreatAuthorityCopy
}

func (m *_BACnetConstructedDataThreatAuthority) String() string {
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
