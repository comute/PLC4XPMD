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

// BACnetConstructedDataEventParameters is the corresponding interface of BACnetConstructedDataEventParameters
type BACnetConstructedDataEventParameters interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetEventParameter returns EventParameter (property field)
	GetEventParameter() BACnetEventParameter
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetEventParameter
	// IsBACnetConstructedDataEventParameters is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataEventParameters()
	// CreateBuilder creates a BACnetConstructedDataEventParametersBuilder
	CreateBACnetConstructedDataEventParametersBuilder() BACnetConstructedDataEventParametersBuilder
}

// _BACnetConstructedDataEventParameters is the data-structure of this message
type _BACnetConstructedDataEventParameters struct {
	BACnetConstructedDataContract
	EventParameter BACnetEventParameter
}

var _ BACnetConstructedDataEventParameters = (*_BACnetConstructedDataEventParameters)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataEventParameters)(nil)

// NewBACnetConstructedDataEventParameters factory function for _BACnetConstructedDataEventParameters
func NewBACnetConstructedDataEventParameters(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, eventParameter BACnetEventParameter, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataEventParameters {
	if eventParameter == nil {
		panic("eventParameter of type BACnetEventParameter for BACnetConstructedDataEventParameters must not be nil")
	}
	_result := &_BACnetConstructedDataEventParameters{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		EventParameter:                eventParameter,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataEventParametersBuilder is a builder for BACnetConstructedDataEventParameters
type BACnetConstructedDataEventParametersBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(eventParameter BACnetEventParameter) BACnetConstructedDataEventParametersBuilder
	// WithEventParameter adds EventParameter (property field)
	WithEventParameter(BACnetEventParameter) BACnetConstructedDataEventParametersBuilder
	// WithEventParameterBuilder adds EventParameter (property field) which is build by the builder
	WithEventParameterBuilder(func(BACnetEventParameterBuilder) BACnetEventParameterBuilder) BACnetConstructedDataEventParametersBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataEventParameters or returns an error if something is wrong
	Build() (BACnetConstructedDataEventParameters, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataEventParameters
}

// NewBACnetConstructedDataEventParametersBuilder() creates a BACnetConstructedDataEventParametersBuilder
func NewBACnetConstructedDataEventParametersBuilder() BACnetConstructedDataEventParametersBuilder {
	return &_BACnetConstructedDataEventParametersBuilder{_BACnetConstructedDataEventParameters: new(_BACnetConstructedDataEventParameters)}
}

type _BACnetConstructedDataEventParametersBuilder struct {
	*_BACnetConstructedDataEventParameters

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataEventParametersBuilder) = (*_BACnetConstructedDataEventParametersBuilder)(nil)

func (b *_BACnetConstructedDataEventParametersBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataEventParameters
}

func (b *_BACnetConstructedDataEventParametersBuilder) WithMandatoryFields(eventParameter BACnetEventParameter) BACnetConstructedDataEventParametersBuilder {
	return b.WithEventParameter(eventParameter)
}

func (b *_BACnetConstructedDataEventParametersBuilder) WithEventParameter(eventParameter BACnetEventParameter) BACnetConstructedDataEventParametersBuilder {
	b.EventParameter = eventParameter
	return b
}

func (b *_BACnetConstructedDataEventParametersBuilder) WithEventParameterBuilder(builderSupplier func(BACnetEventParameterBuilder) BACnetEventParameterBuilder) BACnetConstructedDataEventParametersBuilder {
	builder := builderSupplier(b.EventParameter.CreateBACnetEventParameterBuilder())
	var err error
	b.EventParameter, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetEventParameterBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataEventParametersBuilder) Build() (BACnetConstructedDataEventParameters, error) {
	if b.EventParameter == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'eventParameter' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataEventParameters.deepCopy(), nil
}

func (b *_BACnetConstructedDataEventParametersBuilder) MustBuild() BACnetConstructedDataEventParameters {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataEventParametersBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataEventParametersBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataEventParametersBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataEventParametersBuilder().(*_BACnetConstructedDataEventParametersBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataEventParametersBuilder creates a BACnetConstructedDataEventParametersBuilder
func (b *_BACnetConstructedDataEventParameters) CreateBACnetConstructedDataEventParametersBuilder() BACnetConstructedDataEventParametersBuilder {
	if b == nil {
		return NewBACnetConstructedDataEventParametersBuilder()
	}
	return &_BACnetConstructedDataEventParametersBuilder{_BACnetConstructedDataEventParameters: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataEventParameters) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataEventParameters) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_EVENT_PARAMETERS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataEventParameters) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataEventParameters) GetEventParameter() BACnetEventParameter {
	return m.EventParameter
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataEventParameters) GetActualValue() BACnetEventParameter {
	ctx := context.Background()
	_ = ctx
	return CastBACnetEventParameter(m.GetEventParameter())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataEventParameters(structType any) BACnetConstructedDataEventParameters {
	if casted, ok := structType.(BACnetConstructedDataEventParameters); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataEventParameters); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataEventParameters) GetTypeName() string {
	return "BACnetConstructedDataEventParameters"
}

func (m *_BACnetConstructedDataEventParameters) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (eventParameter)
	lengthInBits += m.EventParameter.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataEventParameters) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataEventParameters) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataEventParameters BACnetConstructedDataEventParameters, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataEventParameters"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataEventParameters")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	eventParameter, err := ReadSimpleField[BACnetEventParameter](ctx, "eventParameter", ReadComplex[BACnetEventParameter](BACnetEventParameterParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'eventParameter' field"))
	}
	m.EventParameter = eventParameter

	actualValue, err := ReadVirtualField[BACnetEventParameter](ctx, "actualValue", (*BACnetEventParameter)(nil), eventParameter)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataEventParameters"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataEventParameters")
	}

	return m, nil
}

func (m *_BACnetConstructedDataEventParameters) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataEventParameters) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataEventParameters"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataEventParameters")
		}

		if err := WriteSimpleField[BACnetEventParameter](ctx, "eventParameter", m.GetEventParameter(), WriteComplex[BACnetEventParameter](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'eventParameter' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataEventParameters"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataEventParameters")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataEventParameters) IsBACnetConstructedDataEventParameters() {}

func (m *_BACnetConstructedDataEventParameters) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataEventParameters) deepCopy() *_BACnetConstructedDataEventParameters {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataEventParametersCopy := &_BACnetConstructedDataEventParameters{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetEventParameter](m.EventParameter),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataEventParametersCopy
}

func (m *_BACnetConstructedDataEventParameters) String() string {
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
