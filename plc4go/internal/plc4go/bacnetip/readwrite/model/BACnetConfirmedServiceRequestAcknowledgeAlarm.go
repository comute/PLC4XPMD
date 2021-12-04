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
type BACnetConfirmedServiceRequestAcknowledgeAlarm struct {
	*BACnetConfirmedServiceRequest
}

// The corresponding interface
type IBACnetConfirmedServiceRequestAcknowledgeAlarm interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BACnetConfirmedServiceRequestAcknowledgeAlarm) ServiceChoice() uint8 {
	return 0x00
}

func (m *BACnetConfirmedServiceRequestAcknowledgeAlarm) InitializeParent(parent *BACnetConfirmedServiceRequest) {
}

func NewBACnetConfirmedServiceRequestAcknowledgeAlarm() *BACnetConfirmedServiceRequest {
	child := &BACnetConfirmedServiceRequestAcknowledgeAlarm{
		BACnetConfirmedServiceRequest: NewBACnetConfirmedServiceRequest(),
	}
	child.Child = child
	return child.BACnetConfirmedServiceRequest
}

func CastBACnetConfirmedServiceRequestAcknowledgeAlarm(structType interface{}) *BACnetConfirmedServiceRequestAcknowledgeAlarm {
	castFunc := func(typ interface{}) *BACnetConfirmedServiceRequestAcknowledgeAlarm {
		if casted, ok := typ.(BACnetConfirmedServiceRequestAcknowledgeAlarm); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetConfirmedServiceRequestAcknowledgeAlarm); ok {
			return casted
		}
		if casted, ok := typ.(BACnetConfirmedServiceRequest); ok {
			return CastBACnetConfirmedServiceRequestAcknowledgeAlarm(casted.Child)
		}
		if casted, ok := typ.(*BACnetConfirmedServiceRequest); ok {
			return CastBACnetConfirmedServiceRequestAcknowledgeAlarm(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetConfirmedServiceRequestAcknowledgeAlarm) GetTypeName() string {
	return "BACnetConfirmedServiceRequestAcknowledgeAlarm"
}

func (m *BACnetConfirmedServiceRequestAcknowledgeAlarm) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetConfirmedServiceRequestAcknowledgeAlarm) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.ParentLengthInBits())

	return lengthInBits
}

func (m *BACnetConfirmedServiceRequestAcknowledgeAlarm) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetConfirmedServiceRequestAcknowledgeAlarmParse(readBuffer utils.ReadBuffer, len uint16) (*BACnetConfirmedServiceRequest, error) {
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestAcknowledgeAlarm"); pullErr != nil {
		return nil, pullErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestAcknowledgeAlarm"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConfirmedServiceRequestAcknowledgeAlarm{
		BACnetConfirmedServiceRequest: &BACnetConfirmedServiceRequest{},
	}
	_child.BACnetConfirmedServiceRequest.Child = _child
	return _child.BACnetConfirmedServiceRequest, nil
}

func (m *BACnetConfirmedServiceRequestAcknowledgeAlarm) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestAcknowledgeAlarm"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestAcknowledgeAlarm"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConfirmedServiceRequestAcknowledgeAlarm) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
