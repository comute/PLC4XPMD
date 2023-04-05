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

// BACnetConstructedDataAccumulatorMinPresValue is the corresponding interface of BACnetConstructedDataAccumulatorMinPresValue
type BACnetConstructedDataAccumulatorMinPresValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetMinPresValue returns MinPresValue (property field)
	GetMinPresValue() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataAccumulatorMinPresValueExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataAccumulatorMinPresValue.
// This is useful for switch cases.
type BACnetConstructedDataAccumulatorMinPresValueExactly interface {
	BACnetConstructedDataAccumulatorMinPresValue
	isBACnetConstructedDataAccumulatorMinPresValue() bool
}

// _BACnetConstructedDataAccumulatorMinPresValue is the data-structure of this message
type _BACnetConstructedDataAccumulatorMinPresValue struct {
	*_BACnetConstructedData
	MinPresValue BACnetApplicationTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAccumulatorMinPresValue) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_ACCUMULATOR
}

func (m *_BACnetConstructedDataAccumulatorMinPresValue) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_MIN_PRES_VALUE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAccumulatorMinPresValue) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataAccumulatorMinPresValue) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataAccumulatorMinPresValue) GetMinPresValue() BACnetApplicationTagUnsignedInteger {
	return m.MinPresValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataAccumulatorMinPresValue) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetMinPresValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataAccumulatorMinPresValue factory function for _BACnetConstructedDataAccumulatorMinPresValue
func NewBACnetConstructedDataAccumulatorMinPresValue(minPresValue BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAccumulatorMinPresValue {
	_result := &_BACnetConstructedDataAccumulatorMinPresValue{
		MinPresValue:           minPresValue,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAccumulatorMinPresValue(structType interface{}) BACnetConstructedDataAccumulatorMinPresValue {
	if casted, ok := structType.(BACnetConstructedDataAccumulatorMinPresValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAccumulatorMinPresValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAccumulatorMinPresValue) GetTypeName() string {
	return "BACnetConstructedDataAccumulatorMinPresValue"
}

func (m *_BACnetConstructedDataAccumulatorMinPresValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (minPresValue)
	lengthInBits += m.MinPresValue.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataAccumulatorMinPresValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataAccumulatorMinPresValueParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataAccumulatorMinPresValue, error) {
	return BACnetConstructedDataAccumulatorMinPresValueParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataAccumulatorMinPresValueParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataAccumulatorMinPresValue, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAccumulatorMinPresValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAccumulatorMinPresValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (minPresValue)
	if pullErr := readBuffer.PullContext("minPresValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for minPresValue")
	}
	_minPresValue, _minPresValueErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _minPresValueErr != nil {
		return nil, errors.Wrap(_minPresValueErr, "Error parsing 'minPresValue' field of BACnetConstructedDataAccumulatorMinPresValue")
	}
	minPresValue := _minPresValue.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("minPresValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for minPresValue")
	}

	// Virtual field
	_actualValue := minPresValue
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAccumulatorMinPresValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAccumulatorMinPresValue")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataAccumulatorMinPresValue{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		MinPresValue: minPresValue,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataAccumulatorMinPresValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataAccumulatorMinPresValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAccumulatorMinPresValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAccumulatorMinPresValue")
		}

		// Simple Field (minPresValue)
		if pushErr := writeBuffer.PushContext("minPresValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for minPresValue")
		}
		_minPresValueErr := writeBuffer.WriteSerializable(ctx, m.GetMinPresValue())
		if popErr := writeBuffer.PopContext("minPresValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for minPresValue")
		}
		if _minPresValueErr != nil {
			return errors.Wrap(_minPresValueErr, "Error serializing 'minPresValue' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAccumulatorMinPresValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAccumulatorMinPresValue")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAccumulatorMinPresValue) isBACnetConstructedDataAccumulatorMinPresValue() bool {
	return true
}

func (m *_BACnetConstructedDataAccumulatorMinPresValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
