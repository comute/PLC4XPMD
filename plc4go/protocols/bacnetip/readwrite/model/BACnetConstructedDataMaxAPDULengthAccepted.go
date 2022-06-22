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

// BACnetConstructedDataMaxAPDULengthAccepted is the corresponding interface of BACnetConstructedDataMaxAPDULengthAccepted
type BACnetConstructedDataMaxAPDULengthAccepted interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetMaxApduLengthAccepted returns MaxApduLengthAccepted (property field)
	GetMaxApduLengthAccepted() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataMaxAPDULengthAcceptedExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataMaxAPDULengthAccepted.
// This is useful for switch cases.
type BACnetConstructedDataMaxAPDULengthAcceptedExactly interface {
	BACnetConstructedDataMaxAPDULengthAccepted
	isBACnetConstructedDataMaxAPDULengthAccepted() bool
}

// _BACnetConstructedDataMaxAPDULengthAccepted is the data-structure of this message
type _BACnetConstructedDataMaxAPDULengthAccepted struct {
	*_BACnetConstructedData
	MaxApduLengthAccepted BACnetApplicationTagUnsignedInteger

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument BACnetTagPayloadUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataMaxAPDULengthAccepted) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataMaxAPDULengthAccepted) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_MAX_APDU_LENGTH_ACCEPTED
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataMaxAPDULengthAccepted) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataMaxAPDULengthAccepted) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataMaxAPDULengthAccepted) GetMaxApduLengthAccepted() BACnetApplicationTagUnsignedInteger {
	return m.MaxApduLengthAccepted
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataMaxAPDULengthAccepted) GetActualValue() BACnetApplicationTagUnsignedInteger {
	return CastBACnetApplicationTagUnsignedInteger(m.GetMaxApduLengthAccepted())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataMaxAPDULengthAccepted factory function for _BACnetConstructedDataMaxAPDULengthAccepted
func NewBACnetConstructedDataMaxAPDULengthAccepted(maxApduLengthAccepted BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataMaxAPDULengthAccepted {
	_result := &_BACnetConstructedDataMaxAPDULengthAccepted{
		MaxApduLengthAccepted:  maxApduLengthAccepted,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataMaxAPDULengthAccepted(structType interface{}) BACnetConstructedDataMaxAPDULengthAccepted {
	if casted, ok := structType.(BACnetConstructedDataMaxAPDULengthAccepted); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataMaxAPDULengthAccepted); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataMaxAPDULengthAccepted) GetTypeName() string {
	return "BACnetConstructedDataMaxAPDULengthAccepted"
}

func (m *_BACnetConstructedDataMaxAPDULengthAccepted) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataMaxAPDULengthAccepted) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (maxApduLengthAccepted)
	lengthInBits += m.MaxApduLengthAccepted.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataMaxAPDULengthAccepted) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataMaxAPDULengthAcceptedParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataMaxAPDULengthAccepted, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataMaxAPDULengthAccepted"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataMaxAPDULengthAccepted")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (maxApduLengthAccepted)
	if pullErr := readBuffer.PullContext("maxApduLengthAccepted"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for maxApduLengthAccepted")
	}
	_maxApduLengthAccepted, _maxApduLengthAcceptedErr := BACnetApplicationTagParse(readBuffer)
	if _maxApduLengthAcceptedErr != nil {
		return nil, errors.Wrap(_maxApduLengthAcceptedErr, "Error parsing 'maxApduLengthAccepted' field")
	}
	maxApduLengthAccepted := _maxApduLengthAccepted.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("maxApduLengthAccepted"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for maxApduLengthAccepted")
	}

	// Virtual field
	_actualValue := maxApduLengthAccepted
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataMaxAPDULengthAccepted"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataMaxAPDULengthAccepted")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataMaxAPDULengthAccepted{
		MaxApduLengthAccepted:  maxApduLengthAccepted,
		_BACnetConstructedData: &_BACnetConstructedData{},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataMaxAPDULengthAccepted) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataMaxAPDULengthAccepted"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataMaxAPDULengthAccepted")
		}

		// Simple Field (maxApduLengthAccepted)
		if pushErr := writeBuffer.PushContext("maxApduLengthAccepted"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for maxApduLengthAccepted")
		}
		_maxApduLengthAcceptedErr := writeBuffer.WriteSerializable(m.GetMaxApduLengthAccepted())
		if popErr := writeBuffer.PopContext("maxApduLengthAccepted"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for maxApduLengthAccepted")
		}
		if _maxApduLengthAcceptedErr != nil {
			return errors.Wrap(_maxApduLengthAcceptedErr, "Error serializing 'maxApduLengthAccepted' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataMaxAPDULengthAccepted"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataMaxAPDULengthAccepted")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataMaxAPDULengthAccepted) isBACnetConstructedDataMaxAPDULengthAccepted() bool {
	return true
}

func (m *_BACnetConstructedDataMaxAPDULengthAccepted) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
