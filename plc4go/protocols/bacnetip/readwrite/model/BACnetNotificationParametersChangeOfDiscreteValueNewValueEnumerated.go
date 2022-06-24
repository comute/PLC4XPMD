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

// BACnetNotificationParametersChangeOfDiscreteValueNewValueEnumerated is the corresponding interface of BACnetNotificationParametersChangeOfDiscreteValueNewValueEnumerated
type BACnetNotificationParametersChangeOfDiscreteValueNewValueEnumerated interface {
	utils.LengthAware
	utils.Serializable
	BACnetNotificationParametersChangeOfDiscreteValueNewValue
	// GetEnumeratedValue returns EnumeratedValue (property field)
	GetEnumeratedValue() BACnetApplicationTagEnumerated
}

// BACnetNotificationParametersChangeOfDiscreteValueNewValueEnumeratedExactly can be used when we want exactly this type and not a type which fulfills BACnetNotificationParametersChangeOfDiscreteValueNewValueEnumerated.
// This is useful for switch cases.
type BACnetNotificationParametersChangeOfDiscreteValueNewValueEnumeratedExactly interface {
	BACnetNotificationParametersChangeOfDiscreteValueNewValueEnumerated
	isBACnetNotificationParametersChangeOfDiscreteValueNewValueEnumerated() bool
}

// _BACnetNotificationParametersChangeOfDiscreteValueNewValueEnumerated is the data-structure of this message
type _BACnetNotificationParametersChangeOfDiscreteValueNewValueEnumerated struct {
	*_BACnetNotificationParametersChangeOfDiscreteValueNewValue
	EnumeratedValue BACnetApplicationTagEnumerated
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueEnumerated) InitializeParent(parent BACnetNotificationParametersChangeOfDiscreteValueNewValue, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueEnumerated) GetParent() BACnetNotificationParametersChangeOfDiscreteValueNewValue {
	return m._BACnetNotificationParametersChangeOfDiscreteValueNewValue
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueEnumerated) GetEnumeratedValue() BACnetApplicationTagEnumerated {
	return m.EnumeratedValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetNotificationParametersChangeOfDiscreteValueNewValueEnumerated factory function for _BACnetNotificationParametersChangeOfDiscreteValueNewValueEnumerated
func NewBACnetNotificationParametersChangeOfDiscreteValueNewValueEnumerated(enumeratedValue BACnetApplicationTagEnumerated, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetNotificationParametersChangeOfDiscreteValueNewValueEnumerated {
	_result := &_BACnetNotificationParametersChangeOfDiscreteValueNewValueEnumerated{
		EnumeratedValue: enumeratedValue,
		_BACnetNotificationParametersChangeOfDiscreteValueNewValue: NewBACnetNotificationParametersChangeOfDiscreteValueNewValue(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result._BACnetNotificationParametersChangeOfDiscreteValueNewValue._BACnetNotificationParametersChangeOfDiscreteValueNewValueChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetNotificationParametersChangeOfDiscreteValueNewValueEnumerated(structType interface{}) BACnetNotificationParametersChangeOfDiscreteValueNewValueEnumerated {
	if casted, ok := structType.(BACnetNotificationParametersChangeOfDiscreteValueNewValueEnumerated); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetNotificationParametersChangeOfDiscreteValueNewValueEnumerated); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueEnumerated) GetTypeName() string {
	return "BACnetNotificationParametersChangeOfDiscreteValueNewValueEnumerated"
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueEnumerated) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueEnumerated) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (enumeratedValue)
	lengthInBits += m.EnumeratedValue.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueEnumerated) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetNotificationParametersChangeOfDiscreteValueNewValueEnumeratedParse(readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetNotificationParametersChangeOfDiscreteValueNewValueEnumerated, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetNotificationParametersChangeOfDiscreteValueNewValueEnumerated"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetNotificationParametersChangeOfDiscreteValueNewValueEnumerated")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (enumeratedValue)
	if pullErr := readBuffer.PullContext("enumeratedValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for enumeratedValue")
	}
	_enumeratedValue, _enumeratedValueErr := BACnetApplicationTagParse(readBuffer)
	if _enumeratedValueErr != nil {
		return nil, errors.Wrap(_enumeratedValueErr, "Error parsing 'enumeratedValue' field")
	}
	enumeratedValue := _enumeratedValue.(BACnetApplicationTagEnumerated)
	if closeErr := readBuffer.CloseContext("enumeratedValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for enumeratedValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetNotificationParametersChangeOfDiscreteValueNewValueEnumerated"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetNotificationParametersChangeOfDiscreteValueNewValueEnumerated")
	}

	// Create a partially initialized instance
	_child := &_BACnetNotificationParametersChangeOfDiscreteValueNewValueEnumerated{
		EnumeratedValue: enumeratedValue,
		_BACnetNotificationParametersChangeOfDiscreteValueNewValue: &_BACnetNotificationParametersChangeOfDiscreteValueNewValue{
			TagNumber: tagNumber,
		},
	}
	_child._BACnetNotificationParametersChangeOfDiscreteValueNewValue._BACnetNotificationParametersChangeOfDiscreteValueNewValueChildRequirements = _child
	return _child, nil
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueEnumerated) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetNotificationParametersChangeOfDiscreteValueNewValueEnumerated"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetNotificationParametersChangeOfDiscreteValueNewValueEnumerated")
		}

		// Simple Field (enumeratedValue)
		if pushErr := writeBuffer.PushContext("enumeratedValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for enumeratedValue")
		}
		_enumeratedValueErr := writeBuffer.WriteSerializable(m.GetEnumeratedValue())
		if popErr := writeBuffer.PopContext("enumeratedValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for enumeratedValue")
		}
		if _enumeratedValueErr != nil {
			return errors.Wrap(_enumeratedValueErr, "Error serializing 'enumeratedValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetNotificationParametersChangeOfDiscreteValueNewValueEnumerated"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetNotificationParametersChangeOfDiscreteValueNewValueEnumerated")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueEnumerated) isBACnetNotificationParametersChangeOfDiscreteValueNewValueEnumerated() bool {
	return true
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueEnumerated) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
