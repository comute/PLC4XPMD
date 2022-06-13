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

// BACnetPropertyStatesWriteStatus is the data-structure of this message
type BACnetPropertyStatesWriteStatus struct {
	*BACnetPropertyStates
	WriteStatus *BACnetWriteStatusTagged
}

// IBACnetPropertyStatesWriteStatus is the corresponding interface of BACnetPropertyStatesWriteStatus
type IBACnetPropertyStatesWriteStatus interface {
	IBACnetPropertyStates
	// GetWriteStatus returns WriteStatus (property field)
	GetWriteStatus() *BACnetWriteStatusTagged
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

func (m *BACnetPropertyStatesWriteStatus) InitializeParent(parent *BACnetPropertyStates, peekedTagHeader *BACnetTagHeader) {
	m.BACnetPropertyStates.PeekedTagHeader = peekedTagHeader
}

func (m *BACnetPropertyStatesWriteStatus) GetParent() *BACnetPropertyStates {
	return m.BACnetPropertyStates
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetPropertyStatesWriteStatus) GetWriteStatus() *BACnetWriteStatusTagged {
	return m.WriteStatus
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStatesWriteStatus factory function for BACnetPropertyStatesWriteStatus
func NewBACnetPropertyStatesWriteStatus(writeStatus *BACnetWriteStatusTagged, peekedTagHeader *BACnetTagHeader) *BACnetPropertyStatesWriteStatus {
	_result := &BACnetPropertyStatesWriteStatus{
		WriteStatus:          writeStatus,
		BACnetPropertyStates: NewBACnetPropertyStates(peekedTagHeader),
	}
	_result.Child = _result
	return _result
}

func CastBACnetPropertyStatesWriteStatus(structType interface{}) *BACnetPropertyStatesWriteStatus {
	if casted, ok := structType.(BACnetPropertyStatesWriteStatus); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesWriteStatus); ok {
		return casted
	}
	if casted, ok := structType.(BACnetPropertyStates); ok {
		return CastBACnetPropertyStatesWriteStatus(casted.Child)
	}
	if casted, ok := structType.(*BACnetPropertyStates); ok {
		return CastBACnetPropertyStatesWriteStatus(casted.Child)
	}
	return nil
}

func (m *BACnetPropertyStatesWriteStatus) GetTypeName() string {
	return "BACnetPropertyStatesWriteStatus"
}

func (m *BACnetPropertyStatesWriteStatus) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetPropertyStatesWriteStatus) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (writeStatus)
	lengthInBits += m.WriteStatus.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetPropertyStatesWriteStatus) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetPropertyStatesWriteStatusParse(readBuffer utils.ReadBuffer, peekedTagNumber uint8) (*BACnetPropertyStatesWriteStatus, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesWriteStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesWriteStatus")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (writeStatus)
	if pullErr := readBuffer.PullContext("writeStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for writeStatus")
	}
	_writeStatus, _writeStatusErr := BACnetWriteStatusTaggedParse(readBuffer, uint8(peekedTagNumber), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _writeStatusErr != nil {
		return nil, errors.Wrap(_writeStatusErr, "Error parsing 'writeStatus' field")
	}
	writeStatus := CastBACnetWriteStatusTagged(_writeStatus)
	if closeErr := readBuffer.CloseContext("writeStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for writeStatus")
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesWriteStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesWriteStatus")
	}

	// Create a partially initialized instance
	_child := &BACnetPropertyStatesWriteStatus{
		WriteStatus:          CastBACnetWriteStatusTagged(writeStatus),
		BACnetPropertyStates: &BACnetPropertyStates{},
	}
	_child.BACnetPropertyStates.Child = _child
	return _child, nil
}

func (m *BACnetPropertyStatesWriteStatus) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesWriteStatus"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesWriteStatus")
		}

		// Simple Field (writeStatus)
		if pushErr := writeBuffer.PushContext("writeStatus"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for writeStatus")
		}
		_writeStatusErr := writeBuffer.WriteSerializable(m.WriteStatus)
		if popErr := writeBuffer.PopContext("writeStatus"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for writeStatus")
		}
		if _writeStatusErr != nil {
			return errors.Wrap(_writeStatusErr, "Error serializing 'writeStatus' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesWriteStatus"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesWriteStatus")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetPropertyStatesWriteStatus) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
