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

// BACnetConstructedDataWindowSamples is the corresponding interface of BACnetConstructedDataWindowSamples
type BACnetConstructedDataWindowSamples interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetWindowSamples returns WindowSamples (property field)
	GetWindowSamples() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetConstructedDataWindowSamples is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataWindowSamples()
	// CreateBuilder creates a BACnetConstructedDataWindowSamplesBuilder
	CreateBACnetConstructedDataWindowSamplesBuilder() BACnetConstructedDataWindowSamplesBuilder
}

// _BACnetConstructedDataWindowSamples is the data-structure of this message
type _BACnetConstructedDataWindowSamples struct {
	BACnetConstructedDataContract
	WindowSamples BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataWindowSamples = (*_BACnetConstructedDataWindowSamples)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataWindowSamples)(nil)

// NewBACnetConstructedDataWindowSamples factory function for _BACnetConstructedDataWindowSamples
func NewBACnetConstructedDataWindowSamples(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, windowSamples BACnetApplicationTagUnsignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataWindowSamples {
	if windowSamples == nil {
		panic("windowSamples of type BACnetApplicationTagUnsignedInteger for BACnetConstructedDataWindowSamples must not be nil")
	}
	_result := &_BACnetConstructedDataWindowSamples{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		WindowSamples:                 windowSamples,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataWindowSamplesBuilder is a builder for BACnetConstructedDataWindowSamples
type BACnetConstructedDataWindowSamplesBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(windowSamples BACnetApplicationTagUnsignedInteger) BACnetConstructedDataWindowSamplesBuilder
	// WithWindowSamples adds WindowSamples (property field)
	WithWindowSamples(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataWindowSamplesBuilder
	// WithWindowSamplesBuilder adds WindowSamples (property field) which is build by the builder
	WithWindowSamplesBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataWindowSamplesBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataWindowSamples or returns an error if something is wrong
	Build() (BACnetConstructedDataWindowSamples, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataWindowSamples
}

// NewBACnetConstructedDataWindowSamplesBuilder() creates a BACnetConstructedDataWindowSamplesBuilder
func NewBACnetConstructedDataWindowSamplesBuilder() BACnetConstructedDataWindowSamplesBuilder {
	return &_BACnetConstructedDataWindowSamplesBuilder{_BACnetConstructedDataWindowSamples: new(_BACnetConstructedDataWindowSamples)}
}

type _BACnetConstructedDataWindowSamplesBuilder struct {
	*_BACnetConstructedDataWindowSamples

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataWindowSamplesBuilder) = (*_BACnetConstructedDataWindowSamplesBuilder)(nil)

func (b *_BACnetConstructedDataWindowSamplesBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataWindowSamples
}

func (b *_BACnetConstructedDataWindowSamplesBuilder) WithMandatoryFields(windowSamples BACnetApplicationTagUnsignedInteger) BACnetConstructedDataWindowSamplesBuilder {
	return b.WithWindowSamples(windowSamples)
}

func (b *_BACnetConstructedDataWindowSamplesBuilder) WithWindowSamples(windowSamples BACnetApplicationTagUnsignedInteger) BACnetConstructedDataWindowSamplesBuilder {
	b.WindowSamples = windowSamples
	return b
}

func (b *_BACnetConstructedDataWindowSamplesBuilder) WithWindowSamplesBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataWindowSamplesBuilder {
	builder := builderSupplier(b.WindowSamples.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	b.WindowSamples, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataWindowSamplesBuilder) Build() (BACnetConstructedDataWindowSamples, error) {
	if b.WindowSamples == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'windowSamples' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataWindowSamples.deepCopy(), nil
}

func (b *_BACnetConstructedDataWindowSamplesBuilder) MustBuild() BACnetConstructedDataWindowSamples {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataWindowSamplesBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataWindowSamplesBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataWindowSamplesBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataWindowSamplesBuilder().(*_BACnetConstructedDataWindowSamplesBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataWindowSamplesBuilder creates a BACnetConstructedDataWindowSamplesBuilder
func (b *_BACnetConstructedDataWindowSamples) CreateBACnetConstructedDataWindowSamplesBuilder() BACnetConstructedDataWindowSamplesBuilder {
	if b == nil {
		return NewBACnetConstructedDataWindowSamplesBuilder()
	}
	return &_BACnetConstructedDataWindowSamplesBuilder{_BACnetConstructedDataWindowSamples: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataWindowSamples) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataWindowSamples) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_WINDOW_SAMPLES
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataWindowSamples) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataWindowSamples) GetWindowSamples() BACnetApplicationTagUnsignedInteger {
	return m.WindowSamples
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataWindowSamples) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetWindowSamples())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataWindowSamples(structType any) BACnetConstructedDataWindowSamples {
	if casted, ok := structType.(BACnetConstructedDataWindowSamples); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataWindowSamples); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataWindowSamples) GetTypeName() string {
	return "BACnetConstructedDataWindowSamples"
}

func (m *_BACnetConstructedDataWindowSamples) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (windowSamples)
	lengthInBits += m.WindowSamples.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataWindowSamples) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataWindowSamples) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataWindowSamples BACnetConstructedDataWindowSamples, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataWindowSamples"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataWindowSamples")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	windowSamples, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "windowSamples", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'windowSamples' field"))
	}
	m.WindowSamples = windowSamples

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), windowSamples)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataWindowSamples"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataWindowSamples")
	}

	return m, nil
}

func (m *_BACnetConstructedDataWindowSamples) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataWindowSamples) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataWindowSamples"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataWindowSamples")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "windowSamples", m.GetWindowSamples(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'windowSamples' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataWindowSamples"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataWindowSamples")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataWindowSamples) IsBACnetConstructedDataWindowSamples() {}

func (m *_BACnetConstructedDataWindowSamples) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataWindowSamples) deepCopy() *_BACnetConstructedDataWindowSamples {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataWindowSamplesCopy := &_BACnetConstructedDataWindowSamples{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagUnsignedInteger](m.WindowSamples),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataWindowSamplesCopy
}

func (m *_BACnetConstructedDataWindowSamples) String() string {
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
