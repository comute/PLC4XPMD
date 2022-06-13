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

// BACnetAssignedLandingCallsLandingCallsListEntry is the data-structure of this message
type BACnetAssignedLandingCallsLandingCallsListEntry struct {
	FloorNumber *BACnetContextTagUnsignedInteger
	Direction   *BACnetLiftCarDirectionTagged
}

// IBACnetAssignedLandingCallsLandingCallsListEntry is the corresponding interface of BACnetAssignedLandingCallsLandingCallsListEntry
type IBACnetAssignedLandingCallsLandingCallsListEntry interface {
	// GetFloorNumber returns FloorNumber (property field)
	GetFloorNumber() *BACnetContextTagUnsignedInteger
	// GetDirection returns Direction (property field)
	GetDirection() *BACnetLiftCarDirectionTagged
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetAssignedLandingCallsLandingCallsListEntry) GetFloorNumber() *BACnetContextTagUnsignedInteger {
	return m.FloorNumber
}

func (m *BACnetAssignedLandingCallsLandingCallsListEntry) GetDirection() *BACnetLiftCarDirectionTagged {
	return m.Direction
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetAssignedLandingCallsLandingCallsListEntry factory function for BACnetAssignedLandingCallsLandingCallsListEntry
func NewBACnetAssignedLandingCallsLandingCallsListEntry(floorNumber *BACnetContextTagUnsignedInteger, direction *BACnetLiftCarDirectionTagged) *BACnetAssignedLandingCallsLandingCallsListEntry {
	return &BACnetAssignedLandingCallsLandingCallsListEntry{FloorNumber: floorNumber, Direction: direction}
}

func CastBACnetAssignedLandingCallsLandingCallsListEntry(structType interface{}) *BACnetAssignedLandingCallsLandingCallsListEntry {
	if casted, ok := structType.(BACnetAssignedLandingCallsLandingCallsListEntry); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetAssignedLandingCallsLandingCallsListEntry); ok {
		return casted
	}
	return nil
}

func (m *BACnetAssignedLandingCallsLandingCallsListEntry) GetTypeName() string {
	return "BACnetAssignedLandingCallsLandingCallsListEntry"
}

func (m *BACnetAssignedLandingCallsLandingCallsListEntry) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetAssignedLandingCallsLandingCallsListEntry) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (floorNumber)
	lengthInBits += m.FloorNumber.GetLengthInBits()

	// Simple field (direction)
	lengthInBits += m.Direction.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetAssignedLandingCallsLandingCallsListEntry) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetAssignedLandingCallsLandingCallsListEntryParse(readBuffer utils.ReadBuffer) (*BACnetAssignedLandingCallsLandingCallsListEntry, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetAssignedLandingCallsLandingCallsListEntry"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetAssignedLandingCallsLandingCallsListEntry")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (floorNumber)
	if pullErr := readBuffer.PullContext("floorNumber"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for floorNumber")
	}
	_floorNumber, _floorNumberErr := BACnetContextTagParse(readBuffer, uint8(uint8(0)), BACnetDataType(BACnetDataType_UNSIGNED_INTEGER))
	if _floorNumberErr != nil {
		return nil, errors.Wrap(_floorNumberErr, "Error parsing 'floorNumber' field")
	}
	floorNumber := CastBACnetContextTagUnsignedInteger(_floorNumber)
	if closeErr := readBuffer.CloseContext("floorNumber"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for floorNumber")
	}

	// Simple Field (direction)
	if pullErr := readBuffer.PullContext("direction"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for direction")
	}
	_direction, _directionErr := BACnetLiftCarDirectionTaggedParse(readBuffer, uint8(uint8(1)), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _directionErr != nil {
		return nil, errors.Wrap(_directionErr, "Error parsing 'direction' field")
	}
	direction := CastBACnetLiftCarDirectionTagged(_direction)
	if closeErr := readBuffer.CloseContext("direction"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for direction")
	}

	if closeErr := readBuffer.CloseContext("BACnetAssignedLandingCallsLandingCallsListEntry"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetAssignedLandingCallsLandingCallsListEntry")
	}

	// Create the instance
	return NewBACnetAssignedLandingCallsLandingCallsListEntry(floorNumber, direction), nil
}

func (m *BACnetAssignedLandingCallsLandingCallsListEntry) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetAssignedLandingCallsLandingCallsListEntry"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetAssignedLandingCallsLandingCallsListEntry")
	}

	// Simple Field (floorNumber)
	if pushErr := writeBuffer.PushContext("floorNumber"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for floorNumber")
	}
	_floorNumberErr := writeBuffer.WriteSerializable(m.FloorNumber)
	if popErr := writeBuffer.PopContext("floorNumber"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for floorNumber")
	}
	if _floorNumberErr != nil {
		return errors.Wrap(_floorNumberErr, "Error serializing 'floorNumber' field")
	}

	// Simple Field (direction)
	if pushErr := writeBuffer.PushContext("direction"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for direction")
	}
	_directionErr := writeBuffer.WriteSerializable(m.Direction)
	if popErr := writeBuffer.PopContext("direction"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for direction")
	}
	if _directionErr != nil {
		return errors.Wrap(_directionErr, "Error serializing 'direction' field")
	}

	if popErr := writeBuffer.PopContext("BACnetAssignedLandingCallsLandingCallsListEntry"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetAssignedLandingCallsLandingCallsListEntry")
	}
	return nil
}

func (m *BACnetAssignedLandingCallsLandingCallsListEntry) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
