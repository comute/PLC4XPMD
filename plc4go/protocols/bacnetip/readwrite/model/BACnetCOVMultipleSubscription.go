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

// BACnetCOVMultipleSubscription is the corresponding interface of BACnetCOVMultipleSubscription
type BACnetCOVMultipleSubscription interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetRecipient returns Recipient (property field)
	GetRecipient() BACnetRecipientProcessEnclosed
	// GetIssueConfirmedNotifications returns IssueConfirmedNotifications (property field)
	GetIssueConfirmedNotifications() BACnetContextTagBoolean
	// GetTimeRemaining returns TimeRemaining (property field)
	GetTimeRemaining() BACnetContextTagUnsignedInteger
	// GetMaxNotificationDelay returns MaxNotificationDelay (property field)
	GetMaxNotificationDelay() BACnetContextTagUnsignedInteger
	// GetListOfCovSubscriptionSpecification returns ListOfCovSubscriptionSpecification (property field)
	GetListOfCovSubscriptionSpecification() BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification
}

// BACnetCOVMultipleSubscriptionExactly can be used when we want exactly this type and not a type which fulfills BACnetCOVMultipleSubscription.
// This is useful for switch cases.
type BACnetCOVMultipleSubscriptionExactly interface {
	BACnetCOVMultipleSubscription
	isBACnetCOVMultipleSubscription() bool
}

// _BACnetCOVMultipleSubscription is the data-structure of this message
type _BACnetCOVMultipleSubscription struct {
	Recipient                          BACnetRecipientProcessEnclosed
	IssueConfirmedNotifications        BACnetContextTagBoolean
	TimeRemaining                      BACnetContextTagUnsignedInteger
	MaxNotificationDelay               BACnetContextTagUnsignedInteger
	ListOfCovSubscriptionSpecification BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetCOVMultipleSubscription) GetRecipient() BACnetRecipientProcessEnclosed {
	return m.Recipient
}

func (m *_BACnetCOVMultipleSubscription) GetIssueConfirmedNotifications() BACnetContextTagBoolean {
	return m.IssueConfirmedNotifications
}

func (m *_BACnetCOVMultipleSubscription) GetTimeRemaining() BACnetContextTagUnsignedInteger {
	return m.TimeRemaining
}

func (m *_BACnetCOVMultipleSubscription) GetMaxNotificationDelay() BACnetContextTagUnsignedInteger {
	return m.MaxNotificationDelay
}

func (m *_BACnetCOVMultipleSubscription) GetListOfCovSubscriptionSpecification() BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification {
	return m.ListOfCovSubscriptionSpecification
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetCOVMultipleSubscription factory function for _BACnetCOVMultipleSubscription
func NewBACnetCOVMultipleSubscription(recipient BACnetRecipientProcessEnclosed, issueConfirmedNotifications BACnetContextTagBoolean, timeRemaining BACnetContextTagUnsignedInteger, maxNotificationDelay BACnetContextTagUnsignedInteger, listOfCovSubscriptionSpecification BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification) *_BACnetCOVMultipleSubscription {
	return &_BACnetCOVMultipleSubscription{Recipient: recipient, IssueConfirmedNotifications: issueConfirmedNotifications, TimeRemaining: timeRemaining, MaxNotificationDelay: maxNotificationDelay, ListOfCovSubscriptionSpecification: listOfCovSubscriptionSpecification}
}

// Deprecated: use the interface for direct cast
func CastBACnetCOVMultipleSubscription(structType any) BACnetCOVMultipleSubscription {
	if casted, ok := structType.(BACnetCOVMultipleSubscription); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetCOVMultipleSubscription); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetCOVMultipleSubscription) GetTypeName() string {
	return "BACnetCOVMultipleSubscription"
}

func (m *_BACnetCOVMultipleSubscription) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (recipient)
	lengthInBits += m.Recipient.GetLengthInBits(ctx)

	// Simple field (issueConfirmedNotifications)
	lengthInBits += m.IssueConfirmedNotifications.GetLengthInBits(ctx)

	// Simple field (timeRemaining)
	lengthInBits += m.TimeRemaining.GetLengthInBits(ctx)

	// Simple field (maxNotificationDelay)
	lengthInBits += m.MaxNotificationDelay.GetLengthInBits(ctx)

	// Simple field (listOfCovSubscriptionSpecification)
	lengthInBits += m.ListOfCovSubscriptionSpecification.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetCOVMultipleSubscription) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetCOVMultipleSubscriptionParse(ctx context.Context, theBytes []byte) (BACnetCOVMultipleSubscription, error) {
	return BACnetCOVMultipleSubscriptionParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetCOVMultipleSubscriptionParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetCOVMultipleSubscription, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetCOVMultipleSubscription, error) {
		return BACnetCOVMultipleSubscriptionParseWithBuffer(ctx, readBuffer)
	}
}

func BACnetCOVMultipleSubscriptionParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetCOVMultipleSubscription, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetCOVMultipleSubscription"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetCOVMultipleSubscription")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	recipient, err := ReadSimpleField[BACnetRecipientProcessEnclosed](ctx, "recipient", ReadComplex[BACnetRecipientProcessEnclosed](BACnetRecipientProcessEnclosedParseWithBufferProducer((uint8)(uint8(0))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'recipient' field"))
	}

	issueConfirmedNotifications, err := ReadSimpleField[BACnetContextTagBoolean](ctx, "issueConfirmedNotifications", ReadComplex[BACnetContextTagBoolean](BACnetContextTagParseWithBufferProducer[BACnetContextTagBoolean]((uint8)(uint8(1)), (BACnetDataType)(BACnetDataType_BOOLEAN)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'issueConfirmedNotifications' field"))
	}

	timeRemaining, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "timeRemaining", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(2)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timeRemaining' field"))
	}

	maxNotificationDelay, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "maxNotificationDelay", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(3)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'maxNotificationDelay' field"))
	}

	listOfCovSubscriptionSpecification, err := ReadSimpleField[BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification](ctx, "listOfCovSubscriptionSpecification", ReadComplex[BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification](BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationParseWithBufferProducer((uint8)(uint8(4))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'listOfCovSubscriptionSpecification' field"))
	}

	if closeErr := readBuffer.CloseContext("BACnetCOVMultipleSubscription"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetCOVMultipleSubscription")
	}

	// Create the instance
	return &_BACnetCOVMultipleSubscription{
		Recipient:                          recipient,
		IssueConfirmedNotifications:        issueConfirmedNotifications,
		TimeRemaining:                      timeRemaining,
		MaxNotificationDelay:               maxNotificationDelay,
		ListOfCovSubscriptionSpecification: listOfCovSubscriptionSpecification,
	}, nil
}

func (m *_BACnetCOVMultipleSubscription) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetCOVMultipleSubscription) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetCOVMultipleSubscription"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetCOVMultipleSubscription")
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

	// Simple Field (timeRemaining)
	if pushErr := writeBuffer.PushContext("timeRemaining"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for timeRemaining")
	}
	_timeRemainingErr := writeBuffer.WriteSerializable(ctx, m.GetTimeRemaining())
	if popErr := writeBuffer.PopContext("timeRemaining"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for timeRemaining")
	}
	if _timeRemainingErr != nil {
		return errors.Wrap(_timeRemainingErr, "Error serializing 'timeRemaining' field")
	}

	// Simple Field (maxNotificationDelay)
	if pushErr := writeBuffer.PushContext("maxNotificationDelay"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for maxNotificationDelay")
	}
	_maxNotificationDelayErr := writeBuffer.WriteSerializable(ctx, m.GetMaxNotificationDelay())
	if popErr := writeBuffer.PopContext("maxNotificationDelay"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for maxNotificationDelay")
	}
	if _maxNotificationDelayErr != nil {
		return errors.Wrap(_maxNotificationDelayErr, "Error serializing 'maxNotificationDelay' field")
	}

	// Simple Field (listOfCovSubscriptionSpecification)
	if pushErr := writeBuffer.PushContext("listOfCovSubscriptionSpecification"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for listOfCovSubscriptionSpecification")
	}
	_listOfCovSubscriptionSpecificationErr := writeBuffer.WriteSerializable(ctx, m.GetListOfCovSubscriptionSpecification())
	if popErr := writeBuffer.PopContext("listOfCovSubscriptionSpecification"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for listOfCovSubscriptionSpecification")
	}
	if _listOfCovSubscriptionSpecificationErr != nil {
		return errors.Wrap(_listOfCovSubscriptionSpecificationErr, "Error serializing 'listOfCovSubscriptionSpecification' field")
	}

	if popErr := writeBuffer.PopContext("BACnetCOVMultipleSubscription"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetCOVMultipleSubscription")
	}
	return nil
}

func (m *_BACnetCOVMultipleSubscription) isBACnetCOVMultipleSubscription() bool {
	return true
}

func (m *_BACnetCOVMultipleSubscription) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
