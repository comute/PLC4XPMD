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
	"github.com/rs/zerolog/log"
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataEventMessageTexts is the data-structure of this message
type BACnetConstructedDataEventMessageTexts struct {
	*BACnetConstructedData
	NumberOfDataElements *BACnetApplicationTagUnsignedInteger
	EventMessageTexts    []*BACnetOptionalCharacterString

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataEventMessageTexts is the corresponding interface of BACnetConstructedDataEventMessageTexts
type IBACnetConstructedDataEventMessageTexts interface {
	IBACnetConstructedData
	// GetNumberOfDataElements returns NumberOfDataElements (property field)
	GetNumberOfDataElements() *BACnetApplicationTagUnsignedInteger
	// GetEventMessageTexts returns EventMessageTexts (property field)
	GetEventMessageTexts() []*BACnetOptionalCharacterString
	// GetZero returns Zero (virtual field)
	GetZero() uint64
	// GetToOffnormalText returns ToOffnormalText (virtual field)
	GetToOffnormalText() *BACnetOptionalCharacterString
	// GetToFaultText returns ToFaultText (virtual field)
	GetToFaultText() *BACnetOptionalCharacterString
	// GetToNormalText returns ToNormalText (virtual field)
	GetToNormalText() *BACnetOptionalCharacterString
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

func (m *BACnetConstructedDataEventMessageTexts) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataEventMessageTexts) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_EVENT_MESSAGE_TEXTS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataEventMessageTexts) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataEventMessageTexts) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataEventMessageTexts) GetNumberOfDataElements() *BACnetApplicationTagUnsignedInteger {
	return m.NumberOfDataElements
}

func (m *BACnetConstructedDataEventMessageTexts) GetEventMessageTexts() []*BACnetOptionalCharacterString {
	return m.EventMessageTexts
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *BACnetConstructedDataEventMessageTexts) GetZero() uint64 {
	numberOfDataElements := m.NumberOfDataElements
	_ = numberOfDataElements
	return uint64(uint64(0))
}

func (m *BACnetConstructedDataEventMessageTexts) GetToOffnormalText() *BACnetOptionalCharacterString {
	numberOfDataElements := m.NumberOfDataElements
	_ = numberOfDataElements
	return CastBACnetOptionalCharacterString(CastBACnetOptionalCharacterString(utils.InlineIf(bool((len(m.GetEventMessageTexts())) == (3)), func() interface{} { return CastBACnetOptionalCharacterString(m.GetEventMessageTexts()[0]) }, func() interface{} { return CastBACnetOptionalCharacterString(nil) })))
}

func (m *BACnetConstructedDataEventMessageTexts) GetToFaultText() *BACnetOptionalCharacterString {
	numberOfDataElements := m.NumberOfDataElements
	_ = numberOfDataElements
	return CastBACnetOptionalCharacterString(CastBACnetOptionalCharacterString(utils.InlineIf(bool((len(m.GetEventMessageTexts())) == (3)), func() interface{} { return CastBACnetOptionalCharacterString(m.GetEventMessageTexts()[1]) }, func() interface{} { return CastBACnetOptionalCharacterString(nil) })))
}

func (m *BACnetConstructedDataEventMessageTexts) GetToNormalText() *BACnetOptionalCharacterString {
	numberOfDataElements := m.NumberOfDataElements
	_ = numberOfDataElements
	return CastBACnetOptionalCharacterString(CastBACnetOptionalCharacterString(utils.InlineIf(bool((len(m.GetEventMessageTexts())) == (3)), func() interface{} { return CastBACnetOptionalCharacterString(m.GetEventMessageTexts()[2]) }, func() interface{} { return CastBACnetOptionalCharacterString(nil) })))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataEventMessageTexts factory function for BACnetConstructedDataEventMessageTexts
func NewBACnetConstructedDataEventMessageTexts(numberOfDataElements *BACnetApplicationTagUnsignedInteger, eventMessageTexts []*BACnetOptionalCharacterString, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataEventMessageTexts {
	_result := &BACnetConstructedDataEventMessageTexts{
		NumberOfDataElements:  numberOfDataElements,
		EventMessageTexts:     eventMessageTexts,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataEventMessageTexts(structType interface{}) *BACnetConstructedDataEventMessageTexts {
	if casted, ok := structType.(BACnetConstructedDataEventMessageTexts); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataEventMessageTexts); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataEventMessageTexts(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataEventMessageTexts(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataEventMessageTexts) GetTypeName() string {
	return "BACnetConstructedDataEventMessageTexts"
}

func (m *BACnetConstructedDataEventMessageTexts) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataEventMessageTexts) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// A virtual field doesn't have any in- or output.

	// Optional Field (numberOfDataElements)
	if m.NumberOfDataElements != nil {
		lengthInBits += (*m.NumberOfDataElements).GetLengthInBits()
	}

	// Array field
	if len(m.EventMessageTexts) > 0 {
		for _, element := range m.EventMessageTexts {
			lengthInBits += element.GetLengthInBits()
		}
	}

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *BACnetConstructedDataEventMessageTexts) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataEventMessageTextsParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataEventMessageTexts, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataEventMessageTexts"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataEventMessageTexts")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Virtual field
	_zero := uint64(0)
	zero := uint64(_zero)
	_ = zero

	// Optional Field (numberOfDataElements) (Can be skipped, if a given expression evaluates to false)
	var numberOfDataElements *BACnetApplicationTagUnsignedInteger = nil
	if bool(bool((arrayIndexArgument) != (nil))) && bool(bool((arrayIndexArgument.GetActualValue()) == (zero))) {
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("numberOfDataElements"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for numberOfDataElements")
		}
		_val, _err := BACnetApplicationTagParse(readBuffer)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'numberOfDataElements' field")
		default:
			numberOfDataElements = CastBACnetApplicationTagUnsignedInteger(_val)
			if closeErr := readBuffer.CloseContext("numberOfDataElements"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for numberOfDataElements")
			}
		}
	}

	// Array field (eventMessageTexts)
	if pullErr := readBuffer.PullContext("eventMessageTexts", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for eventMessageTexts")
	}
	// Terminated array
	eventMessageTexts := make([]*BACnetOptionalCharacterString, 0)
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetOptionalCharacterStringParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'eventMessageTexts' field")
			}
			eventMessageTexts = append(eventMessageTexts, CastBACnetOptionalCharacterString(_item))

		}
	}
	if closeErr := readBuffer.CloseContext("eventMessageTexts", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for eventMessageTexts")
	}

	// Virtual field
	_toOffnormalText := CastBACnetOptionalCharacterString(utils.InlineIf(bool((len(eventMessageTexts)) == (3)), func() interface{} { return CastBACnetOptionalCharacterString(eventMessageTexts[0]) }, func() interface{} { return CastBACnetOptionalCharacterString(nil) }))
	toOffnormalText := CastBACnetOptionalCharacterString(_toOffnormalText)
	_ = toOffnormalText

	// Virtual field
	_toFaultText := CastBACnetOptionalCharacterString(utils.InlineIf(bool((len(eventMessageTexts)) == (3)), func() interface{} { return CastBACnetOptionalCharacterString(eventMessageTexts[1]) }, func() interface{} { return CastBACnetOptionalCharacterString(nil) }))
	toFaultText := CastBACnetOptionalCharacterString(_toFaultText)
	_ = toFaultText

	// Virtual field
	_toNormalText := CastBACnetOptionalCharacterString(utils.InlineIf(bool((len(eventMessageTexts)) == (3)), func() interface{} { return CastBACnetOptionalCharacterString(eventMessageTexts[2]) }, func() interface{} { return CastBACnetOptionalCharacterString(nil) }))
	toNormalText := CastBACnetOptionalCharacterString(_toNormalText)
	_ = toNormalText

	// Validation
	if !(bool(bool((arrayIndexArgument) != (nil))) || bool(bool((len(eventMessageTexts)) == (3)))) {
		return nil, errors.WithStack(utils.ParseValidationError{"eventMessageTexts should have exactly 3 values"})
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataEventMessageTexts"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataEventMessageTexts")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataEventMessageTexts{
		NumberOfDataElements:  CastBACnetApplicationTagUnsignedInteger(numberOfDataElements),
		EventMessageTexts:     eventMessageTexts,
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataEventMessageTexts) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataEventMessageTexts"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataEventMessageTexts")
		}
		// Virtual field
		if _zeroErr := writeBuffer.WriteVirtual("zero", m.GetZero()); _zeroErr != nil {
			return errors.Wrap(_zeroErr, "Error serializing 'zero' field")
		}

		// Optional Field (numberOfDataElements) (Can be skipped, if the value is null)
		var numberOfDataElements *BACnetApplicationTagUnsignedInteger = nil
		if m.NumberOfDataElements != nil {
			if pushErr := writeBuffer.PushContext("numberOfDataElements"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for numberOfDataElements")
			}
			numberOfDataElements = m.NumberOfDataElements
			_numberOfDataElementsErr := numberOfDataElements.Serialize(writeBuffer)
			if popErr := writeBuffer.PopContext("numberOfDataElements"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for numberOfDataElements")
			}
			if _numberOfDataElementsErr != nil {
				return errors.Wrap(_numberOfDataElementsErr, "Error serializing 'numberOfDataElements' field")
			}
		}

		// Array Field (eventMessageTexts)
		if m.EventMessageTexts != nil {
			if pushErr := writeBuffer.PushContext("eventMessageTexts", utils.WithRenderAsList(true)); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for eventMessageTexts")
			}
			for _, _element := range m.EventMessageTexts {
				_elementErr := _element.Serialize(writeBuffer)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'eventMessageTexts' field")
				}
			}
			if popErr := writeBuffer.PopContext("eventMessageTexts", utils.WithRenderAsList(true)); popErr != nil {
				return errors.Wrap(popErr, "Error popping for eventMessageTexts")
			}
		}
		// Virtual field
		if _toOffnormalTextErr := writeBuffer.WriteVirtual("toOffnormalText", m.GetToOffnormalText()); _toOffnormalTextErr != nil {
			return errors.Wrap(_toOffnormalTextErr, "Error serializing 'toOffnormalText' field")
		}
		// Virtual field
		if _toFaultTextErr := writeBuffer.WriteVirtual("toFaultText", m.GetToFaultText()); _toFaultTextErr != nil {
			return errors.Wrap(_toFaultTextErr, "Error serializing 'toFaultText' field")
		}
		// Virtual field
		if _toNormalTextErr := writeBuffer.WriteVirtual("toNormalText", m.GetToNormalText()); _toNormalTextErr != nil {
			return errors.Wrap(_toNormalTextErr, "Error serializing 'toNormalText' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataEventMessageTexts"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataEventMessageTexts")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataEventMessageTexts) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
