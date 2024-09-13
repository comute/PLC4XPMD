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

// BACnetCalendarEntryWeekNDay is the corresponding interface of BACnetCalendarEntryWeekNDay
type BACnetCalendarEntryWeekNDay interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetCalendarEntry
	// GetWeekNDay returns WeekNDay (property field)
	GetWeekNDay() BACnetWeekNDayTagged
	// IsBACnetCalendarEntryWeekNDay is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetCalendarEntryWeekNDay()
}

// _BACnetCalendarEntryWeekNDay is the data-structure of this message
type _BACnetCalendarEntryWeekNDay struct {
	BACnetCalendarEntryContract
	WeekNDay BACnetWeekNDayTagged
}

var _ BACnetCalendarEntryWeekNDay = (*_BACnetCalendarEntryWeekNDay)(nil)
var _ BACnetCalendarEntryRequirements = (*_BACnetCalendarEntryWeekNDay)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetCalendarEntryWeekNDay) GetParent() BACnetCalendarEntryContract {
	return m.BACnetCalendarEntryContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetCalendarEntryWeekNDay) GetWeekNDay() BACnetWeekNDayTagged {
	return m.WeekNDay
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetCalendarEntryWeekNDay factory function for _BACnetCalendarEntryWeekNDay
func NewBACnetCalendarEntryWeekNDay(weekNDay BACnetWeekNDayTagged, peekedTagHeader BACnetTagHeader) *_BACnetCalendarEntryWeekNDay {
	if weekNDay == nil {
		panic("weekNDay of type BACnetWeekNDayTagged for BACnetCalendarEntryWeekNDay must not be nil")
	}
	_result := &_BACnetCalendarEntryWeekNDay{
		BACnetCalendarEntryContract: NewBACnetCalendarEntry(peekedTagHeader),
		WeekNDay:                    weekNDay,
	}
	_result.BACnetCalendarEntryContract.(*_BACnetCalendarEntry)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetCalendarEntryWeekNDay(structType any) BACnetCalendarEntryWeekNDay {
	if casted, ok := structType.(BACnetCalendarEntryWeekNDay); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetCalendarEntryWeekNDay); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetCalendarEntryWeekNDay) GetTypeName() string {
	return "BACnetCalendarEntryWeekNDay"
}

func (m *_BACnetCalendarEntryWeekNDay) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetCalendarEntryContract.(*_BACnetCalendarEntry).getLengthInBits(ctx))

	// Simple field (weekNDay)
	lengthInBits += m.WeekNDay.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetCalendarEntryWeekNDay) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetCalendarEntryWeekNDay) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetCalendarEntry) (__bACnetCalendarEntryWeekNDay BACnetCalendarEntryWeekNDay, err error) {
	m.BACnetCalendarEntryContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetCalendarEntryWeekNDay"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetCalendarEntryWeekNDay")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	weekNDay, err := ReadSimpleField[BACnetWeekNDayTagged](ctx, "weekNDay", ReadComplex[BACnetWeekNDayTagged](BACnetWeekNDayTaggedParseWithBufferProducer((uint8)(uint8(2)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'weekNDay' field"))
	}
	m.WeekNDay = weekNDay

	if closeErr := readBuffer.CloseContext("BACnetCalendarEntryWeekNDay"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetCalendarEntryWeekNDay")
	}

	return m, nil
}

func (m *_BACnetCalendarEntryWeekNDay) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetCalendarEntryWeekNDay) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetCalendarEntryWeekNDay"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetCalendarEntryWeekNDay")
		}

		if err := WriteSimpleField[BACnetWeekNDayTagged](ctx, "weekNDay", m.GetWeekNDay(), WriteComplex[BACnetWeekNDayTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'weekNDay' field")
		}

		if popErr := writeBuffer.PopContext("BACnetCalendarEntryWeekNDay"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetCalendarEntryWeekNDay")
		}
		return nil
	}
	return m.BACnetCalendarEntryContract.(*_BACnetCalendarEntry).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetCalendarEntryWeekNDay) IsBACnetCalendarEntryWeekNDay() {}

func (m *_BACnetCalendarEntryWeekNDay) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
