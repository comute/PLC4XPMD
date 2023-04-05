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

// BACnetTagPayloadDate is the corresponding interface of BACnetTagPayloadDate
type BACnetTagPayloadDate interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetYearMinus1900 returns YearMinus1900 (property field)
	GetYearMinus1900() uint8
	// GetMonth returns Month (property field)
	GetMonth() uint8
	// GetDayOfMonth returns DayOfMonth (property field)
	GetDayOfMonth() uint8
	// GetDayOfWeek returns DayOfWeek (property field)
	GetDayOfWeek() uint8
	// GetWildcard returns Wildcard (virtual field)
	GetWildcard() uint8
	// GetYearIsWildcard returns YearIsWildcard (virtual field)
	GetYearIsWildcard() bool
	// GetYear returns Year (virtual field)
	GetYear() uint16
	// GetMonthIsWildcard returns MonthIsWildcard (virtual field)
	GetMonthIsWildcard() bool
	// GetOddMonthWildcard returns OddMonthWildcard (virtual field)
	GetOddMonthWildcard() bool
	// GetEvenMonthWildcard returns EvenMonthWildcard (virtual field)
	GetEvenMonthWildcard() bool
	// GetDayOfMonthIsWildcard returns DayOfMonthIsWildcard (virtual field)
	GetDayOfMonthIsWildcard() bool
	// GetLastDayOfMonthWildcard returns LastDayOfMonthWildcard (virtual field)
	GetLastDayOfMonthWildcard() bool
	// GetOddDayOfMonthWildcard returns OddDayOfMonthWildcard (virtual field)
	GetOddDayOfMonthWildcard() bool
	// GetEvenDayOfMonthWildcard returns EvenDayOfMonthWildcard (virtual field)
	GetEvenDayOfMonthWildcard() bool
	// GetDayOfWeekIsWildcard returns DayOfWeekIsWildcard (virtual field)
	GetDayOfWeekIsWildcard() bool
}

// BACnetTagPayloadDateExactly can be used when we want exactly this type and not a type which fulfills BACnetTagPayloadDate.
// This is useful for switch cases.
type BACnetTagPayloadDateExactly interface {
	BACnetTagPayloadDate
	isBACnetTagPayloadDate() bool
}

// _BACnetTagPayloadDate is the data-structure of this message
type _BACnetTagPayloadDate struct {
	YearMinus1900 uint8
	Month         uint8
	DayOfMonth    uint8
	DayOfWeek     uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetTagPayloadDate) GetYearMinus1900() uint8 {
	return m.YearMinus1900
}

func (m *_BACnetTagPayloadDate) GetMonth() uint8 {
	return m.Month
}

func (m *_BACnetTagPayloadDate) GetDayOfMonth() uint8 {
	return m.DayOfMonth
}

func (m *_BACnetTagPayloadDate) GetDayOfWeek() uint8 {
	return m.DayOfWeek
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetTagPayloadDate) GetWildcard() uint8 {
	ctx := context.Background()
	_ = ctx
	return uint8(0xFF)
}

func (m *_BACnetTagPayloadDate) GetYearIsWildcard() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetYearMinus1900()) == (m.GetWildcard())))
}

func (m *_BACnetTagPayloadDate) GetYear() uint16 {
	ctx := context.Background()
	_ = ctx
	return uint16(uint16(m.GetYearMinus1900()) + uint16(uint16(1900)))
}

func (m *_BACnetTagPayloadDate) GetMonthIsWildcard() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetMonth()) == (m.GetWildcard())))
}

func (m *_BACnetTagPayloadDate) GetOddMonthWildcard() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetMonth()) == (13)))
}

func (m *_BACnetTagPayloadDate) GetEvenMonthWildcard() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetMonth()) == (14)))
}

func (m *_BACnetTagPayloadDate) GetDayOfMonthIsWildcard() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetDayOfMonth()) == (m.GetWildcard())))
}

func (m *_BACnetTagPayloadDate) GetLastDayOfMonthWildcard() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetDayOfMonth()) == (32)))
}

func (m *_BACnetTagPayloadDate) GetOddDayOfMonthWildcard() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetDayOfMonth()) == (33)))
}

func (m *_BACnetTagPayloadDate) GetEvenDayOfMonthWildcard() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetDayOfMonth()) == (34)))
}

func (m *_BACnetTagPayloadDate) GetDayOfWeekIsWildcard() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetDayOfWeek()) == (m.GetWildcard())))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetTagPayloadDate factory function for _BACnetTagPayloadDate
func NewBACnetTagPayloadDate(yearMinus1900 uint8, month uint8, dayOfMonth uint8, dayOfWeek uint8) *_BACnetTagPayloadDate {
	return &_BACnetTagPayloadDate{YearMinus1900: yearMinus1900, Month: month, DayOfMonth: dayOfMonth, DayOfWeek: dayOfWeek}
}

// Deprecated: use the interface for direct cast
func CastBACnetTagPayloadDate(structType interface{}) BACnetTagPayloadDate {
	if casted, ok := structType.(BACnetTagPayloadDate); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetTagPayloadDate); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetTagPayloadDate) GetTypeName() string {
	return "BACnetTagPayloadDate"
}

func (m *_BACnetTagPayloadDate) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// A virtual field doesn't have any in- or output.

	// Simple field (yearMinus1900)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// Simple field (month)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// Simple field (dayOfMonth)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// Simple field (dayOfWeek)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetTagPayloadDate) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetTagPayloadDateParse(theBytes []byte) (BACnetTagPayloadDate, error) {
	return BACnetTagPayloadDateParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func BACnetTagPayloadDateParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetTagPayloadDate, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetTagPayloadDate"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetTagPayloadDate")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Virtual field
	_wildcard := 0xFF
	wildcard := uint8(_wildcard)
	_ = wildcard

	// Simple Field (yearMinus1900)
	_yearMinus1900, _yearMinus1900Err := readBuffer.ReadUint8("yearMinus1900", 8)
	if _yearMinus1900Err != nil {
		return nil, errors.Wrap(_yearMinus1900Err, "Error parsing 'yearMinus1900' field of BACnetTagPayloadDate")
	}
	yearMinus1900 := _yearMinus1900

	// Virtual field
	_yearIsWildcard := bool((yearMinus1900) == (wildcard))
	yearIsWildcard := bool(_yearIsWildcard)
	_ = yearIsWildcard

	// Virtual field
	_year := uint16(yearMinus1900) + uint16(uint16(1900))
	year := uint16(_year)
	_ = year

	// Simple Field (month)
	_month, _monthErr := readBuffer.ReadUint8("month", 8)
	if _monthErr != nil {
		return nil, errors.Wrap(_monthErr, "Error parsing 'month' field of BACnetTagPayloadDate")
	}
	month := _month

	// Virtual field
	_monthIsWildcard := bool((month) == (wildcard))
	monthIsWildcard := bool(_monthIsWildcard)
	_ = monthIsWildcard

	// Virtual field
	_oddMonthWildcard := bool((month) == (13))
	oddMonthWildcard := bool(_oddMonthWildcard)
	_ = oddMonthWildcard

	// Virtual field
	_evenMonthWildcard := bool((month) == (14))
	evenMonthWildcard := bool(_evenMonthWildcard)
	_ = evenMonthWildcard

	// Simple Field (dayOfMonth)
	_dayOfMonth, _dayOfMonthErr := readBuffer.ReadUint8("dayOfMonth", 8)
	if _dayOfMonthErr != nil {
		return nil, errors.Wrap(_dayOfMonthErr, "Error parsing 'dayOfMonth' field of BACnetTagPayloadDate")
	}
	dayOfMonth := _dayOfMonth

	// Virtual field
	_dayOfMonthIsWildcard := bool((dayOfMonth) == (wildcard))
	dayOfMonthIsWildcard := bool(_dayOfMonthIsWildcard)
	_ = dayOfMonthIsWildcard

	// Virtual field
	_lastDayOfMonthWildcard := bool((dayOfMonth) == (32))
	lastDayOfMonthWildcard := bool(_lastDayOfMonthWildcard)
	_ = lastDayOfMonthWildcard

	// Virtual field
	_oddDayOfMonthWildcard := bool((dayOfMonth) == (33))
	oddDayOfMonthWildcard := bool(_oddDayOfMonthWildcard)
	_ = oddDayOfMonthWildcard

	// Virtual field
	_evenDayOfMonthWildcard := bool((dayOfMonth) == (34))
	evenDayOfMonthWildcard := bool(_evenDayOfMonthWildcard)
	_ = evenDayOfMonthWildcard

	// Simple Field (dayOfWeek)
	_dayOfWeek, _dayOfWeekErr := readBuffer.ReadUint8("dayOfWeek", 8)
	if _dayOfWeekErr != nil {
		return nil, errors.Wrap(_dayOfWeekErr, "Error parsing 'dayOfWeek' field of BACnetTagPayloadDate")
	}
	dayOfWeek := _dayOfWeek

	// Virtual field
	_dayOfWeekIsWildcard := bool((dayOfWeek) == (wildcard))
	dayOfWeekIsWildcard := bool(_dayOfWeekIsWildcard)
	_ = dayOfWeekIsWildcard

	if closeErr := readBuffer.CloseContext("BACnetTagPayloadDate"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetTagPayloadDate")
	}

	// Create the instance
	return &_BACnetTagPayloadDate{
		YearMinus1900: yearMinus1900,
		Month:         month,
		DayOfMonth:    dayOfMonth,
		DayOfWeek:     dayOfWeek,
	}, nil
}

func (m *_BACnetTagPayloadDate) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetTagPayloadDate) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetTagPayloadDate"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetTagPayloadDate")
	}
	// Virtual field
	if _wildcardErr := writeBuffer.WriteVirtual(ctx, "wildcard", m.GetWildcard()); _wildcardErr != nil {
		return errors.Wrap(_wildcardErr, "Error serializing 'wildcard' field")
	}

	// Simple Field (yearMinus1900)
	yearMinus1900 := uint8(m.GetYearMinus1900())
	_yearMinus1900Err := writeBuffer.WriteUint8("yearMinus1900", 8, (yearMinus1900))
	if _yearMinus1900Err != nil {
		return errors.Wrap(_yearMinus1900Err, "Error serializing 'yearMinus1900' field")
	}
	// Virtual field
	if _yearIsWildcardErr := writeBuffer.WriteVirtual(ctx, "yearIsWildcard", m.GetYearIsWildcard()); _yearIsWildcardErr != nil {
		return errors.Wrap(_yearIsWildcardErr, "Error serializing 'yearIsWildcard' field")
	}
	// Virtual field
	if _yearErr := writeBuffer.WriteVirtual(ctx, "year", m.GetYear()); _yearErr != nil {
		return errors.Wrap(_yearErr, "Error serializing 'year' field")
	}

	// Simple Field (month)
	month := uint8(m.GetMonth())
	_monthErr := writeBuffer.WriteUint8("month", 8, (month))
	if _monthErr != nil {
		return errors.Wrap(_monthErr, "Error serializing 'month' field")
	}
	// Virtual field
	if _monthIsWildcardErr := writeBuffer.WriteVirtual(ctx, "monthIsWildcard", m.GetMonthIsWildcard()); _monthIsWildcardErr != nil {
		return errors.Wrap(_monthIsWildcardErr, "Error serializing 'monthIsWildcard' field")
	}
	// Virtual field
	if _oddMonthWildcardErr := writeBuffer.WriteVirtual(ctx, "oddMonthWildcard", m.GetOddMonthWildcard()); _oddMonthWildcardErr != nil {
		return errors.Wrap(_oddMonthWildcardErr, "Error serializing 'oddMonthWildcard' field")
	}
	// Virtual field
	if _evenMonthWildcardErr := writeBuffer.WriteVirtual(ctx, "evenMonthWildcard", m.GetEvenMonthWildcard()); _evenMonthWildcardErr != nil {
		return errors.Wrap(_evenMonthWildcardErr, "Error serializing 'evenMonthWildcard' field")
	}

	// Simple Field (dayOfMonth)
	dayOfMonth := uint8(m.GetDayOfMonth())
	_dayOfMonthErr := writeBuffer.WriteUint8("dayOfMonth", 8, (dayOfMonth))
	if _dayOfMonthErr != nil {
		return errors.Wrap(_dayOfMonthErr, "Error serializing 'dayOfMonth' field")
	}
	// Virtual field
	if _dayOfMonthIsWildcardErr := writeBuffer.WriteVirtual(ctx, "dayOfMonthIsWildcard", m.GetDayOfMonthIsWildcard()); _dayOfMonthIsWildcardErr != nil {
		return errors.Wrap(_dayOfMonthIsWildcardErr, "Error serializing 'dayOfMonthIsWildcard' field")
	}
	// Virtual field
	if _lastDayOfMonthWildcardErr := writeBuffer.WriteVirtual(ctx, "lastDayOfMonthWildcard", m.GetLastDayOfMonthWildcard()); _lastDayOfMonthWildcardErr != nil {
		return errors.Wrap(_lastDayOfMonthWildcardErr, "Error serializing 'lastDayOfMonthWildcard' field")
	}
	// Virtual field
	if _oddDayOfMonthWildcardErr := writeBuffer.WriteVirtual(ctx, "oddDayOfMonthWildcard", m.GetOddDayOfMonthWildcard()); _oddDayOfMonthWildcardErr != nil {
		return errors.Wrap(_oddDayOfMonthWildcardErr, "Error serializing 'oddDayOfMonthWildcard' field")
	}
	// Virtual field
	if _evenDayOfMonthWildcardErr := writeBuffer.WriteVirtual(ctx, "evenDayOfMonthWildcard", m.GetEvenDayOfMonthWildcard()); _evenDayOfMonthWildcardErr != nil {
		return errors.Wrap(_evenDayOfMonthWildcardErr, "Error serializing 'evenDayOfMonthWildcard' field")
	}

	// Simple Field (dayOfWeek)
	dayOfWeek := uint8(m.GetDayOfWeek())
	_dayOfWeekErr := writeBuffer.WriteUint8("dayOfWeek", 8, (dayOfWeek))
	if _dayOfWeekErr != nil {
		return errors.Wrap(_dayOfWeekErr, "Error serializing 'dayOfWeek' field")
	}
	// Virtual field
	if _dayOfWeekIsWildcardErr := writeBuffer.WriteVirtual(ctx, "dayOfWeekIsWildcard", m.GetDayOfWeekIsWildcard()); _dayOfWeekIsWildcardErr != nil {
		return errors.Wrap(_dayOfWeekIsWildcardErr, "Error serializing 'dayOfWeekIsWildcard' field")
	}

	if popErr := writeBuffer.PopContext("BACnetTagPayloadDate"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetTagPayloadDate")
	}
	return nil
}

func (m *_BACnetTagPayloadDate) isBACnetTagPayloadDate() bool {
	return true
}

func (m *_BACnetTagPayloadDate) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
