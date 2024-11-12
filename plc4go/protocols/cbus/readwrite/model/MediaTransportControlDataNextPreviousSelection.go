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

// MediaTransportControlDataNextPreviousSelection is the corresponding interface of MediaTransportControlDataNextPreviousSelection
type MediaTransportControlDataNextPreviousSelection interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	MediaTransportControlData
	// GetOperation returns Operation (property field)
	GetOperation() byte
	// GetIsSetThePreviousSelection returns IsSetThePreviousSelection (virtual field)
	GetIsSetThePreviousSelection() bool
	// GetIsSetTheNextSelection returns IsSetTheNextSelection (virtual field)
	GetIsSetTheNextSelection() bool
	// IsMediaTransportControlDataNextPreviousSelection is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsMediaTransportControlDataNextPreviousSelection()
	// CreateBuilder creates a MediaTransportControlDataNextPreviousSelectionBuilder
	CreateMediaTransportControlDataNextPreviousSelectionBuilder() MediaTransportControlDataNextPreviousSelectionBuilder
}

// _MediaTransportControlDataNextPreviousSelection is the data-structure of this message
type _MediaTransportControlDataNextPreviousSelection struct {
	MediaTransportControlDataContract
	Operation byte
}

var _ MediaTransportControlDataNextPreviousSelection = (*_MediaTransportControlDataNextPreviousSelection)(nil)
var _ MediaTransportControlDataRequirements = (*_MediaTransportControlDataNextPreviousSelection)(nil)

// NewMediaTransportControlDataNextPreviousSelection factory function for _MediaTransportControlDataNextPreviousSelection
func NewMediaTransportControlDataNextPreviousSelection(commandTypeContainer MediaTransportControlCommandTypeContainer, mediaLinkGroup byte, operation byte) *_MediaTransportControlDataNextPreviousSelection {
	_result := &_MediaTransportControlDataNextPreviousSelection{
		MediaTransportControlDataContract: NewMediaTransportControlData(commandTypeContainer, mediaLinkGroup),
		Operation:                         operation,
	}
	_result.MediaTransportControlDataContract.(*_MediaTransportControlData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// MediaTransportControlDataNextPreviousSelectionBuilder is a builder for MediaTransportControlDataNextPreviousSelection
type MediaTransportControlDataNextPreviousSelectionBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(operation byte) MediaTransportControlDataNextPreviousSelectionBuilder
	// WithOperation adds Operation (property field)
	WithOperation(byte) MediaTransportControlDataNextPreviousSelectionBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() MediaTransportControlDataBuilder
	// Build builds the MediaTransportControlDataNextPreviousSelection or returns an error if something is wrong
	Build() (MediaTransportControlDataNextPreviousSelection, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() MediaTransportControlDataNextPreviousSelection
}

// NewMediaTransportControlDataNextPreviousSelectionBuilder() creates a MediaTransportControlDataNextPreviousSelectionBuilder
func NewMediaTransportControlDataNextPreviousSelectionBuilder() MediaTransportControlDataNextPreviousSelectionBuilder {
	return &_MediaTransportControlDataNextPreviousSelectionBuilder{_MediaTransportControlDataNextPreviousSelection: new(_MediaTransportControlDataNextPreviousSelection)}
}

type _MediaTransportControlDataNextPreviousSelectionBuilder struct {
	*_MediaTransportControlDataNextPreviousSelection

	parentBuilder *_MediaTransportControlDataBuilder

	err *utils.MultiError
}

var _ (MediaTransportControlDataNextPreviousSelectionBuilder) = (*_MediaTransportControlDataNextPreviousSelectionBuilder)(nil)

func (b *_MediaTransportControlDataNextPreviousSelectionBuilder) setParent(contract MediaTransportControlDataContract) {
	b.MediaTransportControlDataContract = contract
	contract.(*_MediaTransportControlData)._SubType = b._MediaTransportControlDataNextPreviousSelection
}

func (b *_MediaTransportControlDataNextPreviousSelectionBuilder) WithMandatoryFields(operation byte) MediaTransportControlDataNextPreviousSelectionBuilder {
	return b.WithOperation(operation)
}

func (b *_MediaTransportControlDataNextPreviousSelectionBuilder) WithOperation(operation byte) MediaTransportControlDataNextPreviousSelectionBuilder {
	b.Operation = operation
	return b
}

func (b *_MediaTransportControlDataNextPreviousSelectionBuilder) Build() (MediaTransportControlDataNextPreviousSelection, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._MediaTransportControlDataNextPreviousSelection.deepCopy(), nil
}

func (b *_MediaTransportControlDataNextPreviousSelectionBuilder) MustBuild() MediaTransportControlDataNextPreviousSelection {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_MediaTransportControlDataNextPreviousSelectionBuilder) Done() MediaTransportControlDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewMediaTransportControlDataBuilder().(*_MediaTransportControlDataBuilder)
	}
	return b.parentBuilder
}

func (b *_MediaTransportControlDataNextPreviousSelectionBuilder) buildForMediaTransportControlData() (MediaTransportControlData, error) {
	return b.Build()
}

func (b *_MediaTransportControlDataNextPreviousSelectionBuilder) DeepCopy() any {
	_copy := b.CreateMediaTransportControlDataNextPreviousSelectionBuilder().(*_MediaTransportControlDataNextPreviousSelectionBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateMediaTransportControlDataNextPreviousSelectionBuilder creates a MediaTransportControlDataNextPreviousSelectionBuilder
func (b *_MediaTransportControlDataNextPreviousSelection) CreateMediaTransportControlDataNextPreviousSelectionBuilder() MediaTransportControlDataNextPreviousSelectionBuilder {
	if b == nil {
		return NewMediaTransportControlDataNextPreviousSelectionBuilder()
	}
	return &_MediaTransportControlDataNextPreviousSelectionBuilder{_MediaTransportControlDataNextPreviousSelection: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_MediaTransportControlDataNextPreviousSelection) GetParent() MediaTransportControlDataContract {
	return m.MediaTransportControlDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_MediaTransportControlDataNextPreviousSelection) GetOperation() byte {
	return m.Operation
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_MediaTransportControlDataNextPreviousSelection) GetIsSetThePreviousSelection() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetOperation()) == (0x00)))
}

func (m *_MediaTransportControlDataNextPreviousSelection) GetIsSetTheNextSelection() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetOperation()) != (0x00)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastMediaTransportControlDataNextPreviousSelection(structType any) MediaTransportControlDataNextPreviousSelection {
	if casted, ok := structType.(MediaTransportControlDataNextPreviousSelection); ok {
		return casted
	}
	if casted, ok := structType.(*MediaTransportControlDataNextPreviousSelection); ok {
		return *casted
	}
	return nil
}

func (m *_MediaTransportControlDataNextPreviousSelection) GetTypeName() string {
	return "MediaTransportControlDataNextPreviousSelection"
}

func (m *_MediaTransportControlDataNextPreviousSelection) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.MediaTransportControlDataContract.(*_MediaTransportControlData).getLengthInBits(ctx))

	// Simple field (operation)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_MediaTransportControlDataNextPreviousSelection) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_MediaTransportControlDataNextPreviousSelection) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_MediaTransportControlData) (__mediaTransportControlDataNextPreviousSelection MediaTransportControlDataNextPreviousSelection, err error) {
	m.MediaTransportControlDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MediaTransportControlDataNextPreviousSelection"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MediaTransportControlDataNextPreviousSelection")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	operation, err := ReadSimpleField(ctx, "operation", ReadByte(readBuffer, 8))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'operation' field"))
	}
	m.Operation = operation

	isSetThePreviousSelection, err := ReadVirtualField[bool](ctx, "isSetThePreviousSelection", (*bool)(nil), bool((operation) == (0x00)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isSetThePreviousSelection' field"))
	}
	_ = isSetThePreviousSelection

	isSetTheNextSelection, err := ReadVirtualField[bool](ctx, "isSetTheNextSelection", (*bool)(nil), bool((operation) != (0x00)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isSetTheNextSelection' field"))
	}
	_ = isSetTheNextSelection

	if closeErr := readBuffer.CloseContext("MediaTransportControlDataNextPreviousSelection"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MediaTransportControlDataNextPreviousSelection")
	}

	return m, nil
}

func (m *_MediaTransportControlDataNextPreviousSelection) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MediaTransportControlDataNextPreviousSelection) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MediaTransportControlDataNextPreviousSelection"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MediaTransportControlDataNextPreviousSelection")
		}

		if err := WriteSimpleField[byte](ctx, "operation", m.GetOperation(), WriteByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'operation' field")
		}
		// Virtual field
		isSetThePreviousSelection := m.GetIsSetThePreviousSelection()
		_ = isSetThePreviousSelection
		if _isSetThePreviousSelectionErr := writeBuffer.WriteVirtual(ctx, "isSetThePreviousSelection", m.GetIsSetThePreviousSelection()); _isSetThePreviousSelectionErr != nil {
			return errors.Wrap(_isSetThePreviousSelectionErr, "Error serializing 'isSetThePreviousSelection' field")
		}
		// Virtual field
		isSetTheNextSelection := m.GetIsSetTheNextSelection()
		_ = isSetTheNextSelection
		if _isSetTheNextSelectionErr := writeBuffer.WriteVirtual(ctx, "isSetTheNextSelection", m.GetIsSetTheNextSelection()); _isSetTheNextSelectionErr != nil {
			return errors.Wrap(_isSetTheNextSelectionErr, "Error serializing 'isSetTheNextSelection' field")
		}

		if popErr := writeBuffer.PopContext("MediaTransportControlDataNextPreviousSelection"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MediaTransportControlDataNextPreviousSelection")
		}
		return nil
	}
	return m.MediaTransportControlDataContract.(*_MediaTransportControlData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_MediaTransportControlDataNextPreviousSelection) IsMediaTransportControlDataNextPreviousSelection() {
}

func (m *_MediaTransportControlDataNextPreviousSelection) DeepCopy() any {
	return m.deepCopy()
}

func (m *_MediaTransportControlDataNextPreviousSelection) deepCopy() *_MediaTransportControlDataNextPreviousSelection {
	if m == nil {
		return nil
	}
	_MediaTransportControlDataNextPreviousSelectionCopy := &_MediaTransportControlDataNextPreviousSelection{
		m.MediaTransportControlDataContract.(*_MediaTransportControlData).deepCopy(),
		m.Operation,
	}
	m.MediaTransportControlDataContract.(*_MediaTransportControlData)._SubType = m
	return _MediaTransportControlDataNextPreviousSelectionCopy
}

func (m *_MediaTransportControlDataNextPreviousSelection) String() string {
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
