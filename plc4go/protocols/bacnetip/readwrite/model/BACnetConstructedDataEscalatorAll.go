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

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataEscalatorAll is the corresponding interface of BACnetConstructedDataEscalatorAll
type BACnetConstructedDataEscalatorAll interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// IsBACnetConstructedDataEscalatorAll is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataEscalatorAll()
	// CreateBuilder creates a BACnetConstructedDataEscalatorAllBuilder
	CreateBACnetConstructedDataEscalatorAllBuilder() BACnetConstructedDataEscalatorAllBuilder
}

// _BACnetConstructedDataEscalatorAll is the data-structure of this message
type _BACnetConstructedDataEscalatorAll struct {
	BACnetConstructedDataContract
}

var _ BACnetConstructedDataEscalatorAll = (*_BACnetConstructedDataEscalatorAll)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataEscalatorAll)(nil)

// NewBACnetConstructedDataEscalatorAll factory function for _BACnetConstructedDataEscalatorAll
func NewBACnetConstructedDataEscalatorAll(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataEscalatorAll {
	_result := &_BACnetConstructedDataEscalatorAll{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataEscalatorAllBuilder is a builder for BACnetConstructedDataEscalatorAll
type BACnetConstructedDataEscalatorAllBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() BACnetConstructedDataEscalatorAllBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataEscalatorAll or returns an error if something is wrong
	Build() (BACnetConstructedDataEscalatorAll, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataEscalatorAll
}

// NewBACnetConstructedDataEscalatorAllBuilder() creates a BACnetConstructedDataEscalatorAllBuilder
func NewBACnetConstructedDataEscalatorAllBuilder() BACnetConstructedDataEscalatorAllBuilder {
	return &_BACnetConstructedDataEscalatorAllBuilder{_BACnetConstructedDataEscalatorAll: new(_BACnetConstructedDataEscalatorAll)}
}

type _BACnetConstructedDataEscalatorAllBuilder struct {
	*_BACnetConstructedDataEscalatorAll

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataEscalatorAllBuilder) = (*_BACnetConstructedDataEscalatorAllBuilder)(nil)

func (b *_BACnetConstructedDataEscalatorAllBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataEscalatorAll
}

func (b *_BACnetConstructedDataEscalatorAllBuilder) WithMandatoryFields() BACnetConstructedDataEscalatorAllBuilder {
	return b
}

func (b *_BACnetConstructedDataEscalatorAllBuilder) Build() (BACnetConstructedDataEscalatorAll, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataEscalatorAll.deepCopy(), nil
}

func (b *_BACnetConstructedDataEscalatorAllBuilder) MustBuild() BACnetConstructedDataEscalatorAll {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataEscalatorAllBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataEscalatorAllBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataEscalatorAllBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataEscalatorAllBuilder().(*_BACnetConstructedDataEscalatorAllBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataEscalatorAllBuilder creates a BACnetConstructedDataEscalatorAllBuilder
func (b *_BACnetConstructedDataEscalatorAll) CreateBACnetConstructedDataEscalatorAllBuilder() BACnetConstructedDataEscalatorAllBuilder {
	if b == nil {
		return NewBACnetConstructedDataEscalatorAllBuilder()
	}
	return &_BACnetConstructedDataEscalatorAllBuilder{_BACnetConstructedDataEscalatorAll: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataEscalatorAll) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_ESCALATOR
}

func (m *_BACnetConstructedDataEscalatorAll) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ALL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataEscalatorAll) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataEscalatorAll(structType any) BACnetConstructedDataEscalatorAll {
	if casted, ok := structType.(BACnetConstructedDataEscalatorAll); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataEscalatorAll); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataEscalatorAll) GetTypeName() string {
	return "BACnetConstructedDataEscalatorAll"
}

func (m *_BACnetConstructedDataEscalatorAll) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_BACnetConstructedDataEscalatorAll) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataEscalatorAll) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataEscalatorAll BACnetConstructedDataEscalatorAll, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataEscalatorAll"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataEscalatorAll")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(bool((1) == (2))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "All should never occur in context of constructed data. If it does please report"})
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataEscalatorAll"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataEscalatorAll")
	}

	return m, nil
}

func (m *_BACnetConstructedDataEscalatorAll) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataEscalatorAll) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataEscalatorAll"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataEscalatorAll")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataEscalatorAll"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataEscalatorAll")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataEscalatorAll) IsBACnetConstructedDataEscalatorAll() {}

func (m *_BACnetConstructedDataEscalatorAll) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataEscalatorAll) deepCopy() *_BACnetConstructedDataEscalatorAll {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataEscalatorAllCopy := &_BACnetConstructedDataEscalatorAll{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataEscalatorAllCopy
}

func (m *_BACnetConstructedDataEscalatorAll) String() string {
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
