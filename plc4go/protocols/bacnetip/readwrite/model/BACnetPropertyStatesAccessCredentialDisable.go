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

// BACnetPropertyStatesAccessCredentialDisable is the data-structure of this message
type BACnetPropertyStatesAccessCredentialDisable struct {
	*BACnetPropertyStates
	AccessCredentialDisable *BACnetAccessCredentialDisableTagged
}

// IBACnetPropertyStatesAccessCredentialDisable is the corresponding interface of BACnetPropertyStatesAccessCredentialDisable
type IBACnetPropertyStatesAccessCredentialDisable interface {
	IBACnetPropertyStates
	// GetAccessCredentialDisable returns AccessCredentialDisable (property field)
	GetAccessCredentialDisable() *BACnetAccessCredentialDisableTagged
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

func (m *BACnetPropertyStatesAccessCredentialDisable) InitializeParent(parent *BACnetPropertyStates, peekedTagHeader *BACnetTagHeader) {
	m.BACnetPropertyStates.PeekedTagHeader = peekedTagHeader
}

func (m *BACnetPropertyStatesAccessCredentialDisable) GetParent() *BACnetPropertyStates {
	return m.BACnetPropertyStates
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetPropertyStatesAccessCredentialDisable) GetAccessCredentialDisable() *BACnetAccessCredentialDisableTagged {
	return m.AccessCredentialDisable
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStatesAccessCredentialDisable factory function for BACnetPropertyStatesAccessCredentialDisable
func NewBACnetPropertyStatesAccessCredentialDisable(accessCredentialDisable *BACnetAccessCredentialDisableTagged, peekedTagHeader *BACnetTagHeader) *BACnetPropertyStatesAccessCredentialDisable {
	_result := &BACnetPropertyStatesAccessCredentialDisable{
		AccessCredentialDisable: accessCredentialDisable,
		BACnetPropertyStates:    NewBACnetPropertyStates(peekedTagHeader),
	}
	_result.Child = _result
	return _result
}

func CastBACnetPropertyStatesAccessCredentialDisable(structType interface{}) *BACnetPropertyStatesAccessCredentialDisable {
	if casted, ok := structType.(BACnetPropertyStatesAccessCredentialDisable); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesAccessCredentialDisable); ok {
		return casted
	}
	if casted, ok := structType.(BACnetPropertyStates); ok {
		return CastBACnetPropertyStatesAccessCredentialDisable(casted.Child)
	}
	if casted, ok := structType.(*BACnetPropertyStates); ok {
		return CastBACnetPropertyStatesAccessCredentialDisable(casted.Child)
	}
	return nil
}

func (m *BACnetPropertyStatesAccessCredentialDisable) GetTypeName() string {
	return "BACnetPropertyStatesAccessCredentialDisable"
}

func (m *BACnetPropertyStatesAccessCredentialDisable) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetPropertyStatesAccessCredentialDisable) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (accessCredentialDisable)
	lengthInBits += m.AccessCredentialDisable.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetPropertyStatesAccessCredentialDisable) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetPropertyStatesAccessCredentialDisableParse(readBuffer utils.ReadBuffer, peekedTagNumber uint8) (*BACnetPropertyStatesAccessCredentialDisable, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesAccessCredentialDisable"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesAccessCredentialDisable")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (accessCredentialDisable)
	if pullErr := readBuffer.PullContext("accessCredentialDisable"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for accessCredentialDisable")
	}
	_accessCredentialDisable, _accessCredentialDisableErr := BACnetAccessCredentialDisableTaggedParse(readBuffer, uint8(peekedTagNumber), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _accessCredentialDisableErr != nil {
		return nil, errors.Wrap(_accessCredentialDisableErr, "Error parsing 'accessCredentialDisable' field")
	}
	accessCredentialDisable := CastBACnetAccessCredentialDisableTagged(_accessCredentialDisable)
	if closeErr := readBuffer.CloseContext("accessCredentialDisable"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for accessCredentialDisable")
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesAccessCredentialDisable"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesAccessCredentialDisable")
	}

	// Create a partially initialized instance
	_child := &BACnetPropertyStatesAccessCredentialDisable{
		AccessCredentialDisable: CastBACnetAccessCredentialDisableTagged(accessCredentialDisable),
		BACnetPropertyStates:    &BACnetPropertyStates{},
	}
	_child.BACnetPropertyStates.Child = _child
	return _child, nil
}

func (m *BACnetPropertyStatesAccessCredentialDisable) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesAccessCredentialDisable"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesAccessCredentialDisable")
		}

		// Simple Field (accessCredentialDisable)
		if pushErr := writeBuffer.PushContext("accessCredentialDisable"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for accessCredentialDisable")
		}
		_accessCredentialDisableErr := writeBuffer.WriteSerializable(m.AccessCredentialDisable)
		if popErr := writeBuffer.PopContext("accessCredentialDisable"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for accessCredentialDisable")
		}
		if _accessCredentialDisableErr != nil {
			return errors.Wrap(_accessCredentialDisableErr, "Error serializing 'accessCredentialDisable' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesAccessCredentialDisable"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesAccessCredentialDisable")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetPropertyStatesAccessCredentialDisable) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
