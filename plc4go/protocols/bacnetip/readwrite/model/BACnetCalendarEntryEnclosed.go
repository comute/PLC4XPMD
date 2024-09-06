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

// BACnetCalendarEntryEnclosed is the corresponding interface of BACnetCalendarEntryEnclosed
type BACnetCalendarEntryEnclosed interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetCalendarEntry returns CalendarEntry (property field)
	GetCalendarEntry() BACnetCalendarEntry
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
	// IsBACnetCalendarEntryEnclosed is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetCalendarEntryEnclosed()
}

// _BACnetCalendarEntryEnclosed is the data-structure of this message
type _BACnetCalendarEntryEnclosed struct {
	OpeningTag    BACnetOpeningTag
	CalendarEntry BACnetCalendarEntry
	ClosingTag    BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

var _ BACnetCalendarEntryEnclosed = (*_BACnetCalendarEntryEnclosed)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetCalendarEntryEnclosed) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetCalendarEntryEnclosed) GetCalendarEntry() BACnetCalendarEntry {
	return m.CalendarEntry
}

func (m *_BACnetCalendarEntryEnclosed) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetCalendarEntryEnclosed factory function for _BACnetCalendarEntryEnclosed
func NewBACnetCalendarEntryEnclosed(openingTag BACnetOpeningTag, calendarEntry BACnetCalendarEntry, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetCalendarEntryEnclosed {
	if openingTag == nil {
		panic("openingTag of type BACnetOpeningTag for BACnetCalendarEntryEnclosed must not be nil")
	}
	if calendarEntry == nil {
		panic("calendarEntry of type BACnetCalendarEntry for BACnetCalendarEntryEnclosed must not be nil")
	}
	if closingTag == nil {
		panic("closingTag of type BACnetClosingTag for BACnetCalendarEntryEnclosed must not be nil")
	}
	return &_BACnetCalendarEntryEnclosed{OpeningTag: openingTag, CalendarEntry: calendarEntry, ClosingTag: closingTag, TagNumber: tagNumber}
}

// Deprecated: use the interface for direct cast
func CastBACnetCalendarEntryEnclosed(structType any) BACnetCalendarEntryEnclosed {
	if casted, ok := structType.(BACnetCalendarEntryEnclosed); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetCalendarEntryEnclosed); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetCalendarEntryEnclosed) GetTypeName() string {
	return "BACnetCalendarEntryEnclosed"
}

func (m *_BACnetCalendarEntryEnclosed) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Simple field (calendarEntry)
	lengthInBits += m.CalendarEntry.GetLengthInBits(ctx)

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetCalendarEntryEnclosed) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetCalendarEntryEnclosedParse(ctx context.Context, theBytes []byte, tagNumber uint8) (BACnetCalendarEntryEnclosed, error) {
	return BACnetCalendarEntryEnclosedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetCalendarEntryEnclosedParseWithBufferProducer(tagNumber uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetCalendarEntryEnclosed, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetCalendarEntryEnclosed, error) {
		return BACnetCalendarEntryEnclosedParseWithBuffer(ctx, readBuffer, tagNumber)
	}
}

func BACnetCalendarEntryEnclosedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetCalendarEntryEnclosed, error) {
	v, err := (&_BACnetCalendarEntryEnclosed{TagNumber: tagNumber}).parse(ctx, readBuffer, tagNumber)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *_BACnetCalendarEntryEnclosed) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (__bACnetCalendarEntryEnclosed BACnetCalendarEntryEnclosed, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetCalendarEntryEnclosed"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetCalendarEntryEnclosed")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	m.OpeningTag = openingTag

	calendarEntry, err := ReadSimpleField[BACnetCalendarEntry](ctx, "calendarEntry", ReadComplex[BACnetCalendarEntry](BACnetCalendarEntryParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'calendarEntry' field"))
	}
	m.CalendarEntry = calendarEntry

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("BACnetCalendarEntryEnclosed"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetCalendarEntryEnclosed")
	}

	return m, nil
}

func (m *_BACnetCalendarEntryEnclosed) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetCalendarEntryEnclosed) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetCalendarEntryEnclosed"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetCalendarEntryEnclosed")
	}

	if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'openingTag' field")
	}

	if err := WriteSimpleField[BACnetCalendarEntry](ctx, "calendarEntry", m.GetCalendarEntry(), WriteComplex[BACnetCalendarEntry](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'calendarEntry' field")
	}

	if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetCalendarEntryEnclosed"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetCalendarEntryEnclosed")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetCalendarEntryEnclosed) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetCalendarEntryEnclosed) IsBACnetCalendarEntryEnclosed() {}

func (m *_BACnetCalendarEntryEnclosed) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
