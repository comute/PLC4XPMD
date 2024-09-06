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

// BACnetEventParameterChangeOfTimerAlarmValue is the corresponding interface of BACnetEventParameterChangeOfTimerAlarmValue
type BACnetEventParameterChangeOfTimerAlarmValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetAlarmValues returns AlarmValues (property field)
	GetAlarmValues() []BACnetTimerStateTagged
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
	// IsBACnetEventParameterChangeOfTimerAlarmValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetEventParameterChangeOfTimerAlarmValue()
}

// _BACnetEventParameterChangeOfTimerAlarmValue is the data-structure of this message
type _BACnetEventParameterChangeOfTimerAlarmValue struct {
	OpeningTag  BACnetOpeningTag
	AlarmValues []BACnetTimerStateTagged
	ClosingTag  BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

var _ BACnetEventParameterChangeOfTimerAlarmValue = (*_BACnetEventParameterChangeOfTimerAlarmValue)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetEventParameterChangeOfTimerAlarmValue) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetEventParameterChangeOfTimerAlarmValue) GetAlarmValues() []BACnetTimerStateTagged {
	return m.AlarmValues
}

func (m *_BACnetEventParameterChangeOfTimerAlarmValue) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetEventParameterChangeOfTimerAlarmValue factory function for _BACnetEventParameterChangeOfTimerAlarmValue
func NewBACnetEventParameterChangeOfTimerAlarmValue(openingTag BACnetOpeningTag, alarmValues []BACnetTimerStateTagged, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetEventParameterChangeOfTimerAlarmValue {
	if openingTag == nil {
		panic("openingTag of type BACnetOpeningTag for BACnetEventParameterChangeOfTimerAlarmValue must not be nil")
	}
	if closingTag == nil {
		panic("closingTag of type BACnetClosingTag for BACnetEventParameterChangeOfTimerAlarmValue must not be nil")
	}
	return &_BACnetEventParameterChangeOfTimerAlarmValue{OpeningTag: openingTag, AlarmValues: alarmValues, ClosingTag: closingTag, TagNumber: tagNumber}
}

// Deprecated: use the interface for direct cast
func CastBACnetEventParameterChangeOfTimerAlarmValue(structType any) BACnetEventParameterChangeOfTimerAlarmValue {
	if casted, ok := structType.(BACnetEventParameterChangeOfTimerAlarmValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetEventParameterChangeOfTimerAlarmValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetEventParameterChangeOfTimerAlarmValue) GetTypeName() string {
	return "BACnetEventParameterChangeOfTimerAlarmValue"
}

func (m *_BACnetEventParameterChangeOfTimerAlarmValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Array field
	if len(m.AlarmValues) > 0 {
		for _, element := range m.AlarmValues {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetEventParameterChangeOfTimerAlarmValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetEventParameterChangeOfTimerAlarmValueParse(ctx context.Context, theBytes []byte, tagNumber uint8) (BACnetEventParameterChangeOfTimerAlarmValue, error) {
	return BACnetEventParameterChangeOfTimerAlarmValueParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetEventParameterChangeOfTimerAlarmValueParseWithBufferProducer(tagNumber uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetEventParameterChangeOfTimerAlarmValue, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetEventParameterChangeOfTimerAlarmValue, error) {
		return BACnetEventParameterChangeOfTimerAlarmValueParseWithBuffer(ctx, readBuffer, tagNumber)
	}
}

func BACnetEventParameterChangeOfTimerAlarmValueParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetEventParameterChangeOfTimerAlarmValue, error) {
	v, err := (&_BACnetEventParameterChangeOfTimerAlarmValue{TagNumber: tagNumber}).parse(ctx, readBuffer, tagNumber)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *_BACnetEventParameterChangeOfTimerAlarmValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (__bACnetEventParameterChangeOfTimerAlarmValue BACnetEventParameterChangeOfTimerAlarmValue, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetEventParameterChangeOfTimerAlarmValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetEventParameterChangeOfTimerAlarmValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	m.OpeningTag = openingTag

	alarmValues, err := ReadTerminatedArrayField[BACnetTimerStateTagged](ctx, "alarmValues", ReadComplex[BACnetTimerStateTagged](BACnetTimerStateTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'alarmValues' field"))
	}
	m.AlarmValues = alarmValues

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("BACnetEventParameterChangeOfTimerAlarmValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetEventParameterChangeOfTimerAlarmValue")
	}

	return m, nil
}

func (m *_BACnetEventParameterChangeOfTimerAlarmValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetEventParameterChangeOfTimerAlarmValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetEventParameterChangeOfTimerAlarmValue"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetEventParameterChangeOfTimerAlarmValue")
	}

	if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'openingTag' field")
	}

	if err := WriteComplexTypeArrayField(ctx, "alarmValues", m.GetAlarmValues(), writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'alarmValues' field")
	}

	if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetEventParameterChangeOfTimerAlarmValue"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetEventParameterChangeOfTimerAlarmValue")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetEventParameterChangeOfTimerAlarmValue) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetEventParameterChangeOfTimerAlarmValue) IsBACnetEventParameterChangeOfTimerAlarmValue() {
}

func (m *_BACnetEventParameterChangeOfTimerAlarmValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
