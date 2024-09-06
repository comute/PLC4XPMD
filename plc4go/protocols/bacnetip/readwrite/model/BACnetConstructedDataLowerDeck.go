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

// BACnetConstructedDataLowerDeck is the corresponding interface of BACnetConstructedDataLowerDeck
type BACnetConstructedDataLowerDeck interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetLowerDeck returns LowerDeck (property field)
	GetLowerDeck() BACnetApplicationTagObjectIdentifier
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagObjectIdentifier
	// IsBACnetConstructedDataLowerDeck is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataLowerDeck()
}

// _BACnetConstructedDataLowerDeck is the data-structure of this message
type _BACnetConstructedDataLowerDeck struct {
	BACnetConstructedDataContract
	LowerDeck BACnetApplicationTagObjectIdentifier
}

var _ BACnetConstructedDataLowerDeck = (*_BACnetConstructedDataLowerDeck)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataLowerDeck)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLowerDeck) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataLowerDeck) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LOWER_DECK
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLowerDeck) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLowerDeck) GetLowerDeck() BACnetApplicationTagObjectIdentifier {
	return m.LowerDeck
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataLowerDeck) GetActualValue() BACnetApplicationTagObjectIdentifier {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagObjectIdentifier(m.GetLowerDeck())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataLowerDeck factory function for _BACnetConstructedDataLowerDeck
func NewBACnetConstructedDataLowerDeck(lowerDeck BACnetApplicationTagObjectIdentifier, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLowerDeck {
	if lowerDeck == nil {
		panic("lowerDeck of type BACnetApplicationTagObjectIdentifier for BACnetConstructedDataLowerDeck must not be nil")
	}
	_result := &_BACnetConstructedDataLowerDeck{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		LowerDeck:                     lowerDeck,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLowerDeck(structType any) BACnetConstructedDataLowerDeck {
	if casted, ok := structType.(BACnetConstructedDataLowerDeck); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLowerDeck); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLowerDeck) GetTypeName() string {
	return "BACnetConstructedDataLowerDeck"
}

func (m *_BACnetConstructedDataLowerDeck) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (lowerDeck)
	lengthInBits += m.LowerDeck.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataLowerDeck) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataLowerDeck) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataLowerDeck BACnetConstructedDataLowerDeck, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLowerDeck"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLowerDeck")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	lowerDeck, err := ReadSimpleField[BACnetApplicationTagObjectIdentifier](ctx, "lowerDeck", ReadComplex[BACnetApplicationTagObjectIdentifier](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagObjectIdentifier](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lowerDeck' field"))
	}
	m.LowerDeck = lowerDeck

	actualValue, err := ReadVirtualField[BACnetApplicationTagObjectIdentifier](ctx, "actualValue", (*BACnetApplicationTagObjectIdentifier)(nil), lowerDeck)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLowerDeck"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLowerDeck")
	}

	return m, nil
}

func (m *_BACnetConstructedDataLowerDeck) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataLowerDeck) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLowerDeck"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLowerDeck")
		}

		if err := WriteSimpleField[BACnetApplicationTagObjectIdentifier](ctx, "lowerDeck", m.GetLowerDeck(), WriteComplex[BACnetApplicationTagObjectIdentifier](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'lowerDeck' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLowerDeck"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLowerDeck")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLowerDeck) IsBACnetConstructedDataLowerDeck() {}

func (m *_BACnetConstructedDataLowerDeck) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
