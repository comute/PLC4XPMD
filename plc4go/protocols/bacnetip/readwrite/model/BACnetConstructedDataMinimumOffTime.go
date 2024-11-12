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

// BACnetConstructedDataMinimumOffTime is the corresponding interface of BACnetConstructedDataMinimumOffTime
type BACnetConstructedDataMinimumOffTime interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetMinimumOffTime returns MinimumOffTime (property field)
	GetMinimumOffTime() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetConstructedDataMinimumOffTime is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataMinimumOffTime()
	// CreateBuilder creates a BACnetConstructedDataMinimumOffTimeBuilder
	CreateBACnetConstructedDataMinimumOffTimeBuilder() BACnetConstructedDataMinimumOffTimeBuilder
}

// _BACnetConstructedDataMinimumOffTime is the data-structure of this message
type _BACnetConstructedDataMinimumOffTime struct {
	BACnetConstructedDataContract
	MinimumOffTime BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataMinimumOffTime = (*_BACnetConstructedDataMinimumOffTime)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataMinimumOffTime)(nil)

// NewBACnetConstructedDataMinimumOffTime factory function for _BACnetConstructedDataMinimumOffTime
func NewBACnetConstructedDataMinimumOffTime(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, minimumOffTime BACnetApplicationTagUnsignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataMinimumOffTime {
	if minimumOffTime == nil {
		panic("minimumOffTime of type BACnetApplicationTagUnsignedInteger for BACnetConstructedDataMinimumOffTime must not be nil")
	}
	_result := &_BACnetConstructedDataMinimumOffTime{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		MinimumOffTime:                minimumOffTime,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataMinimumOffTimeBuilder is a builder for BACnetConstructedDataMinimumOffTime
type BACnetConstructedDataMinimumOffTimeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(minimumOffTime BACnetApplicationTagUnsignedInteger) BACnetConstructedDataMinimumOffTimeBuilder
	// WithMinimumOffTime adds MinimumOffTime (property field)
	WithMinimumOffTime(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataMinimumOffTimeBuilder
	// WithMinimumOffTimeBuilder adds MinimumOffTime (property field) which is build by the builder
	WithMinimumOffTimeBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataMinimumOffTimeBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataMinimumOffTime or returns an error if something is wrong
	Build() (BACnetConstructedDataMinimumOffTime, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataMinimumOffTime
}

// NewBACnetConstructedDataMinimumOffTimeBuilder() creates a BACnetConstructedDataMinimumOffTimeBuilder
func NewBACnetConstructedDataMinimumOffTimeBuilder() BACnetConstructedDataMinimumOffTimeBuilder {
	return &_BACnetConstructedDataMinimumOffTimeBuilder{_BACnetConstructedDataMinimumOffTime: new(_BACnetConstructedDataMinimumOffTime)}
}

type _BACnetConstructedDataMinimumOffTimeBuilder struct {
	*_BACnetConstructedDataMinimumOffTime

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataMinimumOffTimeBuilder) = (*_BACnetConstructedDataMinimumOffTimeBuilder)(nil)

func (b *_BACnetConstructedDataMinimumOffTimeBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataMinimumOffTime
}

func (b *_BACnetConstructedDataMinimumOffTimeBuilder) WithMandatoryFields(minimumOffTime BACnetApplicationTagUnsignedInteger) BACnetConstructedDataMinimumOffTimeBuilder {
	return b.WithMinimumOffTime(minimumOffTime)
}

func (b *_BACnetConstructedDataMinimumOffTimeBuilder) WithMinimumOffTime(minimumOffTime BACnetApplicationTagUnsignedInteger) BACnetConstructedDataMinimumOffTimeBuilder {
	b.MinimumOffTime = minimumOffTime
	return b
}

func (b *_BACnetConstructedDataMinimumOffTimeBuilder) WithMinimumOffTimeBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataMinimumOffTimeBuilder {
	builder := builderSupplier(b.MinimumOffTime.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	b.MinimumOffTime, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataMinimumOffTimeBuilder) Build() (BACnetConstructedDataMinimumOffTime, error) {
	if b.MinimumOffTime == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'minimumOffTime' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataMinimumOffTime.deepCopy(), nil
}

func (b *_BACnetConstructedDataMinimumOffTimeBuilder) MustBuild() BACnetConstructedDataMinimumOffTime {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataMinimumOffTimeBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataMinimumOffTimeBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataMinimumOffTimeBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataMinimumOffTimeBuilder().(*_BACnetConstructedDataMinimumOffTimeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataMinimumOffTimeBuilder creates a BACnetConstructedDataMinimumOffTimeBuilder
func (b *_BACnetConstructedDataMinimumOffTime) CreateBACnetConstructedDataMinimumOffTimeBuilder() BACnetConstructedDataMinimumOffTimeBuilder {
	if b == nil {
		return NewBACnetConstructedDataMinimumOffTimeBuilder()
	}
	return &_BACnetConstructedDataMinimumOffTimeBuilder{_BACnetConstructedDataMinimumOffTime: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataMinimumOffTime) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataMinimumOffTime) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_MINIMUM_OFF_TIME
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataMinimumOffTime) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataMinimumOffTime) GetMinimumOffTime() BACnetApplicationTagUnsignedInteger {
	return m.MinimumOffTime
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataMinimumOffTime) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetMinimumOffTime())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataMinimumOffTime(structType any) BACnetConstructedDataMinimumOffTime {
	if casted, ok := structType.(BACnetConstructedDataMinimumOffTime); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataMinimumOffTime); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataMinimumOffTime) GetTypeName() string {
	return "BACnetConstructedDataMinimumOffTime"
}

func (m *_BACnetConstructedDataMinimumOffTime) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (minimumOffTime)
	lengthInBits += m.MinimumOffTime.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataMinimumOffTime) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataMinimumOffTime) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataMinimumOffTime BACnetConstructedDataMinimumOffTime, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataMinimumOffTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataMinimumOffTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	minimumOffTime, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "minimumOffTime", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'minimumOffTime' field"))
	}
	m.MinimumOffTime = minimumOffTime

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), minimumOffTime)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataMinimumOffTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataMinimumOffTime")
	}

	return m, nil
}

func (m *_BACnetConstructedDataMinimumOffTime) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataMinimumOffTime) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataMinimumOffTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataMinimumOffTime")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "minimumOffTime", m.GetMinimumOffTime(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'minimumOffTime' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataMinimumOffTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataMinimumOffTime")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataMinimumOffTime) IsBACnetConstructedDataMinimumOffTime() {}

func (m *_BACnetConstructedDataMinimumOffTime) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataMinimumOffTime) deepCopy() *_BACnetConstructedDataMinimumOffTime {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataMinimumOffTimeCopy := &_BACnetConstructedDataMinimumOffTime{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagUnsignedInteger](m.MinimumOffTime),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataMinimumOffTimeCopy
}

func (m *_BACnetConstructedDataMinimumOffTime) String() string {
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
