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
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetAuthenticationPolicyList is the corresponding interface of BACnetAuthenticationPolicyList
type BACnetAuthenticationPolicyList interface {
	fmt.Stringer
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
func CastBACnetAuthenticationPolicyList(structType any) BACnetAuthenticationPolicyList {
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

func (m *_BACnetAuthenticationPolicyList) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Array field
	if len(m.Entries) > 0 {
		for _, element := range m.Entries {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetAuthenticationPolicyList) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetAuthenticationPolicyListParse(ctx context.Context, theBytes []byte, tagNumber uint8) (BACnetAuthenticationPolicyList, error) {
	return BACnetAuthenticationPolicyListParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetAuthenticationPolicyListParseWithBufferProducer(tagNumber uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAuthenticationPolicyList, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAuthenticationPolicyList, error) {
		return BACnetAuthenticationPolicyListParseWithBuffer(ctx, readBuffer, tagNumber)
	}
}

func BACnetAuthenticationPolicyListParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetAuthenticationPolicyList, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetAuthenticationPolicyList"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetAuthenticationPolicyList")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}

	entries, err := ReadTerminatedArrayField[BACnetAuthenticationPolicyListEntry](ctx, "entries", ReadComplex[BACnetAuthenticationPolicyListEntry](BACnetAuthenticationPolicyListEntryParseWithBuffer, readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'entries' field"))
	}

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}

	if closeErr := readBuffer.CloseContext("BACnetAuthenticationPolicyList"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetAuthenticationPolicyList")
	}

	// Create the instance
	return &_BACnetAuthenticationPolicyList{
		TagNumber:  tagNumber,
		OpeningTag: openingTag,
		Entries:    entries,
		ClosingTag: closingTag,
	}, nil
}

func (m *_BACnetAuthenticationPolicyList) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetAuthenticationPolicyList) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetAuthenticationPolicyList"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetAuthenticationPolicyList")
	}

	// Simple Field (openingTag)
	if pushErr := writeBuffer.PushContext("openingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for openingTag")
	}
	_openingTagErr := writeBuffer.WriteSerializable(ctx, m.GetOpeningTag())
	if popErr := writeBuffer.PopContext("openingTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for openingTag")
	}
	if _openingTagErr != nil {
		return errors.Wrap(_openingTagErr, "Error serializing 'openingTag' field")
	}

	if err := WriteComplexTypeArrayField(ctx, "entries", m.GetEntries(), writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'entries' field")
	}

	// Simple Field (closingTag)
	if pushErr := writeBuffer.PushContext("closingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for closingTag")
	}
	_closingTagErr := writeBuffer.WriteSerializable(ctx, m.GetClosingTag())
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
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
