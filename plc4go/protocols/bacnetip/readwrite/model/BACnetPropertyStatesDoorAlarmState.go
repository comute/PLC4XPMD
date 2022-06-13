/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
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

// BACnetPropertyStatesDoorAlarmState is the data-structure of this message
type BACnetPropertyStatesDoorAlarmState struct {
	*BACnetPropertyStates
	DoorAlarmState *BACnetDoorAlarmStateTagged
}

// IBACnetPropertyStatesDoorAlarmState is the corresponding interface of BACnetPropertyStatesDoorAlarmState
type IBACnetPropertyStatesDoorAlarmState interface {
	IBACnetPropertyStates
	// GetDoorAlarmState returns DoorAlarmState (property field)
	GetDoorAlarmState() *BACnetDoorAlarmStateTagged
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetPropertyStatesDoorAlarmState) InitializeParent(parent *BACnetPropertyStates, peekedTagHeader *BACnetTagHeader) {
	m.BACnetPropertyStates.PeekedTagHeader = peekedTagHeader
}

func (m *BACnetPropertyStatesDoorAlarmState) GetParent() *BACnetPropertyStates {
	return m.BACnetPropertyStates
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetPropertyStatesDoorAlarmState) GetDoorAlarmState() *BACnetDoorAlarmStateTagged {
	return m.DoorAlarmState
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStatesDoorAlarmState factory function for BACnetPropertyStatesDoorAlarmState
func NewBACnetPropertyStatesDoorAlarmState(doorAlarmState *BACnetDoorAlarmStateTagged, peekedTagHeader *BACnetTagHeader) *BACnetPropertyStatesDoorAlarmState {
	_result := &BACnetPropertyStatesDoorAlarmState{
		DoorAlarmState:       doorAlarmState,
		BACnetPropertyStates: NewBACnetPropertyStates(peekedTagHeader),
	}
	_result.Child = _result
	return _result
}

func CastBACnetPropertyStatesDoorAlarmState(structType interface{}) *BACnetPropertyStatesDoorAlarmState {
	if casted, ok := structType.(BACnetPropertyStatesDoorAlarmState); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesDoorAlarmState); ok {
		return casted
	}
	if casted, ok := structType.(BACnetPropertyStates); ok {
		return CastBACnetPropertyStatesDoorAlarmState(casted.Child)
	}
	if casted, ok := structType.(*BACnetPropertyStates); ok {
		return CastBACnetPropertyStatesDoorAlarmState(casted.Child)
	}
	return nil
}

func (m *BACnetPropertyStatesDoorAlarmState) GetTypeName() string {
	return "BACnetPropertyStatesDoorAlarmState"
}

func (m *BACnetPropertyStatesDoorAlarmState) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetPropertyStatesDoorAlarmState) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (doorAlarmState)
	lengthInBits += m.DoorAlarmState.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetPropertyStatesDoorAlarmState) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetPropertyStatesDoorAlarmStateParse(readBuffer utils.ReadBuffer, peekedTagNumber uint8) (*BACnetPropertyStatesDoorAlarmState, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesDoorAlarmState"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesDoorAlarmState")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (doorAlarmState)
	if pullErr := readBuffer.PullContext("doorAlarmState"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for doorAlarmState")
	}
	_doorAlarmState, _doorAlarmStateErr := BACnetDoorAlarmStateTaggedParse(readBuffer, uint8(peekedTagNumber), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _doorAlarmStateErr != nil {
		return nil, errors.Wrap(_doorAlarmStateErr, "Error parsing 'doorAlarmState' field")
	}
	doorAlarmState := CastBACnetDoorAlarmStateTagged(_doorAlarmState)
	if closeErr := readBuffer.CloseContext("doorAlarmState"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for doorAlarmState")
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesDoorAlarmState"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesDoorAlarmState")
	}

	// Create a partially initialized instance
	_child := &BACnetPropertyStatesDoorAlarmState{
		DoorAlarmState:       CastBACnetDoorAlarmStateTagged(doorAlarmState),
		BACnetPropertyStates: &BACnetPropertyStates{},
	}
	_child.BACnetPropertyStates.Child = _child
	return _child, nil
}

func (m *BACnetPropertyStatesDoorAlarmState) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesDoorAlarmState"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesDoorAlarmState")
		}

		// Simple Field (doorAlarmState)
		if pushErr := writeBuffer.PushContext("doorAlarmState"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for doorAlarmState")
		}
		_doorAlarmStateErr := writeBuffer.WriteSerializable(m.DoorAlarmState)
		if popErr := writeBuffer.PopContext("doorAlarmState"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for doorAlarmState")
		}
		if _doorAlarmStateErr != nil {
			return errors.Wrap(_doorAlarmStateErr, "Error serializing 'doorAlarmState' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesDoorAlarmState"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesDoorAlarmState")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetPropertyStatesDoorAlarmState) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
