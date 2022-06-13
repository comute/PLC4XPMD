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

// BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferences is the data-structure of this message
type BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferences struct {
	OpeningTag          *BACnetOpeningTag
	ListOfCovReferences []*BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferencesEntry
	ClosingTag          *BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

// IBACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferences is the corresponding interface of BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferences
type IBACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferences interface {
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() *BACnetOpeningTag
	// GetListOfCovReferences returns ListOfCovReferences (property field)
	GetListOfCovReferences() []*BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferencesEntry
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

func (m *BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferences) GetOpeningTag() *BACnetOpeningTag {
	return m.OpeningTag
}

func (m *BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferences) GetListOfCovReferences() []*BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferencesEntry {
	return m.ListOfCovReferences
}

func (m *BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferences) GetClosingTag() *BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferences factory function for BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferences
func NewBACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferences(openingTag *BACnetOpeningTag, listOfCovReferences []*BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferencesEntry, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferences {
	return &BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferences{OpeningTag: openingTag, ListOfCovReferences: listOfCovReferences, ClosingTag: closingTag, TagNumber: tagNumber}
}

func CastBACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferences(structType interface{}) *BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferences {
	if casted, ok := structType.(BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferences); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferences); ok {
		return casted
	}
	return nil
}

func (m *BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferences) GetTypeName() string {
	return "BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferences"
}

func (m *BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferences) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferences) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits()

	// Array field
	if len(m.ListOfCovReferences) > 0 {
		for _, element := range m.ListOfCovReferences {
			lengthInBits += element.GetLengthInBits()
		}
	}

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferences) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferencesParse(readBuffer utils.ReadBuffer, tagNumber uint8) (*BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferences, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferences"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferences")
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

	// Array field (listOfCovReferences)
	if pullErr := readBuffer.PullContext("listOfCovReferences", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for listOfCovReferences")
	}
	// Terminated array
	listOfCovReferences := make([]*BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferencesEntry, 0)
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferencesEntryParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'listOfCovReferences' field")
			}
			listOfCovReferences = append(listOfCovReferences, CastBACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferencesEntry(_item))

		}
	}
	if closeErr := readBuffer.CloseContext("listOfCovReferences", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for listOfCovReferences")
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

	if closeErr := readBuffer.CloseContext("BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferences"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferences")
	}

	// Create the instance
	return NewBACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferences(openingTag, listOfCovReferences, closingTag, tagNumber), nil
}

func (m *BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferences) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferences"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferences")
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

	// Array Field (listOfCovReferences)
	if m.ListOfCovReferences != nil {
		if pushErr := writeBuffer.PushContext("listOfCovReferences", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for listOfCovReferences")
		}
		for _, _element := range m.ListOfCovReferences {
			_elementErr := writeBuffer.WriteSerializable(_element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'listOfCovReferences' field")
			}
		}
		if popErr := writeBuffer.PopContext("listOfCovReferences", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for listOfCovReferences")
		}
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

	if popErr := writeBuffer.PopContext("BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferences"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferences")
	}
	return nil
}

func (m *BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferences) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
