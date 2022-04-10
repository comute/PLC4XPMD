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
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type BACnetConfirmedServiceACKGetAlarmSummary struct {
	*BACnetConfirmedServiceACK
}

// The corresponding interface
type IBACnetConfirmedServiceACKGetAlarmSummary interface {
	IBACnetConfirmedServiceACK
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
func (m *BACnetConfirmedServiceACKGetAlarmSummary) GetServiceChoice() uint8 {
	return 0x03
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConfirmedServiceACKGetAlarmSummary) InitializeParent(parent *BACnetConfirmedServiceACK) {
}

func (m *BACnetConfirmedServiceACKGetAlarmSummary) GetParent() *BACnetConfirmedServiceACK {
	return m.BACnetConfirmedServiceACK
}

// NewBACnetConfirmedServiceACKGetAlarmSummary factory function for BACnetConfirmedServiceACKGetAlarmSummary
func NewBACnetConfirmedServiceACKGetAlarmSummary() *BACnetConfirmedServiceACKGetAlarmSummary {
	_result := &BACnetConfirmedServiceACKGetAlarmSummary{
		BACnetConfirmedServiceACK: NewBACnetConfirmedServiceACK(),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConfirmedServiceACKGetAlarmSummary(structType interface{}) *BACnetConfirmedServiceACKGetAlarmSummary {
	if casted, ok := structType.(BACnetConfirmedServiceACKGetAlarmSummary); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceACKGetAlarmSummary); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConfirmedServiceACK); ok {
		return CastBACnetConfirmedServiceACKGetAlarmSummary(casted.Child)
	}
	if casted, ok := structType.(*BACnetConfirmedServiceACK); ok {
		return CastBACnetConfirmedServiceACKGetAlarmSummary(casted.Child)
	}
	return nil
}

func (m *BACnetConfirmedServiceACKGetAlarmSummary) GetTypeName() string {
	return "BACnetConfirmedServiceACKGetAlarmSummary"
}

func (m *BACnetConfirmedServiceACKGetAlarmSummary) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConfirmedServiceACKGetAlarmSummary) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *BACnetConfirmedServiceACKGetAlarmSummary) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConfirmedServiceACKGetAlarmSummaryParse(readBuffer utils.ReadBuffer) (*BACnetConfirmedServiceACKGetAlarmSummary, error) {
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceACKGetAlarmSummary"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceACKGetAlarmSummary"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConfirmedServiceACKGetAlarmSummary{
		BACnetConfirmedServiceACK: &BACnetConfirmedServiceACK{},
	}
	_child.BACnetConfirmedServiceACK.Child = _child
	return _child, nil
}

func (m *BACnetConfirmedServiceACKGetAlarmSummary) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceACKGetAlarmSummary"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceACKGetAlarmSummary"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConfirmedServiceACKGetAlarmSummary) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
