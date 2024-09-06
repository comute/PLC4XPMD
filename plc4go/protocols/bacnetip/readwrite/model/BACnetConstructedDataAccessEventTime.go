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

// BACnetConstructedDataAccessEventTime is the corresponding interface of BACnetConstructedDataAccessEventTime
type BACnetConstructedDataAccessEventTime interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetAccessEventTime returns AccessEventTime (property field)
	GetAccessEventTime() BACnetTimeStamp
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetTimeStamp
	// IsBACnetConstructedDataAccessEventTime is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataAccessEventTime()
}

// _BACnetConstructedDataAccessEventTime is the data-structure of this message
type _BACnetConstructedDataAccessEventTime struct {
	BACnetConstructedDataContract
	AccessEventTime BACnetTimeStamp
}

var _ BACnetConstructedDataAccessEventTime = (*_BACnetConstructedDataAccessEventTime)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataAccessEventTime)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAccessEventTime) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataAccessEventTime) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ACCESS_EVENT_TIME
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAccessEventTime) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataAccessEventTime) GetAccessEventTime() BACnetTimeStamp {
	return m.AccessEventTime
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataAccessEventTime) GetActualValue() BACnetTimeStamp {
	ctx := context.Background()
	_ = ctx
	return CastBACnetTimeStamp(m.GetAccessEventTime())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataAccessEventTime factory function for _BACnetConstructedDataAccessEventTime
func NewBACnetConstructedDataAccessEventTime(accessEventTime BACnetTimeStamp, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAccessEventTime {
	if accessEventTime == nil {
		panic("accessEventTime of type BACnetTimeStamp for BACnetConstructedDataAccessEventTime must not be nil")
	}
	_result := &_BACnetConstructedDataAccessEventTime{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		AccessEventTime:               accessEventTime,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAccessEventTime(structType any) BACnetConstructedDataAccessEventTime {
	if casted, ok := structType.(BACnetConstructedDataAccessEventTime); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAccessEventTime); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAccessEventTime) GetTypeName() string {
	return "BACnetConstructedDataAccessEventTime"
}

func (m *_BACnetConstructedDataAccessEventTime) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (accessEventTime)
	lengthInBits += m.AccessEventTime.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataAccessEventTime) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataAccessEventTime) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataAccessEventTime BACnetConstructedDataAccessEventTime, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAccessEventTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAccessEventTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	accessEventTime, err := ReadSimpleField[BACnetTimeStamp](ctx, "accessEventTime", ReadComplex[BACnetTimeStamp](BACnetTimeStampParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'accessEventTime' field"))
	}
	m.AccessEventTime = accessEventTime

	actualValue, err := ReadVirtualField[BACnetTimeStamp](ctx, "actualValue", (*BACnetTimeStamp)(nil), accessEventTime)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAccessEventTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAccessEventTime")
	}

	return m, nil
}

func (m *_BACnetConstructedDataAccessEventTime) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataAccessEventTime) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAccessEventTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAccessEventTime")
		}

		if err := WriteSimpleField[BACnetTimeStamp](ctx, "accessEventTime", m.GetAccessEventTime(), WriteComplex[BACnetTimeStamp](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'accessEventTime' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAccessEventTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAccessEventTime")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAccessEventTime) IsBACnetConstructedDataAccessEventTime() {}

func (m *_BACnetConstructedDataAccessEventTime) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
