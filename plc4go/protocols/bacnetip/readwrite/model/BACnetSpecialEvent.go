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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetSpecialEvent is the corresponding interface of BACnetSpecialEvent
type BACnetSpecialEvent interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetPeriod returns Period (property field)
	GetPeriod() BACnetSpecialEventPeriod
	// GetListOfTimeValues returns ListOfTimeValues (property field)
	GetListOfTimeValues() BACnetSpecialEventListOfTimeValues
	// GetEventPriority returns EventPriority (property field)
	GetEventPriority() BACnetContextTagUnsignedInteger
}

// BACnetSpecialEventExactly can be used when we want exactly this type and not a type which fulfills BACnetSpecialEvent.
// This is useful for switch cases.
type BACnetSpecialEventExactly interface {
	BACnetSpecialEvent
	isBACnetSpecialEvent() bool
}

// _BACnetSpecialEvent is the data-structure of this message
type _BACnetSpecialEvent struct {
	Period           BACnetSpecialEventPeriod
	ListOfTimeValues BACnetSpecialEventListOfTimeValues
	EventPriority    BACnetContextTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetSpecialEvent) GetPeriod() BACnetSpecialEventPeriod {
	return m.Period
}

func (m *_BACnetSpecialEvent) GetListOfTimeValues() BACnetSpecialEventListOfTimeValues {
	return m.ListOfTimeValues
}

func (m *_BACnetSpecialEvent) GetEventPriority() BACnetContextTagUnsignedInteger {
	return m.EventPriority
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetSpecialEvent factory function for _BACnetSpecialEvent
func NewBACnetSpecialEvent(period BACnetSpecialEventPeriod, listOfTimeValues BACnetSpecialEventListOfTimeValues, eventPriority BACnetContextTagUnsignedInteger) *_BACnetSpecialEvent {
	return &_BACnetSpecialEvent{Period: period, ListOfTimeValues: listOfTimeValues, EventPriority: eventPriority}
}

// Deprecated: use the interface for direct cast
func CastBACnetSpecialEvent(structType interface{}) BACnetSpecialEvent {
	if casted, ok := structType.(BACnetSpecialEvent); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetSpecialEvent); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetSpecialEvent) GetTypeName() string {
	return "BACnetSpecialEvent"
}

func (m *_BACnetSpecialEvent) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (period)
	lengthInBits += m.Period.GetLengthInBits(ctx)

	// Simple field (listOfTimeValues)
	lengthInBits += m.ListOfTimeValues.GetLengthInBits(ctx)

	// Simple field (eventPriority)
	lengthInBits += m.EventPriority.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetSpecialEvent) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetSpecialEventParse(theBytes []byte) (BACnetSpecialEvent, error) {
	return BACnetSpecialEventParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func BACnetSpecialEventParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetSpecialEvent, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetSpecialEvent"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetSpecialEvent")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (period)
	if pullErr := readBuffer.PullContext("period"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for period")
	}
	_period, _periodErr := BACnetSpecialEventPeriodParseWithBuffer(ctx, readBuffer)
	if _periodErr != nil {
		return nil, errors.Wrap(_periodErr, "Error parsing 'period' field of BACnetSpecialEvent")
	}
	period := _period.(BACnetSpecialEventPeriod)
	if closeErr := readBuffer.CloseContext("period"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for period")
	}

	// Simple Field (listOfTimeValues)
	if pullErr := readBuffer.PullContext("listOfTimeValues"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for listOfTimeValues")
	}
	_listOfTimeValues, _listOfTimeValuesErr := BACnetSpecialEventListOfTimeValuesParseWithBuffer(ctx, readBuffer, uint8(uint8(2)))
	if _listOfTimeValuesErr != nil {
		return nil, errors.Wrap(_listOfTimeValuesErr, "Error parsing 'listOfTimeValues' field of BACnetSpecialEvent")
	}
	listOfTimeValues := _listOfTimeValues.(BACnetSpecialEventListOfTimeValues)
	if closeErr := readBuffer.CloseContext("listOfTimeValues"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for listOfTimeValues")
	}

	// Simple Field (eventPriority)
	if pullErr := readBuffer.PullContext("eventPriority"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for eventPriority")
	}
	_eventPriority, _eventPriorityErr := BACnetContextTagParseWithBuffer(ctx, readBuffer, uint8(uint8(3)), BACnetDataType(BACnetDataType_UNSIGNED_INTEGER))
	if _eventPriorityErr != nil {
		return nil, errors.Wrap(_eventPriorityErr, "Error parsing 'eventPriority' field of BACnetSpecialEvent")
	}
	eventPriority := _eventPriority.(BACnetContextTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("eventPriority"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for eventPriority")
	}

	if closeErr := readBuffer.CloseContext("BACnetSpecialEvent"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetSpecialEvent")
	}

	// Create the instance
	return &_BACnetSpecialEvent{
		Period:           period,
		ListOfTimeValues: listOfTimeValues,
		EventPriority:    eventPriority,
	}, nil
}

func (m *_BACnetSpecialEvent) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetSpecialEvent) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetSpecialEvent"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetSpecialEvent")
	}

	// Simple Field (period)
	if pushErr := writeBuffer.PushContext("period"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for period")
	}
	_periodErr := writeBuffer.WriteSerializable(ctx, m.GetPeriod())
	if popErr := writeBuffer.PopContext("period"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for period")
	}
	if _periodErr != nil {
		return errors.Wrap(_periodErr, "Error serializing 'period' field")
	}

	// Simple Field (listOfTimeValues)
	if pushErr := writeBuffer.PushContext("listOfTimeValues"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for listOfTimeValues")
	}
	_listOfTimeValuesErr := writeBuffer.WriteSerializable(ctx, m.GetListOfTimeValues())
	if popErr := writeBuffer.PopContext("listOfTimeValues"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for listOfTimeValues")
	}
	if _listOfTimeValuesErr != nil {
		return errors.Wrap(_listOfTimeValuesErr, "Error serializing 'listOfTimeValues' field")
	}

	// Simple Field (eventPriority)
	if pushErr := writeBuffer.PushContext("eventPriority"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for eventPriority")
	}
	_eventPriorityErr := writeBuffer.WriteSerializable(ctx, m.GetEventPriority())
	if popErr := writeBuffer.PopContext("eventPriority"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for eventPriority")
	}
	if _eventPriorityErr != nil {
		return errors.Wrap(_eventPriorityErr, "Error serializing 'eventPriority' field")
	}

	if popErr := writeBuffer.PopContext("BACnetSpecialEvent"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetSpecialEvent")
	}
	return nil
}

func (m *_BACnetSpecialEvent) isBACnetSpecialEvent() bool {
	return true
}

func (m *_BACnetSpecialEvent) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
