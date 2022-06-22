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
	"github.com/rs/zerolog/log"
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataEventTimeStamps is the corresponding interface of BACnetConstructedDataEventTimeStamps
type BACnetConstructedDataEventTimeStamps interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetNumberOfDataElements returns NumberOfDataElements (property field)
	GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger
	// GetEventTimeStamps returns EventTimeStamps (property field)
	GetEventTimeStamps() []BACnetTimeStamp
	// GetZero returns Zero (virtual field)
	GetZero() uint64
	// GetToOffnormal returns ToOffnormal (virtual field)
	GetToOffnormal() BACnetTimeStamp
	// GetToFault returns ToFault (virtual field)
	GetToFault() BACnetTimeStamp
	// GetToNormal returns ToNormal (virtual field)
	GetToNormal() BACnetTimeStamp
}

// BACnetConstructedDataEventTimeStampsExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataEventTimeStamps.
// This is useful for switch cases.
type BACnetConstructedDataEventTimeStampsExactly interface {
	BACnetConstructedDataEventTimeStamps
	isBACnetConstructedDataEventTimeStamps() bool
}

// _BACnetConstructedDataEventTimeStamps is the data-structure of this message
type _BACnetConstructedDataEventTimeStamps struct {
	*_BACnetConstructedData
	NumberOfDataElements BACnetApplicationTagUnsignedInteger
	EventTimeStamps      []BACnetTimeStamp

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument BACnetTagPayloadUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataEventTimeStamps) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataEventTimeStamps) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_EVENT_TIME_STAMPS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataEventTimeStamps) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataEventTimeStamps) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataEventTimeStamps) GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger {
	return m.NumberOfDataElements
}

func (m *_BACnetConstructedDataEventTimeStamps) GetEventTimeStamps() []BACnetTimeStamp {
	return m.EventTimeStamps
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataEventTimeStamps) GetZero() uint64 {
	numberOfDataElements := m.NumberOfDataElements
	_ = numberOfDataElements
	return uint64(uint64(0))
}

func (m *_BACnetConstructedDataEventTimeStamps) GetToOffnormal() BACnetTimeStamp {
	numberOfDataElements := m.NumberOfDataElements
	_ = numberOfDataElements
	return CastBACnetTimeStamp(CastBACnetTimeStamp(utils.InlineIf(bool((len(m.GetEventTimeStamps())) == (3)), func() interface{} { return CastBACnetTimeStamp(m.GetEventTimeStamps()[0]) }, func() interface{} { return CastBACnetTimeStamp(nil) })))
}

func (m *_BACnetConstructedDataEventTimeStamps) GetToFault() BACnetTimeStamp {
	numberOfDataElements := m.NumberOfDataElements
	_ = numberOfDataElements
	return CastBACnetTimeStamp(CastBACnetTimeStamp(utils.InlineIf(bool((len(m.GetEventTimeStamps())) == (3)), func() interface{} { return CastBACnetTimeStamp(m.GetEventTimeStamps()[1]) }, func() interface{} { return CastBACnetTimeStamp(nil) })))
}

func (m *_BACnetConstructedDataEventTimeStamps) GetToNormal() BACnetTimeStamp {
	numberOfDataElements := m.NumberOfDataElements
	_ = numberOfDataElements
	return CastBACnetTimeStamp(CastBACnetTimeStamp(utils.InlineIf(bool((len(m.GetEventTimeStamps())) == (3)), func() interface{} { return CastBACnetTimeStamp(m.GetEventTimeStamps()[2]) }, func() interface{} { return CastBACnetTimeStamp(nil) })))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataEventTimeStamps factory function for _BACnetConstructedDataEventTimeStamps
func NewBACnetConstructedDataEventTimeStamps(numberOfDataElements BACnetApplicationTagUnsignedInteger, eventTimeStamps []BACnetTimeStamp, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataEventTimeStamps {
	_result := &_BACnetConstructedDataEventTimeStamps{
		NumberOfDataElements:   numberOfDataElements,
		EventTimeStamps:        eventTimeStamps,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataEventTimeStamps(structType interface{}) BACnetConstructedDataEventTimeStamps {
	if casted, ok := structType.(BACnetConstructedDataEventTimeStamps); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataEventTimeStamps); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataEventTimeStamps) GetTypeName() string {
	return "BACnetConstructedDataEventTimeStamps"
}

func (m *_BACnetConstructedDataEventTimeStamps) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataEventTimeStamps) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// A virtual field doesn't have any in- or output.

	// Optional Field (numberOfDataElements)
	if m.NumberOfDataElements != nil {
		lengthInBits += m.NumberOfDataElements.GetLengthInBits()
	}

	// Array field
	if len(m.EventTimeStamps) > 0 {
		for _, element := range m.EventTimeStamps {
			lengthInBits += element.GetLengthInBits()
		}
	}

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataEventTimeStamps) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataEventTimeStampsParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataEventTimeStamps, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataEventTimeStamps"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataEventTimeStamps")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Virtual field
	_zero := uint64(0)
	zero := uint64(_zero)
	_ = zero

	// Optional Field (numberOfDataElements) (Can be skipped, if a given expression evaluates to false)
	var numberOfDataElements BACnetApplicationTagUnsignedInteger = nil
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
			numberOfDataElements = _val.(BACnetApplicationTagUnsignedInteger)
			if closeErr := readBuffer.CloseContext("numberOfDataElements"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for numberOfDataElements")
			}
		}
	}

	// Array field (eventTimeStamps)
	if pullErr := readBuffer.PullContext("eventTimeStamps", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for eventTimeStamps")
	}
	// Terminated array
	eventTimeStamps := make([]BACnetTimeStamp, 0)
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetTimeStampParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'eventTimeStamps' field")
			}
			eventTimeStamps = append(eventTimeStamps, _item.(BACnetTimeStamp))

		}
	}
	if closeErr := readBuffer.CloseContext("eventTimeStamps", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for eventTimeStamps")
	}

	// Virtual field
	_toOffnormal := CastBACnetTimeStamp(utils.InlineIf(bool((len(eventTimeStamps)) == (3)), func() interface{} { return CastBACnetTimeStamp(eventTimeStamps[0]) }, func() interface{} { return CastBACnetTimeStamp(nil) }))
	toOffnormal := _toOffnormal
	_ = toOffnormal

	// Virtual field
	_toFault := CastBACnetTimeStamp(utils.InlineIf(bool((len(eventTimeStamps)) == (3)), func() interface{} { return CastBACnetTimeStamp(eventTimeStamps[1]) }, func() interface{} { return CastBACnetTimeStamp(nil) }))
	toFault := _toFault
	_ = toFault

	// Virtual field
	_toNormal := CastBACnetTimeStamp(utils.InlineIf(bool((len(eventTimeStamps)) == (3)), func() interface{} { return CastBACnetTimeStamp(eventTimeStamps[2]) }, func() interface{} { return CastBACnetTimeStamp(nil) }))
	toNormal := _toNormal
	_ = toNormal

	// Validation
	if !(bool(bool((arrayIndexArgument) != (nil))) || bool(bool((len(eventTimeStamps)) == (3)))) {
		return nil, errors.WithStack(utils.ParseValidationError{"eventTimeStamps should have exactly 3 values"})
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataEventTimeStamps"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataEventTimeStamps")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataEventTimeStamps{
		NumberOfDataElements:   numberOfDataElements,
		EventTimeStamps:        eventTimeStamps,
		_BACnetConstructedData: &_BACnetConstructedData{},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataEventTimeStamps) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataEventTimeStamps"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataEventTimeStamps")
		}
		// Virtual field
		if _zeroErr := writeBuffer.WriteVirtual("zero", m.GetZero()); _zeroErr != nil {
			return errors.Wrap(_zeroErr, "Error serializing 'zero' field")
		}

		// Optional Field (numberOfDataElements) (Can be skipped, if the value is null)
		var numberOfDataElements BACnetApplicationTagUnsignedInteger = nil
		if m.GetNumberOfDataElements() != nil {
			if pushErr := writeBuffer.PushContext("numberOfDataElements"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for numberOfDataElements")
			}
			numberOfDataElements = m.GetNumberOfDataElements()
			_numberOfDataElementsErr := writeBuffer.WriteSerializable(numberOfDataElements)
			if popErr := writeBuffer.PopContext("numberOfDataElements"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for numberOfDataElements")
			}
			if _numberOfDataElementsErr != nil {
				return errors.Wrap(_numberOfDataElementsErr, "Error serializing 'numberOfDataElements' field")
			}
		}

		// Array Field (eventTimeStamps)
		if m.GetEventTimeStamps() != nil {
			if pushErr := writeBuffer.PushContext("eventTimeStamps", utils.WithRenderAsList(true)); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for eventTimeStamps")
			}
			for _, _element := range m.GetEventTimeStamps() {
				_elementErr := writeBuffer.WriteSerializable(_element)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'eventTimeStamps' field")
				}
			}
			if popErr := writeBuffer.PopContext("eventTimeStamps", utils.WithRenderAsList(true)); popErr != nil {
				return errors.Wrap(popErr, "Error popping for eventTimeStamps")
			}
		}
		// Virtual field
		if _toOffnormalErr := writeBuffer.WriteVirtual("toOffnormal", m.GetToOffnormal()); _toOffnormalErr != nil {
			return errors.Wrap(_toOffnormalErr, "Error serializing 'toOffnormal' field")
		}
		// Virtual field
		if _toFaultErr := writeBuffer.WriteVirtual("toFault", m.GetToFault()); _toFaultErr != nil {
			return errors.Wrap(_toFaultErr, "Error serializing 'toFault' field")
		}
		// Virtual field
		if _toNormalErr := writeBuffer.WriteVirtual("toNormal", m.GetToNormal()); _toNormalErr != nil {
			return errors.Wrap(_toNormalErr, "Error serializing 'toNormal' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataEventTimeStamps"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataEventTimeStamps")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataEventTimeStamps) isBACnetConstructedDataEventTimeStamps() bool {
	return true
}

func (m *_BACnetConstructedDataEventTimeStamps) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
