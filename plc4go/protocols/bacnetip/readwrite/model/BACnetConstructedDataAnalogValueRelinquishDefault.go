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

// BACnetConstructedDataAnalogValueRelinquishDefault is the data-structure of this message
type BACnetConstructedDataAnalogValueRelinquishDefault struct {
	*BACnetConstructedData
	RelinquishDefault *BACnetApplicationTagReal

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataAnalogValueRelinquishDefault is the corresponding interface of BACnetConstructedDataAnalogValueRelinquishDefault
type IBACnetConstructedDataAnalogValueRelinquishDefault interface {
	IBACnetConstructedData
	// GetRelinquishDefault returns RelinquishDefault (property field)
	GetRelinquishDefault() *BACnetApplicationTagReal
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

func (m *BACnetConstructedDataAnalogValueRelinquishDefault) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_ANALOG_VALUE
}

func (m *BACnetConstructedDataAnalogValueRelinquishDefault) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_RELINQUISH_DEFAULT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataAnalogValueRelinquishDefault) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataAnalogValueRelinquishDefault) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataAnalogValueRelinquishDefault) GetRelinquishDefault() *BACnetApplicationTagReal {
	return m.RelinquishDefault
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataAnalogValueRelinquishDefault factory function for BACnetConstructedDataAnalogValueRelinquishDefault
func NewBACnetConstructedDataAnalogValueRelinquishDefault(relinquishDefault *BACnetApplicationTagReal, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataAnalogValueRelinquishDefault {
	_result := &BACnetConstructedDataAnalogValueRelinquishDefault{
		RelinquishDefault:     relinquishDefault,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataAnalogValueRelinquishDefault(structType interface{}) *BACnetConstructedDataAnalogValueRelinquishDefault {
	if casted, ok := structType.(BACnetConstructedDataAnalogValueRelinquishDefault); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAnalogValueRelinquishDefault); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataAnalogValueRelinquishDefault(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataAnalogValueRelinquishDefault(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataAnalogValueRelinquishDefault) GetTypeName() string {
	return "BACnetConstructedDataAnalogValueRelinquishDefault"
}

func (m *BACnetConstructedDataAnalogValueRelinquishDefault) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataAnalogValueRelinquishDefault) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (relinquishDefault)
	lengthInBits += m.RelinquishDefault.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataAnalogValueRelinquishDefault) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataAnalogValueRelinquishDefaultParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataAnalogValueRelinquishDefault, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAnalogValueRelinquishDefault"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAnalogValueRelinquishDefault")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (relinquishDefault)
	if pullErr := readBuffer.PullContext("relinquishDefault"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for relinquishDefault")
	}
	_relinquishDefault, _relinquishDefaultErr := BACnetApplicationTagParse(readBuffer)
	if _relinquishDefaultErr != nil {
		return nil, errors.Wrap(_relinquishDefaultErr, "Error parsing 'relinquishDefault' field")
	}
	relinquishDefault := CastBACnetApplicationTagReal(_relinquishDefault)
	if closeErr := readBuffer.CloseContext("relinquishDefault"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for relinquishDefault")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAnalogValueRelinquishDefault"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAnalogValueRelinquishDefault")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataAnalogValueRelinquishDefault{
		RelinquishDefault:     CastBACnetApplicationTagReal(relinquishDefault),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataAnalogValueRelinquishDefault) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAnalogValueRelinquishDefault"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAnalogValueRelinquishDefault")
		}

		// Simple Field (relinquishDefault)
		if pushErr := writeBuffer.PushContext("relinquishDefault"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for relinquishDefault")
		}
		_relinquishDefaultErr := m.RelinquishDefault.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("relinquishDefault"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for relinquishDefault")
		}
		if _relinquishDefaultErr != nil {
			return errors.Wrap(_relinquishDefaultErr, "Error serializing 'relinquishDefault' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAnalogValueRelinquishDefault"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAnalogValueRelinquishDefault")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataAnalogValueRelinquishDefault) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
