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

// BACnetConstructedDataDoorOpenTooLongTime is the corresponding interface of BACnetConstructedDataDoorOpenTooLongTime
type BACnetConstructedDataDoorOpenTooLongTime interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetDoorOpenTooLongTime returns DoorOpenTooLongTime (property field)
	GetDoorOpenTooLongTime() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetConstructedDataDoorOpenTooLongTime is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataDoorOpenTooLongTime()
	// CreateBuilder creates a BACnetConstructedDataDoorOpenTooLongTimeBuilder
	CreateBACnetConstructedDataDoorOpenTooLongTimeBuilder() BACnetConstructedDataDoorOpenTooLongTimeBuilder
}

// _BACnetConstructedDataDoorOpenTooLongTime is the data-structure of this message
type _BACnetConstructedDataDoorOpenTooLongTime struct {
	BACnetConstructedDataContract
	DoorOpenTooLongTime BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataDoorOpenTooLongTime = (*_BACnetConstructedDataDoorOpenTooLongTime)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataDoorOpenTooLongTime)(nil)

// NewBACnetConstructedDataDoorOpenTooLongTime factory function for _BACnetConstructedDataDoorOpenTooLongTime
func NewBACnetConstructedDataDoorOpenTooLongTime(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, doorOpenTooLongTime BACnetApplicationTagUnsignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataDoorOpenTooLongTime {
	if doorOpenTooLongTime == nil {
		panic("doorOpenTooLongTime of type BACnetApplicationTagUnsignedInteger for BACnetConstructedDataDoorOpenTooLongTime must not be nil")
	}
	_result := &_BACnetConstructedDataDoorOpenTooLongTime{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		DoorOpenTooLongTime:           doorOpenTooLongTime,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataDoorOpenTooLongTimeBuilder is a builder for BACnetConstructedDataDoorOpenTooLongTime
type BACnetConstructedDataDoorOpenTooLongTimeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(doorOpenTooLongTime BACnetApplicationTagUnsignedInteger) BACnetConstructedDataDoorOpenTooLongTimeBuilder
	// WithDoorOpenTooLongTime adds DoorOpenTooLongTime (property field)
	WithDoorOpenTooLongTime(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataDoorOpenTooLongTimeBuilder
	// WithDoorOpenTooLongTimeBuilder adds DoorOpenTooLongTime (property field) which is build by the builder
	WithDoorOpenTooLongTimeBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataDoorOpenTooLongTimeBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataDoorOpenTooLongTime or returns an error if something is wrong
	Build() (BACnetConstructedDataDoorOpenTooLongTime, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataDoorOpenTooLongTime
}

// NewBACnetConstructedDataDoorOpenTooLongTimeBuilder() creates a BACnetConstructedDataDoorOpenTooLongTimeBuilder
func NewBACnetConstructedDataDoorOpenTooLongTimeBuilder() BACnetConstructedDataDoorOpenTooLongTimeBuilder {
	return &_BACnetConstructedDataDoorOpenTooLongTimeBuilder{_BACnetConstructedDataDoorOpenTooLongTime: new(_BACnetConstructedDataDoorOpenTooLongTime)}
}

type _BACnetConstructedDataDoorOpenTooLongTimeBuilder struct {
	*_BACnetConstructedDataDoorOpenTooLongTime

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataDoorOpenTooLongTimeBuilder) = (*_BACnetConstructedDataDoorOpenTooLongTimeBuilder)(nil)

func (b *_BACnetConstructedDataDoorOpenTooLongTimeBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataDoorOpenTooLongTime
}

func (b *_BACnetConstructedDataDoorOpenTooLongTimeBuilder) WithMandatoryFields(doorOpenTooLongTime BACnetApplicationTagUnsignedInteger) BACnetConstructedDataDoorOpenTooLongTimeBuilder {
	return b.WithDoorOpenTooLongTime(doorOpenTooLongTime)
}

func (b *_BACnetConstructedDataDoorOpenTooLongTimeBuilder) WithDoorOpenTooLongTime(doorOpenTooLongTime BACnetApplicationTagUnsignedInteger) BACnetConstructedDataDoorOpenTooLongTimeBuilder {
	b.DoorOpenTooLongTime = doorOpenTooLongTime
	return b
}

func (b *_BACnetConstructedDataDoorOpenTooLongTimeBuilder) WithDoorOpenTooLongTimeBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataDoorOpenTooLongTimeBuilder {
	builder := builderSupplier(b.DoorOpenTooLongTime.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	b.DoorOpenTooLongTime, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataDoorOpenTooLongTimeBuilder) Build() (BACnetConstructedDataDoorOpenTooLongTime, error) {
	if b.DoorOpenTooLongTime == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'doorOpenTooLongTime' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataDoorOpenTooLongTime.deepCopy(), nil
}

func (b *_BACnetConstructedDataDoorOpenTooLongTimeBuilder) MustBuild() BACnetConstructedDataDoorOpenTooLongTime {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataDoorOpenTooLongTimeBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataDoorOpenTooLongTimeBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataDoorOpenTooLongTimeBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataDoorOpenTooLongTimeBuilder().(*_BACnetConstructedDataDoorOpenTooLongTimeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataDoorOpenTooLongTimeBuilder creates a BACnetConstructedDataDoorOpenTooLongTimeBuilder
func (b *_BACnetConstructedDataDoorOpenTooLongTime) CreateBACnetConstructedDataDoorOpenTooLongTimeBuilder() BACnetConstructedDataDoorOpenTooLongTimeBuilder {
	if b == nil {
		return NewBACnetConstructedDataDoorOpenTooLongTimeBuilder()
	}
	return &_BACnetConstructedDataDoorOpenTooLongTimeBuilder{_BACnetConstructedDataDoorOpenTooLongTime: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataDoorOpenTooLongTime) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataDoorOpenTooLongTime) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_DOOR_OPEN_TOO_LONG_TIME
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataDoorOpenTooLongTime) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataDoorOpenTooLongTime) GetDoorOpenTooLongTime() BACnetApplicationTagUnsignedInteger {
	return m.DoorOpenTooLongTime
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataDoorOpenTooLongTime) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetDoorOpenTooLongTime())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataDoorOpenTooLongTime(structType any) BACnetConstructedDataDoorOpenTooLongTime {
	if casted, ok := structType.(BACnetConstructedDataDoorOpenTooLongTime); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataDoorOpenTooLongTime); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataDoorOpenTooLongTime) GetTypeName() string {
	return "BACnetConstructedDataDoorOpenTooLongTime"
}

func (m *_BACnetConstructedDataDoorOpenTooLongTime) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (doorOpenTooLongTime)
	lengthInBits += m.DoorOpenTooLongTime.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataDoorOpenTooLongTime) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataDoorOpenTooLongTime) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataDoorOpenTooLongTime BACnetConstructedDataDoorOpenTooLongTime, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataDoorOpenTooLongTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataDoorOpenTooLongTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	doorOpenTooLongTime, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "doorOpenTooLongTime", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'doorOpenTooLongTime' field"))
	}
	m.DoorOpenTooLongTime = doorOpenTooLongTime

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), doorOpenTooLongTime)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataDoorOpenTooLongTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataDoorOpenTooLongTime")
	}

	return m, nil
}

func (m *_BACnetConstructedDataDoorOpenTooLongTime) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataDoorOpenTooLongTime) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataDoorOpenTooLongTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataDoorOpenTooLongTime")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "doorOpenTooLongTime", m.GetDoorOpenTooLongTime(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'doorOpenTooLongTime' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataDoorOpenTooLongTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataDoorOpenTooLongTime")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataDoorOpenTooLongTime) IsBACnetConstructedDataDoorOpenTooLongTime() {}

func (m *_BACnetConstructedDataDoorOpenTooLongTime) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataDoorOpenTooLongTime) deepCopy() *_BACnetConstructedDataDoorOpenTooLongTime {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataDoorOpenTooLongTimeCopy := &_BACnetConstructedDataDoorOpenTooLongTime{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagUnsignedInteger](m.DoorOpenTooLongTime),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataDoorOpenTooLongTimeCopy
}

func (m *_BACnetConstructedDataDoorOpenTooLongTime) String() string {
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
