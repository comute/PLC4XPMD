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

// BACnetEventNotificationSubscription is the corresponding interface of BACnetEventNotificationSubscription
type BACnetEventNotificationSubscription interface {
	utils.LengthAware
	utils.Serializable
	// GetRecipient returns Recipient (property field)
	GetRecipient() BACnetRecipientEnclosed
	// GetProcessIdentifier returns ProcessIdentifier (property field)
	GetProcessIdentifier() BACnetContextTagUnsignedInteger
	// GetIssueConfirmedNotifications returns IssueConfirmedNotifications (property field)
	GetIssueConfirmedNotifications() BACnetContextTagBoolean
	// GetTimeRemaining returns TimeRemaining (property field)
	GetTimeRemaining() BACnetContextTagUnsignedInteger
}

// BACnetEventNotificationSubscriptionExactly can be used when we want exactly this type and not a type which fulfills BACnetEventNotificationSubscription.
// This is useful for switch cases.
type BACnetEventNotificationSubscriptionExactly interface {
	BACnetEventNotificationSubscription
	isBACnetEventNotificationSubscription() bool
}

// _BACnetEventNotificationSubscription is the data-structure of this message
type _BACnetEventNotificationSubscription struct {
	Recipient                   BACnetRecipientEnclosed
	ProcessIdentifier           BACnetContextTagUnsignedInteger
	IssueConfirmedNotifications BACnetContextTagBoolean
	TimeRemaining               BACnetContextTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetEventNotificationSubscription) GetRecipient() BACnetRecipientEnclosed {
	return m.Recipient
}

func (m *_BACnetEventNotificationSubscription) GetProcessIdentifier() BACnetContextTagUnsignedInteger {
	return m.ProcessIdentifier
}

func (m *_BACnetEventNotificationSubscription) GetIssueConfirmedNotifications() BACnetContextTagBoolean {
	return m.IssueConfirmedNotifications
}

func (m *_BACnetEventNotificationSubscription) GetTimeRemaining() BACnetContextTagUnsignedInteger {
	return m.TimeRemaining
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetEventNotificationSubscription factory function for _BACnetEventNotificationSubscription
func NewBACnetEventNotificationSubscription(recipient BACnetRecipientEnclosed, processIdentifier BACnetContextTagUnsignedInteger, issueConfirmedNotifications BACnetContextTagBoolean, timeRemaining BACnetContextTagUnsignedInteger) *_BACnetEventNotificationSubscription {
	return &_BACnetEventNotificationSubscription{Recipient: recipient, ProcessIdentifier: processIdentifier, IssueConfirmedNotifications: issueConfirmedNotifications, TimeRemaining: timeRemaining}
}

// Deprecated: use the interface for direct cast
func CastBACnetEventNotificationSubscription(structType interface{}) BACnetEventNotificationSubscription {
	if casted, ok := structType.(BACnetEventNotificationSubscription); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetEventNotificationSubscription); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetEventNotificationSubscription) GetTypeName() string {
	return "BACnetEventNotificationSubscription"
}

func (m *_BACnetEventNotificationSubscription) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetEventNotificationSubscription) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (recipient)
	lengthInBits += m.Recipient.GetLengthInBits()

	// Simple field (processIdentifier)
	lengthInBits += m.ProcessIdentifier.GetLengthInBits()

	// Optional Field (issueConfirmedNotifications)
	if m.IssueConfirmedNotifications != nil {
		lengthInBits += m.IssueConfirmedNotifications.GetLengthInBits()
	}

	// Simple field (timeRemaining)
	lengthInBits += m.TimeRemaining.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetEventNotificationSubscription) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetEventNotificationSubscriptionParse(readBuffer utils.ReadBuffer) (BACnetEventNotificationSubscription, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetEventNotificationSubscription"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetEventNotificationSubscription")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (recipient)
	if pullErr := readBuffer.PullContext("recipient"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for recipient")
	}
	_recipient, _recipientErr := BACnetRecipientEnclosedParse(readBuffer, uint8(uint8(0)))
	if _recipientErr != nil {
		return nil, errors.Wrap(_recipientErr, "Error parsing 'recipient' field")
	}
	recipient := _recipient.(BACnetRecipientEnclosed)
	if closeErr := readBuffer.CloseContext("recipient"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for recipient")
	}

	// Simple Field (processIdentifier)
	if pullErr := readBuffer.PullContext("processIdentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for processIdentifier")
	}
	_processIdentifier, _processIdentifierErr := BACnetContextTagParse(readBuffer, uint8(uint8(1)), BACnetDataType(BACnetDataType_UNSIGNED_INTEGER))
	if _processIdentifierErr != nil {
		return nil, errors.Wrap(_processIdentifierErr, "Error parsing 'processIdentifier' field")
	}
	processIdentifier := _processIdentifier.(BACnetContextTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("processIdentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for processIdentifier")
	}

	// Optional Field (issueConfirmedNotifications) (Can be skipped, if a given expression evaluates to false)
	var issueConfirmedNotifications BACnetContextTagBoolean = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("issueConfirmedNotifications"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for issueConfirmedNotifications")
		}
		_val, _err := BACnetContextTagParse(readBuffer, uint8(2), BACnetDataType_BOOLEAN)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'issueConfirmedNotifications' field")
		default:
			issueConfirmedNotifications = _val.(BACnetContextTagBoolean)
			if closeErr := readBuffer.CloseContext("issueConfirmedNotifications"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for issueConfirmedNotifications")
			}
		}
	}

	// Simple Field (timeRemaining)
	if pullErr := readBuffer.PullContext("timeRemaining"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for timeRemaining")
	}
	_timeRemaining, _timeRemainingErr := BACnetContextTagParse(readBuffer, uint8(uint8(3)), BACnetDataType(BACnetDataType_UNSIGNED_INTEGER))
	if _timeRemainingErr != nil {
		return nil, errors.Wrap(_timeRemainingErr, "Error parsing 'timeRemaining' field")
	}
	timeRemaining := _timeRemaining.(BACnetContextTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("timeRemaining"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for timeRemaining")
	}

	if closeErr := readBuffer.CloseContext("BACnetEventNotificationSubscription"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetEventNotificationSubscription")
	}

	// Create the instance
	return NewBACnetEventNotificationSubscription(recipient, processIdentifier, issueConfirmedNotifications, timeRemaining), nil
}

func (m *_BACnetEventNotificationSubscription) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetEventNotificationSubscription"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetEventNotificationSubscription")
	}

	// Simple Field (recipient)
	if pushErr := writeBuffer.PushContext("recipient"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for recipient")
	}
	_recipientErr := writeBuffer.WriteSerializable(m.GetRecipient())
	if popErr := writeBuffer.PopContext("recipient"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for recipient")
	}
	if _recipientErr != nil {
		return errors.Wrap(_recipientErr, "Error serializing 'recipient' field")
	}

	// Simple Field (processIdentifier)
	if pushErr := writeBuffer.PushContext("processIdentifier"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for processIdentifier")
	}
	_processIdentifierErr := writeBuffer.WriteSerializable(m.GetProcessIdentifier())
	if popErr := writeBuffer.PopContext("processIdentifier"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for processIdentifier")
	}
	if _processIdentifierErr != nil {
		return errors.Wrap(_processIdentifierErr, "Error serializing 'processIdentifier' field")
	}

	// Optional Field (issueConfirmedNotifications) (Can be skipped, if the value is null)
	var issueConfirmedNotifications BACnetContextTagBoolean = nil
	if m.GetIssueConfirmedNotifications() != nil {
		if pushErr := writeBuffer.PushContext("issueConfirmedNotifications"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for issueConfirmedNotifications")
		}
		issueConfirmedNotifications = m.GetIssueConfirmedNotifications()
		_issueConfirmedNotificationsErr := writeBuffer.WriteSerializable(issueConfirmedNotifications)
		if popErr := writeBuffer.PopContext("issueConfirmedNotifications"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for issueConfirmedNotifications")
		}
		if _issueConfirmedNotificationsErr != nil {
			return errors.Wrap(_issueConfirmedNotificationsErr, "Error serializing 'issueConfirmedNotifications' field")
		}
	}

	// Simple Field (timeRemaining)
	if pushErr := writeBuffer.PushContext("timeRemaining"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for timeRemaining")
	}
	_timeRemainingErr := writeBuffer.WriteSerializable(m.GetTimeRemaining())
	if popErr := writeBuffer.PopContext("timeRemaining"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for timeRemaining")
	}
	if _timeRemainingErr != nil {
		return errors.Wrap(_timeRemainingErr, "Error serializing 'timeRemaining' field")
	}

	if popErr := writeBuffer.PopContext("BACnetEventNotificationSubscription"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetEventNotificationSubscription")
	}
	return nil
}

func (m *_BACnetEventNotificationSubscription) isBACnetEventNotificationSubscription() bool {
	return true
}

func (m *_BACnetEventNotificationSubscription) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
