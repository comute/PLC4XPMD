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

// BACnetEventNotificationSubscription is the data-structure of this message
type BACnetEventNotificationSubscription struct {
	Recipient                   *BACnetRecipientEnclosed
	ProcessIdentifier           *BACnetContextTagUnsignedInteger
	IssueConfirmedNotifications *BACnetContextTagBoolean
	TimeRemaining               *BACnetContextTagUnsignedInteger
}

// IBACnetEventNotificationSubscription is the corresponding interface of BACnetEventNotificationSubscription
type IBACnetEventNotificationSubscription interface {
	// GetRecipient returns Recipient (property field)
	GetRecipient() *BACnetRecipientEnclosed
	// GetProcessIdentifier returns ProcessIdentifier (property field)
	GetProcessIdentifier() *BACnetContextTagUnsignedInteger
	// GetIssueConfirmedNotifications returns IssueConfirmedNotifications (property field)
	GetIssueConfirmedNotifications() *BACnetContextTagBoolean
	// GetTimeRemaining returns TimeRemaining (property field)
	GetTimeRemaining() *BACnetContextTagUnsignedInteger
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

func (m *BACnetEventNotificationSubscription) GetRecipient() *BACnetRecipientEnclosed {
	return m.Recipient
}

func (m *BACnetEventNotificationSubscription) GetProcessIdentifier() *BACnetContextTagUnsignedInteger {
	return m.ProcessIdentifier
}

func (m *BACnetEventNotificationSubscription) GetIssueConfirmedNotifications() *BACnetContextTagBoolean {
	return m.IssueConfirmedNotifications
}

func (m *BACnetEventNotificationSubscription) GetTimeRemaining() *BACnetContextTagUnsignedInteger {
	return m.TimeRemaining
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetEventNotificationSubscription factory function for BACnetEventNotificationSubscription
func NewBACnetEventNotificationSubscription(recipient *BACnetRecipientEnclosed, processIdentifier *BACnetContextTagUnsignedInteger, issueConfirmedNotifications *BACnetContextTagBoolean, timeRemaining *BACnetContextTagUnsignedInteger) *BACnetEventNotificationSubscription {
	return &BACnetEventNotificationSubscription{Recipient: recipient, ProcessIdentifier: processIdentifier, IssueConfirmedNotifications: issueConfirmedNotifications, TimeRemaining: timeRemaining}
}

func CastBACnetEventNotificationSubscription(structType interface{}) *BACnetEventNotificationSubscription {
	if casted, ok := structType.(BACnetEventNotificationSubscription); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetEventNotificationSubscription); ok {
		return casted
	}
	return nil
}

func (m *BACnetEventNotificationSubscription) GetTypeName() string {
	return "BACnetEventNotificationSubscription"
}

func (m *BACnetEventNotificationSubscription) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetEventNotificationSubscription) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (recipient)
	lengthInBits += m.Recipient.GetLengthInBits()

	// Simple field (processIdentifier)
	lengthInBits += m.ProcessIdentifier.GetLengthInBits()

	// Optional Field (issueConfirmedNotifications)
	if m.IssueConfirmedNotifications != nil {
		lengthInBits += (*m.IssueConfirmedNotifications).GetLengthInBits()
	}

	// Simple field (timeRemaining)
	lengthInBits += m.TimeRemaining.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetEventNotificationSubscription) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetEventNotificationSubscriptionParse(readBuffer utils.ReadBuffer) (*BACnetEventNotificationSubscription, error) {
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
	recipient := CastBACnetRecipientEnclosed(_recipient)
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
	processIdentifier := CastBACnetContextTagUnsignedInteger(_processIdentifier)
	if closeErr := readBuffer.CloseContext("processIdentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for processIdentifier")
	}

	// Optional Field (issueConfirmedNotifications) (Can be skipped, if a given expression evaluates to false)
	var issueConfirmedNotifications *BACnetContextTagBoolean = nil
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
			issueConfirmedNotifications = CastBACnetContextTagBoolean(_val)
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
	timeRemaining := CastBACnetContextTagUnsignedInteger(_timeRemaining)
	if closeErr := readBuffer.CloseContext("timeRemaining"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for timeRemaining")
	}

	if closeErr := readBuffer.CloseContext("BACnetEventNotificationSubscription"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetEventNotificationSubscription")
	}

	// Create the instance
	return NewBACnetEventNotificationSubscription(recipient, processIdentifier, issueConfirmedNotifications, timeRemaining), nil
}

func (m *BACnetEventNotificationSubscription) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetEventNotificationSubscription"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetEventNotificationSubscription")
	}

	// Simple Field (recipient)
	if pushErr := writeBuffer.PushContext("recipient"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for recipient")
	}
	_recipientErr := writeBuffer.WriteSerializable(m.Recipient)
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
	_processIdentifierErr := writeBuffer.WriteSerializable(m.ProcessIdentifier)
	if popErr := writeBuffer.PopContext("processIdentifier"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for processIdentifier")
	}
	if _processIdentifierErr != nil {
		return errors.Wrap(_processIdentifierErr, "Error serializing 'processIdentifier' field")
	}

	// Optional Field (issueConfirmedNotifications) (Can be skipped, if the value is null)
	var issueConfirmedNotifications *BACnetContextTagBoolean = nil
	if m.IssueConfirmedNotifications != nil {
		if pushErr := writeBuffer.PushContext("issueConfirmedNotifications"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for issueConfirmedNotifications")
		}
		issueConfirmedNotifications = m.IssueConfirmedNotifications
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
	_timeRemainingErr := writeBuffer.WriteSerializable(m.TimeRemaining)
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

func (m *BACnetEventNotificationSubscription) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
