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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// VTCloseErrorListOfVTSessionIdentifiers is the corresponding interface of VTCloseErrorListOfVTSessionIdentifiers
type VTCloseErrorListOfVTSessionIdentifiers interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetListOfVtSessionIdentifiers returns ListOfVtSessionIdentifiers (property field)
	GetListOfVtSessionIdentifiers() []BACnetApplicationTagUnsignedInteger
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
}

// VTCloseErrorListOfVTSessionIdentifiersExactly can be used when we want exactly this type and not a type which fulfills VTCloseErrorListOfVTSessionIdentifiers.
// This is useful for switch cases.
type VTCloseErrorListOfVTSessionIdentifiersExactly interface {
	VTCloseErrorListOfVTSessionIdentifiers
	isVTCloseErrorListOfVTSessionIdentifiers() bool
}

// _VTCloseErrorListOfVTSessionIdentifiers is the data-structure of this message
type _VTCloseErrorListOfVTSessionIdentifiers struct {
	OpeningTag                 BACnetOpeningTag
	ListOfVtSessionIdentifiers []BACnetApplicationTagUnsignedInteger
	ClosingTag                 BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_VTCloseErrorListOfVTSessionIdentifiers) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_VTCloseErrorListOfVTSessionIdentifiers) GetListOfVtSessionIdentifiers() []BACnetApplicationTagUnsignedInteger {
	return m.ListOfVtSessionIdentifiers
}

func (m *_VTCloseErrorListOfVTSessionIdentifiers) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewVTCloseErrorListOfVTSessionIdentifiers factory function for _VTCloseErrorListOfVTSessionIdentifiers
func NewVTCloseErrorListOfVTSessionIdentifiers(openingTag BACnetOpeningTag, listOfVtSessionIdentifiers []BACnetApplicationTagUnsignedInteger, closingTag BACnetClosingTag, tagNumber uint8) *_VTCloseErrorListOfVTSessionIdentifiers {
	return &_VTCloseErrorListOfVTSessionIdentifiers{OpeningTag: openingTag, ListOfVtSessionIdentifiers: listOfVtSessionIdentifiers, ClosingTag: closingTag, TagNumber: tagNumber}
}

// Deprecated: use the interface for direct cast
func CastVTCloseErrorListOfVTSessionIdentifiers(structType any) VTCloseErrorListOfVTSessionIdentifiers {
	if casted, ok := structType.(VTCloseErrorListOfVTSessionIdentifiers); ok {
		return casted
	}
	if casted, ok := structType.(*VTCloseErrorListOfVTSessionIdentifiers); ok {
		return *casted
	}
	return nil
}

func (m *_VTCloseErrorListOfVTSessionIdentifiers) GetTypeName() string {
	return "VTCloseErrorListOfVTSessionIdentifiers"
}

func (m *_VTCloseErrorListOfVTSessionIdentifiers) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Array field
	if len(m.ListOfVtSessionIdentifiers) > 0 {
		for _, element := range m.ListOfVtSessionIdentifiers {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_VTCloseErrorListOfVTSessionIdentifiers) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func VTCloseErrorListOfVTSessionIdentifiersParse(theBytes []byte, tagNumber uint8) (VTCloseErrorListOfVTSessionIdentifiers, error) {
	return VTCloseErrorListOfVTSessionIdentifiersParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func VTCloseErrorListOfVTSessionIdentifiersParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (VTCloseErrorListOfVTSessionIdentifiers, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("VTCloseErrorListOfVTSessionIdentifiers"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for VTCloseErrorListOfVTSessionIdentifiers")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for openingTag")
	}
	_openingTag, _openingTagErr := BACnetOpeningTagParseWithBuffer(ctx, readBuffer, uint8(tagNumber))
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field of VTCloseErrorListOfVTSessionIdentifiers")
	}
	openingTag := _openingTag.(BACnetOpeningTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for openingTag")
	}

	// Array field (listOfVtSessionIdentifiers)
	if pullErr := readBuffer.PullContext("listOfVtSessionIdentifiers", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for listOfVtSessionIdentifiers")
	}
	// Terminated array
	var listOfVtSessionIdentifiers []BACnetApplicationTagUnsignedInteger
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, 1)) {
			_item, _err := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'listOfVtSessionIdentifiers' field of VTCloseErrorListOfVTSessionIdentifiers")
			}
			listOfVtSessionIdentifiers = append(listOfVtSessionIdentifiers, _item.(BACnetApplicationTagUnsignedInteger))
		}
	}
	if closeErr := readBuffer.CloseContext("listOfVtSessionIdentifiers", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for listOfVtSessionIdentifiers")
	}

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for closingTag")
	}
	_closingTag, _closingTagErr := BACnetClosingTagParseWithBuffer(ctx, readBuffer, uint8(tagNumber))
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field of VTCloseErrorListOfVTSessionIdentifiers")
	}
	closingTag := _closingTag.(BACnetClosingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for closingTag")
	}

	if closeErr := readBuffer.CloseContext("VTCloseErrorListOfVTSessionIdentifiers"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for VTCloseErrorListOfVTSessionIdentifiers")
	}

	// Create the instance
	return &_VTCloseErrorListOfVTSessionIdentifiers{
		TagNumber:                  tagNumber,
		OpeningTag:                 openingTag,
		ListOfVtSessionIdentifiers: listOfVtSessionIdentifiers,
		ClosingTag:                 closingTag,
	}, nil
}

func (m *_VTCloseErrorListOfVTSessionIdentifiers) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_VTCloseErrorListOfVTSessionIdentifiers) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("VTCloseErrorListOfVTSessionIdentifiers"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for VTCloseErrorListOfVTSessionIdentifiers")
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

	// Array Field (listOfVtSessionIdentifiers)
	if pushErr := writeBuffer.PushContext("listOfVtSessionIdentifiers", utils.WithRenderAsList(true)); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for listOfVtSessionIdentifiers")
	}
	for _curItem, _element := range m.GetListOfVtSessionIdentifiers() {
		_ = _curItem
		arrayCtx := utils.CreateArrayContext(ctx, len(m.GetListOfVtSessionIdentifiers()), _curItem)
		_ = arrayCtx
		_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
		if _elementErr != nil {
			return errors.Wrap(_elementErr, "Error serializing 'listOfVtSessionIdentifiers' field")
		}
	}
	if popErr := writeBuffer.PopContext("listOfVtSessionIdentifiers", utils.WithRenderAsList(true)); popErr != nil {
		return errors.Wrap(popErr, "Error popping for listOfVtSessionIdentifiers")
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

	if popErr := writeBuffer.PopContext("VTCloseErrorListOfVTSessionIdentifiers"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for VTCloseErrorListOfVTSessionIdentifiers")
	}
	return nil
}

////
// Arguments Getter

func (m *_VTCloseErrorListOfVTSessionIdentifiers) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_VTCloseErrorListOfVTSessionIdentifiers) isVTCloseErrorListOfVTSessionIdentifiers() bool {
	return true
}

func (m *_VTCloseErrorListOfVTSessionIdentifiers) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
