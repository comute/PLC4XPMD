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
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataIntegerValueMinPresValue is the corresponding interface of BACnetConstructedDataIntegerValueMinPresValue
type BACnetConstructedDataIntegerValueMinPresValue interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetMinPresValue returns MinPresValue (property field)
	GetMinPresValue() BACnetApplicationTagSignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagSignedInteger
}

// BACnetConstructedDataIntegerValueMinPresValueExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataIntegerValueMinPresValue.
// This is useful for switch cases.
type BACnetConstructedDataIntegerValueMinPresValueExactly interface {
	BACnetConstructedDataIntegerValueMinPresValue
	isBACnetConstructedDataIntegerValueMinPresValue() bool
}

// _BACnetConstructedDataIntegerValueMinPresValue is the data-structure of this message
type _BACnetConstructedDataIntegerValueMinPresValue struct {
	*_BACnetConstructedData
	MinPresValue BACnetApplicationTagSignedInteger

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument BACnetTagPayloadUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataIntegerValueMinPresValue) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_INTEGER_VALUE
}

func (m *_BACnetConstructedDataIntegerValueMinPresValue) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_MIN_PRES_VALUE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataIntegerValueMinPresValue) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataIntegerValueMinPresValue) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataIntegerValueMinPresValue) GetMinPresValue() BACnetApplicationTagSignedInteger {
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

func (m *_BACnetConstructedDataIntegerValueMinPresValue) GetActualValue() BACnetApplicationTagSignedInteger {
	return CastBACnetApplicationTagSignedInteger(m.GetMinPresValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataIntegerValueMinPresValue factory function for _BACnetConstructedDataIntegerValueMinPresValue
func NewBACnetConstructedDataIntegerValueMinPresValue(minPresValue BACnetApplicationTagSignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataIntegerValueMinPresValue {
	_result := &_BACnetConstructedDataIntegerValueMinPresValue{
		MinPresValue:           minPresValue,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataIntegerValueMinPresValue(structType interface{}) BACnetConstructedDataIntegerValueMinPresValue {
	if casted, ok := structType.(BACnetConstructedDataIntegerValueMinPresValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataIntegerValueMinPresValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataIntegerValueMinPresValue) GetTypeName() string {
	return "BACnetConstructedDataIntegerValueMinPresValue"
}

func (m *_BACnetConstructedDataIntegerValueMinPresValue) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataIntegerValueMinPresValue) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (minPresValue)
	lengthInBits += m.MinPresValue.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataIntegerValueMinPresValue) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataIntegerValueMinPresValueParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataIntegerValueMinPresValue, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataIntegerValueMinPresValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataIntegerValueMinPresValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (minPresValue)
	if pullErr := readBuffer.PullContext("minPresValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for minPresValue")
	}
	_minPresValue, _minPresValueErr := BACnetApplicationTagParse(readBuffer)
	if _minPresValueErr != nil {
		return nil, errors.Wrap(_minPresValueErr, "Error parsing 'minPresValue' field")
	}
	minPresValue := _minPresValue.(BACnetApplicationTagSignedInteger)
	if closeErr := readBuffer.CloseContext("minPresValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for minPresValue")
	}

	// Virtual field
	_actualValue := minPresValue
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataIntegerValueMinPresValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataIntegerValueMinPresValue")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataIntegerValueMinPresValue{
		MinPresValue:           minPresValue,
		_BACnetConstructedData: &_BACnetConstructedData{},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataIntegerValueMinPresValue) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataIntegerValueMinPresValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataIntegerValueMinPresValue")
		}

		// Simple Field (minPresValue)
		if pushErr := writeBuffer.PushContext("minPresValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for minPresValue")
		}
		_minPresValueErr := writeBuffer.WriteSerializable(m.GetMinPresValue())
		if popErr := writeBuffer.PopContext("minPresValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for minPresValue")
		}
		if _minPresValueErr != nil {
			return errors.Wrap(_minPresValueErr, "Error serializing 'minPresValue' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataIntegerValueMinPresValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataIntegerValueMinPresValue")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataIntegerValueMinPresValue) isBACnetConstructedDataIntegerValueMinPresValue() bool {
	return true
}

func (m *_BACnetConstructedDataIntegerValueMinPresValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
