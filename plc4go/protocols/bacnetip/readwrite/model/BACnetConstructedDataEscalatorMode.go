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

// BACnetConstructedDataEscalatorMode is the corresponding interface of BACnetConstructedDataEscalatorMode
type BACnetConstructedDataEscalatorMode interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetEscalatorMode returns EscalatorMode (property field)
	GetEscalatorMode() BACnetEscalatorModeTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetEscalatorModeTagged
	// IsBACnetConstructedDataEscalatorMode is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataEscalatorMode()
	// CreateBuilder creates a BACnetConstructedDataEscalatorModeBuilder
	CreateBACnetConstructedDataEscalatorModeBuilder() BACnetConstructedDataEscalatorModeBuilder
}

// _BACnetConstructedDataEscalatorMode is the data-structure of this message
type _BACnetConstructedDataEscalatorMode struct {
	BACnetConstructedDataContract
	EscalatorMode BACnetEscalatorModeTagged
}

var _ BACnetConstructedDataEscalatorMode = (*_BACnetConstructedDataEscalatorMode)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataEscalatorMode)(nil)

// NewBACnetConstructedDataEscalatorMode factory function for _BACnetConstructedDataEscalatorMode
func NewBACnetConstructedDataEscalatorMode(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, escalatorMode BACnetEscalatorModeTagged, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataEscalatorMode {
	if escalatorMode == nil {
		panic("escalatorMode of type BACnetEscalatorModeTagged for BACnetConstructedDataEscalatorMode must not be nil")
	}
	_result := &_BACnetConstructedDataEscalatorMode{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		EscalatorMode:                 escalatorMode,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataEscalatorModeBuilder is a builder for BACnetConstructedDataEscalatorMode
type BACnetConstructedDataEscalatorModeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(escalatorMode BACnetEscalatorModeTagged) BACnetConstructedDataEscalatorModeBuilder
	// WithEscalatorMode adds EscalatorMode (property field)
	WithEscalatorMode(BACnetEscalatorModeTagged) BACnetConstructedDataEscalatorModeBuilder
	// WithEscalatorModeBuilder adds EscalatorMode (property field) which is build by the builder
	WithEscalatorModeBuilder(func(BACnetEscalatorModeTaggedBuilder) BACnetEscalatorModeTaggedBuilder) BACnetConstructedDataEscalatorModeBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataEscalatorMode or returns an error if something is wrong
	Build() (BACnetConstructedDataEscalatorMode, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataEscalatorMode
}

// NewBACnetConstructedDataEscalatorModeBuilder() creates a BACnetConstructedDataEscalatorModeBuilder
func NewBACnetConstructedDataEscalatorModeBuilder() BACnetConstructedDataEscalatorModeBuilder {
	return &_BACnetConstructedDataEscalatorModeBuilder{_BACnetConstructedDataEscalatorMode: new(_BACnetConstructedDataEscalatorMode)}
}

type _BACnetConstructedDataEscalatorModeBuilder struct {
	*_BACnetConstructedDataEscalatorMode

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataEscalatorModeBuilder) = (*_BACnetConstructedDataEscalatorModeBuilder)(nil)

func (b *_BACnetConstructedDataEscalatorModeBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataEscalatorMode
}

func (b *_BACnetConstructedDataEscalatorModeBuilder) WithMandatoryFields(escalatorMode BACnetEscalatorModeTagged) BACnetConstructedDataEscalatorModeBuilder {
	return b.WithEscalatorMode(escalatorMode)
}

func (b *_BACnetConstructedDataEscalatorModeBuilder) WithEscalatorMode(escalatorMode BACnetEscalatorModeTagged) BACnetConstructedDataEscalatorModeBuilder {
	b.EscalatorMode = escalatorMode
	return b
}

func (b *_BACnetConstructedDataEscalatorModeBuilder) WithEscalatorModeBuilder(builderSupplier func(BACnetEscalatorModeTaggedBuilder) BACnetEscalatorModeTaggedBuilder) BACnetConstructedDataEscalatorModeBuilder {
	builder := builderSupplier(b.EscalatorMode.CreateBACnetEscalatorModeTaggedBuilder())
	var err error
	b.EscalatorMode, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetEscalatorModeTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataEscalatorModeBuilder) Build() (BACnetConstructedDataEscalatorMode, error) {
	if b.EscalatorMode == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'escalatorMode' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataEscalatorMode.deepCopy(), nil
}

func (b *_BACnetConstructedDataEscalatorModeBuilder) MustBuild() BACnetConstructedDataEscalatorMode {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataEscalatorModeBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataEscalatorModeBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataEscalatorModeBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataEscalatorModeBuilder().(*_BACnetConstructedDataEscalatorModeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataEscalatorModeBuilder creates a BACnetConstructedDataEscalatorModeBuilder
func (b *_BACnetConstructedDataEscalatorMode) CreateBACnetConstructedDataEscalatorModeBuilder() BACnetConstructedDataEscalatorModeBuilder {
	if b == nil {
		return NewBACnetConstructedDataEscalatorModeBuilder()
	}
	return &_BACnetConstructedDataEscalatorModeBuilder{_BACnetConstructedDataEscalatorMode: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataEscalatorMode) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataEscalatorMode) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ESCALATOR_MODE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataEscalatorMode) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataEscalatorMode) GetEscalatorMode() BACnetEscalatorModeTagged {
	return m.EscalatorMode
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataEscalatorMode) GetActualValue() BACnetEscalatorModeTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetEscalatorModeTagged(m.GetEscalatorMode())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataEscalatorMode(structType any) BACnetConstructedDataEscalatorMode {
	if casted, ok := structType.(BACnetConstructedDataEscalatorMode); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataEscalatorMode); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataEscalatorMode) GetTypeName() string {
	return "BACnetConstructedDataEscalatorMode"
}

func (m *_BACnetConstructedDataEscalatorMode) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (escalatorMode)
	lengthInBits += m.EscalatorMode.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataEscalatorMode) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataEscalatorMode) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataEscalatorMode BACnetConstructedDataEscalatorMode, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataEscalatorMode"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataEscalatorMode")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	escalatorMode, err := ReadSimpleField[BACnetEscalatorModeTagged](ctx, "escalatorMode", ReadComplex[BACnetEscalatorModeTagged](BACnetEscalatorModeTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'escalatorMode' field"))
	}
	m.EscalatorMode = escalatorMode

	actualValue, err := ReadVirtualField[BACnetEscalatorModeTagged](ctx, "actualValue", (*BACnetEscalatorModeTagged)(nil), escalatorMode)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataEscalatorMode"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataEscalatorMode")
	}

	return m, nil
}

func (m *_BACnetConstructedDataEscalatorMode) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataEscalatorMode) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataEscalatorMode"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataEscalatorMode")
		}

		if err := WriteSimpleField[BACnetEscalatorModeTagged](ctx, "escalatorMode", m.GetEscalatorMode(), WriteComplex[BACnetEscalatorModeTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'escalatorMode' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataEscalatorMode"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataEscalatorMode")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataEscalatorMode) IsBACnetConstructedDataEscalatorMode() {}

func (m *_BACnetConstructedDataEscalatorMode) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataEscalatorMode) deepCopy() *_BACnetConstructedDataEscalatorMode {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataEscalatorModeCopy := &_BACnetConstructedDataEscalatorMode{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetEscalatorModeTagged](m.EscalatorMode),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataEscalatorModeCopy
}

func (m *_BACnetConstructedDataEscalatorMode) String() string {
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
