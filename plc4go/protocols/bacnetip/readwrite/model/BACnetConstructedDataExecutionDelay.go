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

// BACnetConstructedDataExecutionDelay is the corresponding interface of BACnetConstructedDataExecutionDelay
type BACnetConstructedDataExecutionDelay interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetNumberOfDataElements returns NumberOfDataElements (property field)
	GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger
	// GetExecutionDelay returns ExecutionDelay (property field)
	GetExecutionDelay() []BACnetApplicationTagUnsignedInteger
	// GetZero returns Zero (virtual field)
	GetZero() uint64
	// IsBACnetConstructedDataExecutionDelay is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataExecutionDelay()
	// CreateBuilder creates a BACnetConstructedDataExecutionDelayBuilder
	CreateBACnetConstructedDataExecutionDelayBuilder() BACnetConstructedDataExecutionDelayBuilder
}

// _BACnetConstructedDataExecutionDelay is the data-structure of this message
type _BACnetConstructedDataExecutionDelay struct {
	BACnetConstructedDataContract
	NumberOfDataElements BACnetApplicationTagUnsignedInteger
	ExecutionDelay       []BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataExecutionDelay = (*_BACnetConstructedDataExecutionDelay)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataExecutionDelay)(nil)

// NewBACnetConstructedDataExecutionDelay factory function for _BACnetConstructedDataExecutionDelay
func NewBACnetConstructedDataExecutionDelay(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, numberOfDataElements BACnetApplicationTagUnsignedInteger, executionDelay []BACnetApplicationTagUnsignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataExecutionDelay {
	_result := &_BACnetConstructedDataExecutionDelay{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		NumberOfDataElements:          numberOfDataElements,
		ExecutionDelay:                executionDelay,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataExecutionDelayBuilder is a builder for BACnetConstructedDataExecutionDelay
type BACnetConstructedDataExecutionDelayBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(executionDelay []BACnetApplicationTagUnsignedInteger) BACnetConstructedDataExecutionDelayBuilder
	// WithNumberOfDataElements adds NumberOfDataElements (property field)
	WithOptionalNumberOfDataElements(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataExecutionDelayBuilder
	// WithOptionalNumberOfDataElementsBuilder adds NumberOfDataElements (property field) which is build by the builder
	WithOptionalNumberOfDataElementsBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataExecutionDelayBuilder
	// WithExecutionDelay adds ExecutionDelay (property field)
	WithExecutionDelay(...BACnetApplicationTagUnsignedInteger) BACnetConstructedDataExecutionDelayBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataExecutionDelay or returns an error if something is wrong
	Build() (BACnetConstructedDataExecutionDelay, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataExecutionDelay
}

// NewBACnetConstructedDataExecutionDelayBuilder() creates a BACnetConstructedDataExecutionDelayBuilder
func NewBACnetConstructedDataExecutionDelayBuilder() BACnetConstructedDataExecutionDelayBuilder {
	return &_BACnetConstructedDataExecutionDelayBuilder{_BACnetConstructedDataExecutionDelay: new(_BACnetConstructedDataExecutionDelay)}
}

type _BACnetConstructedDataExecutionDelayBuilder struct {
	*_BACnetConstructedDataExecutionDelay

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataExecutionDelayBuilder) = (*_BACnetConstructedDataExecutionDelayBuilder)(nil)

func (b *_BACnetConstructedDataExecutionDelayBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataExecutionDelay
}

func (b *_BACnetConstructedDataExecutionDelayBuilder) WithMandatoryFields(executionDelay []BACnetApplicationTagUnsignedInteger) BACnetConstructedDataExecutionDelayBuilder {
	return b.WithExecutionDelay(executionDelay...)
}

func (b *_BACnetConstructedDataExecutionDelayBuilder) WithOptionalNumberOfDataElements(numberOfDataElements BACnetApplicationTagUnsignedInteger) BACnetConstructedDataExecutionDelayBuilder {
	b.NumberOfDataElements = numberOfDataElements
	return b
}

func (b *_BACnetConstructedDataExecutionDelayBuilder) WithOptionalNumberOfDataElementsBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataExecutionDelayBuilder {
	builder := builderSupplier(b.NumberOfDataElements.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	b.NumberOfDataElements, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataExecutionDelayBuilder) WithExecutionDelay(executionDelay ...BACnetApplicationTagUnsignedInteger) BACnetConstructedDataExecutionDelayBuilder {
	b.ExecutionDelay = executionDelay
	return b
}

func (b *_BACnetConstructedDataExecutionDelayBuilder) Build() (BACnetConstructedDataExecutionDelay, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataExecutionDelay.deepCopy(), nil
}

func (b *_BACnetConstructedDataExecutionDelayBuilder) MustBuild() BACnetConstructedDataExecutionDelay {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataExecutionDelayBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataExecutionDelayBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataExecutionDelayBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataExecutionDelayBuilder().(*_BACnetConstructedDataExecutionDelayBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataExecutionDelayBuilder creates a BACnetConstructedDataExecutionDelayBuilder
func (b *_BACnetConstructedDataExecutionDelay) CreateBACnetConstructedDataExecutionDelayBuilder() BACnetConstructedDataExecutionDelayBuilder {
	if b == nil {
		return NewBACnetConstructedDataExecutionDelayBuilder()
	}
	return &_BACnetConstructedDataExecutionDelayBuilder{_BACnetConstructedDataExecutionDelay: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataExecutionDelay) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataExecutionDelay) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_EXECUTION_DELAY
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataExecutionDelay) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataExecutionDelay) GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger {
	return m.NumberOfDataElements
}

func (m *_BACnetConstructedDataExecutionDelay) GetExecutionDelay() []BACnetApplicationTagUnsignedInteger {
	return m.ExecutionDelay
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataExecutionDelay) GetZero() uint64 {
	ctx := context.Background()
	_ = ctx
	numberOfDataElements := m.GetNumberOfDataElements()
	_ = numberOfDataElements
	return uint64(uint64(0))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataExecutionDelay(structType any) BACnetConstructedDataExecutionDelay {
	if casted, ok := structType.(BACnetConstructedDataExecutionDelay); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataExecutionDelay); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataExecutionDelay) GetTypeName() string {
	return "BACnetConstructedDataExecutionDelay"
}

func (m *_BACnetConstructedDataExecutionDelay) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// A virtual field doesn't have any in- or output.

	// Optional Field (numberOfDataElements)
	if m.NumberOfDataElements != nil {
		lengthInBits += m.NumberOfDataElements.GetLengthInBits(ctx)
	}

	// Array field
	if len(m.ExecutionDelay) > 0 {
		for _, element := range m.ExecutionDelay {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataExecutionDelay) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataExecutionDelay) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataExecutionDelay BACnetConstructedDataExecutionDelay, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataExecutionDelay"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataExecutionDelay")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	zero, err := ReadVirtualField[uint64](ctx, "zero", (*uint64)(nil), uint64(0))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'zero' field"))
	}
	_ = zero

	var numberOfDataElements BACnetApplicationTagUnsignedInteger
	_numberOfDataElements, err := ReadOptionalField[BACnetApplicationTagUnsignedInteger](ctx, "numberOfDataElements", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer), bool(bool((arrayIndexArgument) != (nil))) && bool(bool((arrayIndexArgument.GetActualValue()) == (zero))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'numberOfDataElements' field"))
	}
	if _numberOfDataElements != nil {
		numberOfDataElements = *_numberOfDataElements
		m.NumberOfDataElements = numberOfDataElements
	}

	executionDelay, err := ReadTerminatedArrayField[BACnetApplicationTagUnsignedInteger](ctx, "executionDelay", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'executionDelay' field"))
	}
	m.ExecutionDelay = executionDelay

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataExecutionDelay"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataExecutionDelay")
	}

	return m, nil
}

func (m *_BACnetConstructedDataExecutionDelay) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataExecutionDelay) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataExecutionDelay"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataExecutionDelay")
		}
		// Virtual field
		zero := m.GetZero()
		_ = zero
		if _zeroErr := writeBuffer.WriteVirtual(ctx, "zero", m.GetZero()); _zeroErr != nil {
			return errors.Wrap(_zeroErr, "Error serializing 'zero' field")
		}

		if err := WriteOptionalField[BACnetApplicationTagUnsignedInteger](ctx, "numberOfDataElements", GetRef(m.GetNumberOfDataElements()), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'numberOfDataElements' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "executionDelay", m.GetExecutionDelay(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'executionDelay' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataExecutionDelay"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataExecutionDelay")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataExecutionDelay) IsBACnetConstructedDataExecutionDelay() {}

func (m *_BACnetConstructedDataExecutionDelay) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataExecutionDelay) deepCopy() *_BACnetConstructedDataExecutionDelay {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataExecutionDelayCopy := &_BACnetConstructedDataExecutionDelay{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagUnsignedInteger](m.NumberOfDataElements),
		utils.DeepCopySlice[BACnetApplicationTagUnsignedInteger, BACnetApplicationTagUnsignedInteger](m.ExecutionDelay),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataExecutionDelayCopy
}

func (m *_BACnetConstructedDataExecutionDelay) String() string {
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
