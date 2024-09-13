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

// BACnetConstructedDataNextStoppingFloor is the corresponding interface of BACnetConstructedDataNextStoppingFloor
type BACnetConstructedDataNextStoppingFloor interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetNextStoppingFloor returns NextStoppingFloor (property field)
	GetNextStoppingFloor() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetConstructedDataNextStoppingFloor is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataNextStoppingFloor()
}

// _BACnetConstructedDataNextStoppingFloor is the data-structure of this message
type _BACnetConstructedDataNextStoppingFloor struct {
	BACnetConstructedDataContract
	NextStoppingFloor BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataNextStoppingFloor = (*_BACnetConstructedDataNextStoppingFloor)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataNextStoppingFloor)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataNextStoppingFloor) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataNextStoppingFloor) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_NEXT_STOPPING_FLOOR
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataNextStoppingFloor) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataNextStoppingFloor) GetNextStoppingFloor() BACnetApplicationTagUnsignedInteger {
	return m.NextStoppingFloor
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataNextStoppingFloor) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetNextStoppingFloor())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataNextStoppingFloor factory function for _BACnetConstructedDataNextStoppingFloor
func NewBACnetConstructedDataNextStoppingFloor(nextStoppingFloor BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataNextStoppingFloor {
	if nextStoppingFloor == nil {
		panic("nextStoppingFloor of type BACnetApplicationTagUnsignedInteger for BACnetConstructedDataNextStoppingFloor must not be nil")
	}
	_result := &_BACnetConstructedDataNextStoppingFloor{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		NextStoppingFloor:             nextStoppingFloor,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataNextStoppingFloor(structType any) BACnetConstructedDataNextStoppingFloor {
	if casted, ok := structType.(BACnetConstructedDataNextStoppingFloor); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataNextStoppingFloor); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataNextStoppingFloor) GetTypeName() string {
	return "BACnetConstructedDataNextStoppingFloor"
}

func (m *_BACnetConstructedDataNextStoppingFloor) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (nextStoppingFloor)
	lengthInBits += m.NextStoppingFloor.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataNextStoppingFloor) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataNextStoppingFloor) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataNextStoppingFloor BACnetConstructedDataNextStoppingFloor, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataNextStoppingFloor"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataNextStoppingFloor")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	nextStoppingFloor, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "nextStoppingFloor", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'nextStoppingFloor' field"))
	}
	m.NextStoppingFloor = nextStoppingFloor

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), nextStoppingFloor)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataNextStoppingFloor"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataNextStoppingFloor")
	}

	return m, nil
}

func (m *_BACnetConstructedDataNextStoppingFloor) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataNextStoppingFloor) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataNextStoppingFloor"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataNextStoppingFloor")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "nextStoppingFloor", m.GetNextStoppingFloor(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'nextStoppingFloor' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataNextStoppingFloor"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataNextStoppingFloor")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataNextStoppingFloor) IsBACnetConstructedDataNextStoppingFloor() {}

func (m *_BACnetConstructedDataNextStoppingFloor) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
