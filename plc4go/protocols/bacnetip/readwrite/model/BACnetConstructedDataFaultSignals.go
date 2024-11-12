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

// BACnetConstructedDataFaultSignals is the corresponding interface of BACnetConstructedDataFaultSignals
type BACnetConstructedDataFaultSignals interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetFaultSignals returns FaultSignals (property field)
	GetFaultSignals() []BACnetLiftFaultTagged
	// IsBACnetConstructedDataFaultSignals is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataFaultSignals()
	// CreateBuilder creates a BACnetConstructedDataFaultSignalsBuilder
	CreateBACnetConstructedDataFaultSignalsBuilder() BACnetConstructedDataFaultSignalsBuilder
}

// _BACnetConstructedDataFaultSignals is the data-structure of this message
type _BACnetConstructedDataFaultSignals struct {
	BACnetConstructedDataContract
	FaultSignals []BACnetLiftFaultTagged
}

var _ BACnetConstructedDataFaultSignals = (*_BACnetConstructedDataFaultSignals)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataFaultSignals)(nil)

// NewBACnetConstructedDataFaultSignals factory function for _BACnetConstructedDataFaultSignals
func NewBACnetConstructedDataFaultSignals(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, faultSignals []BACnetLiftFaultTagged, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataFaultSignals {
	_result := &_BACnetConstructedDataFaultSignals{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		FaultSignals:                  faultSignals,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataFaultSignalsBuilder is a builder for BACnetConstructedDataFaultSignals
type BACnetConstructedDataFaultSignalsBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(faultSignals []BACnetLiftFaultTagged) BACnetConstructedDataFaultSignalsBuilder
	// WithFaultSignals adds FaultSignals (property field)
	WithFaultSignals(...BACnetLiftFaultTagged) BACnetConstructedDataFaultSignalsBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataFaultSignals or returns an error if something is wrong
	Build() (BACnetConstructedDataFaultSignals, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataFaultSignals
}

// NewBACnetConstructedDataFaultSignalsBuilder() creates a BACnetConstructedDataFaultSignalsBuilder
func NewBACnetConstructedDataFaultSignalsBuilder() BACnetConstructedDataFaultSignalsBuilder {
	return &_BACnetConstructedDataFaultSignalsBuilder{_BACnetConstructedDataFaultSignals: new(_BACnetConstructedDataFaultSignals)}
}

type _BACnetConstructedDataFaultSignalsBuilder struct {
	*_BACnetConstructedDataFaultSignals

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataFaultSignalsBuilder) = (*_BACnetConstructedDataFaultSignalsBuilder)(nil)

func (b *_BACnetConstructedDataFaultSignalsBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataFaultSignals
}

func (b *_BACnetConstructedDataFaultSignalsBuilder) WithMandatoryFields(faultSignals []BACnetLiftFaultTagged) BACnetConstructedDataFaultSignalsBuilder {
	return b.WithFaultSignals(faultSignals...)
}

func (b *_BACnetConstructedDataFaultSignalsBuilder) WithFaultSignals(faultSignals ...BACnetLiftFaultTagged) BACnetConstructedDataFaultSignalsBuilder {
	b.FaultSignals = faultSignals
	return b
}

func (b *_BACnetConstructedDataFaultSignalsBuilder) Build() (BACnetConstructedDataFaultSignals, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataFaultSignals.deepCopy(), nil
}

func (b *_BACnetConstructedDataFaultSignalsBuilder) MustBuild() BACnetConstructedDataFaultSignals {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataFaultSignalsBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataFaultSignalsBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataFaultSignalsBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataFaultSignalsBuilder().(*_BACnetConstructedDataFaultSignalsBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataFaultSignalsBuilder creates a BACnetConstructedDataFaultSignalsBuilder
func (b *_BACnetConstructedDataFaultSignals) CreateBACnetConstructedDataFaultSignalsBuilder() BACnetConstructedDataFaultSignalsBuilder {
	if b == nil {
		return NewBACnetConstructedDataFaultSignalsBuilder()
	}
	return &_BACnetConstructedDataFaultSignalsBuilder{_BACnetConstructedDataFaultSignals: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataFaultSignals) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataFaultSignals) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_FAULT_SIGNALS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataFaultSignals) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataFaultSignals) GetFaultSignals() []BACnetLiftFaultTagged {
	return m.FaultSignals
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataFaultSignals(structType any) BACnetConstructedDataFaultSignals {
	if casted, ok := structType.(BACnetConstructedDataFaultSignals); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataFaultSignals); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataFaultSignals) GetTypeName() string {
	return "BACnetConstructedDataFaultSignals"
}

func (m *_BACnetConstructedDataFaultSignals) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Array field
	if len(m.FaultSignals) > 0 {
		for _, element := range m.FaultSignals {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataFaultSignals) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataFaultSignals) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataFaultSignals BACnetConstructedDataFaultSignals, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataFaultSignals"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataFaultSignals")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	faultSignals, err := ReadTerminatedArrayField[BACnetLiftFaultTagged](ctx, "faultSignals", ReadComplex[BACnetLiftFaultTagged](BACnetLiftFaultTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'faultSignals' field"))
	}
	m.FaultSignals = faultSignals

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataFaultSignals"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataFaultSignals")
	}

	return m, nil
}

func (m *_BACnetConstructedDataFaultSignals) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataFaultSignals) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataFaultSignals"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataFaultSignals")
		}

		if err := WriteComplexTypeArrayField(ctx, "faultSignals", m.GetFaultSignals(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'faultSignals' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataFaultSignals"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataFaultSignals")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataFaultSignals) IsBACnetConstructedDataFaultSignals() {}

func (m *_BACnetConstructedDataFaultSignals) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataFaultSignals) deepCopy() *_BACnetConstructedDataFaultSignals {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataFaultSignalsCopy := &_BACnetConstructedDataFaultSignals{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopySlice[BACnetLiftFaultTagged, BACnetLiftFaultTagged](m.FaultSignals),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataFaultSignalsCopy
}

func (m *_BACnetConstructedDataFaultSignals) String() string {
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
