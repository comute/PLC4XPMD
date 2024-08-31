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

// BACnetAccumulatorRecord is the corresponding interface of BACnetAccumulatorRecord
type BACnetAccumulatorRecord interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetTimestamp returns Timestamp (property field)
	GetTimestamp() BACnetDateTimeEnclosed
	// GetPresentValue returns PresentValue (property field)
	GetPresentValue() BACnetContextTagSignedInteger
	// GetAccumulatedValue returns AccumulatedValue (property field)
	GetAccumulatedValue() BACnetContextTagSignedInteger
	// GetAccumulatorStatus returns AccumulatorStatus (property field)
	GetAccumulatorStatus() BACnetAccumulatorRecordAccumulatorStatusTagged
}

// BACnetAccumulatorRecordExactly can be used when we want exactly this type and not a type which fulfills BACnetAccumulatorRecord.
// This is useful for switch cases.
type BACnetAccumulatorRecordExactly interface {
	BACnetAccumulatorRecord
	isBACnetAccumulatorRecord() bool
}

// _BACnetAccumulatorRecord is the data-structure of this message
type _BACnetAccumulatorRecord struct {
	Timestamp         BACnetDateTimeEnclosed
	PresentValue      BACnetContextTagSignedInteger
	AccumulatedValue  BACnetContextTagSignedInteger
	AccumulatorStatus BACnetAccumulatorRecordAccumulatorStatusTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetAccumulatorRecord) GetTimestamp() BACnetDateTimeEnclosed {
	return m.Timestamp
}

func (m *_BACnetAccumulatorRecord) GetPresentValue() BACnetContextTagSignedInteger {
	return m.PresentValue
}

func (m *_BACnetAccumulatorRecord) GetAccumulatedValue() BACnetContextTagSignedInteger {
	return m.AccumulatedValue
}

func (m *_BACnetAccumulatorRecord) GetAccumulatorStatus() BACnetAccumulatorRecordAccumulatorStatusTagged {
	return m.AccumulatorStatus
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetAccumulatorRecord factory function for _BACnetAccumulatorRecord
func NewBACnetAccumulatorRecord(timestamp BACnetDateTimeEnclosed, presentValue BACnetContextTagSignedInteger, accumulatedValue BACnetContextTagSignedInteger, accumulatorStatus BACnetAccumulatorRecordAccumulatorStatusTagged) *_BACnetAccumulatorRecord {
	return &_BACnetAccumulatorRecord{Timestamp: timestamp, PresentValue: presentValue, AccumulatedValue: accumulatedValue, AccumulatorStatus: accumulatorStatus}
}

// Deprecated: use the interface for direct cast
func CastBACnetAccumulatorRecord(structType any) BACnetAccumulatorRecord {
	if casted, ok := structType.(BACnetAccumulatorRecord); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetAccumulatorRecord); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetAccumulatorRecord) GetTypeName() string {
	return "BACnetAccumulatorRecord"
}

func (m *_BACnetAccumulatorRecord) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (timestamp)
	lengthInBits += m.Timestamp.GetLengthInBits(ctx)

	// Simple field (presentValue)
	lengthInBits += m.PresentValue.GetLengthInBits(ctx)

	// Simple field (accumulatedValue)
	lengthInBits += m.AccumulatedValue.GetLengthInBits(ctx)

	// Simple field (accumulatorStatus)
	lengthInBits += m.AccumulatorStatus.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetAccumulatorRecord) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetAccumulatorRecordParse(ctx context.Context, theBytes []byte) (BACnetAccumulatorRecord, error) {
	return BACnetAccumulatorRecordParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetAccumulatorRecordParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAccumulatorRecord, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAccumulatorRecord, error) {
		return BACnetAccumulatorRecordParseWithBuffer(ctx, readBuffer)
	}
}

func BACnetAccumulatorRecordParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAccumulatorRecord, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetAccumulatorRecord"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetAccumulatorRecord")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	timestamp, err := ReadSimpleField[BACnetDateTimeEnclosed](ctx, "timestamp", ReadComplex[BACnetDateTimeEnclosed](BACnetDateTimeEnclosedParseWithBufferProducer((uint8)(uint8(0))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timestamp' field"))
	}

	presentValue, err := ReadSimpleField[BACnetContextTagSignedInteger](ctx, "presentValue", ReadComplex[BACnetContextTagSignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagSignedInteger]((uint8)(uint8(1)), (BACnetDataType)(BACnetDataType_SIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'presentValue' field"))
	}

	accumulatedValue, err := ReadSimpleField[BACnetContextTagSignedInteger](ctx, "accumulatedValue", ReadComplex[BACnetContextTagSignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagSignedInteger]((uint8)(uint8(2)), (BACnetDataType)(BACnetDataType_SIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'accumulatedValue' field"))
	}

	accumulatorStatus, err := ReadSimpleField[BACnetAccumulatorRecordAccumulatorStatusTagged](ctx, "accumulatorStatus", ReadComplex[BACnetAccumulatorRecordAccumulatorStatusTagged](BACnetAccumulatorRecordAccumulatorStatusTaggedParseWithBufferProducer((uint8)(uint8(3)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'accumulatorStatus' field"))
	}

	if closeErr := readBuffer.CloseContext("BACnetAccumulatorRecord"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetAccumulatorRecord")
	}

	// Create the instance
	return &_BACnetAccumulatorRecord{
		Timestamp:         timestamp,
		PresentValue:      presentValue,
		AccumulatedValue:  accumulatedValue,
		AccumulatorStatus: accumulatorStatus,
	}, nil
}

func (m *_BACnetAccumulatorRecord) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetAccumulatorRecord) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetAccumulatorRecord"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetAccumulatorRecord")
	}

	// Simple Field (timestamp)
	if pushErr := writeBuffer.PushContext("timestamp"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for timestamp")
	}
	_timestampErr := writeBuffer.WriteSerializable(ctx, m.GetTimestamp())
	if popErr := writeBuffer.PopContext("timestamp"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for timestamp")
	}
	if _timestampErr != nil {
		return errors.Wrap(_timestampErr, "Error serializing 'timestamp' field")
	}

	// Simple Field (presentValue)
	if pushErr := writeBuffer.PushContext("presentValue"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for presentValue")
	}
	_presentValueErr := writeBuffer.WriteSerializable(ctx, m.GetPresentValue())
	if popErr := writeBuffer.PopContext("presentValue"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for presentValue")
	}
	if _presentValueErr != nil {
		return errors.Wrap(_presentValueErr, "Error serializing 'presentValue' field")
	}

	// Simple Field (accumulatedValue)
	if pushErr := writeBuffer.PushContext("accumulatedValue"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for accumulatedValue")
	}
	_accumulatedValueErr := writeBuffer.WriteSerializable(ctx, m.GetAccumulatedValue())
	if popErr := writeBuffer.PopContext("accumulatedValue"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for accumulatedValue")
	}
	if _accumulatedValueErr != nil {
		return errors.Wrap(_accumulatedValueErr, "Error serializing 'accumulatedValue' field")
	}

	// Simple Field (accumulatorStatus)
	if pushErr := writeBuffer.PushContext("accumulatorStatus"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for accumulatorStatus")
	}
	_accumulatorStatusErr := writeBuffer.WriteSerializable(ctx, m.GetAccumulatorStatus())
	if popErr := writeBuffer.PopContext("accumulatorStatus"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for accumulatorStatus")
	}
	if _accumulatorStatusErr != nil {
		return errors.Wrap(_accumulatorStatusErr, "Error serializing 'accumulatorStatus' field")
	}

	if popErr := writeBuffer.PopContext("BACnetAccumulatorRecord"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetAccumulatorRecord")
	}
	return nil
}

func (m *_BACnetAccumulatorRecord) isBACnetAccumulatorRecord() bool {
	return true
}

func (m *_BACnetAccumulatorRecord) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
