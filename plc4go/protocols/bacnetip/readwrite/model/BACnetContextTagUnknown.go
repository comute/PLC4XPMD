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

// BACnetContextTagUnknown is the corresponding interface of BACnetContextTagUnknown
type BACnetContextTagUnknown interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetContextTag
	// GetUnknownData returns UnknownData (property field)
	GetUnknownData() []byte
	// IsBACnetContextTagUnknown is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetContextTagUnknown()
	// CreateBuilder creates a BACnetContextTagUnknownBuilder
	CreateBACnetContextTagUnknownBuilder() BACnetContextTagUnknownBuilder
}

// _BACnetContextTagUnknown is the data-structure of this message
type _BACnetContextTagUnknown struct {
	BACnetContextTagContract
	UnknownData []byte

	// Arguments.
	ActualLength uint32
}

var _ BACnetContextTagUnknown = (*_BACnetContextTagUnknown)(nil)
var _ BACnetContextTagRequirements = (*_BACnetContextTagUnknown)(nil)

// NewBACnetContextTagUnknown factory function for _BACnetContextTagUnknown
func NewBACnetContextTagUnknown(header BACnetTagHeader, unknownData []byte, actualLength uint32, tagNumberArgument uint8) *_BACnetContextTagUnknown {
	_result := &_BACnetContextTagUnknown{
		BACnetContextTagContract: NewBACnetContextTag(header, tagNumberArgument),
		UnknownData:              unknownData,
	}
	_result.BACnetContextTagContract.(*_BACnetContextTag)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetContextTagUnknownBuilder is a builder for BACnetContextTagUnknown
type BACnetContextTagUnknownBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(unknownData []byte) BACnetContextTagUnknownBuilder
	// WithUnknownData adds UnknownData (property field)
	WithUnknownData(...byte) BACnetContextTagUnknownBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetContextTagBuilder
	// Build builds the BACnetContextTagUnknown or returns an error if something is wrong
	Build() (BACnetContextTagUnknown, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetContextTagUnknown
}

// NewBACnetContextTagUnknownBuilder() creates a BACnetContextTagUnknownBuilder
func NewBACnetContextTagUnknownBuilder() BACnetContextTagUnknownBuilder {
	return &_BACnetContextTagUnknownBuilder{_BACnetContextTagUnknown: new(_BACnetContextTagUnknown)}
}

type _BACnetContextTagUnknownBuilder struct {
	*_BACnetContextTagUnknown

	parentBuilder *_BACnetContextTagBuilder

	err *utils.MultiError
}

var _ (BACnetContextTagUnknownBuilder) = (*_BACnetContextTagUnknownBuilder)(nil)

func (b *_BACnetContextTagUnknownBuilder) setParent(contract BACnetContextTagContract) {
	b.BACnetContextTagContract = contract
	contract.(*_BACnetContextTag)._SubType = b._BACnetContextTagUnknown
}

func (b *_BACnetContextTagUnknownBuilder) WithMandatoryFields(unknownData []byte) BACnetContextTagUnknownBuilder {
	return b.WithUnknownData(unknownData...)
}

func (b *_BACnetContextTagUnknownBuilder) WithUnknownData(unknownData ...byte) BACnetContextTagUnknownBuilder {
	b.UnknownData = unknownData
	return b
}

func (b *_BACnetContextTagUnknownBuilder) Build() (BACnetContextTagUnknown, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetContextTagUnknown.deepCopy(), nil
}

func (b *_BACnetContextTagUnknownBuilder) MustBuild() BACnetContextTagUnknown {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetContextTagUnknownBuilder) Done() BACnetContextTagBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetContextTagBuilder().(*_BACnetContextTagBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetContextTagUnknownBuilder) buildForBACnetContextTag() (BACnetContextTag, error) {
	return b.Build()
}

func (b *_BACnetContextTagUnknownBuilder) DeepCopy() any {
	_copy := b.CreateBACnetContextTagUnknownBuilder().(*_BACnetContextTagUnknownBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetContextTagUnknownBuilder creates a BACnetContextTagUnknownBuilder
func (b *_BACnetContextTagUnknown) CreateBACnetContextTagUnknownBuilder() BACnetContextTagUnknownBuilder {
	if b == nil {
		return NewBACnetContextTagUnknownBuilder()
	}
	return &_BACnetContextTagUnknownBuilder{_BACnetContextTagUnknown: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetContextTagUnknown) GetDataType() BACnetDataType {
	return BACnetDataType_UNKNOWN
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetContextTagUnknown) GetParent() BACnetContextTagContract {
	return m.BACnetContextTagContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetContextTagUnknown) GetUnknownData() []byte {
	return m.UnknownData
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetContextTagUnknown(structType any) BACnetContextTagUnknown {
	if casted, ok := structType.(BACnetContextTagUnknown); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetContextTagUnknown); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetContextTagUnknown) GetTypeName() string {
	return "BACnetContextTagUnknown"
}

func (m *_BACnetContextTagUnknown) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetContextTagContract.(*_BACnetContextTag).getLengthInBits(ctx))

	// Array field
	if len(m.UnknownData) > 0 {
		lengthInBits += 8 * uint16(len(m.UnknownData))
	}

	return lengthInBits
}

func (m *_BACnetContextTagUnknown) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetContextTagUnknown) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetContextTag, actualLength uint32, tagNumberArgument uint8, dataType BACnetDataType) (__bACnetContextTagUnknown BACnetContextTagUnknown, err error) {
	m.BACnetContextTagContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetContextTagUnknown"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetContextTagUnknown")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	unknownData, err := readBuffer.ReadByteArray("unknownData", int(actualLength))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'unknownData' field"))
	}
	m.UnknownData = unknownData

	if closeErr := readBuffer.CloseContext("BACnetContextTagUnknown"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetContextTagUnknown")
	}

	return m, nil
}

func (m *_BACnetContextTagUnknown) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetContextTagUnknown) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetContextTagUnknown"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetContextTagUnknown")
		}

		if err := WriteByteArrayField(ctx, "unknownData", m.GetUnknownData(), WriteByteArray(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'unknownData' field")
		}

		if popErr := writeBuffer.PopContext("BACnetContextTagUnknown"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetContextTagUnknown")
		}
		return nil
	}
	return m.BACnetContextTagContract.(*_BACnetContextTag).serializeParent(ctx, writeBuffer, m, ser)
}

////
// Arguments Getter

func (m *_BACnetContextTagUnknown) GetActualLength() uint32 {
	return m.ActualLength
}

//
////

func (m *_BACnetContextTagUnknown) IsBACnetContextTagUnknown() {}

func (m *_BACnetContextTagUnknown) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetContextTagUnknown) deepCopy() *_BACnetContextTagUnknown {
	if m == nil {
		return nil
	}
	_BACnetContextTagUnknownCopy := &_BACnetContextTagUnknown{
		m.BACnetContextTagContract.(*_BACnetContextTag).deepCopy(),
		utils.DeepCopySlice[byte, byte](m.UnknownData),
		m.ActualLength,
	}
	m.BACnetContextTagContract.(*_BACnetContextTag)._SubType = m
	return _BACnetContextTagUnknownCopy
}

func (m *_BACnetContextTagUnknown) String() string {
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
