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

// BACnetConstructedDataEventLogLogBuffer is the data-structure of this message
type BACnetConstructedDataEventLogLogBuffer struct {
	*BACnetConstructedData
	FloorText []*BACnetEventLogRecord

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataEventLogLogBuffer is the corresponding interface of BACnetConstructedDataEventLogLogBuffer
type IBACnetConstructedDataEventLogLogBuffer interface {
	IBACnetConstructedData
	// GetFloorText returns FloorText (property field)
	GetFloorText() []*BACnetEventLogRecord
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

func (m *BACnetConstructedDataEventLogLogBuffer) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_EVENT_LOG
}

func (m *BACnetConstructedDataEventLogLogBuffer) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LOG_BUFFER
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataEventLogLogBuffer) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataEventLogLogBuffer) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataEventLogLogBuffer) GetFloorText() []*BACnetEventLogRecord {
	return m.FloorText
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataEventLogLogBuffer factory function for BACnetConstructedDataEventLogLogBuffer
func NewBACnetConstructedDataEventLogLogBuffer(floorText []*BACnetEventLogRecord, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataEventLogLogBuffer {
	_result := &BACnetConstructedDataEventLogLogBuffer{
		FloorText:             floorText,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataEventLogLogBuffer(structType interface{}) *BACnetConstructedDataEventLogLogBuffer {
	if casted, ok := structType.(BACnetConstructedDataEventLogLogBuffer); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataEventLogLogBuffer); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataEventLogLogBuffer(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataEventLogLogBuffer(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataEventLogLogBuffer) GetTypeName() string {
	return "BACnetConstructedDataEventLogLogBuffer"
}

func (m *BACnetConstructedDataEventLogLogBuffer) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataEventLogLogBuffer) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.FloorText) > 0 {
		for _, element := range m.FloorText {
			lengthInBits += element.GetLengthInBits()
		}
	}

	return lengthInBits
}

func (m *BACnetConstructedDataEventLogLogBuffer) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataEventLogLogBufferParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataEventLogLogBuffer, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataEventLogLogBuffer"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataEventLogLogBuffer")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (floorText)
	if pullErr := readBuffer.PullContext("floorText", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for floorText")
	}
	// Terminated array
	floorText := make([]*BACnetEventLogRecord, 0)
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetEventLogRecordParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'floorText' field")
			}
			floorText = append(floorText, CastBACnetEventLogRecord(_item))

		}
	}
	if closeErr := readBuffer.CloseContext("floorText", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for floorText")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataEventLogLogBuffer"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataEventLogLogBuffer")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataEventLogLogBuffer{
		FloorText:             floorText,
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataEventLogLogBuffer) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataEventLogLogBuffer"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataEventLogLogBuffer")
		}

		// Array Field (floorText)
		if m.FloorText != nil {
			if pushErr := writeBuffer.PushContext("floorText", utils.WithRenderAsList(true)); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for floorText")
			}
			for _, _element := range m.FloorText {
				_elementErr := writeBuffer.WriteSerializable(_element)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'floorText' field")
				}
			}
			if popErr := writeBuffer.PopContext("floorText", utils.WithRenderAsList(true)); popErr != nil {
				return errors.Wrap(popErr, "Error popping for floorText")
			}
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataEventLogLogBuffer"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataEventLogLogBuffer")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataEventLogLogBuffer) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
