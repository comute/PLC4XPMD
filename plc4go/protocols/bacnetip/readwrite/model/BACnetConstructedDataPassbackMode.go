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

// BACnetConstructedDataPassbackMode is the data-structure of this message
type BACnetConstructedDataPassbackMode struct {
	*BACnetConstructedData
	PassbackMode *BACnetAccessPassbackModeTagged

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataPassbackMode is the corresponding interface of BACnetConstructedDataPassbackMode
type IBACnetConstructedDataPassbackMode interface {
	IBACnetConstructedData
	// GetPassbackMode returns PassbackMode (property field)
	GetPassbackMode() *BACnetAccessPassbackModeTagged
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

func (m *BACnetConstructedDataPassbackMode) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataPassbackMode) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_PASSBACK_MODE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataPassbackMode) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataPassbackMode) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataPassbackMode) GetPassbackMode() *BACnetAccessPassbackModeTagged {
	return m.PassbackMode
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataPassbackMode factory function for BACnetConstructedDataPassbackMode
func NewBACnetConstructedDataPassbackMode(passbackMode *BACnetAccessPassbackModeTagged, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataPassbackMode {
	_result := &BACnetConstructedDataPassbackMode{
		PassbackMode:          passbackMode,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataPassbackMode(structType interface{}) *BACnetConstructedDataPassbackMode {
	if casted, ok := structType.(BACnetConstructedDataPassbackMode); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataPassbackMode); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataPassbackMode(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataPassbackMode(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataPassbackMode) GetTypeName() string {
	return "BACnetConstructedDataPassbackMode"
}

func (m *BACnetConstructedDataPassbackMode) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataPassbackMode) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (passbackMode)
	lengthInBits += m.PassbackMode.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataPassbackMode) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataPassbackModeParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataPassbackMode, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataPassbackMode"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataPassbackMode")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (passbackMode)
	if pullErr := readBuffer.PullContext("passbackMode"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for passbackMode")
	}
	_passbackMode, _passbackModeErr := BACnetAccessPassbackModeTaggedParse(readBuffer, uint8(uint8(0)), TagClass(TagClass_APPLICATION_TAGS))
	if _passbackModeErr != nil {
		return nil, errors.Wrap(_passbackModeErr, "Error parsing 'passbackMode' field")
	}
	passbackMode := CastBACnetAccessPassbackModeTagged(_passbackMode)
	if closeErr := readBuffer.CloseContext("passbackMode"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for passbackMode")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataPassbackMode"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataPassbackMode")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataPassbackMode{
		PassbackMode:          CastBACnetAccessPassbackModeTagged(passbackMode),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataPassbackMode) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataPassbackMode"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataPassbackMode")
		}

		// Simple Field (passbackMode)
		if pushErr := writeBuffer.PushContext("passbackMode"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for passbackMode")
		}
		_passbackModeErr := writeBuffer.WriteSerializable(m.PassbackMode)
		if popErr := writeBuffer.PopContext("passbackMode"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for passbackMode")
		}
		if _passbackModeErr != nil {
			return errors.Wrap(_passbackModeErr, "Error serializing 'passbackMode' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataPassbackMode"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataPassbackMode")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataPassbackMode) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
