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

// BACnetAuthenticationPolicyList is the corresponding interface of BACnetAuthenticationPolicyList
type BACnetAuthenticationPolicyList interface {
	utils.LengthAware
	utils.Serializable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetEntries returns Entries (property field)
	GetEntries() []BACnetAuthenticationPolicyListEntry
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
}

// BACnetAuthenticationPolicyListExactly can be used when we want exactly this type and not a type which fulfills BACnetAuthenticationPolicyList.
// This is useful for switch cases.
type BACnetAuthenticationPolicyListExactly interface {
	BACnetAuthenticationPolicyList
	isBACnetAuthenticationPolicyList() bool
}

// _BACnetAuthenticationPolicyList is the data-structure of this message
type _BACnetAuthenticationPolicyList struct {
	OpeningTag BACnetOpeningTag
	Entries    []BACnetAuthenticationPolicyListEntry
	ClosingTag BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetAuthenticationPolicyList) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetAuthenticationPolicyList) GetEntries() []BACnetAuthenticationPolicyListEntry {
	return m.Entries
}

func (m *_BACnetAuthenticationPolicyList) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetAuthenticationPolicyList factory function for _BACnetAuthenticationPolicyList
func NewBACnetAuthenticationPolicyList(openingTag BACnetOpeningTag, entries []BACnetAuthenticationPolicyListEntry, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetAuthenticationPolicyList {
	return &_BACnetAuthenticationPolicyList{OpeningTag: openingTag, Entries: entries, ClosingTag: closingTag, TagNumber: tagNumber}
}

// Deprecated: use the interface for direct cast
func CastBACnetAuthenticationPolicyList(structType interface{}) BACnetAuthenticationPolicyList {
	if casted, ok := structType.(BACnetAuthenticationPolicyList); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetAuthenticationPolicyList); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetAuthenticationPolicyList) GetTypeName() string {
	return "BACnetAuthenticationPolicyList"
}

func (m *_BACnetAuthenticationPolicyList) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetAuthenticationPolicyList) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits()

	// Array field
	if len(m.Entries) > 0 {
		for _, element := range m.Entries {
			lengthInBits += element.GetLengthInBits()
		}
	}

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetAuthenticationPolicyList) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetAuthenticationPolicyListParse(readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetAuthenticationPolicyList, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetAuthenticationPolicyList"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetAuthenticationPolicyList")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for openingTag")
	}
	_openingTag, _openingTagErr := BACnetOpeningTagParse(readBuffer, uint8(tagNumber))
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field of BACnetAuthenticationPolicyList")
	}
	openingTag := _openingTag.(BACnetOpeningTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for openingTag")
	}

	// Array field (entries)
	if pullErr := readBuffer.PullContext("entries", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for entries")
	}
	// Terminated array
	var entries []BACnetAuthenticationPolicyListEntry
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetAuthenticationPolicyListEntryParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'entries' field of BACnetAuthenticationPolicyList")
			}
			entries = append(entries, _item.(BACnetAuthenticationPolicyListEntry))

		}
	}
	if closeErr := readBuffer.CloseContext("entries", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for entries")
	}

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for closingTag")
	}
	_closingTag, _closingTagErr := BACnetClosingTagParse(readBuffer, uint8(tagNumber))
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field of BACnetAuthenticationPolicyList")
	}
	closingTag := _closingTag.(BACnetClosingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for closingTag")
	}

	if closeErr := readBuffer.CloseContext("BACnetAuthenticationPolicyList"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetAuthenticationPolicyList")
	}

	// Create the instance
	return NewBACnetAuthenticationPolicyList(openingTag, entries, closingTag, tagNumber), nil
}

func (m *_BACnetAuthenticationPolicyList) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetAuthenticationPolicyList"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetAuthenticationPolicyList")
	}

	// Simple Field (openingTag)
	if pushErr := writeBuffer.PushContext("openingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for openingTag")
	}
	_openingTagErr := writeBuffer.WriteSerializable(m.GetOpeningTag())
	if popErr := writeBuffer.PopContext("openingTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for openingTag")
	}
	if _openingTagErr != nil {
		return errors.Wrap(_openingTagErr, "Error serializing 'openingTag' field")
	}

	// Array Field (entries)
	if pushErr := writeBuffer.PushContext("entries", utils.WithRenderAsList(true)); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for entries")
	}
	for _, _element := range m.GetEntries() {
		_elementErr := writeBuffer.WriteSerializable(_element)
		if _elementErr != nil {
			return errors.Wrap(_elementErr, "Error serializing 'entries' field")
		}
	}
	if popErr := writeBuffer.PopContext("entries", utils.WithRenderAsList(true)); popErr != nil {
		return errors.Wrap(popErr, "Error popping for entries")
	}

	// Simple Field (closingTag)
	if pushErr := writeBuffer.PushContext("closingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for closingTag")
	}
	_closingTagErr := writeBuffer.WriteSerializable(m.GetClosingTag())
	if popErr := writeBuffer.PopContext("closingTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for closingTag")
	}
	if _closingTagErr != nil {
		return errors.Wrap(_closingTagErr, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetAuthenticationPolicyList"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetAuthenticationPolicyList")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetAuthenticationPolicyList) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetAuthenticationPolicyList) isBACnetAuthenticationPolicyList() bool {
	return true
}

func (m *_BACnetAuthenticationPolicyList) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
