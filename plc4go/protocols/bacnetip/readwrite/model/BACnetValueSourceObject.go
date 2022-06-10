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

// BACnetValueSourceObject is the data-structure of this message
type BACnetValueSourceObject struct {
	*BACnetValueSource
	Object *BACnetDeviceObjectReferenceEnclosed
}

// IBACnetValueSourceObject is the corresponding interface of BACnetValueSourceObject
type IBACnetValueSourceObject interface {
	IBACnetValueSource
	// GetObject returns Object (property field)
	GetObject() *BACnetDeviceObjectReferenceEnclosed
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

func (m *BACnetValueSourceObject) InitializeParent(parent *BACnetValueSource, peekedTagHeader *BACnetTagHeader) {
	m.BACnetValueSource.PeekedTagHeader = peekedTagHeader
}

func (m *BACnetValueSourceObject) GetParent() *BACnetValueSource {
	return m.BACnetValueSource
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetValueSourceObject) GetObject() *BACnetDeviceObjectReferenceEnclosed {
	return m.Object
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetValueSourceObject factory function for BACnetValueSourceObject
func NewBACnetValueSourceObject(object *BACnetDeviceObjectReferenceEnclosed, peekedTagHeader *BACnetTagHeader) *BACnetValueSourceObject {
	_result := &BACnetValueSourceObject{
		Object:            object,
		BACnetValueSource: NewBACnetValueSource(peekedTagHeader),
	}
	_result.Child = _result
	return _result
}

func CastBACnetValueSourceObject(structType interface{}) *BACnetValueSourceObject {
	if casted, ok := structType.(BACnetValueSourceObject); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetValueSourceObject); ok {
		return casted
	}
	if casted, ok := structType.(BACnetValueSource); ok {
		return CastBACnetValueSourceObject(casted.Child)
	}
	if casted, ok := structType.(*BACnetValueSource); ok {
		return CastBACnetValueSourceObject(casted.Child)
	}
	return nil
}

func (m *BACnetValueSourceObject) GetTypeName() string {
	return "BACnetValueSourceObject"
}

func (m *BACnetValueSourceObject) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetValueSourceObject) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (object)
	lengthInBits += m.Object.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetValueSourceObject) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetValueSourceObjectParse(readBuffer utils.ReadBuffer) (*BACnetValueSourceObject, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetValueSourceObject"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetValueSourceObject")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (object)
	if pullErr := readBuffer.PullContext("object"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for object")
	}
	_object, _objectErr := BACnetDeviceObjectReferenceEnclosedParse(readBuffer, uint8(uint8(1)))
	if _objectErr != nil {
		return nil, errors.Wrap(_objectErr, "Error parsing 'object' field")
	}
	object := CastBACnetDeviceObjectReferenceEnclosed(_object)
	if closeErr := readBuffer.CloseContext("object"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for object")
	}

	if closeErr := readBuffer.CloseContext("BACnetValueSourceObject"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetValueSourceObject")
	}

	// Create a partially initialized instance
	_child := &BACnetValueSourceObject{
		Object:            CastBACnetDeviceObjectReferenceEnclosed(object),
		BACnetValueSource: &BACnetValueSource{},
	}
	_child.BACnetValueSource.Child = _child
	return _child, nil
}

func (m *BACnetValueSourceObject) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetValueSourceObject"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetValueSourceObject")
		}

		// Simple Field (object)
		if pushErr := writeBuffer.PushContext("object"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for object")
		}
		_objectErr := m.Object.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("object"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for object")
		}
		if _objectErr != nil {
			return errors.Wrap(_objectErr, "Error serializing 'object' field")
		}

		if popErr := writeBuffer.PopContext("BACnetValueSourceObject"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetValueSourceObject")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetValueSourceObject) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
