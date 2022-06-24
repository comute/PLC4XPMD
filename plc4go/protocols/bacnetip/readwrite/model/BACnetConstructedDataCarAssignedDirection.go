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

// BACnetConstructedDataCarAssignedDirection is the corresponding interface of BACnetConstructedDataCarAssignedDirection
type BACnetConstructedDataCarAssignedDirection interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetAssignedDirection returns AssignedDirection (property field)
	GetAssignedDirection() BACnetLiftCarDirectionTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetLiftCarDirectionTagged
}

// BACnetConstructedDataCarAssignedDirectionExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataCarAssignedDirection.
// This is useful for switch cases.
type BACnetConstructedDataCarAssignedDirectionExactly interface {
	BACnetConstructedDataCarAssignedDirection
	isBACnetConstructedDataCarAssignedDirection() bool
}

// _BACnetConstructedDataCarAssignedDirection is the data-structure of this message
type _BACnetConstructedDataCarAssignedDirection struct {
	*_BACnetConstructedData
	AssignedDirection BACnetLiftCarDirectionTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataCarAssignedDirection) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataCarAssignedDirection) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_CAR_ASSIGNED_DIRECTION
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataCarAssignedDirection) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataCarAssignedDirection) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataCarAssignedDirection) GetAssignedDirection() BACnetLiftCarDirectionTagged {
	return m.AssignedDirection
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataCarAssignedDirection) GetActualValue() BACnetLiftCarDirectionTagged {
	return CastBACnetLiftCarDirectionTagged(m.GetAssignedDirection())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataCarAssignedDirection factory function for _BACnetConstructedDataCarAssignedDirection
func NewBACnetConstructedDataCarAssignedDirection(assignedDirection BACnetLiftCarDirectionTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataCarAssignedDirection {
	_result := &_BACnetConstructedDataCarAssignedDirection{
		AssignedDirection:      assignedDirection,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataCarAssignedDirection(structType interface{}) BACnetConstructedDataCarAssignedDirection {
	if casted, ok := structType.(BACnetConstructedDataCarAssignedDirection); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataCarAssignedDirection); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataCarAssignedDirection) GetTypeName() string {
	return "BACnetConstructedDataCarAssignedDirection"
}

func (m *_BACnetConstructedDataCarAssignedDirection) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataCarAssignedDirection) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (assignedDirection)
	lengthInBits += m.AssignedDirection.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataCarAssignedDirection) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataCarAssignedDirectionParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataCarAssignedDirection, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataCarAssignedDirection"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataCarAssignedDirection")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (assignedDirection)
	if pullErr := readBuffer.PullContext("assignedDirection"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for assignedDirection")
	}
	_assignedDirection, _assignedDirectionErr := BACnetLiftCarDirectionTaggedParse(readBuffer, uint8(uint8(0)), TagClass(TagClass_APPLICATION_TAGS))
	if _assignedDirectionErr != nil {
		return nil, errors.Wrap(_assignedDirectionErr, "Error parsing 'assignedDirection' field")
	}
	assignedDirection := _assignedDirection.(BACnetLiftCarDirectionTagged)
	if closeErr := readBuffer.CloseContext("assignedDirection"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for assignedDirection")
	}

	// Virtual field
	_actualValue := assignedDirection
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataCarAssignedDirection"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataCarAssignedDirection")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataCarAssignedDirection{
		AssignedDirection: assignedDirection,
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataCarAssignedDirection) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataCarAssignedDirection"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataCarAssignedDirection")
		}

		// Simple Field (assignedDirection)
		if pushErr := writeBuffer.PushContext("assignedDirection"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for assignedDirection")
		}
		_assignedDirectionErr := writeBuffer.WriteSerializable(m.GetAssignedDirection())
		if popErr := writeBuffer.PopContext("assignedDirection"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for assignedDirection")
		}
		if _assignedDirectionErr != nil {
			return errors.Wrap(_assignedDirectionErr, "Error serializing 'assignedDirection' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataCarAssignedDirection"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataCarAssignedDirection")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataCarAssignedDirection) isBACnetConstructedDataCarAssignedDirection() bool {
	return true
}

func (m *_BACnetConstructedDataCarAssignedDirection) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
