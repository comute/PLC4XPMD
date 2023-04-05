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

// BACnetConstructedDataIntegerValueLowLimit is the corresponding interface of BACnetConstructedDataIntegerValueLowLimit
type BACnetConstructedDataIntegerValueLowLimit interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetLowLimit returns LowLimit (property field)
	GetLowLimit() BACnetApplicationTagSignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagSignedInteger
}

// BACnetConstructedDataIntegerValueLowLimitExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataIntegerValueLowLimit.
// This is useful for switch cases.
type BACnetConstructedDataIntegerValueLowLimitExactly interface {
	BACnetConstructedDataIntegerValueLowLimit
	isBACnetConstructedDataIntegerValueLowLimit() bool
}

// _BACnetConstructedDataIntegerValueLowLimit is the data-structure of this message
type _BACnetConstructedDataIntegerValueLowLimit struct {
	*_BACnetConstructedData
	LowLimit BACnetApplicationTagSignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataIntegerValueLowLimit) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_INTEGER_VALUE
}

func (m *_BACnetConstructedDataIntegerValueLowLimit) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LOW_LIMIT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataIntegerValueLowLimit) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataIntegerValueLowLimit) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataIntegerValueLowLimit) GetLowLimit() BACnetApplicationTagSignedInteger {
	return m.LowLimit
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataIntegerValueLowLimit) GetActualValue() BACnetApplicationTagSignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagSignedInteger(m.GetLowLimit())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataIntegerValueLowLimit factory function for _BACnetConstructedDataIntegerValueLowLimit
func NewBACnetConstructedDataIntegerValueLowLimit(lowLimit BACnetApplicationTagSignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataIntegerValueLowLimit {
	_result := &_BACnetConstructedDataIntegerValueLowLimit{
		LowLimit:               lowLimit,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataIntegerValueLowLimit(structType interface{}) BACnetConstructedDataIntegerValueLowLimit {
	if casted, ok := structType.(BACnetConstructedDataIntegerValueLowLimit); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataIntegerValueLowLimit); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataIntegerValueLowLimit) GetTypeName() string {
	return "BACnetConstructedDataIntegerValueLowLimit"
}

func (m *_BACnetConstructedDataIntegerValueLowLimit) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (lowLimit)
	lengthInBits += m.LowLimit.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataIntegerValueLowLimit) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataIntegerValueLowLimitParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataIntegerValueLowLimit, error) {
	return BACnetConstructedDataIntegerValueLowLimitParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataIntegerValueLowLimitParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataIntegerValueLowLimit, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataIntegerValueLowLimit"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataIntegerValueLowLimit")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (lowLimit)
	if pullErr := readBuffer.PullContext("lowLimit"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for lowLimit")
	}
	_lowLimit, _lowLimitErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _lowLimitErr != nil {
		return nil, errors.Wrap(_lowLimitErr, "Error parsing 'lowLimit' field of BACnetConstructedDataIntegerValueLowLimit")
	}
	lowLimit := _lowLimit.(BACnetApplicationTagSignedInteger)
	if closeErr := readBuffer.CloseContext("lowLimit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for lowLimit")
	}

	// Virtual field
	_actualValue := lowLimit
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataIntegerValueLowLimit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataIntegerValueLowLimit")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataIntegerValueLowLimit{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		LowLimit: lowLimit,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataIntegerValueLowLimit) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataIntegerValueLowLimit) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataIntegerValueLowLimit"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataIntegerValueLowLimit")
		}

		// Simple Field (lowLimit)
		if pushErr := writeBuffer.PushContext("lowLimit"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for lowLimit")
		}
		_lowLimitErr := writeBuffer.WriteSerializable(ctx, m.GetLowLimit())
		if popErr := writeBuffer.PopContext("lowLimit"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for lowLimit")
		}
		if _lowLimitErr != nil {
			return errors.Wrap(_lowLimitErr, "Error serializing 'lowLimit' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataIntegerValueLowLimit"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataIntegerValueLowLimit")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataIntegerValueLowLimit) isBACnetConstructedDataIntegerValueLowLimit() bool {
	return true
}

func (m *_BACnetConstructedDataIntegerValueLowLimit) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
