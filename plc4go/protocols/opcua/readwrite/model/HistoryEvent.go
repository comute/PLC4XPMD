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

// HistoryEvent is the corresponding interface of HistoryEvent
type HistoryEvent interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetEvents returns Events (property field)
	GetEvents() []HistoryEventFieldList
	// IsHistoryEvent is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsHistoryEvent()
	// CreateBuilder creates a HistoryEventBuilder
	CreateHistoryEventBuilder() HistoryEventBuilder
}

// _HistoryEvent is the data-structure of this message
type _HistoryEvent struct {
	ExtensionObjectDefinitionContract
	Events []HistoryEventFieldList
}

var _ HistoryEvent = (*_HistoryEvent)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_HistoryEvent)(nil)

// NewHistoryEvent factory function for _HistoryEvent
func NewHistoryEvent(events []HistoryEventFieldList) *_HistoryEvent {
	_result := &_HistoryEvent{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		Events:                            events,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// HistoryEventBuilder is a builder for HistoryEvent
type HistoryEventBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(events []HistoryEventFieldList) HistoryEventBuilder
	// WithEvents adds Events (property field)
	WithEvents(...HistoryEventFieldList) HistoryEventBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ExtensionObjectDefinitionBuilder
	// Build builds the HistoryEvent or returns an error if something is wrong
	Build() (HistoryEvent, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() HistoryEvent
}

// NewHistoryEventBuilder() creates a HistoryEventBuilder
func NewHistoryEventBuilder() HistoryEventBuilder {
	return &_HistoryEventBuilder{_HistoryEvent: new(_HistoryEvent)}
}

type _HistoryEventBuilder struct {
	*_HistoryEvent

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (HistoryEventBuilder) = (*_HistoryEventBuilder)(nil)

func (b *_HistoryEventBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
	contract.(*_ExtensionObjectDefinition)._SubType = b._HistoryEvent
}

func (b *_HistoryEventBuilder) WithMandatoryFields(events []HistoryEventFieldList) HistoryEventBuilder {
	return b.WithEvents(events...)
}

func (b *_HistoryEventBuilder) WithEvents(events ...HistoryEventFieldList) HistoryEventBuilder {
	b.Events = events
	return b
}

func (b *_HistoryEventBuilder) Build() (HistoryEvent, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._HistoryEvent.deepCopy(), nil
}

func (b *_HistoryEventBuilder) MustBuild() HistoryEvent {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_HistoryEventBuilder) Done() ExtensionObjectDefinitionBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewExtensionObjectDefinitionBuilder().(*_ExtensionObjectDefinitionBuilder)
	}
	return b.parentBuilder
}

func (b *_HistoryEventBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_HistoryEventBuilder) DeepCopy() any {
	_copy := b.CreateHistoryEventBuilder().(*_HistoryEventBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateHistoryEventBuilder creates a HistoryEventBuilder
func (b *_HistoryEvent) CreateHistoryEventBuilder() HistoryEventBuilder {
	if b == nil {
		return NewHistoryEventBuilder()
	}
	return &_HistoryEventBuilder{_HistoryEvent: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_HistoryEvent) GetExtensionId() int32 {
	return int32(661)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_HistoryEvent) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_HistoryEvent) GetEvents() []HistoryEventFieldList {
	return m.Events
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastHistoryEvent(structType any) HistoryEvent {
	if casted, ok := structType.(HistoryEvent); ok {
		return casted
	}
	if casted, ok := structType.(*HistoryEvent); ok {
		return *casted
	}
	return nil
}

func (m *_HistoryEvent) GetTypeName() string {
	return "HistoryEvent"
}

func (m *_HistoryEvent) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Implicit Field (noOfEvents)
	lengthInBits += 32

	// Array field
	if len(m.Events) > 0 {
		for _curItem, element := range m.Events {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.Events), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_HistoryEvent) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_HistoryEvent) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__historyEvent HistoryEvent, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("HistoryEvent"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for HistoryEvent")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	noOfEvents, err := ReadImplicitField[int32](ctx, "noOfEvents", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfEvents' field"))
	}
	_ = noOfEvents

	events, err := ReadCountArrayField[HistoryEventFieldList](ctx, "events", ReadComplex[HistoryEventFieldList](ExtensionObjectDefinitionParseWithBufferProducer[HistoryEventFieldList]((int32)(int32(922))), readBuffer), uint64(noOfEvents))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'events' field"))
	}
	m.Events = events

	if closeErr := readBuffer.CloseContext("HistoryEvent"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for HistoryEvent")
	}

	return m, nil
}

func (m *_HistoryEvent) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_HistoryEvent) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("HistoryEvent"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for HistoryEvent")
		}
		noOfEvents := int32(utils.InlineIf(bool((m.GetEvents()) == (nil)), func() any { return int32(-(int32(1))) }, func() any { return int32(int32(len(m.GetEvents()))) }).(int32))
		if err := WriteImplicitField(ctx, "noOfEvents", noOfEvents, WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfEvents' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "events", m.GetEvents(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'events' field")
		}

		if popErr := writeBuffer.PopContext("HistoryEvent"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for HistoryEvent")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_HistoryEvent) IsHistoryEvent() {}

func (m *_HistoryEvent) DeepCopy() any {
	return m.deepCopy()
}

func (m *_HistoryEvent) deepCopy() *_HistoryEvent {
	if m == nil {
		return nil
	}
	_HistoryEventCopy := &_HistoryEvent{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		utils.DeepCopySlice[HistoryEventFieldList, HistoryEventFieldList](m.Events),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _HistoryEventCopy
}

func (m *_HistoryEvent) String() string {
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
