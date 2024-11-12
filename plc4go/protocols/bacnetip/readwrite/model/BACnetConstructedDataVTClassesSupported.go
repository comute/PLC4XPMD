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

// BACnetConstructedDataVTClassesSupported is the corresponding interface of BACnetConstructedDataVTClassesSupported
type BACnetConstructedDataVTClassesSupported interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetVtClassesSupported returns VtClassesSupported (property field)
	GetVtClassesSupported() []BACnetVTClassTagged
	// IsBACnetConstructedDataVTClassesSupported is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataVTClassesSupported()
	// CreateBuilder creates a BACnetConstructedDataVTClassesSupportedBuilder
	CreateBACnetConstructedDataVTClassesSupportedBuilder() BACnetConstructedDataVTClassesSupportedBuilder
}

// _BACnetConstructedDataVTClassesSupported is the data-structure of this message
type _BACnetConstructedDataVTClassesSupported struct {
	BACnetConstructedDataContract
	VtClassesSupported []BACnetVTClassTagged
}

var _ BACnetConstructedDataVTClassesSupported = (*_BACnetConstructedDataVTClassesSupported)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataVTClassesSupported)(nil)

// NewBACnetConstructedDataVTClassesSupported factory function for _BACnetConstructedDataVTClassesSupported
func NewBACnetConstructedDataVTClassesSupported(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, vtClassesSupported []BACnetVTClassTagged, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataVTClassesSupported {
	_result := &_BACnetConstructedDataVTClassesSupported{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		VtClassesSupported:            vtClassesSupported,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataVTClassesSupportedBuilder is a builder for BACnetConstructedDataVTClassesSupported
type BACnetConstructedDataVTClassesSupportedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(vtClassesSupported []BACnetVTClassTagged) BACnetConstructedDataVTClassesSupportedBuilder
	// WithVtClassesSupported adds VtClassesSupported (property field)
	WithVtClassesSupported(...BACnetVTClassTagged) BACnetConstructedDataVTClassesSupportedBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataVTClassesSupported or returns an error if something is wrong
	Build() (BACnetConstructedDataVTClassesSupported, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataVTClassesSupported
}

// NewBACnetConstructedDataVTClassesSupportedBuilder() creates a BACnetConstructedDataVTClassesSupportedBuilder
func NewBACnetConstructedDataVTClassesSupportedBuilder() BACnetConstructedDataVTClassesSupportedBuilder {
	return &_BACnetConstructedDataVTClassesSupportedBuilder{_BACnetConstructedDataVTClassesSupported: new(_BACnetConstructedDataVTClassesSupported)}
}

type _BACnetConstructedDataVTClassesSupportedBuilder struct {
	*_BACnetConstructedDataVTClassesSupported

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataVTClassesSupportedBuilder) = (*_BACnetConstructedDataVTClassesSupportedBuilder)(nil)

func (b *_BACnetConstructedDataVTClassesSupportedBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataVTClassesSupported
}

func (b *_BACnetConstructedDataVTClassesSupportedBuilder) WithMandatoryFields(vtClassesSupported []BACnetVTClassTagged) BACnetConstructedDataVTClassesSupportedBuilder {
	return b.WithVtClassesSupported(vtClassesSupported...)
}

func (b *_BACnetConstructedDataVTClassesSupportedBuilder) WithVtClassesSupported(vtClassesSupported ...BACnetVTClassTagged) BACnetConstructedDataVTClassesSupportedBuilder {
	b.VtClassesSupported = vtClassesSupported
	return b
}

func (b *_BACnetConstructedDataVTClassesSupportedBuilder) Build() (BACnetConstructedDataVTClassesSupported, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataVTClassesSupported.deepCopy(), nil
}

func (b *_BACnetConstructedDataVTClassesSupportedBuilder) MustBuild() BACnetConstructedDataVTClassesSupported {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataVTClassesSupportedBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataVTClassesSupportedBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataVTClassesSupportedBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataVTClassesSupportedBuilder().(*_BACnetConstructedDataVTClassesSupportedBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataVTClassesSupportedBuilder creates a BACnetConstructedDataVTClassesSupportedBuilder
func (b *_BACnetConstructedDataVTClassesSupported) CreateBACnetConstructedDataVTClassesSupportedBuilder() BACnetConstructedDataVTClassesSupportedBuilder {
	if b == nil {
		return NewBACnetConstructedDataVTClassesSupportedBuilder()
	}
	return &_BACnetConstructedDataVTClassesSupportedBuilder{_BACnetConstructedDataVTClassesSupported: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataVTClassesSupported) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataVTClassesSupported) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_VT_CLASSES_SUPPORTED
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataVTClassesSupported) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataVTClassesSupported) GetVtClassesSupported() []BACnetVTClassTagged {
	return m.VtClassesSupported
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataVTClassesSupported(structType any) BACnetConstructedDataVTClassesSupported {
	if casted, ok := structType.(BACnetConstructedDataVTClassesSupported); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataVTClassesSupported); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataVTClassesSupported) GetTypeName() string {
	return "BACnetConstructedDataVTClassesSupported"
}

func (m *_BACnetConstructedDataVTClassesSupported) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Array field
	if len(m.VtClassesSupported) > 0 {
		for _, element := range m.VtClassesSupported {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataVTClassesSupported) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataVTClassesSupported) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataVTClassesSupported BACnetConstructedDataVTClassesSupported, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataVTClassesSupported"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataVTClassesSupported")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	vtClassesSupported, err := ReadTerminatedArrayField[BACnetVTClassTagged](ctx, "vtClassesSupported", ReadComplex[BACnetVTClassTagged](BACnetVTClassTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'vtClassesSupported' field"))
	}
	m.VtClassesSupported = vtClassesSupported

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataVTClassesSupported"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataVTClassesSupported")
	}

	return m, nil
}

func (m *_BACnetConstructedDataVTClassesSupported) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataVTClassesSupported) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataVTClassesSupported"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataVTClassesSupported")
		}

		if err := WriteComplexTypeArrayField(ctx, "vtClassesSupported", m.GetVtClassesSupported(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'vtClassesSupported' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataVTClassesSupported"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataVTClassesSupported")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataVTClassesSupported) IsBACnetConstructedDataVTClassesSupported() {}

func (m *_BACnetConstructedDataVTClassesSupported) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataVTClassesSupported) deepCopy() *_BACnetConstructedDataVTClassesSupported {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataVTClassesSupportedCopy := &_BACnetConstructedDataVTClassesSupported{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopySlice[BACnetVTClassTagged, BACnetVTClassTagged](m.VtClassesSupported),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataVTClassesSupportedCopy
}

func (m *_BACnetConstructedDataVTClassesSupported) String() string {
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
