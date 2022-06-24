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

// BACnetNotificationParametersUnsignedOutOfRange is the corresponding interface of BACnetNotificationParametersUnsignedOutOfRange
type BACnetNotificationParametersUnsignedOutOfRange interface {
	utils.LengthAware
	utils.Serializable
	BACnetNotificationParameters
	// GetInnerOpeningTag returns InnerOpeningTag (property field)
	GetInnerOpeningTag() BACnetOpeningTag
	// GetExceedingValue returns ExceedingValue (property field)
	GetExceedingValue() BACnetContextTagUnsignedInteger
	// GetStatusFlags returns StatusFlags (property field)
	GetStatusFlags() BACnetStatusFlagsTagged
	// GetDeadband returns Deadband (property field)
	GetDeadband() BACnetContextTagUnsignedInteger
	// GetExceededLimit returns ExceededLimit (property field)
	GetExceededLimit() BACnetContextTagUnsignedInteger
	// GetInnerClosingTag returns InnerClosingTag (property field)
	GetInnerClosingTag() BACnetClosingTag
}

// BACnetNotificationParametersUnsignedOutOfRangeExactly can be used when we want exactly this type and not a type which fulfills BACnetNotificationParametersUnsignedOutOfRange.
// This is useful for switch cases.
type BACnetNotificationParametersUnsignedOutOfRangeExactly interface {
	BACnetNotificationParametersUnsignedOutOfRange
	isBACnetNotificationParametersUnsignedOutOfRange() bool
}

// _BACnetNotificationParametersUnsignedOutOfRange is the data-structure of this message
type _BACnetNotificationParametersUnsignedOutOfRange struct {
	*_BACnetNotificationParameters
	InnerOpeningTag BACnetOpeningTag
	ExceedingValue  BACnetContextTagUnsignedInteger
	StatusFlags     BACnetStatusFlagsTagged
	Deadband        BACnetContextTagUnsignedInteger
	ExceededLimit   BACnetContextTagUnsignedInteger
	InnerClosingTag BACnetClosingTag
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetNotificationParametersUnsignedOutOfRange) InitializeParent(parent BACnetNotificationParameters, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetNotificationParametersUnsignedOutOfRange) GetParent() BACnetNotificationParameters {
	return m._BACnetNotificationParameters
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetNotificationParametersUnsignedOutOfRange) GetInnerOpeningTag() BACnetOpeningTag {
	return m.InnerOpeningTag
}

func (m *_BACnetNotificationParametersUnsignedOutOfRange) GetExceedingValue() BACnetContextTagUnsignedInteger {
	return m.ExceedingValue
}

func (m *_BACnetNotificationParametersUnsignedOutOfRange) GetStatusFlags() BACnetStatusFlagsTagged {
	return m.StatusFlags
}

func (m *_BACnetNotificationParametersUnsignedOutOfRange) GetDeadband() BACnetContextTagUnsignedInteger {
	return m.Deadband
}

func (m *_BACnetNotificationParametersUnsignedOutOfRange) GetExceededLimit() BACnetContextTagUnsignedInteger {
	return m.ExceededLimit
}

func (m *_BACnetNotificationParametersUnsignedOutOfRange) GetInnerClosingTag() BACnetClosingTag {
	return m.InnerClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetNotificationParametersUnsignedOutOfRange factory function for _BACnetNotificationParametersUnsignedOutOfRange
func NewBACnetNotificationParametersUnsignedOutOfRange(innerOpeningTag BACnetOpeningTag, exceedingValue BACnetContextTagUnsignedInteger, statusFlags BACnetStatusFlagsTagged, deadband BACnetContextTagUnsignedInteger, exceededLimit BACnetContextTagUnsignedInteger, innerClosingTag BACnetClosingTag, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, objectTypeArgument BACnetObjectType) *_BACnetNotificationParametersUnsignedOutOfRange {
	_result := &_BACnetNotificationParametersUnsignedOutOfRange{
		InnerOpeningTag:               innerOpeningTag,
		ExceedingValue:                exceedingValue,
		StatusFlags:                   statusFlags,
		Deadband:                      deadband,
		ExceededLimit:                 exceededLimit,
		InnerClosingTag:               innerClosingTag,
		_BACnetNotificationParameters: NewBACnetNotificationParameters(openingTag, peekedTagHeader, closingTag, tagNumber, objectTypeArgument),
	}
	_result._BACnetNotificationParameters._BACnetNotificationParametersChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetNotificationParametersUnsignedOutOfRange(structType interface{}) BACnetNotificationParametersUnsignedOutOfRange {
	if casted, ok := structType.(BACnetNotificationParametersUnsignedOutOfRange); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetNotificationParametersUnsignedOutOfRange); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetNotificationParametersUnsignedOutOfRange) GetTypeName() string {
	return "BACnetNotificationParametersUnsignedOutOfRange"
}

func (m *_BACnetNotificationParametersUnsignedOutOfRange) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetNotificationParametersUnsignedOutOfRange) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (innerOpeningTag)
	lengthInBits += m.InnerOpeningTag.GetLengthInBits()

	// Simple field (exceedingValue)
	lengthInBits += m.ExceedingValue.GetLengthInBits()

	// Simple field (statusFlags)
	lengthInBits += m.StatusFlags.GetLengthInBits()

	// Simple field (deadband)
	lengthInBits += m.Deadband.GetLengthInBits()

	// Simple field (exceededLimit)
	lengthInBits += m.ExceededLimit.GetLengthInBits()

	// Simple field (innerClosingTag)
	lengthInBits += m.InnerClosingTag.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetNotificationParametersUnsignedOutOfRange) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetNotificationParametersUnsignedOutOfRangeParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, peekedTagNumber uint8) (BACnetNotificationParametersUnsignedOutOfRange, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetNotificationParametersUnsignedOutOfRange"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetNotificationParametersUnsignedOutOfRange")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (innerOpeningTag)
	if pullErr := readBuffer.PullContext("innerOpeningTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for innerOpeningTag")
	}
	_innerOpeningTag, _innerOpeningTagErr := BACnetOpeningTagParse(readBuffer, uint8(peekedTagNumber))
	if _innerOpeningTagErr != nil {
		return nil, errors.Wrap(_innerOpeningTagErr, "Error parsing 'innerOpeningTag' field")
	}
	innerOpeningTag := _innerOpeningTag.(BACnetOpeningTag)
	if closeErr := readBuffer.CloseContext("innerOpeningTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for innerOpeningTag")
	}

	// Simple Field (exceedingValue)
	if pullErr := readBuffer.PullContext("exceedingValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for exceedingValue")
	}
	_exceedingValue, _exceedingValueErr := BACnetContextTagParse(readBuffer, uint8(uint8(0)), BACnetDataType(BACnetDataType_UNSIGNED_INTEGER))
	if _exceedingValueErr != nil {
		return nil, errors.Wrap(_exceedingValueErr, "Error parsing 'exceedingValue' field")
	}
	exceedingValue := _exceedingValue.(BACnetContextTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("exceedingValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for exceedingValue")
	}

	// Simple Field (statusFlags)
	if pullErr := readBuffer.PullContext("statusFlags"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for statusFlags")
	}
	_statusFlags, _statusFlagsErr := BACnetStatusFlagsTaggedParse(readBuffer, uint8(uint8(1)), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _statusFlagsErr != nil {
		return nil, errors.Wrap(_statusFlagsErr, "Error parsing 'statusFlags' field")
	}
	statusFlags := _statusFlags.(BACnetStatusFlagsTagged)
	if closeErr := readBuffer.CloseContext("statusFlags"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for statusFlags")
	}

	// Simple Field (deadband)
	if pullErr := readBuffer.PullContext("deadband"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for deadband")
	}
	_deadband, _deadbandErr := BACnetContextTagParse(readBuffer, uint8(uint8(2)), BACnetDataType(BACnetDataType_UNSIGNED_INTEGER))
	if _deadbandErr != nil {
		return nil, errors.Wrap(_deadbandErr, "Error parsing 'deadband' field")
	}
	deadband := _deadband.(BACnetContextTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("deadband"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for deadband")
	}

	// Simple Field (exceededLimit)
	if pullErr := readBuffer.PullContext("exceededLimit"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for exceededLimit")
	}
	_exceededLimit, _exceededLimitErr := BACnetContextTagParse(readBuffer, uint8(uint8(3)), BACnetDataType(BACnetDataType_UNSIGNED_INTEGER))
	if _exceededLimitErr != nil {
		return nil, errors.Wrap(_exceededLimitErr, "Error parsing 'exceededLimit' field")
	}
	exceededLimit := _exceededLimit.(BACnetContextTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("exceededLimit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for exceededLimit")
	}

	// Simple Field (innerClosingTag)
	if pullErr := readBuffer.PullContext("innerClosingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for innerClosingTag")
	}
	_innerClosingTag, _innerClosingTagErr := BACnetClosingTagParse(readBuffer, uint8(peekedTagNumber))
	if _innerClosingTagErr != nil {
		return nil, errors.Wrap(_innerClosingTagErr, "Error parsing 'innerClosingTag' field")
	}
	innerClosingTag := _innerClosingTag.(BACnetClosingTag)
	if closeErr := readBuffer.CloseContext("innerClosingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for innerClosingTag")
	}

	if closeErr := readBuffer.CloseContext("BACnetNotificationParametersUnsignedOutOfRange"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetNotificationParametersUnsignedOutOfRange")
	}

	// Create a partially initialized instance
	_child := &_BACnetNotificationParametersUnsignedOutOfRange{
		InnerOpeningTag: innerOpeningTag,
		ExceedingValue:  exceedingValue,
		StatusFlags:     statusFlags,
		Deadband:        deadband,
		ExceededLimit:   exceededLimit,
		InnerClosingTag: innerClosingTag,
		_BACnetNotificationParameters: &_BACnetNotificationParameters{
			TagNumber:          tagNumber,
			ObjectTypeArgument: objectTypeArgument,
		},
	}
	_child._BACnetNotificationParameters._BACnetNotificationParametersChildRequirements = _child
	return _child, nil
}

func (m *_BACnetNotificationParametersUnsignedOutOfRange) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetNotificationParametersUnsignedOutOfRange"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetNotificationParametersUnsignedOutOfRange")
		}

		// Simple Field (innerOpeningTag)
		if pushErr := writeBuffer.PushContext("innerOpeningTag"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for innerOpeningTag")
		}
		_innerOpeningTagErr := writeBuffer.WriteSerializable(m.GetInnerOpeningTag())
		if popErr := writeBuffer.PopContext("innerOpeningTag"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for innerOpeningTag")
		}
		if _innerOpeningTagErr != nil {
			return errors.Wrap(_innerOpeningTagErr, "Error serializing 'innerOpeningTag' field")
		}

		// Simple Field (exceedingValue)
		if pushErr := writeBuffer.PushContext("exceedingValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for exceedingValue")
		}
		_exceedingValueErr := writeBuffer.WriteSerializable(m.GetExceedingValue())
		if popErr := writeBuffer.PopContext("exceedingValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for exceedingValue")
		}
		if _exceedingValueErr != nil {
			return errors.Wrap(_exceedingValueErr, "Error serializing 'exceedingValue' field")
		}

		// Simple Field (statusFlags)
		if pushErr := writeBuffer.PushContext("statusFlags"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for statusFlags")
		}
		_statusFlagsErr := writeBuffer.WriteSerializable(m.GetStatusFlags())
		if popErr := writeBuffer.PopContext("statusFlags"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for statusFlags")
		}
		if _statusFlagsErr != nil {
			return errors.Wrap(_statusFlagsErr, "Error serializing 'statusFlags' field")
		}

		// Simple Field (deadband)
		if pushErr := writeBuffer.PushContext("deadband"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for deadband")
		}
		_deadbandErr := writeBuffer.WriteSerializable(m.GetDeadband())
		if popErr := writeBuffer.PopContext("deadband"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for deadband")
		}
		if _deadbandErr != nil {
			return errors.Wrap(_deadbandErr, "Error serializing 'deadband' field")
		}

		// Simple Field (exceededLimit)
		if pushErr := writeBuffer.PushContext("exceededLimit"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for exceededLimit")
		}
		_exceededLimitErr := writeBuffer.WriteSerializable(m.GetExceededLimit())
		if popErr := writeBuffer.PopContext("exceededLimit"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for exceededLimit")
		}
		if _exceededLimitErr != nil {
			return errors.Wrap(_exceededLimitErr, "Error serializing 'exceededLimit' field")
		}

		// Simple Field (innerClosingTag)
		if pushErr := writeBuffer.PushContext("innerClosingTag"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for innerClosingTag")
		}
		_innerClosingTagErr := writeBuffer.WriteSerializable(m.GetInnerClosingTag())
		if popErr := writeBuffer.PopContext("innerClosingTag"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for innerClosingTag")
		}
		if _innerClosingTagErr != nil {
			return errors.Wrap(_innerClosingTagErr, "Error serializing 'innerClosingTag' field")
		}

		if popErr := writeBuffer.PopContext("BACnetNotificationParametersUnsignedOutOfRange"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetNotificationParametersUnsignedOutOfRange")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetNotificationParametersUnsignedOutOfRange) isBACnetNotificationParametersUnsignedOutOfRange() bool {
	return true
}

func (m *_BACnetNotificationParametersUnsignedOutOfRange) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
