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

// BACnetConstructedDataEventDetectionEnable is the corresponding interface of BACnetConstructedDataEventDetectionEnable
type BACnetConstructedDataEventDetectionEnable interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetEventDetectionEnable returns EventDetectionEnable (property field)
	GetEventDetectionEnable() BACnetApplicationTagBoolean
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagBoolean
	// IsBACnetConstructedDataEventDetectionEnable is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataEventDetectionEnable()
	// CreateBuilder creates a BACnetConstructedDataEventDetectionEnableBuilder
	CreateBACnetConstructedDataEventDetectionEnableBuilder() BACnetConstructedDataEventDetectionEnableBuilder
}

// _BACnetConstructedDataEventDetectionEnable is the data-structure of this message
type _BACnetConstructedDataEventDetectionEnable struct {
	BACnetConstructedDataContract
	EventDetectionEnable BACnetApplicationTagBoolean
}

var _ BACnetConstructedDataEventDetectionEnable = (*_BACnetConstructedDataEventDetectionEnable)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataEventDetectionEnable)(nil)

// NewBACnetConstructedDataEventDetectionEnable factory function for _BACnetConstructedDataEventDetectionEnable
func NewBACnetConstructedDataEventDetectionEnable(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, eventDetectionEnable BACnetApplicationTagBoolean, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataEventDetectionEnable {
	if eventDetectionEnable == nil {
		panic("eventDetectionEnable of type BACnetApplicationTagBoolean for BACnetConstructedDataEventDetectionEnable must not be nil")
	}
	_result := &_BACnetConstructedDataEventDetectionEnable{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		EventDetectionEnable:          eventDetectionEnable,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataEventDetectionEnableBuilder is a builder for BACnetConstructedDataEventDetectionEnable
type BACnetConstructedDataEventDetectionEnableBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(eventDetectionEnable BACnetApplicationTagBoolean) BACnetConstructedDataEventDetectionEnableBuilder
	// WithEventDetectionEnable adds EventDetectionEnable (property field)
	WithEventDetectionEnable(BACnetApplicationTagBoolean) BACnetConstructedDataEventDetectionEnableBuilder
	// WithEventDetectionEnableBuilder adds EventDetectionEnable (property field) which is build by the builder
	WithEventDetectionEnableBuilder(func(BACnetApplicationTagBooleanBuilder) BACnetApplicationTagBooleanBuilder) BACnetConstructedDataEventDetectionEnableBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataEventDetectionEnable or returns an error if something is wrong
	Build() (BACnetConstructedDataEventDetectionEnable, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataEventDetectionEnable
}

// NewBACnetConstructedDataEventDetectionEnableBuilder() creates a BACnetConstructedDataEventDetectionEnableBuilder
func NewBACnetConstructedDataEventDetectionEnableBuilder() BACnetConstructedDataEventDetectionEnableBuilder {
	return &_BACnetConstructedDataEventDetectionEnableBuilder{_BACnetConstructedDataEventDetectionEnable: new(_BACnetConstructedDataEventDetectionEnable)}
}

type _BACnetConstructedDataEventDetectionEnableBuilder struct {
	*_BACnetConstructedDataEventDetectionEnable

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataEventDetectionEnableBuilder) = (*_BACnetConstructedDataEventDetectionEnableBuilder)(nil)

func (b *_BACnetConstructedDataEventDetectionEnableBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataEventDetectionEnable
}

func (b *_BACnetConstructedDataEventDetectionEnableBuilder) WithMandatoryFields(eventDetectionEnable BACnetApplicationTagBoolean) BACnetConstructedDataEventDetectionEnableBuilder {
	return b.WithEventDetectionEnable(eventDetectionEnable)
}

func (b *_BACnetConstructedDataEventDetectionEnableBuilder) WithEventDetectionEnable(eventDetectionEnable BACnetApplicationTagBoolean) BACnetConstructedDataEventDetectionEnableBuilder {
	b.EventDetectionEnable = eventDetectionEnable
	return b
}

func (b *_BACnetConstructedDataEventDetectionEnableBuilder) WithEventDetectionEnableBuilder(builderSupplier func(BACnetApplicationTagBooleanBuilder) BACnetApplicationTagBooleanBuilder) BACnetConstructedDataEventDetectionEnableBuilder {
	builder := builderSupplier(b.EventDetectionEnable.CreateBACnetApplicationTagBooleanBuilder())
	var err error
	b.EventDetectionEnable, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagBooleanBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataEventDetectionEnableBuilder) Build() (BACnetConstructedDataEventDetectionEnable, error) {
	if b.EventDetectionEnable == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'eventDetectionEnable' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataEventDetectionEnable.deepCopy(), nil
}

func (b *_BACnetConstructedDataEventDetectionEnableBuilder) MustBuild() BACnetConstructedDataEventDetectionEnable {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataEventDetectionEnableBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataEventDetectionEnableBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataEventDetectionEnableBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataEventDetectionEnableBuilder().(*_BACnetConstructedDataEventDetectionEnableBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataEventDetectionEnableBuilder creates a BACnetConstructedDataEventDetectionEnableBuilder
func (b *_BACnetConstructedDataEventDetectionEnable) CreateBACnetConstructedDataEventDetectionEnableBuilder() BACnetConstructedDataEventDetectionEnableBuilder {
	if b == nil {
		return NewBACnetConstructedDataEventDetectionEnableBuilder()
	}
	return &_BACnetConstructedDataEventDetectionEnableBuilder{_BACnetConstructedDataEventDetectionEnable: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataEventDetectionEnable) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataEventDetectionEnable) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_EVENT_DETECTION_ENABLE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataEventDetectionEnable) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataEventDetectionEnable) GetEventDetectionEnable() BACnetApplicationTagBoolean {
	return m.EventDetectionEnable
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataEventDetectionEnable) GetActualValue() BACnetApplicationTagBoolean {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagBoolean(m.GetEventDetectionEnable())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataEventDetectionEnable(structType any) BACnetConstructedDataEventDetectionEnable {
	if casted, ok := structType.(BACnetConstructedDataEventDetectionEnable); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataEventDetectionEnable); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataEventDetectionEnable) GetTypeName() string {
	return "BACnetConstructedDataEventDetectionEnable"
}

func (m *_BACnetConstructedDataEventDetectionEnable) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (eventDetectionEnable)
	lengthInBits += m.EventDetectionEnable.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataEventDetectionEnable) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataEventDetectionEnable) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataEventDetectionEnable BACnetConstructedDataEventDetectionEnable, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataEventDetectionEnable"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataEventDetectionEnable")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	eventDetectionEnable, err := ReadSimpleField[BACnetApplicationTagBoolean](ctx, "eventDetectionEnable", ReadComplex[BACnetApplicationTagBoolean](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagBoolean](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'eventDetectionEnable' field"))
	}
	m.EventDetectionEnable = eventDetectionEnable

	actualValue, err := ReadVirtualField[BACnetApplicationTagBoolean](ctx, "actualValue", (*BACnetApplicationTagBoolean)(nil), eventDetectionEnable)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataEventDetectionEnable"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataEventDetectionEnable")
	}

	return m, nil
}

func (m *_BACnetConstructedDataEventDetectionEnable) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataEventDetectionEnable) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataEventDetectionEnable"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataEventDetectionEnable")
		}

		if err := WriteSimpleField[BACnetApplicationTagBoolean](ctx, "eventDetectionEnable", m.GetEventDetectionEnable(), WriteComplex[BACnetApplicationTagBoolean](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'eventDetectionEnable' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataEventDetectionEnable"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataEventDetectionEnable")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataEventDetectionEnable) IsBACnetConstructedDataEventDetectionEnable() {}

func (m *_BACnetConstructedDataEventDetectionEnable) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataEventDetectionEnable) deepCopy() *_BACnetConstructedDataEventDetectionEnable {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataEventDetectionEnableCopy := &_BACnetConstructedDataEventDetectionEnable{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagBoolean](m.EventDetectionEnable),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataEventDetectionEnableCopy
}

func (m *_BACnetConstructedDataEventDetectionEnable) String() string {
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
