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

// EventNotificationList is the corresponding interface of EventNotificationList
type EventNotificationList interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetNoOfEvents returns NoOfEvents (property field)
	GetNoOfEvents() int32
	// GetEvents returns Events (property field)
	GetEvents() []ExtensionObjectDefinition
	// IsEventNotificationList is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsEventNotificationList()
}

// _EventNotificationList is the data-structure of this message
type _EventNotificationList struct {
	ExtensionObjectDefinitionContract
	NoOfEvents int32
	Events     []ExtensionObjectDefinition
}

var _ EventNotificationList = (*_EventNotificationList)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_EventNotificationList)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_EventNotificationList) GetIdentifier() string {
	return "916"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_EventNotificationList) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_EventNotificationList) GetNoOfEvents() int32 {
	return m.NoOfEvents
}

func (m *_EventNotificationList) GetEvents() []ExtensionObjectDefinition {
	return m.Events
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewEventNotificationList factory function for _EventNotificationList
func NewEventNotificationList(noOfEvents int32, events []ExtensionObjectDefinition) *_EventNotificationList {
	_result := &_EventNotificationList{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		NoOfEvents:                        noOfEvents,
		Events:                            events,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastEventNotificationList(structType any) EventNotificationList {
	if casted, ok := structType.(EventNotificationList); ok {
		return casted
	}
	if casted, ok := structType.(*EventNotificationList); ok {
		return *casted
	}
	return nil
}

func (m *_EventNotificationList) GetTypeName() string {
	return "EventNotificationList"
}

func (m *_EventNotificationList) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Implicit Field (notificationLength)
	lengthInBits += 32

	// Simple field (noOfEvents)
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

func (m *_EventNotificationList) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_EventNotificationList) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__eventNotificationList EventNotificationList, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("EventNotificationList"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for EventNotificationList")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	notificationLength, err := ReadImplicitField[int32](ctx, "notificationLength", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'notificationLength' field"))
	}
	_ = notificationLength

	noOfEvents, err := ReadSimpleField(ctx, "noOfEvents", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfEvents' field"))
	}
	m.NoOfEvents = noOfEvents

	events, err := ReadCountArrayField[ExtensionObjectDefinition](ctx, "events", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("919")), readBuffer), uint64(noOfEvents))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'events' field"))
	}
	m.Events = events

	if closeErr := readBuffer.CloseContext("EventNotificationList"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for EventNotificationList")
	}

	return m, nil
}

func (m *_EventNotificationList) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_EventNotificationList) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("EventNotificationList"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for EventNotificationList")
		}
		notificationLength := int32(int32(m.GetLengthInBytes(ctx)))
		if err := WriteImplicitField(ctx, "notificationLength", notificationLength, WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'notificationLength' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfEvents", m.GetNoOfEvents(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfEvents' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "events", m.GetEvents(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'events' field")
		}

		if popErr := writeBuffer.PopContext("EventNotificationList"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for EventNotificationList")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_EventNotificationList) IsEventNotificationList() {}

func (m *_EventNotificationList) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
