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

// BACnetPropertyStatesLifeSafetyOperations is the data-structure of this message
type BACnetPropertyStatesLifeSafetyOperations struct {
	*BACnetPropertyStates
	LifeSafetyOperations *BACnetLifeSafetyOperationTagged
}

// IBACnetPropertyStatesLifeSafetyOperations is the corresponding interface of BACnetPropertyStatesLifeSafetyOperations
type IBACnetPropertyStatesLifeSafetyOperations interface {
	IBACnetPropertyStates
	// GetLifeSafetyOperations returns LifeSafetyOperations (property field)
	GetLifeSafetyOperations() *BACnetLifeSafetyOperationTagged
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

func (m *BACnetPropertyStatesLifeSafetyOperations) InitializeParent(parent *BACnetPropertyStates, peekedTagHeader *BACnetTagHeader) {
	m.BACnetPropertyStates.PeekedTagHeader = peekedTagHeader
}

func (m *BACnetPropertyStatesLifeSafetyOperations) GetParent() *BACnetPropertyStates {
	return m.BACnetPropertyStates
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetPropertyStatesLifeSafetyOperations) GetLifeSafetyOperations() *BACnetLifeSafetyOperationTagged {
	return m.LifeSafetyOperations
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStatesLifeSafetyOperations factory function for BACnetPropertyStatesLifeSafetyOperations
func NewBACnetPropertyStatesLifeSafetyOperations(lifeSafetyOperations *BACnetLifeSafetyOperationTagged, peekedTagHeader *BACnetTagHeader) *BACnetPropertyStatesLifeSafetyOperations {
	_result := &BACnetPropertyStatesLifeSafetyOperations{
		LifeSafetyOperations: lifeSafetyOperations,
		BACnetPropertyStates: NewBACnetPropertyStates(peekedTagHeader),
	}
	_result.Child = _result
	return _result
}

func CastBACnetPropertyStatesLifeSafetyOperations(structType interface{}) *BACnetPropertyStatesLifeSafetyOperations {
	if casted, ok := structType.(BACnetPropertyStatesLifeSafetyOperations); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesLifeSafetyOperations); ok {
		return casted
	}
	if casted, ok := structType.(BACnetPropertyStates); ok {
		return CastBACnetPropertyStatesLifeSafetyOperations(casted.Child)
	}
	if casted, ok := structType.(*BACnetPropertyStates); ok {
		return CastBACnetPropertyStatesLifeSafetyOperations(casted.Child)
	}
	return nil
}

func (m *BACnetPropertyStatesLifeSafetyOperations) GetTypeName() string {
	return "BACnetPropertyStatesLifeSafetyOperations"
}

func (m *BACnetPropertyStatesLifeSafetyOperations) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetPropertyStatesLifeSafetyOperations) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (lifeSafetyOperations)
	lengthInBits += m.LifeSafetyOperations.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetPropertyStatesLifeSafetyOperations) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetPropertyStatesLifeSafetyOperationsParse(readBuffer utils.ReadBuffer, peekedTagNumber uint8) (*BACnetPropertyStatesLifeSafetyOperations, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesLifeSafetyOperations"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesLifeSafetyOperations")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (lifeSafetyOperations)
	if pullErr := readBuffer.PullContext("lifeSafetyOperations"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for lifeSafetyOperations")
	}
	_lifeSafetyOperations, _lifeSafetyOperationsErr := BACnetLifeSafetyOperationTaggedParse(readBuffer, uint8(peekedTagNumber), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _lifeSafetyOperationsErr != nil {
		return nil, errors.Wrap(_lifeSafetyOperationsErr, "Error parsing 'lifeSafetyOperations' field")
	}
	lifeSafetyOperations := CastBACnetLifeSafetyOperationTagged(_lifeSafetyOperations)
	if closeErr := readBuffer.CloseContext("lifeSafetyOperations"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for lifeSafetyOperations")
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesLifeSafetyOperations"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesLifeSafetyOperations")
	}

	// Create a partially initialized instance
	_child := &BACnetPropertyStatesLifeSafetyOperations{
		LifeSafetyOperations: CastBACnetLifeSafetyOperationTagged(lifeSafetyOperations),
		BACnetPropertyStates: &BACnetPropertyStates{},
	}
	_child.BACnetPropertyStates.Child = _child
	return _child, nil
}

func (m *BACnetPropertyStatesLifeSafetyOperations) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesLifeSafetyOperations"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesLifeSafetyOperations")
		}

		// Simple Field (lifeSafetyOperations)
		if pushErr := writeBuffer.PushContext("lifeSafetyOperations"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for lifeSafetyOperations")
		}
		_lifeSafetyOperationsErr := writeBuffer.WriteSerializable(m.LifeSafetyOperations)
		if popErr := writeBuffer.PopContext("lifeSafetyOperations"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for lifeSafetyOperations")
		}
		if _lifeSafetyOperationsErr != nil {
			return errors.Wrap(_lifeSafetyOperationsErr, "Error serializing 'lifeSafetyOperations' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesLifeSafetyOperations"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesLifeSafetyOperations")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetPropertyStatesLifeSafetyOperations) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
