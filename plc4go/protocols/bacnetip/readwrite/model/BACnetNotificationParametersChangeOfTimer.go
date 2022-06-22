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

// BACnetNotificationParametersChangeOfTimer is the corresponding interface of BACnetNotificationParametersChangeOfTimer
type BACnetNotificationParametersChangeOfTimer interface {
	utils.LengthAware
	utils.Serializable
	BACnetNotificationParameters
	// GetInnerOpeningTag returns InnerOpeningTag (property field)
	GetInnerOpeningTag() BACnetOpeningTag
	// GetNewValue returns NewValue (property field)
	GetNewValue() BACnetTimerStateTagged
	// GetStatusFlags returns StatusFlags (property field)
	GetStatusFlags() BACnetStatusFlagsTagged
	// GetUpdateTime returns UpdateTime (property field)
	GetUpdateTime() BACnetDateTimeEnclosed
	// GetLastStateChange returns LastStateChange (property field)
	GetLastStateChange() BACnetTimerTransitionTagged
	// GetInitialTimeout returns InitialTimeout (property field)
	GetInitialTimeout() BACnetContextTagUnsignedInteger
	// GetExpirationTime returns ExpirationTime (property field)
	GetExpirationTime() BACnetDateTimeEnclosed
	// GetInnerClosingTag returns InnerClosingTag (property field)
	GetInnerClosingTag() BACnetClosingTag
}

// BACnetNotificationParametersChangeOfTimerExactly can be used when we want exactly this type and not a type which fulfills BACnetNotificationParametersChangeOfTimer.
// This is useful for switch cases.
type BACnetNotificationParametersChangeOfTimerExactly interface {
	BACnetNotificationParametersChangeOfTimer
	isBACnetNotificationParametersChangeOfTimer() bool
}

// _BACnetNotificationParametersChangeOfTimer is the data-structure of this message
type _BACnetNotificationParametersChangeOfTimer struct {
	*_BACnetNotificationParameters
	InnerOpeningTag BACnetOpeningTag
	NewValue        BACnetTimerStateTagged
	StatusFlags     BACnetStatusFlagsTagged
	UpdateTime      BACnetDateTimeEnclosed
	LastStateChange BACnetTimerTransitionTagged
	InitialTimeout  BACnetContextTagUnsignedInteger
	ExpirationTime  BACnetDateTimeEnclosed
	InnerClosingTag BACnetClosingTag

	// Arguments.
	TagNumber          uint8
	ObjectTypeArgument BACnetObjectType
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetNotificationParametersChangeOfTimer) InitializeParent(parent BACnetNotificationParameters, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetNotificationParametersChangeOfTimer) GetParent() BACnetNotificationParameters {
	return m._BACnetNotificationParameters
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetNotificationParametersChangeOfTimer) GetInnerOpeningTag() BACnetOpeningTag {
	return m.InnerOpeningTag
}

func (m *_BACnetNotificationParametersChangeOfTimer) GetNewValue() BACnetTimerStateTagged {
	return m.NewValue
}

func (m *_BACnetNotificationParametersChangeOfTimer) GetStatusFlags() BACnetStatusFlagsTagged {
	return m.StatusFlags
}

func (m *_BACnetNotificationParametersChangeOfTimer) GetUpdateTime() BACnetDateTimeEnclosed {
	return m.UpdateTime
}

func (m *_BACnetNotificationParametersChangeOfTimer) GetLastStateChange() BACnetTimerTransitionTagged {
	return m.LastStateChange
}

func (m *_BACnetNotificationParametersChangeOfTimer) GetInitialTimeout() BACnetContextTagUnsignedInteger {
	return m.InitialTimeout
}

func (m *_BACnetNotificationParametersChangeOfTimer) GetExpirationTime() BACnetDateTimeEnclosed {
	return m.ExpirationTime
}

func (m *_BACnetNotificationParametersChangeOfTimer) GetInnerClosingTag() BACnetClosingTag {
	return m.InnerClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetNotificationParametersChangeOfTimer factory function for _BACnetNotificationParametersChangeOfTimer
func NewBACnetNotificationParametersChangeOfTimer(innerOpeningTag BACnetOpeningTag, newValue BACnetTimerStateTagged, statusFlags BACnetStatusFlagsTagged, updateTime BACnetDateTimeEnclosed, lastStateChange BACnetTimerTransitionTagged, initialTimeout BACnetContextTagUnsignedInteger, expirationTime BACnetDateTimeEnclosed, innerClosingTag BACnetClosingTag, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, objectTypeArgument BACnetObjectType) *_BACnetNotificationParametersChangeOfTimer {
	_result := &_BACnetNotificationParametersChangeOfTimer{
		InnerOpeningTag:               innerOpeningTag,
		NewValue:                      newValue,
		StatusFlags:                   statusFlags,
		UpdateTime:                    updateTime,
		LastStateChange:               lastStateChange,
		InitialTimeout:                initialTimeout,
		ExpirationTime:                expirationTime,
		InnerClosingTag:               innerClosingTag,
		_BACnetNotificationParameters: NewBACnetNotificationParameters(openingTag, peekedTagHeader, closingTag, tagNumber, objectTypeArgument),
	}
	_result._BACnetNotificationParameters._BACnetNotificationParametersChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetNotificationParametersChangeOfTimer(structType interface{}) BACnetNotificationParametersChangeOfTimer {
	if casted, ok := structType.(BACnetNotificationParametersChangeOfTimer); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetNotificationParametersChangeOfTimer); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetNotificationParametersChangeOfTimer) GetTypeName() string {
	return "BACnetNotificationParametersChangeOfTimer"
}

func (m *_BACnetNotificationParametersChangeOfTimer) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetNotificationParametersChangeOfTimer) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (innerOpeningTag)
	lengthInBits += m.InnerOpeningTag.GetLengthInBits()

	// Simple field (newValue)
	lengthInBits += m.NewValue.GetLengthInBits()

	// Simple field (statusFlags)
	lengthInBits += m.StatusFlags.GetLengthInBits()

	// Simple field (updateTime)
	lengthInBits += m.UpdateTime.GetLengthInBits()

	// Optional Field (lastStateChange)
	if m.LastStateChange != nil {
		lengthInBits += m.LastStateChange.GetLengthInBits()
	}

	// Optional Field (initialTimeout)
	if m.InitialTimeout != nil {
		lengthInBits += m.InitialTimeout.GetLengthInBits()
	}

	// Optional Field (expirationTime)
	if m.ExpirationTime != nil {
		lengthInBits += m.ExpirationTime.GetLengthInBits()
	}

	// Simple field (innerClosingTag)
	lengthInBits += m.InnerClosingTag.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetNotificationParametersChangeOfTimer) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetNotificationParametersChangeOfTimerParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, peekedTagNumber uint8) (BACnetNotificationParametersChangeOfTimer, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetNotificationParametersChangeOfTimer"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetNotificationParametersChangeOfTimer")
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

	// Simple Field (newValue)
	if pullErr := readBuffer.PullContext("newValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for newValue")
	}
	_newValue, _newValueErr := BACnetTimerStateTaggedParse(readBuffer, uint8(uint8(0)), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _newValueErr != nil {
		return nil, errors.Wrap(_newValueErr, "Error parsing 'newValue' field")
	}
	newValue := _newValue.(BACnetTimerStateTagged)
	if closeErr := readBuffer.CloseContext("newValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for newValue")
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

	// Simple Field (updateTime)
	if pullErr := readBuffer.PullContext("updateTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for updateTime")
	}
	_updateTime, _updateTimeErr := BACnetDateTimeEnclosedParse(readBuffer, uint8(uint8(2)))
	if _updateTimeErr != nil {
		return nil, errors.Wrap(_updateTimeErr, "Error parsing 'updateTime' field")
	}
	updateTime := _updateTime.(BACnetDateTimeEnclosed)
	if closeErr := readBuffer.CloseContext("updateTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for updateTime")
	}

	// Optional Field (lastStateChange) (Can be skipped, if a given expression evaluates to false)
	var lastStateChange BACnetTimerTransitionTagged = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("lastStateChange"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for lastStateChange")
		}
		_val, _err := BACnetTimerTransitionTaggedParse(readBuffer, uint8(3), TagClass_CONTEXT_SPECIFIC_TAGS)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'lastStateChange' field")
		default:
			lastStateChange = _val.(BACnetTimerTransitionTagged)
			if closeErr := readBuffer.CloseContext("lastStateChange"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for lastStateChange")
			}
		}
	}

	// Optional Field (initialTimeout) (Can be skipped, if a given expression evaluates to false)
	var initialTimeout BACnetContextTagUnsignedInteger = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("initialTimeout"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for initialTimeout")
		}
		_val, _err := BACnetContextTagParse(readBuffer, uint8(4), BACnetDataType_UNSIGNED_INTEGER)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'initialTimeout' field")
		default:
			initialTimeout = _val.(BACnetContextTagUnsignedInteger)
			if closeErr := readBuffer.CloseContext("initialTimeout"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for initialTimeout")
			}
		}
	}

	// Optional Field (expirationTime) (Can be skipped, if a given expression evaluates to false)
	var expirationTime BACnetDateTimeEnclosed = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("expirationTime"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for expirationTime")
		}
		_val, _err := BACnetDateTimeEnclosedParse(readBuffer, uint8(5))
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'expirationTime' field")
		default:
			expirationTime = _val.(BACnetDateTimeEnclosed)
			if closeErr := readBuffer.CloseContext("expirationTime"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for expirationTime")
			}
		}
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

	if closeErr := readBuffer.CloseContext("BACnetNotificationParametersChangeOfTimer"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetNotificationParametersChangeOfTimer")
	}

	// Create a partially initialized instance
	_child := &_BACnetNotificationParametersChangeOfTimer{
		InnerOpeningTag:               innerOpeningTag,
		NewValue:                      newValue,
		StatusFlags:                   statusFlags,
		UpdateTime:                    updateTime,
		LastStateChange:               lastStateChange,
		InitialTimeout:                initialTimeout,
		ExpirationTime:                expirationTime,
		InnerClosingTag:               innerClosingTag,
		_BACnetNotificationParameters: &_BACnetNotificationParameters{},
	}
	_child._BACnetNotificationParameters._BACnetNotificationParametersChildRequirements = _child
	return _child, nil
}

func (m *_BACnetNotificationParametersChangeOfTimer) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetNotificationParametersChangeOfTimer"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetNotificationParametersChangeOfTimer")
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

		// Simple Field (newValue)
		if pushErr := writeBuffer.PushContext("newValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for newValue")
		}
		_newValueErr := writeBuffer.WriteSerializable(m.GetNewValue())
		if popErr := writeBuffer.PopContext("newValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for newValue")
		}
		if _newValueErr != nil {
			return errors.Wrap(_newValueErr, "Error serializing 'newValue' field")
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

		// Simple Field (updateTime)
		if pushErr := writeBuffer.PushContext("updateTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for updateTime")
		}
		_updateTimeErr := writeBuffer.WriteSerializable(m.GetUpdateTime())
		if popErr := writeBuffer.PopContext("updateTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for updateTime")
		}
		if _updateTimeErr != nil {
			return errors.Wrap(_updateTimeErr, "Error serializing 'updateTime' field")
		}

		// Optional Field (lastStateChange) (Can be skipped, if the value is null)
		var lastStateChange BACnetTimerTransitionTagged = nil
		if m.GetLastStateChange() != nil {
			if pushErr := writeBuffer.PushContext("lastStateChange"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for lastStateChange")
			}
			lastStateChange = m.GetLastStateChange()
			_lastStateChangeErr := writeBuffer.WriteSerializable(lastStateChange)
			if popErr := writeBuffer.PopContext("lastStateChange"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for lastStateChange")
			}
			if _lastStateChangeErr != nil {
				return errors.Wrap(_lastStateChangeErr, "Error serializing 'lastStateChange' field")
			}
		}

		// Optional Field (initialTimeout) (Can be skipped, if the value is null)
		var initialTimeout BACnetContextTagUnsignedInteger = nil
		if m.GetInitialTimeout() != nil {
			if pushErr := writeBuffer.PushContext("initialTimeout"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for initialTimeout")
			}
			initialTimeout = m.GetInitialTimeout()
			_initialTimeoutErr := writeBuffer.WriteSerializable(initialTimeout)
			if popErr := writeBuffer.PopContext("initialTimeout"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for initialTimeout")
			}
			if _initialTimeoutErr != nil {
				return errors.Wrap(_initialTimeoutErr, "Error serializing 'initialTimeout' field")
			}
		}

		// Optional Field (expirationTime) (Can be skipped, if the value is null)
		var expirationTime BACnetDateTimeEnclosed = nil
		if m.GetExpirationTime() != nil {
			if pushErr := writeBuffer.PushContext("expirationTime"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for expirationTime")
			}
			expirationTime = m.GetExpirationTime()
			_expirationTimeErr := writeBuffer.WriteSerializable(expirationTime)
			if popErr := writeBuffer.PopContext("expirationTime"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for expirationTime")
			}
			if _expirationTimeErr != nil {
				return errors.Wrap(_expirationTimeErr, "Error serializing 'expirationTime' field")
			}
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

		if popErr := writeBuffer.PopContext("BACnetNotificationParametersChangeOfTimer"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetNotificationParametersChangeOfTimer")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetNotificationParametersChangeOfTimer) isBACnetNotificationParametersChangeOfTimer() bool {
	return true
}

func (m *_BACnetNotificationParametersChangeOfTimer) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
