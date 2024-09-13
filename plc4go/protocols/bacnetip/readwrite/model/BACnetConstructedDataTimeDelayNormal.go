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

// BACnetConstructedDataTimeDelayNormal is the corresponding interface of BACnetConstructedDataTimeDelayNormal
type BACnetConstructedDataTimeDelayNormal interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetTimeDelayNormal returns TimeDelayNormal (property field)
	GetTimeDelayNormal() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetConstructedDataTimeDelayNormal is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataTimeDelayNormal()
}

// _BACnetConstructedDataTimeDelayNormal is the data-structure of this message
type _BACnetConstructedDataTimeDelayNormal struct {
	BACnetConstructedDataContract
	TimeDelayNormal BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataTimeDelayNormal = (*_BACnetConstructedDataTimeDelayNormal)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataTimeDelayNormal)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataTimeDelayNormal) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataTimeDelayNormal) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_TIME_DELAY_NORMAL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataTimeDelayNormal) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataTimeDelayNormal) GetTimeDelayNormal() BACnetApplicationTagUnsignedInteger {
	return m.TimeDelayNormal
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataTimeDelayNormal) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetTimeDelayNormal())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataTimeDelayNormal factory function for _BACnetConstructedDataTimeDelayNormal
func NewBACnetConstructedDataTimeDelayNormal(timeDelayNormal BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataTimeDelayNormal {
	if timeDelayNormal == nil {
		panic("timeDelayNormal of type BACnetApplicationTagUnsignedInteger for BACnetConstructedDataTimeDelayNormal must not be nil")
	}
	_result := &_BACnetConstructedDataTimeDelayNormal{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		TimeDelayNormal:               timeDelayNormal,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataTimeDelayNormal(structType any) BACnetConstructedDataTimeDelayNormal {
	if casted, ok := structType.(BACnetConstructedDataTimeDelayNormal); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataTimeDelayNormal); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataTimeDelayNormal) GetTypeName() string {
	return "BACnetConstructedDataTimeDelayNormal"
}

func (m *_BACnetConstructedDataTimeDelayNormal) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (timeDelayNormal)
	lengthInBits += m.TimeDelayNormal.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataTimeDelayNormal) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataTimeDelayNormal) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataTimeDelayNormal BACnetConstructedDataTimeDelayNormal, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataTimeDelayNormal"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataTimeDelayNormal")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	timeDelayNormal, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "timeDelayNormal", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timeDelayNormal' field"))
	}
	m.TimeDelayNormal = timeDelayNormal

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), timeDelayNormal)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataTimeDelayNormal"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataTimeDelayNormal")
	}

	return m, nil
}

func (m *_BACnetConstructedDataTimeDelayNormal) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataTimeDelayNormal) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataTimeDelayNormal"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataTimeDelayNormal")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "timeDelayNormal", m.GetTimeDelayNormal(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'timeDelayNormal' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataTimeDelayNormal"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataTimeDelayNormal")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataTimeDelayNormal) IsBACnetConstructedDataTimeDelayNormal() {}

func (m *_BACnetConstructedDataTimeDelayNormal) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
