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

// BACnetTimeStampTime is the corresponding interface of BACnetTimeStampTime
type BACnetTimeStampTime interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetTimeStamp
	// GetTimeValue returns TimeValue (property field)
	GetTimeValue() BACnetContextTagTime
}

// BACnetTimeStampTimeExactly can be used when we want exactly this type and not a type which fulfills BACnetTimeStampTime.
// This is useful for switch cases.
type BACnetTimeStampTimeExactly interface {
	BACnetTimeStampTime
	isBACnetTimeStampTime() bool
}

// _BACnetTimeStampTime is the data-structure of this message
type _BACnetTimeStampTime struct {
	*_BACnetTimeStamp
	TimeValue BACnetContextTagTime
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetTimeStampTime) InitializeParent(parent BACnetTimeStamp, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetTimeStampTime) GetParent() BACnetTimeStamp {
	return m._BACnetTimeStamp
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetTimeStampTime) GetTimeValue() BACnetContextTagTime {
	return m.TimeValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetTimeStampTime factory function for _BACnetTimeStampTime
func NewBACnetTimeStampTime(timeValue BACnetContextTagTime, peekedTagHeader BACnetTagHeader) *_BACnetTimeStampTime {
	_result := &_BACnetTimeStampTime{
		TimeValue:        timeValue,
		_BACnetTimeStamp: NewBACnetTimeStamp(peekedTagHeader),
	}
	_result._BACnetTimeStamp._BACnetTimeStampChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetTimeStampTime(structType any) BACnetTimeStampTime {
	if casted, ok := structType.(BACnetTimeStampTime); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetTimeStampTime); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetTimeStampTime) GetTypeName() string {
	return "BACnetTimeStampTime"
}

func (m *_BACnetTimeStampTime) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (timeValue)
	lengthInBits += m.TimeValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetTimeStampTime) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetTimeStampTimeParse(ctx context.Context, theBytes []byte) (BACnetTimeStampTime, error) {
	return BACnetTimeStampTimeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetTimeStampTimeParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetTimeStampTime, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetTimeStampTime, error) {
		return BACnetTimeStampTimeParseWithBuffer(ctx, readBuffer)
	}
}

func BACnetTimeStampTimeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetTimeStampTime, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetTimeStampTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetTimeStampTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	timeValue, err := ReadSimpleField[BACnetContextTagTime](ctx, "timeValue", ReadComplex[BACnetContextTagTime](BACnetContextTagParseWithBufferProducer[BACnetContextTagTime]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_TIME)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timeValue' field"))
	}

	if closeErr := readBuffer.CloseContext("BACnetTimeStampTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetTimeStampTime")
	}

	// Create a partially initialized instance
	_child := &_BACnetTimeStampTime{
		_BACnetTimeStamp: &_BACnetTimeStamp{},
		TimeValue:        timeValue,
	}
	_child._BACnetTimeStamp._BACnetTimeStampChildRequirements = _child
	return _child, nil
}

func (m *_BACnetTimeStampTime) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetTimeStampTime) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetTimeStampTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetTimeStampTime")
		}

		// Simple Field (timeValue)
		if pushErr := writeBuffer.PushContext("timeValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for timeValue")
		}
		_timeValueErr := writeBuffer.WriteSerializable(ctx, m.GetTimeValue())
		if popErr := writeBuffer.PopContext("timeValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for timeValue")
		}
		if _timeValueErr != nil {
			return errors.Wrap(_timeValueErr, "Error serializing 'timeValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetTimeStampTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetTimeStampTime")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetTimeStampTime) isBACnetTimeStampTime() bool {
	return true
}

func (m *_BACnetTimeStampTime) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
