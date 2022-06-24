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

// BACnetTimerStateChangeValueTime is the corresponding interface of BACnetTimerStateChangeValueTime
type BACnetTimerStateChangeValueTime interface {
	utils.LengthAware
	utils.Serializable
	BACnetTimerStateChangeValue
	// GetTimeValue returns TimeValue (property field)
	GetTimeValue() BACnetApplicationTagTime
}

// BACnetTimerStateChangeValueTimeExactly can be used when we want exactly this type and not a type which fulfills BACnetTimerStateChangeValueTime.
// This is useful for switch cases.
type BACnetTimerStateChangeValueTimeExactly interface {
	BACnetTimerStateChangeValueTime
	isBACnetTimerStateChangeValueTime() bool
}

// _BACnetTimerStateChangeValueTime is the data-structure of this message
type _BACnetTimerStateChangeValueTime struct {
	*_BACnetTimerStateChangeValue
	TimeValue BACnetApplicationTagTime
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetTimerStateChangeValueTime) InitializeParent(parent BACnetTimerStateChangeValue, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetTimerStateChangeValueTime) GetParent() BACnetTimerStateChangeValue {
	return m._BACnetTimerStateChangeValue
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetTimerStateChangeValueTime) GetTimeValue() BACnetApplicationTagTime {
	return m.TimeValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetTimerStateChangeValueTime factory function for _BACnetTimerStateChangeValueTime
func NewBACnetTimerStateChangeValueTime(timeValue BACnetApplicationTagTime, peekedTagHeader BACnetTagHeader, objectTypeArgument BACnetObjectType) *_BACnetTimerStateChangeValueTime {
	_result := &_BACnetTimerStateChangeValueTime{
		TimeValue:                    timeValue,
		_BACnetTimerStateChangeValue: NewBACnetTimerStateChangeValue(peekedTagHeader, objectTypeArgument),
	}
	_result._BACnetTimerStateChangeValue._BACnetTimerStateChangeValueChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetTimerStateChangeValueTime(structType interface{}) BACnetTimerStateChangeValueTime {
	if casted, ok := structType.(BACnetTimerStateChangeValueTime); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetTimerStateChangeValueTime); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetTimerStateChangeValueTime) GetTypeName() string {
	return "BACnetTimerStateChangeValueTime"
}

func (m *_BACnetTimerStateChangeValueTime) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetTimerStateChangeValueTime) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (timeValue)
	lengthInBits += m.TimeValue.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetTimerStateChangeValueTime) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetTimerStateChangeValueTimeParse(readBuffer utils.ReadBuffer, objectTypeArgument BACnetObjectType) (BACnetTimerStateChangeValueTime, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetTimerStateChangeValueTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetTimerStateChangeValueTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (timeValue)
	if pullErr := readBuffer.PullContext("timeValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for timeValue")
	}
	_timeValue, _timeValueErr := BACnetApplicationTagParse(readBuffer)
	if _timeValueErr != nil {
		return nil, errors.Wrap(_timeValueErr, "Error parsing 'timeValue' field")
	}
	timeValue := _timeValue.(BACnetApplicationTagTime)
	if closeErr := readBuffer.CloseContext("timeValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for timeValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetTimerStateChangeValueTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetTimerStateChangeValueTime")
	}

	// Create a partially initialized instance
	_child := &_BACnetTimerStateChangeValueTime{
		TimeValue: timeValue,
		_BACnetTimerStateChangeValue: &_BACnetTimerStateChangeValue{
			ObjectTypeArgument: objectTypeArgument,
		},
	}
	_child._BACnetTimerStateChangeValue._BACnetTimerStateChangeValueChildRequirements = _child
	return _child, nil
}

func (m *_BACnetTimerStateChangeValueTime) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetTimerStateChangeValueTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetTimerStateChangeValueTime")
		}

		// Simple Field (timeValue)
		if pushErr := writeBuffer.PushContext("timeValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for timeValue")
		}
		_timeValueErr := writeBuffer.WriteSerializable(m.GetTimeValue())
		if popErr := writeBuffer.PopContext("timeValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for timeValue")
		}
		if _timeValueErr != nil {
			return errors.Wrap(_timeValueErr, "Error serializing 'timeValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetTimerStateChangeValueTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetTimerStateChangeValueTime")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetTimerStateChangeValueTime) isBACnetTimerStateChangeValueTime() bool {
	return true
}

func (m *_BACnetTimerStateChangeValueTime) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
