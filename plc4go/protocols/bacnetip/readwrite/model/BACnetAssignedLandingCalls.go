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

// BACnetAssignedLandingCalls is the data-structure of this message
type BACnetAssignedLandingCalls struct {
	LandingCalls *BACnetAssignedLandingCallsLandingCallsList
}

// IBACnetAssignedLandingCalls is the corresponding interface of BACnetAssignedLandingCalls
type IBACnetAssignedLandingCalls interface {
	// GetLandingCalls returns LandingCalls (property field)
	GetLandingCalls() *BACnetAssignedLandingCallsLandingCallsList
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

func (m *BACnetAssignedLandingCalls) GetLandingCalls() *BACnetAssignedLandingCallsLandingCallsList {
	return m.LandingCalls
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetAssignedLandingCalls factory function for BACnetAssignedLandingCalls
func NewBACnetAssignedLandingCalls(landingCalls *BACnetAssignedLandingCallsLandingCallsList) *BACnetAssignedLandingCalls {
	return &BACnetAssignedLandingCalls{LandingCalls: landingCalls}
}

func CastBACnetAssignedLandingCalls(structType interface{}) *BACnetAssignedLandingCalls {
	if casted, ok := structType.(BACnetAssignedLandingCalls); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetAssignedLandingCalls); ok {
		return casted
	}
	return nil
}

func (m *BACnetAssignedLandingCalls) GetTypeName() string {
	return "BACnetAssignedLandingCalls"
}

func (m *BACnetAssignedLandingCalls) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetAssignedLandingCalls) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (landingCalls)
	lengthInBits += m.LandingCalls.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetAssignedLandingCalls) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetAssignedLandingCallsParse(readBuffer utils.ReadBuffer) (*BACnetAssignedLandingCalls, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetAssignedLandingCalls"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetAssignedLandingCalls")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (landingCalls)
	if pullErr := readBuffer.PullContext("landingCalls"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for landingCalls")
	}
	_landingCalls, _landingCallsErr := BACnetAssignedLandingCallsLandingCallsListParse(readBuffer, uint8(uint8(0)))
	if _landingCallsErr != nil {
		return nil, errors.Wrap(_landingCallsErr, "Error parsing 'landingCalls' field")
	}
	landingCalls := CastBACnetAssignedLandingCallsLandingCallsList(_landingCalls)
	if closeErr := readBuffer.CloseContext("landingCalls"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for landingCalls")
	}

	if closeErr := readBuffer.CloseContext("BACnetAssignedLandingCalls"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetAssignedLandingCalls")
	}

	// Create the instance
	return NewBACnetAssignedLandingCalls(landingCalls), nil
}

func (m *BACnetAssignedLandingCalls) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetAssignedLandingCalls"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetAssignedLandingCalls")
	}

	// Simple Field (landingCalls)
	if pushErr := writeBuffer.PushContext("landingCalls"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for landingCalls")
	}
	_landingCallsErr := m.LandingCalls.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("landingCalls"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for landingCalls")
	}
	if _landingCallsErr != nil {
		return errors.Wrap(_landingCallsErr, "Error serializing 'landingCalls' field")
	}

	if popErr := writeBuffer.PopContext("BACnetAssignedLandingCalls"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetAssignedLandingCalls")
	}
	return nil
}

func (m *BACnetAssignedLandingCalls) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
