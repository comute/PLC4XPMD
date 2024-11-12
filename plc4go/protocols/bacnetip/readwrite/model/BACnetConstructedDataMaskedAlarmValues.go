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

// BACnetConstructedDataMaskedAlarmValues is the corresponding interface of BACnetConstructedDataMaskedAlarmValues
type BACnetConstructedDataMaskedAlarmValues interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetMaskedAlarmValues returns MaskedAlarmValues (property field)
	GetMaskedAlarmValues() []BACnetDoorAlarmStateTagged
	// IsBACnetConstructedDataMaskedAlarmValues is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataMaskedAlarmValues()
	// CreateBuilder creates a BACnetConstructedDataMaskedAlarmValuesBuilder
	CreateBACnetConstructedDataMaskedAlarmValuesBuilder() BACnetConstructedDataMaskedAlarmValuesBuilder
}

// _BACnetConstructedDataMaskedAlarmValues is the data-structure of this message
type _BACnetConstructedDataMaskedAlarmValues struct {
	BACnetConstructedDataContract
	MaskedAlarmValues []BACnetDoorAlarmStateTagged
}

var _ BACnetConstructedDataMaskedAlarmValues = (*_BACnetConstructedDataMaskedAlarmValues)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataMaskedAlarmValues)(nil)

// NewBACnetConstructedDataMaskedAlarmValues factory function for _BACnetConstructedDataMaskedAlarmValues
func NewBACnetConstructedDataMaskedAlarmValues(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, maskedAlarmValues []BACnetDoorAlarmStateTagged, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataMaskedAlarmValues {
	_result := &_BACnetConstructedDataMaskedAlarmValues{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		MaskedAlarmValues:             maskedAlarmValues,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataMaskedAlarmValuesBuilder is a builder for BACnetConstructedDataMaskedAlarmValues
type BACnetConstructedDataMaskedAlarmValuesBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(maskedAlarmValues []BACnetDoorAlarmStateTagged) BACnetConstructedDataMaskedAlarmValuesBuilder
	// WithMaskedAlarmValues adds MaskedAlarmValues (property field)
	WithMaskedAlarmValues(...BACnetDoorAlarmStateTagged) BACnetConstructedDataMaskedAlarmValuesBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataMaskedAlarmValues or returns an error if something is wrong
	Build() (BACnetConstructedDataMaskedAlarmValues, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataMaskedAlarmValues
}

// NewBACnetConstructedDataMaskedAlarmValuesBuilder() creates a BACnetConstructedDataMaskedAlarmValuesBuilder
func NewBACnetConstructedDataMaskedAlarmValuesBuilder() BACnetConstructedDataMaskedAlarmValuesBuilder {
	return &_BACnetConstructedDataMaskedAlarmValuesBuilder{_BACnetConstructedDataMaskedAlarmValues: new(_BACnetConstructedDataMaskedAlarmValues)}
}

type _BACnetConstructedDataMaskedAlarmValuesBuilder struct {
	*_BACnetConstructedDataMaskedAlarmValues

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataMaskedAlarmValuesBuilder) = (*_BACnetConstructedDataMaskedAlarmValuesBuilder)(nil)

func (b *_BACnetConstructedDataMaskedAlarmValuesBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataMaskedAlarmValues
}

func (b *_BACnetConstructedDataMaskedAlarmValuesBuilder) WithMandatoryFields(maskedAlarmValues []BACnetDoorAlarmStateTagged) BACnetConstructedDataMaskedAlarmValuesBuilder {
	return b.WithMaskedAlarmValues(maskedAlarmValues...)
}

func (b *_BACnetConstructedDataMaskedAlarmValuesBuilder) WithMaskedAlarmValues(maskedAlarmValues ...BACnetDoorAlarmStateTagged) BACnetConstructedDataMaskedAlarmValuesBuilder {
	b.MaskedAlarmValues = maskedAlarmValues
	return b
}

func (b *_BACnetConstructedDataMaskedAlarmValuesBuilder) Build() (BACnetConstructedDataMaskedAlarmValues, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataMaskedAlarmValues.deepCopy(), nil
}

func (b *_BACnetConstructedDataMaskedAlarmValuesBuilder) MustBuild() BACnetConstructedDataMaskedAlarmValues {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataMaskedAlarmValuesBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataMaskedAlarmValuesBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataMaskedAlarmValuesBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataMaskedAlarmValuesBuilder().(*_BACnetConstructedDataMaskedAlarmValuesBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataMaskedAlarmValuesBuilder creates a BACnetConstructedDataMaskedAlarmValuesBuilder
func (b *_BACnetConstructedDataMaskedAlarmValues) CreateBACnetConstructedDataMaskedAlarmValuesBuilder() BACnetConstructedDataMaskedAlarmValuesBuilder {
	if b == nil {
		return NewBACnetConstructedDataMaskedAlarmValuesBuilder()
	}
	return &_BACnetConstructedDataMaskedAlarmValuesBuilder{_BACnetConstructedDataMaskedAlarmValues: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataMaskedAlarmValues) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataMaskedAlarmValues) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_MASKED_ALARM_VALUES
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataMaskedAlarmValues) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataMaskedAlarmValues) GetMaskedAlarmValues() []BACnetDoorAlarmStateTagged {
	return m.MaskedAlarmValues
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataMaskedAlarmValues(structType any) BACnetConstructedDataMaskedAlarmValues {
	if casted, ok := structType.(BACnetConstructedDataMaskedAlarmValues); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataMaskedAlarmValues); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataMaskedAlarmValues) GetTypeName() string {
	return "BACnetConstructedDataMaskedAlarmValues"
}

func (m *_BACnetConstructedDataMaskedAlarmValues) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Array field
	if len(m.MaskedAlarmValues) > 0 {
		for _, element := range m.MaskedAlarmValues {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataMaskedAlarmValues) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataMaskedAlarmValues) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataMaskedAlarmValues BACnetConstructedDataMaskedAlarmValues, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataMaskedAlarmValues"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataMaskedAlarmValues")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	maskedAlarmValues, err := ReadTerminatedArrayField[BACnetDoorAlarmStateTagged](ctx, "maskedAlarmValues", ReadComplex[BACnetDoorAlarmStateTagged](BACnetDoorAlarmStateTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'maskedAlarmValues' field"))
	}
	m.MaskedAlarmValues = maskedAlarmValues

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataMaskedAlarmValues"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataMaskedAlarmValues")
	}

	return m, nil
}

func (m *_BACnetConstructedDataMaskedAlarmValues) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataMaskedAlarmValues) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataMaskedAlarmValues"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataMaskedAlarmValues")
		}

		if err := WriteComplexTypeArrayField(ctx, "maskedAlarmValues", m.GetMaskedAlarmValues(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'maskedAlarmValues' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataMaskedAlarmValues"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataMaskedAlarmValues")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataMaskedAlarmValues) IsBACnetConstructedDataMaskedAlarmValues() {}

func (m *_BACnetConstructedDataMaskedAlarmValues) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataMaskedAlarmValues) deepCopy() *_BACnetConstructedDataMaskedAlarmValues {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataMaskedAlarmValuesCopy := &_BACnetConstructedDataMaskedAlarmValues{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopySlice[BACnetDoorAlarmStateTagged, BACnetDoorAlarmStateTagged](m.MaskedAlarmValues),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataMaskedAlarmValuesCopy
}

func (m *_BACnetConstructedDataMaskedAlarmValues) String() string {
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
