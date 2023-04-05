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

// BACnetConstructedDataBinaryOutputPresentValue is the corresponding interface of BACnetConstructedDataBinaryOutputPresentValue
type BACnetConstructedDataBinaryOutputPresentValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetPresentValue returns PresentValue (property field)
	GetPresentValue() BACnetBinaryPVTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetBinaryPVTagged
}

// BACnetConstructedDataBinaryOutputPresentValueExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataBinaryOutputPresentValue.
// This is useful for switch cases.
type BACnetConstructedDataBinaryOutputPresentValueExactly interface {
	BACnetConstructedDataBinaryOutputPresentValue
	isBACnetConstructedDataBinaryOutputPresentValue() bool
}

// _BACnetConstructedDataBinaryOutputPresentValue is the data-structure of this message
type _BACnetConstructedDataBinaryOutputPresentValue struct {
	*_BACnetConstructedData
	PresentValue BACnetBinaryPVTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataBinaryOutputPresentValue) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_BINARY_OUTPUT
}

func (m *_BACnetConstructedDataBinaryOutputPresentValue) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_PRESENT_VALUE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataBinaryOutputPresentValue) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataBinaryOutputPresentValue) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataBinaryOutputPresentValue) GetPresentValue() BACnetBinaryPVTagged {
	return m.PresentValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataBinaryOutputPresentValue) GetActualValue() BACnetBinaryPVTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetBinaryPVTagged(m.GetPresentValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataBinaryOutputPresentValue factory function for _BACnetConstructedDataBinaryOutputPresentValue
func NewBACnetConstructedDataBinaryOutputPresentValue(presentValue BACnetBinaryPVTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataBinaryOutputPresentValue {
	_result := &_BACnetConstructedDataBinaryOutputPresentValue{
		PresentValue:           presentValue,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataBinaryOutputPresentValue(structType interface{}) BACnetConstructedDataBinaryOutputPresentValue {
	if casted, ok := structType.(BACnetConstructedDataBinaryOutputPresentValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataBinaryOutputPresentValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataBinaryOutputPresentValue) GetTypeName() string {
	return "BACnetConstructedDataBinaryOutputPresentValue"
}

func (m *_BACnetConstructedDataBinaryOutputPresentValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (presentValue)
	lengthInBits += m.PresentValue.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataBinaryOutputPresentValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataBinaryOutputPresentValueParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataBinaryOutputPresentValue, error) {
	return BACnetConstructedDataBinaryOutputPresentValueParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataBinaryOutputPresentValueParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataBinaryOutputPresentValue, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataBinaryOutputPresentValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataBinaryOutputPresentValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (presentValue)
	if pullErr := readBuffer.PullContext("presentValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for presentValue")
	}
	_presentValue, _presentValueErr := BACnetBinaryPVTaggedParseWithBuffer(ctx, readBuffer, uint8(uint8(0)), TagClass(TagClass_APPLICATION_TAGS))
	if _presentValueErr != nil {
		return nil, errors.Wrap(_presentValueErr, "Error parsing 'presentValue' field of BACnetConstructedDataBinaryOutputPresentValue")
	}
	presentValue := _presentValue.(BACnetBinaryPVTagged)
	if closeErr := readBuffer.CloseContext("presentValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for presentValue")
	}

	// Virtual field
	_actualValue := presentValue
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataBinaryOutputPresentValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataBinaryOutputPresentValue")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataBinaryOutputPresentValue{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		PresentValue: presentValue,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataBinaryOutputPresentValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataBinaryOutputPresentValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataBinaryOutputPresentValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataBinaryOutputPresentValue")
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
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataBinaryOutputPresentValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataBinaryOutputPresentValue")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataBinaryOutputPresentValue) isBACnetConstructedDataBinaryOutputPresentValue() bool {
	return true
}

func (m *_BACnetConstructedDataBinaryOutputPresentValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
