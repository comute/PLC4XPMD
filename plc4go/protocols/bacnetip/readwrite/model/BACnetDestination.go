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

// BACnetDestination is the corresponding interface of BACnetDestination
type BACnetDestination interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetValidDays returns ValidDays (property field)
	GetValidDays() BACnetDaysOfWeekTagged
	// GetFromTime returns FromTime (property field)
	GetFromTime() BACnetApplicationTagTime
	// GetToTime returns ToTime (property field)
	GetToTime() BACnetApplicationTagTime
	// GetRecipient returns Recipient (property field)
	GetRecipient() BACnetRecipient
	// GetProcessIdentifier returns ProcessIdentifier (property field)
	GetProcessIdentifier() BACnetApplicationTagUnsignedInteger
	// GetIssueConfirmedNotifications returns IssueConfirmedNotifications (property field)
	GetIssueConfirmedNotifications() BACnetApplicationTagBoolean
	// GetTransitions returns Transitions (property field)
	GetTransitions() BACnetEventTransitionBitsTagged
}

// BACnetDestinationExactly can be used when we want exactly this type and not a type which fulfills BACnetDestination.
// This is useful for switch cases.
type BACnetDestinationExactly interface {
	BACnetDestination
	isBACnetDestination() bool
}

// _BACnetDestination is the data-structure of this message
type _BACnetDestination struct {
	ValidDays                   BACnetDaysOfWeekTagged
	FromTime                    BACnetApplicationTagTime
	ToTime                      BACnetApplicationTagTime
	Recipient                   BACnetRecipient
	ProcessIdentifier           BACnetApplicationTagUnsignedInteger
	IssueConfirmedNotifications BACnetApplicationTagBoolean
	Transitions                 BACnetEventTransitionBitsTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetDestination) GetValidDays() BACnetDaysOfWeekTagged {
	return m.ValidDays
}

func (m *_BACnetDestination) GetFromTime() BACnetApplicationTagTime {
	return m.FromTime
}

func (m *_BACnetDestination) GetToTime() BACnetApplicationTagTime {
	return m.ToTime
}

func (m *_BACnetDestination) GetRecipient() BACnetRecipient {
	return m.Recipient
}

func (m *_BACnetDestination) GetProcessIdentifier() BACnetApplicationTagUnsignedInteger {
	return m.ProcessIdentifier
}

func (m *_BACnetDestination) GetIssueConfirmedNotifications() BACnetApplicationTagBoolean {
	return m.IssueConfirmedNotifications
}

func (m *_BACnetDestination) GetTransitions() BACnetEventTransitionBitsTagged {
	return m.Transitions
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetDestination factory function for _BACnetDestination
func NewBACnetDestination(validDays BACnetDaysOfWeekTagged, fromTime BACnetApplicationTagTime, toTime BACnetApplicationTagTime, recipient BACnetRecipient, processIdentifier BACnetApplicationTagUnsignedInteger, issueConfirmedNotifications BACnetApplicationTagBoolean, transitions BACnetEventTransitionBitsTagged) *_BACnetDestination {
	return &_BACnetDestination{ValidDays: validDays, FromTime: fromTime, ToTime: toTime, Recipient: recipient, ProcessIdentifier: processIdentifier, IssueConfirmedNotifications: issueConfirmedNotifications, Transitions: transitions}
}

// Deprecated: use the interface for direct cast
func CastBACnetDestination(structType any) BACnetDestination {
	if casted, ok := structType.(BACnetDestination); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetDestination); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetDestination) GetTypeName() string {
	return "BACnetDestination"
}

func (m *_BACnetDestination) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (validDays)
	lengthInBits += m.ValidDays.GetLengthInBits(ctx)

	// Simple field (fromTime)
	lengthInBits += m.FromTime.GetLengthInBits(ctx)

	// Simple field (toTime)
	lengthInBits += m.ToTime.GetLengthInBits(ctx)

	// Simple field (recipient)
	lengthInBits += m.Recipient.GetLengthInBits(ctx)

	// Simple field (processIdentifier)
	lengthInBits += m.ProcessIdentifier.GetLengthInBits(ctx)

	// Simple field (issueConfirmedNotifications)
	lengthInBits += m.IssueConfirmedNotifications.GetLengthInBits(ctx)

	// Simple field (transitions)
	lengthInBits += m.Transitions.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetDestination) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetDestinationParse(ctx context.Context, theBytes []byte) (BACnetDestination, error) {
	return BACnetDestinationParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetDestinationParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetDestination, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetDestination, error) {
		return BACnetDestinationParseWithBuffer(ctx, readBuffer)
	}
}

func BACnetDestinationParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetDestination, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetDestination"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetDestination")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	validDays, err := ReadSimpleField[BACnetDaysOfWeekTagged](ctx, "validDays", ReadComplex[BACnetDaysOfWeekTagged](BACnetDaysOfWeekTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'validDays' field"))
	}

	fromTime, err := ReadSimpleField[BACnetApplicationTagTime](ctx, "fromTime", ReadComplex[BACnetApplicationTagTime](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagTime](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'fromTime' field"))
	}

	toTime, err := ReadSimpleField[BACnetApplicationTagTime](ctx, "toTime", ReadComplex[BACnetApplicationTagTime](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagTime](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'toTime' field"))
	}

	recipient, err := ReadSimpleField[BACnetRecipient](ctx, "recipient", ReadComplex[BACnetRecipient](BACnetRecipientParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'recipient' field"))
	}

	processIdentifier, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "processIdentifier", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'processIdentifier' field"))
	}

	issueConfirmedNotifications, err := ReadSimpleField[BACnetApplicationTagBoolean](ctx, "issueConfirmedNotifications", ReadComplex[BACnetApplicationTagBoolean](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagBoolean](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'issueConfirmedNotifications' field"))
	}

	transitions, err := ReadSimpleField[BACnetEventTransitionBitsTagged](ctx, "transitions", ReadComplex[BACnetEventTransitionBitsTagged](BACnetEventTransitionBitsTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'transitions' field"))
	}

	if closeErr := readBuffer.CloseContext("BACnetDestination"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetDestination")
	}

	// Create the instance
	return &_BACnetDestination{
		ValidDays:                   validDays,
		FromTime:                    fromTime,
		ToTime:                      toTime,
		Recipient:                   recipient,
		ProcessIdentifier:           processIdentifier,
		IssueConfirmedNotifications: issueConfirmedNotifications,
		Transitions:                 transitions,
	}, nil
}

func (m *_BACnetDestination) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetDestination) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetDestination"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetDestination")
	}

	// Simple Field (validDays)
	if pushErr := writeBuffer.PushContext("validDays"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for validDays")
	}
	_validDaysErr := writeBuffer.WriteSerializable(ctx, m.GetValidDays())
	if popErr := writeBuffer.PopContext("validDays"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for validDays")
	}
	if _validDaysErr != nil {
		return errors.Wrap(_validDaysErr, "Error serializing 'validDays' field")
	}

	// Simple Field (fromTime)
	if pushErr := writeBuffer.PushContext("fromTime"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for fromTime")
	}
	_fromTimeErr := writeBuffer.WriteSerializable(ctx, m.GetFromTime())
	if popErr := writeBuffer.PopContext("fromTime"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for fromTime")
	}
	if _fromTimeErr != nil {
		return errors.Wrap(_fromTimeErr, "Error serializing 'fromTime' field")
	}

	// Simple Field (toTime)
	if pushErr := writeBuffer.PushContext("toTime"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for toTime")
	}
	_toTimeErr := writeBuffer.WriteSerializable(ctx, m.GetToTime())
	if popErr := writeBuffer.PopContext("toTime"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for toTime")
	}
	if _toTimeErr != nil {
		return errors.Wrap(_toTimeErr, "Error serializing 'toTime' field")
	}

	// Simple Field (recipient)
	if pushErr := writeBuffer.PushContext("recipient"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for recipient")
	}
	_recipientErr := writeBuffer.WriteSerializable(ctx, m.GetRecipient())
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
	_processIdentifierErr := writeBuffer.WriteSerializable(ctx, m.GetProcessIdentifier())
	if popErr := writeBuffer.PopContext("processIdentifier"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for processIdentifier")
	}
	if _processIdentifierErr != nil {
		return errors.Wrap(_processIdentifierErr, "Error serializing 'processIdentifier' field")
	}

	// Simple Field (issueConfirmedNotifications)
	if pushErr := writeBuffer.PushContext("issueConfirmedNotifications"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for issueConfirmedNotifications")
	}
	_issueConfirmedNotificationsErr := writeBuffer.WriteSerializable(ctx, m.GetIssueConfirmedNotifications())
	if popErr := writeBuffer.PopContext("issueConfirmedNotifications"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for issueConfirmedNotifications")
	}
	if _issueConfirmedNotificationsErr != nil {
		return errors.Wrap(_issueConfirmedNotificationsErr, "Error serializing 'issueConfirmedNotifications' field")
	}

	// Simple Field (transitions)
	if pushErr := writeBuffer.PushContext("transitions"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for transitions")
	}
	_transitionsErr := writeBuffer.WriteSerializable(ctx, m.GetTransitions())
	if popErr := writeBuffer.PopContext("transitions"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for transitions")
	}
	if _transitionsErr != nil {
		return errors.Wrap(_transitionsErr, "Error serializing 'transitions' field")
	}

	if popErr := writeBuffer.PopContext("BACnetDestination"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetDestination")
	}
	return nil
}

func (m *_BACnetDestination) isBACnetDestination() bool {
	return true
}

func (m *_BACnetDestination) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
