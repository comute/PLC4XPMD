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

// BACnetPropertyReferenceEnclosed is the data-structure of this message
type BACnetPropertyReferenceEnclosed struct {
	OpeningTag *BACnetOpeningTag
	Reference  *BACnetPropertyReference
	ClosingTag *BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

// IBACnetPropertyReferenceEnclosed is the corresponding interface of BACnetPropertyReferenceEnclosed
type IBACnetPropertyReferenceEnclosed interface {
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() *BACnetOpeningTag
	// GetReference returns Reference (property field)
	GetReference() *BACnetPropertyReference
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() *BACnetClosingTag
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetPropertyReferenceEnclosed) GetOpeningTag() *BACnetOpeningTag {
	return m.OpeningTag
}

func (m *BACnetPropertyReferenceEnclosed) GetReference() *BACnetPropertyReference {
	return m.Reference
}

func (m *BACnetPropertyReferenceEnclosed) GetClosingTag() *BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyReferenceEnclosed factory function for BACnetPropertyReferenceEnclosed
func NewBACnetPropertyReferenceEnclosed(openingTag *BACnetOpeningTag, reference *BACnetPropertyReference, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetPropertyReferenceEnclosed {
	return &BACnetPropertyReferenceEnclosed{OpeningTag: openingTag, Reference: reference, ClosingTag: closingTag, TagNumber: tagNumber}
}

func CastBACnetPropertyReferenceEnclosed(structType interface{}) *BACnetPropertyReferenceEnclosed {
	if casted, ok := structType.(BACnetPropertyReferenceEnclosed); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetPropertyReferenceEnclosed); ok {
		return casted
	}
	return nil
}

func (m *BACnetPropertyReferenceEnclosed) GetTypeName() string {
	return "BACnetPropertyReferenceEnclosed"
}

func (m *BACnetPropertyReferenceEnclosed) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetPropertyReferenceEnclosed) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits()

	// Simple field (reference)
	lengthInBits += m.Reference.GetLengthInBits()

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetPropertyReferenceEnclosed) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetPropertyReferenceEnclosedParse(readBuffer utils.ReadBuffer, tagNumber uint8) (*BACnetPropertyReferenceEnclosed, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyReferenceEnclosed"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyReferenceEnclosed")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for openingTag")
	}
	_openingTag, _openingTagErr := BACnetOpeningTagParse(readBuffer, uint8(tagNumber))
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field")
	}
	openingTag := CastBACnetOpeningTag(_openingTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for openingTag")
	}

	// Simple Field (reference)
	if pullErr := readBuffer.PullContext("reference"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for reference")
	}
	_reference, _referenceErr := BACnetPropertyReferenceParse(readBuffer)
	if _referenceErr != nil {
		return nil, errors.Wrap(_referenceErr, "Error parsing 'reference' field")
	}
	reference := CastBACnetPropertyReference(_reference)
	if closeErr := readBuffer.CloseContext("reference"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for reference")
	}

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for closingTag")
	}
	_closingTag, _closingTagErr := BACnetClosingTagParse(readBuffer, uint8(tagNumber))
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field")
	}
	closingTag := CastBACnetClosingTag(_closingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for closingTag")
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyReferenceEnclosed"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyReferenceEnclosed")
	}

	// Create the instance
	return NewBACnetPropertyReferenceEnclosed(openingTag, reference, closingTag, tagNumber), nil
}

func (m *BACnetPropertyReferenceEnclosed) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetPropertyReferenceEnclosed"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetPropertyReferenceEnclosed")
	}

	// Simple Field (openingTag)
	if pushErr := writeBuffer.PushContext("openingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for openingTag")
	}
	_openingTagErr := writeBuffer.WriteSerializable(m.OpeningTag)
	if popErr := writeBuffer.PopContext("openingTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for openingTag")
	}
	if _openingTagErr != nil {
		return errors.Wrap(_openingTagErr, "Error serializing 'openingTag' field")
	}

	// Simple Field (reference)
	if pushErr := writeBuffer.PushContext("reference"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for reference")
	}
	_referenceErr := writeBuffer.WriteSerializable(m.Reference)
	if popErr := writeBuffer.PopContext("reference"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for reference")
	}
	if _referenceErr != nil {
		return errors.Wrap(_referenceErr, "Error serializing 'reference' field")
	}

	// Simple Field (closingTag)
	if pushErr := writeBuffer.PushContext("closingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for closingTag")
	}
	_closingTagErr := writeBuffer.WriteSerializable(m.ClosingTag)
	if popErr := writeBuffer.PopContext("closingTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for closingTag")
	}
	if _closingTagErr != nil {
		return errors.Wrap(_closingTagErr, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetPropertyReferenceEnclosed"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetPropertyReferenceEnclosed")
	}
	return nil
}

func (m *BACnetPropertyReferenceEnclosed) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
