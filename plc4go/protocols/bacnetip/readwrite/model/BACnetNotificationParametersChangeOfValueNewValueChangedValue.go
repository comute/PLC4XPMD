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

// BACnetNotificationParametersChangeOfValueNewValueChangedValue is the corresponding interface of BACnetNotificationParametersChangeOfValueNewValueChangedValue
type BACnetNotificationParametersChangeOfValueNewValueChangedValue interface {
	utils.LengthAware
	utils.Serializable
	BACnetNotificationParametersChangeOfValueNewValue
	// GetChangedValue returns ChangedValue (property field)
	GetChangedValue() BACnetContextTagReal
}

// BACnetNotificationParametersChangeOfValueNewValueChangedValueExactly can be used when we want exactly this type and not a type which fulfills BACnetNotificationParametersChangeOfValueNewValueChangedValue.
// This is useful for switch cases.
type BACnetNotificationParametersChangeOfValueNewValueChangedValueExactly interface {
	BACnetNotificationParametersChangeOfValueNewValueChangedValue
	isBACnetNotificationParametersChangeOfValueNewValueChangedValue() bool
}

// _BACnetNotificationParametersChangeOfValueNewValueChangedValue is the data-structure of this message
type _BACnetNotificationParametersChangeOfValueNewValueChangedValue struct {
	*_BACnetNotificationParametersChangeOfValueNewValue
	ChangedValue BACnetContextTagReal

	// Arguments.
	TagNumber uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetNotificationParametersChangeOfValueNewValueChangedValue) InitializeParent(parent BACnetNotificationParametersChangeOfValueNewValue, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetNotificationParametersChangeOfValueNewValueChangedValue) GetParent() BACnetNotificationParametersChangeOfValueNewValue {
	return m._BACnetNotificationParametersChangeOfValueNewValue
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetNotificationParametersChangeOfValueNewValueChangedValue) GetChangedValue() BACnetContextTagReal {
	return m.ChangedValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetNotificationParametersChangeOfValueNewValueChangedValue factory function for _BACnetNotificationParametersChangeOfValueNewValueChangedValue
func NewBACnetNotificationParametersChangeOfValueNewValueChangedValue(changedValue BACnetContextTagReal, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetNotificationParametersChangeOfValueNewValueChangedValue {
	_result := &_BACnetNotificationParametersChangeOfValueNewValueChangedValue{
		ChangedValue: changedValue,
		_BACnetNotificationParametersChangeOfValueNewValue: NewBACnetNotificationParametersChangeOfValueNewValue(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result._BACnetNotificationParametersChangeOfValueNewValue._BACnetNotificationParametersChangeOfValueNewValueChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetNotificationParametersChangeOfValueNewValueChangedValue(structType interface{}) BACnetNotificationParametersChangeOfValueNewValueChangedValue {
	if casted, ok := structType.(BACnetNotificationParametersChangeOfValueNewValueChangedValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetNotificationParametersChangeOfValueNewValueChangedValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetNotificationParametersChangeOfValueNewValueChangedValue) GetTypeName() string {
	return "BACnetNotificationParametersChangeOfValueNewValueChangedValue"
}

func (m *_BACnetNotificationParametersChangeOfValueNewValueChangedValue) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetNotificationParametersChangeOfValueNewValueChangedValue) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (changedValue)
	lengthInBits += m.ChangedValue.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetNotificationParametersChangeOfValueNewValueChangedValue) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetNotificationParametersChangeOfValueNewValueChangedValueParse(readBuffer utils.ReadBuffer, tagNumber uint8, peekedTagNumber uint8) (BACnetNotificationParametersChangeOfValueNewValueChangedValue, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetNotificationParametersChangeOfValueNewValueChangedValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetNotificationParametersChangeOfValueNewValueChangedValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (changedValue)
	if pullErr := readBuffer.PullContext("changedValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for changedValue")
	}
	_changedValue, _changedValueErr := BACnetContextTagParse(readBuffer, uint8(uint8(1)), BACnetDataType(BACnetDataType_REAL))
	if _changedValueErr != nil {
		return nil, errors.Wrap(_changedValueErr, "Error parsing 'changedValue' field")
	}
	changedValue := _changedValue.(BACnetContextTagReal)
	if closeErr := readBuffer.CloseContext("changedValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for changedValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetNotificationParametersChangeOfValueNewValueChangedValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetNotificationParametersChangeOfValueNewValueChangedValue")
	}

	// Create a partially initialized instance
	_child := &_BACnetNotificationParametersChangeOfValueNewValueChangedValue{
		ChangedValue: changedValue,
		_BACnetNotificationParametersChangeOfValueNewValue: &_BACnetNotificationParametersChangeOfValueNewValue{},
	}
	_child._BACnetNotificationParametersChangeOfValueNewValue._BACnetNotificationParametersChangeOfValueNewValueChildRequirements = _child
	return _child, nil
}

func (m *_BACnetNotificationParametersChangeOfValueNewValueChangedValue) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetNotificationParametersChangeOfValueNewValueChangedValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetNotificationParametersChangeOfValueNewValueChangedValue")
		}

		// Simple Field (changedValue)
		if pushErr := writeBuffer.PushContext("changedValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for changedValue")
		}
		_changedValueErr := writeBuffer.WriteSerializable(m.GetChangedValue())
		if popErr := writeBuffer.PopContext("changedValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for changedValue")
		}
		if _changedValueErr != nil {
			return errors.Wrap(_changedValueErr, "Error serializing 'changedValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetNotificationParametersChangeOfValueNewValueChangedValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetNotificationParametersChangeOfValueNewValueChangedValue")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetNotificationParametersChangeOfValueNewValueChangedValue) isBACnetNotificationParametersChangeOfValueNewValueChangedValue() bool {
	return true
}

func (m *_BACnetNotificationParametersChangeOfValueNewValueChangedValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
