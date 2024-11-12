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

// AdsMultiRequestItemWrite is the corresponding interface of AdsMultiRequestItemWrite
type AdsMultiRequestItemWrite interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	AdsMultiRequestItem
	// GetItemIndexGroup returns ItemIndexGroup (property field)
	GetItemIndexGroup() uint32
	// GetItemIndexOffset returns ItemIndexOffset (property field)
	GetItemIndexOffset() uint32
	// GetItemWriteLength returns ItemWriteLength (property field)
	GetItemWriteLength() uint32
	// IsAdsMultiRequestItemWrite is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsAdsMultiRequestItemWrite()
	// CreateBuilder creates a AdsMultiRequestItemWriteBuilder
	CreateAdsMultiRequestItemWriteBuilder() AdsMultiRequestItemWriteBuilder
}

// _AdsMultiRequestItemWrite is the data-structure of this message
type _AdsMultiRequestItemWrite struct {
	AdsMultiRequestItemContract
	ItemIndexGroup  uint32
	ItemIndexOffset uint32
	ItemWriteLength uint32
}

var _ AdsMultiRequestItemWrite = (*_AdsMultiRequestItemWrite)(nil)
var _ AdsMultiRequestItemRequirements = (*_AdsMultiRequestItemWrite)(nil)

// NewAdsMultiRequestItemWrite factory function for _AdsMultiRequestItemWrite
func NewAdsMultiRequestItemWrite(itemIndexGroup uint32, itemIndexOffset uint32, itemWriteLength uint32) *_AdsMultiRequestItemWrite {
	_result := &_AdsMultiRequestItemWrite{
		AdsMultiRequestItemContract: NewAdsMultiRequestItem(),
		ItemIndexGroup:              itemIndexGroup,
		ItemIndexOffset:             itemIndexOffset,
		ItemWriteLength:             itemWriteLength,
	}
	_result.AdsMultiRequestItemContract.(*_AdsMultiRequestItem)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// AdsMultiRequestItemWriteBuilder is a builder for AdsMultiRequestItemWrite
type AdsMultiRequestItemWriteBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(itemIndexGroup uint32, itemIndexOffset uint32, itemWriteLength uint32) AdsMultiRequestItemWriteBuilder
	// WithItemIndexGroup adds ItemIndexGroup (property field)
	WithItemIndexGroup(uint32) AdsMultiRequestItemWriteBuilder
	// WithItemIndexOffset adds ItemIndexOffset (property field)
	WithItemIndexOffset(uint32) AdsMultiRequestItemWriteBuilder
	// WithItemWriteLength adds ItemWriteLength (property field)
	WithItemWriteLength(uint32) AdsMultiRequestItemWriteBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() AdsMultiRequestItemBuilder
	// Build builds the AdsMultiRequestItemWrite or returns an error if something is wrong
	Build() (AdsMultiRequestItemWrite, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() AdsMultiRequestItemWrite
}

// NewAdsMultiRequestItemWriteBuilder() creates a AdsMultiRequestItemWriteBuilder
func NewAdsMultiRequestItemWriteBuilder() AdsMultiRequestItemWriteBuilder {
	return &_AdsMultiRequestItemWriteBuilder{_AdsMultiRequestItemWrite: new(_AdsMultiRequestItemWrite)}
}

type _AdsMultiRequestItemWriteBuilder struct {
	*_AdsMultiRequestItemWrite

	parentBuilder *_AdsMultiRequestItemBuilder

	err *utils.MultiError
}

var _ (AdsMultiRequestItemWriteBuilder) = (*_AdsMultiRequestItemWriteBuilder)(nil)

func (b *_AdsMultiRequestItemWriteBuilder) setParent(contract AdsMultiRequestItemContract) {
	b.AdsMultiRequestItemContract = contract
	contract.(*_AdsMultiRequestItem)._SubType = b._AdsMultiRequestItemWrite
}

func (b *_AdsMultiRequestItemWriteBuilder) WithMandatoryFields(itemIndexGroup uint32, itemIndexOffset uint32, itemWriteLength uint32) AdsMultiRequestItemWriteBuilder {
	return b.WithItemIndexGroup(itemIndexGroup).WithItemIndexOffset(itemIndexOffset).WithItemWriteLength(itemWriteLength)
}

func (b *_AdsMultiRequestItemWriteBuilder) WithItemIndexGroup(itemIndexGroup uint32) AdsMultiRequestItemWriteBuilder {
	b.ItemIndexGroup = itemIndexGroup
	return b
}

func (b *_AdsMultiRequestItemWriteBuilder) WithItemIndexOffset(itemIndexOffset uint32) AdsMultiRequestItemWriteBuilder {
	b.ItemIndexOffset = itemIndexOffset
	return b
}

func (b *_AdsMultiRequestItemWriteBuilder) WithItemWriteLength(itemWriteLength uint32) AdsMultiRequestItemWriteBuilder {
	b.ItemWriteLength = itemWriteLength
	return b
}

func (b *_AdsMultiRequestItemWriteBuilder) Build() (AdsMultiRequestItemWrite, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._AdsMultiRequestItemWrite.deepCopy(), nil
}

func (b *_AdsMultiRequestItemWriteBuilder) MustBuild() AdsMultiRequestItemWrite {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_AdsMultiRequestItemWriteBuilder) Done() AdsMultiRequestItemBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewAdsMultiRequestItemBuilder().(*_AdsMultiRequestItemBuilder)
	}
	return b.parentBuilder
}

func (b *_AdsMultiRequestItemWriteBuilder) buildForAdsMultiRequestItem() (AdsMultiRequestItem, error) {
	return b.Build()
}

func (b *_AdsMultiRequestItemWriteBuilder) DeepCopy() any {
	_copy := b.CreateAdsMultiRequestItemWriteBuilder().(*_AdsMultiRequestItemWriteBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateAdsMultiRequestItemWriteBuilder creates a AdsMultiRequestItemWriteBuilder
func (b *_AdsMultiRequestItemWrite) CreateAdsMultiRequestItemWriteBuilder() AdsMultiRequestItemWriteBuilder {
	if b == nil {
		return NewAdsMultiRequestItemWriteBuilder()
	}
	return &_AdsMultiRequestItemWriteBuilder{_AdsMultiRequestItemWrite: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_AdsMultiRequestItemWrite) GetIndexGroup() uint32 {
	return uint32(61569)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AdsMultiRequestItemWrite) GetParent() AdsMultiRequestItemContract {
	return m.AdsMultiRequestItemContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AdsMultiRequestItemWrite) GetItemIndexGroup() uint32 {
	return m.ItemIndexGroup
}

func (m *_AdsMultiRequestItemWrite) GetItemIndexOffset() uint32 {
	return m.ItemIndexOffset
}

func (m *_AdsMultiRequestItemWrite) GetItemWriteLength() uint32 {
	return m.ItemWriteLength
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastAdsMultiRequestItemWrite(structType any) AdsMultiRequestItemWrite {
	if casted, ok := structType.(AdsMultiRequestItemWrite); ok {
		return casted
	}
	if casted, ok := structType.(*AdsMultiRequestItemWrite); ok {
		return *casted
	}
	return nil
}

func (m *_AdsMultiRequestItemWrite) GetTypeName() string {
	return "AdsMultiRequestItemWrite"
}

func (m *_AdsMultiRequestItemWrite) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.AdsMultiRequestItemContract.(*_AdsMultiRequestItem).getLengthInBits(ctx))

	// Simple field (itemIndexGroup)
	lengthInBits += 32

	// Simple field (itemIndexOffset)
	lengthInBits += 32

	// Simple field (itemWriteLength)
	lengthInBits += 32

	return lengthInBits
}

func (m *_AdsMultiRequestItemWrite) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_AdsMultiRequestItemWrite) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_AdsMultiRequestItem, indexGroup uint32) (__adsMultiRequestItemWrite AdsMultiRequestItemWrite, err error) {
	m.AdsMultiRequestItemContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AdsMultiRequestItemWrite"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AdsMultiRequestItemWrite")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	itemIndexGroup, err := ReadSimpleField(ctx, "itemIndexGroup", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'itemIndexGroup' field"))
	}
	m.ItemIndexGroup = itemIndexGroup

	itemIndexOffset, err := ReadSimpleField(ctx, "itemIndexOffset", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'itemIndexOffset' field"))
	}
	m.ItemIndexOffset = itemIndexOffset

	itemWriteLength, err := ReadSimpleField(ctx, "itemWriteLength", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'itemWriteLength' field"))
	}
	m.ItemWriteLength = itemWriteLength

	if closeErr := readBuffer.CloseContext("AdsMultiRequestItemWrite"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AdsMultiRequestItemWrite")
	}

	return m, nil
}

func (m *_AdsMultiRequestItemWrite) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AdsMultiRequestItemWrite) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AdsMultiRequestItemWrite"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AdsMultiRequestItemWrite")
		}

		if err := WriteSimpleField[uint32](ctx, "itemIndexGroup", m.GetItemIndexGroup(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'itemIndexGroup' field")
		}

		if err := WriteSimpleField[uint32](ctx, "itemIndexOffset", m.GetItemIndexOffset(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'itemIndexOffset' field")
		}

		if err := WriteSimpleField[uint32](ctx, "itemWriteLength", m.GetItemWriteLength(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'itemWriteLength' field")
		}

		if popErr := writeBuffer.PopContext("AdsMultiRequestItemWrite"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AdsMultiRequestItemWrite")
		}
		return nil
	}
	return m.AdsMultiRequestItemContract.(*_AdsMultiRequestItem).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AdsMultiRequestItemWrite) IsAdsMultiRequestItemWrite() {}

func (m *_AdsMultiRequestItemWrite) DeepCopy() any {
	return m.deepCopy()
}

func (m *_AdsMultiRequestItemWrite) deepCopy() *_AdsMultiRequestItemWrite {
	if m == nil {
		return nil
	}
	_AdsMultiRequestItemWriteCopy := &_AdsMultiRequestItemWrite{
		m.AdsMultiRequestItemContract.(*_AdsMultiRequestItem).deepCopy(),
		m.ItemIndexGroup,
		m.ItemIndexOffset,
		m.ItemWriteLength,
	}
	m.AdsMultiRequestItemContract.(*_AdsMultiRequestItem)._SubType = m
	return _AdsMultiRequestItemWriteCopy
}

func (m *_AdsMultiRequestItemWrite) String() string {
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
